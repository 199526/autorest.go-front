//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"encoding/xml"
	"time"
)

// MarshalXML implements the xml.Marshaller interface for type AccessPolicy.
func (a AccessPolicy) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias AccessPolicy
	aux := &struct {
		*alias
		Expiry *timeRFC3339 `xml:"Expiry"`
		Start  *timeRFC3339 `xml:"Start"`
	}{
		alias:  (*alias)(&a),
		Expiry: (*timeRFC3339)(a.Expiry),
		Start:  (*timeRFC3339)(a.Start),
	}
	return e.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type AccessPolicy.
func (a *AccessPolicy) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias AccessPolicy
	aux := &struct {
		*alias
		Expiry *timeRFC3339 `xml:"Expiry"`
		Start  *timeRFC3339 `xml:"Start"`
	}{
		alias: (*alias)(a),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	a.Expiry = (*time.Time)(aux.Expiry)
	a.Start = (*time.Time)(aux.Start)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type ArrowConfiguration.
func (a ArrowConfiguration) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias ArrowConfiguration
	aux := &struct {
		*alias
		Schema *[]*ArrowField `xml:"Schema>Field"`
	}{
		alias: (*alias)(&a),
	}
	if a.Schema != nil {
		aux.Schema = &a.Schema
	}
	return e.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type BlockList.
func (b BlockList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias BlockList
	aux := &struct {
		*alias
		CommittedBlocks   *[]*Block `xml:"CommittedBlocks>Block"`
		UncommittedBlocks *[]*Block `xml:"UncommittedBlocks>Block"`
	}{
		alias: (*alias)(&b),
	}
	if b.CommittedBlocks != nil {
		aux.CommittedBlocks = &b.CommittedBlocks
	}
	if b.UncommittedBlocks != nil {
		aux.UncommittedBlocks = &b.UncommittedBlocks
	}
	return e.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type BlockLookupList.
func (b BlockLookupList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "BlockList"
	type alias BlockLookupList
	aux := &struct {
		*alias
		Committed   *[]*string `xml:"Committed"`
		Latest      *[]*string `xml:"Latest"`
		Uncommitted *[]*string `xml:"Uncommitted"`
	}{
		alias: (*alias)(&b),
	}
	if b.Committed != nil {
		aux.Committed = &b.Committed
	}
	if b.Latest != nil {
		aux.Latest = &b.Latest
	}
	if b.Uncommitted != nil {
		aux.Uncommitted = &b.Uncommitted
	}
	return e.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type ContainerItem.
func (c *ContainerItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias ContainerItem
	aux := &struct {
		*alias
		Metadata additionalProperties `xml:"Metadata"`
	}{
		alias: (*alias)(c),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	c.Metadata = (map[string]*string)(aux.Metadata)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type ContainerProperties.
func (c ContainerProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias ContainerProperties
	aux := &struct {
		*alias
		DeletedTime  *timeRFC1123 `xml:"DeletedTime"`
		LastModified *timeRFC1123 `xml:"Last-Modified"`
	}{
		alias:        (*alias)(&c),
		DeletedTime:  (*timeRFC1123)(c.DeletedTime),
		LastModified: (*timeRFC1123)(c.LastModified),
	}
	return e.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type ContainerProperties.
func (c *ContainerProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias ContainerProperties
	aux := &struct {
		*alias
		DeletedTime  *timeRFC1123 `xml:"DeletedTime"`
		LastModified *timeRFC1123 `xml:"Last-Modified"`
	}{
		alias: (*alias)(c),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	c.DeletedTime = (*time.Time)(aux.DeletedTime)
	c.LastModified = (*time.Time)(aux.LastModified)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type FilterBlobSegment.
func (f FilterBlobSegment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias FilterBlobSegment
	aux := &struct {
		*alias
		Blobs *[]*FilterBlobItem `xml:"Blobs>Blob"`
	}{
		alias: (*alias)(&f),
	}
	if f.Blobs != nil {
		aux.Blobs = &f.Blobs
	}
	return e.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type FlatListSegment.
func (f FlatListSegment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias FlatListSegment
	aux := &struct {
		*alias
		BlobItems *[]*ItemInternal `xml:"Blob"`
	}{
		alias: (*alias)(&f),
	}
	if f.BlobItems != nil {
		aux.BlobItems = &f.BlobItems
	}
	return e.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type GeoReplication.
func (g GeoReplication) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias GeoReplication
	aux := &struct {
		*alias
		LastSyncTime *timeRFC1123 `xml:"LastSyncTime"`
	}{
		alias:        (*alias)(&g),
		LastSyncTime: (*timeRFC1123)(g.LastSyncTime),
	}
	return e.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type GeoReplication.
func (g *GeoReplication) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias GeoReplication
	aux := &struct {
		*alias
		LastSyncTime *timeRFC1123 `xml:"LastSyncTime"`
	}{
		alias: (*alias)(g),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	g.LastSyncTime = (*time.Time)(aux.LastSyncTime)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type HierarchyListSegment.
func (h HierarchyListSegment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias HierarchyListSegment
	aux := &struct {
		*alias
		BlobItems    *[]*ItemInternal `xml:"Blob"`
		BlobPrefixes *[]*Prefix       `xml:"Prefix"`
	}{
		alias: (*alias)(&h),
	}
	if h.BlobItems != nil {
		aux.BlobItems = &h.BlobItems
	}
	if h.BlobPrefixes != nil {
		aux.BlobPrefixes = &h.BlobPrefixes
	}
	return e.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type ItemInternal.
func (i *ItemInternal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias ItemInternal
	aux := &struct {
		*alias
		ObjectReplicationMetadata additionalProperties `xml:"OrMetadata"`
	}{
		alias: (*alias)(i),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	i.ObjectReplicationMetadata = (map[string]*string)(aux.ObjectReplicationMetadata)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type ListContainersSegmentResponse.
func (l ListContainersSegmentResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias ListContainersSegmentResponse
	aux := &struct {
		*alias
		ContainerItems *[]*ContainerItem `xml:"Containers>Container"`
	}{
		alias: (*alias)(&l),
	}
	if l.ContainerItems != nil {
		aux.ContainerItems = &l.ContainerItems
	}
	return e.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type PageList.
func (p PageList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias PageList
	aux := &struct {
		*alias
		ClearRange *[]*ClearRange `xml:"ClearRange"`
		PageRange  *[]*PageRange  `xml:"PageRange"`
	}{
		alias: (*alias)(&p),
	}
	if p.ClearRange != nil {
		aux.ClearRange = &p.ClearRange
	}
	if p.PageRange != nil {
		aux.PageRange = &p.PageRange
	}
	return e.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type PropertiesInternal.
func (p PropertiesInternal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias PropertiesInternal
	aux := &struct {
		*alias
		AccessTierChangeTime        *timeRFC1123 `xml:"AccessTierChangeTime"`
		ContentMD5                  *[]byte      `xml:"Content-MD5"`
		CopyCompletionTime          *timeRFC1123 `xml:"CopyCompletionTime"`
		CreationTime                *timeRFC1123 `xml:"Creation-Time"`
		DeletedTime                 *timeRFC1123 `xml:"DeletedTime"`
		ExpiresOn                   *timeRFC1123 `xml:"Expiry-Time"`
		ImmutabilityPolicyExpiresOn *timeRFC1123 `xml:"ImmutabilityPolicyUntilDate"`
		LastAccessedOn              *timeRFC1123 `xml:"LastAccessTime"`
		LastModified                *timeRFC1123 `xml:"Last-Modified"`
	}{
		alias:                       (*alias)(&p),
		AccessTierChangeTime:        (*timeRFC1123)(p.AccessTierChangeTime),
		CopyCompletionTime:          (*timeRFC1123)(p.CopyCompletionTime),
		CreationTime:                (*timeRFC1123)(p.CreationTime),
		DeletedTime:                 (*timeRFC1123)(p.DeletedTime),
		ExpiresOn:                   (*timeRFC1123)(p.ExpiresOn),
		ImmutabilityPolicyExpiresOn: (*timeRFC1123)(p.ImmutabilityPolicyExpiresOn),
		LastAccessedOn:              (*timeRFC1123)(p.LastAccessedOn),
		LastModified:                (*timeRFC1123)(p.LastModified),
	}
	if p.ContentMD5 != nil {
		aux.ContentMD5 = &p.ContentMD5
	}
	return e.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type PropertiesInternal.
func (p *PropertiesInternal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias PropertiesInternal
	aux := &struct {
		*alias
		AccessTierChangeTime        *timeRFC1123 `xml:"AccessTierChangeTime"`
		ContentMD5                  *[]byte      `xml:"Content-MD5"`
		CopyCompletionTime          *timeRFC1123 `xml:"CopyCompletionTime"`
		CreationTime                *timeRFC1123 `xml:"Creation-Time"`
		DeletedTime                 *timeRFC1123 `xml:"DeletedTime"`
		ExpiresOn                   *timeRFC1123 `xml:"Expiry-Time"`
		ImmutabilityPolicyExpiresOn *timeRFC1123 `xml:"ImmutabilityPolicyUntilDate"`
		LastAccessedOn              *timeRFC1123 `xml:"LastAccessTime"`
		LastModified                *timeRFC1123 `xml:"Last-Modified"`
	}{
		alias: (*alias)(p),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	p.AccessTierChangeTime = (*time.Time)(aux.AccessTierChangeTime)
	p.CopyCompletionTime = (*time.Time)(aux.CopyCompletionTime)
	p.CreationTime = (*time.Time)(aux.CreationTime)
	p.DeletedTime = (*time.Time)(aux.DeletedTime)
	p.ExpiresOn = (*time.Time)(aux.ExpiresOn)
	p.ImmutabilityPolicyExpiresOn = (*time.Time)(aux.ImmutabilityPolicyExpiresOn)
	p.LastAccessedOn = (*time.Time)(aux.LastAccessedOn)
	p.LastModified = (*time.Time)(aux.LastModified)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type QueryRequest.
func (q QueryRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "QueryRequest"
	type alias QueryRequest
	aux := &struct {
		*alias
	}{
		alias: (*alias)(&q),
	}
	return e.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type StorageServiceProperties.
func (s StorageServiceProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias StorageServiceProperties
	aux := &struct {
		*alias
		Cors *[]*CorsRule `xml:"Cors>CorsRule"`
	}{
		alias: (*alias)(&s),
	}
	if s.Cors != nil {
		aux.Cors = &s.Cors
	}
	return e.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type Tags.
func (t Tags) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Tags"
	type alias Tags
	aux := &struct {
		*alias
		BlobTagSet *[]*Tag `xml:"TagSet>Tag"`
	}{
		alias: (*alias)(&t),
	}
	if t.BlobTagSet != nil {
		aux.BlobTagSet = &t.BlobTagSet
	}
	return e.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type UserDelegationKey.
func (u UserDelegationKey) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias UserDelegationKey
	aux := &struct {
		*alias
		SignedExpiry *timeRFC3339 `xml:"SignedExpiry"`
		SignedStart  *timeRFC3339 `xml:"SignedStart"`
	}{
		alias:        (*alias)(&u),
		SignedExpiry: (*timeRFC3339)(u.SignedExpiry),
		SignedStart:  (*timeRFC3339)(u.SignedStart),
	}
	return e.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type UserDelegationKey.
func (u *UserDelegationKey) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias UserDelegationKey
	aux := &struct {
		*alias
		SignedExpiry *timeRFC3339 `xml:"SignedExpiry"`
		SignedStart  *timeRFC3339 `xml:"SignedStart"`
	}{
		alias: (*alias)(u),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	u.SignedExpiry = (*time.Time)(aux.SignedExpiry)
	u.SignedStart = (*time.Time)(aux.SignedStart)
	return nil
}
