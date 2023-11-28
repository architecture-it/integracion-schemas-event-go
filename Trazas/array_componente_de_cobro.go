// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     GestionCobranzaV2.avsc
 */
package TrazasEvents

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayComponenteDeCobro(r []ComponenteDeCobro, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeComponenteDeCobro(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayComponenteDeCobroWrapper struct {
	Target *[]ComponenteDeCobro
}

func (_ ArrayComponenteDeCobroWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ ArrayComponenteDeCobroWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ ArrayComponenteDeCobroWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ ArrayComponenteDeCobroWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ ArrayComponenteDeCobroWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ ArrayComponenteDeCobroWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ ArrayComponenteDeCobroWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ ArrayComponenteDeCobroWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ ArrayComponenteDeCobroWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ ArrayComponenteDeCobroWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayComponenteDeCobroWrapper) Finalize()        {}
func (_ ArrayComponenteDeCobroWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r ArrayComponenteDeCobroWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]ComponenteDeCobro, 0, s)
	}
}
func (r ArrayComponenteDeCobroWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayComponenteDeCobroWrapper) AppendArray() types.Field {
	var v ComponenteDeCobro
	v = NewComponenteDeCobro()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
