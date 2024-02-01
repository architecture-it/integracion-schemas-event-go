// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ValidacionAtributosDeLote.avsc
 */
package EmpaquetadoEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullBoolTypeEnum int

const (
	UnionNullBoolTypeEnumBool UnionNullBoolTypeEnum = 1
)

type UnionNullBool struct {
	Null      *types.NullVal
	Bool      bool
	UnionType UnionNullBoolTypeEnum
}

func writeUnionNullBool(r *UnionNullBool, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullBoolTypeEnumBool:
		return vm.WriteBool(r.Bool, w)
	}
	return fmt.Errorf("invalid value for *UnionNullBool")
}

func NewUnionNullBool() *UnionNullBool {
	return &UnionNullBool{}
}

func (r *UnionNullBool) Serialize(w io.Writer) error {
	return writeUnionNullBool(r, w)
}

func DeserializeUnionNullBool(r io.Reader) (*UnionNullBool, error) {
	t := NewUnionNullBool()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionNullBoolFromSchema(r io.Reader, schema string) (*UnionNullBool, error) {
	t := NewUnionNullBool()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r *UnionNullBool) Schema() string {
	return "[\"null\",\"boolean\"]"
}

func (_ *UnionNullBool) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullBool) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullBool) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullBool) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullBool) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullBool) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullBool) SetLong(v int64) {

	r.UnionType = (UnionNullBoolTypeEnum)(v)
}

func (r *UnionNullBool) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		return &types.Boolean{Target: (&r.Bool)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullBool) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullBool) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullBool) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullBool) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullBool) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullBool) Finalize()                        {}

func (r *UnionNullBool) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullBoolTypeEnumBool:
		return json.Marshal(map[string]interface{}{"boolean": r.Bool})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullBool")
}

func (r *UnionNullBool) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["boolean"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Bool)
	}
	return fmt.Errorf("invalid value for *UnionNullBool")
}
