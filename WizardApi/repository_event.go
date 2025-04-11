// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     RepositoryEvent.avsc
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

type RepositoryEvent struct {
	Id int32 `json:"Id"`

	Name string `json:"Name"`

	Description string `json:"Description"`

	TeamGroup TeamGroupEvent `json:"TeamGroup"`

	User GithubUserEvent `json:"User"`

	Organization OrganizationEvent `json:"Organization"`

	Variables []VariableEvent `json:"Variables"`

	Secrets []SecretEvent `json:"Secrets"`

	Environments []EnvironmentEvent `json:"Environments"`

	AuditInfo AuditEvent `json:"AuditInfo"`
}

const RepositoryEventAvroCRC64Fingerprint = "\x10i\xbe\x87\x95\x11+\xb1"

func NewRepositoryEvent() RepositoryEvent {
	r := RepositoryEvent{}
	r.TeamGroup = NewTeamGroupEvent()

	r.User = NewGithubUserEvent()

	r.Organization = NewOrganizationEvent()

	r.Variables = make([]VariableEvent, 0)

	r.Variables = make([]VariableEvent, 0)

	r.Secrets = make([]SecretEvent, 0)

	r.Secrets = make([]SecretEvent, 0)

	r.Environments = make([]EnvironmentEvent, 0)

	r.Environments = make([]EnvironmentEvent, 0)

	r.AuditInfo = NewAuditEvent()

	return r
}

func DeserializeRepositoryEvent(r io.Reader) (RepositoryEvent, error) {
	t := NewRepositoryEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRepositoryEventFromSchema(r io.Reader, schema string) (RepositoryEvent, error) {
	t := NewRepositoryEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRepositoryEvent(r RepositoryEvent, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Description, w)
	if err != nil {
		return err
	}
	err = writeTeamGroupEvent(r.TeamGroup, w)
	if err != nil {
		return err
	}
	err = writeGithubUserEvent(r.User, w)
	if err != nil {
		return err
	}
	err = writeOrganizationEvent(r.Organization, w)
	if err != nil {
		return err
	}
	err = writeArrayVariableEvent(r.Variables, w)
	if err != nil {
		return err
	}
	err = writeArraySecretEvent(r.Secrets, w)
	if err != nil {
		return err
	}
	err = writeArrayEnvironmentEvent(r.Environments, w)
	if err != nil {
		return err
	}
	err = writeAuditEvent(r.AuditInfo, w)
	if err != nil {
		return err
	}
	return err
}

func (r RepositoryEvent) Serialize(w io.Writer) error {
	return writeRepositoryEvent(r, w)
}

func (r RepositoryEvent) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Description\",\"type\":\"string\"},{\"name\":\"TeamGroup\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Name\",\"type\":\"string\"},{\"default\":[],\"name\":\"TeamGroupTeams\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Role\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"},{\"default\":null,\"name\":\"IsUserRole\",\"type\":[\"null\",\"boolean\"]}],\"name\":\"RoleEvent\",\"type\":\"record\"}},{\"name\":\"Team\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"long\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"SlugName\",\"type\":\"string\"},{\"name\":\"Organization\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"long\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Freeze\",\"type\":\"boolean\"},{\"name\":\"UserOrganizationRole\",\"type\":{\"fields\":[{\"name\":\"Role\",\"type\":\"Andreani.WizardApi.Events.Record.RoleEvent\"},{\"name\":\"User\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"LoginName\",\"type\":\"string\"},{\"name\":\"UserName\",\"type\":\"string\"},{\"default\":null,\"name\":\"Email\",\"type\":[\"null\",\"string\"]}],\"name\":\"GithubUserEvent\",\"namespace\":\"Andreani.WizardApi.Events.Common\",\"type\":\"record\"}}],\"name\":\"UserOrganizationRoleEvent\",\"type\":\"record\"}}],\"name\":\"OrganizationEvent\",\"type\":\"record\"}},{\"default\":[],\"name\":\"Users\",\"type\":{\"items\":\"Andreani.WizardApi.Events.Common.GithubUserEvent\",\"type\":\"array\"}}],\"name\":\"TeamEvent\",\"type\":\"record\"}}],\"name\":\"TeamGroupTeamEvent\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"TeamGroupEvent\",\"type\":\"record\"}},{\"name\":\"User\",\"type\":\"Andreani.WizardApi.Events.Common.GithubUserEvent\"},{\"name\":\"Organization\",\"type\":\"Andreani.WizardApi.Events.Record.OrganizationEvent\"},{\"default\":[],\"name\":\"Variables\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Id\",\"type\":\"long\"},{\"name\":\"Key\",\"type\":\"string\"},{\"name\":\"Value\",\"type\":\"string\"},{\"name\":\"AuditInfo\",\"type\":{\"fields\":[{\"name\":\"CreateBy\",\"type\":\"string\"},{\"name\":\"CreateDate\",\"type\":{\"logicalType\":\"date\",\"type\":\"int\"}},{\"default\":null,\"name\":\"UpdateBy\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"UpdateDate\",\"type\":[\"null\",{\"logicalType\":\"date\",\"type\":\"int\"}]},{\"name\":\"Deleted\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"DeletedBy\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DeletedDate\",\"type\":[\"null\",{\"logicalType\":\"date\",\"type\":\"int\"}]}],\"name\":\"AuditEvent\",\"namespace\":\"Andreani.WizardApi.Events.Common\",\"type\":\"record\"}}],\"name\":\"VariableEvent\",\"type\":\"record\"},\"type\":\"array\"}},{\"default\":[],\"name\":\"Secrets\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Key\",\"type\":\"string\"},{\"name\":\"AuditInfo\",\"type\":\"Andreani.WizardApi.Events.Common.AuditEvent\"}],\"name\":\"SecretEvent\",\"type\":\"record\"},\"type\":\"array\"}},{\"default\":[],\"name\":\"Environments\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Id\",\"type\":\"long\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"HostName\",\"type\":\"string\"},{\"name\":\"IsProduction\",\"type\":\"boolean\"},{\"name\":\"Schedule\",\"type\":\"string\"},{\"name\":\"MatrixDeploy\",\"type\":\"string\"},{\"default\":[],\"name\":\"EnvironmentVariables\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Id\",\"type\":\"long\"},{\"name\":\"Key\",\"type\":\"string\"},{\"name\":\"Value\",\"type\":\"string\"},{\"name\":\"AuditInfo\",\"type\":\"Andreani.WizardApi.Events.Common.AuditEvent\"}],\"name\":\"EnvironmentVariableEvent\",\"type\":\"record\"},\"type\":\"array\"}},{\"default\":[],\"name\":\"EnvironmentSecrets\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Key\",\"type\":\"string\"},{\"name\":\"AuditInfo\",\"type\":\"Andreani.WizardApi.Events.Common.AuditEvent\"}],\"name\":\"EnvironmentSecretEvent\",\"type\":\"record\"},\"type\":\"array\"}},{\"default\":[],\"name\":\"ApprovalTeams\",\"type\":{\"items\":\"Andreani.WizardApi.Events.Record.TeamEvent\",\"type\":\"array\"}},{\"name\":\"AuditInfo\",\"type\":\"Andreani.WizardApi.Events.Common.AuditEvent\"}],\"name\":\"EnvironmentEvent\",\"type\":\"record\"},\"type\":\"array\"}},{\"name\":\"AuditInfo\",\"type\":\"Andreani.WizardApi.Events.Common.AuditEvent\"}],\"name\":\"Andreani.WizardApi.Events.Record.RepositoryEvent\",\"type\":\"record\"}"
}

func (r RepositoryEvent) SchemaName() string {
	return "Andreani.WizardApi.Events.Record.RepositoryEvent"
}

func (_ RepositoryEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ RepositoryEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ RepositoryEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ RepositoryEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ RepositoryEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ RepositoryEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ RepositoryEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ RepositoryEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RepositoryEvent) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Name}

		return w

	case 2:
		w := types.String{Target: &r.Description}

		return w

	case 3:
		r.TeamGroup = NewTeamGroupEvent()

		w := types.Record{Target: &r.TeamGroup}

		return w

	case 4:
		r.User = NewGithubUserEvent()

		w := types.Record{Target: &r.User}

		return w

	case 5:
		r.Organization = NewOrganizationEvent()

		w := types.Record{Target: &r.Organization}

		return w

	case 6:
		r.Variables = make([]VariableEvent, 0)

		w := ArrayVariableEventWrapper{Target: &r.Variables}

		return w

	case 7:
		r.Secrets = make([]SecretEvent, 0)

		w := ArraySecretEventWrapper{Target: &r.Secrets}

		return w

	case 8:
		r.Environments = make([]EnvironmentEvent, 0)

		w := ArrayEnvironmentEventWrapper{Target: &r.Environments}

		return w

	case 9:
		r.AuditInfo = NewAuditEvent()

		w := types.Record{Target: &r.AuditInfo}

		return w

	}
	panic("Unknown field index")
}

func (r *RepositoryEvent) SetDefault(i int) {
	switch i {
	case 6:
		r.Variables = make([]VariableEvent, 0)

		return
	case 7:
		r.Secrets = make([]SecretEvent, 0)

		return
	case 8:
		r.Environments = make([]EnvironmentEvent, 0)

		return
	}
	panic("Unknown field index")
}

func (r *RepositoryEvent) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ RepositoryEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ RepositoryEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ RepositoryEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ RepositoryEvent) Finalize()                        {}

func (_ RepositoryEvent) AvroCRC64Fingerprint() []byte {
	return []byte(RepositoryEventAvroCRC64Fingerprint)
}

func (r RepositoryEvent) MarshalJSON() ([]byte, error) {
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
	output["Description"], err = json.Marshal(r.Description)
	if err != nil {
		return nil, err
	}
	output["TeamGroup"], err = json.Marshal(r.TeamGroup)
	if err != nil {
		return nil, err
	}
	output["User"], err = json.Marshal(r.User)
	if err != nil {
		return nil, err
	}
	output["Organization"], err = json.Marshal(r.Organization)
	if err != nil {
		return nil, err
	}
	output["Variables"], err = json.Marshal(r.Variables)
	if err != nil {
		return nil, err
	}
	output["Secrets"], err = json.Marshal(r.Secrets)
	if err != nil {
		return nil, err
	}
	output["Environments"], err = json.Marshal(r.Environments)
	if err != nil {
		return nil, err
	}
	output["AuditInfo"], err = json.Marshal(r.AuditInfo)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *RepositoryEvent) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["Description"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Description); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Description")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TeamGroup"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TeamGroup); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TeamGroup")
	}
	val = func() json.RawMessage {
		if v, ok := fields["User"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.User); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for User")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Organization"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Organization); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Organization")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Variables"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Variables); err != nil {
			return err
		}
	} else {
		r.Variables = make([]VariableEvent, 0)

		r.Variables = make([]VariableEvent, 0)

	}
	val = func() json.RawMessage {
		if v, ok := fields["Secrets"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Secrets); err != nil {
			return err
		}
	} else {
		r.Secrets = make([]SecretEvent, 0)

		r.Secrets = make([]SecretEvent, 0)

	}
	val = func() json.RawMessage {
		if v, ok := fields["Environments"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Environments); err != nil {
			return err
		}
	} else {
		r.Environments = make([]EnvironmentEvent, 0)

		r.Environments = make([]EnvironmentEvent, 0)

	}
	val = func() json.RawMessage {
		if v, ok := fields["AuditInfo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AuditInfo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for AuditInfo")
	}
	return nil
}
