//
//
//
//
//

package geo

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_modules_geo_GeoStateDefinitions_proto protoreflect.FileDescriptor

var file_modules_geo_GeoStateDefinitions_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6f, 0x2f, 0x47, 0x65,
	0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x24, 0x5a, 0x22, 0x70, 0x69, 0x78, 0x65, 0x6c,
	0x70, 0x6c, 0x75, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x63,
	0x6b, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_modules_geo_GeoStateDefinitions_proto_goTypes = []interface{}{}
var file_modules_geo_GeoStateDefinitions_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_modules_geo_GeoStateDefinitions_proto_init() }
func file_modules_geo_GeoStateDefinitions_proto_init() {
	if File_modules_geo_GeoStateDefinitions_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_modules_geo_GeoStateDefinitions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_modules_geo_GeoStateDefinitions_proto_goTypes,
		DependencyIndexes: file_modules_geo_GeoStateDefinitions_proto_depIdxs,
	}.Build()
	File_modules_geo_GeoStateDefinitions_proto = out.File
	file_modules_geo_GeoStateDefinitions_proto_rawDesc = nil
	file_modules_geo_GeoStateDefinitions_proto_goTypes = nil
	file_modules_geo_GeoStateDefinitions_proto_depIdxs = nil
}
