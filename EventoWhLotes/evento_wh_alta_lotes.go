// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhAltaLotes.avsc
 */
package EventoWhLotesEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EventoWhAltaLotes struct {
	Identificacion Identificacion `json:"Identificacion"`

	Detalle Detalle `json:"Detalle"`
}

const EventoWhAltaLotesAvroCRC64Fingerprint = "<\t\xe7I=\x83\xffU"

func NewEventoWhAltaLotes() EventoWhAltaLotes {
	r := EventoWhAltaLotes{}
	r.Identificacion = NewIdentificacion()

	r.Detalle = NewDetalle()

	return r
}

func DeserializeEventoWhAltaLotes(r io.Reader) (EventoWhAltaLotes, error) {
	t := NewEventoWhAltaLotes()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoWhAltaLotesFromSchema(r io.Reader, schema string) (EventoWhAltaLotes, error) {
	t := NewEventoWhAltaLotes()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoWhAltaLotes(r EventoWhAltaLotes, w io.Writer) error {
	var err error
	err = writeIdentificacion(r.Identificacion, w)
	if err != nil {
		return err
	}
	err = writeDetalle(r.Detalle, w)
	if err != nil {
		return err
	}
	return err
}

func (r EventoWhAltaLotes) Serialize(w io.Writer) error {
	return writeEventoWhAltaLotes(r, w)
}

func (r EventoWhAltaLotes) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.EventoWhLotes.Events.AltaLoteCommon\",\"type\":\"record\"}},{\"name\":\"Detalle\",\"type\":{\"fields\":[{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"SKU\",\"type\":\"string\"},{\"name\":\"PaqueteLote\",\"type\":\"string\"},{\"name\":\"LoteCajitaFabricante\",\"type\":\"string\"},{\"name\":\"LoteSecundario\",\"type\":\"string\"},{\"default\":null,\"name\":\"FechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"FechaVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ProductoTrazable\",\"type\":\"string\"},{\"name\":\"AlmacenConsumo\",\"type\":\"string\"},{\"name\":\"EstadoLote\",\"type\":\"string\"},{\"name\":\"BloqueoUbicacion\",\"type\":\"string\"},{\"name\":\"VidaUtilLote\",\"type\":\"string\"},{\"default\":null,\"name\":\"EntregaAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"ConsumoAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaCreacion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"Cantidad\",\"type\":[\"null\",\"string\"]}],\"name\":\"Detalle\",\"namespace\":\"Andreani.EventoWhLotes.Events.AltaLoteCommon\",\"type\":\"record\"}}],\"name\":\"Andreani.EventoWhLotes.Events.Record.EventoWhAltaLotes\",\"type\":\"record\"}"
}

func (r EventoWhAltaLotes) SchemaName() string {
	return "Andreani.EventoWhLotes.Events.Record.EventoWhAltaLotes"
}

func (_ EventoWhAltaLotes) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoWhAltaLotes) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoWhAltaLotes) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoWhAltaLotes) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoWhAltaLotes) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoWhAltaLotes) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoWhAltaLotes) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoWhAltaLotes) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoWhAltaLotes) Get(i int) types.Field {
	switch i {
	case 0:
		r.Identificacion = NewIdentificacion()

		w := types.Record{Target: &r.Identificacion}

		return w

	case 1:
		r.Detalle = NewDetalle()

		w := types.Record{Target: &r.Detalle}

		return w

	}
	panic("Unknown field index")
}

func (r *EventoWhAltaLotes) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventoWhAltaLotes) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EventoWhAltaLotes) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EventoWhAltaLotes) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EventoWhAltaLotes) HintSize(int)                     { panic("Unsupported operation") }
func (_ EventoWhAltaLotes) Finalize()                        {}

func (_ EventoWhAltaLotes) AvroCRC64Fingerprint() []byte {
	return []byte(EventoWhAltaLotesAvroCRC64Fingerprint)
}

func (r EventoWhAltaLotes) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Identificacion"], err = json.Marshal(r.Identificacion)
	if err != nil {
		return nil, err
	}
	output["Detalle"], err = json.Marshal(r.Detalle)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EventoWhAltaLotes) UnmarshalJSON(data []byte) error {
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
