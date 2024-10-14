// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Envio.avsc
 */
package MobileOperacionesProducerEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Envio struct {
	NumeroSeguimiento string `json:"numeroSeguimiento"`

	UnidadOperativa int32 `json:"unidadOperativa"`

	TareaId int32 `json:"tareaId"`

	MontoSolicitado float64 `json:"montoSolicitado"`
}

const EnvioAvroCRC64Fingerprint = "2\xdd`\x1a\xcf*0\xdf"

func NewEnvio() Envio {
	r := Envio{}
	return r
}

func DeserializeEnvio(r io.Reader) (Envio, error) {
	t := NewEnvio()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnvioFromSchema(r io.Reader, schema string) (Envio, error) {
	t := NewEnvio()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnvio(r Envio, w io.Writer) error {
	var err error
	err = vm.WriteString(r.NumeroSeguimiento, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.UnidadOperativa, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.TareaId, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.MontoSolicitado, w)
	if err != nil {
		return err
	}
	return err
}

func (r Envio) Serialize(w io.Writer) error {
	return writeEnvio(r, w)
}

func (r Envio) Schema() string {
	return "{\"fields\":[{\"name\":\"numeroSeguimiento\",\"type\":\"string\"},{\"name\":\"unidadOperativa\",\"type\":\"int\"},{\"name\":\"tareaId\",\"type\":\"int\"},{\"name\":\"montoSolicitado\",\"type\":\"double\"}],\"name\":\"Andreani.MobileOperacionesProducer.Events.Record.Envio\",\"type\":\"record\"}"
}

func (r Envio) SchemaName() string {
	return "Andreani.MobileOperacionesProducer.Events.Record.Envio"
}

func (_ Envio) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Envio) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Envio) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Envio) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Envio) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Envio) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Envio) SetString(v string)   { panic("Unsupported operation") }
func (_ Envio) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Envio) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.NumeroSeguimiento}

		return w

	case 1:
		w := types.Int{Target: &r.UnidadOperativa}

		return w

	case 2:
		w := types.Int{Target: &r.TareaId}

		return w

	case 3:
		w := types.Double{Target: &r.MontoSolicitado}

		return w

	}
	panic("Unknown field index")
}

func (r *Envio) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Envio) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Envio) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Envio) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Envio) HintSize(int)                     { panic("Unsupported operation") }
func (_ Envio) Finalize()                        {}

func (_ Envio) AvroCRC64Fingerprint() []byte {
	return []byte(EnvioAvroCRC64Fingerprint)
}

func (r Envio) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["numeroSeguimiento"], err = json.Marshal(r.NumeroSeguimiento)
	if err != nil {
		return nil, err
	}
	output["unidadOperativa"], err = json.Marshal(r.UnidadOperativa)
	if err != nil {
		return nil, err
	}
	output["tareaId"], err = json.Marshal(r.TareaId)
	if err != nil {
		return nil, err
	}
	output["montoSolicitado"], err = json.Marshal(r.MontoSolicitado)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Envio) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["numeroSeguimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroSeguimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroSeguimiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["unidadOperativa"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UnidadOperativa); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for unidadOperativa")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tareaId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TareaId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tareaId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["montoSolicitado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MontoSolicitado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for montoSolicitado")
	}
	return nil
}
