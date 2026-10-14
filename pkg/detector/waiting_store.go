/*
Copyright 2026 The Karmada Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package detector

import (
	"sync"
	"sync/atomic"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"

	policyv1alpha1 "github.com/karmada-io/karmada/pkg/apis/policy/v1alpha1"
	"github.com/karmada-io/karmada/pkg/util/fedinformer/keys"
)

// waitingObjectStore keeps the resource identity and labels needed by ResourceSelector matching.
// Its zero value is ready for use.
type waitingObjectStore struct {
	count atomic.Int64
	// cache stores the namespaces seen for each GVK. GVK and namespace
	// maps are intentionally retained even when empty, so readers and writers
	// can use stable map pointers without locks or delete/recreate races. The
	// number of observed GVK/namespace pairs is expected to remain below 1,000
	// in normal operation; even 1,000 pairs retain only about 600 KiB of map
	// structure after their objects are removed.
	// structure: GVK -> *sync.Map(namespace -> *sync.Map(ClusterWideKey -> labels))
	cache sync.Map
}

type waitingObjectSelector struct {
	gvk           schema.GroupVersionKind
	namespace     string
	name          string
	labelSelector labels.Selector
}

// upsert takes ownership of objectLabels. The caller must not mutate it after the call.
func (s *waitingObjectStore) upsert(key keys.ClusterWideKey, objectLabels map[string]string) (bool, int) {
	objects := s.getOrCreateObjectsBucket(key.GroupVersionKind(), key.Namespace)
	if _, exists := objects.Swap(key, objectLabels); exists {
		return false, int(s.count.Load())
	}
	return true, int(s.count.Add(1))
}

func (s *waitingObjectStore) delete(key keys.ClusterWideKey) {
	objects := s.getObjectsBucket(key.GroupVersionKind(), key.Namespace)
	if objects == nil {
		return
	}
	if _, exists := objects.LoadAndDelete(key); !exists {
		return
	}
	s.count.Add(-1)
}

func (s *waitingObjectStore) len() int {
	return int(s.count.Load())
}

func (s *waitingObjectStore) match(resourceSelectors []policyv1alpha1.ResourceSelector) []keys.ClusterWideKey {
	selectors := compileWaitingObjectSelectors(resourceSelectors)
	matched := sets.New[keys.ClusterWideKey]()

	for _, selector := range selectors {
		if selector.name != "" && selector.namespace != "" {
			key := waitingObjectKey(selector)
			if objects := s.getObjectsBucket(selector.gvk, selector.namespace); objects != nil {
				if _, exists := objects.Load(key); exists {
					matched.Insert(key)
				}
			}
			continue
		}

		if selector.namespace != "" {
			matchObjects(s.getObjectsBucket(selector.gvk, selector.namespace), selector, matched)
			continue
		}

		namespaces, exists := s.cache.Load(selector.gvk)
		if !exists {
			continue
		}

		namespaces.(*sync.Map).Range(func(_, objects any) bool {
			matchObjects(objects.(*sync.Map), selector, matched)
			return true
		})
	}

	return matched.UnsortedList()
}

func matchObjects(
	objects *sync.Map,
	selector waitingObjectSelector,
	matched sets.Set[keys.ClusterWideKey],
) {
	if objects == nil {
		return
	}

	// Range can observe an entry concurrently removed from or replaced in the
	// indexes. Matching only schedules another resource reconciliation, which
	// validates the current object and policy state before taking action.
	objects.Range(func(keyValue, objectValue any) bool {
		key := keyValue.(keys.ClusterWideKey)
		if matched.Has(key) {
			return true
		}

		matches := selector.name == key.Name
		if selector.name == "" {
			matches = selector.labelSelector == nil || selector.labelSelector.Matches(labels.Set(objectValue.(map[string]string)))
		}
		if matches {
			matched.Insert(key)
		}
		return true
	})
}

func waitingObjectKey(selector waitingObjectSelector) keys.ClusterWideKey {
	return keys.ClusterWideKey{
		Group:     selector.gvk.Group,
		Version:   selector.gvk.Version,
		Kind:      selector.gvk.Kind,
		Namespace: selector.namespace,
		Name:      selector.name,
	}
}

func compileWaitingObjectSelectors(resourceSelectors []policyv1alpha1.ResourceSelector) []waitingObjectSelector {
	selectors := make([]waitingObjectSelector, 0, len(resourceSelectors))
	for _, resourceSelector := range resourceSelectors {
		groupVersion, err := schema.ParseGroupVersion(resourceSelector.APIVersion)
		if err != nil {
			continue
		}

		selector := waitingObjectSelector{
			gvk:       groupVersion.WithKind(resourceSelector.Kind),
			namespace: resourceSelector.Namespace,
			name:      resourceSelector.Name,
		}
		if selector.name == "" && resourceSelector.LabelSelector != nil {
			selector.labelSelector, err = metav1.LabelSelectorAsSelector(resourceSelector.LabelSelector)
			if err != nil {
				continue
			}
		}
		selectors = append(selectors, selector)
	}
	return selectors
}

func (s *waitingObjectStore) getObjectsBucket(gvk schema.GroupVersionKind, namespace string) *sync.Map {
	namespaces, exists := s.cache.Load(gvk)
	if !exists {
		return nil
	}
	objects, exists := namespaces.(*sync.Map).Load(namespace)
	if !exists {
		return nil
	}
	return objects.(*sync.Map)
}

func (s *waitingObjectStore) getOrCreateObjectsBucket(gvk schema.GroupVersionKind, namespace string) *sync.Map {
	namespacesAny, exists := s.cache.Load(gvk)
	if !exists {
		namespacesAny, _ = s.cache.LoadOrStore(gvk, &sync.Map{})
	}
	namespaces := namespacesAny.(*sync.Map)
	objects, exists := namespaces.Load(namespace)
	if !exists {
		objects, _ = namespaces.LoadOrStore(namespace, &sync.Map{})
	}
	return objects.(*sync.Map)
}
