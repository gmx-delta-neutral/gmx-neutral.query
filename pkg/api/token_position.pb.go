// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
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

type TokenExposure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount   string  `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Leverage float64 `protobuf:"fixed64,2,opt,name=leverage,proto3" json:"leverage,omitempty"`
}

func (x *TokenExposure) Reset() {
	*x = TokenExposure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_position_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenExposure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenExposure) ProtoMessage() {}

func (x *TokenExposure) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TokenExposure.ProtoReflect.Descriptor instead.
func (*TokenExposure) Descriptor() ([]byte, []int) {
	return file_token_position_proto_rawDescGZIP(), []int{0}
}

func (x *TokenExposure) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *TokenExposure) GetLeverage() float64 {
	if x != nil {
		return x.Leverage
	}
	return 0
}

type TokenPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenAddress  string           `protobuf:"bytes,1,opt,name=tokenAddress,proto3" json:"tokenAddress,omitempty"`
	Symbol        string           `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Amount        string           `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Worth         string           `protobuf:"bytes,4,opt,name=worth,proto3" json:"worth,omitempty"`
	Pnl           string           `protobuf:"bytes,5,opt,name=pnl,proto3" json:"pnl,omitempty"`
	PnlPercentage string           `protobuf:"bytes,6,opt,name=pnlPercentage,proto3" json:"pnlPercentage,omitempty"`
	CostBasis     string           `protobuf:"bytes,7,opt,name=costBasis,proto3" json:"costBasis,omitempty"`
	Decimals      int32            `protobuf:"varint,8,opt,name=decimals,proto3" json:"decimals,omitempty"`
	TokenExposure []*TokenExposure `protobuf:"bytes,9,rep,name=tokenExposure,proto3" json:"tokenExposure,omitempty"`
}

func (x *TokenPosition) Reset() {
	*x = TokenPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_position_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenPosition) ProtoMessage() {}

func (x *TokenPosition) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TokenPosition.ProtoReflect.Descriptor instead.
func (*TokenPosition) Descriptor() ([]byte, []int) {
	return file_token_position_proto_rawDescGZIP(), []int{1}
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

func (x *TokenPosition) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *TokenPosition) GetWorth() string {
	if x != nil {
		return x.Worth
	}
	return ""
}

func (x *TokenPosition) GetPnl() string {
	if x != nil {
		return x.Pnl
	}
	return ""
}

func (x *TokenPosition) GetPnlPercentage() string {
	if x != nil {
		return x.PnlPercentage
	}
	return ""
}

func (x *TokenPosition) GetCostBasis() string {
	if x != nil {
		return x.CostBasis
	}
	return ""
}

func (x *TokenPosition) GetDecimals() int32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *TokenPosition) GetTokenExposure() []*TokenExposure {
	if x != nil {
		return x.TokenExposure
	}
	return nil
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
		mi := &file_token_position_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenPositionRequest) ProtoMessage() {}

func (x *GetTokenPositionRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetTokenPositionRequest.ProtoReflect.Descriptor instead.
func (*GetTokenPositionRequest) Descriptor() ([]byte, []int) {
	return file_token_position_proto_rawDescGZIP(), []int{2}
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
		mi := &file_token_position_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenPositionResponse) ProtoMessage() {}

func (x *GetTokenPositionResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetTokenPositionResponse.ProtoReflect.Descriptor instead.
func (*GetTokenPositionResponse) Descriptor() ([]byte, []int) {
	return file_token_position_proto_rawDescGZIP(), []int{3}
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
		mi := &file_token_position_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenPositionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenPositionsRequest) ProtoMessage() {}

func (x *GetTokenPositionsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetTokenPositionsRequest.ProtoReflect.Descriptor instead.
func (*GetTokenPositionsRequest) Descriptor() ([]byte, []int) {
	return file_token_position_proto_rawDescGZIP(), []int{4}
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
		mi := &file_token_position_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenPositionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenPositionsResponse) ProtoMessage() {}

func (x *GetTokenPositionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_position_proto_msgTypes[5]
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
	return file_token_position_proto_rawDescGZIP(), []int{5}
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
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6d, 0x78, 0x2e, 0x6e, 0x65, 0x75, 0x74,
	0x72, 0x61, 0x6c, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x43, 0x0a,
	0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x22, 0xb7, 0x02, 0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x6f, 0x72, 0x74,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x77, 0x6f, 0x72, 0x74, 0x68, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x6e, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x6e, 0x6c,
	0x12, 0x24, 0x0a, 0x0d, 0x70, 0x6e, 0x6c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x6e, 0x6c, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x42, 0x61,
	0x73, 0x69, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x42,
	0x61, 0x73, 0x69, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73,
	0x12, 0x4a, 0x0a, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72,
	0x65, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6d, 0x78, 0x2e, 0x6e, 0x65,
	0x75, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x52, 0x0d, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x22, 0x3d, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x66, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x67, 0x6d, 0x78, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x3e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x69, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4c, 0x0a, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6d, 0x78, 0x2e, 0x6e,
	0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x82,
	0x02, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x78, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x2e, 0x67, 0x6d, 0x78, 0x2e, 0x6e, 0x65,
	0x75, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67, 0x6d, 0x78, 0x2e, 0x6e,
	0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x75, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2e, 0x2e, 0x67, 0x6d, 0x78, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2f, 0x2e, 0x67, 0x6d, 0x78, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
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

var file_token_position_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_token_position_proto_goTypes = []interface{}{
	(*TokenExposure)(nil),             // 0: gmx.neutral.query.api.TokenExposure
	(*TokenPosition)(nil),             // 1: gmx.neutral.query.api.TokenPosition
	(*GetTokenPositionRequest)(nil),   // 2: gmx.neutral.query.api.GetTokenPositionRequest
	(*GetTokenPositionResponse)(nil),  // 3: gmx.neutral.query.api.GetTokenPositionResponse
	(*GetTokenPositionsRequest)(nil),  // 4: gmx.neutral.query.api.GetTokenPositionsRequest
	(*GetTokenPositionsResponse)(nil), // 5: gmx.neutral.query.api.GetTokenPositionsResponse
}
var file_token_position_proto_depIdxs = []int32{
	0, // 0: gmx.neutral.query.api.TokenPosition.tokenExposure:type_name -> gmx.neutral.query.api.TokenExposure
	1, // 1: gmx.neutral.query.api.GetTokenPositionResponse.tokenPosition:type_name -> gmx.neutral.query.api.TokenPosition
	1, // 2: gmx.neutral.query.api.GetTokenPositionsResponse.tokenPositions:type_name -> gmx.neutral.query.api.TokenPosition
	4, // 3: gmx.neutral.query.api.PositionService.GetTokenPositions:input_type -> gmx.neutral.query.api.GetTokenPositionsRequest
	2, // 4: gmx.neutral.query.api.PositionService.GetTokenPosition:input_type -> gmx.neutral.query.api.GetTokenPositionRequest
	5, // 5: gmx.neutral.query.api.PositionService.GetTokenPositions:output_type -> gmx.neutral.query.api.GetTokenPositionsResponse
	3, // 6: gmx.neutral.query.api.PositionService.GetTokenPosition:output_type -> gmx.neutral.query.api.GetTokenPositionResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_token_position_proto_init() }
func file_token_position_proto_init() {
	if File_token_position_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_token_position_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenExposure); i {
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
		file_token_position_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_token_position_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_token_position_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_token_position_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   6,
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
