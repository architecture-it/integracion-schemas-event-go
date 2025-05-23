// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TeamGroupTeamEvent.avsc
 */
package WizardApiEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type OrganizationEvent struct {
	Id int64 `json:"Id"`

	Name string `json:"Name"`

	Freeze bool `json:"Freeze"`

	UserOrganizationRoles *UnionNullArrayUserOrganizationRoleEvent `json:"UserOrganizationRoles"`
}

const OrganizationEventAvroCRC64Fingerprint = "\xd8UA\xf9P\a/\xb3"

func NewOrganizationEvent() OrganizationEvent {
	r := OrganizationEvent{}
	r.UserOrganizationRoles = nil
	return r
}

func DeserializeOrganizationEvent(r io.Reader) (OrganizationEvent, error) {
	t := NewOrganizationEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeOrganizationEventFromSchema(r io.Reader, schema string) (OrganizationEvent, error) {
	t := NewOrganizationEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeOrganizationEvent(r OrganizationEvent, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Freeze, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayUserOrganizationRoleEvent(r.UserOrganizationRoles, w)
	if err != nil {
		return err
	}
	return err
}

func (r OrganizationEvent) Serialize(w io.Writer) error {
	return writeOrganizationEvent(r, w)
}

func (r OrganizationEvent) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"long\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Freeze\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"UserOrganizationRoles\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"Role\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"},{\"default\":null,\"name\":\"IsUserRole\",\"type\":[\"null\",\"boolean\"]}],\"name\":\"RoleEvent\",\"type\":\"record\"}},{\"name\":\"User\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"LoginName\",\"type\":\"string\"},{\"name\":\"UserName\",\"type\":\"string\"},{\"default\":null,\"name\":\"Email\",\"type\":[\"null\",\"string\"]}],\"name\":\"GithubUserEvent\",\"namespace\":\"Andreani.WizardApi.Events.Common\",\"type\":\"record\"}}],\"name\":\"UserOrganizationRoleEvent\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"Andreani.WizardApi.Events.Record.OrganizationEvent\",\"type\":\"record\"}"
}

func (r OrganizationEvent) SchemaName() string {
	return "Andreani.WizardApi.Events.Record.OrganizationEvent"
}

func (_ OrganizationEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ OrganizationEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ OrganizationEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ OrganizationEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ OrganizationEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ OrganizationEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ OrganizationEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ OrganizationEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *OrganizationEvent) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Name}

		return w

	case 2:
		w := types.Boolean{Target: &r.Freeze}

		return w

	case 3:
		r.UserOrganizationRoles = NewUnionNullArrayUserOrganizationRoleEvent()

		return r.UserOrganizationRoles
	}
	panic("Unknown field index")
}

func (r *OrganizationEvent) SetDefault(i int) {
	switch i {
	case 3:
		r.UserOrganizationRoles = nil
		return
	}
	panic("Unknown field index")
}

func (r *OrganizationEvent) NullField(i int) {
	switch i {
	case 3:
		r.UserOrganizationRoles = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ OrganizationEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ OrganizationEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ OrganizationEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ OrganizationEvent) Finalize()                        {}

func (_ OrganizationEvent) AvroCRC64Fingerprint() []byte {
	return []byte(OrganizationEventAvroCRC64Fingerprint)
}

func (r OrganizationEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["Name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	output["Freeze"], err = json.Marshal(r.Freeze)
	if err != nil {
		return nil, err
	}
	output["UserOrganizationRoles"], err = json.Marshal(r.UserOrganizationRoles)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *OrganizationEvent) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Freeze"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Freeze); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Freeze")
	}
	val = func() json.RawMessage {
		if v, ok := fields["UserOrganizationRoles"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UserOrganizationRoles); err != nil {
			return err
		}
	} else {
		r.UserOrganizationRoles = NewUnionNullArrayUserOrganizationRoleEvent()

		r.UserOrganizationRoles = nil
	}
	return nil
}
