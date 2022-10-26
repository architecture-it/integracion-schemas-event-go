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

type UnionArrayMonitorEstadoOtrosSemaforosTypeEnum int

const (
	UnionArrayMonitorEstadoOtrosSemaforosTypeEnumArrayMonitorEstadoOtrosSemaforos UnionArrayMonitorEstadoOtrosSemaforosTypeEnum = 0
)

type UnionArrayMonitorEstadoOtrosSemaforos struct {
	ArrayMonitorEstadoOtrosSemaforos []MonitorEstadoOtrosSemaforos
	UnionType                        UnionArrayMonitorEstadoOtrosSemaforosTypeEnum
}

func writeUnionArrayMonitorEstadoOtrosSemaforos(r UnionArrayMonitorEstadoOtrosSemaforos, w io.Writer) error {

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionArrayMonitorEstadoOtrosSemaforosTypeEnumArrayMonitorEstadoOtrosSemaforos:
		return writeArrayMonitorEstadoOtrosSemaforos(r.ArrayMonitorEstadoOtrosSemaforos, w)
	}
	return fmt.Errorf("invalid value for UnionArrayMonitorEstadoOtrosSemaforos")
}

func NewUnionArrayMonitorEstadoOtrosSemaforos() UnionArrayMonitorEstadoOtrosSemaforos {
	return UnionArrayMonitorEstadoOtrosSemaforos{}
}

func (r UnionArrayMonitorEstadoOtrosSemaforos) Serialize(w io.Writer) error {
	return writeUnionArrayMonitorEstadoOtrosSemaforos(r, w)
}

func DeserializeUnionArrayMonitorEstadoOtrosSemaforos(r io.Reader) (UnionArrayMonitorEstadoOtrosSemaforos, error) {
	t := NewUnionArrayMonitorEstadoOtrosSemaforos()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionArrayMonitorEstadoOtrosSemaforosFromSchema(r io.Reader, schema string) (UnionArrayMonitorEstadoOtrosSemaforos, error) {
	t := NewUnionArrayMonitorEstadoOtrosSemaforos()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r UnionArrayMonitorEstadoOtrosSemaforos) Schema() string {
	return "[{\"items\":{\"fields\":[{\"name\":\"Semaforo\",\"type\":\"string\"},{\"name\":\"Valor\",\"type\":\"boolean\"}],\"name\":\"MonitorEstadoOtrosSemaforos\",\"type\":\"record\"},\"type\":\"array\"}]"
}

func (_ UnionArrayMonitorEstadoOtrosSemaforos) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ UnionArrayMonitorEstadoOtrosSemaforos) SetInt(v int32)      { panic("Unsupported operation") }
func (_ UnionArrayMonitorEstadoOtrosSemaforos) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ UnionArrayMonitorEstadoOtrosSemaforos) SetDouble(v float64) { panic("Unsupported operation") }
func (_ UnionArrayMonitorEstadoOtrosSemaforos) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ UnionArrayMonitorEstadoOtrosSemaforos) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionArrayMonitorEstadoOtrosSemaforos) SetLong(v int64) {

	r.UnionType = (UnionArrayMonitorEstadoOtrosSemaforosTypeEnum)(v)
}

func (r *UnionArrayMonitorEstadoOtrosSemaforos) Get(i int) types.Field {

	switch i {
	case 0:
		r.ArrayMonitorEstadoOtrosSemaforos = make([]MonitorEstadoOtrosSemaforos, 0)
		return &ArrayMonitorEstadoOtrosSemaforosWrapper{Target: (&r.ArrayMonitorEstadoOtrosSemaforos)}
	}
	panic("Unknown field index")
}
func (_ UnionArrayMonitorEstadoOtrosSemaforos) NullField(i int)  { panic("Unsupported operation") }
func (_ UnionArrayMonitorEstadoOtrosSemaforos) HintSize(i int)   { panic("Unsupported operation") }
func (_ UnionArrayMonitorEstadoOtrosSemaforos) SetDefault(i int) { panic("Unsupported operation") }
func (_ UnionArrayMonitorEstadoOtrosSemaforos) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ UnionArrayMonitorEstadoOtrosSemaforos) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ UnionArrayMonitorEstadoOtrosSemaforos) Finalize() {}

func (r UnionArrayMonitorEstadoOtrosSemaforos) MarshalJSON() ([]byte, error) {

	switch r.UnionType {
	case UnionArrayMonitorEstadoOtrosSemaforosTypeEnumArrayMonitorEstadoOtrosSemaforos:
		return json.Marshal(map[string]interface{}{"array": r.ArrayMonitorEstadoOtrosSemaforos})
	}
	return nil, fmt.Errorf("invalid value for UnionArrayMonitorEstadoOtrosSemaforos")
}

func (r *UnionArrayMonitorEstadoOtrosSemaforos) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 0
		return json.Unmarshal([]byte(value), &r.ArrayMonitorEstadoOtrosSemaforos)
	}
	return fmt.Errorf("invalid value for UnionArrayMonitorEstadoOtrosSemaforos")
}
