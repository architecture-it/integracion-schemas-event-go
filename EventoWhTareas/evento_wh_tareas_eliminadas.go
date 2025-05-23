// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhTareasEliminadas.avsc
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

type EventoWhTareasEliminadas struct {
	Identificacion Identificacion `json:"Identificacion"`

	Tarea Tarea `json:"Tarea"`
}

const EventoWhTareasEliminadasAvroCRC64Fingerprint = "\x9bfA\xe3p\xed>\x89"

func NewEventoWhTareasEliminadas() EventoWhTareasEliminadas {
	r := EventoWhTareasEliminadas{}
	r.Identificacion = NewIdentificacion()

	r.Tarea = NewTarea()

	return r
}

func DeserializeEventoWhTareasEliminadas(r io.Reader) (EventoWhTareasEliminadas, error) {
	t := NewEventoWhTareasEliminadas()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoWhTareasEliminadasFromSchema(r io.Reader, schema string) (EventoWhTareasEliminadas, error) {
	t := NewEventoWhTareasEliminadas()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoWhTareasEliminadas(r EventoWhTareasEliminadas, w io.Writer) error {
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

func (r EventoWhTareasEliminadas) Serialize(w io.Writer) error {
	return writeEventoWhTareasEliminadas(r, w)
}

func (r EventoWhTareasEliminadas) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.EventoWhTareas.Events.TareasEliminadasCommon\",\"type\":\"record\"}},{\"name\":\"Tarea\",\"type\":{\"fields\":[{\"name\":\"IDTarea\",\"type\":\"string\"},{\"default\":null,\"name\":\"OrdenWH\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoDeTarea\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoTarea\",\"type\":[\"null\",\"string\"]},{\"name\":\"Propietario\",\"type\":[\"null\",\"string\"]},{\"name\":\"SKU\",\"type\":[\"null\",\"string\"]},{\"name\":\"LoteInternoWH\",\"type\":[\"null\",\"string\"]},{\"name\":\"CantidadTarea\",\"type\":\"float\"},{\"name\":\"UbicacionInicial\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"IDInicial\",\"type\":[\"null\",\"string\"]},{\"name\":\"UbicacionFinal\",\"type\":[\"null\",\"string\"]},{\"name\":\"IDFinal\",\"type\":[\"null\",\"string\"]},{\"name\":\"Estado\",\"type\":[\"null\",\"string\"]},{\"name\":\"Prioridad\",\"type\":[\"null\",\"string\"]},{\"name\":\"Usuario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaInicioTarea\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"FechaFinTarea\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"LineaDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeOleadaWH\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaDeCreacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"UsuarioCreacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaEdicion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"UsuarioEdicion\",\"type\":[\"null\",\"string\"]},{\"name\":\"PesoTara\",\"type\":\"float\"},{\"name\":\"PesoNeto\",\"type\":\"float\"},{\"name\":\"PesoBruto\",\"type\":\"float\"},{\"default\":null,\"name\":\"NumeroDeAsignacionWH\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoDeOrigen\",\"type\":[\"null\",\"string\"]},{\"name\":\"LoteSecundario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"FechaVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ProductoTrazable\",\"type\":[\"null\",\"string\"]},{\"name\":\"AlmacenConsumo\",\"type\":[\"null\",\"string\"]},{\"name\":\"EstadoLote\",\"type\":[\"null\",\"string\"]},{\"name\":\"BloqueoUbicacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"VidaUtilLote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"EntregaAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"ConsumoAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"Tarea\",\"namespace\":\"Andreani.EventoWhTareas.Events.TareasEliminadasCommon\",\"type\":\"record\"}}],\"name\":\"Andreani.EventoWhTareas.Events.Record.EventoWhTareasEliminadas\",\"type\":\"record\"}"
}

func (r EventoWhTareasEliminadas) SchemaName() string {
	return "Andreani.EventoWhTareas.Events.Record.EventoWhTareasEliminadas"
}

func (_ EventoWhTareasEliminadas) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoWhTareasEliminadas) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoWhTareasEliminadas) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoWhTareasEliminadas) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoWhTareasEliminadas) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoWhTareasEliminadas) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoWhTareasEliminadas) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoWhTareasEliminadas) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoWhTareasEliminadas) Get(i int) types.Field {
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

func (r *EventoWhTareasEliminadas) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventoWhTareasEliminadas) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EventoWhTareasEliminadas) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EventoWhTareasEliminadas) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EventoWhTareasEliminadas) HintSize(int)                     { panic("Unsupported operation") }
func (_ EventoWhTareasEliminadas) Finalize()                        {}

func (_ EventoWhTareasEliminadas) AvroCRC64Fingerprint() []byte {
	return []byte(EventoWhTareasEliminadasAvroCRC64Fingerprint)
}

func (r EventoWhTareasEliminadas) MarshalJSON() ([]byte, error) {
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

func (r *EventoWhTareasEliminadas) UnmarshalJSON(data []byte) error {
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
