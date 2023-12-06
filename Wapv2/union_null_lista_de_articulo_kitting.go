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

type UnionNullListaDeArticuloKittingTypeEnum int

const (
	UnionNullListaDeArticuloKittingTypeEnumListaDeArticuloKitting UnionNullListaDeArticuloKittingTypeEnum = 1
)

type UnionNullListaDeArticuloKitting struct {
	Null                   *types.NullVal
	ListaDeArticuloKitting ListaDeArticuloKitting
	UnionType              UnionNullListaDeArticuloKittingTypeEnum
}

func writeUnionNullListaDeArticuloKitting(r *UnionNullListaDeArticuloKitting, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullListaDeArticuloKittingTypeEnumListaDeArticuloKitting:
		return writeListaDeArticuloKitting(r.ListaDeArticuloKitting, w)
	}
	return fmt.Errorf("invalid value for *UnionNullListaDeArticuloKitting")
}

func NewUnionNullListaDeArticuloKitting() *UnionNullListaDeArticuloKitting {
	return &UnionNullListaDeArticuloKitting{}
}

func (r *UnionNullListaDeArticuloKitting) Serialize(w io.Writer) error {
	return writeUnionNullListaDeArticuloKitting(r, w)
}

func DeserializeUnionNullListaDeArticuloKitting(r io.Reader) (*UnionNullListaDeArticuloKitting, error) {
	t := NewUnionNullListaDeArticuloKitting()
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

func DeserializeUnionNullListaDeArticuloKittingFromSchema(r io.Reader, schema string) (*UnionNullListaDeArticuloKitting, error) {
	t := NewUnionNullListaDeArticuloKitting()
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

func (r *UnionNullListaDeArticuloKitting) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"listaDeArticuloKitting\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"cantidad\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"loteFabricante\",\"type\":\"string\"},{\"name\":\"loteSecundario\",\"type\":\"string\"},{\"name\":\"loteAlmacen\",\"type\":\"string\"},{\"name\":\"loteEstado\",\"type\":\"string\"}],\"name\":\"ArticuloKitting\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDeArticuloKitting\",\"type\":\"record\"}]"
}

func (_ *UnionNullListaDeArticuloKitting) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullListaDeArticuloKitting) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullListaDeArticuloKitting) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullListaDeArticuloKitting) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullListaDeArticuloKitting) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullListaDeArticuloKitting) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullListaDeArticuloKitting) SetLong(v int64) {

	r.UnionType = (UnionNullListaDeArticuloKittingTypeEnum)(v)
}

func (r *UnionNullListaDeArticuloKitting) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ListaDeArticuloKitting = NewListaDeArticuloKitting()
		return &types.Record{Target: (&r.ListaDeArticuloKitting)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullListaDeArticuloKitting) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullListaDeArticuloKitting) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullListaDeArticuloKitting) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullListaDeArticuloKitting) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullListaDeArticuloKitting) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullListaDeArticuloKitting) Finalize()                {}

func (r *UnionNullListaDeArticuloKitting) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullListaDeArticuloKittingTypeEnumListaDeArticuloKitting:
		return json.Marshal(map[string]interface{}{"Andreani.Wapv2.Events.Record.ListaDeArticuloKitting": r.ListaDeArticuloKitting})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullListaDeArticuloKitting")
}

func (r *UnionNullListaDeArticuloKitting) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.Wapv2.Events.Record.ListaDeArticuloKitting"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ListaDeArticuloKitting)
	}
	return fmt.Errorf("invalid value for *UnionNullListaDeArticuloKitting")
}