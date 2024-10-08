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

func writeArrayDestinatario(r []Destinatario, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeDestinatario(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayDestinatarioWrapper struct {
	Target *[]Destinatario
}

func (_ ArrayDestinatarioWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayDestinatarioWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayDestinatarioWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayDestinatarioWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayDestinatarioWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayDestinatarioWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayDestinatarioWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayDestinatarioWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayDestinatarioWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayDestinatarioWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayDestinatarioWrapper) Finalize()                        {}
func (_ ArrayDestinatarioWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayDestinatarioWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]Destinatario, 0, s)
	}
}
func (r ArrayDestinatarioWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayDestinatarioWrapper) AppendArray() types.Field {
	var v Destinatario
	v = NewDestinatario()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
