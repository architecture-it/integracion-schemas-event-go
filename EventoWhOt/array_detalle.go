// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhOtRecepcionFinalizada.avsc
 */
package EventoWhOtEvents

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayDetalle(r []Detalle, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeDetalle(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayDetalleWrapper struct {
	Target *[]Detalle
}

func (_ ArrayDetalleWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayDetalleWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayDetalleWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayDetalleWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayDetalleWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayDetalleWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayDetalleWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayDetalleWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayDetalleWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayDetalleWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayDetalleWrapper) Finalize()                        {}
func (_ ArrayDetalleWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayDetalleWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]Detalle, 0, s)
	}
}
func (r ArrayDetalleWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayDetalleWrapper) AppendArray() types.Field {
	var v Detalle
	v = NewDetalle()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
