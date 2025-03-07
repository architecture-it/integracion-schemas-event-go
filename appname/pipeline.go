// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Repository.avsc
 */
package appnameEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Pipeline struct {
	Id int32 `json:"Id"`

	Name string `json:"Name"`

	Description string `json:"Description"`

	RepositoryConfig string `json:"RepositoryConfig"`

	Deleted bool `json:"Deleted"`
}

const PipelineAvroCRC64Fingerprint = "E\x16\x13\xb2\xc8 \xa1\xb0"

func NewPipeline() Pipeline {
	r := Pipeline{}
	return r
}

func DeserializePipeline(r io.Reader) (Pipeline, error) {
	t := NewPipeline()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePipelineFromSchema(r io.Reader, schema string) (Pipeline, error) {
	t := NewPipeline()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePipeline(r Pipeline, w io.Writer) error {
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
	err = vm.WriteString(r.RepositoryConfig, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Deleted, w)
	if err != nil {
		return err
	}
	return err
}

func (r Pipeline) Serialize(w io.Writer) error {
	return writePipeline(r, w)
}

func (r Pipeline) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Description\",\"type\":\"string\"},{\"name\":\"RepositoryConfig\",\"type\":\"string\"},{\"name\":\"Deleted\",\"type\":\"boolean\"}],\"name\":\"Andreani.GithubIntegration.Events.Record.Pipeline\",\"type\":\"record\"}"
}

func (r Pipeline) SchemaName() string {
	return "Andreani.GithubIntegration.Events.Record.Pipeline"
}

func (_ Pipeline) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Pipeline) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Pipeline) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Pipeline) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Pipeline) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Pipeline) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Pipeline) SetString(v string)   { panic("Unsupported operation") }
func (_ Pipeline) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Pipeline) Get(i int) types.Field {
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
		w := types.String{Target: &r.RepositoryConfig}

		return w

	case 4:
		w := types.Boolean{Target: &r.Deleted}

		return w

	}
	panic("Unknown field index")
}

func (r *Pipeline) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Pipeline) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Pipeline) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Pipeline) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Pipeline) HintSize(int)                     { panic("Unsupported operation") }
func (_ Pipeline) Finalize()                        {}

func (_ Pipeline) AvroCRC64Fingerprint() []byte {
	return []byte(PipelineAvroCRC64Fingerprint)
}

func (r Pipeline) MarshalJSON() ([]byte, error) {
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
	output["RepositoryConfig"], err = json.Marshal(r.RepositoryConfig)
	if err != nil {
		return nil, err
	}
	output["Deleted"], err = json.Marshal(r.Deleted)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Pipeline) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["RepositoryConfig"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RepositoryConfig); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for RepositoryConfig")
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
	return nil
}
