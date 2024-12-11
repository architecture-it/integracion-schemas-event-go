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

type UnionNullWorkOrderTypeEnum int

const (
	UnionNullWorkOrderTypeEnumWorkOrder UnionNullWorkOrderTypeEnum = 1
)

type UnionNullWorkOrder struct {
	Null      *types.NullVal
	WorkOrder WorkOrder
	UnionType UnionNullWorkOrderTypeEnum
}

func writeUnionNullWorkOrder(r *UnionNullWorkOrder, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullWorkOrderTypeEnumWorkOrder:
		return writeWorkOrder(r.WorkOrder, w)
	}
	return fmt.Errorf("invalid value for *UnionNullWorkOrder")
}

func NewUnionNullWorkOrder() *UnionNullWorkOrder {
	return &UnionNullWorkOrder{}
}

func (r *UnionNullWorkOrder) Serialize(w io.Writer) error {
	return writeUnionNullWorkOrder(r, w)
}

func DeserializeUnionNullWorkOrder(r io.Reader) (*UnionNullWorkOrder, error) {
	t := NewUnionNullWorkOrder()
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

func DeserializeUnionNullWorkOrderFromSchema(r io.Reader, schema string) (*UnionNullWorkOrder, error) {
	t := NewUnionNullWorkOrder()
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

func (r *UnionNullWorkOrder) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"id_equipo\",\"type\":\"string\"},{\"name\":\"planta\",\"type\":\"string\"},{\"name\":\"descripcion\",\"type\":\"string\"},{\"name\":\"user_report\",\"type\":\"string\"}],\"name\":\"WorkOrder\",\"namespace\":\"Andreani.EAM.Events.Sharepoint\",\"type\":\"record\"}]"
}

func (_ *UnionNullWorkOrder) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullWorkOrder) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullWorkOrder) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullWorkOrder) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullWorkOrder) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullWorkOrder) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullWorkOrder) SetLong(v int64) {

	r.UnionType = (UnionNullWorkOrderTypeEnum)(v)
}

func (r *UnionNullWorkOrder) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.WorkOrder = NewWorkOrder()
		return &types.Record{Target: (&r.WorkOrder)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullWorkOrder) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullWorkOrder) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullWorkOrder) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullWorkOrder) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullWorkOrder) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullWorkOrder) Finalize()                        {}

func (r *UnionNullWorkOrder) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullWorkOrderTypeEnumWorkOrder:
		return json.Marshal(map[string]interface{}{"Andreani.EAM.Events.Sharepoint.WorkOrder": r.WorkOrder})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullWorkOrder")
}

func (r *UnionNullWorkOrder) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.EAM.Events.Sharepoint.WorkOrder"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.WorkOrder)
	}
	return fmt.Errorf("invalid value for *UnionNullWorkOrder")
}
