// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     GestionCobranza.avsc
 */
package TrazasEvents

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayDetalleDelCobro(r []DetalleDelCobro, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeDetalleDelCobro(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayDetalleDelCobroWrapper struct {
	Target *[]DetalleDelCobro
}

func (_ ArrayDetalleDelCobroWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ ArrayDetalleDelCobroWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ ArrayDetalleDelCobroWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ ArrayDetalleDelCobroWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ ArrayDetalleDelCobroWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ ArrayDetalleDelCobroWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ ArrayDetalleDelCobroWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ ArrayDetalleDelCobroWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ ArrayDetalleDelCobroWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ ArrayDetalleDelCobroWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayDetalleDelCobroWrapper) Finalize()        {}
func (_ ArrayDetalleDelCobroWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r ArrayDetalleDelCobroWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]DetalleDelCobro, 0, s)
	}
}
func (r ArrayDetalleDelCobroWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayDetalleDelCobroWrapper) AppendArray() types.Field {
	var v DetalleDelCobro
	v = NewDetalleDelCobro()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}