// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: token_position.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TokenPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenAddress  string `protobuf:"bytes,1,opt,name=tokenAddress,proto3" json:"tokenAddress,omitempty"`
	Symbol        string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Amount        []byte `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Worth         []byte `protobuf:"bytes,4,opt,name=worth,proto3" json:"worth,omitempty"`
	Pnl           []byte `protobuf:"bytes,5,opt,name=pnl,proto3" json:"pnl,omitempty"`
	PnlPercentage string `protobuf:"bytes,6,opt,name=pnlPercentage,proto3" json:"pnlPercentage,omitempty"`
}

func (x *TokenPosition) Reset() {
	*x = TokenPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_position_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenPosition) ProtoMessage() {}

func (x *TokenPosition) ProtoReflect() protoreflect.Message {
	mi := &file_token_position_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenPosition.ProtoReflect.Descriptor instead.
func (*TokenPosition) Descriptor() ([]byte, []int) {
	return file_token_position_proto_rawDescGZIP(), []int{0}
}

func (x *TokenPosition) GetTokenAddress() string {
	if x != nil {
		return x.TokenAddress
	}
	return ""
}

func (x *TokenPosition) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *TokenPosition) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *TokenPosition) GetWorth() []byte {
	if x != nil {
		return x.Worth
	}
	return nil
}

func (x *TokenPosition) GetPnl() []byte {
	if x != nil {
		return x.Pnl
	}
	return nil
}

func (x *TokenPosition) GetPnlPercentage() string {
	if x != nil {
		return x.PnlPercentage
	}
	return ""
}

type GetTokenPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenAddress string `protobuf:"bytes,1,opt,name=tokenAddress,proto3" json:"tokenAddress,omitempty"`
}

func (x *GetTokenPositionRequest) Reset() {
	*x = GetTokenPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_position_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenPositionRequest) ProtoMessage() {}

func (x *GetTokenPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_position_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenPositionRequest.ProtoReflect.Descriptor instead.
func (*GetTokenPositionRequest) Descriptor() ([]byte, []int) {
	return file_token_position_proto_rawDescGZIP(), []int{1}
}

func (x *GetTokenPositionRequest) GetTokenAddress() string {
	if x != nil {
		return x.TokenAddress
	}
	return ""
}

type GetTokenPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenPosition *TokenPosition `protobuf:"bytes,1,opt,name=tokenPosition,proto3" json:"tokenPosition,omitempty"`
}

func (x *GetTokenPositionResponse) Reset() {
	*x = GetTokenPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_position_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenPositionResponse) ProtoMessage() {}

func (x *GetTokenPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_position_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenPositionResponse.ProtoReflect.Descriptor instead.
func (*GetTokenPositionResponse) Descriptor() ([]byte, []int) {
	return file_token_position_proto_rawDescGZIP(), []int{2}
}

func (x *GetTokenPositionResponse) GetTokenPosition() *TokenPosition {
	if x != nil {
		return x.TokenPosition
	}
	return nil
}

type GetTokenPositionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenAddress string `protobuf:"bytes,1,opt,name=tokenAddress,proto3" json:"tokenAddress,omitempty"`
}

func (x *GetTokenPositionsRequest) Reset() {
	*x = GetTokenPositionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_position_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenPositionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenPositionsRequest) ProtoMessage() {}

func (x *GetTokenPositionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_position_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenPositionsRequest.ProtoReflect.Descriptor instead.
func (*GetTokenPositionsRequest) Descriptor() ([]byte, []int) {
	return file_token_position_proto_rawDescGZIP(), []int{3}
}

func (x *GetTokenPositionsRequest) GetTokenAddress() string {
	if x != nil {
		return x.TokenAddress
	}
	return ""
}

type GetTokenPositionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenPositions []*TokenPosition `protobuf:"bytes,1,rep,name=tokenPositions,proto3" json:"tokenPositions,omitempty"`
}

func (x *GetTokenPositionsResponse) Reset() {
	*x = GetTokenPositionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_position_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenPositionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenPositionsResponse) ProtoMessage() {}

func (x *GetTokenPositionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_position_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenPositionsResponse.ProtoReflect.Descriptor instead.
func (*GetTokenPositionsResponse) Descriptor() ([]byte, []int) {
	return file_token_position_proto_rawDescGZIP(), []int{4}
}

func (x *GetTokenPositionsResponse) GetTokenPositions() []*TokenPosition {
	if x != nil {
		return x.TokenPositions
	}
	return nil
}

var File_token_position_proto protoreflect.FileDescriptor

var file_token_position_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb1, 0x01, 0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x77, 0x6f, 0x72, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x77, 0x6f, 0x72,
	0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6e, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x70, 0x6e, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x6e, 0x6c, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x6e, 0x6c,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x22, 0x3d, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5f, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3e, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x62, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xe6,
	0x01, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x6a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_position_proto_rawDescOnce sync.Once
	file_token_position_proto_rawDescData = file_token_position_proto_rawDesc
)

func file_token_position_proto_rawDescGZIP() []byte {
	file_token_position_proto_rawDescOnce.Do(func() {
		file_token_position_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_position_proto_rawDescData)
	})
	return file_token_position_proto_rawDescData
}

var file_token_position_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_token_position_proto_goTypes = []interface{}{
	(*TokenPosition)(nil),             // 0: token_position.TokenPosition
	(*GetTokenPositionRequest)(nil),   // 1: token_position.GetTokenPositionRequest
	(*GetTokenPositionResponse)(nil),  // 2: token_position.GetTokenPositionResponse
	(*GetTokenPositionsRequest)(nil),  // 3: token_position.GetTokenPositionsRequest
	(*GetTokenPositionsResponse)(nil), // 4: token_position.GetTokenPositionsResponse
}
var file_token_position_proto_depIdxs = []int32{
	0, // 0: token_position.GetTokenPositionResponse.tokenPosition:type_name -> token_position.TokenPosition
	0, // 1: token_position.GetTokenPositionsResponse.tokenPositions:type_name -> token_position.TokenPosition
	3, // 2: token_position.PositionService.GetTokenPositions:input_type -> token_position.GetTokenPositionsRequest
	1, // 3: token_position.PositionService.GetTokenPosition:input_type -> token_position.GetTokenPositionRequest
	4, // 4: token_position.PositionService.GetTokenPositions:output_type -> token_position.GetTokenPositionsResponse
	2, // 5: token_position.PositionService.GetTokenPosition:output_type -> token_position.GetTokenPositionResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_token_position_proto_init() }
func file_token_position_proto_init() {
	if File_token_position_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_token_position_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenPosition); i {
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
		file_token_position_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenPositionRequest); i {
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
		file_token_position_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenPositionResponse); i {
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
		file_token_position_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenPositionsRequest); i {
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
		file_token_position_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenPositionsResponse); i {
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
			RawDescriptor: file_token_position_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_token_position_proto_goTypes,
		DependencyIndexes: file_token_position_proto_depIdxs,
		MessageInfos:      file_token_position_proto_msgTypes,
	}.Build()
	File_token_position_proto = out.File
	file_token_position_proto_rawDesc = nil
	file_token_position_proto_goTypes = nil
	file_token_position_proto_depIdxs = nil
}
