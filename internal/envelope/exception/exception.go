// Code generated by thriftrw v1.20.2. DO NOT EDIT.
// @generated

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package exception

import (
	bytes "bytes"
	json "encoding/json"
	fmt "fmt"
	multierr "go.uber.org/multierr"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	math "math"
	strconv "strconv"
	strings "strings"
)

type ExceptionType int32

const (
	ExceptionTypeUnknown               ExceptionType = 0
	ExceptionTypeUnknownMethod         ExceptionType = 1
	ExceptionTypeInvalidMessageType    ExceptionType = 2
	ExceptionTypeWrongMethodName       ExceptionType = 3
	ExceptionTypeBadSequenceID         ExceptionType = 4
	ExceptionTypeMissingResult         ExceptionType = 5
	ExceptionTypeInternalError         ExceptionType = 6
	ExceptionTypeProtocolError         ExceptionType = 7
	ExceptionTypeInvalidTransform      ExceptionType = 8
	ExceptionTypeInvalidProtocol       ExceptionType = 9
	ExceptionTypeUnsupportedClientType ExceptionType = 10
)

// ExceptionType_Values returns all recognized values of ExceptionType.
func ExceptionType_Values() []ExceptionType {
	return []ExceptionType{
		ExceptionTypeUnknown,
		ExceptionTypeUnknownMethod,
		ExceptionTypeInvalidMessageType,
		ExceptionTypeWrongMethodName,
		ExceptionTypeBadSequenceID,
		ExceptionTypeMissingResult,
		ExceptionTypeInternalError,
		ExceptionTypeProtocolError,
		ExceptionTypeInvalidTransform,
		ExceptionTypeInvalidProtocol,
		ExceptionTypeUnsupportedClientType,
	}
}

// UnmarshalText tries to decode ExceptionType from a byte slice
// containing its name.
//
//   var v ExceptionType
//   err := v.UnmarshalText([]byte("UNKNOWN"))
func (v *ExceptionType) UnmarshalText(value []byte) error {
	switch s := string(value); s {
	case "UNKNOWN":
		*v = ExceptionTypeUnknown
		return nil
	case "UNKNOWN_METHOD":
		*v = ExceptionTypeUnknownMethod
		return nil
	case "INVALID_MESSAGE_TYPE":
		*v = ExceptionTypeInvalidMessageType
		return nil
	case "WRONG_METHOD_NAME":
		*v = ExceptionTypeWrongMethodName
		return nil
	case "BAD_SEQUENCE_ID":
		*v = ExceptionTypeBadSequenceID
		return nil
	case "MISSING_RESULT":
		*v = ExceptionTypeMissingResult
		return nil
	case "INTERNAL_ERROR":
		*v = ExceptionTypeInternalError
		return nil
	case "PROTOCOL_ERROR":
		*v = ExceptionTypeProtocolError
		return nil
	case "INVALID_TRANSFORM":
		*v = ExceptionTypeInvalidTransform
		return nil
	case "INVALID_PROTOCOL":
		*v = ExceptionTypeInvalidProtocol
		return nil
	case "UNSUPPORTED_CLIENT_TYPE":
		*v = ExceptionTypeUnsupportedClientType
		return nil
	default:
		val, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return fmt.Errorf("unknown enum value %q for %q: %v", s, "ExceptionType", err)
		}
		*v = ExceptionType(val)
		return nil
	}
}

// MarshalText encodes ExceptionType to text.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements the TextMarshaler interface.
func (v ExceptionType) MarshalText() ([]byte, error) {
	switch int32(v) {
	case 0:
		return []byte("UNKNOWN"), nil
	case 1:
		return []byte("UNKNOWN_METHOD"), nil
	case 2:
		return []byte("INVALID_MESSAGE_TYPE"), nil
	case 3:
		return []byte("WRONG_METHOD_NAME"), nil
	case 4:
		return []byte("BAD_SEQUENCE_ID"), nil
	case 5:
		return []byte("MISSING_RESULT"), nil
	case 6:
		return []byte("INTERNAL_ERROR"), nil
	case 7:
		return []byte("PROTOCOL_ERROR"), nil
	case 8:
		return []byte("INVALID_TRANSFORM"), nil
	case 9:
		return []byte("INVALID_PROTOCOL"), nil
	case 10:
		return []byte("UNSUPPORTED_CLIENT_TYPE"), nil
	}
	return []byte(strconv.FormatInt(int64(v), 10)), nil
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ExceptionType.
// Enums are logged as objects, where the value is logged with key "value", and
// if this value's name is known, the name is logged with key "name".
func (v ExceptionType) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt32("value", int32(v))
	switch int32(v) {
	case 0:
		enc.AddString("name", "UNKNOWN")
	case 1:
		enc.AddString("name", "UNKNOWN_METHOD")
	case 2:
		enc.AddString("name", "INVALID_MESSAGE_TYPE")
	case 3:
		enc.AddString("name", "WRONG_METHOD_NAME")
	case 4:
		enc.AddString("name", "BAD_SEQUENCE_ID")
	case 5:
		enc.AddString("name", "MISSING_RESULT")
	case 6:
		enc.AddString("name", "INTERNAL_ERROR")
	case 7:
		enc.AddString("name", "PROTOCOL_ERROR")
	case 8:
		enc.AddString("name", "INVALID_TRANSFORM")
	case 9:
		enc.AddString("name", "INVALID_PROTOCOL")
	case 10:
		enc.AddString("name", "UNSUPPORTED_CLIENT_TYPE")
	}
	return nil
}

// Ptr returns a pointer to this enum value.
func (v ExceptionType) Ptr() *ExceptionType {
	return &v
}

// ToWire translates ExceptionType into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// Enums are represented as 32-bit integers over the wire.
func (v ExceptionType) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

// FromWire deserializes ExceptionType from its Thrift-level
// representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TI32)
//   if err != nil {
//     return ExceptionType(0), err
//   }
//
//   var v ExceptionType
//   if err := v.FromWire(x); err != nil {
//     return ExceptionType(0), err
//   }
//   return v, nil
func (v *ExceptionType) FromWire(w wire.Value) error {
	*v = (ExceptionType)(w.GetI32())
	return nil
}

// String returns a readable string representation of ExceptionType.
func (v ExceptionType) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "UNKNOWN"
	case 1:
		return "UNKNOWN_METHOD"
	case 2:
		return "INVALID_MESSAGE_TYPE"
	case 3:
		return "WRONG_METHOD_NAME"
	case 4:
		return "BAD_SEQUENCE_ID"
	case 5:
		return "MISSING_RESULT"
	case 6:
		return "INTERNAL_ERROR"
	case 7:
		return "PROTOCOL_ERROR"
	case 8:
		return "INVALID_TRANSFORM"
	case 9:
		return "INVALID_PROTOCOL"
	case 10:
		return "UNSUPPORTED_CLIENT_TYPE"
	}
	return fmt.Sprintf("ExceptionType(%d)", w)
}

// Equals returns true if this ExceptionType value matches the provided
// value.
func (v ExceptionType) Equals(rhs ExceptionType) bool {
	return v == rhs
}

// MarshalJSON serializes ExceptionType into JSON.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements json.Marshaler.
func (v ExceptionType) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"UNKNOWN\""), nil
	case 1:
		return ([]byte)("\"UNKNOWN_METHOD\""), nil
	case 2:
		return ([]byte)("\"INVALID_MESSAGE_TYPE\""), nil
	case 3:
		return ([]byte)("\"WRONG_METHOD_NAME\""), nil
	case 4:
		return ([]byte)("\"BAD_SEQUENCE_ID\""), nil
	case 5:
		return ([]byte)("\"MISSING_RESULT\""), nil
	case 6:
		return ([]byte)("\"INTERNAL_ERROR\""), nil
	case 7:
		return ([]byte)("\"PROTOCOL_ERROR\""), nil
	case 8:
		return ([]byte)("\"INVALID_TRANSFORM\""), nil
	case 9:
		return ([]byte)("\"INVALID_PROTOCOL\""), nil
	case 10:
		return ([]byte)("\"UNSUPPORTED_CLIENT_TYPE\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

// UnmarshalJSON attempts to decode ExceptionType from its JSON
// representation.
//
// This implementation supports both, numeric and string inputs. If a
// string is provided, it must be a known enum name.
//
// This implements json.Unmarshaler.
func (v *ExceptionType) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "ExceptionType")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "ExceptionType")
		}
		*v = (ExceptionType)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "ExceptionType")
	}
}

type TApplicationException struct {
	Message *string        `json:"message,omitempty"`
	Type    *ExceptionType `json:"type,omitempty"`
}

// ToWire translates a TApplicationException struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *TApplicationException) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Message != nil {
		w, err = wire.NewValueString(*(v.Message)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.Type != nil {
		w, err = v.Type.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ExceptionType_Read(w wire.Value) (ExceptionType, error) {
	var v ExceptionType
	err := v.FromWire(w)
	return v, err
}

// FromWire deserializes a TApplicationException struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a TApplicationException struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v TApplicationException
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *TApplicationException) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Message = &x
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TI32 {
				var x ExceptionType
				x, err = _ExceptionType_Read(field.Value)
				v.Type = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a TApplicationException
// struct.
func (v *TApplicationException) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.Message != nil {
		fields[i] = fmt.Sprintf("Message: %v", *(v.Message))
		i++
	}
	if v.Type != nil {
		fields[i] = fmt.Sprintf("Type: %v", *(v.Type))
		i++
	}

	return fmt.Sprintf("TApplicationException{%v}", strings.Join(fields[:i], ", "))
}

// ErrorName is the name of this type as defined in the Thrift
// file.
func (*TApplicationException) ErrorName() string {
	return "TApplicationException"
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _ExceptionType_EqualsPtr(lhs, rhs *ExceptionType) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return x.Equals(y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this TApplicationException match the
// provided TApplicationException.
//
// This function performs a deep comparison.
func (v *TApplicationException) Equals(rhs *TApplicationException) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Message, rhs.Message) {
		return false
	}
	if !_ExceptionType_EqualsPtr(v.Type, rhs.Type) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of TApplicationException.
func (v *TApplicationException) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Message != nil {
		enc.AddString("message", *v.Message)
	}
	if v.Type != nil {
		err = multierr.Append(err, enc.AddObject("type", *v.Type))
	}
	return err
}

// GetMessage returns the value of Message if it is set or its
// zero value if it is unset.
func (v *TApplicationException) GetMessage() (o string) {
	if v != nil && v.Message != nil {
		return *v.Message
	}

	return
}

// IsSetMessage returns true if Message is not nil.
func (v *TApplicationException) IsSetMessage() bool {
	return v != nil && v.Message != nil
}

// GetType returns the value of Type if it is set or its
// zero value if it is unset.
func (v *TApplicationException) GetType() (o ExceptionType) {
	if v != nil && v.Type != nil {
		return *v.Type
	}

	return
}

// IsSetType returns true if Type is not nil.
func (v *TApplicationException) IsSetType() bool {
	return v != nil && v.Type != nil
}

func (v *TApplicationException) Error() string {
	return v.String()
}

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "exception",
	Package:  "go.uber.org/thriftrw/internal/envelope/exception",
	FilePath: "exception.thrift",
	SHA1:     "88105bcd404d4aee06542af9452f7cf76647ae98",
	Raw:      rawIDL,
}

const rawIDL = "enum ExceptionType {\n  UNKNOWN = 0\n  UNKNOWN_METHOD = 1\n  INVALID_MESSAGE_TYPE = 2\n  WRONG_METHOD_NAME = 3\n  BAD_SEQUENCE_ID = 4\n  MISSING_RESULT = 5\n  INTERNAL_ERROR = 6\n  PROTOCOL_ERROR = 7\n  INVALID_TRANSFORM = 8\n  INVALID_PROTOCOL = 9\n  UNSUPPORTED_CLIENT_TYPE = 10\n}\n\nexception TApplicationException {\n  1: optional string message\n  2: optional ExceptionType type\n}\n"
