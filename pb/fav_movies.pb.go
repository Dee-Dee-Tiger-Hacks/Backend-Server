// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: fav_movies.proto

package pb

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

type FavMovies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MovieIds []string `protobuf:"bytes,2,rep,name=movie_ids,json=movieIds,proto3" json:"movie_ids,omitempty"`
}

func (x *FavMovies) Reset() {
	*x = FavMovies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fav_movies_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavMovies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavMovies) ProtoMessage() {}

func (x *FavMovies) ProtoReflect() protoreflect.Message {
	mi := &file_fav_movies_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavMovies.ProtoReflect.Descriptor instead.
func (*FavMovies) Descriptor() ([]byte, []int) {
	return file_fav_movies_proto_rawDescGZIP(), []int{0}
}

func (x *FavMovies) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FavMovies) GetMovieIds() []string {
	if x != nil {
		return x.MovieIds
	}
	return nil
}

var File_fav_movies_proto protoreflect.FileDescriptor

var file_fav_movies_proto_rawDesc = []byte{
	0x0a, 0x10, 0x66, 0x61, 0x76, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x41, 0x0a, 0x09, 0x46, 0x61, 0x76, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x73, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x65, 0x70,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fav_movies_proto_rawDescOnce sync.Once
	file_fav_movies_proto_rawDescData = file_fav_movies_proto_rawDesc
)

func file_fav_movies_proto_rawDescGZIP() []byte {
	file_fav_movies_proto_rawDescOnce.Do(func() {
		file_fav_movies_proto_rawDescData = protoimpl.X.CompressGZIP(file_fav_movies_proto_rawDescData)
	})
	return file_fav_movies_proto_rawDescData
}

var file_fav_movies_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fav_movies_proto_goTypes = []interface{}{
	(*FavMovies)(nil), // 0: pb.FavMovies
}
var file_fav_movies_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fav_movies_proto_init() }
func file_fav_movies_proto_init() {
	if File_fav_movies_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fav_movies_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavMovies); i {
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
			RawDescriptor: file_fav_movies_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fav_movies_proto_goTypes,
		DependencyIndexes: file_fav_movies_proto_depIdxs,
		MessageInfos:      file_fav_movies_proto_msgTypes,
	}.Build()
	File_fav_movies_proto = out.File
	file_fav_movies_proto_rawDesc = nil
	file_fav_movies_proto_goTypes = nil
	file_fav_movies_proto_depIdxs = nil
}