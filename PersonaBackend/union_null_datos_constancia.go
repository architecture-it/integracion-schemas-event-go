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

type UnionNullDatosConstanciaTypeEnum int

const (
	UnionNullDatosConstanciaTypeEnumDatosConstancia UnionNullDatosConstanciaTypeEnum = 1
)

type UnionNullDatosConstancia struct {
	Null            *types.NullVal
	DatosConstancia DatosConstancia
	UnionType       UnionNullDatosConstanciaTypeEnum
}

func writeUnionNullDatosConstancia(r *UnionNullDatosConstancia, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullDatosConstanciaTypeEnumDatosConstancia:
		return writeDatosConstancia(r.DatosConstancia, w)
	}
	return fmt.Errorf("invalid value for *UnionNullDatosConstancia")
}

func NewUnionNullDatosConstancia() *UnionNullDatosConstancia {
	return &UnionNullDatosConstancia{}
}

func (r *UnionNullDatosConstancia) Serialize(w io.Writer) error {
	return writeUnionNullDatosConstancia(r, w)
}

func DeserializeUnionNullDatosConstancia(r io.Reader) (*UnionNullDatosConstancia, error) {
	t := NewUnionNullDatosConstancia()
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

func DeserializeUnionNullDatosConstanciaFromSchema(r io.Reader, schema string) (*UnionNullDatosConstancia, error) {
	t := NewUnionNullDatosConstancia()
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

func (r *UnionNullDatosConstancia) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"Url\",\"type\":\"string\"},{\"name\":\"NumeroPermisionaria\",\"type\":\"string\"},{\"name\":\"SucursalDistribucion\",\"type\":\"string\"},{\"name\":\"SucursalDistribucionNomenclatura\",\"type\":\"string\"},{\"name\":\"SucursalDistribucionDescripcion\",\"type\":\"string\"},{\"name\":\"SucursalDistribucionId\",\"type\":\"string\"},{\"name\":\"SucursalRendicion\",\"type\":\"string\"},{\"name\":\"SucursalRendicionNomenclatura\",\"type\":\"string\"},{\"name\":\"SucursalRendicionDescripcion\",\"type\":\"string\"},{\"name\":\"SucursalRendicionId\",\"type\":\"string\"},{\"name\":\"CodigoSucursalCabecera\",\"type\":\"string\"},{\"name\":\"SucursalAbastecedoraDescripcion\",\"type\":\"string\"},{\"name\":\"SucursalAbastecedoraId\",\"type\":\"string\"},{\"name\":\"CodigoZonaReparto\",\"type\":\"string\"},{\"default\":null,\"name\":\"Region\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Provincia\",\"type\":[\"null\",\"string\"]}],\"name\":\"DatosConstancia\",\"type\":\"record\"}]"
}

func (_ *UnionNullDatosConstancia) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullDatosConstancia) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullDatosConstancia) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullDatosConstancia) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullDatosConstancia) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullDatosConstancia) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullDatosConstancia) SetLong(v int64) {

	r.UnionType = (UnionNullDatosConstanciaTypeEnum)(v)
}

func (r *UnionNullDatosConstancia) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.DatosConstancia = NewDatosConstancia()
		return &types.Record{Target: (&r.DatosConstancia)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullDatosConstancia) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullDatosConstancia) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullDatosConstancia) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullDatosConstancia) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullDatosConstancia) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullDatosConstancia) Finalize()                        {}

func (r *UnionNullDatosConstancia) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullDatosConstanciaTypeEnumDatosConstancia:
		return json.Marshal(map[string]interface{}{"Andreani.PersonaBackend.Events.Common.DatosConstancia": r.DatosConstancia})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullDatosConstancia")
}

func (r *UnionNullDatosConstancia) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.PersonaBackend.Events.Common.DatosConstancia"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.DatosConstancia)
	}
	return fmt.Errorf("invalid value for *UnionNullDatosConstancia")
}
