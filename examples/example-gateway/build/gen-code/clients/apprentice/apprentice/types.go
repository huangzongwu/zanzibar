// Code generated by thriftrw v1.12.0. DO NOT EDIT.
// @generated

package apprentice

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type OperationError struct {
	Message string `json:"message,required"`
}

// ToWire translates a OperationError struct into a Thrift-level intermediate
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
func (v *OperationError) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a OperationError struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a OperationError struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v OperationError
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *OperationError) FromWire(w wire.Value) error {
	var err error

	messageIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		}
	}

	if !messageIsSet {
		return errors.New("field Message of OperationError is required")
	}

	return nil
}

// String returns a readable string representation of a OperationError
// struct.
func (v *OperationError) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++

	return fmt.Sprintf("OperationError{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this OperationError match the
// provided OperationError.
//
// This function performs a deep comparison.
func (v *OperationError) Equals(rhs *OperationError) bool {
	if !(v.Message == rhs.Message) {
		return false
	}

	return true
}

// GetMessage returns the value of Message if it is set or its
// zero value if it is unset.
func (v *OperationError) GetMessage() (o string) { return v.Message }

func (v *OperationError) Error() string {
	return v.String()
}