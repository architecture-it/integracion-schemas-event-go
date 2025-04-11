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

type LanguageEvent struct {
	Id int32 `json:"Id"`

	Name string `json:"Name"`
}

const LanguageEventAvroCRC64Fingerprint = "\xa61\xd7.\xbcZ\xcc\xeb"

func NewLanguageEvent() LanguageEvent {
	r := LanguageEvent{}
	return r
}

func DeserializeLanguageEvent(r io.Reader) (LanguageEvent, error) {
	t := NewLanguageEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeLanguageEventFromSchema(r io.Reader, schema string) (LanguageEvent, error) {
	t := NewLanguageEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeLanguageEvent(r LanguageEvent, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	return err
}

func (r LanguageEvent) Serialize(w io.Writer) error {
	return writeLanguageEvent(r, w)
}

func (r LanguageEvent) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Name\",\"type\":\"string\"}],\"name\":\"Andreani.WizardApi.Events.Record.LanguageEvent\",\"type\":\"record\"}"
}

func (r LanguageEvent) SchemaName() string {
	return "Andreani.WizardApi.Events.Record.LanguageEvent"
}

func (_ LanguageEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ LanguageEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ LanguageEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ LanguageEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ LanguageEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ LanguageEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ LanguageEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ LanguageEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *LanguageEvent) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Name}

		return w

	}
	panic("Unknown field index")
}

func (r *LanguageEvent) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *LanguageEvent) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ LanguageEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ LanguageEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ LanguageEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ LanguageEvent) Finalize()                        {}

func (_ LanguageEvent) AvroCRC64Fingerprint() []byte {
	return []byte(LanguageEventAvroCRC64Fingerprint)
}

func (r LanguageEvent) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(output)
}

func (r *LanguageEvent) UnmarshalJSON(data []byte) error {
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
	return nil
}
