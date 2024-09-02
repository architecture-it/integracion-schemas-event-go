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

func writeArrayCliente(r []Cliente, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeCliente(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayClienteWrapper struct {
	Target *[]Cliente
}

func (_ ArrayClienteWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayClienteWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayClienteWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayClienteWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayClienteWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayClienteWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayClienteWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayClienteWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayClienteWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayClienteWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayClienteWrapper) Finalize()                        {}
func (_ ArrayClienteWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayClienteWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]Cliente, 0, s)
	}
}
func (r ArrayClienteWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayClienteWrapper) AppendArray() types.Field {
	var v Cliente
	v = NewCliente()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
