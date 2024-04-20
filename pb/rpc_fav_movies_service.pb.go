// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: rpc_fav_movies_service.proto

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

// Create Fav Movies
type CreateFavMoviesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MovieIds []string `protobuf:"bytes,2,rep,name=movie_ids,json=movieIds,proto3" json:"movie_ids,omitempty"`
}

func (x *CreateFavMoviesRequest) Reset() {
	*x = CreateFavMoviesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fav_movies_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFavMoviesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFavMoviesRequest) ProtoMessage() {}

func (x *CreateFavMoviesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fav_movies_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFavMoviesRequest.ProtoReflect.Descriptor instead.
func (*CreateFavMoviesRequest) Descriptor() ([]byte, []int) {
	return file_rpc_fav_movies_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFavMoviesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateFavMoviesRequest) GetMovieIds() []string {
	if x != nil {
		return x.MovieIds
	}
	return nil
}

type CreateFavMoviesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FavMovies *FavMovies `protobuf:"bytes,1,opt,name=fav_movies,json=favMovies,proto3" json:"fav_movies,omitempty"`
}

func (x *CreateFavMoviesResponse) Reset() {
	*x = CreateFavMoviesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fav_movies_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFavMoviesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFavMoviesResponse) ProtoMessage() {}

func (x *CreateFavMoviesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fav_movies_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFavMoviesResponse.ProtoReflect.Descriptor instead.
func (*CreateFavMoviesResponse) Descriptor() ([]byte, []int) {
	return file_rpc_fav_movies_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFavMoviesResponse) GetFavMovies() *FavMovies {
	if x != nil {
		return x.FavMovies
	}
	return nil
}

// Get Fav Movies
type GetFavMoviesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetFavMoviesRequest) Reset() {
	*x = GetFavMoviesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fav_movies_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFavMoviesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFavMoviesRequest) ProtoMessage() {}

func (x *GetFavMoviesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fav_movies_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFavMoviesRequest.ProtoReflect.Descriptor instead.
func (*GetFavMoviesRequest) Descriptor() ([]byte, []int) {
	return file_rpc_fav_movies_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetFavMoviesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetFavMoviesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movies []*Movie `protobuf:"bytes,1,rep,name=movies,proto3" json:"movies,omitempty"`
}

func (x *GetFavMoviesResponse) Reset() {
	*x = GetFavMoviesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fav_movies_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFavMoviesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFavMoviesResponse) ProtoMessage() {}

func (x *GetFavMoviesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fav_movies_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFavMoviesResponse.ProtoReflect.Descriptor instead.
func (*GetFavMoviesResponse) Descriptor() ([]byte, []int) {
	return file_rpc_fav_movies_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetFavMoviesResponse) GetMovies() []*Movie {
	if x != nil {
		return x.Movies
	}
	return nil
}

// Update Fav Movies
type UpdateFavMoviesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MovieIds []string `protobuf:"bytes,2,rep,name=movie_ids,json=movieIds,proto3" json:"movie_ids,omitempty"`
}

func (x *UpdateFavMoviesRequest) Reset() {
	*x = UpdateFavMoviesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fav_movies_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFavMoviesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFavMoviesRequest) ProtoMessage() {}

func (x *UpdateFavMoviesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fav_movies_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFavMoviesRequest.ProtoReflect.Descriptor instead.
func (*UpdateFavMoviesRequest) Descriptor() ([]byte, []int) {
	return file_rpc_fav_movies_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateFavMoviesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateFavMoviesRequest) GetMovieIds() []string {
	if x != nil {
		return x.MovieIds
	}
	return nil
}

type UpdateFavMoviesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FavMovies *FavMovies `protobuf:"bytes,1,opt,name=fav_movies,json=favMovies,proto3" json:"fav_movies,omitempty"`
}

func (x *UpdateFavMoviesResponse) Reset() {
	*x = UpdateFavMoviesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fav_movies_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFavMoviesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFavMoviesResponse) ProtoMessage() {}

func (x *UpdateFavMoviesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fav_movies_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFavMoviesResponse.ProtoReflect.Descriptor instead.
func (*UpdateFavMoviesResponse) Descriptor() ([]byte, []int) {
	return file_rpc_fav_movies_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateFavMoviesResponse) GetFavMovies() *FavMovies {
	if x != nil {
		return x.FavMovies
	}
	return nil
}

var File_rpc_fav_movies_service_proto protoreflect.FileDescriptor

var file_rpc_fav_movies_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x70, 0x63, 0x5f, 0x66, 0x61, 0x76, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x10, 0x66, 0x61, 0x76, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4e, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x76, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64,
	0x73, 0x22, 0x47, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x76, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0a,
	0x66, 0x61, 0x76, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x61, 0x76, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52,
	0x09, 0x66, 0x61, 0x76, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x46, 0x61, 0x76, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x46, 0x61, 0x76, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x06, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x73, 0x22, 0x4e, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x61, 0x76, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x49, 0x64, 0x73, 0x22, 0x47, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x61, 0x76, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x0a, 0x66, 0x61, 0x76, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x61, 0x76, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x73, 0x52, 0x09, 0x66, 0x61, 0x76, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x42, 0x2c,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x69, 0x6e,
	0x65, 0x44, 0x65, 0x65, 0x70, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x42, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_fav_movies_service_proto_rawDescOnce sync.Once
	file_rpc_fav_movies_service_proto_rawDescData = file_rpc_fav_movies_service_proto_rawDesc
)

func file_rpc_fav_movies_service_proto_rawDescGZIP() []byte {
	file_rpc_fav_movies_service_proto_rawDescOnce.Do(func() {
		file_rpc_fav_movies_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_fav_movies_service_proto_rawDescData)
	})
	return file_rpc_fav_movies_service_proto_rawDescData
}

var file_rpc_fav_movies_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rpc_fav_movies_service_proto_goTypes = []interface{}{
	(*CreateFavMoviesRequest)(nil),  // 0: pb.CreateFavMoviesRequest
	(*CreateFavMoviesResponse)(nil), // 1: pb.CreateFavMoviesResponse
	(*GetFavMoviesRequest)(nil),     // 2: pb.GetFavMoviesRequest
	(*GetFavMoviesResponse)(nil),    // 3: pb.GetFavMoviesResponse
	(*UpdateFavMoviesRequest)(nil),  // 4: pb.UpdateFavMoviesRequest
	(*UpdateFavMoviesResponse)(nil), // 5: pb.UpdateFavMoviesResponse
	(*FavMovies)(nil),               // 6: pb.FavMovies
	(*Movie)(nil),                   // 7: pb.Movie
}
var file_rpc_fav_movies_service_proto_depIdxs = []int32{
	6, // 0: pb.CreateFavMoviesResponse.fav_movies:type_name -> pb.FavMovies
	7, // 1: pb.GetFavMoviesResponse.movies:type_name -> pb.Movie
	6, // 2: pb.UpdateFavMoviesResponse.fav_movies:type_name -> pb.FavMovies
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rpc_fav_movies_service_proto_init() }
func file_rpc_fav_movies_service_proto_init() {
	if File_rpc_fav_movies_service_proto != nil {
		return
	}
	file_fav_movies_proto_init()
	file_movie_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_fav_movies_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFavMoviesRequest); i {
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
		file_rpc_fav_movies_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFavMoviesResponse); i {
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
		file_rpc_fav_movies_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFavMoviesRequest); i {
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
		file_rpc_fav_movies_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFavMoviesResponse); i {
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
		file_rpc_fav_movies_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFavMoviesRequest); i {
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
		file_rpc_fav_movies_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFavMoviesResponse); i {
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
			RawDescriptor: file_rpc_fav_movies_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_fav_movies_service_proto_goTypes,
		DependencyIndexes: file_rpc_fav_movies_service_proto_depIdxs,
		MessageInfos:      file_rpc_fav_movies_service_proto_msgTypes,
	}.Build()
	File_rpc_fav_movies_service_proto = out.File
	file_rpc_fav_movies_service_proto_rawDesc = nil
	file_rpc_fav_movies_service_proto_goTypes = nil
	file_rpc_fav_movies_service_proto_depIdxs = nil
}