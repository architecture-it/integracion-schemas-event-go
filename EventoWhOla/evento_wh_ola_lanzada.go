// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhOlaLanzada.avsc
 */
package EventoWhOlaEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EventoWhOlaLanzada struct {
	Identificacion Identificacion `json:"Identificacion"`

	Cabecera Cabecera `json:"Cabecera"`

	Detalle []Detalle `json:"Detalle"`
}

const EventoWhOlaLanzadaAvroCRC64Fingerprint = "yD\x1a\xfb\v\x84\x04\x1c"

func NewEventoWhOlaLanzada() EventoWhOlaLanzada {
	r := EventoWhOlaLanzada{}
	r.Identificacion = NewIdentificacion()

	r.Cabecera = NewCabecera()

	r.Detalle = make([]Detalle, 0)

	return r
}

func DeserializeEventoWhOlaLanzada(r io.Reader) (EventoWhOlaLanzada, error) {
	t := NewEventoWhOlaLanzada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoWhOlaLanzadaFromSchema(r io.Reader, schema string) (EventoWhOlaLanzada, error) {
	t := NewEventoWhOlaLanzada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoWhOlaLanzada(r EventoWhOlaLanzada, w io.Writer) error {
	var err error
	err = writeIdentificacion(r.Identificacion, w)
	if err != nil {
		return err
	}
	err = writeCabecera(r.Cabecera, w)
	if err != nil {
		return err
	}
	err = writeArrayDetalle(r.Detalle, w)
	if err != nil {
		return err
	}
	return err
}

func (r EventoWhOlaLanzada) Serialize(w io.Writer) error {
	return writeEventoWhOlaLanzada(r, w)
}

func (r EventoWhOlaLanzada) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.EventoWhOla.Events.LanzadaCommon\",\"type\":\"record\"}},{\"name\":\"Cabecera\",\"type\":{\"fields\":[{\"name\":\"Ola\",\"type\":\"string\"}],\"name\":\"Cabecera\",\"namespace\":\"Andreani.EventoWhOla.Events.LanzadaCommon\",\"type\":\"record\"}},{\"name\":\"Detalle\",\"type\":{\"items\":{\"fields\":[{\"name\":\"OrdenSCE\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Ola\",\"type\":\"string\"}],\"name\":\"Detalle\",\"namespace\":\"Andreani.EventoWhOla.Events.LanzadaCommon\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.EventoWhOla.Events.Record.EventoWhOlaLanzada\",\"type\":\"record\"}"
}

func (r EventoWhOlaLanzada) SchemaName() string {
	return "Andreani.EventoWhOla.Events.Record.EventoWhOlaLanzada"
}

func (_ EventoWhOlaLanzada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoWhOlaLanzada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoWhOlaLanzada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoWhOlaLanzada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoWhOlaLanzada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoWhOlaLanzada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoWhOlaLanzada) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoWhOlaLanzada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoWhOlaLanzada) Get(i int) types.Field {
	switch i {
	case 0:
		r.Identificacion = NewIdentificacion()

		w := types.Record{Target: &r.Identificacion}

		return w

	case 1:
		r.Cabecera = NewCabecera()

		w := types.Record{Target: &r.Cabecera}

		return w

	case 2:
		r.Detalle = make([]Detalle, 0)

		w := ArrayDetalleWrapper{Target: &r.Detalle}

		return w

	}
	panic("Unknown field index")
}

func (r *EventoWhOlaLanzada) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventoWhOlaLanzada) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EventoWhOlaLanzada) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EventoWhOlaLanzada) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EventoWhOlaLanzada) HintSize(int)                     { panic("Unsupported operation") }
func (_ EventoWhOlaLanzada) Finalize()                        {}

func (_ EventoWhOlaLanzada) AvroCRC64Fingerprint() []byte {
	return []byte(EventoWhOlaLanzadaAvroCRC64Fingerprint)
}

func (r EventoWhOlaLanzada) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Identificacion"], err = json.Marshal(r.Identificacion)
	if err != nil {
		return nil, err
	}
	output["Cabecera"], err = json.Marshal(r.Cabecera)
	if err != nil {
		return nil, err
	}
	output["Detalle"], err = json.Marshal(r.Detalle)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EventoWhOlaLanzada) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Identificacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Identificacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Identificacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Cabecera"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cabecera); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Cabecera")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Detalle"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Detalle); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Detalle")
	}
	return nil
}