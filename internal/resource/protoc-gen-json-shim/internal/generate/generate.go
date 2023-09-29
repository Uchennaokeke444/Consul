// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package generate

import (
	"path"
	"strings"
	"sync"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

func Generate(gen *protogen.Plugin) error {
	for _, file := range gen.Files {
		if file.Generate != true {
			continue
		}
		filename := file.GeneratedFilenamePrefix + "_json.gen.go"
		genFile := gen.NewGeneratedFile(filename, file.GoImportPath)

		genFile.P("// Code generated by protoc-json-shim. DO NOT EDIT.")
		genFile.P("package ", file.GoPackageName)
		var process func([]*protogen.Message)

		marshalerName := FileName(file) + "Marshaler"
		unmarshalerName := FileName(file) + "Unmarshaler"

		process = func(messages []*protogen.Message) {
			for _, message := range messages {
				// skip maps in protos.
				if message.Desc.Options().(*descriptorpb.MessageOptions).GetMapEntry() {
					continue
				}
				typeName := message.GoIdent.GoName
				genFile.P(`// MarshalJSON is a custom marshaler for `, typeName)
				genFile.P(`func (this *`, typeName, `) MarshalJSON() ([]byte, error) {`)
				genFile.P(`str, err := `, marshalerName, `.Marshal(this)`)
				genFile.P(`return []byte(str), err`)
				genFile.P(`}`)
				// Generate UnmarshalJSON() method for this type
				genFile.P(`// UnmarshalJSON is a custom unmarshaler for `, typeName)
				genFile.P(`func (this *`, typeName, `) UnmarshalJSON(b []byte) error {`)
				genFile.P(`return `, unmarshalerName, `.Unmarshal(b, this)`)
				genFile.P(`}`)
				process(message.Messages)
			}
		}
		process(file.Messages)

		// write out globals
		genFile.P(`var (`)
		genFile.P(marshalerName, ` = &`, protogen.GoIdent{GoName: "MarshalOptions", GoImportPath: "google.golang.org/protobuf/encoding/protojson"}, `{}`)
		genFile.P(unmarshalerName, ` = &`, protogen.GoIdent{GoName: "UnmarshalOptions", GoImportPath: "google.golang.org/protobuf/encoding/protojson"}, `{DiscardUnknown: false}`)
		genFile.P(`)`)
	}
	return nil
}

func FileName(file *protogen.File) string {
	fname := path.Base(file.Proto.GetName())
	fname = strings.Replace(fname, ".proto", "", -1)
	fname = strings.Replace(fname, "-", "_", -1)
	fname = strings.Replace(fname, ".", "_", -1)
	return toCamelInitCase(fname, true)
}

var uppercaseAcronym = sync.Map{}

// Converts a string to CamelCase
func toCamelInitCase(s string, initCase bool) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	a, hasAcronym := uppercaseAcronym.Load(s)
	if hasAcronym {
		s = a.(string)
	}

	n := strings.Builder{}
	n.Grow(len(s))
	capNext := initCase
	prevIsCap := false
	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if capNext {
			if vIsLow {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsCap {
				v += 'a'
				v -= 'A'
			}
		} else if prevIsCap && vIsCap && !hasAcronym {
			v += 'a'
			v -= 'A'
		}
		prevIsCap = vIsCap

		if vIsCap || vIsLow {
			n.WriteByte(v)
			capNext = false
		} else if vIsNum := v >= '0' && v <= '9'; vIsNum {
			n.WriteByte(v)
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}
	return n.String()
}
