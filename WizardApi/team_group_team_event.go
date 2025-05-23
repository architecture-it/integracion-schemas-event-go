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

type TeamGroupTeamEvent struct {
	Role RoleEvent `json:"Role"`

	Team TeamEvent `json:"Team"`
}

const TeamGroupTeamEventAvroCRC64Fingerprint = "X\x1b\xf9ɲ\xe1\r4"

func NewTeamGroupTeamEvent() TeamGroupTeamEvent {
	r := TeamGroupTeamEvent{}
	r.Role = NewRoleEvent()

	r.Team = NewTeamEvent()

	return r
}

func DeserializeTeamGroupTeamEvent(r io.Reader) (TeamGroupTeamEvent, error) {
	t := NewTeamGroupTeamEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeTeamGroupTeamEventFromSchema(r io.Reader, schema string) (TeamGroupTeamEvent, error) {
	t := NewTeamGroupTeamEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeTeamGroupTeamEvent(r TeamGroupTeamEvent, w io.Writer) error {
	var err error
	err = writeRoleEvent(r.Role, w)
	if err != nil {
		return err
	}
	err = writeTeamEvent(r.Team, w)
	if err != nil {
		return err
	}
	return err
}

func (r TeamGroupTeamEvent) Serialize(w io.Writer) error {
	return writeTeamGroupTeamEvent(r, w)
}

func (r TeamGroupTeamEvent) Schema() string {
	return "{\"fields\":[{\"name\":\"Role\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"},{\"default\":null,\"name\":\"IsUserRole\",\"type\":[\"null\",\"boolean\"]}],\"name\":\"RoleEvent\",\"type\":\"record\"}},{\"name\":\"Team\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"long\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"SlugName\",\"type\":\"string\"},{\"name\":\"Organization\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"long\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Freeze\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"UserOrganizationRoles\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"Role\",\"type\":\"Andreani.WizardApi.Events.Record.RoleEvent\"},{\"name\":\"User\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"LoginName\",\"type\":\"string\"},{\"name\":\"UserName\",\"type\":\"string\"},{\"default\":null,\"name\":\"Email\",\"type\":[\"null\",\"string\"]}],\"name\":\"GithubUserEvent\",\"namespace\":\"Andreani.WizardApi.Events.Common\",\"type\":\"record\"}}],\"name\":\"UserOrganizationRoleEvent\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"OrganizationEvent\",\"type\":\"record\"}},{\"default\":[],\"name\":\"Users\",\"type\":{\"items\":\"Andreani.WizardApi.Events.Common.GithubUserEvent\",\"type\":\"array\"}}],\"name\":\"TeamEvent\",\"type\":\"record\"}}],\"name\":\"Andreani.WizardApi.Events.Record.TeamGroupTeamEvent\",\"type\":\"record\"}"
}

func (r TeamGroupTeamEvent) SchemaName() string {
	return "Andreani.WizardApi.Events.Record.TeamGroupTeamEvent"
}

func (_ TeamGroupTeamEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ TeamGroupTeamEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ TeamGroupTeamEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ TeamGroupTeamEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ TeamGroupTeamEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ TeamGroupTeamEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ TeamGroupTeamEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ TeamGroupTeamEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *TeamGroupTeamEvent) Get(i int) types.Field {
	switch i {
	case 0:
		r.Role = NewRoleEvent()

		w := types.Record{Target: &r.Role}

		return w

	case 1:
		r.Team = NewTeamEvent()

		w := types.Record{Target: &r.Team}

		return w

	}
	panic("Unknown field index")
}

func (r *TeamGroupTeamEvent) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *TeamGroupTeamEvent) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ TeamGroupTeamEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ TeamGroupTeamEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ TeamGroupTeamEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ TeamGroupTeamEvent) Finalize()                        {}

func (_ TeamGroupTeamEvent) AvroCRC64Fingerprint() []byte {
	return []byte(TeamGroupTeamEventAvroCRC64Fingerprint)
}

func (r TeamGroupTeamEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Role"], err = json.Marshal(r.Role)
	if err != nil {
		return nil, err
	}
	output["Team"], err = json.Marshal(r.Team)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *TeamGroupTeamEvent) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Role"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Role); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Role")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Team"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Team); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Team")
	}
	return nil
}
