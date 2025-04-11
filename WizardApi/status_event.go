// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TemplateEvent.avsc
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

type StatusEvent struct {
	Id int32 `json:"Id"`

	Description string `json:"Description"`
}

const StatusEventAvroCRC64Fingerprint = "-2\xa2\xcd\xe2\xde\xec\x93"

func NewStatusEvent() StatusEvent {
	r := StatusEvent{}
	return r
}

func DeserializeStatusEvent(r io.Reader) (StatusEvent, error) {
	t := NewStatusEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeStatusEventFromSchema(r io.Reader, schema string) (StatusEvent, error) {
	t := NewStatusEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeStatusEvent(r StatusEvent, w io.Writer) error {
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

func (r StatusEvent) Serialize(w io.Writer) error {
	return writeStatusEvent(r, w)
}

func (r StatusEvent) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"}],\"name\":\"Andreani.WizardApi.Events.Common.StatusEvent\",\"type\":\"record\"}"
}

func (r StatusEvent) SchemaName() string {
	return "Andreani.WizardApi.Events.Common.StatusEvent"
}

func (_ StatusEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ StatusEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ StatusEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ StatusEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ StatusEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ StatusEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ StatusEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ StatusEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *StatusEvent) Get(i int) types.Field {
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

func (r *StatusEvent) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *StatusEvent) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ StatusEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ StatusEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ StatusEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ StatusEvent) Finalize()                        {}

func (_ StatusEvent) AvroCRC64Fingerprint() []byte {
	return []byte(StatusEventAvroCRC64Fingerprint)
}

func (r StatusEvent) MarshalJSON() ([]byte, error) {
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

func (r *StatusEvent) UnmarshalJSON(data []byte) error {
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
