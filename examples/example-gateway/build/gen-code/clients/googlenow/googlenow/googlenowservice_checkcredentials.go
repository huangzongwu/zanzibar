// Code generated by thriftrw v1.3.0
// @generated

package googlenow

import (
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type GoogleNowService_CheckCredentials_Args struct{}

func (v *GoogleNowService_CheckCredentials_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *GoogleNowService_CheckCredentials_Args) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *GoogleNowService_CheckCredentials_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [0]string
	i := 0
	return fmt.Sprintf("GoogleNowService_CheckCredentials_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *GoogleNowService_CheckCredentials_Args) Equals(rhs *GoogleNowService_CheckCredentials_Args) bool {
	return true
}

func (v *GoogleNowService_CheckCredentials_Args) MethodName() string {
	return "checkCredentials"
}

func (v *GoogleNowService_CheckCredentials_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var GoogleNowService_CheckCredentials_Helper = struct {
	Args           func() *GoogleNowService_CheckCredentials_Args
	IsException    func(error) bool
	WrapResponse   func(error) (*GoogleNowService_CheckCredentials_Result, error)
	UnwrapResponse func(*GoogleNowService_CheckCredentials_Result) error
}{}

func init() {
	GoogleNowService_CheckCredentials_Helper.Args = func() *GoogleNowService_CheckCredentials_Args {
		return &GoogleNowService_CheckCredentials_Args{}
	}
	GoogleNowService_CheckCredentials_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	GoogleNowService_CheckCredentials_Helper.WrapResponse = func(err error) (*GoogleNowService_CheckCredentials_Result, error) {
		if err == nil {
			return &GoogleNowService_CheckCredentials_Result{}, nil
		}
		return nil, err
	}
	GoogleNowService_CheckCredentials_Helper.UnwrapResponse = func(result *GoogleNowService_CheckCredentials_Result) (err error) {
		return
	}
}

type GoogleNowService_CheckCredentials_Result struct{}

func (v *GoogleNowService_CheckCredentials_Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *GoogleNowService_CheckCredentials_Result) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *GoogleNowService_CheckCredentials_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [0]string
	i := 0
	return fmt.Sprintf("GoogleNowService_CheckCredentials_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *GoogleNowService_CheckCredentials_Result) Equals(rhs *GoogleNowService_CheckCredentials_Result) bool {
	return true
}

func (v *GoogleNowService_CheckCredentials_Result) MethodName() string {
	return "checkCredentials"
}

func (v *GoogleNowService_CheckCredentials_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
