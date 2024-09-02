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

func writeArrayEnvio(r []Envio, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeEnvio(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayEnvioWrapper struct {
	Target *[]Envio
}

func (_ ArrayEnvioWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayEnvioWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayEnvioWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayEnvioWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayEnvioWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayEnvioWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayEnvioWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayEnvioWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayEnvioWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayEnvioWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayEnvioWrapper) Finalize()                        {}
func (_ ArrayEnvioWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayEnvioWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]Envio, 0, s)
	}
}
func (r ArrayEnvioWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayEnvioWrapper) AppendArray() types.Field {
	var v Envio
	v = NewEnvio()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
