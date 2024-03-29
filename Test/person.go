// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Person.avsc
 */
package TestEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Person struct {
	Name string `json:"name"`

	Surname string `json:"surname"`

	Seniority string `json:"seniority"`

	OnSite bool `json:"onSite"`

	Team *UnionNullTeam `json:"team"`

	Age int32 `json:"age"`

	Direccion *UnionNullString `json:"direccion"`
}

const PersonAvroCRC64Fingerprint = "\xe0\xb9J\x15\x98>#<"

func NewPerson() Person {
	r := Person{}
	r.Team = nil
	r.Direccion = nil
	return r
}

func DeserializePerson(r io.Reader) (Person, error) {
	t := NewPerson()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePersonFromSchema(r io.Reader, schema string) (Person, error) {
	t := NewPerson()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePerson(r Person, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Surname, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Seniority, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.OnSite, w)
	if err != nil {
		return err
	}
	err = writeUnionNullTeam(r.Team, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Age, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Direccion, w)
	if err != nil {
		return err
	}
	return err
}

func (r Person) Serialize(w io.Writer) error {
	return writePerson(r, w)
}

func (r Person) Schema() string {
	return "{\"fields\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"surname\",\"type\":\"string\"},{\"name\":\"seniority\",\"type\":\"string\"},{\"name\":\"onSite\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"team\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"tl\",\"type\":[\"null\",\"string\"]},{\"name\":\"boss\",\"type\":\"string\"}],\"name\":\"Team\",\"type\":\"record\"}]},{\"name\":\"age\",\"type\":\"int\"},{\"default\":null,\"name\":\"direccion\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.Test.Events.Record.Common.Person\",\"type\":\"record\"}"
}

func (r Person) SchemaName() string {
	return "Andreani.Test.Events.Record.Common.Person"
}

func (_ Person) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Person) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Person) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Person) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Person) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Person) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Person) SetString(v string)   { panic("Unsupported operation") }
func (_ Person) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Person) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Name}

		return w

	case 1:
		w := types.String{Target: &r.Surname}

		return w

	case 2:
		w := types.String{Target: &r.Seniority}

		return w

	case 3:
		w := types.Boolean{Target: &r.OnSite}

		return w

	case 4:
		r.Team = NewUnionNullTeam()

		return r.Team
	case 5:
		w := types.Int{Target: &r.Age}

		return w

	case 6:
		r.Direccion = NewUnionNullString()

		return r.Direccion
	}
	panic("Unknown field index")
}

func (r *Person) SetDefault(i int) {
	switch i {
	case 4:
		r.Team = nil
		return
	case 6:
		r.Direccion = nil
		return
	}
	panic("Unknown field index")
}

func (r *Person) NullField(i int) {
	switch i {
	case 4:
		r.Team = nil
		return
	case 6:
		r.Direccion = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Person) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Person) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Person) HintSize(int)                     { panic("Unsupported operation") }
func (_ Person) Finalize()                        {}

func (_ Person) AvroCRC64Fingerprint() []byte {
	return []byte(PersonAvroCRC64Fingerprint)
}

func (r Person) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	output["surname"], err = json.Marshal(r.Surname)
	if err != nil {
		return nil, err
	}
	output["seniority"], err = json.Marshal(r.Seniority)
	if err != nil {
		return nil, err
	}
	output["onSite"], err = json.Marshal(r.OnSite)
	if err != nil {
		return nil, err
	}
	output["team"], err = json.Marshal(r.Team)
	if err != nil {
		return nil, err
	}
	output["age"], err = json.Marshal(r.Age)
	if err != nil {
		return nil, err
	}
	output["direccion"], err = json.Marshal(r.Direccion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Person) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["surname"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Surname); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for surname")
	}
	val = func() json.RawMessage {
		if v, ok := fields["seniority"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Seniority); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for seniority")
	}
	val = func() json.RawMessage {
		if v, ok := fields["onSite"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OnSite); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for onSite")
	}
	val = func() json.RawMessage {
		if v, ok := fields["team"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Team); err != nil {
			return err
		}
	} else {
		r.Team = NewUnionNullTeam()

		r.Team = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["age"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Age); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for age")
	}
	val = func() json.RawMessage {
		if v, ok := fields["direccion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Direccion); err != nil {
			return err
		}
	} else {
		r.Direccion = NewUnionNullString()

		r.Direccion = nil
	}
	return nil
}
