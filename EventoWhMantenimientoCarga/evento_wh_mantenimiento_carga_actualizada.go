// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhMantenimientoCargaActualizada.avsc
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

type EventoWhMantenimientoCargaActualizada struct {
	Identificacion Identificacion `json:"Identificacion"`

	Cabecera Cabecera `json:"Cabecera"`

	Detalle []Detalle `json:"Detalle"`
}

const EventoWhMantenimientoCargaActualizadaAvroCRC64Fingerprint = "\xcfV5\xfd\x13\\\xe7)"

func NewEventoWhMantenimientoCargaActualizada() EventoWhMantenimientoCargaActualizada {
	r := EventoWhMantenimientoCargaActualizada{}
	r.Identificacion = NewIdentificacion()

	r.Cabecera = NewCabecera()

	r.Detalle = make([]Detalle, 0)

	return r
}

func DeserializeEventoWhMantenimientoCargaActualizada(r io.Reader) (EventoWhMantenimientoCargaActualizada, error) {
	t := NewEventoWhMantenimientoCargaActualizada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoWhMantenimientoCargaActualizadaFromSchema(r io.Reader, schema string) (EventoWhMantenimientoCargaActualizada, error) {
	t := NewEventoWhMantenimientoCargaActualizada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoWhMantenimientoCargaActualizada(r EventoWhMantenimientoCargaActualizada, w io.Writer) error {
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

func (r EventoWhMantenimientoCargaActualizada) Serialize(w io.Writer) error {
	return writeEventoWhMantenimientoCargaActualizada(r, w)
}

func (r EventoWhMantenimientoCargaActualizada) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.EventoWhMantenimientoCarga.Events.ActualizadaCommon\",\"type\":\"record\"}},{\"name\":\"Cabecera\",\"type\":{\"fields\":[{\"name\":\"CargaWH\",\"type\":[\"null\",\"string\"]},{\"name\":\"Ruta\",\"type\":[\"null\",\"string\"]},{\"name\":\"MuelleDeCarga\",\"type\":[\"null\",\"string\"]},{\"name\":\"Estado\",\"type\":[\"null\",\"string\"]},{\"name\":\"FechaDePartida\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"IDTransportista\",\"type\":[\"null\",\"string\"]},{\"name\":\"PatenteDelTransporte\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoDeTransporte\",\"type\":[\"null\",\"string\"]},{\"name\":\"TotalUnidades\",\"type\":\"float\"},{\"name\":\"NumeroDeCargaExterno\",\"type\":[\"null\",\"string\"]},{\"name\":\"FechaInicioDeCarga\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaFinDeCarga\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"NumerosDeIdCargados\",\"type\":[\"null\",\"string\"]},{\"name\":\"FechaDeCierreDeCarga\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"NombreDelConductor\",\"type\":[\"null\",\"string\"]},{\"name\":\"NumeroDelPrecinto\",\"type\":[\"null\",\"string\"]},{\"name\":\"FechaExpedicion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"Cabecera\",\"namespace\":\"Andreani.EventoWhMantenimientoCarga.Events.ActualizadaCommon\",\"type\":\"record\"}},{\"name\":\"Detalle\",\"type\":{\"items\":{\"fields\":[{\"name\":\"UnidadesDelPedido\",\"type\":\"float\"},{\"name\":\"OrdenCliente\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrdenWH\",\"type\":[\"null\",\"string\"]}],\"name\":\"Detalle\",\"namespace\":\"Andreani.EventoWhMantenimientoCarga.Events.ActualizadaCommon\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.EventoWhMantenimientoCarga.Events.Record.EventoWhMantenimientoCargaActualizada\",\"type\":\"record\"}"
}

func (r EventoWhMantenimientoCargaActualizada) SchemaName() string {
	return "Andreani.EventoWhMantenimientoCarga.Events.Record.EventoWhMantenimientoCargaActualizada"
}

func (_ EventoWhMantenimientoCargaActualizada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoWhMantenimientoCargaActualizada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoWhMantenimientoCargaActualizada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoWhMantenimientoCargaActualizada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoWhMantenimientoCargaActualizada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoWhMantenimientoCargaActualizada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoWhMantenimientoCargaActualizada) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoWhMantenimientoCargaActualizada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoWhMantenimientoCargaActualizada) Get(i int) types.Field {
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

func (r *EventoWhMantenimientoCargaActualizada) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventoWhMantenimientoCargaActualizada) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EventoWhMantenimientoCargaActualizada) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ EventoWhMantenimientoCargaActualizada) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ EventoWhMantenimientoCargaActualizada) HintSize(int) { panic("Unsupported operation") }
func (_ EventoWhMantenimientoCargaActualizada) Finalize()    {}

func (_ EventoWhMantenimientoCargaActualizada) AvroCRC64Fingerprint() []byte {
	return []byte(EventoWhMantenimientoCargaActualizadaAvroCRC64Fingerprint)
}

func (r EventoWhMantenimientoCargaActualizada) MarshalJSON() ([]byte, error) {
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

func (r *EventoWhMantenimientoCargaActualizada) UnmarshalJSON(data []byte) error {
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
