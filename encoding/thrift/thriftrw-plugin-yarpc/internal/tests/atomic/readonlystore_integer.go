// Code generated by thriftrw v1.5.0. DO NOT EDIT.
// @generated

package atomic

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type ReadOnlyStore_Integer_Args struct {
	Key *string `json:"key,omitempty"`
}

func (v *ReadOnlyStore_Integer_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Key != nil {
		w, err = wire.NewValueString(*(v.Key)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *ReadOnlyStore_Integer_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Key = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *ReadOnlyStore_Integer_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Key != nil {
		fields[i] = fmt.Sprintf("Key: %v", *(v.Key))
		i++
	}
	return fmt.Sprintf("ReadOnlyStore_Integer_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *ReadOnlyStore_Integer_Args) Equals(rhs *ReadOnlyStore_Integer_Args) bool {
	if !_String_EqualsPtr(v.Key, rhs.Key) {
		return false
	}
	return true
}

func (v *ReadOnlyStore_Integer_Args) MethodName() string {
	return "integer"
}

func (v *ReadOnlyStore_Integer_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var ReadOnlyStore_Integer_Helper = struct {
	Args           func(key *string) *ReadOnlyStore_Integer_Args
	IsException    func(error) bool
	WrapResponse   func(int64, error) (*ReadOnlyStore_Integer_Result, error)
	UnwrapResponse func(*ReadOnlyStore_Integer_Result) (int64, error)
}{}

func init() {
	ReadOnlyStore_Integer_Helper.Args = func(key *string) *ReadOnlyStore_Integer_Args {
		return &ReadOnlyStore_Integer_Args{Key: key}
	}
	ReadOnlyStore_Integer_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *KeyDoesNotExist:
			return true
		default:
			return false
		}
	}
	ReadOnlyStore_Integer_Helper.WrapResponse = func(success int64, err error) (*ReadOnlyStore_Integer_Result, error) {
		if err == nil {
			return &ReadOnlyStore_Integer_Result{Success: &success}, nil
		}
		switch e := err.(type) {
		case *KeyDoesNotExist:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for ReadOnlyStore_Integer_Result.DoesNotExist")
			}
			return &ReadOnlyStore_Integer_Result{DoesNotExist: e}, nil
		}
		return nil, err
	}
	ReadOnlyStore_Integer_Helper.UnwrapResponse = func(result *ReadOnlyStore_Integer_Result) (success int64, err error) {
		if result.DoesNotExist != nil {
			err = result.DoesNotExist
			return
		}
		if result.Success != nil {
			success = *result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type ReadOnlyStore_Integer_Result struct {
	Success      *int64           `json:"success,omitempty"`
	DoesNotExist *KeyDoesNotExist `json:"doesNotExist,omitempty"`
}

func (v *ReadOnlyStore_Integer_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = wire.NewValueI64(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.DoesNotExist != nil {
		w, err = v.DoesNotExist.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("ReadOnlyStore_Integer_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _KeyDoesNotExist_Read(w wire.Value) (*KeyDoesNotExist, error) {
	var v KeyDoesNotExist
	err := v.FromWire(w)
	return &v, err
}

func (v *ReadOnlyStore_Integer_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.Success = &x
				if err != nil {
					return err
				}
			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.DoesNotExist, err = _KeyDoesNotExist_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if v.DoesNotExist != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("ReadOnlyStore_Integer_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *ReadOnlyStore_Integer_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}
	if v.DoesNotExist != nil {
		fields[i] = fmt.Sprintf("DoesNotExist: %v", v.DoesNotExist)
		i++
	}
	return fmt.Sprintf("ReadOnlyStore_Integer_Result{%v}", strings.Join(fields[:i], ", "))
}

func _I64_EqualsPtr(lhs, rhs *int64) bool {
	if lhs != nil && rhs != nil {
		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func (v *ReadOnlyStore_Integer_Result) Equals(rhs *ReadOnlyStore_Integer_Result) bool {
	if !_I64_EqualsPtr(v.Success, rhs.Success) {
		return false
	}
	if !((v.DoesNotExist == nil && rhs.DoesNotExist == nil) || (v.DoesNotExist != nil && rhs.DoesNotExist != nil && v.DoesNotExist.Equals(rhs.DoesNotExist))) {
		return false
	}
	return true
}

func (v *ReadOnlyStore_Integer_Result) MethodName() string {
	return "integer"
}

func (v *ReadOnlyStore_Integer_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
