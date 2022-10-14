// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoDeAlmacenSCE.avsc
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

type UnionNullListaDePropiedadesTypeEnum int

const (
	UnionNullListaDePropiedadesTypeEnumListaDePropiedades UnionNullListaDePropiedadesTypeEnum = 1
)

type UnionNullListaDePropiedades struct {
	Null               *types.NullVal
	ListaDePropiedades ListaDePropiedades
	UnionType          UnionNullListaDePropiedadesTypeEnum
}

func writeUnionNullListaDePropiedades(r *UnionNullListaDePropiedades, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullListaDePropiedadesTypeEnumListaDePropiedades:
		return writeListaDePropiedades(r.ListaDePropiedades, w)
	}
	return fmt.Errorf("invalid value for *UnionNullListaDePropiedades")
}

func NewUnionNullListaDePropiedades() *UnionNullListaDePropiedades {
	return &UnionNullListaDePropiedades{}
}

func (r *UnionNullListaDePropiedades) Serialize(w io.Writer) error {
	return writeUnionNullListaDePropiedades(r, w)
}

func DeserializeUnionNullListaDePropiedades(r io.Reader) (*UnionNullListaDePropiedades, error) {
	t := NewUnionNullListaDePropiedades()
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

func DeserializeUnionNullListaDePropiedadesFromSchema(r io.Reader, schema string) (*UnionNullListaDePropiedades, error) {
	t := NewUnionNullListaDePropiedades()
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

func (r *UnionNullListaDePropiedades) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]"
}

func (_ *UnionNullListaDePropiedades) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullListaDePropiedades) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullListaDePropiedades) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullListaDePropiedades) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullListaDePropiedades) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullListaDePropiedades) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullListaDePropiedades) SetLong(v int64) {

	r.UnionType = (UnionNullListaDePropiedadesTypeEnum)(v)
}

func (r *UnionNullListaDePropiedades) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ListaDePropiedades = NewListaDePropiedades()
		return &types.Record{Target: (&r.ListaDePropiedades)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullListaDePropiedades) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullListaDePropiedades) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullListaDePropiedades) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullListaDePropiedades) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullListaDePropiedades) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullListaDePropiedades) Finalize()                {}

func (r *UnionNullListaDePropiedades) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullListaDePropiedadesTypeEnumListaDePropiedades:
		return json.Marshal(map[string]interface{}{"Andreani.Wapv2.Events.Record.ListaDePropiedades": r.ListaDePropiedades})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullListaDePropiedades")
}

func (r *UnionNullListaDePropiedades) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.Wapv2.Events.Record.ListaDePropiedades"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ListaDePropiedades)
	}
	return fmt.Errorf("invalid value for *UnionNullListaDePropiedades")
}
