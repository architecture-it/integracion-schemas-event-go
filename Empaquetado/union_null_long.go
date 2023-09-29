// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoEmpaquetar.avsc
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

type UnionNullLongTypeEnum int

const (
	UnionNullLongTypeEnumLong UnionNullLongTypeEnum = 1
)

type UnionNullLong struct {
	Null      *types.NullVal
	Long      int64
	UnionType UnionNullLongTypeEnum
}

func writeUnionNullLong(r *UnionNullLong, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullLongTypeEnumLong:
		return vm.WriteLong(r.Long, w)
	}
	return fmt.Errorf("invalid value for *UnionNullLong")
}

func NewUnionNullLong() *UnionNullLong {
	return &UnionNullLong{}
}

func (r *UnionNullLong) Serialize(w io.Writer) error {
	return writeUnionNullLong(r, w)
}

func DeserializeUnionNullLong(r io.Reader) (*UnionNullLong, error) {
	t := NewUnionNullLong()
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

func DeserializeUnionNullLongFromSchema(r io.Reader, schema string) (*UnionNullLong, error) {
	t := NewUnionNullLong()
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

func (r *UnionNullLong) Schema() string {
	return "[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]"
}

func (_ *UnionNullLong) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullLong) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullLong) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullLong) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullLong) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullLong) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullLong) SetLong(v int64) {

	r.UnionType = (UnionNullLongTypeEnum)(v)
}

func (r *UnionNullLong) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		return &types.Long{Target: (&r.Long)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullLong) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullLong) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullLong) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullLong) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullLong) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullLong) Finalize()                        {}

func (r *UnionNullLong) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullLongTypeEnumLong:
		return json.Marshal(map[string]interface{}{"long": r.Long})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullLong")
}

func (r *UnionNullLong) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["long"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Long)
	}
	return fmt.Errorf("invalid value for *UnionNullLong")
}
