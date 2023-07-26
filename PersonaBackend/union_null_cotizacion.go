// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnviosEvent.avsc
 */
package PersonaBackendEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullCotizacionTypeEnum int

const (
	UnionNullCotizacionTypeEnumCotizacion UnionNullCotizacionTypeEnum = 1
)

type UnionNullCotizacion struct {
	Null       *types.NullVal
	Cotizacion Cotizacion
	UnionType  UnionNullCotizacionTypeEnum
}

func writeUnionNullCotizacion(r *UnionNullCotizacion, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullCotizacionTypeEnumCotizacion:
		return writeCotizacion(r.Cotizacion, w)
	}
	return fmt.Errorf("invalid value for *UnionNullCotizacion")
}

func NewUnionNullCotizacion() *UnionNullCotizacion {
	return &UnionNullCotizacion{}
}

func (r *UnionNullCotizacion) Serialize(w io.Writer) error {
	return writeUnionNullCotizacion(r, w)
}

func DeserializeUnionNullCotizacion(r io.Reader) (*UnionNullCotizacion, error) {
	t := NewUnionNullCotizacion()
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

func DeserializeUnionNullCotizacionFromSchema(r io.Reader, schema string) (*UnionNullCotizacion, error) {
	t := NewUnionNullCotizacion()
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

func (r *UnionNullCotizacion) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"IdMercadoPago\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Iva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PesoAforado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SeguroDistribucionSinIva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SeguroDistribucionConIva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DistribucionSinIva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DistribucionConIva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TarifaSinIva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TarifaConIva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Descuento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Cupon\",\"type\":[\"null\",\"string\"]}],\"name\":\"Cotizacion\",\"type\":\"record\"}]"
}

func (_ *UnionNullCotizacion) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullCotizacion) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullCotizacion) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullCotizacion) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullCotizacion) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullCotizacion) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullCotizacion) SetLong(v int64) {

	r.UnionType = (UnionNullCotizacionTypeEnum)(v)
}

func (r *UnionNullCotizacion) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Cotizacion = NewCotizacion()
		return &types.Record{Target: (&r.Cotizacion)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullCotizacion) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullCotizacion) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullCotizacion) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullCotizacion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullCotizacion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullCotizacion) Finalize()                        {}

func (r *UnionNullCotizacion) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullCotizacionTypeEnumCotizacion:
		return json.Marshal(map[string]interface{}{"Andreani.PersonaBackend.Events.Common.Cotizacion": r.Cotizacion})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullCotizacion")
}

func (r *UnionNullCotizacion) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.PersonaBackend.Events.Common.Cotizacion"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Cotizacion)
	}
	return fmt.Errorf("invalid value for *UnionNullCotizacion")
}
