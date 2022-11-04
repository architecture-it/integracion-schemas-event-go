// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NumeracionAgrupadorBultosEvent.avsc
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

type NumeracionAgrupadorBultosEvent struct {
	Numero string `json:"Numero"`
}

const NumeracionAgrupadorBultosEventAvroCRC64Fingerprint = "G\x01[\xf1F\x9a<z"

func NewNumeracionAgrupadorBultosEvent() NumeracionAgrupadorBultosEvent {
	r := NumeracionAgrupadorBultosEvent{}
	return r
}

func DeserializeNumeracionAgrupadorBultosEvent(r io.Reader) (NumeracionAgrupadorBultosEvent, error) {
	t := NewNumeracionAgrupadorBultosEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNumeracionAgrupadorBultosEventFromSchema(r io.Reader, schema string) (NumeracionAgrupadorBultosEvent, error) {
	t := NewNumeracionAgrupadorBultosEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNumeracionAgrupadorBultosEvent(r NumeracionAgrupadorBultosEvent, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Numero, w)
	if err != nil {
		return err
	}
	return err
}

func (r NumeracionAgrupadorBultosEvent) Serialize(w io.Writer) error {
	return writeNumeracionAgrupadorBultosEvent(r, w)
}

func (r NumeracionAgrupadorBultosEvent) Schema() string {
	return "{\"fields\":[{\"name\":\"Numero\",\"type\":\"string\"}],\"name\":\"Andreani.AltaOrdenEnvio.Events.Record.NumeracionAgrupadorBultosEvent\",\"type\":\"record\"}"
}

func (r NumeracionAgrupadorBultosEvent) SchemaName() string {
	return "Andreani.AltaOrdenEnvio.Events.Record.NumeracionAgrupadorBultosEvent"
}

func (_ NumeracionAgrupadorBultosEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NumeracionAgrupadorBultosEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NumeracionAgrupadorBultosEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NumeracionAgrupadorBultosEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NumeracionAgrupadorBultosEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NumeracionAgrupadorBultosEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NumeracionAgrupadorBultosEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ NumeracionAgrupadorBultosEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NumeracionAgrupadorBultosEvent) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Numero}

		return w

	}
	panic("Unknown field index")
}

func (r *NumeracionAgrupadorBultosEvent) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NumeracionAgrupadorBultosEvent) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ NumeracionAgrupadorBultosEvent) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ NumeracionAgrupadorBultosEvent) AppendArray() types.Field { panic("Unsupported operation") }
func (_ NumeracionAgrupadorBultosEvent) HintSize(int)             { panic("Unsupported operation") }
func (_ NumeracionAgrupadorBultosEvent) Finalize()                {}

func (_ NumeracionAgrupadorBultosEvent) AvroCRC64Fingerprint() []byte {
	return []byte(NumeracionAgrupadorBultosEventAvroCRC64Fingerprint)
}

func (r NumeracionAgrupadorBultosEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Numero"], err = json.Marshal(r.Numero)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NumeracionAgrupadorBultosEvent) UnmarshalJSON(data []byte) error {
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
