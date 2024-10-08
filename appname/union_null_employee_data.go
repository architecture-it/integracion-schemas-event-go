// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Employee.avsc
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

type UnionNullEmployeeDataTypeEnum int

const (
	UnionNullEmployeeDataTypeEnumEmployeeData UnionNullEmployeeDataTypeEnum = 1
)

type UnionNullEmployeeData struct {
	Null         *types.NullVal
	EmployeeData EmployeeData
	UnionType    UnionNullEmployeeDataTypeEnum
}

func writeUnionNullEmployeeData(r *UnionNullEmployeeData, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullEmployeeDataTypeEnumEmployeeData:
		return writeEmployeeData(r.EmployeeData, w)
	}
	return fmt.Errorf("invalid value for *UnionNullEmployeeData")
}

func NewUnionNullEmployeeData() *UnionNullEmployeeData {
	return &UnionNullEmployeeData{}
}

func (r *UnionNullEmployeeData) Serialize(w io.Writer) error {
	return writeUnionNullEmployeeData(r, w)
}

func DeserializeUnionNullEmployeeData(r io.Reader) (*UnionNullEmployeeData, error) {
	t := NewUnionNullEmployeeData()
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

func DeserializeUnionNullEmployeeDataFromSchema(r io.Reader, schema string) (*UnionNullEmployeeData, error) {
	t := NewUnionNullEmployeeData()
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

func (r *UnionNullEmployeeData) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"DateOfBirth\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"HiringDate\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"ExternalId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FamilyName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FirstName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ImageUrl\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LastName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"MiddleName\",\"type\":[\"null\",\"string\"]},{\"name\":\"IsActive\",\"type\":\"boolean\"},{\"name\":\"Id\",\"type\":\"int\"}],\"name\":\"EmployeeData\",\"namespace\":\"Andreani.RHpro.Events.Common\",\"type\":\"record\"}]"
}

func (_ *UnionNullEmployeeData) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullEmployeeData) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullEmployeeData) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullEmployeeData) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullEmployeeData) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullEmployeeData) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullEmployeeData) SetLong(v int64) {

	r.UnionType = (UnionNullEmployeeDataTypeEnum)(v)
}

func (r *UnionNullEmployeeData) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.EmployeeData = NewEmployeeData()
		return &types.Record{Target: (&r.EmployeeData)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullEmployeeData) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullEmployeeData) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullEmployeeData) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullEmployeeData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullEmployeeData) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullEmployeeData) Finalize()                        {}

func (r *UnionNullEmployeeData) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullEmployeeDataTypeEnumEmployeeData:
		return json.Marshal(map[string]interface{}{"Andreani.RHpro.Events.Common.EmployeeData": r.EmployeeData})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullEmployeeData")
}

func (r *UnionNullEmployeeData) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.RHpro.Events.Common.EmployeeData"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.EmployeeData)
	}
	return fmt.Errorf("invalid value for *UnionNullEmployeeData")
}
