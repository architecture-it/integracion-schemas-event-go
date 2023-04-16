// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoDeAlmacenSCE.avsc
 */
package Wapv2Events

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayTiposDeAcondicionamientoSecundario(r []TiposDeAcondicionamientoSecundario, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeTiposDeAcondicionamientoSecundario(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayTiposDeAcondicionamientoSecundarioWrapper struct {
	Target *[]TiposDeAcondicionamientoSecundario
}

func (_ ArrayTiposDeAcondicionamientoSecundarioWrapper) SetBoolean(v bool) {
	panic("Unsupported operation")
}
func (_ ArrayTiposDeAcondicionamientoSecundarioWrapper) SetInt(v int32) {
	panic("Unsupported operation")
}
func (_ ArrayTiposDeAcondicionamientoSecundarioWrapper) SetLong(v int64) {
	panic("Unsupported operation")
}
func (_ ArrayTiposDeAcondicionamientoSecundarioWrapper) SetFloat(v float32) {
	panic("Unsupported operation")
}
func (_ ArrayTiposDeAcondicionamientoSecundarioWrapper) SetDouble(v float64) {
	panic("Unsupported operation")
}
func (_ ArrayTiposDeAcondicionamientoSecundarioWrapper) SetBytes(v []byte) {
	panic("Unsupported operation")
}
func (_ ArrayTiposDeAcondicionamientoSecundarioWrapper) SetString(v string) {
	panic("Unsupported operation")
}
func (_ ArrayTiposDeAcondicionamientoSecundarioWrapper) SetUnionElem(v int64) {
	panic("Unsupported operation")
}
func (_ ArrayTiposDeAcondicionamientoSecundarioWrapper) Get(i int) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayTiposDeAcondicionamientoSecundarioWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayTiposDeAcondicionamientoSecundarioWrapper) Finalize() {}
func (_ ArrayTiposDeAcondicionamientoSecundarioWrapper) SetDefault(i int) {
	panic("Unsupported operation")
}
func (r ArrayTiposDeAcondicionamientoSecundarioWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]TiposDeAcondicionamientoSecundario, 0, s)
	}
}
func (r ArrayTiposDeAcondicionamientoSecundarioWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayTiposDeAcondicionamientoSecundarioWrapper) AppendArray() types.Field {
	var v TiposDeAcondicionamientoSecundario
	v = NewTiposDeAcondicionamientoSecundario()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}