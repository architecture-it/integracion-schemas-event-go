// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ApplicationEvent.avsc
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

type ApplicationEvent struct {
	Id int32 `json:"Id"`

	Name string `json:"Name"`

	Description string `json:"Description"`

	IsMigration bool `json:"IsMigration"`

	JsonData string `json:"JsonData"`

	Project ProjectEvent `json:"Project"`

	Template TemplateEvent `json:"Template"`

	PipelineVersion PipelineVersionEvent `json:"PipelineVersion"`

	Repository RepositoryEvent `json:"Repository"`

	Status StatusEvent `json:"Status"`

	AuditInfo AuditEvent `json:"AuditInfo"`
}

const ApplicationEventAvroCRC64Fingerprint = "\x8f\xa2\xa3:\\`\xee2"

func NewApplicationEvent() ApplicationEvent {
	r := ApplicationEvent{}
	r.Project = NewProjectEvent()

	r.Template = NewTemplateEvent()

	r.PipelineVersion = NewPipelineVersionEvent()

	r.Repository = NewRepositoryEvent()

	r.Status = NewStatusEvent()

	r.AuditInfo = NewAuditEvent()

	return r
}

func DeserializeApplicationEvent(r io.Reader) (ApplicationEvent, error) {
	t := NewApplicationEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeApplicationEventFromSchema(r io.Reader, schema string) (ApplicationEvent, error) {
	t := NewApplicationEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeApplicationEvent(r ApplicationEvent, w io.Writer) error {
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
	err = vm.WriteBool(r.IsMigration, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.JsonData, w)
	if err != nil {
		return err
	}
	err = writeProjectEvent(r.Project, w)
	if err != nil {
		return err
	}
	err = writeTemplateEvent(r.Template, w)
	if err != nil {
		return err
	}
	err = writePipelineVersionEvent(r.PipelineVersion, w)
	if err != nil {
		return err
	}
	err = writeRepositoryEvent(r.Repository, w)
	if err != nil {
		return err
	}
	err = writeStatusEvent(r.Status, w)
	if err != nil {
		return err
	}
	err = writeAuditEvent(r.AuditInfo, w)
	if err != nil {
		return err
	}
	return err
}

func (r ApplicationEvent) Serialize(w io.Writer) error {
	return writeApplicationEvent(r, w)
}

func (r ApplicationEvent) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Description\",\"type\":\"string\"},{\"name\":\"IsMigration\",\"type\":\"boolean\"},{\"name\":\"JsonData\",\"type\":\"string\"},{\"name\":\"Project\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Acronym\",\"type\":\"string\"},{\"name\":\"Description\",\"type\":\"string\"},{\"name\":\"OwnerMail\",\"type\":\"string\"},{\"name\":\"OrganizationId\",\"type\":\"long\"},{\"name\":\"JsonData\",\"type\":\"string\"},{\"name\":\"AuditInfo\",\"type\":{\"fields\":[{\"name\":\"CreateBy\",\"type\":\"string\"},{\"name\":\"CreateDate\",\"type\":{\"logicalType\":\"date\",\"type\":\"int\"}},{\"default\":null,\"name\":\"UpdateBy\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"UpdateDate\",\"type\":[\"null\",{\"logicalType\":\"date\",\"type\":\"int\"}]},{\"name\":\"Deleted\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"DeletedBy\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DeletedDate\",\"type\":[\"null\",{\"logicalType\":\"date\",\"type\":\"int\"}]}],\"name\":\"AuditEvent\",\"namespace\":\"Andreani.WizardApi.Events.Common\",\"type\":\"record\"}}],\"name\":\"ProjectEvent\",\"type\":\"record\"}},{\"name\":\"Template\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Alias\",\"type\":\"string\"},{\"name\":\"IsTemplate\",\"type\":\"boolean\"},{\"name\":\"Active\",\"type\":\"boolean\"},{\"name\":\"FrameworkVersion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"},{\"name\":\"Latest\",\"type\":\"boolean\"},{\"name\":\"StatusVersion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"}],\"name\":\"StatusEvent\",\"namespace\":\"Andreani.WizardApi.Events.Common\",\"type\":\"record\"}},{\"name\":\"Framework\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Active\",\"type\":\"boolean\"},{\"name\":\"Language\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Name\",\"type\":\"string\"}],\"name\":\"LanguageEvent\",\"type\":\"record\"}}],\"name\":\"FrameworkEvent\",\"type\":\"record\"}}],\"name\":\"FrameworkVersionEvent\",\"type\":\"record\"}},{\"name\":\"Workflow\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"}],\"name\":\"WorkflowCiEvent\",\"type\":\"record\"}}],\"name\":\"TemplateEvent\",\"type\":\"record\"}},{\"name\":\"PipelineVersion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Version\",\"type\":\"string\"},{\"name\":\"ReleaseNotes\",\"type\":\"string\"},{\"name\":\"Latest\",\"type\":\"boolean\"},{\"name\":\"StatusVersion\",\"type\":\"Andreani.WizardApi.Events.Common.StatusEvent\"},{\"name\":\"Pipeline\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Description\",\"type\":\"string\"},{\"name\":\"Ref\",\"type\":\"string\"},{\"name\":\"RepositoryConfig\",\"type\":\"string\"},{\"name\":\"AuditInfo\",\"type\":\"Andreani.WizardApi.Events.Common.AuditEvent\"}],\"name\":\"PipelineEvent\",\"type\":\"record\"}},{\"name\":\"AuditInfo\",\"type\":\"Andreani.WizardApi.Events.Common.AuditEvent\"}],\"name\":\"PipelineVersionEvent\",\"type\":\"record\"}},{\"name\":\"Repository\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"RepositoryId\",\"type\":\"long\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Description\",\"type\":\"string\"},{\"name\":\"TeamGroup\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Name\",\"type\":\"string\"},{\"default\":[],\"name\":\"TeamGroupTeams\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Role\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"},{\"default\":null,\"name\":\"IsUserRole\",\"type\":[\"null\",\"boolean\"]}],\"name\":\"RoleEvent\",\"type\":\"record\"}},{\"name\":\"Team\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"long\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"SlugName\",\"type\":\"string\"},{\"name\":\"Organization\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"long\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Freeze\",\"type\":\"boolean\"},{\"name\":\"UserOrganizationRoles\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Role\",\"type\":\"Andreani.WizardApi.Events.Record.RoleEvent\"},{\"name\":\"User\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"LoginName\",\"type\":\"string\"},{\"name\":\"UserName\",\"type\":\"string\"},{\"default\":null,\"name\":\"Email\",\"type\":[\"null\",\"string\"]}],\"name\":\"GithubUserEvent\",\"namespace\":\"Andreani.WizardApi.Events.Common\",\"type\":\"record\"}}],\"name\":\"UserOrganizationRoleEvent\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"OrganizationEvent\",\"type\":\"record\"}},{\"default\":[],\"name\":\"Users\",\"type\":{\"items\":\"Andreani.WizardApi.Events.Common.GithubUserEvent\",\"type\":\"array\"}}],\"name\":\"TeamEvent\",\"type\":\"record\"}}],\"name\":\"TeamGroupTeamEvent\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"TeamGroupEvent\",\"type\":\"record\"}},{\"name\":\"User\",\"type\":\"Andreani.WizardApi.Events.Common.GithubUserEvent\"},{\"name\":\"Organization\",\"type\":\"Andreani.WizardApi.Events.Record.OrganizationEvent\"},{\"default\":[],\"name\":\"Variables\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Id\",\"type\":\"long\"},{\"name\":\"Key\",\"type\":\"string\"},{\"name\":\"Value\",\"type\":\"string\"},{\"name\":\"AuditInfo\",\"type\":\"Andreani.WizardApi.Events.Common.AuditEvent\"}],\"name\":\"VariableEvent\",\"type\":\"record\"},\"type\":\"array\"}},{\"default\":[],\"name\":\"Secrets\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Key\",\"type\":\"string\"},{\"name\":\"AuditInfo\",\"type\":\"Andreani.WizardApi.Events.Common.AuditEvent\"}],\"name\":\"SecretEvent\",\"type\":\"record\"},\"type\":\"array\"}},{\"default\":[],\"name\":\"Environments\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Id\",\"type\":\"long\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"HostName\",\"type\":\"string\"},{\"name\":\"IsProduction\",\"type\":\"boolean\"},{\"name\":\"Schedule\",\"type\":\"string\"},{\"name\":\"MatrixDeploy\",\"type\":\"string\"},{\"default\":[],\"name\":\"EnvironmentVariables\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Id\",\"type\":\"long\"},{\"name\":\"Key\",\"type\":\"string\"},{\"name\":\"Value\",\"type\":\"string\"},{\"name\":\"AuditInfo\",\"type\":\"Andreani.WizardApi.Events.Common.AuditEvent\"}],\"name\":\"EnvironmentVariableEvent\",\"type\":\"record\"},\"type\":\"array\"}},{\"default\":[],\"name\":\"EnvironmentSecrets\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Key\",\"type\":\"string\"},{\"name\":\"AuditInfo\",\"type\":\"Andreani.WizardApi.Events.Common.AuditEvent\"}],\"name\":\"EnvironmentSecretEvent\",\"type\":\"record\"},\"type\":\"array\"}},{\"default\":[],\"name\":\"ApprovalTeams\",\"type\":{\"items\":\"Andreani.WizardApi.Events.Record.TeamEvent\",\"type\":\"array\"}},{\"name\":\"AuditInfo\",\"type\":\"Andreani.WizardApi.Events.Common.AuditEvent\"}],\"name\":\"EnvironmentEvent\",\"type\":\"record\"},\"type\":\"array\"}},{\"name\":\"AuditInfo\",\"type\":\"Andreani.WizardApi.Events.Common.AuditEvent\"}],\"name\":\"RepositoryEvent\",\"type\":\"record\"}},{\"name\":\"Status\",\"type\":\"Andreani.WizardApi.Events.Common.StatusEvent\"},{\"name\":\"AuditInfo\",\"type\":\"Andreani.WizardApi.Events.Common.AuditEvent\"}],\"name\":\"Andreani.WizardApi.Events.Record.ApplicationEvent\",\"type\":\"record\"}"
}

func (r ApplicationEvent) SchemaName() string {
	return "Andreani.WizardApi.Events.Record.ApplicationEvent"
}

func (_ ApplicationEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ApplicationEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ApplicationEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ApplicationEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ApplicationEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ApplicationEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ApplicationEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ ApplicationEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ApplicationEvent) Get(i int) types.Field {
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
		w := types.Boolean{Target: &r.IsMigration}

		return w

	case 4:
		w := types.String{Target: &r.JsonData}

		return w

	case 5:
		r.Project = NewProjectEvent()

		w := types.Record{Target: &r.Project}

		return w

	case 6:
		r.Template = NewTemplateEvent()

		w := types.Record{Target: &r.Template}

		return w

	case 7:
		r.PipelineVersion = NewPipelineVersionEvent()

		w := types.Record{Target: &r.PipelineVersion}

		return w

	case 8:
		r.Repository = NewRepositoryEvent()

		w := types.Record{Target: &r.Repository}

		return w

	case 9:
		r.Status = NewStatusEvent()

		w := types.Record{Target: &r.Status}

		return w

	case 10:
		r.AuditInfo = NewAuditEvent()

		w := types.Record{Target: &r.AuditInfo}

		return w

	}
	panic("Unknown field index")
}

func (r *ApplicationEvent) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ApplicationEvent) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ApplicationEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ApplicationEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ApplicationEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ ApplicationEvent) Finalize()                        {}

func (_ ApplicationEvent) AvroCRC64Fingerprint() []byte {
	return []byte(ApplicationEventAvroCRC64Fingerprint)
}

func (r ApplicationEvent) MarshalJSON() ([]byte, error) {
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
	output["IsMigration"], err = json.Marshal(r.IsMigration)
	if err != nil {
		return nil, err
	}
	output["JsonData"], err = json.Marshal(r.JsonData)
	if err != nil {
		return nil, err
	}
	output["Project"], err = json.Marshal(r.Project)
	if err != nil {
		return nil, err
	}
	output["Template"], err = json.Marshal(r.Template)
	if err != nil {
		return nil, err
	}
	output["PipelineVersion"], err = json.Marshal(r.PipelineVersion)
	if err != nil {
		return nil, err
	}
	output["Repository"], err = json.Marshal(r.Repository)
	if err != nil {
		return nil, err
	}
	output["Status"], err = json.Marshal(r.Status)
	if err != nil {
		return nil, err
	}
	output["AuditInfo"], err = json.Marshal(r.AuditInfo)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ApplicationEvent) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["IsMigration"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IsMigration); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IsMigration")
	}
	val = func() json.RawMessage {
		if v, ok := fields["JsonData"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.JsonData); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for JsonData")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Project"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Project); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Project")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Template"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Template); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Template")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PipelineVersion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PipelineVersion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PipelineVersion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Repository"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Repository); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Repository")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Status"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Status); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Status")
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
