// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     WorkflowCiEvent.avsc
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

type WorkflowCiEvent struct {
	Id int32 `json:"Id"`

	Description string `json:"Description"`
}

const WorkflowCiEventAvroCRC64Fingerprint = "Rg\xde\x19\xe44\x01N"

func NewWorkflowCiEvent() WorkflowCiEvent {
	r := WorkflowCiEvent{}
	return r
}

func DeserializeWorkflowCiEvent(r io.Reader) (WorkflowCiEvent, error) {
	t := NewWorkflowCiEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeWorkflowCiEventFromSchema(r io.Reader, schema string) (WorkflowCiEvent, error) {
	t := NewWorkflowCiEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeWorkflowCiEvent(r WorkflowCiEvent, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Description, w)
	if err != nil {
		return err
	}
	return err
}

func (r WorkflowCiEvent) Serialize(w io.Writer) error {
	return writeWorkflowCiEvent(r, w)
}

func (r WorkflowCiEvent) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"}],\"name\":\"Andreani.WizardApi.Events.Record.WorkflowCiEvent\",\"type\":\"record\"}"
}

func (r WorkflowCiEvent) SchemaName() string {
	return "Andreani.WizardApi.Events.Record.WorkflowCiEvent"
}

func (_ WorkflowCiEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ WorkflowCiEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ WorkflowCiEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ WorkflowCiEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ WorkflowCiEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ WorkflowCiEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ WorkflowCiEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ WorkflowCiEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *WorkflowCiEvent) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Description}

		return w

	}
	panic("Unknown field index")
}

func (r *WorkflowCiEvent) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *WorkflowCiEvent) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ WorkflowCiEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ WorkflowCiEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ WorkflowCiEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ WorkflowCiEvent) Finalize()                        {}

func (_ WorkflowCiEvent) AvroCRC64Fingerprint() []byte {
	return []byte(WorkflowCiEventAvroCRC64Fingerprint)
}

func (r WorkflowCiEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["Description"], err = json.Marshal(r.Description)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *WorkflowCiEvent) UnmarshalJSON(data []byte) error {
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
	return nil
}
