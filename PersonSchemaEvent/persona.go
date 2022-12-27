// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Persona.avsc
 */
package PersonSchemaEventEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Persona struct {
	Name string `json:"Name"`

	Lastname string `json:"Lastname"`

	Number int32 `json:"number"`
}

const PersonaAvroCRC64Fingerprint = "\xa9OPL`\x9b\x8e\xa9"

func NewPersona() Persona {
	r := Persona{}
	return r
}

func DeserializePersona(r io.Reader) (Persona, error) {
	t := NewPersona()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePersonaFromSchema(r io.Reader, schema string) (Persona, error) {
	t := NewPersona()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePersona(r Persona, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Lastname, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Number, w)
	if err != nil {
		return err
	}
	return err
}

func (r Persona) Serialize(w io.Writer) error {
	return writePersona(r, w)
}

func (r Persona) Schema() string {
	return "{\"fields\":[{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Lastname\",\"type\":\"string\"},{\"name\":\"number\",\"type\":\"int\"}],\"name\":\"Andreani.PersonSchemaEvent.Record.Persona\",\"type\":\"record\"}"
}

func (r Persona) SchemaName() string {
	return "Andreani.PersonSchemaEvent.Record.Persona"
}

func (_ Persona) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Persona) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Persona) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Persona) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Persona) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Persona) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Persona) SetString(v string)   { panic("Unsupported operation") }
func (_ Persona) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Persona) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Name}

		return w

	case 1:
		w := types.String{Target: &r.Lastname}

		return w

	case 2:
		w := types.Int{Target: &r.Number}

		return w

	}
	panic("Unknown field index")
}

func (r *Persona) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Persona) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Persona) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Persona) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Persona) HintSize(int)                     { panic("Unsupported operation") }
func (_ Persona) Finalize()                        {}

func (_ Persona) AvroCRC64Fingerprint() []byte {
	return []byte(PersonaAvroCRC64Fingerprint)
}

func (r Persona) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	output["Lastname"], err = json.Marshal(r.Lastname)
	if err != nil {
		return nil, err
	}
	output["number"], err = json.Marshal(r.Number)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Persona) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
		if v, ok := fields["Lastname"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Lastname); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Lastname")
	}
	val = func() json.RawMessage {
		if v, ok := fields["number"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Number); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for number")
	}
	return nil
}