// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoDeAlmacenSCE.avsc
 */
package WarehousePedidoEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullListaDeTiposDeAcondicionamientoSecundarioTypeEnum int

const (
	UnionNullListaDeTiposDeAcondicionamientoSecundarioTypeEnumListaDeTiposDeAcondicionamientoSecundario UnionNullListaDeTiposDeAcondicionamientoSecundarioTypeEnum = 1
)

type UnionNullListaDeTiposDeAcondicionamientoSecundario struct {
	Null                                      *types.NullVal
	ListaDeTiposDeAcondicionamientoSecundario ListaDeTiposDeAcondicionamientoSecundario
	UnionType                                 UnionNullListaDeTiposDeAcondicionamientoSecundarioTypeEnum
}

func writeUnionNullListaDeTiposDeAcondicionamientoSecundario(r *UnionNullListaDeTiposDeAcondicionamientoSecundario, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullListaDeTiposDeAcondicionamientoSecundarioTypeEnumListaDeTiposDeAcondicionamientoSecundario:
		return writeListaDeTiposDeAcondicionamientoSecundario(r.ListaDeTiposDeAcondicionamientoSecundario, w)
	}
	return fmt.Errorf("invalid value for *UnionNullListaDeTiposDeAcondicionamientoSecundario")
}

func NewUnionNullListaDeTiposDeAcondicionamientoSecundario() *UnionNullListaDeTiposDeAcondicionamientoSecundario {
	return &UnionNullListaDeTiposDeAcondicionamientoSecundario{}
}

func (r *UnionNullListaDeTiposDeAcondicionamientoSecundario) Serialize(w io.Writer) error {
	return writeUnionNullListaDeTiposDeAcondicionamientoSecundario(r, w)
}

func DeserializeUnionNullListaDeTiposDeAcondicionamientoSecundario(r io.Reader) (*UnionNullListaDeTiposDeAcondicionamientoSecundario, error) {
	t := NewUnionNullListaDeTiposDeAcondicionamientoSecundario()
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

func DeserializeUnionNullListaDeTiposDeAcondicionamientoSecundarioFromSchema(r io.Reader, schema string) (*UnionNullListaDeTiposDeAcondicionamientoSecundario, error) {
	t := NewUnionNullListaDeTiposDeAcondicionamientoSecundario()
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

func (r *UnionNullListaDeTiposDeAcondicionamientoSecundario) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"tiposDeAcondicionamientoSecundario\",\"type\":{\"items\":{\"fields\":[{\"name\":\"tipoDeAcondi\",\"type\":\"string\"}],\"name\":\"TiposDeAcondicionamientoSecundario\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDeTiposDeAcondicionamientoSecundario\",\"type\":\"record\"}]"
}

func (_ *UnionNullListaDeTiposDeAcondicionamientoSecundario) SetBoolean(v bool) {
	panic("Unsupported operation")
}
func (_ *UnionNullListaDeTiposDeAcondicionamientoSecundario) SetInt(v int32) {
	panic("Unsupported operation")
}
func (_ *UnionNullListaDeTiposDeAcondicionamientoSecundario) SetFloat(v float32) {
	panic("Unsupported operation")
}
func (_ *UnionNullListaDeTiposDeAcondicionamientoSecundario) SetDouble(v float64) {
	panic("Unsupported operation")
}
func (_ *UnionNullListaDeTiposDeAcondicionamientoSecundario) SetBytes(v []byte) {
	panic("Unsupported operation")
}
func (_ *UnionNullListaDeTiposDeAcondicionamientoSecundario) SetString(v string) {
	panic("Unsupported operation")
}

func (r *UnionNullListaDeTiposDeAcondicionamientoSecundario) SetLong(v int64) {

	r.UnionType = (UnionNullListaDeTiposDeAcondicionamientoSecundarioTypeEnum)(v)
}

func (r *UnionNullListaDeTiposDeAcondicionamientoSecundario) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ListaDeTiposDeAcondicionamientoSecundario = NewListaDeTiposDeAcondicionamientoSecundario()
		return &types.Record{Target: (&r.ListaDeTiposDeAcondicionamientoSecundario)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullListaDeTiposDeAcondicionamientoSecundario) NullField(i int) {
	panic("Unsupported operation")
}
func (_ *UnionNullListaDeTiposDeAcondicionamientoSecundario) HintSize(i int) {
	panic("Unsupported operation")
}
func (_ *UnionNullListaDeTiposDeAcondicionamientoSecundario) SetDefault(i int) {
	panic("Unsupported operation")
}
func (_ *UnionNullListaDeTiposDeAcondicionamientoSecundario) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullListaDeTiposDeAcondicionamientoSecundario) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullListaDeTiposDeAcondicionamientoSecundario) Finalize() {}

func (r *UnionNullListaDeTiposDeAcondicionamientoSecundario) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullListaDeTiposDeAcondicionamientoSecundarioTypeEnumListaDeTiposDeAcondicionamientoSecundario:
		return json.Marshal(map[string]interface{}{"Andreani.WarehousePedido.Events.Record.ListaDeTiposDeAcondicionamientoSecundario": r.ListaDeTiposDeAcondicionamientoSecundario})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullListaDeTiposDeAcondicionamientoSecundario")
}

func (r *UnionNullListaDeTiposDeAcondicionamientoSecundario) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.WarehousePedido.Events.Record.ListaDeTiposDeAcondicionamientoSecundario"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ListaDeTiposDeAcondicionamientoSecundario)
	}
	return fmt.Errorf("invalid value for *UnionNullListaDeTiposDeAcondicionamientoSecundario")
}
