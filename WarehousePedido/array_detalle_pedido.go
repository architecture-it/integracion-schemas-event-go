// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoDeAlmacenSCE.avsc
 */
package WarehousePedidoEvents

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayDetallePedido(r []DetallePedido, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeDetallePedido(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayDetallePedidoWrapper struct {
	Target *[]DetallePedido
}

func (_ ArrayDetallePedidoWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayDetallePedidoWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayDetallePedidoWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayDetallePedidoWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayDetallePedidoWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayDetallePedidoWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayDetallePedidoWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayDetallePedidoWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayDetallePedidoWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayDetallePedidoWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayDetallePedidoWrapper) Finalize()                        {}
func (_ ArrayDetallePedidoWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayDetallePedidoWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]DetallePedido, 0, s)
	}
}
func (r ArrayDetallePedidoWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayDetallePedidoWrapper) AppendArray() types.Field {
	var v DetallePedido
	v = NewDetallePedido()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
