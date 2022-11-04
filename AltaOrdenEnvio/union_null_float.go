// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TransactionEvent.avsc
 */
package AltaOrdenEnvioEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullFloatTypeEnum int

const (
	UnionNullFloatTypeEnumFloat UnionNullFloatTypeEnum = 1
)

type UnionNullFloat struct {
	Null      *types.NullVal
	Float     float32
	UnionType UnionNullFloatTypeEnum
}

func writeUnionNullFloat(r *UnionNullFloat, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullFloatTypeEnumFloat:
		return vm.WriteFloat(r.Float, w)
	}
	return fmt.Errorf("invalid value for *UnionNullFloat")
}

func NewUnionNullFloat() *UnionNullFloat {
	return &UnionNullFloat{}
}

func (r *UnionNullFloat) Serialize(w io.Writer) error {
	return writeUnionNullFloat(r, w)
}

func DeserializeUnionNullFloat(r io.Reader) (*UnionNullFloat, error) {
	t := NewUnionNullFloat()
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

func DeserializeUnionNullFloatFromSchema(r io.Reader, schema string) (*UnionNullFloat, error) {
	t := NewUnionNullFloat()
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

func (r *UnionNullFloat) Schema() string {
	return "[\"null\",\"float\"]"
}

func (_ *UnionNullFloat) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullFloat) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullFloat) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullFloat) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullFloat) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullFloat) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullFloat) SetLong(v int64) {

	r.UnionType = (UnionNullFloatTypeEnum)(v)
}

func (r *UnionNullFloat) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		return &types.Float{Target: (&r.Float)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullFloat) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullFloat) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullFloat) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullFloat) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullFloat) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullFloat) Finalize()                        {}

func (r *UnionNullFloat) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullFloatTypeEnumFloat:
		return json.Marshal(map[string]interface{}{"float": r.Float})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullFloat")
}

func (r *UnionNullFloat) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["float"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Float)
	}
	return fmt.Errorf("invalid value for *UnionNullFloat")
}
