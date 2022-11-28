// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhPosicionesAlta.avsc
 */
package EventoWhPosicionesEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EventoWhPosicionesAlta struct {
	Identificacion Identificacion `json:"Identificacion"`

	Cabecera Cabecera `json:"Cabecera"`
}

const EventoWhPosicionesAltaAvroCRC64Fingerprint = "\x1b\xd0\xc4\xd4\xe79\xc4\x06"

func NewEventoWhPosicionesAlta() EventoWhPosicionesAlta {
	r := EventoWhPosicionesAlta{}
	r.Identificacion = NewIdentificacion()

	r.Cabecera = NewCabecera()

	return r
}

func DeserializeEventoWhPosicionesAlta(r io.Reader) (EventoWhPosicionesAlta, error) {
	t := NewEventoWhPosicionesAlta()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoWhPosicionesAltaFromSchema(r io.Reader, schema string) (EventoWhPosicionesAlta, error) {
	t := NewEventoWhPosicionesAlta()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoWhPosicionesAlta(r EventoWhPosicionesAlta, w io.Writer) error {
	var err error
	err = writeIdentificacion(r.Identificacion, w)
	if err != nil {
		return err
	}
	err = writeCabecera(r.Cabecera, w)
	if err != nil {
		return err
	}
	return err
}

func (r EventoWhPosicionesAlta) Serialize(w io.Writer) error {
	return writeEventoWhPosicionesAlta(r, w)
}

func (r EventoWhPosicionesAlta) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.EventoWhPosiciones.Events.AltaCommon\",\"type\":\"record\"}},{\"name\":\"Cabecera\",\"type\":{\"fields\":[{\"name\":\"Ubicacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoUbicacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"CategoriaUbicacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"ABC\",\"type\":[\"null\",\"string\"]},{\"name\":\"Zona\",\"type\":[\"null\",\"string\"]},{\"name\":\"MezclaSKU\",\"type\":[\"null\",\"string\"]},{\"name\":\"MezclaLote\",\"type\":[\"null\",\"string\"]},{\"name\":\"PierdaLpn\",\"type\":[\"null\",\"string\"]},{\"name\":\"CapacidadCubica\",\"type\":[\"null\",\"string\"]},{\"name\":\"CapacidadPeso\",\"type\":[\"null\",\"string\"]},{\"name\":\"Altura\",\"type\":[\"null\",\"string\"]},{\"name\":\"Longitud\",\"type\":[\"null\",\"string\"]},{\"name\":\"Anchura\",\"type\":[\"null\",\"string\"]},{\"name\":\"Nivel\",\"type\":[\"null\",\"string\"]},{\"name\":\"Estado\",\"type\":[\"null\",\"string\"]},{\"name\":\"Temperatura\",\"type\":[\"null\",\"string\"]},{\"name\":\"FormatoUbicacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"FamiliaProducto\",\"type\":[\"null\",\"string\"]},{\"name\":\"Cliente\",\"type\":[\"null\",\"string\"]},{\"name\":\"Situacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoUbicacion2\",\"type\":[\"null\",\"string\"]},{\"name\":\"EuroEquivalente\",\"type\":[\"null\",\"string\"]},{\"name\":\"Llena\",\"type\":[\"null\",\"string\"]},{\"name\":\"Vacia\",\"type\":[\"null\",\"string\"]},{\"name\":\"PorcentajeLleno\",\"type\":[\"null\",\"string\"]},{\"name\":\"CantidadEvento\",\"type\":\"float\"},{\"name\":\"CantidadStock\",\"type\":\"float\"}],\"name\":\"Cabecera\",\"namespace\":\"Andreani.EventoWhPosiciones.Events.AltaCommon\",\"type\":\"record\"}}],\"name\":\"Andreani.EventoWhPosiciones.Events.Record.EventoWhPosicionesAlta\",\"type\":\"record\"}"
}

func (r EventoWhPosicionesAlta) SchemaName() string {
	return "Andreani.EventoWhPosiciones.Events.Record.EventoWhPosicionesAlta"
}

func (_ EventoWhPosicionesAlta) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoWhPosicionesAlta) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoWhPosicionesAlta) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoWhPosicionesAlta) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoWhPosicionesAlta) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoWhPosicionesAlta) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoWhPosicionesAlta) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoWhPosicionesAlta) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoWhPosicionesAlta) Get(i int) types.Field {
	switch i {
	case 0:
		r.Identificacion = NewIdentificacion()

		w := types.Record{Target: &r.Identificacion}

		return w

	case 1:
		r.Cabecera = NewCabecera()

		w := types.Record{Target: &r.Cabecera}

		return w

	}
	panic("Unknown field index")
}

func (r *EventoWhPosicionesAlta) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventoWhPosicionesAlta) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EventoWhPosicionesAlta) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EventoWhPosicionesAlta) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EventoWhPosicionesAlta) HintSize(int)                     { panic("Unsupported operation") }
func (_ EventoWhPosicionesAlta) Finalize()                        {}

func (_ EventoWhPosicionesAlta) AvroCRC64Fingerprint() []byte {
	return []byte(EventoWhPosicionesAltaAvroCRC64Fingerprint)
}

func (r EventoWhPosicionesAlta) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(output)
}

func (r *EventoWhPosicionesAlta) UnmarshalJSON(data []byte) error {
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
	return nil
}