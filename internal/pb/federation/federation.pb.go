// Copyright 2020 the Exposure Notifications Server authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: internal/pb/federation/federation.proto

package federation

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExposureKey_ReportType int32

const (
	ExposureKey_UNKNOWN                      ExposureKey_ReportType = 0 // Not used by this protocol
	ExposureKey_CONFIRMED_TEST               ExposureKey_ReportType = 1
	ExposureKey_CONFIRMED_CLINICAL_DIAGNOSIS ExposureKey_ReportType = 2
	ExposureKey_SELF_REPORT                  ExposureKey_ReportType = 3 // Not used by this protocol
	ExposureKey_RECURSIVE                    ExposureKey_ReportType = 4 // Not used by this protocol
	ExposureKey_REVOKED                      ExposureKey_ReportType = 5 // Used to revoke a key, never returned by client API.
)

// Enum value maps for ExposureKey_ReportType.
var (
	ExposureKey_ReportType_name = map[int32]string{
		0: "UNKNOWN",
		1: "CONFIRMED_TEST",
		2: "CONFIRMED_CLINICAL_DIAGNOSIS",
		3: "SELF_REPORT",
		4: "RECURSIVE",
		5: "REVOKED",
	}
	ExposureKey_ReportType_value = map[string]int32{
		"UNKNOWN":                      0,
		"CONFIRMED_TEST":               1,
		"CONFIRMED_CLINICAL_DIAGNOSIS": 2,
		"SELF_REPORT":                  3,
		"RECURSIVE":                    4,
		"REVOKED":                      5,
	}
)

func (x ExposureKey_ReportType) Enum() *ExposureKey_ReportType {
	p := new(ExposureKey_ReportType)
	*p = x
	return p
}

func (x ExposureKey_ReportType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExposureKey_ReportType) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_pb_federation_federation_proto_enumTypes[0].Descriptor()
}

func (ExposureKey_ReportType) Type() protoreflect.EnumType {
	return &file_internal_pb_federation_federation_proto_enumTypes[0]
}

func (x ExposureKey_ReportType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExposureKey_ReportType.Descriptor instead.
func (ExposureKey_ReportType) EnumDescriptor() ([]byte, []int) {
	return file_internal_pb_federation_federation_proto_rawDescGZIP(), []int{4, 0}
}

type FederationFetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncludeRegions      []string `protobuf:"bytes,1,rep,name=includeRegions,proto3" json:"includeRegions,omitempty"`
	ExcludeRegions      []string `protobuf:"bytes,2,rep,name=excludeRegions,proto3" json:"excludeRegions,omitempty"`
	OnlyTravelers       bool     `protobuf:"varint,3,opt,name=onlyTravelers,proto3" json:"onlyTravelers,omitempty"`
	OnlyLocalProvenance bool     `protobuf:"varint,4,opt,name=onlyLocalProvenance,proto3" json:"onlyLocalProvenance,omitempty"`
	// Max overall exposure keys to fetch.
	// Upper limit is 500.
	MaxExposureKeys uint32 `protobuf:"varint,5,opt,name=maxExposureKeys,proto3" json:"maxExposureKeys,omitempty"`
	// region, includeTravelers, onlyTravelers must be stable to send a fetchToken.
	// initial query should send an empty fetch state token.
	State *FetchState `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *FederationFetchRequest) Reset() {
	*x = FederationFetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_federation_federation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederationFetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederationFetchRequest) ProtoMessage() {}

func (x *FederationFetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_federation_federation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederationFetchRequest.ProtoReflect.Descriptor instead.
func (*FederationFetchRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_federation_federation_proto_rawDescGZIP(), []int{0}
}

func (x *FederationFetchRequest) GetIncludeRegions() []string {
	if x != nil {
		return x.IncludeRegions
	}
	return nil
}

func (x *FederationFetchRequest) GetExcludeRegions() []string {
	if x != nil {
		return x.ExcludeRegions
	}
	return nil
}

func (x *FederationFetchRequest) GetOnlyTravelers() bool {
	if x != nil {
		return x.OnlyTravelers
	}
	return false
}

func (x *FederationFetchRequest) GetOnlyLocalProvenance() bool {
	if x != nil {
		return x.OnlyLocalProvenance
	}
	return false
}

func (x *FederationFetchRequest) GetMaxExposureKeys() uint32 {
	if x != nil {
		return x.MaxExposureKeys
	}
	return 0
}

func (x *FederationFetchRequest) GetState() *FetchState {
	if x != nil {
		return x.State
	}
	return nil
}

type FederationFetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys            []*ExposureKey `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	RevisedKeys     []*ExposureKey `protobuf:"bytes,2,rep,name=revisedKeys,proto3" json:"revisedKeys,omitempty"`
	PartialResponse bool           `protobuf:"varint,3,opt,name=partialResponse,proto3" json:"partialResponse,omitempty"` // required
	NextFetchState  *FetchState    `protobuf:"bytes,4,opt,name=nextFetchState,proto3" json:"nextFetchState,omitempty"`    // nextFetchState allows the query to be continued.
}

func (x *FederationFetchResponse) Reset() {
	*x = FederationFetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_federation_federation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederationFetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederationFetchResponse) ProtoMessage() {}

func (x *FederationFetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_federation_federation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederationFetchResponse.ProtoReflect.Descriptor instead.
func (*FederationFetchResponse) Descriptor() ([]byte, []int) {
	return file_internal_pb_federation_federation_proto_rawDescGZIP(), []int{1}
}

func (x *FederationFetchResponse) GetKeys() []*ExposureKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *FederationFetchResponse) GetRevisedKeys() []*ExposureKey {
	if x != nil {
		return x.RevisedKeys
	}
	return nil
}

func (x *FederationFetchResponse) GetPartialResponse() bool {
	if x != nil {
		return x.PartialResponse
	}
	return false
}

func (x *FederationFetchResponse) GetNextFetchState() *FetchState {
	if x != nil {
		return x.NextFetchState
	}
	return nil
}

type FetchState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyCursor        *Cursor `protobuf:"bytes,1,opt,name=keyCursor,proto3" json:"keyCursor,omitempty"`
	RevisedKeyCursor *Cursor `protobuf:"bytes,2,opt,name=revisedKeyCursor,proto3" json:"revisedKeyCursor,omitempty"`
}

func (x *FetchState) Reset() {
	*x = FetchState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_federation_federation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchState) ProtoMessage() {}

func (x *FetchState) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_federation_federation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchState.ProtoReflect.Descriptor instead.
func (*FetchState) Descriptor() ([]byte, []int) {
	return file_internal_pb_federation_federation_proto_rawDescGZIP(), []int{2}
}

func (x *FetchState) GetKeyCursor() *Cursor {
	if x != nil {
		return x.KeyCursor
	}
	return nil
}

func (x *FetchState) GetRevisedKeyCursor() *Cursor {
	if x != nil {
		return x.RevisedKeyCursor
	}
	return nil
}

type Cursor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	NextToken string `protobuf:"bytes,2,opt,name=nextToken,proto3" json:"nextToken,omitempty"`
}

func (x *Cursor) Reset() {
	*x = Cursor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_federation_federation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cursor) ProtoMessage() {}

func (x *Cursor) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_federation_federation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cursor.ProtoReflect.Descriptor instead.
func (*Cursor) Descriptor() ([]byte, []int) {
	return file_internal_pb_federation_federation_proto_rawDescGZIP(), []int{3}
}

func (x *Cursor) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Cursor) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

type ExposureKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExposureKey              []byte                 `protobuf:"bytes,1,opt,name=exposureKey,proto3" json:"exposureKey,omitempty"` // required
	TransmissionRisk         int32                  `protobuf:"varint,2,opt,name=transmissionRisk,proto3" json:"transmissionRisk,omitempty"`
	IntervalNumber           int32                  `protobuf:"varint,3,opt,name=intervalNumber,proto3" json:"intervalNumber,omitempty"` // required
	IntervalCount            int32                  `protobuf:"varint,4,opt,name=intervalCount,proto3" json:"intervalCount,omitempty"`   // required
	ReportType               ExposureKey_ReportType `protobuf:"varint,5,opt,name=reportType,proto3,enum=ExposureKey_ReportType" json:"reportType,omitempty"`
	DaysSinceOnsetOfSymptoms int32                  `protobuf:"zigzag32,6,opt,name=daysSinceOnsetOfSymptoms,proto3" json:"daysSinceOnsetOfSymptoms,omitempty"` // Valid values are -14 ... 14
	HasSymptomOnset          bool                   `protobuf:"varint,7,opt,name=hasSymptomOnset,proto3" json:"hasSymptomOnset,omitempty"`                     // Used to disambiguate between 0 and missing.
	Traveler                 bool                   `protobuf:"varint,8,opt,name=traveler,proto3" json:"traveler,omitempty"`
	Regions                  []string               `protobuf:"bytes,9,rep,name=regions,proto3" json:"regions,omitempty"`
}

func (x *ExposureKey) Reset() {
	*x = ExposureKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_federation_federation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExposureKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExposureKey) ProtoMessage() {}

func (x *ExposureKey) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_federation_federation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExposureKey.ProtoReflect.Descriptor instead.
func (*ExposureKey) Descriptor() ([]byte, []int) {
	return file_internal_pb_federation_federation_proto_rawDescGZIP(), []int{4}
}

func (x *ExposureKey) GetExposureKey() []byte {
	if x != nil {
		return x.ExposureKey
	}
	return nil
}

func (x *ExposureKey) GetTransmissionRisk() int32 {
	if x != nil {
		return x.TransmissionRisk
	}
	return 0
}

func (x *ExposureKey) GetIntervalNumber() int32 {
	if x != nil {
		return x.IntervalNumber
	}
	return 0
}

func (x *ExposureKey) GetIntervalCount() int32 {
	if x != nil {
		return x.IntervalCount
	}
	return 0
}

func (x *ExposureKey) GetReportType() ExposureKey_ReportType {
	if x != nil {
		return x.ReportType
	}
	return ExposureKey_UNKNOWN
}

func (x *ExposureKey) GetDaysSinceOnsetOfSymptoms() int32 {
	if x != nil {
		return x.DaysSinceOnsetOfSymptoms
	}
	return 0
}

func (x *ExposureKey) GetHasSymptomOnset() bool {
	if x != nil {
		return x.HasSymptomOnset
	}
	return false
}

func (x *ExposureKey) GetTraveler() bool {
	if x != nil {
		return x.Traveler
	}
	return false
}

func (x *ExposureKey) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

var File_internal_pb_federation_federation_proto protoreflect.FileDescriptor

var file_internal_pb_federation_federation_proto_rawDesc = []byte{
	0x0a, 0x27, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x66, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x02, 0x0a, 0x16, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0e,
	0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x6e, 0x6c, 0x79, 0x54, 0x72, 0x61, 0x76,
	0x65, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6f, 0x6e, 0x6c,
	0x79, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x6f, 0x6e,
	0x6c, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x6f, 0x6e, 0x6c, 0x79, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x6d, 0x61, 0x78, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75,
	0x72, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x21, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xca, 0x01, 0x0a, 0x17, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x4b, 0x65,
	0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x2e, 0x0a, 0x0b, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45,
	0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x0b, 0x72, 0x65, 0x76, 0x69,
	0x73, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x33, 0x0a, 0x0e, 0x6e, 0x65, 0x78, 0x74, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0e, 0x6e, 0x65, 0x78, 0x74, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x68, 0x0a, 0x0a, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x43, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72,
	0x52, 0x09, 0x6b, 0x65, 0x79, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x33, 0x0a, 0x10, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x10,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72,
	0x22, 0x44, 0x0a, 0x06, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x78,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xfc, 0x03, 0x0a, 0x0b, 0x45, 0x78, 0x70, 0x6f, 0x73,
	0x75, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x75,
	0x72, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x65, 0x78, 0x70,
	0x6f, 0x73, 0x75, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x69, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x69, 0x73, 0x6b, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x37, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72,
	0x65, 0x4b, 0x65, 0x79, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x18, 0x64,
	0x61, 0x79, 0x73, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x4f, 0x6e, 0x73, 0x65, 0x74, 0x4f, 0x66, 0x53,
	0x79, 0x6d, 0x70, 0x74, 0x6f, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x11, 0x52, 0x18, 0x64,
	0x61, 0x79, 0x73, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x4f, 0x6e, 0x73, 0x65, 0x74, 0x4f, 0x66, 0x53,
	0x79, 0x6d, 0x70, 0x74, 0x6f, 0x6d, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x68, 0x61, 0x73, 0x53, 0x79,
	0x6d, 0x70, 0x74, 0x6f, 0x6d, 0x4f, 0x6e, 0x73, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x68, 0x61, 0x73, 0x53, 0x79, 0x6d, 0x70, 0x74, 0x6f, 0x6d, 0x4f, 0x6e, 0x73, 0x65,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x7c, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x5f,
	0x54, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52,
	0x4d, 0x45, 0x44, 0x5f, 0x43, 0x4c, 0x49, 0x4e, 0x49, 0x43, 0x41, 0x4c, 0x5f, 0x44, 0x49, 0x41,
	0x47, 0x4e, 0x4f, 0x53, 0x49, 0x53, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x45, 0x4c, 0x46,
	0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x43,
	0x55, 0x52, 0x53, 0x49, 0x56, 0x45, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x56, 0x4f,
	0x4b, 0x45, 0x44, 0x10, 0x05, 0x32, 0x4a, 0x0a, 0x0a, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x17, 0x2e, 0x46,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x53, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x2d,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62,
	0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x66, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pb_federation_federation_proto_rawDescOnce sync.Once
	file_internal_pb_federation_federation_proto_rawDescData = file_internal_pb_federation_federation_proto_rawDesc
)

func file_internal_pb_federation_federation_proto_rawDescGZIP() []byte {
	file_internal_pb_federation_federation_proto_rawDescOnce.Do(func() {
		file_internal_pb_federation_federation_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pb_federation_federation_proto_rawDescData)
	})
	return file_internal_pb_federation_federation_proto_rawDescData
}

var file_internal_pb_federation_federation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_pb_federation_federation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_pb_federation_federation_proto_goTypes = []interface{}{
	(ExposureKey_ReportType)(0),     // 0: ExposureKey.ReportType
	(*FederationFetchRequest)(nil),  // 1: FederationFetchRequest
	(*FederationFetchResponse)(nil), // 2: FederationFetchResponse
	(*FetchState)(nil),              // 3: FetchState
	(*Cursor)(nil),                  // 4: Cursor
	(*ExposureKey)(nil),             // 5: ExposureKey
}
var file_internal_pb_federation_federation_proto_depIdxs = []int32{
	3, // 0: FederationFetchRequest.state:type_name -> FetchState
	5, // 1: FederationFetchResponse.keys:type_name -> ExposureKey
	5, // 2: FederationFetchResponse.revisedKeys:type_name -> ExposureKey
	3, // 3: FederationFetchResponse.nextFetchState:type_name -> FetchState
	4, // 4: FetchState.keyCursor:type_name -> Cursor
	4, // 5: FetchState.revisedKeyCursor:type_name -> Cursor
	0, // 6: ExposureKey.reportType:type_name -> ExposureKey.ReportType
	1, // 7: Federation.Fetch:input_type -> FederationFetchRequest
	2, // 8: Federation.Fetch:output_type -> FederationFetchResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_internal_pb_federation_federation_proto_init() }
func file_internal_pb_federation_federation_proto_init() {
	if File_internal_pb_federation_federation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pb_federation_federation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederationFetchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_pb_federation_federation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederationFetchResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_pb_federation_federation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_pb_federation_federation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cursor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_pb_federation_federation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExposureKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_pb_federation_federation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_pb_federation_federation_proto_goTypes,
		DependencyIndexes: file_internal_pb_federation_federation_proto_depIdxs,
		EnumInfos:         file_internal_pb_federation_federation_proto_enumTypes,
		MessageInfos:      file_internal_pb_federation_federation_proto_msgTypes,
	}.Build()
	File_internal_pb_federation_federation_proto = out.File
	file_internal_pb_federation_federation_proto_rawDesc = nil
	file_internal_pb_federation_federation_proto_goTypes = nil
	file_internal_pb_federation_federation_proto_depIdxs = nil
}
