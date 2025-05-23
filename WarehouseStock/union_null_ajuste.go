// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CambioDeStock.avsc
 */
package WarehouseStockEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullAjusteTypeEnum int

const (
	UnionNullAjusteTypeEnumAjuste UnionNullAjusteTypeEnum = 1
)

type UnionNullAjuste struct {
	Null      *types.NullVal
	Ajuste    Ajuste
	UnionType UnionNullAjusteTypeEnum
}

func writeUnionNullAjuste(r *UnionNullAjuste, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullAjusteTypeEnumAjuste:
		return writeAjuste(r.Ajuste, w)
	}
	return fmt.Errorf("invalid value for *UnionNullAjuste")
}

func NewUnionNullAjuste() *UnionNullAjuste {
	return &UnionNullAjuste{}
}

func (r *UnionNullAjuste) Serialize(w io.Writer) error {
	return writeUnionNullAjuste(r, w)
}

func DeserializeUnionNullAjuste(r io.Reader) (*UnionNullAjuste, error) {
	t := NewUnionNullAjuste()
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

func DeserializeUnionNullAjusteFromSchema(r io.Reader, schema string) (*UnionNullAjuste, error) {
	t := NewUnionNullAjuste()
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

func (r *UnionNullAjuste) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"StockTotal\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"StockDisponible\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"StockEnTransito\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"StockAnteriorAjuste\",\"type\":[\"null\",\"float\"]}],\"name\":\"Ajuste\",\"namespace\":\"Andreani.WarehouseStock.Events.StockCommon\",\"type\":\"record\"}]"
}

func (_ *UnionNullAjuste) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullAjuste) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullAjuste) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullAjuste) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullAjuste) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullAjuste) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullAjuste) SetLong(v int64) {

	r.UnionType = (UnionNullAjusteTypeEnum)(v)
}

func (r *UnionNullAjuste) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Ajuste = NewAjuste()
		return &types.Record{Target: (&r.Ajuste)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullAjuste) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullAjuste) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullAjuste) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullAjuste) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullAjuste) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullAjuste) Finalize()                        {}

func (r *UnionNullAjuste) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullAjusteTypeEnumAjuste:
		return json.Marshal(map[string]interface{}{"Andreani.WarehouseStock.Events.StockCommon.Ajuste": r.Ajuste})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullAjuste")
}

func (r *UnionNullAjuste) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.WarehouseStock.Events.StockCommon.Ajuste"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Ajuste)
	}
	return fmt.Errorf("invalid value for *UnionNullAjuste")
}
