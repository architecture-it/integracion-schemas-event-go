// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     RepublicOrder.avsc
 */
package RepublicOrder_SchemeEvents

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayString(r []string, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = vm.WriteString(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayStringWrapper struct {
	Target *[]string
}

func (_ ArrayStringWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayStringWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayStringWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayStringWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayStringWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayStringWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayStringWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayStringWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayStringWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayStringWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayStringWrapper) Finalize()                        {}
func (_ ArrayStringWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayStringWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]string, 0, s)
	}
}
func (r ArrayStringWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayStringWrapper) AppendArray() types.Field {
	var v string

	*r.Target = append(*r.Target, v)
	return &types.String{Target: &(*r.Target)[len(*r.Target)-1]}
}
