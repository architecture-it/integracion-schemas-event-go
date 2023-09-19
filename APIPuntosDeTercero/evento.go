// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Evento.avsc
 */
package APIPuntosDeTerceroEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Evento struct {
	PuntoDeTercero *UnionNullPuntoDeTercero `json:"puntoDeTercero"`

	Estado *UnionNullString `json:"estado"`
}

const EventoAvroCRC64Fingerprint = "I\xc3\x00J\xfa\xa15\xb3"

func NewEvento() Evento {
	r := Evento{}
	return r
}

func DeserializeEvento(r io.Reader) (Evento, error) {
	t := NewEvento()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoFromSchema(r io.Reader, schema string) (Evento, error) {
	t := NewEvento()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEvento(r Evento, w io.Writer) error {
	var err error
	err = writeUnionNullPuntoDeTercero(r.PuntoDeTercero, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Estado, w)
	if err != nil {
		return err
	}
	return err
}

func (r Evento) Serialize(w io.Writer) error {
	return writeEvento(r, w)
}

func (r Evento) Schema() string {
	return "{\"fields\":[{\"name\":\"puntoDeTercero\",\"type\":[\"null\",{\"fields\":[{\"name\":\"id\",\"type\":[\"null\",\"int\"]},{\"name\":\"nombre\",\"type\":\"string\"},{\"name\":\"telefono\",\"type\":[\"null\",\"string\"]},{\"name\":\"horarioDeAtencion\",\"type\":\"string\"},{\"name\":\"observaciones\",\"type\":[\"null\",\"string\"]},{\"name\":\"admiteEnvios\",\"type\":[\"null\",\"int\"]},{\"name\":\"entregaEnvios\",\"type\":[\"null\",\"int\"]},{\"name\":\"tipo\",\"type\":\"string\"},{\"name\":\"referencia\",\"type\":[\"null\",\"string\"]},{\"name\":\"responsable\",\"type\":[\"null\",{\"fields\":[{\"name\":\"nombre\",\"type\":\"string\"},{\"name\":\"apellido\",\"type\":[\"null\",\"string\"]},{\"name\":\"mail\",\"type\":\"string\"}],\"name\":\"Responsable\",\"type\":\"record\"}]},{\"name\":\"ubicacion\",\"type\":[\"null\",{\"fields\":[{\"name\":\"calle\",\"type\":\"string\"},{\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"name\":\"codigoPostal\",\"type\":\"string\"},{\"name\":\"latitud\",\"type\":\"string\"},{\"name\":\"longitud\",\"type\":\"string\"},{\"name\":\"region\",\"type\":[\"null\",\"string\"]}],\"name\":\"Ubicacion\",\"type\":\"record\"}]},{\"name\":\"activo\",\"type\":[\"null\",\"boolean\"]}],\"name\":\"PuntoDeTercero\",\"namespace\":\"Andreani.PuntoDeTercero.Events.Common\",\"type\":\"record\"}]},{\"name\":\"estado\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.PuntoDeTercero.Events.Record.Evento\",\"type\":\"record\"}"
}

func (r Evento) SchemaName() string {
	return "Andreani.PuntoDeTercero.Events.Record.Evento"
}

func (_ Evento) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Evento) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Evento) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Evento) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Evento) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Evento) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Evento) SetString(v string)   { panic("Unsupported operation") }
func (_ Evento) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Evento) Get(i int) types.Field {
	switch i {
	case 0:
		r.PuntoDeTercero = NewUnionNullPuntoDeTercero()

		return r.PuntoDeTercero
	case 1:
		r.Estado = NewUnionNullString()

		return r.Estado
	}
	panic("Unknown field index")
}

func (r *Evento) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Evento) NullField(i int) {
	switch i {
	case 0:
		r.PuntoDeTercero = nil
		return
	case 1:
		r.Estado = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Evento) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Evento) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Evento) HintSize(int)                     { panic("Unsupported operation") }
func (_ Evento) Finalize()                        {}

func (_ Evento) AvroCRC64Fingerprint() []byte {
	return []byte(EventoAvroCRC64Fingerprint)
}

func (r Evento) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["puntoDeTercero"], err = json.Marshal(r.PuntoDeTercero)
	if err != nil {
		return nil, err
	}
	output["estado"], err = json.Marshal(r.Estado)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Evento) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["puntoDeTercero"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PuntoDeTercero); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for puntoDeTercero")
	}
	val = func() json.RawMessage {
		if v, ok := fields["estado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for estado")
	}
	return nil
}
