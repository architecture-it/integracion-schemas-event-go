// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoOT.avsc
 */
package EventoWhAltaOtEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EventoOT struct {
	Identificacion Identificacion `json:"Identificacion"`

	Cabecera Cabecera `json:"Cabecera"`

	Detalle []Detalle `json:"Detalle"`
}

const EventoOTAvroCRC64Fingerprint = "#v\\\r\x03\xc7{\xa4"

func NewEventoOT() EventoOT {
	r := EventoOT{}
	r.Identificacion = NewIdentificacion()

	r.Cabecera = NewCabecera()

	r.Detalle = make([]Detalle, 0)

	return r
}

func DeserializeEventoOT(r io.Reader) (EventoOT, error) {
	t := NewEventoOT()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoOTFromSchema(r io.Reader, schema string) (EventoOT, error) {
	t := NewEventoOT()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoOT(r EventoOT, w io.Writer) error {
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

func (r EventoOT) Serialize(w io.Writer) error {
	return writeEventoOT(r, w)
}

func (r EventoOT) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.EventoWhAltaOt.Events.AltaCommon\",\"type\":\"record\"}},{\"name\":\"Cabecera\",\"type\":{\"fields\":[{\"name\":\"IdCliente\",\"type\":\"string\"},{\"name\":\"Remito\",\"type\":\"string\"},{\"name\":\"FechaRemito\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"GlnOrigen\",\"type\":\"string\"},{\"name\":\"ReferenciaExterna\",\"type\":\"string\"},{\"name\":\"IdUsuarioModif\",\"type\":\"string\"},{\"name\":\"FechaModificacionRegistro\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"GlnDestino\",\"type\":\"string\"},{\"name\":\"IdTipoOT\",\"type\":\"int\"},{\"name\":\"Factura\",\"type\":\"string\"},{\"name\":\"IdEstadoLote\",\"type\":\"string\"},{\"name\":\"IdEvento\",\"type\":\"int\"},{\"name\":\"IdEstado\",\"type\":\"int\"},{\"name\":\"Contenedor\",\"type\":\"string\"},{\"name\":\"Pedido\",\"type\":\"string\"},{\"name\":\"IdTipoDocumento\",\"type\":\"int\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Apellido\",\"type\":\"string\"},{\"name\":\"NroDocumento\",\"type\":\"string\"},{\"name\":\"Direccion\",\"type\":\"string\"},{\"name\":\"NroEntrega\",\"type\":\"string\"},{\"name\":\"ClienteDestinatario\",\"type\":\"string\"},{\"name\":\"ClienteFacturacion\",\"type\":\"string\"},{\"name\":\"FechaEnvioBoss\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"IdEstadoTransferenciaBoss\",\"type\":\"int\"},{\"name\":\"CorrecionFix\",\"type\":\"boolean\"},{\"name\":\"FechaCreacionRegistro\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"IdUsuarioCreacionRegistro\",\"type\":\"string\"},{\"name\":\"FormatoNuevo\",\"type\":\"boolean\"}],\"name\":\"Cabecera\",\"namespace\":\"Andreani.EventoWhAltaOt.Events.AltaCommon\",\"type\":\"record\"}},{\"name\":\"Detalle\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Gtin\",\"type\":\"string\"},{\"name\":\"Lote\",\"type\":\"string\"},{\"name\":\"FechaVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"IdUsuarioModificacion\",\"type\":\"string\"},{\"name\":\"FechaModificacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"Cantidad\",\"type\":\"float\"}],\"name\":\"Detalle\",\"namespace\":\"Andreani.EventoWhAltaOt.Events.AltaCommon\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.EventoWhAltaOt.Events.Record.EventoOT\",\"type\":\"record\"}"
}

func (r EventoOT) SchemaName() string {
	return "Andreani.EventoWhAltaOt.Events.Record.EventoOT"
}

func (_ EventoOT) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoOT) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoOT) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoOT) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoOT) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoOT) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoOT) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoOT) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoOT) Get(i int) types.Field {
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

func (r *EventoOT) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventoOT) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EventoOT) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EventoOT) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EventoOT) HintSize(int)                     { panic("Unsupported operation") }
func (_ EventoOT) Finalize()                        {}

func (_ EventoOT) AvroCRC64Fingerprint() []byte {
	return []byte(EventoOTAvroCRC64Fingerprint)
}

func (r EventoOT) MarshalJSON() ([]byte, error) {
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

func (r *EventoOT) UnmarshalJSON(data []byte) error {
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
