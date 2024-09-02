// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     StructureData.avsc
 */
package appnameEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullStructureModelTypeEnum int

const (
	UnionNullStructureModelTypeEnumStructureModel UnionNullStructureModelTypeEnum = 1
)

type UnionNullStructureModel struct {
	Null           *types.NullVal
	StructureModel StructureModel
	UnionType      UnionNullStructureModelTypeEnum
}

func writeUnionNullStructureModel(r *UnionNullStructureModel, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullStructureModelTypeEnumStructureModel:
		return writeStructureModel(r.StructureModel, w)
	}
	return fmt.Errorf("invalid value for *UnionNullStructureModel")
}

func NewUnionNullStructureModel() *UnionNullStructureModel {
	return &UnionNullStructureModel{}
}

func (r *UnionNullStructureModel) Serialize(w io.Writer) error {
	return writeUnionNullStructureModel(r, w)
}

func DeserializeUnionNullStructureModel(r io.Reader) (*UnionNullStructureModel, error) {
	t := NewUnionNullStructureModel()
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

func DeserializeUnionNullStructureModelFromSchema(r io.Reader, schema string) (*UnionNullStructureModel, error) {
	t := NewUnionNullStructureModel()
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

func (r *UnionNullStructureModel) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"}],\"name\":\"StructureModel\",\"type\":\"record\"}]"
}

func (_ *UnionNullStructureModel) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullStructureModel) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullStructureModel) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullStructureModel) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullStructureModel) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullStructureModel) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullStructureModel) SetLong(v int64) {

	r.UnionType = (UnionNullStructureModelTypeEnum)(v)
}

func (r *UnionNullStructureModel) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.StructureModel = NewStructureModel()
		return &types.Record{Target: (&r.StructureModel)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullStructureModel) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullStructureModel) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullStructureModel) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullStructureModel) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullStructureModel) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullStructureModel) Finalize()                        {}

func (r *UnionNullStructureModel) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullStructureModelTypeEnumStructureModel:
		return json.Marshal(map[string]interface{}{"Andreani.RHpro.Events.Common.StructureModel": r.StructureModel})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullStructureModel")
}

func (r *UnionNullStructureModel) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.RHpro.Events.Common.StructureModel"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.StructureModel)
	}
	return fmt.Errorf("invalid value for *UnionNullStructureModel")
}
