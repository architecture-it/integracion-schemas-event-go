// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Incidencia.avsc
 */
package appnameEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullIntTypeEnum int

const (
	UnionNullIntTypeEnumInt UnionNullIntTypeEnum = 1
)

type UnionNullInt struct {
	Null      *types.NullVal
	Int       int32
	UnionType UnionNullIntTypeEnum
}

func writeUnionNullInt(r *UnionNullInt, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullIntTypeEnumInt:
		return vm.WriteInt(r.Int, w)
	}
	return fmt.Errorf("invalid value for *UnionNullInt")
}

func NewUnionNullInt() *UnionNullInt {
	return &UnionNullInt{}
}

func (r *UnionNullInt) Serialize(w io.Writer) error {
	return writeUnionNullInt(r, w)
}

func DeserializeUnionNullInt(r io.Reader) (*UnionNullInt, error) {
	t := NewUnionNullInt()
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

func DeserializeUnionNullIntFromSchema(r io.Reader, schema string) (*UnionNullInt, error) {
	t := NewUnionNullInt()
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

func (r *UnionNullInt) Schema() string {
	return "[\"null\",\"int\"]"
}

func (_ *UnionNullInt) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullInt) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullInt) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullInt) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullInt) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullInt) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullInt) SetLong(v int64) {

	r.UnionType = (UnionNullIntTypeEnum)(v)
}

func (r *UnionNullInt) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		return &types.Int{Target: (&r.Int)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullInt) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullInt) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullInt) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullInt) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullInt) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullInt) Finalize()                        {}

func (r *UnionNullInt) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullIntTypeEnumInt:
		return json.Marshal(map[string]interface{}{"int": r.Int})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullInt")
}

func (r *UnionNullInt) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["int"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Int)
	}
	return fmt.Errorf("invalid value for *UnionNullInt")
}
