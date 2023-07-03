// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: proto/course.proto

package course

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CourseServiceClient is the client API for CourseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourseServiceClient interface {
	CreateCourse(ctx context.Context, in *NewCourse, opts ...grpc.CallOption) (*Course, error)
	GetCourses(ctx context.Context, in *BlankCourse, opts ...grpc.CallOption) (*CoursesResponse, error)
	GetCourseByID(ctx context.Context, in *CourseID, opts ...grpc.CallOption) (*Course, error)
	DeleteCourse(ctx context.Context, in *CourseID, opts ...grpc.CallOption) (*BlankCourse, error)
	UpdateCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Course, error)
}

type courseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseServiceClient(cc grpc.ClientConnInterface) CourseServiceClient {
	return &courseServiceClient{cc}
}

func (c *courseServiceClient) CreateCourse(ctx context.Context, in *NewCourse, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/pb.CourseService/CreateCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) GetCourses(ctx context.Context, in *BlankCourse, opts ...grpc.CallOption) (*CoursesResponse, error) {
	out := new(CoursesResponse)
	err := c.cc.Invoke(ctx, "/pb.CourseService/GetCourses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) GetCourseByID(ctx context.Context, in *CourseID, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/pb.CourseService/GetCourseByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) DeleteCourse(ctx context.Context, in *CourseID, opts ...grpc.CallOption) (*BlankCourse, error) {
	out := new(BlankCourse)
	err := c.cc.Invoke(ctx, "/pb.CourseService/DeleteCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) UpdateCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/pb.CourseService/UpdateCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseServiceServer is the server API for CourseService service.
// All implementations must embed UnimplementedCourseServiceServer
// for forward compatibility
type CourseServiceServer interface {
	CreateCourse(context.Context, *NewCourse) (*Course, error)
	GetCourses(context.Context, *BlankCourse) (*CoursesResponse, error)
	GetCourseByID(context.Context, *CourseID) (*Course, error)
	DeleteCourse(context.Context, *CourseID) (*BlankCourse, error)
	UpdateCourse(context.Context, *Course) (*Course, error)
	mustEmbedUnimplementedCourseServiceServer()
}

// UnimplementedCourseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCourseServiceServer struct {
}

func (UnimplementedCourseServiceServer) CreateCourse(context.Context, *NewCourse) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourse not implemented")
}
func (UnimplementedCourseServiceServer) GetCourses(context.Context, *BlankCourse) (*CoursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourses not implemented")
}
func (UnimplementedCourseServiceServer) GetCourseByID(context.Context, *CourseID) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourseByID not implemented")
}
func (UnimplementedCourseServiceServer) DeleteCourse(context.Context, *CourseID) (*BlankCourse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourse not implemented")
}
func (UnimplementedCourseServiceServer) UpdateCourse(context.Context, *Course) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCourse not implemented")
}
func (UnimplementedCourseServiceServer) mustEmbedUnimplementedCourseServiceServer() {}

// UnsafeCourseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourseServiceServer will
// result in compilation errors.
type UnsafeCourseServiceServer interface {
	mustEmbedUnimplementedCourseServiceServer()
}

func RegisterCourseServiceServer(s grpc.ServiceRegistrar, srv CourseServiceServer) {
	s.RegisterService(&CourseService_ServiceDesc, srv)
}

func _CourseService_CreateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCourse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).CreateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CourseService/CreateCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).CreateCourse(ctx, req.(*NewCourse))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_GetCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlankCourse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).GetCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CourseService/GetCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).GetCourses(ctx, req.(*BlankCourse))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_GetCourseByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).GetCourseByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CourseService/GetCourseByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).GetCourseByID(ctx, req.(*CourseID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_DeleteCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).DeleteCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CourseService/DeleteCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).DeleteCourse(ctx, req.(*CourseID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_UpdateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Course)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).UpdateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CourseService/UpdateCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).UpdateCourse(ctx, req.(*Course))
	}
	return interceptor(ctx, in, info, handler)
}

// CourseService_ServiceDesc is the grpc.ServiceDesc for CourseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CourseService",
	HandlerType: (*CourseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCourse",
			Handler:    _CourseService_CreateCourse_Handler,
		},
		{
			MethodName: "GetCourses",
			Handler:    _CourseService_GetCourses_Handler,
		},
		{
			MethodName: "GetCourseByID",
			Handler:    _CourseService_GetCourseByID_Handler,
		},
		{
			MethodName: "DeleteCourse",
			Handler:    _CourseService_DeleteCourse_Handler,
		},
		{
			MethodName: "UpdateCourse",
			Handler:    _CourseService_UpdateCourse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/course.proto",
}