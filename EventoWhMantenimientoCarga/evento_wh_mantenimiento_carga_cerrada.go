// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhMantenimientoCargaCerrada.avsc
 */
package EventoWhMantenimientoCargaEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EventoWhMantenimientoCargaCerrada struct {
	Identificacion Identificacion `json:"Identificacion"`

	Cabecera Cabecera `json:"Cabecera"`

	Detalle []Detalle `json:"Detalle"`
}

const EventoWhMantenimientoCargaCerradaAvroCRC64Fingerprint = "\x018\"!\x05c\xbd\xbf"

func NewEventoWhMantenimientoCargaCerrada() EventoWhMantenimientoCargaCerrada {
	r := EventoWhMantenimientoCargaCerrada{}
	r.Identificacion = NewIdentificacion()

	r.Cabecera = NewCabecera()

	r.Detalle = make([]Detalle, 0)

	return r
}

func DeserializeEventoWhMantenimientoCargaCerrada(r io.Reader) (EventoWhMantenimientoCargaCerrada, error) {
	t := NewEventoWhMantenimientoCargaCerrada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoWhMantenimientoCargaCerradaFromSchema(r io.Reader, schema string) (EventoWhMantenimientoCargaCerrada, error) {
	t := NewEventoWhMantenimientoCargaCerrada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoWhMantenimientoCargaCerrada(r EventoWhMantenimientoCargaCerrada, w io.Writer) error {
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

func (r EventoWhMantenimientoCargaCerrada) Serialize(w io.Writer) error {
	return writeEventoWhMantenimientoCargaCerrada(r, w)
}

func (r EventoWhMantenimientoCargaCerrada) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.EventoWhMantenimientoCarga.Events.CerradaCommon\",\"type\":\"record\"}},{\"name\":\"Cabecera\",\"type\":{\"fields\":[],\"name\":\"Cabecera\",\"namespace\":\"Andreani.EventoWhMantenimientoCarga.Events.CerradaCommon\",\"type\":\"record\"}},{\"name\":\"Detalle\",\"type\":{\"items\":{\"fields\":[],\"name\":\"Detalle\",\"namespace\":\"Andreani.EventoWhMantenimientoCarga.Events.CerradaCommon\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.EventoWhMantenimientoCarga.Events.Record.EventoWhMantenimientoCargaCerrada\",\"type\":\"record\"}"
}

func (r EventoWhMantenimientoCargaCerrada) SchemaName() string {
	return "Andreani.EventoWhMantenimientoCarga.Events.Record.EventoWhMantenimientoCargaCerrada"
}

func (_ EventoWhMantenimientoCargaCerrada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoWhMantenimientoCargaCerrada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoWhMantenimientoCargaCerrada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoWhMantenimientoCargaCerrada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoWhMantenimientoCargaCerrada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoWhMantenimientoCargaCerrada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoWhMantenimientoCargaCerrada) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoWhMantenimientoCargaCerrada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoWhMantenimientoCargaCerrada) Get(i int) types.Field {
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

func (r *EventoWhMantenimientoCargaCerrada) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventoWhMantenimientoCargaCerrada) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EventoWhMantenimientoCargaCerrada) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ EventoWhMantenimientoCargaCerrada) AppendArray() types.Field { panic("Unsupported operation") }
func (_ EventoWhMantenimientoCargaCerrada) HintSize(int)             { panic("Unsupported operation") }
func (_ EventoWhMantenimientoCargaCerrada) Finalize()                {}

func (_ EventoWhMantenimientoCargaCerrada) AvroCRC64Fingerprint() []byte {
	return []byte(EventoWhMantenimientoCargaCerradaAvroCRC64Fingerprint)
}

func (r EventoWhMantenimientoCargaCerrada) MarshalJSON() ([]byte, error) {
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

func (r *EventoWhMantenimientoCargaCerrada) UnmarshalJSON(data []byte) error {
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
