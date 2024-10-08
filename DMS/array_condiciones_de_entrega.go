// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     HdrCreada.avsc
 */
package DMSEvents

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayCondicionesDeEntrega(r []CondicionesDeEntrega, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeCondicionesDeEntrega(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayCondicionesDeEntregaWrapper struct {
	Target *[]CondicionesDeEntrega
}

func (_ ArrayCondicionesDeEntregaWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ ArrayCondicionesDeEntregaWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ ArrayCondicionesDeEntregaWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ ArrayCondicionesDeEntregaWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ ArrayCondicionesDeEntregaWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ ArrayCondicionesDeEntregaWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ ArrayCondicionesDeEntregaWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ ArrayCondicionesDeEntregaWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ ArrayCondicionesDeEntregaWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ ArrayCondicionesDeEntregaWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayCondicionesDeEntregaWrapper) Finalize()        {}
func (_ ArrayCondicionesDeEntregaWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r ArrayCondicionesDeEntregaWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]CondicionesDeEntrega, 0, s)
	}
}
func (r ArrayCondicionesDeEntregaWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayCondicionesDeEntregaWrapper) AppendArray() types.Field {
	var v CondicionesDeEntrega
	v = NewCondicionesDeEntrega()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
