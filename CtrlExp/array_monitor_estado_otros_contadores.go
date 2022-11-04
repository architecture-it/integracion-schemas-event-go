// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     MonitorEstado.avsc
 */
package CtrlExpEvents

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayMonitorEstadoOtrosContadores(r []MonitorEstadoOtrosContadores, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeMonitorEstadoOtrosContadores(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayMonitorEstadoOtrosContadoresWrapper struct {
	Target *[]MonitorEstadoOtrosContadores
}

func (_ ArrayMonitorEstadoOtrosContadoresWrapper) SetBoolean(v bool)  { panic("Unsupported operation") }
func (_ ArrayMonitorEstadoOtrosContadoresWrapper) SetInt(v int32)     { panic("Unsupported operation") }
func (_ ArrayMonitorEstadoOtrosContadoresWrapper) SetLong(v int64)    { panic("Unsupported operation") }
func (_ ArrayMonitorEstadoOtrosContadoresWrapper) SetFloat(v float32) { panic("Unsupported operation") }
func (_ ArrayMonitorEstadoOtrosContadoresWrapper) SetDouble(v float64) {
	panic("Unsupported operation")
}
func (_ ArrayMonitorEstadoOtrosContadoresWrapper) SetBytes(v []byte)  { panic("Unsupported operation") }
func (_ ArrayMonitorEstadoOtrosContadoresWrapper) SetString(v string) { panic("Unsupported operation") }
func (_ ArrayMonitorEstadoOtrosContadoresWrapper) SetUnionElem(v int64) {
	panic("Unsupported operation")
}
func (_ ArrayMonitorEstadoOtrosContadoresWrapper) Get(i int) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayMonitorEstadoOtrosContadoresWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayMonitorEstadoOtrosContadoresWrapper) Finalize()        {}
func (_ ArrayMonitorEstadoOtrosContadoresWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r ArrayMonitorEstadoOtrosContadoresWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]MonitorEstadoOtrosContadores, 0, s)
	}
}
func (r ArrayMonitorEstadoOtrosContadoresWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayMonitorEstadoOtrosContadoresWrapper) AppendArray() types.Field {
	var v MonitorEstadoOtrosContadores
	v = NewMonitorEstadoOtrosContadores()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}