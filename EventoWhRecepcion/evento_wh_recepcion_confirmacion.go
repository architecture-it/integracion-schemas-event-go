// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhRecepcionConfirmacion.avsc
 */
package EventoWhRecepcionEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EventoWhRecepcionConfirmacion struct {
	Identificacion Identificacion `json:"Identificacion"`

	Cabecera Cabecera `json:"Cabecera"`

	Detalle []Detalle `json:"Detalle"`
}

const EventoWhRecepcionConfirmacionAvroCRC64Fingerprint = "\xd7,\x1d]\xf6\xc7\x17\xdc"

func NewEventoWhRecepcionConfirmacion() EventoWhRecepcionConfirmacion {
	r := EventoWhRecepcionConfirmacion{}
	r.Identificacion = NewIdentificacion()

	r.Cabecera = NewCabecera()

	r.Detalle = make([]Detalle, 0)

	return r
}

func DeserializeEventoWhRecepcionConfirmacion(r io.Reader) (EventoWhRecepcionConfirmacion, error) {
	t := NewEventoWhRecepcionConfirmacion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoWhRecepcionConfirmacionFromSchema(r io.Reader, schema string) (EventoWhRecepcionConfirmacion, error) {
	t := NewEventoWhRecepcionConfirmacion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoWhRecepcionConfirmacion(r EventoWhRecepcionConfirmacion, w io.Writer) error {
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

func (r EventoWhRecepcionConfirmacion) Serialize(w io.Writer) error {
	return writeEventoWhRecepcionConfirmacion(r, w)
}

func (r EventoWhRecepcionConfirmacion) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.EventoWhRecepcion.Events.RecepcionConfirmacionCommon\",\"type\":\"record\"}},{\"name\":\"Cabecera\",\"type\":{\"fields\":[{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"RecepcionWH\",\"type\":\"string\"},{\"name\":\"Remito\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrdenCompra\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoRecepcion\",\"type\":[\"null\",\"string\"]},{\"name\":\"ReferenciaTransportista\",\"type\":[\"null\",\"string\"]},{\"name\":\"ReferenciaContenedor\",\"type\":[\"null\",\"string\"]},{\"name\":\"Muelle\",\"type\":[\"null\",\"string\"]},{\"name\":\"NumeroCita\",\"type\":[\"null\",\"string\"]},{\"name\":\"ReferenciaProveedor\",\"type\":[\"null\",\"string\"]},{\"name\":\"NumeroGuia\",\"type\":[\"null\",\"string\"]},{\"name\":\"Calle\",\"type\":[\"null\",\"string\"]},{\"name\":\"Numero\",\"type\":[\"null\",\"string\"]},{\"name\":\"Piso\",\"type\":[\"null\",\"string\"]},{\"name\":\"Dpto\",\"type\":[\"null\",\"string\"]},{\"name\":\"GLN\",\"type\":[\"null\",\"string\"]},{\"name\":\"CiudadExpedidor\",\"type\":[\"null\",\"string\"]},{\"name\":\"Contacto\",\"type\":[\"null\",\"string\"]},{\"name\":\"Email\",\"type\":[\"null\",\"string\"]},{\"name\":\"CodISOPais\",\"type\":[\"null\",\"string\"]},{\"name\":\"Telefono\",\"type\":[\"null\",\"string\"]},{\"name\":\"EstadoExpedidor\",\"type\":[\"null\",\"string\"]},{\"name\":\"CodPostalExpedidor\",\"type\":[\"null\",\"string\"]},{\"name\":\"FechaLlegada\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaRecepcion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"CantEsperadaTotal\",\"type\":\"float\"},{\"name\":\"CantRecibidaTotal\",\"type\":\"float\"}],\"name\":\"Cabecera\",\"namespace\":\"Andreani.EventoWhRecepcion.Events.RecepcionConfirmacionCommon\",\"type\":\"record\"}},{\"name\":\"Detalle\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"SKU\",\"type\":\"string\"},{\"name\":\"UbicacionDestino\",\"type\":\"string\"},{\"name\":\"LPNDestino\",\"type\":\"string\"},{\"name\":\"PaqueteLote\",\"type\":\"string\"},{\"name\":\"LoteCajitaFabricante\",\"type\":\"string\"},{\"name\":\"LoteSecundario\",\"type\":\"string\"},{\"name\":\"FechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ProductoTrazable\",\"type\":\"string\"},{\"name\":\"AlmacenConsumo\",\"type\":\"string\"},{\"name\":\"EstadoLote\",\"type\":\"string\"},{\"name\":\"BloqueoUbicacion\",\"type\":\"string\"},{\"name\":\"VidaUtilLote\",\"type\":\"string\"},{\"name\":\"EntregaAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ConsumoAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"Contramuestras\",\"type\":[\"null\",\"string\"]},{\"name\":\"EstadoOTAcondi\",\"type\":[\"null\",\"string\"]},{\"name\":\"EstadoOTTraza\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoAcondi\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoTraza\",\"type\":[\"null\",\"string\"]},{\"name\":\"ContratoServicioIngreso\",\"type\":[\"null\",\"string\"]},{\"name\":\"NomenclaturaContratoServicioIngreso\",\"type\":[\"null\",\"string\"]},{\"name\":\"DescripcionServicioIngreso\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoLineaMatriz\",\"type\":[\"null\",\"string\"]},{\"name\":\"CodConCalidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"AccionConCalidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"ResultadoConCalidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"LineaExterna\",\"type\":[\"null\",\"string\"]},{\"name\":\"CantEsperada\",\"type\":\"float\"},{\"name\":\"CantRecibida\",\"type\":\"float\"},{\"name\":\"ValorDeclaradoLinea\",\"type\":[\"null\",\"string\"]},{\"name\":\"UnidadMedida\",\"type\":[\"null\",\"string\"]},{\"name\":\"LineaRecepcionWH\",\"type\":[\"null\",\"string\"]}],\"name\":\"Detalle\",\"namespace\":\"Andreani.EventoWhRecepcion.Events.RecepcionConfirmacionCommon\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.EventoWhRecepcion.Events.Record.EventoWhRecepcionConfirmacion\",\"type\":\"record\"}"
}

func (r EventoWhRecepcionConfirmacion) SchemaName() string {
	return "Andreani.EventoWhRecepcion.Events.Record.EventoWhRecepcionConfirmacion"
}

func (_ EventoWhRecepcionConfirmacion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoWhRecepcionConfirmacion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoWhRecepcionConfirmacion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoWhRecepcionConfirmacion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoWhRecepcionConfirmacion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoWhRecepcionConfirmacion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoWhRecepcionConfirmacion) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoWhRecepcionConfirmacion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoWhRecepcionConfirmacion) Get(i int) types.Field {
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

func (r *EventoWhRecepcionConfirmacion) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventoWhRecepcionConfirmacion) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EventoWhRecepcionConfirmacion) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ EventoWhRecepcionConfirmacion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ EventoWhRecepcionConfirmacion) HintSize(int)             { panic("Unsupported operation") }
func (_ EventoWhRecepcionConfirmacion) Finalize()                {}

func (_ EventoWhRecepcionConfirmacion) AvroCRC64Fingerprint() []byte {
	return []byte(EventoWhRecepcionConfirmacionAvroCRC64Fingerprint)
}

func (r EventoWhRecepcionConfirmacion) MarshalJSON() ([]byte, error) {
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

func (r *EventoWhRecepcionConfirmacion) UnmarshalJSON(data []byte) error {
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
