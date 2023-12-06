// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     UOHdrCreada.avsc
 */
package UOPublisherHdrEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullArrayTipoDeServicioTypeEnum int

const (
	UnionNullArrayTipoDeServicioTypeEnumArrayTipoDeServicio UnionNullArrayTipoDeServicioTypeEnum = 1
)

type UnionNullArrayTipoDeServicio struct {
	Null                *types.NullVal
	ArrayTipoDeServicio []TipoDeServicio
	UnionType           UnionNullArrayTipoDeServicioTypeEnum
}

func writeUnionNullArrayTipoDeServicio(r *UnionNullArrayTipoDeServicio, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullArrayTipoDeServicioTypeEnumArrayTipoDeServicio:
		return writeArrayTipoDeServicio(r.ArrayTipoDeServicio, w)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayTipoDeServicio")
}

func NewUnionNullArrayTipoDeServicio() *UnionNullArrayTipoDeServicio {
	return &UnionNullArrayTipoDeServicio{}
}

func (r *UnionNullArrayTipoDeServicio) Serialize(w io.Writer) error {
	return writeUnionNullArrayTipoDeServicio(r, w)
}

func DeserializeUnionNullArrayTipoDeServicio(r io.Reader) (*UnionNullArrayTipoDeServicio, error) {
	t := NewUnionNullArrayTipoDeServicio()
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

func DeserializeUnionNullArrayTipoDeServicioFromSchema(r io.Reader, schema string) (*UnionNullArrayTipoDeServicio, error) {
	t := NewUnionNullArrayTipoDeServicio()
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

func (r *UnionNullArrayTipoDeServicio) Schema() string {
	return "[\"null\",{\"items\":{\"fields\":[{\"name\":\"TipoDeServicioId\",\"type\":\"int\"},{\"name\":\"Nombre\",\"type\":\"string\"}],\"name\":\"TipoDeServicio\",\"namespace\":\"Andreani.UOPublisherHdr.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}]"
}

func (_ *UnionNullArrayTipoDeServicio) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullArrayTipoDeServicio) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullArrayTipoDeServicio) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullArrayTipoDeServicio) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullArrayTipoDeServicio) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullArrayTipoDeServicio) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullArrayTipoDeServicio) SetLong(v int64) {

	r.UnionType = (UnionNullArrayTipoDeServicioTypeEnum)(v)
}

func (r *UnionNullArrayTipoDeServicio) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ArrayTipoDeServicio = make([]TipoDeServicio, 0)
		return &ArrayTipoDeServicioWrapper{Target: (&r.ArrayTipoDeServicio)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayTipoDeServicio) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullArrayTipoDeServicio) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullArrayTipoDeServicio) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullArrayTipoDeServicio) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayTipoDeServicio) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullArrayTipoDeServicio) Finalize()                {}

func (r *UnionNullArrayTipoDeServicio) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullArrayTipoDeServicioTypeEnumArrayTipoDeServicio:
		return json.Marshal(map[string]interface{}{"array": r.ArrayTipoDeServicio})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullArrayTipoDeServicio")
}

func (r *UnionNullArrayTipoDeServicio) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ArrayTipoDeServicio)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayTipoDeServicio")
}
