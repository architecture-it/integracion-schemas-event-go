// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnviosCotizarEvent.avsc
 */
package CorporativoBackendEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EnviosCotizarEvent struct {
	Envios []Envio `json:"Envios"`
}

const EnviosCotizarEventAvroCRC64Fingerprint = "\xea\x11\x8c!QK\x9aU"

func NewEnviosCotizarEvent() EnviosCotizarEvent {
	r := EnviosCotizarEvent{}
	r.Envios = make([]Envio, 0)

	return r
}

func DeserializeEnviosCotizarEvent(r io.Reader) (EnviosCotizarEvent, error) {
	t := NewEnviosCotizarEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnviosCotizarEventFromSchema(r io.Reader, schema string) (EnviosCotizarEvent, error) {
	t := NewEnviosCotizarEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnviosCotizarEvent(r EnviosCotizarEvent, w io.Writer) error {
	var err error
	err = writeArrayEnvio(r.Envios, w)
	if err != nil {
		return err
	}
	return err
}

func (r EnviosCotizarEvent) Serialize(w io.Writer) error {
	return writeEnviosCotizarEvent(r, w)
}

func (r EnviosCotizarEvent) Schema() string {
	return "{\"fields\":[{\"name\":\"Envios\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"ContratoId\",\"type\":\"string\"},{\"name\":\"CodigoPostalOrigenId\",\"type\":\"string\"},{\"name\":\"CodigoPostalDestinoId\",\"type\":\"string\"},{\"name\":\"Peso\",\"type\":\"float\"},{\"name\":\"Volumen\",\"type\":\"float\"},{\"name\":\"ValorDeclarado\",\"type\":\"float\"},{\"name\":\"AltoCm\",\"type\":\"float\"},{\"name\":\"AnchoCm\",\"type\":\"float\"},{\"name\":\"LargoCm\",\"type\":\"float\"}],\"name\":\"Envio\",\"namespace\":\"Andreani.CorporativoBackend.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.CorporativoBackend.Events.Record.EnviosCotizarEvent\",\"type\":\"record\"}"
}

func (r EnviosCotizarEvent) SchemaName() string {
	return "Andreani.CorporativoBackend.Events.Record.EnviosCotizarEvent"
}

func (_ EnviosCotizarEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EnviosCotizarEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EnviosCotizarEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EnviosCotizarEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EnviosCotizarEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EnviosCotizarEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EnviosCotizarEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ EnviosCotizarEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EnviosCotizarEvent) Get(i int) types.Field {
	switch i {
	case 0:
		r.Envios = make([]Envio, 0)

		w := ArrayEnvioWrapper{Target: &r.Envios}

		return w

	}
	panic("Unknown field index")
}

func (r *EnviosCotizarEvent) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EnviosCotizarEvent) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EnviosCotizarEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EnviosCotizarEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EnviosCotizarEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ EnviosCotizarEvent) Finalize()                        {}

func (_ EnviosCotizarEvent) AvroCRC64Fingerprint() []byte {
	return []byte(EnviosCotizarEventAvroCRC64Fingerprint)
}

func (r EnviosCotizarEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Envios"], err = json.Marshal(r.Envios)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EnviosCotizarEvent) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Envios"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Envios); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Envios")
	}
	return nil
}
