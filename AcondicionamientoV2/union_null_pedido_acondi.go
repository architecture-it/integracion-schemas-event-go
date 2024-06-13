// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     AcondicionamientoV2.avsc
 */
package AcondicionamientoV2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullPedidoAcondiTypeEnum int

const (
	UnionNullPedidoAcondiTypeEnumPedidoAcondi UnionNullPedidoAcondiTypeEnum = 1
)

type UnionNullPedidoAcondi struct {
	Null         *types.NullVal
	PedidoAcondi PedidoAcondi
	UnionType    UnionNullPedidoAcondiTypeEnum
}

func writeUnionNullPedidoAcondi(r *UnionNullPedidoAcondi, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullPedidoAcondiTypeEnumPedidoAcondi:
		return writePedidoAcondi(r.PedidoAcondi, w)
	}
	return fmt.Errorf("invalid value for *UnionNullPedidoAcondi")
}

func NewUnionNullPedidoAcondi() *UnionNullPedidoAcondi {
	return &UnionNullPedidoAcondi{}
}

func (r *UnionNullPedidoAcondi) Serialize(w io.Writer) error {
	return writeUnionNullPedidoAcondi(r, w)
}

func DeserializeUnionNullPedidoAcondi(r io.Reader) (*UnionNullPedidoAcondi, error) {
	t := NewUnionNullPedidoAcondi()
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

func DeserializeUnionNullPedidoAcondiFromSchema(r io.Reader, schema string) (*UnionNullPedidoAcondi, error) {
	t := NewUnionNullPedidoAcondi()
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

func (r *UnionNullPedidoAcondi) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]},{\"name\":\"lineas\",\"type\":{\"items\":{\"fields\":[{\"default\":null,\"name\":\"almacenCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cantidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaDeVencimiento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoLote\",\"type\":[\"null\",\"string\"]}],\"name\":\"LineaPedido\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"PedidoAcondi\",\"namespace\":\"Andreani.ApiAcondicionamientoV2.Events.Common\",\"type\":\"record\"}]"
}

func (_ *UnionNullPedidoAcondi) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullPedidoAcondi) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullPedidoAcondi) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullPedidoAcondi) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullPedidoAcondi) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullPedidoAcondi) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullPedidoAcondi) SetLong(v int64) {

	r.UnionType = (UnionNullPedidoAcondiTypeEnum)(v)
}

func (r *UnionNullPedidoAcondi) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.PedidoAcondi = NewPedidoAcondi()
		return &types.Record{Target: (&r.PedidoAcondi)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullPedidoAcondi) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullPedidoAcondi) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullPedidoAcondi) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullPedidoAcondi) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullPedidoAcondi) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullPedidoAcondi) Finalize()                        {}

func (r *UnionNullPedidoAcondi) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullPedidoAcondiTypeEnumPedidoAcondi:
		return json.Marshal(map[string]interface{}{"Andreani.ApiAcondicionamientoV2.Events.Common.PedidoAcondi": r.PedidoAcondi})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullPedidoAcondi")
}

func (r *UnionNullPedidoAcondi) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.ApiAcondicionamientoV2.Events.Common.PedidoAcondi"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.PedidoAcondi)
	}
	return fmt.Errorf("invalid value for *UnionNullPedidoAcondi")
}
