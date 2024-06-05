// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     OrdenDeKittingSolicitadaV2.avsc
 */
package Wapv2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullArrayStringTypeEnum int

const (
	UnionNullArrayStringTypeEnumArrayString UnionNullArrayStringTypeEnum = 1
)

type UnionNullArrayString struct {
	Null        *types.NullVal
	ArrayString []string
	UnionType   UnionNullArrayStringTypeEnum
}

func writeUnionNullArrayString(r *UnionNullArrayString, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullArrayStringTypeEnumArrayString:
		return writeArrayString(r.ArrayString, w)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayString")
}

func NewUnionNullArrayString() *UnionNullArrayString {
	return &UnionNullArrayString{}
}

func (r *UnionNullArrayString) Serialize(w io.Writer) error {
	return writeUnionNullArrayString(r, w)
}

func DeserializeUnionNullArrayString(r io.Reader) (*UnionNullArrayString, error) {
	t := NewUnionNullArrayString()
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

func DeserializeUnionNullArrayStringFromSchema(r io.Reader, schema string) (*UnionNullArrayString, error) {
	t := NewUnionNullArrayString()
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

func (r *UnionNullArrayString) Schema() string {
	return "[\"null\",{\"items\":\"string\",\"type\":\"array\"}]"
}

func (_ *UnionNullArrayString) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullArrayString) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullArrayString) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullArrayString) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullArrayString) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullArrayString) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullArrayString) SetLong(v int64) {

	r.UnionType = (UnionNullArrayStringTypeEnum)(v)
}

func (r *UnionNullArrayString) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ArrayString = make([]string, 0)
		return &ArrayStringWrapper{Target: (&r.ArrayString)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayString) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullArrayString) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullArrayString) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullArrayString) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullArrayString) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullArrayString) Finalize()                        {}

func (r *UnionNullArrayString) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullArrayStringTypeEnumArrayString:
		return json.Marshal(map[string]interface{}{"array": r.ArrayString})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullArrayString")
}

func (r *UnionNullArrayString) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ArrayString)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayString")
}
