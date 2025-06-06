package dokodemo

import (
	net "github.com/v2fly/v2ray-core/v5/common/net"
	_ "github.com/v2fly/v2ray-core/v5/common/protoext"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Config struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Address *net.IPOrDomain        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port    uint32                 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// List of networks that the Dokodemo accepts.
	// Deprecated. Use networks.
	//
	// Deprecated: Marked as deprecated in proxy/dokodemo/config.proto.
	NetworkList *net.NetworkList `protobuf:"bytes,3,opt,name=network_list,json=networkList,proto3" json:"network_list,omitempty"`
	// List of networks that the Dokodemo accepts.
	Networks []net.Network `protobuf:"varint,7,rep,packed,name=networks,proto3,enum=v2ray.core.common.net.Network" json:"networks,omitempty"`
	// Deprecated: Marked as deprecated in proxy/dokodemo/config.proto.
	Timeout        uint32 `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	FollowRedirect bool   `protobuf:"varint,5,opt,name=follow_redirect,json=followRedirect,proto3" json:"follow_redirect,omitempty"`
	UserLevel      uint32 `protobuf:"varint,6,opt,name=user_level,json=userLevel,proto3" json:"user_level,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_proxy_dokodemo_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_dokodemo_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_proxy_dokodemo_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetAddress() *net.IPOrDomain {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Config) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

// Deprecated: Marked as deprecated in proxy/dokodemo/config.proto.
func (x *Config) GetNetworkList() *net.NetworkList {
	if x != nil {
		return x.NetworkList
	}
	return nil
}

func (x *Config) GetNetworks() []net.Network {
	if x != nil {
		return x.Networks
	}
	return nil
}

// Deprecated: Marked as deprecated in proxy/dokodemo/config.proto.
func (x *Config) GetTimeout() uint32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *Config) GetFollowRedirect() bool {
	if x != nil {
		return x.FollowRedirect
	}
	return false
}

func (x *Config) GetUserLevel() uint32 {
	if x != nil {
		return x.UserLevel
	}
	return 0
}

type SimplifiedConfig struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Address        *net.IPOrDomain        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port           uint32                 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Networks       *net.NetworkList       `protobuf:"bytes,3,opt,name=networks,proto3" json:"networks,omitempty"`
	FollowRedirect bool                   `protobuf:"varint,4,opt,name=follow_redirect,json=followRedirect,proto3" json:"follow_redirect,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *SimplifiedConfig) Reset() {
	*x = SimplifiedConfig{}
	mi := &file_proxy_dokodemo_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimplifiedConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimplifiedConfig) ProtoMessage() {}

func (x *SimplifiedConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_dokodemo_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimplifiedConfig.ProtoReflect.Descriptor instead.
func (*SimplifiedConfig) Descriptor() ([]byte, []int) {
	return file_proxy_dokodemo_config_proto_rawDescGZIP(), []int{1}
}

func (x *SimplifiedConfig) GetAddress() *net.IPOrDomain {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *SimplifiedConfig) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *SimplifiedConfig) GetNetworks() *net.NetworkList {
	if x != nil {
		return x.Networks
	}
	return nil
}

func (x *SimplifiedConfig) GetFollowRedirect() bool {
	if x != nil {
		return x.FollowRedirect
	}
	return false
}

var File_proxy_dokodemo_config_proto protoreflect.FileDescriptor

var file_proxy_dokodemo_config_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x64, 0x6f, 0x6b, 0x6f, 0x64, 0x65, 0x6d, 0x6f,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x64, 0x6f, 0x6b, 0x6f, 0x64, 0x65, 0x6d, 0x6f, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x6e, 0x65, 0x74, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6,
	0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6e,
	0x65, 0x74, 0x2e, 0x49, 0x50, 0x4f, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x49, 0x0a, 0x0c, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x12, 0x1c, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x02, 0x18, 0x01, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12,
	0x27, 0x0a, 0x0f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0xea, 0x01, 0x0a, 0x10, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x49, 0x50, 0x4f, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3e, 0x0a,
	0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x27, 0x0a,
	0x0f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x3a, 0x1c, 0x82, 0xb5, 0x18, 0x18, 0x0a, 0x07, 0x69, 0x6e,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x0d, 0x64, 0x6f, 0x6b, 0x6f, 0x64, 0x65, 0x6d, 0x6f, 0x2d,
	0x64, 0x6f, 0x6f, 0x72, 0x42, 0x6c, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x64, 0x6f, 0x6b,
	0x6f, 0x64, 0x65, 0x6d, 0x6f, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2d,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x35, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x64, 0x6f,
	0x6b, 0x6f, 0x64, 0x65, 0x6d, 0x6f, 0xaa, 0x02, 0x19, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43,
	0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x44, 0x6f, 0x6b, 0x6f, 0x64, 0x65,
	0x6d, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proxy_dokodemo_config_proto_rawDescOnce sync.Once
	file_proxy_dokodemo_config_proto_rawDescData []byte
)

func file_proxy_dokodemo_config_proto_rawDescGZIP() []byte {
	file_proxy_dokodemo_config_proto_rawDescOnce.Do(func() {
		file_proxy_dokodemo_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proxy_dokodemo_config_proto_rawDesc), len(file_proxy_dokodemo_config_proto_rawDesc)))
	})
	return file_proxy_dokodemo_config_proto_rawDescData
}

var file_proxy_dokodemo_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proxy_dokodemo_config_proto_goTypes = []any{
	(*Config)(nil),           // 0: v2ray.core.proxy.dokodemo.Config
	(*SimplifiedConfig)(nil), // 1: v2ray.core.proxy.dokodemo.SimplifiedConfig
	(*net.IPOrDomain)(nil),   // 2: v2ray.core.common.net.IPOrDomain
	(*net.NetworkList)(nil),  // 3: v2ray.core.common.net.NetworkList
	(net.Network)(0),         // 4: v2ray.core.common.net.Network
}
var file_proxy_dokodemo_config_proto_depIdxs = []int32{
	2, // 0: v2ray.core.proxy.dokodemo.Config.address:type_name -> v2ray.core.common.net.IPOrDomain
	3, // 1: v2ray.core.proxy.dokodemo.Config.network_list:type_name -> v2ray.core.common.net.NetworkList
	4, // 2: v2ray.core.proxy.dokodemo.Config.networks:type_name -> v2ray.core.common.net.Network
	2, // 3: v2ray.core.proxy.dokodemo.SimplifiedConfig.address:type_name -> v2ray.core.common.net.IPOrDomain
	3, // 4: v2ray.core.proxy.dokodemo.SimplifiedConfig.networks:type_name -> v2ray.core.common.net.NetworkList
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proxy_dokodemo_config_proto_init() }
func file_proxy_dokodemo_config_proto_init() {
	if File_proxy_dokodemo_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proxy_dokodemo_config_proto_rawDesc), len(file_proxy_dokodemo_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proxy_dokodemo_config_proto_goTypes,
		DependencyIndexes: file_proxy_dokodemo_config_proto_depIdxs,
		MessageInfos:      file_proxy_dokodemo_config_proto_msgTypes,
	}.Build()
	File_proxy_dokodemo_config_proto = out.File
	file_proxy_dokodemo_config_proto_goTypes = nil
	file_proxy_dokodemo_config_proto_depIdxs = nil
}
