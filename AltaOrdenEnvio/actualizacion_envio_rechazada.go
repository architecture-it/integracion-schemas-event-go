// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ActualizacionEnvioRechazada.avsc
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

type ActualizacionEnvioRechazada struct {
	GUID *UnionNullString `json:"GUID"`

	CodigoDeEnvio *UnionNullString `json:"CodigoDeEnvio"`

	Contrato *UnionNullString `json:"Contrato"`

	Cuando *UnionNullString `json:"Cuando"`
}

const ActualizacionEnvioRechazadaAvroCRC64Fingerprint = "\x16\x8d\xf6--\xa0o\xff"

func NewActualizacionEnvioRechazada() ActualizacionEnvioRechazada {
	r := ActualizacionEnvioRechazada{}
	r.Cuando = nil
	return r
}

func DeserializeActualizacionEnvioRechazada(r io.Reader) (ActualizacionEnvioRechazada, error) {
	t := NewActualizacionEnvioRechazada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeActualizacionEnvioRechazadaFromSchema(r io.Reader, schema string) (ActualizacionEnvioRechazada, error) {
	t := NewActualizacionEnvioRechazada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeActualizacionEnvioRechazada(r ActualizacionEnvioRechazada, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.GUID, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoDeEnvio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Contrato, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Cuando, w)
	if err != nil {
		return err
	}
	return err
}

func (r ActualizacionEnvioRechazada) Serialize(w io.Writer) error {
	return writeActualizacionEnvioRechazada(r, w)
}

func (r ActualizacionEnvioRechazada) Schema() string {
	return "{\"fields\":[{\"name\":\"GUID\",\"type\":[\"null\",\"string\"]},{\"name\":\"CodigoDeEnvio\",\"type\":[\"null\",\"string\"]},{\"name\":\"Contrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Cuando\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.AltaOrdenEnvio.Events.Record.ActualizacionEnvioRechazada\",\"type\":\"record\"}"
}

func (r ActualizacionEnvioRechazada) SchemaName() string {
	return "Andreani.AltaOrdenEnvio.Events.Record.ActualizacionEnvioRechazada"
}

func (_ ActualizacionEnvioRechazada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ActualizacionEnvioRechazada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ActualizacionEnvioRechazada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ActualizacionEnvioRechazada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ActualizacionEnvioRechazada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ActualizacionEnvioRechazada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ActualizacionEnvioRechazada) SetString(v string)   { panic("Unsupported operation") }
func (_ ActualizacionEnvioRechazada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ActualizacionEnvioRechazada) Get(i int) types.Field {
	switch i {
	case 0:
		r.GUID = NewUnionNullString()

		return r.GUID
	case 1:
		r.CodigoDeEnvio = NewUnionNullString()

		return r.CodigoDeEnvio
	case 2:
		r.Contrato = NewUnionNullString()

		return r.Contrato
	case 3:
		r.Cuando = NewUnionNullString()

		return r.Cuando
	}
	panic("Unknown field index")
}

func (r *ActualizacionEnvioRechazada) SetDefault(i int) {
	switch i {
	case 3:
		r.Cuando = nil
		return
	}
	panic("Unknown field index")
}

func (r *ActualizacionEnvioRechazada) NullField(i int) {
	switch i {
	case 0:
		r.GUID = nil
		return
	case 1:
		r.CodigoDeEnvio = nil
		return
	case 2:
		r.Contrato = nil
		return
	case 3:
		r.Cuando = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ActualizacionEnvioRechazada) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ActualizacionEnvioRechazada) AppendArray() types.Field { panic("Unsupported operation") }
func (_ ActualizacionEnvioRechazada) HintSize(int)             { panic("Unsupported operation") }
func (_ ActualizacionEnvioRechazada) Finalize()                {}

func (_ ActualizacionEnvioRechazada) AvroCRC64Fingerprint() []byte {
	return []byte(ActualizacionEnvioRechazadaAvroCRC64Fingerprint)
}

func (r ActualizacionEnvioRechazada) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["GUID"], err = json.Marshal(r.GUID)
	if err != nil {
		return nil, err
	}
	output["CodigoDeEnvio"], err = json.Marshal(r.CodigoDeEnvio)
	if err != nil {
		return nil, err
	}
	output["Contrato"], err = json.Marshal(r.Contrato)
	if err != nil {
		return nil, err
	}
	output["Cuando"], err = json.Marshal(r.Cuando)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ActualizacionEnvioRechazada) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["GUID"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.GUID); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for GUID")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoDeEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeEnvio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodigoDeEnvio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Contrato"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contrato); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Contrato")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Cuando"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cuando); err != nil {
			return err
		}
	} else {
		r.Cuando = NewUnionNullString()

		r.Cuando = nil
	}
	return nil
}
