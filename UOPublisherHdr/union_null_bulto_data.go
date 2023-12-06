// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     UOHdrCreada.avsc
 */
package UOPublisherHdrEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullBultoDataTypeEnum int

const (
	UnionNullBultoDataTypeEnumBultoData UnionNullBultoDataTypeEnum = 1
)

type UnionNullBultoData struct {
	Null      *types.NullVal
	BultoData BultoData
	UnionType UnionNullBultoDataTypeEnum
}

func writeUnionNullBultoData(r *UnionNullBultoData, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullBultoDataTypeEnumBultoData:
		return writeBultoData(r.BultoData, w)
	}
	return fmt.Errorf("invalid value for *UnionNullBultoData")
}

func NewUnionNullBultoData() *UnionNullBultoData {
	return &UnionNullBultoData{}
}

func (r *UnionNullBultoData) Serialize(w io.Writer) error {
	return writeUnionNullBultoData(r, w)
}

func DeserializeUnionNullBultoData(r io.Reader) (*UnionNullBultoData, error) {
	t := NewUnionNullBultoData()
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

func DeserializeUnionNullBultoDataFromSchema(r io.Reader, schema string) (*UnionNullBultoData, error) {
	t := NewUnionNullBultoData()
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

func (r *UnionNullBultoData) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"Numero\",\"type\":\"int\"},{\"name\":\"Codigos\",\"type\":{\"items\":\"string\",\"type\":\"array\"}}],\"name\":\"BultoData\",\"type\":\"record\"}]"
}

func (_ *UnionNullBultoData) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullBultoData) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullBultoData) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullBultoData) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullBultoData) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullBultoData) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullBultoData) SetLong(v int64) {

	r.UnionType = (UnionNullBultoDataTypeEnum)(v)
}

func (r *UnionNullBultoData) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.BultoData = NewBultoData()
		return &types.Record{Target: (&r.BultoData)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullBultoData) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullBultoData) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullBultoData) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullBultoData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullBultoData) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullBultoData) Finalize()                        {}

func (r *UnionNullBultoData) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullBultoDataTypeEnumBultoData:
		return json.Marshal(map[string]interface{}{"Andreani.UOPublisherHdr.Events.Common.BultoData": r.BultoData})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullBultoData")
}

func (r *UnionNullBultoData) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.UOPublisherHdr.Events.Common.BultoData"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.BultoData)
	}
	return fmt.Errorf("invalid value for *UnionNullBultoData")
}