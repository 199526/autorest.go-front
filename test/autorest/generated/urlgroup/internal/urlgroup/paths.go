// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package urlgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"time"
)

type PathsOperations struct{}

// ArrayCsvInPathCreateRequest creates the ArrayCsvInPath request.
func (PathsOperations) ArrayCsvInPathCreateRequest(u url.URL, arrayPath []string) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/array/ArrayPath1%2cbegin%21%2A%27%28%29%3B%3A%40%20%26%3D%2B%24%2C%2F%3F%23%5B%5Dend%2c%2c/{arrayPath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// ArrayCsvInPathHandleResponse handles the ArrayCsvInPath response.
func (PathsOperations) ArrayCsvInPathHandleResponse(resp *azcore.Response) (*PathsArrayCsvInPathResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsArrayCsvInPathResponse{StatusCode: resp.StatusCode}, nil
}

// Base64URLCreateRequest creates the Base64URL request.
func (PathsOperations) Base64URLCreateRequest(u url.URL, base64UrlPath []byte) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/string/bG9yZW0/{base64UrlPath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// Base64URLHandleResponse handles the Base64URL response.
func (PathsOperations) Base64URLHandleResponse(resp *azcore.Response) (*PathsBase64URLResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsBase64URLResponse{StatusCode: resp.StatusCode}, nil
}

// ByteEmptyCreateRequest creates the ByteEmpty request.
func (PathsOperations) ByteEmptyCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/byte/empty/{bytePath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// ByteEmptyHandleResponse handles the ByteEmpty response.
func (PathsOperations) ByteEmptyHandleResponse(resp *azcore.Response) (*PathsByteEmptyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsByteEmptyResponse{StatusCode: resp.StatusCode}, nil
}

// ByteMultiByteCreateRequest creates the ByteMultiByte request.
func (PathsOperations) ByteMultiByteCreateRequest(u url.URL, bytePath []byte) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/byte/multibyte/{bytePath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// ByteMultiByteHandleResponse handles the ByteMultiByte response.
func (PathsOperations) ByteMultiByteHandleResponse(resp *azcore.Response) (*PathsByteMultiByteResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsByteMultiByteResponse{StatusCode: resp.StatusCode}, nil
}

// ByteNullCreateRequest creates the ByteNull request.
func (PathsOperations) ByteNullCreateRequest(u url.URL, bytePath []byte) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/byte/null/{bytePath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// ByteNullHandleResponse handles the ByteNull response.
func (PathsOperations) ByteNullHandleResponse(resp *azcore.Response) (*PathsByteNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsByteNullResponse{StatusCode: resp.StatusCode}, nil
}

// DateNullCreateRequest creates the DateNull request.
func (PathsOperations) DateNullCreateRequest(u url.URL, datePath time.Time) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/date/null/{datePath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// DateNullHandleResponse handles the DateNull response.
func (PathsOperations) DateNullHandleResponse(resp *azcore.Response) (*PathsDateNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsDateNullResponse{StatusCode: resp.StatusCode}, nil
}

// DateTimeNullCreateRequest creates the DateTimeNull request.
func (PathsOperations) DateTimeNullCreateRequest(u url.URL, dateTimePath time.Time) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/datetime/null/{dateTimePath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// DateTimeNullHandleResponse handles the DateTimeNull response.
func (PathsOperations) DateTimeNullHandleResponse(resp *azcore.Response) (*PathsDateTimeNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsDateTimeNullResponse{StatusCode: resp.StatusCode}, nil
}

// DateTimeValidCreateRequest creates the DateTimeValid request.
func (PathsOperations) DateTimeValidCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/datetime/2012-01-01T01%3A01%3A01Z/{dateTimePath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// DateTimeValidHandleResponse handles the DateTimeValid response.
func (PathsOperations) DateTimeValidHandleResponse(resp *azcore.Response) (*PathsDateTimeValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsDateTimeValidResponse{StatusCode: resp.StatusCode}, nil
}

// DateValidCreateRequest creates the DateValid request.
func (PathsOperations) DateValidCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/date/2012-01-01/{datePath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// DateValidHandleResponse handles the DateValid response.
func (PathsOperations) DateValidHandleResponse(resp *azcore.Response) (*PathsDateValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsDateValidResponse{StatusCode: resp.StatusCode}, nil
}

// DoubleDecimalNegativeCreateRequest creates the DoubleDecimalNegative request.
func (PathsOperations) DoubleDecimalNegativeCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/double/-9999999.999/{doublePath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// DoubleDecimalNegativeHandleResponse handles the DoubleDecimalNegative response.
func (PathsOperations) DoubleDecimalNegativeHandleResponse(resp *azcore.Response) (*PathsDoubleDecimalNegativeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsDoubleDecimalNegativeResponse{StatusCode: resp.StatusCode}, nil
}

// DoubleDecimalPositiveCreateRequest creates the DoubleDecimalPositive request.
func (PathsOperations) DoubleDecimalPositiveCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/double/9999999.999/{doublePath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// DoubleDecimalPositiveHandleResponse handles the DoubleDecimalPositive response.
func (PathsOperations) DoubleDecimalPositiveHandleResponse(resp *azcore.Response) (*PathsDoubleDecimalPositiveResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsDoubleDecimalPositiveResponse{StatusCode: resp.StatusCode}, nil
}

// EnumNullCreateRequest creates the EnumNull request.
func (PathsOperations) EnumNullCreateRequest(u url.URL, enumPath UriColor) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/string/null/{enumPath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// EnumNullHandleResponse handles the EnumNull response.
func (PathsOperations) EnumNullHandleResponse(resp *azcore.Response) (*PathsEnumNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsEnumNullResponse{StatusCode: resp.StatusCode}, nil
}

// EnumValidCreateRequest creates the EnumValid request.
func (PathsOperations) EnumValidCreateRequest(u url.URL, enumPath UriColor) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/enum/green%20color/{enumPath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// EnumValidHandleResponse handles the EnumValid response.
func (PathsOperations) EnumValidHandleResponse(resp *azcore.Response) (*PathsEnumValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsEnumValidResponse{StatusCode: resp.StatusCode}, nil
}

// FloatScientificNegativeCreateRequest creates the FloatScientificNegative request.
func (PathsOperations) FloatScientificNegativeCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/float/-1.034E-20/{floatPath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// FloatScientificNegativeHandleResponse handles the FloatScientificNegative response.
func (PathsOperations) FloatScientificNegativeHandleResponse(resp *azcore.Response) (*PathsFloatScientificNegativeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsFloatScientificNegativeResponse{StatusCode: resp.StatusCode}, nil
}

// FloatScientificPositiveCreateRequest creates the FloatScientificPositive request.
func (PathsOperations) FloatScientificPositiveCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/float/1.034E+20/{floatPath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// FloatScientificPositiveHandleResponse handles the FloatScientificPositive response.
func (PathsOperations) FloatScientificPositiveHandleResponse(resp *azcore.Response) (*PathsFloatScientificPositiveResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsFloatScientificPositiveResponse{StatusCode: resp.StatusCode}, nil
}

// GetBooleanFalseCreateRequest creates the GetBooleanFalse request.
func (PathsOperations) GetBooleanFalseCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/bool/false/{boolPath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetBooleanFalseHandleResponse handles the GetBooleanFalse response.
func (PathsOperations) GetBooleanFalseHandleResponse(resp *azcore.Response) (*PathsGetBooleanFalseResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsGetBooleanFalseResponse{StatusCode: resp.StatusCode}, nil
}

// GetBooleanTrueCreateRequest creates the GetBooleanTrue request.
func (PathsOperations) GetBooleanTrueCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/bool/true/{boolPath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetBooleanTrueHandleResponse handles the GetBooleanTrue response.
func (PathsOperations) GetBooleanTrueHandleResponse(resp *azcore.Response) (*PathsGetBooleanTrueResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsGetBooleanTrueResponse{StatusCode: resp.StatusCode}, nil
}

// GetIntNegativeOneMillionCreateRequest creates the GetIntNegativeOneMillion request.
func (PathsOperations) GetIntNegativeOneMillionCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/int/-1000000/{intPath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetIntNegativeOneMillionHandleResponse handles the GetIntNegativeOneMillion response.
func (PathsOperations) GetIntNegativeOneMillionHandleResponse(resp *azcore.Response) (*PathsGetIntNegativeOneMillionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsGetIntNegativeOneMillionResponse{StatusCode: resp.StatusCode}, nil
}

// GetIntOneMillionCreateRequest creates the GetIntOneMillion request.
func (PathsOperations) GetIntOneMillionCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/int/1000000/{intPath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetIntOneMillionHandleResponse handles the GetIntOneMillion response.
func (PathsOperations) GetIntOneMillionHandleResponse(resp *azcore.Response) (*PathsGetIntOneMillionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsGetIntOneMillionResponse{StatusCode: resp.StatusCode}, nil
}

// GetNegativeTenBillionCreateRequest creates the GetNegativeTenBillion request.
func (PathsOperations) GetNegativeTenBillionCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/long/-10000000000/{longPath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetNegativeTenBillionHandleResponse handles the GetNegativeTenBillion response.
func (PathsOperations) GetNegativeTenBillionHandleResponse(resp *azcore.Response) (*PathsGetNegativeTenBillionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsGetNegativeTenBillionResponse{StatusCode: resp.StatusCode}, nil
}

// GetTenBillionCreateRequest creates the GetTenBillion request.
func (PathsOperations) GetTenBillionCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/long/10000000000/{longPath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetTenBillionHandleResponse handles the GetTenBillion response.
func (PathsOperations) GetTenBillionHandleResponse(resp *azcore.Response) (*PathsGetTenBillionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsGetTenBillionResponse{StatusCode: resp.StatusCode}, nil
}

// StringEmptyCreateRequest creates the StringEmpty request.
func (PathsOperations) StringEmptyCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/string/empty/{stringPath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// StringEmptyHandleResponse handles the StringEmpty response.
func (PathsOperations) StringEmptyHandleResponse(resp *azcore.Response) (*PathsStringEmptyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsStringEmptyResponse{StatusCode: resp.StatusCode}, nil
}

// StringNullCreateRequest creates the StringNull request.
func (PathsOperations) StringNullCreateRequest(u url.URL, stringPath string) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/string/null/{stringPath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// StringNullHandleResponse handles the StringNull response.
func (PathsOperations) StringNullHandleResponse(resp *azcore.Response) (*PathsStringNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsStringNullResponse{StatusCode: resp.StatusCode}, nil
}

// StringURLEncodedCreateRequest creates the StringURLEncoded request.
func (PathsOperations) StringURLEncodedCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/string/begin%21%2A%27%28%29%3B%3A%40%20%26%3D%2B%24%2C%2F%3F%23%5B%5Dend/{stringPath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// StringURLEncodedHandleResponse handles the StringURLEncoded response.
func (PathsOperations) StringURLEncodedHandleResponse(resp *azcore.Response) (*PathsStringURLEncodedResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsStringURLEncodedResponse{StatusCode: resp.StatusCode}, nil
}

// StringURLNonEncodedCreateRequest creates the StringURLNonEncoded request.
func (PathsOperations) StringURLNonEncodedCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/string/begin!*'();:@&=+$,end/{stringPath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// StringURLNonEncodedHandleResponse handles the StringURLNonEncoded response.
func (PathsOperations) StringURLNonEncodedHandleResponse(resp *azcore.Response) (*PathsStringURLNonEncodedResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsStringURLNonEncodedResponse{StatusCode: resp.StatusCode}, nil
}

// StringUnicodeCreateRequest creates the StringUnicode request.
func (PathsOperations) StringUnicodeCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/string/unicode/{stringPath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// StringUnicodeHandleResponse handles the StringUnicode response.
func (PathsOperations) StringUnicodeHandleResponse(resp *azcore.Response) (*PathsStringUnicodeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsStringUnicodeResponse{StatusCode: resp.StatusCode}, nil
}

// UnixTimeURLCreateRequest creates the UnixTimeURL request.
func (PathsOperations) UnixTimeURLCreateRequest(u url.URL, unixTimeUrlPath time.Time) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/paths/int/1460505600/{unixTimeUrlPath}")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// UnixTimeURLHandleResponse handles the UnixTimeURL response.
func (PathsOperations) UnixTimeURLHandleResponse(resp *azcore.Response) (*PathsUnixTimeURLResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsUnixTimeURLResponse{StatusCode: resp.StatusCode}, nil
}
