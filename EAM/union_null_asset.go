// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NovedadEAMPowerApp.avsc
 */
package EAMEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullAssetTypeEnum int

const (
	UnionNullAssetTypeEnumAsset UnionNullAssetTypeEnum = 1
)

type UnionNullAsset struct {
	Null      *types.NullVal
	Asset     Asset
	UnionType UnionNullAssetTypeEnum
}

func writeUnionNullAsset(r *UnionNullAsset, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullAssetTypeEnumAsset:
		return writeAsset(r.Asset, w)
	}
	return fmt.Errorf("invalid value for *UnionNullAsset")
}

func NewUnionNullAsset() *UnionNullAsset {
	return &UnionNullAsset{}
}

func (r *UnionNullAsset) Serialize(w io.Writer) error {
	return writeUnionNullAsset(r, w)
}

func DeserializeUnionNullAsset(r io.Reader) (*UnionNullAsset, error) {
	t := NewUnionNullAsset()
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

func DeserializeUnionNullAssetFromSchema(r io.Reader, schema string) (*UnionNullAsset, error) {
	t := NewUnionNullAsset()
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

func (r *UnionNullAsset) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"id\",\"type\":\"int\"},{\"name\":\"tipo_objeto\",\"type\":\"string\"},{\"name\":\"descripcion\",\"type\":\"string\"},{\"name\":\"clase\",\"type\":\"string\"},{\"name\":\"codigo_costo\",\"type\":\"string\"},{\"name\":\"estado\",\"type\":\"string\"},{\"name\":\"fecha_alta\",\"type\":\"string\"},{\"name\":\"organizacion\",\"type\":\"string\"},{\"name\":\"fabricante\",\"type\":\"string\"},{\"name\":\"modelo\",\"type\":\"string\"},{\"name\":\"nro_serie\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"fueraDeServicio\",\"type\":\"boolean\"},{\"name\":\"cod_eam\",\"type\":\"string\"}],\"name\":\"Asset\",\"namespace\":\"Andreani.EAM.Events.Sharepoint\",\"type\":\"record\"}]"
}

func (_ *UnionNullAsset) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullAsset) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullAsset) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullAsset) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullAsset) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullAsset) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullAsset) SetLong(v int64) {

	r.UnionType = (UnionNullAssetTypeEnum)(v)
}

func (r *UnionNullAsset) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Asset = NewAsset()
		return &types.Record{Target: (&r.Asset)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullAsset) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullAsset) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullAsset) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullAsset) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullAsset) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullAsset) Finalize()                        {}

func (r *UnionNullAsset) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullAssetTypeEnumAsset:
		return json.Marshal(map[string]interface{}{"Andreani.EAM.Events.Sharepoint.Asset": r.Asset})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullAsset")
}

func (r *UnionNullAsset) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.EAM.Events.Sharepoint.Asset"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Asset)
	}
	return fmt.Errorf("invalid value for *UnionNullAsset")
}
