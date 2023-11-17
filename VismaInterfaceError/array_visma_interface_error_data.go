// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ListVismaInterfaceError.avsc
 */
package VismaInterfaceErrorEvents

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayVismaInterfaceErrorData(r []VismaInterfaceErrorData, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeVismaInterfaceErrorData(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayVismaInterfaceErrorDataWrapper struct {
	Target *[]VismaInterfaceErrorData
}

func (_ ArrayVismaInterfaceErrorDataWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ ArrayVismaInterfaceErrorDataWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ ArrayVismaInterfaceErrorDataWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ ArrayVismaInterfaceErrorDataWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ ArrayVismaInterfaceErrorDataWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ ArrayVismaInterfaceErrorDataWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ ArrayVismaInterfaceErrorDataWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ ArrayVismaInterfaceErrorDataWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ ArrayVismaInterfaceErrorDataWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ ArrayVismaInterfaceErrorDataWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayVismaInterfaceErrorDataWrapper) Finalize()        {}
func (_ ArrayVismaInterfaceErrorDataWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r ArrayVismaInterfaceErrorDataWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]VismaInterfaceErrorData, 0, s)
	}
}
func (r ArrayVismaInterfaceErrorDataWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayVismaInterfaceErrorDataWrapper) AppendArray() types.Field {
	var v VismaInterfaceErrorData
	v = NewVismaInterfaceErrorData()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
