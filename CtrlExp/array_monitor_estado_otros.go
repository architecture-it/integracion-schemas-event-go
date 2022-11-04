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

func writeArrayMonitorEstadoOtros(r []MonitorEstadoOtros, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeMonitorEstadoOtros(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayMonitorEstadoOtrosWrapper struct {
	Target *[]MonitorEstadoOtros
}

func (_ ArrayMonitorEstadoOtrosWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ ArrayMonitorEstadoOtrosWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ ArrayMonitorEstadoOtrosWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ ArrayMonitorEstadoOtrosWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ ArrayMonitorEstadoOtrosWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ ArrayMonitorEstadoOtrosWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ ArrayMonitorEstadoOtrosWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ ArrayMonitorEstadoOtrosWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ ArrayMonitorEstadoOtrosWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ ArrayMonitorEstadoOtrosWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayMonitorEstadoOtrosWrapper) Finalize()        {}
func (_ ArrayMonitorEstadoOtrosWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r ArrayMonitorEstadoOtrosWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]MonitorEstadoOtros, 0, s)
	}
}
func (r ArrayMonitorEstadoOtrosWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayMonitorEstadoOtrosWrapper) AppendArray() types.Field {
	var v MonitorEstadoOtros
	v = NewMonitorEstadoOtros()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}