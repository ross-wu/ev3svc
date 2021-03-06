// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: service.proto

package ev3proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	code "google.golang.org/genproto/googleapis/rpc/code"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Port int32

const (
	Port_UNKNOWN Port = 0
	Port_IN_1    Port = 1
	Port_IN_2    Port = 2
	Port_IN_3    Port = 3
	Port_IN_4    Port = 4
	Port_OUT_A   Port = 5
	Port_OUT_B   Port = 6
	Port_OUT_C   Port = 7
	Port_OUT_D   Port = 8
)

// Enum value maps for Port.
var (
	Port_name = map[int32]string{
		0: "UNKNOWN",
		1: "IN_1",
		2: "IN_2",
		3: "IN_3",
		4: "IN_4",
		5: "OUT_A",
		6: "OUT_B",
		7: "OUT_C",
		8: "OUT_D",
	}
	Port_value = map[string]int32{
		"UNKNOWN": 0,
		"IN_1":    1,
		"IN_2":    2,
		"IN_3":    3,
		"IN_4":    4,
		"OUT_A":   5,
		"OUT_B":   6,
		"OUT_C":   7,
		"OUT_D":   8,
	}
)

func (x Port) Enum() *Port {
	p := new(Port)
	*p = x
	return p
}

func (x Port) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Port) Descriptor() protoreflect.EnumDescriptor {
	return file_service_proto_enumTypes[0].Descriptor()
}

func (Port) Type() protoreflect.EnumType {
	return &file_service_proto_enumTypes[0]
}

func (x Port) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Port.Descriptor instead.
func (Port) EnumDescriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

type InitRequest_DeviceType int32

const (
	InitRequest_UNKNOWN_DEVICE InitRequest_DeviceType = 0
	InitRequest_TACHO_MOTOR    InitRequest_DeviceType = 1
)

// Enum value maps for InitRequest_DeviceType.
var (
	InitRequest_DeviceType_name = map[int32]string{
		0: "UNKNOWN_DEVICE",
		1: "TACHO_MOTOR",
	}
	InitRequest_DeviceType_value = map[string]int32{
		"UNKNOWN_DEVICE": 0,
		"TACHO_MOTOR":    1,
	}
)

func (x InitRequest_DeviceType) Enum() *InitRequest_DeviceType {
	p := new(InitRequest_DeviceType)
	*p = x
	return p
}

func (x InitRequest_DeviceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InitRequest_DeviceType) Descriptor() protoreflect.EnumDescriptor {
	return file_service_proto_enumTypes[1].Descriptor()
}

func (InitRequest_DeviceType) Type() protoreflect.EnumType {
	return &file_service_proto_enumTypes[1]
}

func (x InitRequest_DeviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InitRequest_DeviceType.Descriptor instead.
func (InitRequest_DeviceType) EnumDescriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0, 0}
}

type Request_Attribute int32

const (
	Request_INIT              Request_Attribute = 0
	Request_ADDRESS           Request_Attribute = 1
	Request_COMMANDS          Request_Attribute = 2
	Request_COMMAND           Request_Attribute = 3
	Request_POSITION          Request_Attribute = 4
	Request_POSITION_SP       Request_Attribute = 5
	Request_MAX_SPEED         Request_Attribute = 6
	Request_SPEED             Request_Attribute = 7
	Request_SPEED_SP          Request_Attribute = 8
	Request_STOP_ACTIONS      Request_Attribute = 9
	Request_STOP_ACTION       Request_Attribute = 10
	Request_SET_STOP_ACTION   Request_Attribute = 11
	Request_WAIT_FOR_POSITION Request_Attribute = 100
)

// Enum value maps for Request_Attribute.
var (
	Request_Attribute_name = map[int32]string{
		0:   "INIT",
		1:   "ADDRESS",
		2:   "COMMANDS",
		3:   "COMMAND",
		4:   "POSITION",
		5:   "POSITION_SP",
		6:   "MAX_SPEED",
		7:   "SPEED",
		8:   "SPEED_SP",
		9:   "STOP_ACTIONS",
		10:  "STOP_ACTION",
		11:  "SET_STOP_ACTION",
		100: "WAIT_FOR_POSITION",
	}
	Request_Attribute_value = map[string]int32{
		"INIT":              0,
		"ADDRESS":           1,
		"COMMANDS":          2,
		"COMMAND":           3,
		"POSITION":          4,
		"POSITION_SP":       5,
		"MAX_SPEED":         6,
		"SPEED":             7,
		"SPEED_SP":          8,
		"STOP_ACTIONS":      9,
		"STOP_ACTION":       10,
		"SET_STOP_ACTION":   11,
		"WAIT_FOR_POSITION": 100,
	}
)

func (x Request_Attribute) Enum() *Request_Attribute {
	p := new(Request_Attribute)
	*p = x
	return p
}

func (x Request_Attribute) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Request_Attribute) Descriptor() protoreflect.EnumDescriptor {
	return file_service_proto_enumTypes[2].Descriptor()
}

func (Request_Attribute) Type() protoreflect.EnumType {
	return &file_service_proto_enumTypes[2]
}

func (x Request_Attribute) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Request_Attribute.Descriptor instead.
func (Request_Attribute) EnumDescriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2, 0}
}

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device InitRequest_DeviceType `protobuf:"varint,1,opt,name=device,proto3,enum=ev3proto.InitRequest_DeviceType" json:"device,omitempty"`
	Port   Port                   `protobuf:"varint,2,opt,name=port,proto3,enum=ev3proto.Port" json:"port,omitempty"`
	Driver string                 `protobuf:"bytes,3,opt,name=driver,proto3" json:"driver,omitempty"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *InitRequest) GetDevice() InitRequest_DeviceType {
	if x != nil {
		return x.Device
	}
	return InitRequest_UNKNOWN_DEVICE
}

func (x *InitRequest) GetPort() Port {
	if x != nil {
		return x.Port
	}
	return Port_UNKNOWN
}

func (x *InitRequest) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

type InitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error code.Code `protobuf:"varint,1,opt,name=error,proto3,enum=google.rpc.Code" json:"error,omitempty"`
	Msg   string    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"` // Error message if error != OK.
}

func (x *InitResponse) Reset() {
	*x = InitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitResponse) ProtoMessage() {}

func (x *InitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitResponse.ProtoReflect.Descriptor instead.
func (*InitResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *InitResponse) GetError() code.Code {
	if x != nil {
		return x.Error
	}
	return code.Code_OK
}

func (x *InitResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ops []*Request_Op `protobuf:"bytes,1,rep,name=ops,proto3" json:"ops,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetOps() []*Request_Op {
	if x != nil {
		return x.Ops
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Response_Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetResults() []*Response_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

type Request_Op struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port Port              `protobuf:"varint,1,opt,name=port,proto3,enum=ev3proto.Port" json:"port,omitempty"`
	Attr Request_Attribute `protobuf:"varint,2,opt,name=attr,proto3,enum=ev3proto.Request_Attribute" json:"attr,omitempty"`
	// Types that are assignable to Value:
	//	*Request_Op_Int
	//	*Request_Op_Str
	//	*Request_Op_Duration
	Value isRequest_Op_Value `protobuf_oneof:"value"`
}

func (x *Request_Op) Reset() {
	*x = Request_Op{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Op) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Op) ProtoMessage() {}

func (x *Request_Op) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Op.ProtoReflect.Descriptor instead.
func (*Request_Op) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Request_Op) GetPort() Port {
	if x != nil {
		return x.Port
	}
	return Port_UNKNOWN
}

func (x *Request_Op) GetAttr() Request_Attribute {
	if x != nil {
		return x.Attr
	}
	return Request_INIT
}

func (m *Request_Op) GetValue() isRequest_Op_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Request_Op) GetInt() int32 {
	if x, ok := x.GetValue().(*Request_Op_Int); ok {
		return x.Int
	}
	return 0
}

func (x *Request_Op) GetStr() string {
	if x, ok := x.GetValue().(*Request_Op_Str); ok {
		return x.Str
	}
	return ""
}

func (x *Request_Op) GetDuration() int64 {
	if x, ok := x.GetValue().(*Request_Op_Duration); ok {
		return x.Duration
	}
	return 0
}

type isRequest_Op_Value interface {
	isRequest_Op_Value()
}

type Request_Op_Int struct {
	Int int32 `protobuf:"varint,3,opt,name=int,proto3,oneof"`
}

type Request_Op_Str struct {
	Str string `protobuf:"bytes,4,opt,name=str,proto3,oneof"`
}

type Request_Op_Duration struct {
	Duration int64 `protobuf:"varint,5,opt,name=duration,proto3,oneof"` // time.Duration
}

func (*Request_Op_Int) isRequest_Op_Value() {}

func (*Request_Op_Str) isRequest_Op_Value() {}

func (*Request_Op_Duration) isRequest_Op_Value() {}

type Response_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error code.Code `protobuf:"varint,1,opt,name=error,proto3,enum=google.rpc.Code" json:"error,omitempty"`
	Msg   string    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"` // If error != OK, msg contains the error message.
	// Types that are assignable to Value:
	//	*Response_Result_Int
	//	*Response_Result_Str
	//	*Response_Result_Duration
	Value isResponse_Result_Value `protobuf_oneof:"value"`
	List  []string                `protobuf:"bytes,6,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *Response_Result) Reset() {
	*x = Response_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Result) ProtoMessage() {}

func (x *Response_Result) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Result.ProtoReflect.Descriptor instead.
func (*Response_Result) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Response_Result) GetError() code.Code {
	if x != nil {
		return x.Error
	}
	return code.Code_OK
}

func (x *Response_Result) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (m *Response_Result) GetValue() isResponse_Result_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Response_Result) GetInt() int32 {
	if x, ok := x.GetValue().(*Response_Result_Int); ok {
		return x.Int
	}
	return 0
}

func (x *Response_Result) GetStr() string {
	if x, ok := x.GetValue().(*Response_Result_Str); ok {
		return x.Str
	}
	return ""
}

func (x *Response_Result) GetDuration() int64 {
	if x, ok := x.GetValue().(*Response_Result_Duration); ok {
		return x.Duration
	}
	return 0
}

func (x *Response_Result) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

type isResponse_Result_Value interface {
	isResponse_Result_Value()
}

type Response_Result_Int struct {
	Int int32 `protobuf:"varint,3,opt,name=int,proto3,oneof"`
}

type Response_Result_Str struct {
	Str string `protobuf:"bytes,4,opt,name=str,proto3,oneof"`
}

type Response_Result_Duration struct {
	Duration int64 `protobuf:"varint,5,opt,name=duration,proto3,oneof"` // time.Duration
}

func (*Response_Result_Int) isResponse_Result_Value() {}

func (*Response_Result_Str) isResponse_Result_Value() {}

func (*Response_Result_Duration) isResponse_Result_Value() {}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x65, 0x76, 0x33, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb6, 0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x38, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x65, 0x76, 0x33, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x65, 0x76, 0x33, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f,
	0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x41, 0x43, 0x48,
	0x4f, 0x5f, 0x4d, 0x4f, 0x54, 0x4f, 0x52, 0x10, 0x01, 0x22, 0x48, 0x0a, 0x0c, 0x49, 0x6e, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0xb2, 0x03, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x03, 0x6f, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65,
	0x76, 0x33, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x4f, 0x70, 0x52, 0x03, 0x6f, 0x70, 0x73, 0x1a, 0xa8, 0x01, 0x0a, 0x02, 0x4f, 0x70, 0x12, 0x22,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x65,
	0x76, 0x33, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x61, 0x74, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x65, 0x76, 0x33, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x04, 0x61,
	0x74, 0x74, 0x72, 0x12, 0x12, 0x0a, 0x03, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x03, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x73, 0x74, 0x72, 0x12, 0x1c, 0x0a, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0xd3, 0x01, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x49, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x44,
	0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x4d, 0x41,
	0x4e, 0x44, 0x53, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44,
	0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04,
	0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x50, 0x10,
	0x05, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x41, 0x58, 0x5f, 0x53, 0x50, 0x45, 0x45, 0x44, 0x10, 0x06,
	0x12, 0x09, 0x0a, 0x05, 0x53, 0x50, 0x45, 0x45, 0x44, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x53,
	0x50, 0x45, 0x45, 0x44, 0x5f, 0x53, 0x50, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x4f,
	0x50, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x09, 0x12, 0x0f, 0x0a, 0x0b, 0x53,
	0x54, 0x4f, 0x50, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f,
	0x53, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x0b, 0x12, 0x15, 0x0a, 0x11, 0x57, 0x41, 0x49, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x50, 0x4f,
	0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x64, 0x22, 0xe7, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x76, 0x33, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0xa5, 0x01, 0x0a, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x12, 0x0a, 0x03, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x03,
	0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x03, 0x73, 0x74, 0x72, 0x12, 0x1c, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x2a, 0x67, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x5f, 0x31, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x5f, 0x32, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x49,
	0x4e, 0x5f, 0x33, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x5f, 0x34, 0x10, 0x04, 0x12,
	0x09, 0x0a, 0x05, 0x4f, 0x55, 0x54, 0x5f, 0x41, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x55,
	0x54, 0x5f, 0x42, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x55, 0x54, 0x5f, 0x43, 0x10, 0x07,
	0x12, 0x09, 0x0a, 0x05, 0x4f, 0x55, 0x54, 0x5f, 0x44, 0x10, 0x08, 0x32, 0x6f, 0x0a, 0x03, 0x45,
	0x76, 0x33, 0x12, 0x37, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x15, 0x2e, 0x65, 0x76, 0x33,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x65, 0x76, 0x33, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x04, 0x45,
	0x78, 0x65, 0x63, 0x12, 0x11, 0x2e, 0x65, 0x76, 0x33, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x65, 0x76, 0x33, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x73, 0x73, 0x2d,
	0x77, 0x75, 0x2f, 0x65, 0x76, 0x33, 0x73, 0x76, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x65, 0x76, 0x33, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_proto_goTypes = []interface{}{
	(Port)(0),                   // 0: ev3proto.Port
	(InitRequest_DeviceType)(0), // 1: ev3proto.InitRequest.DeviceType
	(Request_Attribute)(0),      // 2: ev3proto.Request.Attribute
	(*InitRequest)(nil),         // 3: ev3proto.InitRequest
	(*InitResponse)(nil),        // 4: ev3proto.InitResponse
	(*Request)(nil),             // 5: ev3proto.Request
	(*Response)(nil),            // 6: ev3proto.Response
	(*Request_Op)(nil),          // 7: ev3proto.Request.Op
	(*Response_Result)(nil),     // 8: ev3proto.Response.Result
	(code.Code)(0),              // 9: google.rpc.Code
}
var file_service_proto_depIdxs = []int32{
	1,  // 0: ev3proto.InitRequest.device:type_name -> ev3proto.InitRequest.DeviceType
	0,  // 1: ev3proto.InitRequest.port:type_name -> ev3proto.Port
	9,  // 2: ev3proto.InitResponse.error:type_name -> google.rpc.Code
	7,  // 3: ev3proto.Request.ops:type_name -> ev3proto.Request.Op
	8,  // 4: ev3proto.Response.results:type_name -> ev3proto.Response.Result
	0,  // 5: ev3proto.Request.Op.port:type_name -> ev3proto.Port
	2,  // 6: ev3proto.Request.Op.attr:type_name -> ev3proto.Request.Attribute
	9,  // 7: ev3proto.Response.Result.error:type_name -> google.rpc.Code
	3,  // 8: ev3proto.Ev3.Init:input_type -> ev3proto.InitRequest
	5,  // 9: ev3proto.Ev3.Exec:input_type -> ev3proto.Request
	4,  // 10: ev3proto.Ev3.Init:output_type -> ev3proto.InitResponse
	6,  // 11: ev3proto.Ev3.Exec:output_type -> ev3proto.Response
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitRequest); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitResponse); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Op); i {
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
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_Result); i {
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
	file_service_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Request_Op_Int)(nil),
		(*Request_Op_Str)(nil),
		(*Request_Op_Duration)(nil),
	}
	file_service_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Response_Result_Int)(nil),
		(*Response_Result_Str)(nil),
		(*Response_Result_Duration)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		EnumInfos:         file_service_proto_enumTypes,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// Ev3Client is the client API for Ev3 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Ev3Client interface {
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error)
	Exec(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type ev3Client struct {
	cc grpc.ClientConnInterface
}

func NewEv3Client(cc grpc.ClientConnInterface) Ev3Client {
	return &ev3Client{cc}
}

func (c *ev3Client) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error) {
	out := new(InitResponse)
	err := c.cc.Invoke(ctx, "/ev3proto.Ev3/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ev3Client) Exec(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ev3proto.Ev3/Exec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Ev3Server is the server API for Ev3 service.
type Ev3Server interface {
	Init(context.Context, *InitRequest) (*InitResponse, error)
	Exec(context.Context, *Request) (*Response, error)
}

// UnimplementedEv3Server can be embedded to have forward compatible implementations.
type UnimplementedEv3Server struct {
}

func (*UnimplementedEv3Server) Init(context.Context, *InitRequest) (*InitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (*UnimplementedEv3Server) Exec(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}

func RegisterEv3Server(s *grpc.Server, srv Ev3Server) {
	s.RegisterService(&_Ev3_serviceDesc, srv)
}

func _Ev3_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Ev3Server).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ev3proto.Ev3/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Ev3Server).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ev3_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Ev3Server).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ev3proto.Ev3/Exec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Ev3Server).Exec(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ev3_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ev3proto.Ev3",
	HandlerType: (*Ev3Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _Ev3_Init_Handler,
		},
		{
			MethodName: "Exec",
			Handler:    _Ev3_Exec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
