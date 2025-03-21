// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhRecepcionEliminada.avsc
 */
package EventoWhRecepcionEvents

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArraySeries(r []Series, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeSeries(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArraySeriesWrapper struct {
	Target *[]Series
}

func (_ ArraySeriesWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArraySeriesWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArraySeriesWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArraySeriesWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArraySeriesWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArraySeriesWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArraySeriesWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArraySeriesWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArraySeriesWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArraySeriesWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArraySeriesWrapper) Finalize()                        {}
func (_ ArraySeriesWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArraySeriesWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]Series, 0, s)
	}
}
func (r ArraySeriesWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArraySeriesWrapper) AppendArray() types.Field {
	var v Series
	v = NewSeries()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
