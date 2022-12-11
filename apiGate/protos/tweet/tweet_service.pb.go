// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: tweet_service.proto

package tweet

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

type Tweet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Text     string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Tweet) Reset() {
	*x = Tweet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tweet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tweet) ProtoMessage() {}

func (x *Tweet) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tweet.ProtoReflect.Descriptor instead.
func (*Tweet) Descriptor() ([]byte, []int) {
	return file_tweet_service_proto_rawDescGZIP(), []int{0}
}

func (x *Tweet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tweet) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Tweet) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type GetTweetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetTweetRequest) Reset() {
	*x = GetTweetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTweetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTweetRequest) ProtoMessage() {}

func (x *GetTweetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTweetRequest.ProtoReflect.Descriptor instead.
func (*GetTweetRequest) Descriptor() ([]byte, []int) {
	return file_tweet_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetTweetRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetTweetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TweetList []*Tweet `protobuf:"bytes,1,rep,name=tweet_list,json=tweetList,proto3" json:"tweet_list,omitempty"`
}

func (x *GetTweetResponse) Reset() {
	*x = GetTweetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTweetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTweetResponse) ProtoMessage() {}

func (x *GetTweetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTweetResponse.ProtoReflect.Descriptor instead.
func (*GetTweetResponse) Descriptor() ([]byte, []int) {
	return file_tweet_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetTweetResponse) GetTweetList() []*Tweet {
	if x != nil {
		return x.TweetList
	}
	return nil
}

type PostTweetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *PostTweetRequest) Reset() {
	*x = PostTweetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTweetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTweetRequest) ProtoMessage() {}

func (x *PostTweetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTweetRequest.ProtoReflect.Descriptor instead.
func (*PostTweetRequest) Descriptor() ([]byte, []int) {
	return file_tweet_service_proto_rawDescGZIP(), []int{3}
}

func (x *PostTweetRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *PostTweetRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type PostTweetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PostTweetResponse) Reset() {
	*x = PostTweetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTweetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTweetResponse) ProtoMessage() {}

func (x *PostTweetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTweetResponse.ProtoReflect.Descriptor instead.
func (*PostTweetResponse) Descriptor() ([]byte, []int) {
	return file_tweet_service_proto_rawDescGZIP(), []int{4}
}

var File_tweet_service_proto protoreflect.FileDescriptor

var file_tweet_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x77, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x05, 0x54, 0x77, 0x65, 0x65, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x2d,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x0a, 0x74, 0x77, 0x65, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x09, 0x74,
	0x77, 0x65, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74,
	0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x13, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x77,
	0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x78, 0x0a, 0x0c, 0x54,
	0x77, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x73, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x77,
	0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x34, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x77, 0x65, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tweet_service_proto_rawDescOnce sync.Once
	file_tweet_service_proto_rawDescData = file_tweet_service_proto_rawDesc
)

func file_tweet_service_proto_rawDescGZIP() []byte {
	file_tweet_service_proto_rawDescOnce.Do(func() {
		file_tweet_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_tweet_service_proto_rawDescData)
	})
	return file_tweet_service_proto_rawDescData
}

var file_tweet_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tweet_service_proto_goTypes = []interface{}{
	(*Tweet)(nil),             // 0: Tweet
	(*GetTweetRequest)(nil),   // 1: GetTweetRequest
	(*GetTweetResponse)(nil),  // 2: GetTweetResponse
	(*PostTweetRequest)(nil),  // 3: PostTweetRequest
	(*PostTweetResponse)(nil), // 4: PostTweetResponse
}
var file_tweet_service_proto_depIdxs = []int32{
	0, // 0: GetTweetResponse.tweet_list:type_name -> Tweet
	1, // 1: TweetService.GetTweets:input_type -> GetTweetRequest
	3, // 2: TweetService.PostTweet:input_type -> PostTweetRequest
	2, // 3: TweetService.GetTweets:output_type -> GetTweetResponse
	4, // 4: TweetService.PostTweet:output_type -> PostTweetResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tweet_service_proto_init() }
func file_tweet_service_proto_init() {
	if File_tweet_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tweet_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tweet); i {
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
		file_tweet_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTweetRequest); i {
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
		file_tweet_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTweetResponse); i {
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
		file_tweet_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTweetRequest); i {
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
		file_tweet_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTweetResponse); i {
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
			RawDescriptor: file_tweet_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tweet_service_proto_goTypes,
		DependencyIndexes: file_tweet_service_proto_depIdxs,
		MessageInfos:      file_tweet_service_proto_msgTypes,
	}.Build()
	File_tweet_service_proto = out.File
	file_tweet_service_proto_rawDesc = nil
	file_tweet_service_proto_goTypes = nil
	file_tweet_service_proto_depIdxs = nil
}
