package generator

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/vanity"
	cuzproto "github.com/yrzs/protoc-gen-enum/protobuf/utils/gogo"
)

func IsEnumValueCN(field *descriptor.EnumValueDescriptorProto) bool {
	name := GetEnumValueCN(field)
	if name != "" {
		return true
	}
	return false
}

func GetEnumValueCN(field *descriptor.EnumValueDescriptorProto) string {
	if field == nil {
		return ""
	}
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, cuzproto.E_EnumvalueCn)
		if err == nil && v.(*string) != nil {
			return *(v.(*string))
		}
	}
	return ""
}

func IsEnumType(field *descriptor.EnumDescriptorProto) bool {
	name := GetEnumType(field)
	if len(name) > 0 {
		return true
	}
	return false
}

func GetEnumType(field *descriptor.EnumDescriptorProto) string {
	if field == nil {
		return ""
	}
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, cuzproto.E_EnumCustomtype)
		if err == nil && v.(*string) != nil {
			return *(v.(*string))
		}
	}
	return ""
}

func EnabledGoEnumValueMap(file *descriptor.FileDescriptorProto, enum *descriptor.EnumDescriptorProto) bool {
	return proto.GetBoolExtension(enum.Options, cuzproto.E_EnumGenvaluemap, false)
}

func TurnOffGoEnumValueMap(enum *descriptor.EnumDescriptorProto) {
	vanity.SetBoolEnumOption(cuzproto.E_EnumGenvaluemap, false)(enum)
}

func EnabledEnumNumOrder(file *descriptor.FileDescriptorProto, enum *descriptor.EnumDescriptorProto) bool {
	return proto.GetBoolExtension(enum.Options, cuzproto.E_EnumNumorder, false)
}

func TurnOffEnumNumOrder(enum *descriptor.EnumDescriptorProto) {
	vanity.SetBoolEnumOption(cuzproto.E_EnumNumorder, false)(enum)
}

func EnabledEnumJsonMarshal(file *descriptor.FileDescriptorProto, enum *descriptor.EnumDescriptorProto) bool {
	return proto.GetBoolExtension(enum.Options, cuzproto.E_EnumJsonmarshal, false)
}

func TurnOffEnumJsonMarshal(enum *descriptor.EnumDescriptorProto) {
	vanity.SetBoolEnumOption(cuzproto.E_EnumJsonmarshal, false)(enum)
}

func EnabledEnumErrorCode(enum *descriptor.EnumDescriptorProto) bool {
	return proto.GetBoolExtension(enum.Options, cuzproto.E_EnumErrorcode, false)
}

func TurnOffEnumErrorCode(enum *descriptor.EnumDescriptorProto) {
	vanity.SetBoolEnumOption(cuzproto.E_EnumErrorcode, false)(enum)
}

func EnabledEnumGqlGen(enum *descriptor.EnumDescriptorProto) bool {
	return proto.GetBoolExtension(enum.Options, cuzproto.E_EnumGqlgen, false)
}

func TurnOffEnumGqlGen(enum *descriptor.EnumDescriptorProto) {
	vanity.SetBoolEnumOption(cuzproto.E_EnumGqlgen, false)(enum)
}

func EnabledFileEnumGqlGen(file *descriptor.FileDescriptorProto) bool {
	return proto.GetBoolExtension(file.Options, cuzproto.E_EnumGqlgenAll, false)
}

func TurnOffFileEnumGqlGen(file *descriptor.FileDescriptorProto) {
	vanity.SetBoolFileOption(cuzproto.E_EnumGqlgenAll, false)(file)
}

func EnabledGoEnumPrefix(file *descriptor.FileDescriptorProto, enum *descriptor.EnumDescriptorProto) bool {
	return proto.GetBoolExtension(enum.Options, cuzproto.E_EnumPrefix, proto.GetBoolExtension(file.Options, cuzproto.E_EnumPrefixAll, true))
}
