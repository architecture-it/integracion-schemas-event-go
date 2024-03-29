// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TransactionEvent.avsc
 */
package AltaOrdenEnvioEvents

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayBultoRequest(r []BultoRequest, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeBultoRequest(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayBultoRequestWrapper struct {
	Target *[]BultoRequest
}

func (_ ArrayBultoRequestWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayBultoRequestWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayBultoRequestWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayBultoRequestWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayBultoRequestWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayBultoRequestWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayBultoRequestWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayBultoRequestWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayBultoRequestWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayBultoRequestWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayBultoRequestWrapper) Finalize()                        {}
func (_ ArrayBultoRequestWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayBultoRequestWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]BultoRequest, 0, s)
	}
}
func (r ArrayBultoRequestWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayBultoRequestWrapper) AppendArray() types.Field {
	var v BultoRequest
	v = NewBultoRequest()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
