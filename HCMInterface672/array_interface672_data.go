// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ListInterface672.avsc
 */
package HCMInterface672Events

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayInterface672Data(r []Interface672Data, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeInterface672Data(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayInterface672DataWrapper struct {
	Target *[]Interface672Data
}

func (_ ArrayInterface672DataWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ ArrayInterface672DataWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ ArrayInterface672DataWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ ArrayInterface672DataWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ ArrayInterface672DataWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ ArrayInterface672DataWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ ArrayInterface672DataWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ ArrayInterface672DataWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ ArrayInterface672DataWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ ArrayInterface672DataWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayInterface672DataWrapper) Finalize()        {}
func (_ ArrayInterface672DataWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r ArrayInterface672DataWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]Interface672Data, 0, s)
	}
}
func (r ArrayInterface672DataWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayInterface672DataWrapper) AppendArray() types.Field {
	var v Interface672Data
	v = NewInterface672Data()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
