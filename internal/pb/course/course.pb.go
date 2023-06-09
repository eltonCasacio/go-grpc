// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: proto/course.proto

package course

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

type BlankCourse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BlankCourse) Reset() {
	*x = BlankCourse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlankCourse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlankCourse) ProtoMessage() {}

func (x *BlankCourse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlankCourse.ProtoReflect.Descriptor instead.
func (*BlankCourse) Descriptor() ([]byte, []int) {
	return file_proto_course_proto_rawDescGZIP(), []int{0}
}

type CourseID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *CourseID) Reset() {
	*x = CourseID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseID) ProtoMessage() {}

func (x *CourseID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseID.ProtoReflect.Descriptor instead.
func (*CourseID) Descriptor() ([]byte, []int) {
	return file_proto_course_proto_rawDescGZIP(), []int{1}
}

func (x *CourseID) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type NewCourse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	CategoryId  string `protobuf:"bytes,3,opt,name=Category_id,json=CategoryId,proto3" json:"Category_id,omitempty"`
}

func (x *NewCourse) Reset() {
	*x = NewCourse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCourse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCourse) ProtoMessage() {}

func (x *NewCourse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCourse.ProtoReflect.Descriptor instead.
func (*NewCourse) Descriptor() ([]byte, []int) {
	return file_proto_course_proto_rawDescGZIP(), []int{2}
}

func (x *NewCourse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewCourse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewCourse) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	CategoryId  string `protobuf:"bytes,4,opt,name=Category_id,json=CategoryId,proto3" json:"Category_id,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_proto_course_proto_rawDescGZIP(), []int{3}
}

func (x *Course) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Course) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Course) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Course) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type CoursesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Courses []*Course `protobuf:"bytes,1,rep,name=Courses,proto3" json:"Courses,omitempty"`
}

func (x *CoursesResponse) Reset() {
	*x = CoursesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoursesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoursesResponse) ProtoMessage() {}

func (x *CoursesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoursesResponse.ProtoReflect.Descriptor instead.
func (*CoursesResponse) Descriptor() ([]byte, []int) {
	return file_proto_course_proto_rawDescGZIP(), []int{4}
}

func (x *CoursesResponse) GetCourses() []*Course {
	if x != nil {
		return x.Courses
	}
	return nil
}

var File_proto_course_proto protoreflect.FileDescriptor

var file_proto_course_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x0d, 0x0a, 0x0b, 0x62, 0x6c, 0x61, 0x6e,
	0x6b, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x22, 0x62, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x0f, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x07, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x73, 0x32, 0xfa, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x34, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x12, 0x0f,
	0x2e, 0x70, 0x62, 0x2e, 0x62, 0x6c, 0x61, 0x6e, 0x6b, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x49, 0x44, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49,
	0x44, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x62, 0x6c, 0x61, 0x6e, 0x6b, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14,
	0x5a, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_course_proto_rawDescOnce sync.Once
	file_proto_course_proto_rawDescData = file_proto_course_proto_rawDesc
)

func file_proto_course_proto_rawDescGZIP() []byte {
	file_proto_course_proto_rawDescOnce.Do(func() {
		file_proto_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_course_proto_rawDescData)
	})
	return file_proto_course_proto_rawDescData
}

var file_proto_course_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_course_proto_goTypes = []interface{}{
	(*BlankCourse)(nil),     // 0: pb.blankCourse
	(*CourseID)(nil),        // 1: pb.CourseID
	(*NewCourse)(nil),       // 2: pb.NewCourse
	(*Course)(nil),          // 3: pb.Course
	(*CoursesResponse)(nil), // 4: pb.CoursesResponse
}
var file_proto_course_proto_depIdxs = []int32{
	3, // 0: pb.CoursesResponse.Courses:type_name -> pb.Course
	2, // 1: pb.CourseService.CreateCourse:input_type -> pb.NewCourse
	0, // 2: pb.CourseService.GetCourses:input_type -> pb.blankCourse
	1, // 3: pb.CourseService.GetCourseByID:input_type -> pb.CourseID
	1, // 4: pb.CourseService.DeleteCourse:input_type -> pb.CourseID
	3, // 5: pb.CourseService.UpdateCourse:input_type -> pb.Course
	3, // 6: pb.CourseService.CreateCourse:output_type -> pb.Course
	4, // 7: pb.CourseService.GetCourses:output_type -> pb.CoursesResponse
	3, // 8: pb.CourseService.GetCourseByID:output_type -> pb.Course
	0, // 9: pb.CourseService.DeleteCourse:output_type -> pb.blankCourse
	3, // 10: pb.CourseService.UpdateCourse:output_type -> pb.Course
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_course_proto_init() }
func file_proto_course_proto_init() {
	if File_proto_course_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_course_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlankCourse); i {
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
		file_proto_course_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseID); i {
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
		file_proto_course_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewCourse); i {
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
		file_proto_course_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
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
		file_proto_course_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoursesResponse); i {
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
			RawDescriptor: file_proto_course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_course_proto_goTypes,
		DependencyIndexes: file_proto_course_proto_depIdxs,
		MessageInfos:      file_proto_course_proto_msgTypes,
	}.Build()
	File_proto_course_proto = out.File
	file_proto_course_proto_rawDesc = nil
	file_proto_course_proto_goTypes = nil
	file_proto_course_proto_depIdxs = nil
}
