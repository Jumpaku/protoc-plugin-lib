package protocplugin

import (
	"fmt"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
	"io"
)

type GeneratedFile struct {
	Name    string
	Content string
}

type PluginHandler func(req *pluginpb.CodeGeneratorRequest, files map[string]*File) ([]*GeneratedFile, error)

func Run(in io.Reader, out io.Writer, handle PluginHandler) error {
	opts := protogen.Options{}
	inBuf, err := io.ReadAll(in)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	req := &pluginpb.CodeGeneratorRequest{}
	err = proto.Unmarshal(inBuf, req)
	if err != nil {
		return fmt.Errorf("failed to unmarshal request: %w", err)
	}

	p, err := opts.New(req)
	if err != nil {
		return fmt.Errorf("failed to create plugin instance: %w", err)
	}

	inFiles := map[string]*File{}
	for _, f := range p.Files {
		inFiles[f.Desc.Path()] = constructFile(f)
	}
	outFiles, err := handle(req, inFiles)

	resp := &pluginpb.CodeGeneratorResponse{}
	if err != nil {
		resp.Error = proto.String(fmt.Sprintf("%+v", err))
	} else {
		for _, f := range outFiles {
			resp.File = append(resp.File, &pluginpb.CodeGeneratorResponse_File{
				Name:    proto.String(f.Name),
				Content: proto.String(f.Content),
			})
		}
	}

	outBuf, err := proto.Marshal(resp)
	if err != nil {
		return fmt.Errorf("failed to marshal response: %w", err)
	}

	_, err = out.Write(outBuf)
	if err != nil {
		return fmt.Errorf("failed to rwite output: %w", err)
	}

	return nil
}

func constructFile(f *protogen.File) *File {
	file := &File{
		FullName: f.Desc.FullName(),
		Desc:     f.Desc,
	}
	if o := f.Desc.Options().(*descriptorpb.FileOptions); o != nil {
		file.Options = &FileOptions{FileOptions: o}
	}
	for _, e := range f.Enums {
		file.Enums = append(file.Enums, constructEnum(e))
	}
	for _, m := range f.Messages {
		file.Messages = append(file.Messages, constructMessage(m))
	}
	for _, s := range f.Services {
		file.Services = append(file.Services, constructService(file, s))
	}
	return file
}

func constructService(parent *File, s *protogen.Service) *Service {
	service := &Service{
		FullName: s.Desc.FullName(),
		Desc:     s.Desc,
		Parent:   parent,
		Comments: s.Comments,
	}
	if o := s.Desc.Options().(*descriptorpb.ServiceOptions); o != nil {
		service.Options = &ServiceOptions{ServiceOptions: o}
	}
	for _, m := range s.Methods {
		service.Methods = append(service.Methods, constructMethod(service, m))
	}
	return service
}

func constructMethod(parent *Service, m *protogen.Method) *Method {
	method := &Method{
		FullName: m.Desc.FullName(),
		Desc:     m.Desc,
		Parent:   parent,
		Input:    constructMessage(m.Input),
		Output:   constructMessage(m.Output),
		Comments: m.Comments,
	}
	if o := m.Desc.Options().(*descriptorpb.MethodOptions); o != nil {
		method.Options = &MethodOptions{MethodOptions: o}
		if httpRule, ok := proto.GetExtension(o, annotations.E_Http).(*annotations.HttpRule); ok {
			method.Options.Http = httpRule
		}
	}
	return method
}

func constructMessage(m *protogen.Message) *Message {
	message := &Message{
		FullName: m.Desc.FullName(),
		Desc:     m.Desc,
		Comments: m.Comments,
	}
	if o := m.Desc.Options().(*descriptorpb.MessageOptions); o != nil {
		message.Options = &MessageOptions{MessageOptions: o}
	}
	for _, f := range m.Fields {
		message.Fields = append(message.Fields, constructField(message, f))
	}
	return message
}

func constructField(parent *Message, f *protogen.Field) *Field {
	field := &Field{
		FullName: f.Desc.FullName(),
		Desc:     f.Desc,
		Parent:   parent,
		Comments: f.Comments,
	}
	if o := f.Desc.Options().(*descriptorpb.FieldOptions); o != nil {
		field.Options = &FieldOptions{FieldOptions: o}
	}
	return field
}

func constructEnum(e *protogen.Enum) *Enum {
	enum := &Enum{
		FullName: e.Desc.FullName(),
		Desc:     e.Desc,
		Comments: e.Comments,
	}
	if o := e.Desc.Options().(*descriptorpb.EnumOptions); o != nil {
		enum.Options = &EnumOptions{EnumOptions: o}
	}
	for _, v := range e.Values {
		enum.Values = append(enum.Values, constructEnumValue(enum, v))
	}
	return enum
}

func constructEnumValue(parent *Enum, v *protogen.EnumValue) *EnumValue {
	enumValue := &EnumValue{
		FullName: v.Desc.FullName(),
		Desc:     v.Desc,
		Parent:   parent,
		Comments: v.Comments,
	}
	if o := v.Desc.Options().(*descriptorpb.EnumValueOptions); o != nil {
		enumValue.Options = &EnumValueOptions{EnumValueOptions: o}
	}
	return enumValue
}

func panicIfError(err error) {
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
}
