// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Visita.avsc
 */
package TrazasEvents

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayMetadato(r []Metadato, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeMetadato(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayMetadatoWrapper struct {
	Target *[]Metadato
}

func (_ ArrayMetadatoWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayMetadatoWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayMetadatoWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayMetadatoWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayMetadatoWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayMetadatoWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayMetadatoWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayMetadatoWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayMetadatoWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayMetadatoWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayMetadatoWrapper) Finalize()                        {}
func (_ ArrayMetadatoWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayMetadatoWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]Metadato, 0, s)
	}
}
func (r ArrayMetadatoWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayMetadatoWrapper) AppendArray() types.Field {
	var v Metadato
	v = NewMetadato()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
