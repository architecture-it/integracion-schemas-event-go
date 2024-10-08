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

func writeArrayDireccion(r []Direccion, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeDireccion(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayDireccionWrapper struct {
	Target *[]Direccion
}

func (_ ArrayDireccionWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayDireccionWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayDireccionWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayDireccionWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayDireccionWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayDireccionWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayDireccionWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayDireccionWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayDireccionWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayDireccionWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayDireccionWrapper) Finalize()                        {}
func (_ ArrayDireccionWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayDireccionWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]Direccion, 0, s)
	}
}
func (r ArrayDireccionWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayDireccionWrapper) AppendArray() types.Field {
	var v Direccion
	v = NewDireccion()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
