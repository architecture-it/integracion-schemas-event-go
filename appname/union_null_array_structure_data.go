// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Structure.avsc
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

type UnionNullArrayStructureDataTypeEnum int

const (
	UnionNullArrayStructureDataTypeEnumArrayStructureData UnionNullArrayStructureDataTypeEnum = 1
)

type UnionNullArrayStructureData struct {
	Null               *types.NullVal
	ArrayStructureData []StructureData
	UnionType          UnionNullArrayStructureDataTypeEnum
}

func writeUnionNullArrayStructureData(r *UnionNullArrayStructureData, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullArrayStructureDataTypeEnumArrayStructureData:
		return writeArrayStructureData(r.ArrayStructureData, w)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayStructureData")
}

func NewUnionNullArrayStructureData() *UnionNullArrayStructureData {
	return &UnionNullArrayStructureData{}
}

func (r *UnionNullArrayStructureData) Serialize(w io.Writer) error {
	return writeUnionNullArrayStructureData(r, w)
}

func DeserializeUnionNullArrayStructureData(r io.Reader) (*UnionNullArrayStructureData, error) {
	t := NewUnionNullArrayStructureData()
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

func DeserializeUnionNullArrayStructureDataFromSchema(r io.Reader, schema string) (*UnionNullArrayStructureData, error) {
	t := NewUnionNullArrayStructureData()
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

func (r *UnionNullArrayStructureData) Schema() string {
	return "[\"null\",{\"items\":{\"fields\":[{\"name\":\"StructureId\",\"type\":\"int\"},{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"DateFrom\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"DateTo\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"Description\",\"type\":\"string\"},{\"default\":null,\"name\":\"ExternalId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Type\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"}],\"name\":\"StructureType\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Reason\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"},{\"name\":\"Reason\",\"type\":\"string\"}],\"name\":\"StructureReason\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Model\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"}],\"name\":\"StructureModel\",\"type\":\"record\"}]}],\"name\":\"StructureData\",\"namespace\":\"Andreani.RHpro.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}]"
}

func (_ *UnionNullArrayStructureData) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullArrayStructureData) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullArrayStructureData) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullArrayStructureData) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullArrayStructureData) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullArrayStructureData) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullArrayStructureData) SetLong(v int64) {

	r.UnionType = (UnionNullArrayStructureDataTypeEnum)(v)
}

func (r *UnionNullArrayStructureData) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ArrayStructureData = make([]StructureData, 0)
		return &ArrayStructureDataWrapper{Target: (&r.ArrayStructureData)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayStructureData) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullArrayStructureData) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullArrayStructureData) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullArrayStructureData) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayStructureData) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullArrayStructureData) Finalize()                {}

func (r *UnionNullArrayStructureData) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullArrayStructureDataTypeEnumArrayStructureData:
		return json.Marshal(map[string]interface{}{"array": r.ArrayStructureData})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullArrayStructureData")
}

func (r *UnionNullArrayStructureData) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ArrayStructureData)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayStructureData")
}
