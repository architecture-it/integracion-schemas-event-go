// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhArticuloAsnConfirmacion.avsc
 */
package EventoWhArticulosEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EventoWhArticuloAsnConfirmacion struct {
	Identificacion Identificacion `json:"Identificacion"`

	Cabecera Cabecera `json:"Cabecera"`

	Detalle []Detalle `json:"Detalle"`
}

const EventoWhArticuloAsnConfirmacionAvroCRC64Fingerprint = "\x86:Yn\xbb\x91C\xa2"

func NewEventoWhArticuloAsnConfirmacion() EventoWhArticuloAsnConfirmacion {
	r := EventoWhArticuloAsnConfirmacion{}
	r.Identificacion = NewIdentificacion()

	r.Cabecera = NewCabecera()

	r.Detalle = make([]Detalle, 0)

	return r
}

func DeserializeEventoWhArticuloAsnConfirmacion(r io.Reader) (EventoWhArticuloAsnConfirmacion, error) {
	t := NewEventoWhArticuloAsnConfirmacion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoWhArticuloAsnConfirmacionFromSchema(r io.Reader, schema string) (EventoWhArticuloAsnConfirmacion, error) {
	t := NewEventoWhArticuloAsnConfirmacion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoWhArticuloAsnConfirmacion(r EventoWhArticuloAsnConfirmacion, w io.Writer) error {
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

func (r EventoWhArticuloAsnConfirmacion) Serialize(w io.Writer) error {
	return writeEventoWhArticuloAsnConfirmacion(r, w)
}

func (r EventoWhArticuloAsnConfirmacion) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.EventoWhArticulos.Events.AsnConfirmadaCommon\",\"type\":\"record\"}},{\"name\":\"Cabecera\",\"type\":{\"fields\":[{\"name\":\"SKU\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"TipoOrigen\",\"type\":\"string\"},{\"name\":\"CodigoOrigenWH\",\"type\":\"string\"},{\"name\":\"CodigoOrigenExterno\",\"type\":\"string\"},{\"name\":\"ContratoServicioIngreso\",\"type\":[\"null\",\"string\"]},{\"name\":\"NomenclaturaServicioIngreso\",\"type\":[\"null\",\"string\"]},{\"name\":\"DescripcionServicioIngreso\",\"type\":[\"null\",\"string\"]},{\"name\":\"StockAnteriorSKU\",\"type\":\"float\"},{\"name\":\"StockTotalSKU\",\"type\":\"float\"},{\"name\":\"StockDisponibleSKU\",\"type\":\"float\"},{\"name\":\"StockEnTransitoSKU\",\"type\":\"float\"}],\"name\":\"Cabecera\",\"namespace\":\"Andreani.EventoWhArticulos.Events.AsnConfirmadaCommon\",\"type\":\"record\"}},{\"name\":\"Detalle\",\"type\":{\"items\":{\"fields\":[{\"name\":\"LineaExterna\",\"type\":\"string\"},{\"name\":\"PaqueteLote\",\"type\":\"string\"},{\"name\":\"LoteCajitaFabricante\",\"type\":\"string\"},{\"name\":\"LoteSecundario\",\"type\":\"string\"},{\"name\":\"FechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ProductoTrazable\",\"type\":\"string\"},{\"name\":\"AlmacenConsumo\",\"type\":\"string\"},{\"name\":\"EstadoLote\",\"type\":\"string\"},{\"name\":\"BloqueoUbicacion\",\"type\":\"string\"},{\"name\":\"VidaUtilLote\",\"type\":\"string\"},{\"name\":\"Lpn\",\"type\":[\"null\",\"string\"]},{\"name\":\"EntregaAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ConsumoAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"StockEvento\",\"type\":\"float\"},{\"name\":\"StockEnTransito\",\"type\":\"float\"},{\"name\":\"StockAnterior\",\"type\":\"float\"},{\"name\":\"StockTotal\",\"type\":\"float\"}],\"name\":\"Detalle\",\"namespace\":\"Andreani.EventoWhArticulos.Events.AsnConfirmadaCommon\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.EventoWhArticulos.Events.Record.EventoWhArticuloAsnConfirmacion\",\"type\":\"record\"}"
}

func (r EventoWhArticuloAsnConfirmacion) SchemaName() string {
	return "Andreani.EventoWhArticulos.Events.Record.EventoWhArticuloAsnConfirmacion"
}

func (_ EventoWhArticuloAsnConfirmacion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoWhArticuloAsnConfirmacion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoWhArticuloAsnConfirmacion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoWhArticuloAsnConfirmacion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoWhArticuloAsnConfirmacion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoWhArticuloAsnConfirmacion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoWhArticuloAsnConfirmacion) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoWhArticuloAsnConfirmacion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoWhArticuloAsnConfirmacion) Get(i int) types.Field {
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

func (r *EventoWhArticuloAsnConfirmacion) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventoWhArticuloAsnConfirmacion) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EventoWhArticuloAsnConfirmacion) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ EventoWhArticuloAsnConfirmacion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ EventoWhArticuloAsnConfirmacion) HintSize(int)             { panic("Unsupported operation") }
func (_ EventoWhArticuloAsnConfirmacion) Finalize()                {}

func (_ EventoWhArticuloAsnConfirmacion) AvroCRC64Fingerprint() []byte {
	return []byte(EventoWhArticuloAsnConfirmacionAvroCRC64Fingerprint)
}

func (r EventoWhArticuloAsnConfirmacion) MarshalJSON() ([]byte, error) {
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

func (r *EventoWhArticuloAsnConfirmacion) UnmarshalJSON(data []byte) error {
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
