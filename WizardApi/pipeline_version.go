// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PipelineVersion.avsc
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

type PipelineVersion struct {
	Id int32 `json:"Id"`

	Version string `json:"Version"`

	ReleaseNotes string `json:"ReleaseNotes"`

	Deleted bool `json:"Deleted"`

	Latest bool `json:"Latest"`

	Pipeline Pipeline `json:"Pipeline"`
}

const PipelineVersionAvroCRC64Fingerprint = "\x94\x98m\xe9p\xb1\xaf\x93"

func NewPipelineVersion() PipelineVersion {
	r := PipelineVersion{}
	r.Pipeline = NewPipeline()

	return r
}

func DeserializePipelineVersion(r io.Reader) (PipelineVersion, error) {
	t := NewPipelineVersion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePipelineVersionFromSchema(r io.Reader, schema string) (PipelineVersion, error) {
	t := NewPipelineVersion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePipelineVersion(r PipelineVersion, w io.Writer) error {
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
	err = vm.WriteBool(r.Deleted, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Latest, w)
	if err != nil {
		return err
	}
	err = writePipeline(r.Pipeline, w)
	if err != nil {
		return err
	}
	return err
}

func (r PipelineVersion) Serialize(w io.Writer) error {
	return writePipelineVersion(r, w)
}

func (r PipelineVersion) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Version\",\"type\":\"string\"},{\"name\":\"ReleaseNotes\",\"type\":\"string\"},{\"name\":\"Deleted\",\"type\":\"boolean\"},{\"name\":\"Latest\",\"type\":\"boolean\"},{\"name\":\"Pipeline\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Description\",\"type\":\"string\"},{\"name\":\"RepositoryConfig\",\"type\":\"string\"},{\"name\":\"Deleted\",\"type\":\"boolean\"}],\"name\":\"Pipeline\",\"type\":\"record\"}}],\"name\":\"Andreani.WizardApi.Events.Record.PipelineVersion\",\"type\":\"record\"}"
}

func (r PipelineVersion) SchemaName() string {
	return "Andreani.WizardApi.Events.Record.PipelineVersion"
}

func (_ PipelineVersion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ PipelineVersion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ PipelineVersion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ PipelineVersion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ PipelineVersion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ PipelineVersion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ PipelineVersion) SetString(v string)   { panic("Unsupported operation") }
func (_ PipelineVersion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *PipelineVersion) Get(i int) types.Field {
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
		w := types.Boolean{Target: &r.Deleted}

		return w

	case 4:
		w := types.Boolean{Target: &r.Latest}

		return w

	case 5:
		r.Pipeline = NewPipeline()

		w := types.Record{Target: &r.Pipeline}

		return w

	}
	panic("Unknown field index")
}

func (r *PipelineVersion) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *PipelineVersion) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ PipelineVersion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ PipelineVersion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ PipelineVersion) HintSize(int)                     { panic("Unsupported operation") }
func (_ PipelineVersion) Finalize()                        {}

func (_ PipelineVersion) AvroCRC64Fingerprint() []byte {
	return []byte(PipelineVersionAvroCRC64Fingerprint)
}

func (r PipelineVersion) MarshalJSON() ([]byte, error) {
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
	output["Deleted"], err = json.Marshal(r.Deleted)
	if err != nil {
		return nil, err
	}
	output["Latest"], err = json.Marshal(r.Latest)
	if err != nil {
		return nil, err
	}
	output["Pipeline"], err = json.Marshal(r.Pipeline)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *PipelineVersion) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["Deleted"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Deleted); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Deleted")
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
	return nil
}
