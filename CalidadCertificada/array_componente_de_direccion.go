// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Postal.avsc
 */
package CalidadCertificadaEvents

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayComponenteDeDireccion(r []ComponenteDeDireccion, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeComponenteDeDireccion(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayComponenteDeDireccionWrapper struct {
	Target *[]ComponenteDeDireccion
}

func (_ ArrayComponenteDeDireccionWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ ArrayComponenteDeDireccionWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ ArrayComponenteDeDireccionWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ ArrayComponenteDeDireccionWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ ArrayComponenteDeDireccionWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ ArrayComponenteDeDireccionWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ ArrayComponenteDeDireccionWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ ArrayComponenteDeDireccionWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ ArrayComponenteDeDireccionWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ ArrayComponenteDeDireccionWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayComponenteDeDireccionWrapper) Finalize()        {}
func (_ ArrayComponenteDeDireccionWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r ArrayComponenteDeDireccionWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]ComponenteDeDireccion, 0, s)
	}
}
func (r ArrayComponenteDeDireccionWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayComponenteDeDireccionWrapper) AppendArray() types.Field {
	var v ComponenteDeDireccion
	v = NewComponenteDeDireccion()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
