// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoAcondi.avsc
 */
package appnameEvents

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayLineaPedido(r []LineaPedido, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeLineaPedido(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayLineaPedidoWrapper struct {
	Target *[]LineaPedido
}

func (_ ArrayLineaPedidoWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayLineaPedidoWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayLineaPedidoWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayLineaPedidoWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayLineaPedidoWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayLineaPedidoWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayLineaPedidoWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayLineaPedidoWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayLineaPedidoWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayLineaPedidoWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayLineaPedidoWrapper) Finalize()                        {}
func (_ ArrayLineaPedidoWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayLineaPedidoWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]LineaPedido, 0, s)
	}
}
func (r ArrayLineaPedidoWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayLineaPedidoWrapper) AppendArray() types.Field {
	var v LineaPedido
	v = NewLineaPedido()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}