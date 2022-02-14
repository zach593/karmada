package interpreter

import (
	"encoding/json"
	"net/http"

	"gomodules.xyz/jsonpatch/v2"

	configv1alpha1 "github.com/zach593/karmada/pkg/apis/config/v1alpha1"
)

// Errored creates a new Response for error-handling a request.
func Errored(code int32, err error) Response {
	return Response{
		ResourceInterpreterResponse: configv1alpha1.ResourceInterpreterResponse{
			Successful: false,
			Status: &configv1alpha1.RequestStatus{
				Code:    code,
				Message: err.Error(),
			},
		},
	}
}

// Succeeded constructs a response indicating the given operation is handled successfully.
func Succeeded(msg string) Response {
	return ValidationResponse(true, msg)
}

// ValidationResponse returns a response for handle a interpret request.
func ValidationResponse(successful bool, msg string) Response {
	code := http.StatusForbidden
	if successful {
		code = http.StatusOK
	}
	return Response{
		ResourceInterpreterResponse: configv1alpha1.ResourceInterpreterResponse{
			Successful: successful,
			Status: &configv1alpha1.RequestStatus{
				Code:    int32(code),
				Message: msg,
			},
		},
	}
}

// PatchResponseFromRaw takes 2 byte arrays and returns a new response with patch.
func PatchResponseFromRaw(original, current []byte) Response {
	patches, err := jsonpatch.CreatePatch(original, current)
	if err != nil {
		return Errored(http.StatusInternalServerError, err)
	}

	if len(patches) == 0 {
		return Succeeded("")
	}

	patch, err := json.Marshal(patches)
	if err != nil {
		return Errored(http.StatusInternalServerError, err)
	}

	patchType := configv1alpha1.PatchTypeJSONPatch
	return Response{
		ResourceInterpreterResponse: configv1alpha1.ResourceInterpreterResponse{
			Successful: true,
			Patch:      patch,
			PatchType:  &patchType,
		},
	}
}
