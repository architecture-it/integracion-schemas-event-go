// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NotificacionFord.avsc
 */
package NotificacionesEvents

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayDetalleFord(r []DetalleFord, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeDetalleFord(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayDetalleFordWrapper struct {
	Target *[]DetalleFord
}

func (_ ArrayDetalleFordWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayDetalleFordWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayDetalleFordWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayDetalleFordWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayDetalleFordWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayDetalleFordWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayDetalleFordWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayDetalleFordWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayDetalleFordWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayDetalleFordWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayDetalleFordWrapper) Finalize()                        {}
func (_ ArrayDetalleFordWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayDetalleFordWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]DetalleFord, 0, s)
	}
}
func (r ArrayDetalleFordWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayDetalleFordWrapper) AppendArray() types.Field {
	var v DetalleFord
	v = NewDetalleFord()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
