// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"context"
	"encoding/base64"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"strconv"
	"time"
)

type appendBlobClient struct {
	con *connection
}

// AppendBlock - The Append Block operation commits a new block of data to the end of an existing append blob. The Append Block operation is permitted only
// if the blob was created with x-ms-blob-type set to
// AppendBlob. Append Block is supported only on version 2015-02-21 version or later.
func (client *appendBlobClient) AppendBlock(ctx context.Context, contentLength int64, body azcore.ReadSeekCloser, appendBlobAppendBlockOptions *AppendBlobAppendBlockOptions, leaseAccessConditions *LeaseAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (AppendBlobAppendBlockResponse, error) {
	req, err := client.appendBlockCreateRequest(ctx, contentLength, body, appendBlobAppendBlockOptions, leaseAccessConditions, appendPositionAccessConditions, cpkInfo, cpkScopeInfo, modifiedAccessConditions)
	if err != nil {
		return AppendBlobAppendBlockResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return AppendBlobAppendBlockResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return AppendBlobAppendBlockResponse{}, client.appendBlockHandleError(resp)
	}
	return client.appendBlockHandleResponse(resp)
}

// appendBlockCreateRequest creates the AppendBlock request.
func (client *appendBlobClient) appendBlockCreateRequest(ctx context.Context, contentLength int64, body azcore.ReadSeekCloser, appendBlobAppendBlockOptions *AppendBlobAppendBlockOptions, leaseAccessConditions *LeaseAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPut, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("comp", "appendblock")
	if appendBlobAppendBlockOptions != nil && appendBlobAppendBlockOptions.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*appendBlobAppendBlockOptions.Timeout), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if appendBlobAppendBlockOptions != nil && appendBlobAppendBlockOptions.TransactionalContentMD5 != nil {
		req.Header.Set("Content-MD5", base64.StdEncoding.EncodeToString(*appendBlobAppendBlockOptions.TransactionalContentMD5))
	}
	if appendBlobAppendBlockOptions != nil && appendBlobAppendBlockOptions.TransactionalContentCRC64 != nil {
		req.Header.Set("x-ms-content-crc64", base64.StdEncoding.EncodeToString(*appendBlobAppendBlockOptions.TransactionalContentCRC64))
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseID)
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.MaxSize != nil {
		req.Header.Set("x-ms-blob-condition-maxsize", strconv.FormatInt(*appendPositionAccessConditions.MaxSize, 10))
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.AppendPosition != nil {
		req.Header.Set("x-ms-blob-condition-appendpos", strconv.FormatInt(*appendPositionAccessConditions.AppendPosition, 10))
	}
	if cpkInfo != nil && cpkInfo.EncryptionKey != nil {
		req.Header.Set("x-ms-encryption-key", *cpkInfo.EncryptionKey)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKeySHA256 != nil {
		req.Header.Set("x-ms-encryption-key-sha256", *cpkInfo.EncryptionKeySHA256)
	}
	if cpkInfo != nil && cpkInfo.EncryptionAlgorithm != nil {
		req.Header.Set("x-ms-encryption-algorithm", "AES256")
	}
	if cpkScopeInfo != nil && cpkScopeInfo.EncryptionScope != nil {
		req.Header.Set("x-ms-encryption-scope", *cpkScopeInfo.EncryptionScope)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfTags != nil {
		req.Header.Set("x-ms-if-tags", *modifiedAccessConditions.IfTags)
	}
	req.Header.Set("x-ms-version", "2020-06-12")
	if appendBlobAppendBlockOptions != nil && appendBlobAppendBlockOptions.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *appendBlobAppendBlockOptions.RequestID)
	}
	req.Header.Set("Accept", "application/xml")
	return req, req.SetBody(body, "application/octet-stream")
}

// appendBlockHandleResponse handles the AppendBlock response.
func (client *appendBlobClient) appendBlockHandleResponse(resp *azcore.Response) (AppendBlobAppendBlockResponse, error) {
	result := AppendBlobAppendBlockResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobAppendBlockResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("Content-MD5"); val != "" {
		contentMD5, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return AppendBlobAppendBlockResponse{}, err
		}
		result.ContentMD5 = &contentMD5
	}
	if val := resp.Header.Get("x-ms-content-crc64"); val != "" {
		xMSContentCRC64, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return AppendBlobAppendBlockResponse{}, err
		}
		result.XMSContentCRC64 = &xMSContentCRC64
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobAppendBlockResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-blob-append-offset"); val != "" {
		result.BlobAppendOffset = &val
	}
	if val := resp.Header.Get("x-ms-blob-committed-block-count"); val != "" {
		blobCommittedBlockCount32, err := strconv.ParseInt(val, 10, 32)
		blobCommittedBlockCount := int32(blobCommittedBlockCount32)
		if err != nil {
			return AppendBlobAppendBlockResponse{}, err
		}
		result.BlobCommittedBlockCount = &blobCommittedBlockCount
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return AppendBlobAppendBlockResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	if val := resp.Header.Get("x-ms-encryption-key-sha256"); val != "" {
		result.EncryptionKeySHA256 = &val
	}
	if val := resp.Header.Get("x-ms-encryption-scope"); val != "" {
		result.EncryptionScope = &val
	}
	return result, nil
}

// appendBlockHandleError handles the AppendBlock error response.
func (client *appendBlobClient) appendBlockHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// AppendBlockFromURL - The Append Block operation commits a new block of data to the end of an existing append blob where the contents are read from a
// source url. The Append Block operation is permitted only if the blob was
// created with x-ms-blob-type set to AppendBlob. Append Block is supported only on version 2015-02-21 version or later.
func (client *appendBlobClient) AppendBlockFromURL(ctx context.Context, sourceURL string, contentLength int64, appendBlobAppendBlockFromURLOptions *AppendBlobAppendBlockFromURLOptions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, leaseAccessConditions *LeaseAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, sourceModifiedAccessConditions *SourceModifiedAccessConditions) (AppendBlobAppendBlockFromURLResponse, error) {
	req, err := client.appendBlockFromURLCreateRequest(ctx, sourceURL, contentLength, appendBlobAppendBlockFromURLOptions, cpkInfo, cpkScopeInfo, leaseAccessConditions, appendPositionAccessConditions, modifiedAccessConditions, sourceModifiedAccessConditions)
	if err != nil {
		return AppendBlobAppendBlockFromURLResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return AppendBlobAppendBlockFromURLResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return AppendBlobAppendBlockFromURLResponse{}, client.appendBlockFromURLHandleError(resp)
	}
	return client.appendBlockFromURLHandleResponse(resp)
}

// appendBlockFromURLCreateRequest creates the AppendBlockFromURL request.
func (client *appendBlobClient) appendBlockFromURLCreateRequest(ctx context.Context, sourceURL string, contentLength int64, appendBlobAppendBlockFromURLOptions *AppendBlobAppendBlockFromURLOptions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, leaseAccessConditions *LeaseAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, sourceModifiedAccessConditions *SourceModifiedAccessConditions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPut, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("comp", "appendblock")
	if appendBlobAppendBlockFromURLOptions != nil && appendBlobAppendBlockFromURLOptions.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*appendBlobAppendBlockFromURLOptions.Timeout), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-copy-source", sourceURL)
	if appendBlobAppendBlockFromURLOptions != nil && appendBlobAppendBlockFromURLOptions.SourceRange != nil {
		req.Header.Set("x-ms-source-range", *appendBlobAppendBlockFromURLOptions.SourceRange)
	}
	if appendBlobAppendBlockFromURLOptions != nil && appendBlobAppendBlockFromURLOptions.SourceContentMD5 != nil {
		req.Header.Set("x-ms-source-content-md5", base64.StdEncoding.EncodeToString(*appendBlobAppendBlockFromURLOptions.SourceContentMD5))
	}
	if appendBlobAppendBlockFromURLOptions != nil && appendBlobAppendBlockFromURLOptions.SourceContentcrc64 != nil {
		req.Header.Set("x-ms-source-content-crc64", base64.StdEncoding.EncodeToString(*appendBlobAppendBlockFromURLOptions.SourceContentcrc64))
	}
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if appendBlobAppendBlockFromURLOptions != nil && appendBlobAppendBlockFromURLOptions.TransactionalContentMD5 != nil {
		req.Header.Set("Content-MD5", base64.StdEncoding.EncodeToString(*appendBlobAppendBlockFromURLOptions.TransactionalContentMD5))
	}
	if cpkInfo != nil && cpkInfo.EncryptionKey != nil {
		req.Header.Set("x-ms-encryption-key", *cpkInfo.EncryptionKey)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKeySHA256 != nil {
		req.Header.Set("x-ms-encryption-key-sha256", *cpkInfo.EncryptionKeySHA256)
	}
	if cpkInfo != nil && cpkInfo.EncryptionAlgorithm != nil {
		req.Header.Set("x-ms-encryption-algorithm", "AES256")
	}
	if cpkScopeInfo != nil && cpkScopeInfo.EncryptionScope != nil {
		req.Header.Set("x-ms-encryption-scope", *cpkScopeInfo.EncryptionScope)
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseID)
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.MaxSize != nil {
		req.Header.Set("x-ms-blob-condition-maxsize", strconv.FormatInt(*appendPositionAccessConditions.MaxSize, 10))
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.AppendPosition != nil {
		req.Header.Set("x-ms-blob-condition-appendpos", strconv.FormatInt(*appendPositionAccessConditions.AppendPosition, 10))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfTags != nil {
		req.Header.Set("x-ms-if-tags", *modifiedAccessConditions.IfTags)
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfModifiedSince != nil {
		req.Header.Set("x-ms-source-if-modified-since", sourceModifiedAccessConditions.SourceIfModifiedSince.Format(time.RFC1123))
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfUnmodifiedSince != nil {
		req.Header.Set("x-ms-source-if-unmodified-since", sourceModifiedAccessConditions.SourceIfUnmodifiedSince.Format(time.RFC1123))
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfMatch != nil {
		req.Header.Set("x-ms-source-if-match", *sourceModifiedAccessConditions.SourceIfMatch)
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfNoneMatch != nil {
		req.Header.Set("x-ms-source-if-none-match", *sourceModifiedAccessConditions.SourceIfNoneMatch)
	}
	req.Header.Set("x-ms-version", "2020-06-12")
	if appendBlobAppendBlockFromURLOptions != nil && appendBlobAppendBlockFromURLOptions.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *appendBlobAppendBlockFromURLOptions.RequestID)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// appendBlockFromURLHandleResponse handles the AppendBlockFromURL response.
func (client *appendBlobClient) appendBlockFromURLHandleResponse(resp *azcore.Response) (AppendBlobAppendBlockFromURLResponse, error) {
	result := AppendBlobAppendBlockFromURLResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobAppendBlockFromURLResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("Content-MD5"); val != "" {
		contentMD5, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return AppendBlobAppendBlockFromURLResponse{}, err
		}
		result.ContentMD5 = &contentMD5
	}
	if val := resp.Header.Get("x-ms-content-crc64"); val != "" {
		xMSContentCRC64, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return AppendBlobAppendBlockFromURLResponse{}, err
		}
		result.XMSContentCRC64 = &xMSContentCRC64
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobAppendBlockFromURLResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-blob-append-offset"); val != "" {
		result.BlobAppendOffset = &val
	}
	if val := resp.Header.Get("x-ms-blob-committed-block-count"); val != "" {
		blobCommittedBlockCount32, err := strconv.ParseInt(val, 10, 32)
		blobCommittedBlockCount := int32(blobCommittedBlockCount32)
		if err != nil {
			return AppendBlobAppendBlockFromURLResponse{}, err
		}
		result.BlobCommittedBlockCount = &blobCommittedBlockCount
	}
	if val := resp.Header.Get("x-ms-encryption-key-sha256"); val != "" {
		result.EncryptionKeySHA256 = &val
	}
	if val := resp.Header.Get("x-ms-encryption-scope"); val != "" {
		result.EncryptionScope = &val
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return AppendBlobAppendBlockFromURLResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	return result, nil
}

// appendBlockFromURLHandleError handles the AppendBlockFromURL error response.
func (client *appendBlobClient) appendBlockFromURLHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Create - The Create Append Blob operation creates a new append blob.
func (client *appendBlobClient) Create(ctx context.Context, contentLength int64, appendBlobCreateOptions *AppendBlobCreateOptions, blobHTTPHeaders *BlobHTTPHeaders, leaseAccessConditions *LeaseAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (AppendBlobCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, contentLength, appendBlobCreateOptions, blobHTTPHeaders, leaseAccessConditions, cpkInfo, cpkScopeInfo, modifiedAccessConditions)
	if err != nil {
		return AppendBlobCreateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return AppendBlobCreateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return AppendBlobCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *appendBlobClient) createCreateRequest(ctx context.Context, contentLength int64, appendBlobCreateOptions *AppendBlobCreateOptions, blobHTTPHeaders *BlobHTTPHeaders, leaseAccessConditions *LeaseAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPut, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if appendBlobCreateOptions != nil && appendBlobCreateOptions.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*appendBlobCreateOptions.Timeout), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-blob-type", "AppendBlob")
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentType != nil {
		req.Header.Set("x-ms-blob-content-type", *blobHTTPHeaders.BlobContentType)
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentEncoding != nil {
		req.Header.Set("x-ms-blob-content-encoding", *blobHTTPHeaders.BlobContentEncoding)
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentLanguage != nil {
		req.Header.Set("x-ms-blob-content-language", *blobHTTPHeaders.BlobContentLanguage)
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentMD5 != nil {
		req.Header.Set("x-ms-blob-content-md5", base64.StdEncoding.EncodeToString(*blobHTTPHeaders.BlobContentMD5))
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobCacheControl != nil {
		req.Header.Set("x-ms-blob-cache-control", *blobHTTPHeaders.BlobCacheControl)
	}
	if appendBlobCreateOptions != nil && appendBlobCreateOptions.Metadata != nil {
		for k, v := range *appendBlobCreateOptions.Metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseID)
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentDisposition != nil {
		req.Header.Set("x-ms-blob-content-disposition", *blobHTTPHeaders.BlobContentDisposition)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKey != nil {
		req.Header.Set("x-ms-encryption-key", *cpkInfo.EncryptionKey)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKeySHA256 != nil {
		req.Header.Set("x-ms-encryption-key-sha256", *cpkInfo.EncryptionKeySHA256)
	}
	if cpkInfo != nil && cpkInfo.EncryptionAlgorithm != nil {
		req.Header.Set("x-ms-encryption-algorithm", "AES256")
	}
	if cpkScopeInfo != nil && cpkScopeInfo.EncryptionScope != nil {
		req.Header.Set("x-ms-encryption-scope", *cpkScopeInfo.EncryptionScope)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfTags != nil {
		req.Header.Set("x-ms-if-tags", *modifiedAccessConditions.IfTags)
	}
	req.Header.Set("x-ms-version", "2020-06-12")
	if appendBlobCreateOptions != nil && appendBlobCreateOptions.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *appendBlobCreateOptions.RequestID)
	}
	if appendBlobCreateOptions != nil && appendBlobCreateOptions.BlobTagsString != nil {
		req.Header.Set("x-ms-tags", *appendBlobCreateOptions.BlobTagsString)
	}
	if appendBlobCreateOptions != nil && appendBlobCreateOptions.ImmutabilityPolicyExpiry != nil {
		req.Header.Set("x-ms-immutability-policy-until-date", appendBlobCreateOptions.ImmutabilityPolicyExpiry.Format(time.RFC1123))
	}
	if appendBlobCreateOptions != nil && appendBlobCreateOptions.ImmutabilityPolicyMode != nil {
		req.Header.Set("x-ms-immutability-policy-mode", string(*appendBlobCreateOptions.ImmutabilityPolicyMode))
	}
	if appendBlobCreateOptions != nil && appendBlobCreateOptions.LegalHold != nil {
		req.Header.Set("x-ms-legal-hold", strconv.FormatBool(*appendBlobCreateOptions.LegalHold))
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *appendBlobClient) createHandleResponse(resp *azcore.Response) (AppendBlobCreateResponse, error) {
	result := AppendBlobCreateResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobCreateResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("Content-MD5"); val != "" {
		contentMD5, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return AppendBlobCreateResponse{}, err
		}
		result.ContentMD5 = &contentMD5
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("x-ms-version-id"); val != "" {
		result.VersionID = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobCreateResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return AppendBlobCreateResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	if val := resp.Header.Get("x-ms-encryption-key-sha256"); val != "" {
		result.EncryptionKeySHA256 = &val
	}
	if val := resp.Header.Get("x-ms-encryption-scope"); val != "" {
		result.EncryptionScope = &val
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *appendBlobClient) createHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Seal - The Seal operation seals the Append Blob to make it read-only. Seal is supported only on version 2019-12-12 version or later.
func (client *appendBlobClient) Seal(ctx context.Context, appendBlobSealOptions *AppendBlobSealOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions) (AppendBlobSealResponse, error) {
	req, err := client.sealCreateRequest(ctx, appendBlobSealOptions, leaseAccessConditions, modifiedAccessConditions, appendPositionAccessConditions)
	if err != nil {
		return AppendBlobSealResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return AppendBlobSealResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return AppendBlobSealResponse{}, client.sealHandleError(resp)
	}
	return client.sealHandleResponse(resp)
}

// sealCreateRequest creates the Seal request.
func (client *appendBlobClient) sealCreateRequest(ctx context.Context, appendBlobSealOptions *AppendBlobSealOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPut, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("comp", "seal")
	if appendBlobSealOptions != nil && appendBlobSealOptions.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*appendBlobSealOptions.Timeout), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2020-06-12")
	if appendBlobSealOptions != nil && appendBlobSealOptions.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *appendBlobSealOptions.RequestID)
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseID)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.AppendPosition != nil {
		req.Header.Set("x-ms-blob-condition-appendpos", strconv.FormatInt(*appendPositionAccessConditions.AppendPosition, 10))
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// sealHandleResponse handles the Seal response.
func (client *appendBlobClient) sealHandleResponse(resp *azcore.Response) (AppendBlobSealResponse, error) {
	result := AppendBlobSealResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobSealResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobSealResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-blob-sealed"); val != "" {
		isSealed, err := strconv.ParseBool(val)
		if err != nil {
			return AppendBlobSealResponse{}, err
		}
		result.IsSealed = &isSealed
	}
	return result, nil
}

// sealHandleError handles the Seal error response.
func (client *appendBlobClient) sealHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}
