// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ListInterface2018.avsc
 */
package HCMInterface2018Events

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayInterface2018Data(r []Interface2018Data, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeInterface2018Data(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayInterface2018DataWrapper struct {
	Target *[]Interface2018Data
}

func (_ ArrayInterface2018DataWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ ArrayInterface2018DataWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ ArrayInterface2018DataWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ ArrayInterface2018DataWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ ArrayInterface2018DataWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ ArrayInterface2018DataWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ ArrayInterface2018DataWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ ArrayInterface2018DataWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ ArrayInterface2018DataWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ ArrayInterface2018DataWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayInterface2018DataWrapper) Finalize()        {}
func (_ ArrayInterface2018DataWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r ArrayInterface2018DataWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]Interface2018Data, 0, s)
	}
}
func (r ArrayInterface2018DataWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayInterface2018DataWrapper) AppendArray() types.Field {
	var v Interface2018Data
	v = NewInterface2018Data()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}