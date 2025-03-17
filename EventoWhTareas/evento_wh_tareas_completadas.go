// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhTareasCompletadas.avsc
 */
package EventoWhTareasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EventoWhTareasCompletadas struct {
	Identificacion Identificacion `json:"Identificacion"`

	Tarea Tarea `json:"Tarea"`
}

const EventoWhTareasCompletadasAvroCRC64Fingerprint = "-\xf2\x1f\xcb9O\xe5\xdc"

func NewEventoWhTareasCompletadas() EventoWhTareasCompletadas {
	r := EventoWhTareasCompletadas{}
	r.Identificacion = NewIdentificacion()

	r.Tarea = NewTarea()

	return r
}

func DeserializeEventoWhTareasCompletadas(r io.Reader) (EventoWhTareasCompletadas, error) {
	t := NewEventoWhTareasCompletadas()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoWhTareasCompletadasFromSchema(r io.Reader, schema string) (EventoWhTareasCompletadas, error) {
	t := NewEventoWhTareasCompletadas()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoWhTareasCompletadas(r EventoWhTareasCompletadas, w io.Writer) error {
	var err error
	err = writeIdentificacion(r.Identificacion, w)
	if err != nil {
		return err
	}
	err = writeTarea(r.Tarea, w)
	if err != nil {
		return err
	}
	return err
}

func (r EventoWhTareasCompletadas) Serialize(w io.Writer) error {
	return writeEventoWhTareasCompletadas(r, w)
}

func (r EventoWhTareasCompletadas) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.EventoWhTareas.Events.TareasCompletadasCommon\",\"type\":\"record\"}},{\"name\":\"Tarea\",\"type\":{\"fields\":[{\"name\":\"IDTarea\",\"type\":\"string\"},{\"default\":null,\"name\":\"OrdenWH\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoDeTarea\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"SKU\",\"type\":\"string\"},{\"name\":\"LoteInternoWH\",\"type\":\"string\"},{\"name\":\"CantidadTarea\",\"type\":\"float\"},{\"name\":\"UbicacionInicial\",\"type\":\"string\"},{\"default\":null,\"name\":\"IDInicial\",\"type\":[\"null\",\"string\"]},{\"name\":\"UbicacionFinal\",\"type\":\"string\"},{\"name\":\"IDFinal\",\"type\":\"string\"},{\"name\":\"Estado\",\"type\":\"string\"},{\"name\":\"Prioridad\",\"type\":\"string\"},{\"name\":\"Usuario\",\"type\":\"string\"},{\"default\":null,\"name\":\"FechaInicioTarea\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"FechaFinTarea\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"LineaDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeOleadaWH\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaDeCreacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"UsuarioCreacion\",\"type\":\"string\"},{\"default\":null,\"name\":\"FechaEdicion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"UsuarioEdicion\",\"type\":\"string\"},{\"name\":\"PesoTara\",\"type\":\"float\"},{\"name\":\"PesoNeto\",\"type\":\"float\"},{\"name\":\"PesoBruto\",\"type\":\"float\"},{\"default\":null,\"name\":\"NumeroDeAsignacionWH\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoDeOrigen\",\"type\":[\"null\",\"string\"]},{\"name\":\"LoteSecundario\",\"type\":\"string\"},{\"default\":null,\"name\":\"FechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"FechaVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ProductoTrazable\",\"type\":\"string\"},{\"name\":\"AlmacenConsumo\",\"type\":\"string\"},{\"name\":\"EstadoLote\",\"type\":\"string\"},{\"name\":\"BloqueoUbicacion\",\"type\":\"string\"},{\"name\":\"VidaUtilLote\",\"type\":\"string\"},{\"default\":null,\"name\":\"EntregaAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"ConsumoAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"CodigoTarea\",\"type\":[\"null\",\"string\"]}],\"name\":\"Tarea\",\"namespace\":\"Andreani.EventoWhTareas.Events.TareasCompletadasCommon\",\"type\":\"record\"}}],\"name\":\"Andreani.EventoWhTareas.Events.Record.EventoWhTareasCompletadas\",\"type\":\"record\"}"
}

func (r EventoWhTareasCompletadas) SchemaName() string {
	return "Andreani.EventoWhTareas.Events.Record.EventoWhTareasCompletadas"
}

func (_ EventoWhTareasCompletadas) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoWhTareasCompletadas) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoWhTareasCompletadas) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoWhTareasCompletadas) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoWhTareasCompletadas) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoWhTareasCompletadas) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoWhTareasCompletadas) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoWhTareasCompletadas) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoWhTareasCompletadas) Get(i int) types.Field {
	switch i {
	case 0:
		r.Identificacion = NewIdentificacion()

		w := types.Record{Target: &r.Identificacion}

		return w

	case 1:
		r.Tarea = NewTarea()

		w := types.Record{Target: &r.Tarea}

		return w

	}
	panic("Unknown field index")
}

func (r *EventoWhTareasCompletadas) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventoWhTareasCompletadas) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EventoWhTareasCompletadas) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EventoWhTareasCompletadas) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EventoWhTareasCompletadas) HintSize(int)                     { panic("Unsupported operation") }
func (_ EventoWhTareasCompletadas) Finalize()                        {}

func (_ EventoWhTareasCompletadas) AvroCRC64Fingerprint() []byte {
	return []byte(EventoWhTareasCompletadasAvroCRC64Fingerprint)
}

func (r EventoWhTareasCompletadas) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Identificacion"], err = json.Marshal(r.Identificacion)
	if err != nil {
		return nil, err
	}
	output["Tarea"], err = json.Marshal(r.Tarea)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EventoWhTareasCompletadas) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["Tarea"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Tarea); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Tarea")
	}
	return nil
}
