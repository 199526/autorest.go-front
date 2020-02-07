// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"net/http"
	"net/url"
	"path"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// PrimitiveOperations ...
type PrimitiveOperations struct{}

// GetIntCreateRequest creates the GetInt request.
func (PrimitiveOperations) GetIntCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/integer")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetIntHandleResponse handles the GetInt response.
func (PrimitiveOperations) GetIntHandleResponse(resp *azcore.Response) (*GetIntResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetIntResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.IntWrapper)
}

// PutIntCreateRequest creates the PutInt request.
func (PrimitiveOperations) PutIntCreateRequest(u url.URL, complexBody IntWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/integer")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutIntHandleResponse handles the PutInt response.
func (PrimitiveOperations) PutIntHandleResponse(resp *azcore.Response) (*PutIntResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutIntResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.IntWrapper)
}

// GetLongCreateRequest creates the GetLong request.
func (PrimitiveOperations) GetLongCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/long")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetLongHandleResponse handles the GetLong response.
func (PrimitiveOperations) GetLongHandleResponse(resp *azcore.Response) (*GetLongResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetLongResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.LongWrapper)
}

// PutLongCreateRequest creates the PutLong request.
func (PrimitiveOperations) PutLongCreateRequest(u url.URL, complexBody LongWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/long")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutLongHandleResponse handles the PutLong response.
func (PrimitiveOperations) PutLongHandleResponse(resp *azcore.Response) (*PutLongResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutLongResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.LongWrapper)
}

// GetFloatCreateRequest creates the GetFloat request.
func (PrimitiveOperations) GetFloatCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/float")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetFloatHandleResponse handles the GetFloat response.
func (PrimitiveOperations) GetFloatHandleResponse(resp *azcore.Response) (*GetFloatResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetFloatResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.FloatWrapper)
}

// PutFloatCreateRequest creates the PutFloat request.
func (PrimitiveOperations) PutFloatCreateRequest(u url.URL, complexBody FloatWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/float")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutFloatHandleResponse handles the PutFloat response.
func (PrimitiveOperations) PutFloatHandleResponse(resp *azcore.Response) (*PutFloatResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutFloatResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.FloatWrapper)
}

// GetDoubleCreateRequest creates the GetDouble request.
func (PrimitiveOperations) GetDoubleCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/double")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetDoubleHandleResponse handles the GetDouble response.
func (PrimitiveOperations) GetDoubleHandleResponse(resp *azcore.Response) (*GetDoubleResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetDoubleResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.DoubleWrapper)
}

// PutDoubleCreateRequest creates the PutDouble request.
func (PrimitiveOperations) PutDoubleCreateRequest(u url.URL, complexBody DoubleWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/double")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutDoubleHandleResponse handles the PutDouble response.
func (PrimitiveOperations) PutDoubleHandleResponse(resp *azcore.Response) (*PutDoubleResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutDoubleResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.DoubleWrapper)
}

// GetBoolCreateRequest creates the GetBool request.
func (PrimitiveOperations) GetBoolCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/bool")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetBoolHandleResponse handles the GetBool response.
func (PrimitiveOperations) GetBoolHandleResponse(resp *azcore.Response) (*GetBoolResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetBoolResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.BooleanWrapper)
}

// PutBoolCreateRequest creates the PutBool request.
func (PrimitiveOperations) PutBoolCreateRequest(u url.URL, complexBody BooleanWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/bool")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutBoolHandleResponse handles the PutBool response.
func (PrimitiveOperations) PutBoolHandleResponse(resp *azcore.Response) (*PutBoolResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutBoolResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.BooleanWrapper)
}

// GetStringCreateRequest creates the GetString request.
func (PrimitiveOperations) GetStringCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/string")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetStringHandleResponse handles the GetString response.
func (PrimitiveOperations) GetStringHandleResponse(resp *azcore.Response) (*GetStringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetStringResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.StringWrapper)
}

// PutStringCreateRequest creates the PutString request.
func (PrimitiveOperations) PutStringCreateRequest(u url.URL, complexBody StringWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/string")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutStringHandleResponse handles the PutString response.
func (PrimitiveOperations) PutStringHandleResponse(resp *azcore.Response) (*PutStringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutStringResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.StringWrapper)
}

// GetDateCreateRequest creates the GetDate request.
func (PrimitiveOperations) GetDateCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/date")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetDateHandleResponse handles the GetDate response.
func (PrimitiveOperations) GetDateHandleResponse(resp *azcore.Response) (*GetDateResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := GetDateResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.DateWrapper)
}

// PutDateCreateRequest creates the PutDate request.
func (PrimitiveOperations) PutDateCreateRequest(u url.URL, complexBody DateWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/date")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutDateHandleResponse handles the PutDate response.
func (PrimitiveOperations) PutDateHandleResponse(resp *azcore.Response) (*PutDateResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PutDateResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.DateWrapper)
}