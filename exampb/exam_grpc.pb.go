// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: exampb/exam.proto

package exampb

import (
	context "context"
	studentpb "github.com/ChrisCodeX/gRPC/studentpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ExamServiceClient is the client API for ExamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExamServiceClient interface {
	// Client request
	GetExam(ctx context.Context, in *GetExamRequest, opts ...grpc.CallOption) (*Exam, error)
	// Server response
	SetExam(ctx context.Context, in *Exam, opts ...grpc.CallOption) (*SetExamResponse, error)
	// Stream from client - 1 response from server
	SetQuestions(ctx context.Context, opts ...grpc.CallOption) (ExamService_SetQuestionsClient, error)
	// Stream from client - 1 response from server
	EnrollStudents(ctx context.Context, opts ...grpc.CallOption) (ExamService_EnrollStudentsClient, error)
	// Stream for server - 1 response from client
	GetStudentsPerExam(ctx context.Context, in *GetStudentsPerExamRequest, opts ...grpc.CallOption) (ExamService_GetStudentsPerExamClient, error)
}

type examServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExamServiceClient(cc grpc.ClientConnInterface) ExamServiceClient {
	return &examServiceClient{cc}
}

func (c *examServiceClient) GetExam(ctx context.Context, in *GetExamRequest, opts ...grpc.CallOption) (*Exam, error) {
	out := new(Exam)
	err := c.cc.Invoke(ctx, "/exam.ExamService/GetExam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *examServiceClient) SetExam(ctx context.Context, in *Exam, opts ...grpc.CallOption) (*SetExamResponse, error) {
	out := new(SetExamResponse)
	err := c.cc.Invoke(ctx, "/exam.ExamService/SetExam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *examServiceClient) SetQuestions(ctx context.Context, opts ...grpc.CallOption) (ExamService_SetQuestionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExamService_ServiceDesc.Streams[0], "/exam.ExamService/SetQuestions", opts...)
	if err != nil {
		return nil, err
	}
	x := &examServiceSetQuestionsClient{stream}
	return x, nil
}

type ExamService_SetQuestionsClient interface {
	Send(*Question) error
	CloseAndRecv() (*SetQuestionResponse, error)
	grpc.ClientStream
}

type examServiceSetQuestionsClient struct {
	grpc.ClientStream
}

func (x *examServiceSetQuestionsClient) Send(m *Question) error {
	return x.ClientStream.SendMsg(m)
}

func (x *examServiceSetQuestionsClient) CloseAndRecv() (*SetQuestionResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SetQuestionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *examServiceClient) EnrollStudents(ctx context.Context, opts ...grpc.CallOption) (ExamService_EnrollStudentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExamService_ServiceDesc.Streams[1], "/exam.ExamService/EnrollStudents", opts...)
	if err != nil {
		return nil, err
	}
	x := &examServiceEnrollStudentsClient{stream}
	return x, nil
}

type ExamService_EnrollStudentsClient interface {
	Send(*EnrollmentRequest) error
	CloseAndRecv() (*EnrollmentResponse, error)
	grpc.ClientStream
}

type examServiceEnrollStudentsClient struct {
	grpc.ClientStream
}

func (x *examServiceEnrollStudentsClient) Send(m *EnrollmentRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *examServiceEnrollStudentsClient) CloseAndRecv() (*EnrollmentResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EnrollmentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *examServiceClient) GetStudentsPerExam(ctx context.Context, in *GetStudentsPerExamRequest, opts ...grpc.CallOption) (ExamService_GetStudentsPerExamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExamService_ServiceDesc.Streams[2], "/exam.ExamService/GetStudentsPerExam", opts...)
	if err != nil {
		return nil, err
	}
	x := &examServiceGetStudentsPerExamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExamService_GetStudentsPerExamClient interface {
	Recv() (*studentpb.Student, error)
	grpc.ClientStream
}

type examServiceGetStudentsPerExamClient struct {
	grpc.ClientStream
}

func (x *examServiceGetStudentsPerExamClient) Recv() (*studentpb.Student, error) {
	m := new(studentpb.Student)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExamServiceServer is the server API for ExamService service.
// All implementations must embed UnimplementedExamServiceServer
// for forward compatibility
type ExamServiceServer interface {
	// Client request
	GetExam(context.Context, *GetExamRequest) (*Exam, error)
	// Server response
	SetExam(context.Context, *Exam) (*SetExamResponse, error)
	// Stream from client - 1 response from server
	SetQuestions(ExamService_SetQuestionsServer) error
	// Stream from client - 1 response from server
	EnrollStudents(ExamService_EnrollStudentsServer) error
	// Stream for server - 1 response from client
	GetStudentsPerExam(*GetStudentsPerExamRequest, ExamService_GetStudentsPerExamServer) error
	mustEmbedUnimplementedExamServiceServer()
}

// UnimplementedExamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExamServiceServer struct {
}

func (UnimplementedExamServiceServer) GetExam(context.Context, *GetExamRequest) (*Exam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExam not implemented")
}
func (UnimplementedExamServiceServer) SetExam(context.Context, *Exam) (*SetExamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetExam not implemented")
}
func (UnimplementedExamServiceServer) SetQuestions(ExamService_SetQuestionsServer) error {
	return status.Errorf(codes.Unimplemented, "method SetQuestions not implemented")
}
func (UnimplementedExamServiceServer) EnrollStudents(ExamService_EnrollStudentsServer) error {
	return status.Errorf(codes.Unimplemented, "method EnrollStudents not implemented")
}
func (UnimplementedExamServiceServer) GetStudentsPerExam(*GetStudentsPerExamRequest, ExamService_GetStudentsPerExamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStudentsPerExam not implemented")
}
func (UnimplementedExamServiceServer) mustEmbedUnimplementedExamServiceServer() {}

// UnsafeExamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExamServiceServer will
// result in compilation errors.
type UnsafeExamServiceServer interface {
	mustEmbedUnimplementedExamServiceServer()
}

func RegisterExamServiceServer(s grpc.ServiceRegistrar, srv ExamServiceServer) {
	s.RegisterService(&ExamService_ServiceDesc, srv)
}

func _ExamService_GetExam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamServiceServer).GetExam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exam.ExamService/GetExam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamServiceServer).GetExam(ctx, req.(*GetExamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExamService_SetExam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Exam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamServiceServer).SetExam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exam.ExamService/SetExam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamServiceServer).SetExam(ctx, req.(*Exam))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExamService_SetQuestions_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExamServiceServer).SetQuestions(&examServiceSetQuestionsServer{stream})
}

type ExamService_SetQuestionsServer interface {
	SendAndClose(*SetQuestionResponse) error
	Recv() (*Question, error)
	grpc.ServerStream
}

type examServiceSetQuestionsServer struct {
	grpc.ServerStream
}

func (x *examServiceSetQuestionsServer) SendAndClose(m *SetQuestionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *examServiceSetQuestionsServer) Recv() (*Question, error) {
	m := new(Question)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ExamService_EnrollStudents_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExamServiceServer).EnrollStudents(&examServiceEnrollStudentsServer{stream})
}

type ExamService_EnrollStudentsServer interface {
	SendAndClose(*EnrollmentResponse) error
	Recv() (*EnrollmentRequest, error)
	grpc.ServerStream
}

type examServiceEnrollStudentsServer struct {
	grpc.ServerStream
}

func (x *examServiceEnrollStudentsServer) SendAndClose(m *EnrollmentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *examServiceEnrollStudentsServer) Recv() (*EnrollmentRequest, error) {
	m := new(EnrollmentRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ExamService_GetStudentsPerExam_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetStudentsPerExamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExamServiceServer).GetStudentsPerExam(m, &examServiceGetStudentsPerExamServer{stream})
}

type ExamService_GetStudentsPerExamServer interface {
	Send(*studentpb.Student) error
	grpc.ServerStream
}

type examServiceGetStudentsPerExamServer struct {
	grpc.ServerStream
}

func (x *examServiceGetStudentsPerExamServer) Send(m *studentpb.Student) error {
	return x.ServerStream.SendMsg(m)
}

// ExamService_ServiceDesc is the grpc.ServiceDesc for ExamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "exam.ExamService",
	HandlerType: (*ExamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExam",
			Handler:    _ExamService_GetExam_Handler,
		},
		{
			MethodName: "SetExam",
			Handler:    _ExamService_SetExam_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SetQuestions",
			Handler:       _ExamService_SetQuestions_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "EnrollStudents",
			Handler:       _ExamService_EnrollStudents_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetStudentsPerExam",
			Handler:       _ExamService_GetStudentsPerExam_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "exampb/exam.proto",
}