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

type UnionNullNovedadAssetTypeEnum int

const (
	UnionNullNovedadAssetTypeEnumNovedadAsset UnionNullNovedadAssetTypeEnum = 1
)

type UnionNullNovedadAsset struct {
	Null         *types.NullVal
	NovedadAsset NovedadAsset
	UnionType    UnionNullNovedadAssetTypeEnum
}

func writeUnionNullNovedadAsset(r *UnionNullNovedadAsset, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullNovedadAssetTypeEnumNovedadAsset:
		return writeNovedadAsset(r.NovedadAsset, w)
	}
	return fmt.Errorf("invalid value for *UnionNullNovedadAsset")
}

func NewUnionNullNovedadAsset() *UnionNullNovedadAsset {
	return &UnionNullNovedadAsset{}
}

func (r *UnionNullNovedadAsset) Serialize(w io.Writer) error {
	return writeUnionNullNovedadAsset(r, w)
}

func DeserializeUnionNullNovedadAsset(r io.Reader) (*UnionNullNovedadAsset, error) {
	t := NewUnionNullNovedadAsset()
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

func DeserializeUnionNullNovedadAssetFromSchema(r io.Reader, schema string) (*UnionNullNovedadAsset, error) {
	t := NewUnionNullNovedadAsset()
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

func (r *UnionNullNovedadAsset) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"nueva_planta\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ceco\",\"type\":[\"null\",\"string\"]}],\"name\":\"NovedadAsset\",\"namespace\":\"Andreani.EAM.Events.Sharepoint\",\"type\":\"record\"}]"
}

func (_ *UnionNullNovedadAsset) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullNovedadAsset) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullNovedadAsset) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullNovedadAsset) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullNovedadAsset) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullNovedadAsset) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullNovedadAsset) SetLong(v int64) {

	r.UnionType = (UnionNullNovedadAssetTypeEnum)(v)
}

func (r *UnionNullNovedadAsset) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.NovedadAsset = NewNovedadAsset()
		return &types.Record{Target: (&r.NovedadAsset)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullNovedadAsset) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullNovedadAsset) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullNovedadAsset) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullNovedadAsset) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullNovedadAsset) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullNovedadAsset) Finalize()                        {}

func (r *UnionNullNovedadAsset) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullNovedadAssetTypeEnumNovedadAsset:
		return json.Marshal(map[string]interface{}{"Andreani.EAM.Events.Sharepoint.NovedadAsset": r.NovedadAsset})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullNovedadAsset")
}

func (r *UnionNullNovedadAsset) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.EAM.Events.Sharepoint.NovedadAsset"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.NovedadAsset)
	}
	return fmt.Errorf("invalid value for *UnionNullNovedadAsset")
}
