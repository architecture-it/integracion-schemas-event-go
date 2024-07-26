// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     AcondicionamientoV2Procesado.avsc
 */
package ApiAcondicionamientoV2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullAbastecimientoAcondiTypeEnum int

const (
	UnionNullAbastecimientoAcondiTypeEnumAbastecimientoAcondi UnionNullAbastecimientoAcondiTypeEnum = 1
)

type UnionNullAbastecimientoAcondi struct {
	Null                 *types.NullVal
	AbastecimientoAcondi AbastecimientoAcondi
	UnionType            UnionNullAbastecimientoAcondiTypeEnum
}

func writeUnionNullAbastecimientoAcondi(r *UnionNullAbastecimientoAcondi, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullAbastecimientoAcondiTypeEnumAbastecimientoAcondi:
		return writeAbastecimientoAcondi(r.AbastecimientoAcondi, w)
	}
	return fmt.Errorf("invalid value for *UnionNullAbastecimientoAcondi")
}

func NewUnionNullAbastecimientoAcondi() *UnionNullAbastecimientoAcondi {
	return &UnionNullAbastecimientoAcondi{}
}

func (r *UnionNullAbastecimientoAcondi) Serialize(w io.Writer) error {
	return writeUnionNullAbastecimientoAcondi(r, w)
}

func DeserializeUnionNullAbastecimientoAcondi(r io.Reader) (*UnionNullAbastecimientoAcondi, error) {
	t := NewUnionNullAbastecimientoAcondi()
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

func DeserializeUnionNullAbastecimientoAcondiFromSchema(r io.Reader, schema string) (*UnionNullAbastecimientoAcondi, error) {
	t := NewUnionNullAbastecimientoAcondi()
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

func (r *UnionNullAbastecimientoAcondi) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]},{\"name\":\"lineas\",\"type\":{\"items\":{\"fields\":[{\"default\":null,\"name\":\"codigo\",\"type\":[\"null\",\"string\"]},{\"name\":\"cantidadPedida\",\"type\":\"float\"},{\"default\":null,\"name\":\"almacenWMS\",\"type\":[\"null\",\"string\"]},{\"name\":\"tipoAcondicionamiento\",\"type\":{\"items\":\"string\",\"type\":\"array\"}},{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaDeVencimiento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoLote\",\"type\":[\"null\",\"string\"]}],\"name\":\"LineaAbastecimiento\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"AbastecimientoAcondi\",\"namespace\":\"Andreani.ApiAcondicionamientoV2.Events.Common\",\"type\":\"record\"}]"
}

func (_ *UnionNullAbastecimientoAcondi) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullAbastecimientoAcondi) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullAbastecimientoAcondi) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullAbastecimientoAcondi) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullAbastecimientoAcondi) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullAbastecimientoAcondi) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullAbastecimientoAcondi) SetLong(v int64) {

	r.UnionType = (UnionNullAbastecimientoAcondiTypeEnum)(v)
}

func (r *UnionNullAbastecimientoAcondi) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.AbastecimientoAcondi = NewAbastecimientoAcondi()
		return &types.Record{Target: (&r.AbastecimientoAcondi)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullAbastecimientoAcondi) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullAbastecimientoAcondi) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullAbastecimientoAcondi) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullAbastecimientoAcondi) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullAbastecimientoAcondi) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullAbastecimientoAcondi) Finalize()                {}

func (r *UnionNullAbastecimientoAcondi) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullAbastecimientoAcondiTypeEnumAbastecimientoAcondi:
		return json.Marshal(map[string]interface{}{"Andreani.ApiAcondicionamientoV2.Events.Common.AbastecimientoAcondi": r.AbastecimientoAcondi})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullAbastecimientoAcondi")
}

func (r *UnionNullAbastecimientoAcondi) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.ApiAcondicionamientoV2.Events.Common.AbastecimientoAcondi"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.AbastecimientoAcondi)
	}
	return fmt.Errorf("invalid value for *UnionNullAbastecimientoAcondi")
}