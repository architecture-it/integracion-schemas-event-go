// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NumeracionB2BEvent.avsc
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

type NumeracionB2BEvent struct {
	Numero string `json:"Numero"`
}

const NumeracionB2BEventAvroCRC64Fingerprint = "\x03<\x9dc\xb6@Z\xfc"

func NewNumeracionB2BEvent() NumeracionB2BEvent {
	r := NumeracionB2BEvent{}
	return r
}

func DeserializeNumeracionB2BEvent(r io.Reader) (NumeracionB2BEvent, error) {
	t := NewNumeracionB2BEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNumeracionB2BEventFromSchema(r io.Reader, schema string) (NumeracionB2BEvent, error) {
	t := NewNumeracionB2BEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNumeracionB2BEvent(r NumeracionB2BEvent, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Numero, w)
	if err != nil {
		return err
	}
	return err
}

func (r NumeracionB2BEvent) Serialize(w io.Writer) error {
	return writeNumeracionB2BEvent(r, w)
}

func (r NumeracionB2BEvent) Schema() string {
	return "{\"fields\":[{\"name\":\"Numero\",\"type\":\"string\"}],\"name\":\"Andreani.AltaOrdenEnvio.Events.Record.NumeracionB2BEvent\",\"type\":\"record\"}"
}

func (r NumeracionB2BEvent) SchemaName() string {
	return "Andreani.AltaOrdenEnvio.Events.Record.NumeracionB2BEvent"
}

func (_ NumeracionB2BEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NumeracionB2BEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NumeracionB2BEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NumeracionB2BEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NumeracionB2BEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NumeracionB2BEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NumeracionB2BEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ NumeracionB2BEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NumeracionB2BEvent) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Numero}

		return w

	}
	panic("Unknown field index")
}

func (r *NumeracionB2BEvent) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NumeracionB2BEvent) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ NumeracionB2BEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NumeracionB2BEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NumeracionB2BEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ NumeracionB2BEvent) Finalize()                        {}

func (_ NumeracionB2BEvent) AvroCRC64Fingerprint() []byte {
	return []byte(NumeracionB2BEventAvroCRC64Fingerprint)
}

func (r NumeracionB2BEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Numero"], err = json.Marshal(r.Numero)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NumeracionB2BEvent) UnmarshalJSON(data []byte) error {
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
