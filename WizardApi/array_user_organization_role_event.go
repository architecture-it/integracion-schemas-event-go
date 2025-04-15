// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TeamGroupTeamEvent.avsc
 */
package WizardApiEvents

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayUserOrganizationRoleEvent(r []UserOrganizationRoleEvent, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeUserOrganizationRoleEvent(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayUserOrganizationRoleEventWrapper struct {
	Target *[]UserOrganizationRoleEvent
}

func (_ ArrayUserOrganizationRoleEventWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ ArrayUserOrganizationRoleEventWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ ArrayUserOrganizationRoleEventWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ ArrayUserOrganizationRoleEventWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ ArrayUserOrganizationRoleEventWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ ArrayUserOrganizationRoleEventWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ ArrayUserOrganizationRoleEventWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ ArrayUserOrganizationRoleEventWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ ArrayUserOrganizationRoleEventWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ ArrayUserOrganizationRoleEventWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayUserOrganizationRoleEventWrapper) Finalize()        {}
func (_ ArrayUserOrganizationRoleEventWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r ArrayUserOrganizationRoleEventWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]UserOrganizationRoleEvent, 0, s)
	}
}
func (r ArrayUserOrganizationRoleEventWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayUserOrganizationRoleEventWrapper) AppendArray() types.Field {
	var v UserOrganizationRoleEvent
	v = NewUserOrganizationRoleEvent()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
