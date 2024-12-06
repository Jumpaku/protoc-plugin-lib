package protocplugin

import (
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// File represents a protobuf file descriptor.
type File struct {
	FullName protoreflect.FullName       // FullName is the full name of the file.
	Desc     protoreflect.FileDescriptor // Desc is the file descriptor.
	Options  *FileOptions                // Options are the file options.
	Enums    []*Enum                     // Enums are the enums defined in the file.
	Messages []*Message                  // Messages are the messages defined in the file.
	Services []*Service                  // Services are the services defined in the file.
}

// FileOptions represents the options for a protobuf file.
type FileOptions struct {
	*descriptorpb.FileOptions // FileOptions is the embedded protobuf file options.
}

// Enum represents a protobuf enum descriptor.
type Enum struct {
	FullName protoreflect.FullName       // FullName is the full name of the enum.
	Desc     protoreflect.EnumDescriptor // Desc is the enum descriptor.
	Options  *EnumOptions                // Options are the enum options.
	Values   []*EnumValue                // Values are the values defined in the enum.
	Comments protogen.CommentSet         // Comments are the comments associated with the enum.
}

// EnumOptions represents the options for a protobuf enum.
type EnumOptions struct {
	*descriptorpb.EnumOptions // EnumOptions is the embedded protobuf enum options.
}

// EnumValue represents a value in a protobuf enum.
type EnumValue struct {
	FullName protoreflect.FullName            // FullName is the full name of the enum value.
	Desc     protoreflect.EnumValueDescriptor // Desc is the enum value descriptor.
	Options  *EnumValueOptions                // Options are the enum value options.
	Parent   *Enum                            // Parent is the parent enum.
	Comments protogen.CommentSet              // Comments are the comments associated with the enum value.
}

// EnumValueOptions represents the options for a protobuf enum value.
type EnumValueOptions struct {
	*descriptorpb.EnumValueOptions // EnumValueOptions is the embedded protobuf enum value options.
}

// Service represents a protobuf service descriptor.
type Service struct {
	FullName protoreflect.FullName          // FullName is the full name of the service.
	Desc     protoreflect.ServiceDescriptor // Desc is the service descriptor.
	Options  *ServiceOptions                // Options are the service options.
	Parent   *File                          // Parent is the parent file.
	Methods  []*Method                      // Methods are the methods defined in the service.
	Comments protogen.CommentSet            // Comments are the comments associated with the service.
}

// ServiceOptions represents the options for a protobuf service.
type ServiceOptions struct {
	*descriptorpb.ServiceOptions // ServiceOptions is the embedded protobuf service options.
}

// Method represents a method in a protobuf service.
type Method struct {
	FullName protoreflect.FullName         // FullName is the full name of the method.
	Desc     protoreflect.MethodDescriptor // Desc is the method descriptor.
	Options  *MethodOptions                // Options are the method options.
	Parent   *Service                      // Parent is the parent service.
	Input    *Message                      // Input is the input message of the method.
	Output   *Message                      // Output is the output message of the method.
	Comments protogen.CommentSet           // Comments are the comments associated with the method.
}

// MethodOptions represents the options for a protobuf method.
type MethodOptions struct {
	*descriptorpb.MethodOptions                       // MethodOptions is the embedded protobuf method options.
	Http                        *annotations.HttpRule // Http is the HTTP rule associated with the method.
}

// Message represents a protobuf message descriptor.
type Message struct {
	FullName protoreflect.FullName          // FullName is the full name of the message.
	Desc     protoreflect.MessageDescriptor // Desc is the message descriptor.
	Options  *MessageOptions                // Options are the message options.
	Fields   []*Field                       // Fields are the fields defined in the message.
	Oneofs   []*Oneof                       // Oneofs are the oneof fields defined in the message.
	Enums    []*Enum                        // Enums are the enums defined in the message.
	Messages []*Message                     // Messages are the nested messages defined in the message.
	Comments protogen.CommentSet            // Comments are the comments associated with the message.
}

// MessageOptions represents the options for a protobuf message.
type MessageOptions struct {
	*descriptorpb.MessageOptions // MessageOptions is the embedded protobuf message options.
}

// Oneof represents a oneof field in a protobuf message.
type Oneof struct {
	FullName protoreflect.FullName        // FullName is the full name of the oneof field.
	Desc     protoreflect.OneofDescriptor // Desc is the oneof field descriptor.
	Options  *OneofOptions                // Options are the oneof field options.
	Parent   *Message                     // Parent is the parent message.
	Fields   []*Field                     // Fields are the fields defined in the oneof.
	Comments protogen.CommentSet          // Comments are the comments associated with the oneof field.
}

// OneofOptions represents the options for a protobuf oneof field.
type OneofOptions struct {
	descriptorpb.OneofOptions // OneofOptions is the embedded protobuf oneof options.
}

// Field represents a field in a protobuf message.
type Field struct {
	FullName protoreflect.FullName        // FullName is the full name of the field.
	Desc     protoreflect.FieldDescriptor // Desc is the field descriptor.
	Options  *FieldOptions                // Options are the field options.
	Parent   *Message                     // Parent is the parent message.
	Oneof    *Oneof                       // Oneof is the oneof field this field belongs to, if any.
	Extendee *Message                     // Extendee is the extended message, if this is an extension field.
	Enum     *Enum                        // Enum is the enum type of the field, if any.
	Message  *Message                     // Message is the message type of the field, if any.
	Comments protogen.CommentSet          // Comments are the comments associated with the field.
}

// FieldOptions represents the options for a protobuf field.
type FieldOptions struct {
	*descriptorpb.FieldOptions // FieldOptions is the embedded protobuf field options.
}
