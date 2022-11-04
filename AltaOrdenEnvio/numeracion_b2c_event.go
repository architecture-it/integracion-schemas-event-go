// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NumeracionB2CEvent.avsc
 */
package AltaOrdenEnvioEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type NumeracionB2CEvent struct {
	Numero string `json:"Numero"`
}

const NumeracionB2CEventAvroCRC64Fingerprint = "^T\x1c\x1c\xcdש\xf6"

func NewNumeracionB2CEvent() NumeracionB2CEvent {
	r := NumeracionB2CEvent{}
	return r
}

func DeserializeNumeracionB2CEvent(r io.Reader) (NumeracionB2CEvent, error) {
	t := NewNumeracionB2CEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNumeracionB2CEventFromSchema(r io.Reader, schema string) (NumeracionB2CEvent, error) {
	t := NewNumeracionB2CEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNumeracionB2CEvent(r NumeracionB2CEvent, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Numero, w)
	if err != nil {
		return err
	}
	return err
}

func (r NumeracionB2CEvent) Serialize(w io.Writer) error {
	return writeNumeracionB2CEvent(r, w)
}

func (r NumeracionB2CEvent) Schema() string {
	return "{\"fields\":[{\"name\":\"Numero\",\"type\":\"string\"}],\"name\":\"Andreani.AltaOrdenEnvio.Events.Record.NumeracionB2CEvent\",\"type\":\"record\"}"
}

func (r NumeracionB2CEvent) SchemaName() string {
	return "Andreani.AltaOrdenEnvio.Events.Record.NumeracionB2CEvent"
}

func (_ NumeracionB2CEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NumeracionB2CEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NumeracionB2CEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NumeracionB2CEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NumeracionB2CEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NumeracionB2CEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NumeracionB2CEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ NumeracionB2CEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NumeracionB2CEvent) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Numero}

		return w

	}
	panic("Unknown field index")
}

func (r *NumeracionB2CEvent) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NumeracionB2CEvent) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ NumeracionB2CEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NumeracionB2CEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NumeracionB2CEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ NumeracionB2CEvent) Finalize()                        {}

func (_ NumeracionB2CEvent) AvroCRC64Fingerprint() []byte {
	return []byte(NumeracionB2CEventAvroCRC64Fingerprint)
}

func (r NumeracionB2CEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Numero"], err = json.Marshal(r.Numero)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NumeracionB2CEvent) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Numero"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Numero); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Numero")
	}
	return nil
}