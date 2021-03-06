// Code generated by protoc-gen-go. DO NOT EDIT.
// source: executor.proto

package apis

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Command struct {
	Path                 []byte   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Args                 [][]byte `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	Env                  [][]byte `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty"`
	Dir                  []byte   `protobuf:"bytes,4,opt,name=dir,proto3" json:"dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{0}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *Command) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Command) GetEnv() [][]byte {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Command) GetDir() []byte {
	if m != nil {
		return m.Dir
	}
	return nil
}

type Input struct {
	Sn                   uint32   `protobuf:"varint,1,opt,name=sn,proto3" json:"sn,omitempty"`
	Input                []byte   `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{1}
}

func (m *Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Input.Unmarshal(m, b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Input.Marshal(b, m, deterministic)
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return xxx_messageInfo_Input.Size(m)
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetSn() uint32 {
	if m != nil {
		return m.Sn
	}
	return 0
}

func (m *Input) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

type Stdout struct {
	Stdout               []byte   `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Closed               bool     `protobuf:"varint,2,opt,name=closed,proto3" json:"closed,omitempty"`
	RuntimeError         []byte   `protobuf:"bytes,3,opt,name=runtime_error,json=runtimeError,proto3" json:"runtime_error,omitempty"`
	Start                bool     `protobuf:"varint,4,opt,name=start,proto3" json:"start,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stdout) Reset()         { *m = Stdout{} }
func (m *Stdout) String() string { return proto.CompactTextString(m) }
func (*Stdout) ProtoMessage()    {}
func (*Stdout) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{2}
}

func (m *Stdout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stdout.Unmarshal(m, b)
}
func (m *Stdout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stdout.Marshal(b, m, deterministic)
}
func (m *Stdout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stdout.Merge(m, src)
}
func (m *Stdout) XXX_Size() int {
	return xxx_messageInfo_Stdout.Size(m)
}
func (m *Stdout) XXX_DiscardUnknown() {
	xxx_messageInfo_Stdout.DiscardUnknown(m)
}

var xxx_messageInfo_Stdout proto.InternalMessageInfo

func (m *Stdout) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *Stdout) GetClosed() bool {
	if m != nil {
		return m.Closed
	}
	return false
}

func (m *Stdout) GetRuntimeError() []byte {
	if m != nil {
		return m.RuntimeError
	}
	return nil
}

func (m *Stdout) GetStart() bool {
	if m != nil {
		return m.Start
	}
	return false
}

type Stderr struct {
	Stderr               []byte   `protobuf:"bytes,1,opt,name=stderr,proto3" json:"stderr,omitempty"`
	Closed               bool     `protobuf:"varint,2,opt,name=closed,proto3" json:"closed,omitempty"`
	RuntimeError         []byte   `protobuf:"bytes,3,opt,name=runtime_error,json=runtimeError,proto3" json:"runtime_error,omitempty"`
	Start                bool     `protobuf:"varint,4,opt,name=start,proto3" json:"start,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stderr) Reset()         { *m = Stderr{} }
func (m *Stderr) String() string { return proto.CompactTextString(m) }
func (*Stderr) ProtoMessage()    {}
func (*Stderr) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{3}
}

func (m *Stderr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stderr.Unmarshal(m, b)
}
func (m *Stderr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stderr.Marshal(b, m, deterministic)
}
func (m *Stderr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stderr.Merge(m, src)
}
func (m *Stderr) XXX_Size() int {
	return xxx_messageInfo_Stderr.Size(m)
}
func (m *Stderr) XXX_DiscardUnknown() {
	xxx_messageInfo_Stderr.DiscardUnknown(m)
}

var xxx_messageInfo_Stderr proto.InternalMessageInfo

func (m *Stderr) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func (m *Stderr) GetClosed() bool {
	if m != nil {
		return m.Closed
	}
	return false
}

func (m *Stderr) GetRuntimeError() []byte {
	if m != nil {
		return m.RuntimeError
	}
	return nil
}

func (m *Stderr) GetStart() bool {
	if m != nil {
		return m.Start
	}
	return false
}

type StartResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                []byte   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartResponse) Reset()         { *m = StartResponse{} }
func (m *StartResponse) String() string { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()    {}
func (*StartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{4}
}

func (m *StartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartResponse.Unmarshal(m, b)
}
func (m *StartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartResponse.Marshal(b, m, deterministic)
}
func (m *StartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartResponse.Merge(m, src)
}
func (m *StartResponse) XXX_Size() int {
	return xxx_messageInfo_StartResponse.Size(m)
}
func (m *StartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartResponse proto.InternalMessageInfo

func (m *StartResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *StartResponse) GetError() []byte {
	if m != nil {
		return m.Error
	}
	return nil
}

type WaitCommand struct {
	Sn                   uint32   `protobuf:"varint,1,opt,name=sn,proto3" json:"sn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WaitCommand) Reset()         { *m = WaitCommand{} }
func (m *WaitCommand) String() string { return proto.CompactTextString(m) }
func (*WaitCommand) ProtoMessage()    {}
func (*WaitCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{5}
}

func (m *WaitCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaitCommand.Unmarshal(m, b)
}
func (m *WaitCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaitCommand.Marshal(b, m, deterministic)
}
func (m *WaitCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaitCommand.Merge(m, src)
}
func (m *WaitCommand) XXX_Size() int {
	return xxx_messageInfo_WaitCommand.Size(m)
}
func (m *WaitCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_WaitCommand.DiscardUnknown(m)
}

var xxx_messageInfo_WaitCommand proto.InternalMessageInfo

func (m *WaitCommand) GetSn() uint32 {
	if m != nil {
		return m.Sn
	}
	return 0
}

type WaitResponse struct {
	ExitStatus           uint32   `protobuf:"varint,1,opt,name=exit_status,json=exitStatus,proto3" json:"exit_status,omitempty"`
	ErrContent           []byte   `protobuf:"bytes,2,opt,name=err_content,json=errContent,proto3" json:"err_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WaitResponse) Reset()         { *m = WaitResponse{} }
func (m *WaitResponse) String() string { return proto.CompactTextString(m) }
func (*WaitResponse) ProtoMessage()    {}
func (*WaitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{6}
}

func (m *WaitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaitResponse.Unmarshal(m, b)
}
func (m *WaitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaitResponse.Marshal(b, m, deterministic)
}
func (m *WaitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaitResponse.Merge(m, src)
}
func (m *WaitResponse) XXX_Size() int {
	return xxx_messageInfo_WaitResponse.Size(m)
}
func (m *WaitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WaitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WaitResponse proto.InternalMessageInfo

func (m *WaitResponse) GetExitStatus() uint32 {
	if m != nil {
		return m.ExitStatus
	}
	return 0
}

func (m *WaitResponse) GetErrContent() []byte {
	if m != nil {
		return m.ErrContent
	}
	return nil
}

type Sn struct {
	Sn                   uint32   `protobuf:"varint,1,opt,name=sn,proto3" json:"sn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sn) Reset()         { *m = Sn{} }
func (m *Sn) String() string { return proto.CompactTextString(m) }
func (*Sn) ProtoMessage()    {}
func (*Sn) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{7}
}

func (m *Sn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sn.Unmarshal(m, b)
}
func (m *Sn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sn.Marshal(b, m, deterministic)
}
func (m *Sn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sn.Merge(m, src)
}
func (m *Sn) XXX_Size() int {
	return xxx_messageInfo_Sn.Size(m)
}
func (m *Sn) XXX_DiscardUnknown() {
	xxx_messageInfo_Sn.DiscardUnknown(m)
}

var xxx_messageInfo_Sn proto.InternalMessageInfo

func (m *Sn) GetSn() uint32 {
	if m != nil {
		return m.Sn
	}
	return 0
}

type StartInput struct {
	Sn                   uint32   `protobuf:"varint,1,opt,name=sn,proto3" json:"sn,omitempty"`
	HasStdin             bool     `protobuf:"varint,2,opt,name=has_stdin,json=hasStdin,proto3" json:"has_stdin,omitempty"`
	HasStdout            bool     `protobuf:"varint,3,opt,name=has_stdout,json=hasStdout,proto3" json:"has_stdout,omitempty"`
	HasStderr            bool     `protobuf:"varint,4,opt,name=has_stderr,json=hasStderr,proto3" json:"has_stderr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartInput) Reset()         { *m = StartInput{} }
func (m *StartInput) String() string { return proto.CompactTextString(m) }
func (*StartInput) ProtoMessage()    {}
func (*StartInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{8}
}

func (m *StartInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartInput.Unmarshal(m, b)
}
func (m *StartInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartInput.Marshal(b, m, deterministic)
}
func (m *StartInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartInput.Merge(m, src)
}
func (m *StartInput) XXX_Size() int {
	return xxx_messageInfo_StartInput.Size(m)
}
func (m *StartInput) XXX_DiscardUnknown() {
	xxx_messageInfo_StartInput.DiscardUnknown(m)
}

var xxx_messageInfo_StartInput proto.InternalMessageInfo

func (m *StartInput) GetSn() uint32 {
	if m != nil {
		return m.Sn
	}
	return 0
}

func (m *StartInput) GetHasStdin() bool {
	if m != nil {
		return m.HasStdin
	}
	return false
}

func (m *StartInput) GetHasStdout() bool {
	if m != nil {
		return m.HasStdout
	}
	return false
}

func (m *StartInput) GetHasStderr() bool {
	if m != nil {
		return m.HasStderr
	}
	return false
}

type Error struct {
	Error                []byte   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{9}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetError() []byte {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*Command)(nil), "apis.Command")
	proto.RegisterType((*Input)(nil), "apis.Input")
	proto.RegisterType((*Stdout)(nil), "apis.Stdout")
	proto.RegisterType((*Stderr)(nil), "apis.Stderr")
	proto.RegisterType((*StartResponse)(nil), "apis.StartResponse")
	proto.RegisterType((*WaitCommand)(nil), "apis.WaitCommand")
	proto.RegisterType((*WaitResponse)(nil), "apis.WaitResponse")
	proto.RegisterType((*Sn)(nil), "apis.Sn")
	proto.RegisterType((*StartInput)(nil), "apis.StartInput")
	proto.RegisterType((*Error)(nil), "apis.Error")
}

func init() { proto.RegisterFile("executor.proto", fileDescriptor_12d1cdcda51e000f) }

var fileDescriptor_12d1cdcda51e000f = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x6b, 0xdb, 0x4e,
	0x10, 0x45, 0x7f, 0xec, 0xc8, 0x23, 0x3b, 0x84, 0xfd, 0x85, 0x1f, 0xc2, 0xc5, 0x34, 0xa8, 0xa5,
	0xc9, 0xa5, 0x26, 0xb4, 0x1f, 0xa0, 0x87, 0x90, 0x42, 0xe9, 0xa5, 0x48, 0x94, 0x1e, 0x8d, 0x2a,
	0x2d, 0xb5, 0xc0, 0x5e, 0x89, 0x99, 0x55, 0x9b, 0xcf, 0xd3, 0x4f, 0x5a, 0x66, 0x76, 0xe5, 0x2a,
	0x6d, 0x8e, 0xbd, 0xcd, 0x7b, 0x6f, 0x76, 0x66, 0xf7, 0x3d, 0x21, 0x38, 0xd7, 0x0f, 0xba, 0x1e,
	0x6c, 0x87, 0xdb, 0x1e, 0x3b, 0xdb, 0xa9, 0xb8, 0xea, 0x5b, 0xca, 0x3f, 0xc3, 0xd9, 0x5d, 0x77,
	0x3c, 0x56, 0xa6, 0x51, 0x0a, 0xe2, 0xbe, 0xb2, 0xfb, 0x2c, 0xb8, 0x0a, 0x6e, 0x96, 0x85, 0xd4,
	0xcc, 0x55, 0xf8, 0x8d, 0xb2, 0xf0, 0x2a, 0x62, 0x8e, 0x6b, 0x75, 0x01, 0x91, 0x36, 0xdf, 0xb3,
	0x48, 0x28, 0x2e, 0x99, 0x69, 0x5a, 0xcc, 0x62, 0x39, 0xc8, 0x65, 0xfe, 0x1a, 0x66, 0x1f, 0x4c,
	0x3f, 0x58, 0x75, 0x0e, 0x21, 0x19, 0x19, 0xb9, 0x2a, 0x42, 0x32, 0xea, 0x12, 0x66, 0x2d, 0x0b,
	0x59, 0x28, 0xcd, 0x0e, 0xe4, 0x04, 0xf3, 0xd2, 0x36, 0xdd, 0x60, 0xd5, 0xff, 0x30, 0x27, 0xa9,
	0xfc, 0x35, 0x3c, 0x62, 0xbe, 0x3e, 0x74, 0xa4, 0x1b, 0x39, 0x98, 0x14, 0x1e, 0xa9, 0x17, 0xb0,
	0xc2, 0xc1, 0xd8, 0xf6, 0xa8, 0x77, 0x1a, 0xb1, 0xc3, 0x2c, 0x92, 0x63, 0x4b, 0x4f, 0xde, 0x33,
	0xc7, 0x4b, 0xc9, 0x56, 0x68, 0xe5, 0x86, 0x49, 0xe1, 0x80, 0x5f, 0xaa, 0x11, 0xfd, 0x52, 0x8d,
	0x38, 0x59, 0xea, 0xf9, 0x7f, 0xbd, 0xf4, 0x1d, 0xac, 0x4a, 0x2e, 0x0a, 0x4d, 0x7d, 0x67, 0x48,
	0xab, 0x0c, 0xce, 0x68, 0xa8, 0x6b, 0x4d, 0x24, 0xcb, 0x93, 0x62, 0x84, 0x3c, 0xc0, 0x4d, 0xf7,
	0x56, 0x09, 0xc8, 0x37, 0x90, 0x7e, 0xa9, 0x5a, 0x3b, 0x86, 0xf6, 0x87, 0xbf, 0xf9, 0x27, 0x58,
	0xb2, 0x7c, 0x1a, 0xff, 0x1c, 0x52, 0xfd, 0xd0, 0xda, 0x1d, 0xd9, 0xca, 0x0e, 0xe4, 0x1b, 0x81,
	0xa9, 0x52, 0x18, 0x69, 0x40, 0xdc, 0xd5, 0x9d, 0xb1, 0xda, 0x8c, 0xb1, 0x80, 0x46, 0xbc, 0x73,
	0x4c, 0x7e, 0x09, 0x61, 0x69, 0xfe, 0xda, 0xf3, 0x03, 0x40, 0xde, 0xf1, 0x74, 0xca, 0xcf, 0x60,
	0xb1, 0xaf, 0x68, 0x47, 0xb6, 0x69, 0x8d, 0xf7, 0x2e, 0xd9, 0x57, 0x54, 0x32, 0x56, 0x1b, 0x00,
	0x2f, 0x72, 0xcc, 0x91, 0xa8, 0x0b, 0xa7, 0x72, 0xd2, 0xbf, 0x65, 0x0e, 0x24, 0x9e, 0xca, 0x1a,
	0xf9, 0xfd, 0xb3, 0x93, 0xbf, 0xce, 0x9e, 0x60, 0x62, 0xcf, 0x9b, 0x9f, 0x21, 0x24, 0xf7, 0xfe,
	0x43, 0x57, 0xd7, 0xb0, 0x28, 0xb5, 0x69, 0xdc, 0x1d, 0xd3, 0x2d, 0x7f, 0xf0, 0x5b, 0x01, 0x6b,
	0x0f, 0x64, 0xd2, 0x4d, 0xa0, 0xae, 0x21, 0x7d, 0xaf, 0x6d, 0xbd, 0xf7, 0x57, 0x48, 0x9c, 0x5a,
	0x9a, 0xf5, 0xd2, 0x57, 0xc2, 0xdf, 0x3e, 0x6a, 0xe4, 0x0f, 0xe4, 0xa9, 0x46, 0x8d, 0x78, 0x1b,
	0xa8, 0x2d, 0xcc, 0xc4, 0x1f, 0x75, 0x31, 0x0a, 0xa3, 0x59, 0xeb, 0xff, 0x26, 0xcc, 0x29, 0xa7,
	0x97, 0x10, 0x73, 0x6e, 0x93, 0x89, 0xca, 0x55, 0x8f, 0xd2, 0x7c, 0x05, 0x29, 0x3f, 0x6e, 0x0c,
	0x7f, 0xe5, 0x5a, 0x3c, 0x5c, 0x9f, 0xce, 0xaa, 0x0d, 0xc4, 0x1f, 0xdb, 0xc3, 0x61, 0x32, 0x6d,
	0xfa, 0xe0, 0xaf, 0x73, 0xf9, 0x03, 0xbc, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0x60, 0x42, 0x00,
	0x66, 0x13, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExecutorClient is the client API for Executor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExecutorClient interface {
	SendInput(ctx context.Context, opts ...grpc.CallOption) (Executor_SendInputClient, error)
	FetchStdout(ctx context.Context, in *Sn, opts ...grpc.CallOption) (Executor_FetchStdoutClient, error)
	FetchStderr(ctx context.Context, in *Sn, opts ...grpc.CallOption) (Executor_FetchStderrClient, error)
	Start(ctx context.Context, in *StartInput, opts ...grpc.CallOption) (*StartResponse, error)
	Wait(ctx context.Context, in *Sn, opts ...grpc.CallOption) (*WaitResponse, error)
	ExecCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Sn, error)
	Kill(ctx context.Context, in *Sn, opts ...grpc.CallOption) (*Error, error)
}

type executorClient struct {
	cc *grpc.ClientConn
}

func NewExecutorClient(cc *grpc.ClientConn) ExecutorClient {
	return &executorClient{cc}
}

func (c *executorClient) SendInput(ctx context.Context, opts ...grpc.CallOption) (Executor_SendInputClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Executor_serviceDesc.Streams[0], "/apis.Executor/SendInput", opts...)
	if err != nil {
		return nil, err
	}
	x := &executorSendInputClient{stream}
	return x, nil
}

type Executor_SendInputClient interface {
	Send(*Input) error
	CloseAndRecv() (*Error, error)
	grpc.ClientStream
}

type executorSendInputClient struct {
	grpc.ClientStream
}

func (x *executorSendInputClient) Send(m *Input) error {
	return x.ClientStream.SendMsg(m)
}

func (x *executorSendInputClient) CloseAndRecv() (*Error, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Error)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *executorClient) FetchStdout(ctx context.Context, in *Sn, opts ...grpc.CallOption) (Executor_FetchStdoutClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Executor_serviceDesc.Streams[1], "/apis.Executor/FetchStdout", opts...)
	if err != nil {
		return nil, err
	}
	x := &executorFetchStdoutClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Executor_FetchStdoutClient interface {
	Recv() (*Stdout, error)
	grpc.ClientStream
}

type executorFetchStdoutClient struct {
	grpc.ClientStream
}

func (x *executorFetchStdoutClient) Recv() (*Stdout, error) {
	m := new(Stdout)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *executorClient) FetchStderr(ctx context.Context, in *Sn, opts ...grpc.CallOption) (Executor_FetchStderrClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Executor_serviceDesc.Streams[2], "/apis.Executor/FetchStderr", opts...)
	if err != nil {
		return nil, err
	}
	x := &executorFetchStderrClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Executor_FetchStderrClient interface {
	Recv() (*Stderr, error)
	grpc.ClientStream
}

type executorFetchStderrClient struct {
	grpc.ClientStream
}

func (x *executorFetchStderrClient) Recv() (*Stderr, error) {
	m := new(Stderr)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *executorClient) Start(ctx context.Context, in *StartInput, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, "/apis.Executor/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) Wait(ctx context.Context, in *Sn, opts ...grpc.CallOption) (*WaitResponse, error) {
	out := new(WaitResponse)
	err := c.cc.Invoke(ctx, "/apis.Executor/Wait", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) ExecCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Sn, error) {
	out := new(Sn)
	err := c.cc.Invoke(ctx, "/apis.Executor/ExecCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) Kill(ctx context.Context, in *Sn, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/apis.Executor/Kill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutorServer is the server API for Executor service.
type ExecutorServer interface {
	SendInput(Executor_SendInputServer) error
	FetchStdout(*Sn, Executor_FetchStdoutServer) error
	FetchStderr(*Sn, Executor_FetchStderrServer) error
	Start(context.Context, *StartInput) (*StartResponse, error)
	Wait(context.Context, *Sn) (*WaitResponse, error)
	ExecCommand(context.Context, *Command) (*Sn, error)
	Kill(context.Context, *Sn) (*Error, error)
}

// UnimplementedExecutorServer can be embedded to have forward compatible implementations.
type UnimplementedExecutorServer struct {
}

func (*UnimplementedExecutorServer) SendInput(srv Executor_SendInputServer) error {
	return status.Errorf(codes.Unimplemented, "method SendInput not implemented")
}
func (*UnimplementedExecutorServer) FetchStdout(req *Sn, srv Executor_FetchStdoutServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchStdout not implemented")
}
func (*UnimplementedExecutorServer) FetchStderr(req *Sn, srv Executor_FetchStderrServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchStderr not implemented")
}
func (*UnimplementedExecutorServer) Start(ctx context.Context, req *StartInput) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedExecutorServer) Wait(ctx context.Context, req *Sn) (*WaitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Wait not implemented")
}
func (*UnimplementedExecutorServer) ExecCommand(ctx context.Context, req *Command) (*Sn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecCommand not implemented")
}
func (*UnimplementedExecutorServer) Kill(ctx context.Context, req *Sn) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Kill not implemented")
}

func RegisterExecutorServer(s *grpc.Server, srv ExecutorServer) {
	s.RegisterService(&_Executor_serviceDesc, srv)
}

func _Executor_SendInput_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExecutorServer).SendInput(&executorSendInputServer{stream})
}

type Executor_SendInputServer interface {
	SendAndClose(*Error) error
	Recv() (*Input, error)
	grpc.ServerStream
}

type executorSendInputServer struct {
	grpc.ServerStream
}

func (x *executorSendInputServer) SendAndClose(m *Error) error {
	return x.ServerStream.SendMsg(m)
}

func (x *executorSendInputServer) Recv() (*Input, error) {
	m := new(Input)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Executor_FetchStdout_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Sn)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExecutorServer).FetchStdout(m, &executorFetchStdoutServer{stream})
}

type Executor_FetchStdoutServer interface {
	Send(*Stdout) error
	grpc.ServerStream
}

type executorFetchStdoutServer struct {
	grpc.ServerStream
}

func (x *executorFetchStdoutServer) Send(m *Stdout) error {
	return x.ServerStream.SendMsg(m)
}

func _Executor_FetchStderr_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Sn)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExecutorServer).FetchStderr(m, &executorFetchStderrServer{stream})
}

type Executor_FetchStderrServer interface {
	Send(*Stderr) error
	grpc.ServerStream
}

type executorFetchStderrServer struct {
	grpc.ServerStream
}

func (x *executorFetchStderrServer) Send(m *Stderr) error {
	return x.ServerStream.SendMsg(m)
}

func _Executor_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.Executor/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).Start(ctx, req.(*StartInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_Wait_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).Wait(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.Executor/Wait",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).Wait(ctx, req.(*Sn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_ExecCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).ExecCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.Executor/ExecCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).ExecCommand(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_Kill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).Kill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.Executor/Kill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).Kill(ctx, req.(*Sn))
	}
	return interceptor(ctx, in, info, handler)
}

var _Executor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apis.Executor",
	HandlerType: (*ExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Executor_Start_Handler,
		},
		{
			MethodName: "Wait",
			Handler:    _Executor_Wait_Handler,
		},
		{
			MethodName: "ExecCommand",
			Handler:    _Executor_ExecCommand_Handler,
		},
		{
			MethodName: "Kill",
			Handler:    _Executor_Kill_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendInput",
			Handler:       _Executor_SendInput_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FetchStdout",
			Handler:       _Executor_FetchStdout_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FetchStderr",
			Handler:       _Executor_FetchStderr_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "executor.proto",
}
