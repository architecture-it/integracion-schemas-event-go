// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ListInterface282.avsc
 */
package HCMInterface282Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Interface282Data struct {
	Legajo string `json:"Legajo"`

	FechaAcuerdo string `json:"FechaAcuerdo"`

	SueldoJournal string `json:"SueldoJournal"`

	ZonaBandaSalarial string `json:"ZonaBandaSalarial"`
}

const Interface282DataAvroCRC64Fingerprint = "R\xb8\x18\x88\xbd\x8a\x06\x94"

func NewInterface282Data() Interface282Data {
	r := Interface282Data{}
	return r
}

func DeserializeInterface282Data(r io.Reader) (Interface282Data, error) {
	t := NewInterface282Data()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeInterface282DataFromSchema(r io.Reader, schema string) (Interface282Data, error) {
	t := NewInterface282Data()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeInterface282Data(r Interface282Data, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Legajo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaAcuerdo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.SueldoJournal, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.ZonaBandaSalarial, w)
	if err != nil {
		return err
	}
	return err
}

func (r Interface282Data) Serialize(w io.Writer) error {
	return writeInterface282Data(r, w)
}

func (r Interface282Data) Schema() string {
	return "{\"fields\":[{\"name\":\"Legajo\",\"type\":\"string\"},{\"name\":\"FechaAcuerdo\",\"type\":\"string\"},{\"name\":\"SueldoJournal\",\"type\":\"string\"},{\"name\":\"ZonaBandaSalarial\",\"type\":\"string\"}],\"name\":\"Andreani.HCMInterface282.Events.Record.Interface282Data\",\"type\":\"record\"}"
}

func (r Interface282Data) SchemaName() string {
	return "Andreani.HCMInterface282.Events.Record.Interface282Data"
}

func (_ Interface282Data) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Interface282Data) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Interface282Data) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Interface282Data) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Interface282Data) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Interface282Data) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Interface282Data) SetString(v string)   { panic("Unsupported operation") }
func (_ Interface282Data) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Interface282Data) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Legajo}

		return w

	case 1:
		w := types.String{Target: &r.FechaAcuerdo}

		return w

	case 2:
		w := types.String{Target: &r.SueldoJournal}

		return w

	case 3:
		w := types.String{Target: &r.ZonaBandaSalarial}

		return w

	}
	panic("Unknown field index")
}

func (r *Interface282Data) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Interface282Data) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Interface282Data) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Interface282Data) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Interface282Data) HintSize(int)                     { panic("Unsupported operation") }
func (_ Interface282Data) Finalize()                        {}

func (_ Interface282Data) AvroCRC64Fingerprint() []byte {
	return []byte(Interface282DataAvroCRC64Fingerprint)
}

func (r Interface282Data) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Legajo"], err = json.Marshal(r.Legajo)
	if err != nil {
		return nil, err
	}
	output["FechaAcuerdo"], err = json.Marshal(r.FechaAcuerdo)
	if err != nil {
		return nil, err
	}
	output["SueldoJournal"], err = json.Marshal(r.SueldoJournal)
	if err != nil {
		return nil, err
	}
	output["ZonaBandaSalarial"], err = json.Marshal(r.ZonaBandaSalarial)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Interface282Data) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Legajo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Legajo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Legajo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaAcuerdo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaAcuerdo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaAcuerdo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SueldoJournal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SueldoJournal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SueldoJournal")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ZonaBandaSalarial"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ZonaBandaSalarial); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ZonaBandaSalarial")
	}
	return nil
}
