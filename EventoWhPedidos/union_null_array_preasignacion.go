// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhPedidosPreasignados.avsc
 */
package EventoWhPedidosEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullArrayPreasignacionTypeEnum int

const (
	UnionNullArrayPreasignacionTypeEnumArrayPreasignacion UnionNullArrayPreasignacionTypeEnum = 1
)

type UnionNullArrayPreasignacion struct {
	Null               *types.NullVal
	ArrayPreasignacion []Preasignacion
	UnionType          UnionNullArrayPreasignacionTypeEnum
}

func writeUnionNullArrayPreasignacion(r *UnionNullArrayPreasignacion, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullArrayPreasignacionTypeEnumArrayPreasignacion:
		return writeArrayPreasignacion(r.ArrayPreasignacion, w)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayPreasignacion")
}

func NewUnionNullArrayPreasignacion() *UnionNullArrayPreasignacion {
	return &UnionNullArrayPreasignacion{}
}

func (r *UnionNullArrayPreasignacion) Serialize(w io.Writer) error {
	return writeUnionNullArrayPreasignacion(r, w)
}

func DeserializeUnionNullArrayPreasignacion(r io.Reader) (*UnionNullArrayPreasignacion, error) {
	t := NewUnionNullArrayPreasignacion()
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

func DeserializeUnionNullArrayPreasignacionFromSchema(r io.Reader, schema string) (*UnionNullArrayPreasignacion, error) {
	t := NewUnionNullArrayPreasignacion()
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

func (r *UnionNullArrayPreasignacion) Schema() string {
	return "[\"null\",{\"items\":{\"fields\":[{\"name\":\"PaqueteLote\",\"type\":\"string\"},{\"name\":\"LoteCajitaFabricante\",\"type\":\"string\"},{\"name\":\"LoteSecundario\",\"type\":\"string\"},{\"name\":\"FechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ProductoTrazable\",\"type\":\"string\"},{\"name\":\"AlmacenConsumo\",\"type\":\"string\"},{\"name\":\"EstadoLote\",\"type\":\"string\"},{\"name\":\"BloqueoUbicacion\",\"type\":\"string\"},{\"name\":\"VitaUtilLote\",\"type\":\"string\"},{\"name\":\"LoteSce\",\"type\":\"string\"},{\"name\":\"CantidadPreasignada\",\"type\":\"float\"},{\"name\":\"FechaPreasignacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"Preasignacion\",\"type\":\"record\"},\"type\":\"array\"}]"
}

func (_ *UnionNullArrayPreasignacion) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullArrayPreasignacion) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullArrayPreasignacion) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullArrayPreasignacion) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullArrayPreasignacion) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullArrayPreasignacion) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullArrayPreasignacion) SetLong(v int64) {

	r.UnionType = (UnionNullArrayPreasignacionTypeEnum)(v)
}

func (r *UnionNullArrayPreasignacion) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ArrayPreasignacion = make([]Preasignacion, 0)
		return &ArrayPreasignacionWrapper{Target: (&r.ArrayPreasignacion)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayPreasignacion) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullArrayPreasignacion) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullArrayPreasignacion) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullArrayPreasignacion) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayPreasignacion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullArrayPreasignacion) Finalize()                {}

func (r *UnionNullArrayPreasignacion) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullArrayPreasignacionTypeEnumArrayPreasignacion:
		return json.Marshal(map[string]interface{}{"array": r.ArrayPreasignacion})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullArrayPreasignacion")
}

func (r *UnionNullArrayPreasignacion) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ArrayPreasignacion)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayPreasignacion")
}
