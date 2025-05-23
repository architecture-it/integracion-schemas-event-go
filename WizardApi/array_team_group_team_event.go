// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TeamGroupEvent.avsc
 */
package WizardApiEvents

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayTeamGroupTeamEvent(r []TeamGroupTeamEvent, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeTeamGroupTeamEvent(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayTeamGroupTeamEventWrapper struct {
	Target *[]TeamGroupTeamEvent
}

func (_ ArrayTeamGroupTeamEventWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ ArrayTeamGroupTeamEventWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ ArrayTeamGroupTeamEventWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ ArrayTeamGroupTeamEventWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ ArrayTeamGroupTeamEventWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ ArrayTeamGroupTeamEventWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ ArrayTeamGroupTeamEventWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ ArrayTeamGroupTeamEventWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ ArrayTeamGroupTeamEventWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ ArrayTeamGroupTeamEventWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayTeamGroupTeamEventWrapper) Finalize()        {}
func (_ ArrayTeamGroupTeamEventWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r ArrayTeamGroupTeamEventWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]TeamGroupTeamEvent, 0, s)
	}
}
func (r ArrayTeamGroupTeamEventWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayTeamGroupTeamEventWrapper) AppendArray() types.Field {
	var v TeamGroupTeamEvent
	v = NewTeamGroupTeamEvent()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
