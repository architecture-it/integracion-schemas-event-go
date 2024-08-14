// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhAlmacenLotes.avsc
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

type EventoWhAlmacenLotes struct {
	Identificacion Identificacion `json:"Identificacion"`

	Detalle Detalle `json:"Detalle"`
}

const EventoWhAlmacenLotesAvroCRC64Fingerprint = "\xde* \n\xf0\"\x97\x0e"

func NewEventoWhAlmacenLotes() EventoWhAlmacenLotes {
	r := EventoWhAlmacenLotes{}
	r.Identificacion = NewIdentificacion()

	r.Detalle = NewDetalle()

	return r
}

func DeserializeEventoWhAlmacenLotes(r io.Reader) (EventoWhAlmacenLotes, error) {
	t := NewEventoWhAlmacenLotes()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoWhAlmacenLotesFromSchema(r io.Reader, schema string) (EventoWhAlmacenLotes, error) {
	t := NewEventoWhAlmacenLotes()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoWhAlmacenLotes(r EventoWhAlmacenLotes, w io.Writer) error {
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

func (r EventoWhAlmacenLotes) Serialize(w io.Writer) error {
	return writeEventoWhAlmacenLotes(r, w)
}

func (r EventoWhAlmacenLotes) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.EventoWhLotes.Events.AlmacenLoteCommon\",\"type\":\"record\"}},{\"name\":\"Detalle\",\"type\":{\"fields\":[{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"SKU\",\"type\":\"string\"},{\"name\":\"PaqueteLote\",\"type\":\"string\"},{\"name\":\"LoteCajitaFabricante\",\"type\":\"string\"},{\"name\":\"LoteSecundario\",\"type\":\"string\"},{\"default\":null,\"name\":\"FechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"FechaVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ProductoTrazable\",\"type\":\"string\"},{\"name\":\"AlmacenConsumoDestino\",\"type\":\"string\"},{\"name\":\"AlmacenConsumoOrigen\",\"type\":\"string\"},{\"name\":\"EstadoLote\",\"type\":\"string\"},{\"name\":\"BloqueoUbicacion\",\"type\":\"string\"},{\"name\":\"VidaUtilLote\",\"type\":\"string\"},{\"default\":null,\"name\":\"EntregaAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"ConsumoAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaCreacion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"Cantidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Lpn\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ReferenciaCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"IdInterno\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroASNInterna\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroASNExterna\",\"type\":[\"null\",\"string\"]}],\"name\":\"Detalle\",\"namespace\":\"Andreani.EventoWhLotes.Events.AlmacenLoteCommon\",\"type\":\"record\"}}],\"name\":\"Andreani.EventoWhLotes.Events.Record.EventoWhAlmacenLotes\",\"type\":\"record\"}"
}

func (r EventoWhAlmacenLotes) SchemaName() string {
	return "Andreani.EventoWhLotes.Events.Record.EventoWhAlmacenLotes"
}

func (_ EventoWhAlmacenLotes) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoWhAlmacenLotes) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoWhAlmacenLotes) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoWhAlmacenLotes) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoWhAlmacenLotes) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoWhAlmacenLotes) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoWhAlmacenLotes) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoWhAlmacenLotes) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoWhAlmacenLotes) Get(i int) types.Field {
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

func (r *EventoWhAlmacenLotes) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventoWhAlmacenLotes) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EventoWhAlmacenLotes) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EventoWhAlmacenLotes) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EventoWhAlmacenLotes) HintSize(int)                     { panic("Unsupported operation") }
func (_ EventoWhAlmacenLotes) Finalize()                        {}

func (_ EventoWhAlmacenLotes) AvroCRC64Fingerprint() []byte {
	return []byte(EventoWhAlmacenLotesAvroCRC64Fingerprint)
}

func (r EventoWhAlmacenLotes) MarshalJSON() ([]byte, error) {
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

func (r *EventoWhAlmacenLotes) UnmarshalJSON(data []byte) error {
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
