// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     WorkerData.avsc
 */
package HCMWorkersEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullPropertiesDataTypeEnum int

const (
	UnionNullPropertiesDataTypeEnumPropertiesData UnionNullPropertiesDataTypeEnum = 1
)

type UnionNullPropertiesData struct {
	Null           *types.NullVal
	PropertiesData PropertiesData
	UnionType      UnionNullPropertiesDataTypeEnum
}

func writeUnionNullPropertiesData(r *UnionNullPropertiesData, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullPropertiesDataTypeEnumPropertiesData:
		return writePropertiesData(r.PropertiesData, w)
	}
	return fmt.Errorf("invalid value for *UnionNullPropertiesData")
}

func NewUnionNullPropertiesData() *UnionNullPropertiesData {
	return &UnionNullPropertiesData{}
}

func (r *UnionNullPropertiesData) Serialize(w io.Writer) error {
	return writeUnionNullPropertiesData(r, w)
}

func DeserializeUnionNullPropertiesData(r io.Reader) (*UnionNullPropertiesData, error) {
	t := NewUnionNullPropertiesData()
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

func DeserializeUnionNullPropertiesDataFromSchema(r io.Reader, schema string) (*UnionNullPropertiesData, error) {
	t := NewUnionNullPropertiesData()
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

func (r *UnionNullPropertiesData) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"changeIndicator\",\"type\":\"string\"}],\"name\":\"PropertiesData\",\"type\":\"record\"}]"
}

func (_ *UnionNullPropertiesData) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullPropertiesData) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullPropertiesData) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullPropertiesData) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullPropertiesData) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullPropertiesData) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullPropertiesData) SetLong(v int64) {

	r.UnionType = (UnionNullPropertiesDataTypeEnum)(v)
}

func (r *UnionNullPropertiesData) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.PropertiesData = NewPropertiesData()
		return &types.Record{Target: (&r.PropertiesData)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullPropertiesData) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullPropertiesData) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullPropertiesData) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullPropertiesData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullPropertiesData) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullPropertiesData) Finalize()                        {}

func (r *UnionNullPropertiesData) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullPropertiesDataTypeEnumPropertiesData:
		return json.Marshal(map[string]interface{}{"Andreani.HCMWorkers.Events.Record.PropertiesData": r.PropertiesData})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullPropertiesData")
}

func (r *UnionNullPropertiesData) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.HCMWorkers.Events.Record.PropertiesData"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.PropertiesData)
	}
	return fmt.Errorf("invalid value for *UnionNullPropertiesData")
}
