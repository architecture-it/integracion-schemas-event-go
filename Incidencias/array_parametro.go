// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     InsertaIncidencia.avsc
 */
package IncidenciasEvents

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayParametro(r []Parametro, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeParametro(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayParametroWrapper struct {
	Target *[]Parametro
}

func (_ ArrayParametroWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayParametroWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayParametroWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayParametroWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayParametroWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayParametroWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayParametroWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayParametroWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayParametroWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayParametroWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayParametroWrapper) Finalize()                        {}
func (_ ArrayParametroWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayParametroWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]Parametro, 0, s)
	}
}
func (r ArrayParametroWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayParametroWrapper) AppendArray() types.Field {
	var v Parametro
	v = NewParametro()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}