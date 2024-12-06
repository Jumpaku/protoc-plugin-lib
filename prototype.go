package protocplugin

import (
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type File struct {
	FullName protoreflect.FullName
	Desc     protoreflect.FileDescriptor
	Options  *FileOptions
	Enums    []*Enum    // top-level enum declarations
	Messages []*Message // top-level message declarations
	Services []*Service // top-level service declarations
}
type FileOptions struct {
	*descriptorpb.FileOptions
}
type Enum struct {
	FullName protoreflect.FullName
	Desc     protoreflect.EnumDescriptor
	Options  *EnumOptions
	Values   []*EnumValue        // enum value declarations
	Comments protogen.CommentSet // comments associated with this enum
}
type EnumOptions struct {
	*descriptorpb.EnumOptions
}
type EnumValue struct {
	FullName protoreflect.FullName
	Desc     protoreflect.EnumValueDescriptor
	Options  *EnumValueOptions
	Parent   *Enum               // enum in which this value is declared
	Comments protogen.CommentSet // comments associated with this enum value
}
type EnumValueOptions struct {
	*descriptorpb.EnumValueOptions
}
type Service struct {
	FullName protoreflect.FullName
	Desc     protoreflect.ServiceDescriptor
	Options  *ServiceOptions
	Parent   *File               // file in which this service is declared
	Methods  []*Method           // service method declarations
	Comments protogen.CommentSet // comments associated with this service
}
type ServiceOptions struct {
	*descriptorpb.ServiceOptions
}
type Method struct {
	FullName protoreflect.FullName
	Desc     protoreflect.MethodDescriptor
	Options  *MethodOptions
	Parent   *Service // service in which this method is declared
	Input    *Message
	Output   *Message
	Comments protogen.CommentSet // comments associated with this method
}
type MethodOptions struct {
	*descriptorpb.MethodOptions
	Http *annotations.HttpRule
}
type Message struct {
	FullName protoreflect.FullName
	Desc     protoreflect.MessageDescriptor
	Options  *MessageOptions
	Fields   []*Field            // message field declarations
	Oneofs   []*Oneof            // message oneof declarations
	Enums    []*Enum             // nested enum declarations
	Messages []*Message          // nested message declarations
	Comments protogen.CommentSet // comments associated with this message
}
type MessageOptions struct {
	*descriptorpb.MessageOptions
}
type Oneof struct {
	FullName protoreflect.FullName
	Desc     protoreflect.OneofDescriptor
	Options  *OneofOptions
	Parent   *Message            // message in which this oneof is declared
	Fields   []*Field            // fields that are part of this oneof
	Comments protogen.CommentSet // comments associated with this oneof
}
type OneofOptions struct {
	descriptorpb.OneofOptions
}
type Field struct {
	FullName protoreflect.FullName
	Desc     protoreflect.FieldDescriptor
	Options  *FieldOptions
	Parent   *Message            // message in which this field is declared; nil if top-level extension
	Oneof    *Oneof              // containing oneof; nil if not part of a oneof
	Extendee *Message            // extended message for extension fields; nil otherwise
	Enum     *Enum               // type for enum fields; nil otherwise
	Message  *Message            // type for message or group fields; nil otherwise
	Comments protogen.CommentSet // comments associated with this field
}
type FieldOptions struct {
	*descriptorpb.FieldOptions
}
