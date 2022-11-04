// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     MonitorEstado.avsc
 */
package CtrlExpEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullArrayMonitorEstadoOtrosSemaforosTypeEnum int

const (
	UnionNullArrayMonitorEstadoOtrosSemaforosTypeEnumArrayMonitorEstadoOtrosSemaforos UnionNullArrayMonitorEstadoOtrosSemaforosTypeEnum = 1
)

type UnionNullArrayMonitorEstadoOtrosSemaforos struct {
	Null                             *types.NullVal
	ArrayMonitorEstadoOtrosSemaforos []MonitorEstadoOtrosSemaforos
	UnionType                        UnionNullArrayMonitorEstadoOtrosSemaforosTypeEnum
}

func writeUnionNullArrayMonitorEstadoOtrosSemaforos(r *UnionNullArrayMonitorEstadoOtrosSemaforos, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullArrayMonitorEstadoOtrosSemaforosTypeEnumArrayMonitorEstadoOtrosSemaforos:
		return writeArrayMonitorEstadoOtrosSemaforos(r.ArrayMonitorEstadoOtrosSemaforos, w)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayMonitorEstadoOtrosSemaforos")
}

func NewUnionNullArrayMonitorEstadoOtrosSemaforos() *UnionNullArrayMonitorEstadoOtrosSemaforos {
	return &UnionNullArrayMonitorEstadoOtrosSemaforos{}
}

func (r *UnionNullArrayMonitorEstadoOtrosSemaforos) Serialize(w io.Writer) error {
	return writeUnionNullArrayMonitorEstadoOtrosSemaforos(r, w)
}

func DeserializeUnionNullArrayMonitorEstadoOtrosSemaforos(r io.Reader) (*UnionNullArrayMonitorEstadoOtrosSemaforos, error) {
	t := NewUnionNullArrayMonitorEstadoOtrosSemaforos()
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

func DeserializeUnionNullArrayMonitorEstadoOtrosSemaforosFromSchema(r io.Reader, schema string) (*UnionNullArrayMonitorEstadoOtrosSemaforos, error) {
	t := NewUnionNullArrayMonitorEstadoOtrosSemaforos()
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

func (r *UnionNullArrayMonitorEstadoOtrosSemaforos) Schema() string {
	return "[\"null\",{\"items\":{\"fields\":[{\"name\":\"Semaforo\",\"type\":\"string\"},{\"name\":\"Valor\",\"type\":\"boolean\"}],\"name\":\"MonitorEstadoOtrosSemaforos\",\"type\":\"record\"},\"type\":\"array\"}]"
}

func (_ *UnionNullArrayMonitorEstadoOtrosSemaforos) SetBoolean(v bool) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayMonitorEstadoOtrosSemaforos) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionNullArrayMonitorEstadoOtrosSemaforos) SetFloat(v float32) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayMonitorEstadoOtrosSemaforos) SetDouble(v float64) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayMonitorEstadoOtrosSemaforos) SetBytes(v []byte) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayMonitorEstadoOtrosSemaforos) SetString(v string) {
	panic("Unsupported operation")
}

func (r *UnionNullArrayMonitorEstadoOtrosSemaforos) SetLong(v int64) {

	r.UnionType = (UnionNullArrayMonitorEstadoOtrosSemaforosTypeEnum)(v)
}

func (r *UnionNullArrayMonitorEstadoOtrosSemaforos) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ArrayMonitorEstadoOtrosSemaforos = make([]MonitorEstadoOtrosSemaforos, 0)
		return &ArrayMonitorEstadoOtrosSemaforosWrapper{Target: (&r.ArrayMonitorEstadoOtrosSemaforos)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayMonitorEstadoOtrosSemaforos) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullArrayMonitorEstadoOtrosSemaforos) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullArrayMonitorEstadoOtrosSemaforos) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullArrayMonitorEstadoOtrosSemaforos) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayMonitorEstadoOtrosSemaforos) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayMonitorEstadoOtrosSemaforos) Finalize() {}

func (r *UnionNullArrayMonitorEstadoOtrosSemaforos) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullArrayMonitorEstadoOtrosSemaforosTypeEnumArrayMonitorEstadoOtrosSemaforos:
		return json.Marshal(map[string]interface{}{"array": r.ArrayMonitorEstadoOtrosSemaforos})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullArrayMonitorEstadoOtrosSemaforos")
}

func (r *UnionNullArrayMonitorEstadoOtrosSemaforos) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ArrayMonitorEstadoOtrosSemaforos)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayMonitorEstadoOtrosSemaforos")
}