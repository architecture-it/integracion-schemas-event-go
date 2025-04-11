// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PipelineVersionEvent.avsc
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

type PipelineVersionEvent struct {
	Id int32 `json:"Id"`

	Version string `json:"Version"`

	ReleaseNotes string `json:"ReleaseNotes"`

	Latest bool `json:"Latest"`

	StatusVersion StatusEvent `json:"StatusVersion"`

	Pipeline PipelineEvent `json:"Pipeline"`

	AuditInfo AuditEvent `json:"AuditInfo"`
}

const PipelineVersionEventAvroCRC64Fingerprint = "S, hc\x95\xb3\xe7"

func NewPipelineVersionEvent() PipelineVersionEvent {
	r := PipelineVersionEvent{}
	r.StatusVersion = NewStatusEvent()

	r.Pipeline = NewPipelineEvent()

	r.AuditInfo = NewAuditEvent()

	return r
}

func DeserializePipelineVersionEvent(r io.Reader) (PipelineVersionEvent, error) {
	t := NewPipelineVersionEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePipelineVersionEventFromSchema(r io.Reader, schema string) (PipelineVersionEvent, error) {
	t := NewPipelineVersionEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePipelineVersionEvent(r PipelineVersionEvent, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Version, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.ReleaseNotes, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Latest, w)
	if err != nil {
		return err
	}
	err = writeStatusEvent(r.StatusVersion, w)
	if err != nil {
		return err
	}
	err = writePipelineEvent(r.Pipeline, w)
	if err != nil {
		return err
	}
	err = writeAuditEvent(r.AuditInfo, w)
	if err != nil {
		return err
	}
	return err
}

func (r PipelineVersionEvent) Serialize(w io.Writer) error {
	return writePipelineVersionEvent(r, w)
}

func (r PipelineVersionEvent) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Version\",\"type\":\"string\"},{\"name\":\"ReleaseNotes\",\"type\":\"string\"},{\"name\":\"Latest\",\"type\":\"boolean\"},{\"name\":\"StatusVersion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"}],\"name\":\"StatusEvent\",\"namespace\":\"Andreani.WizardApi.Events.Common\",\"type\":\"record\"}},{\"name\":\"Pipeline\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Description\",\"type\":\"string\"},{\"name\":\"Ref\",\"type\":\"string\"},{\"name\":\"RepositoryConfig\",\"type\":\"string\"},{\"name\":\"AuditInfo\",\"type\":{\"fields\":[{\"name\":\"CreateBy\",\"type\":\"string\"},{\"name\":\"CreateDate\",\"type\":{\"logicalType\":\"date\",\"type\":\"int\"}},{\"default\":null,\"name\":\"UpdateBy\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"UpdateDate\",\"type\":[\"null\",{\"logicalType\":\"date\",\"type\":\"int\"}]},{\"name\":\"Deleted\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"DeletedBy\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DeletedDate\",\"type\":[\"null\",{\"logicalType\":\"date\",\"type\":\"int\"}]}],\"name\":\"AuditEvent\",\"namespace\":\"Andreani.WizardApi.Events.Common\",\"type\":\"record\"}}],\"name\":\"PipelineEvent\",\"type\":\"record\"}},{\"name\":\"AuditInfo\",\"type\":\"Andreani.WizardApi.Events.Common.AuditEvent\"}],\"name\":\"Andreani.WizardApi.Events.Record.PipelineVersionEvent\",\"type\":\"record\"}"
}

func (r PipelineVersionEvent) SchemaName() string {
	return "Andreani.WizardApi.Events.Record.PipelineVersionEvent"
}

func (_ PipelineVersionEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ PipelineVersionEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ PipelineVersionEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ PipelineVersionEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ PipelineVersionEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ PipelineVersionEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ PipelineVersionEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ PipelineVersionEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *PipelineVersionEvent) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Version}

		return w

	case 2:
		w := types.String{Target: &r.ReleaseNotes}

		return w

	case 3:
		w := types.Boolean{Target: &r.Latest}

		return w

	case 4:
		r.StatusVersion = NewStatusEvent()

		w := types.Record{Target: &r.StatusVersion}

		return w

	case 5:
		r.Pipeline = NewPipelineEvent()

		w := types.Record{Target: &r.Pipeline}

		return w

	case 6:
		r.AuditInfo = NewAuditEvent()

		w := types.Record{Target: &r.AuditInfo}

		return w

	}
	panic("Unknown field index")
}

func (r *PipelineVersionEvent) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *PipelineVersionEvent) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ PipelineVersionEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ PipelineVersionEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ PipelineVersionEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ PipelineVersionEvent) Finalize()                        {}

func (_ PipelineVersionEvent) AvroCRC64Fingerprint() []byte {
	return []byte(PipelineVersionEventAvroCRC64Fingerprint)
}

func (r PipelineVersionEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["Version"], err = json.Marshal(r.Version)
	if err != nil {
		return nil, err
	}
	output["ReleaseNotes"], err = json.Marshal(r.ReleaseNotes)
	if err != nil {
		return nil, err
	}
	output["Latest"], err = json.Marshal(r.Latest)
	if err != nil {
		return nil, err
	}
	output["StatusVersion"], err = json.Marshal(r.StatusVersion)
	if err != nil {
		return nil, err
	}
	output["Pipeline"], err = json.Marshal(r.Pipeline)
	if err != nil {
		return nil, err
	}
	output["AuditInfo"], err = json.Marshal(r.AuditInfo)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *PipelineVersionEvent) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["Version"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Version); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Version")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ReleaseNotes"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ReleaseNotes); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ReleaseNotes")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Latest"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Latest); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Latest")
	}
	val = func() json.RawMessage {
		if v, ok := fields["StatusVersion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StatusVersion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for StatusVersion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Pipeline"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Pipeline); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Pipeline")
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
