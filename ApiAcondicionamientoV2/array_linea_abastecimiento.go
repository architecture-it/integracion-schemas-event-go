// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     AcondicionamientoV2.avsc
 */
package ApiAcondicionamientoV2Events

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayLineaAbastecimiento(r []LineaAbastecimiento, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeLineaAbastecimiento(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayLineaAbastecimientoWrapper struct {
	Target *[]LineaAbastecimiento
}

func (_ ArrayLineaAbastecimientoWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ ArrayLineaAbastecimientoWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ ArrayLineaAbastecimientoWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ ArrayLineaAbastecimientoWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ ArrayLineaAbastecimientoWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ ArrayLineaAbastecimientoWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ ArrayLineaAbastecimientoWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ ArrayLineaAbastecimientoWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ ArrayLineaAbastecimientoWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ ArrayLineaAbastecimientoWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayLineaAbastecimientoWrapper) Finalize()        {}
func (_ ArrayLineaAbastecimientoWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r ArrayLineaAbastecimientoWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]LineaAbastecimiento, 0, s)
	}
}
func (r ArrayLineaAbastecimientoWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayLineaAbastecimientoWrapper) AppendArray() types.Field {
	var v LineaAbastecimiento
	v = NewLineaAbastecimiento()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
