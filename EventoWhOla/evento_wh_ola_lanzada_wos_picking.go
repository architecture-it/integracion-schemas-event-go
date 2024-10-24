// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhOlaLanzadaWosPicking.avsc
 */
package EventoWhOlaEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EventoWhOlaLanzadaWosPicking struct {
	Identificacion Identificacion `json:"Identificacion"`

	Cabecera Cabecera `json:"Cabecera"`

	Detalle []Detalle `json:"Detalle"`
}

const EventoWhOlaLanzadaWosPickingAvroCRC64Fingerprint = "\xc52\x00\x87)\xa0\x944"

func NewEventoWhOlaLanzadaWosPicking() EventoWhOlaLanzadaWosPicking {
	r := EventoWhOlaLanzadaWosPicking{}
	r.Identificacion = NewIdentificacion()

	r.Cabecera = NewCabecera()

	r.Detalle = make([]Detalle, 0)

	return r
}

func DeserializeEventoWhOlaLanzadaWosPicking(r io.Reader) (EventoWhOlaLanzadaWosPicking, error) {
	t := NewEventoWhOlaLanzadaWosPicking()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoWhOlaLanzadaWosPickingFromSchema(r io.Reader, schema string) (EventoWhOlaLanzadaWosPicking, error) {
	t := NewEventoWhOlaLanzadaWosPicking()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoWhOlaLanzadaWosPicking(r EventoWhOlaLanzadaWosPicking, w io.Writer) error {
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

func (r EventoWhOlaLanzadaWosPicking) Serialize(w io.Writer) error {
	return writeEventoWhOlaLanzadaWosPicking(r, w)
}

func (r EventoWhOlaLanzadaWosPicking) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.EventoWhOla.Events.LanzadaWosPickingCommon\",\"type\":\"record\"}},{\"name\":\"Cabecera\",\"type\":{\"fields\":[{\"name\":\"Ola\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"}],\"name\":\"Cabecera\",\"namespace\":\"Andreani.EventoWhOla.Events.LanzadaWosPickingCommon\",\"type\":\"record\"}},{\"name\":\"Detalle\",\"type\":{\"items\":{\"fields\":[{\"name\":\"TaskDetailKey\",\"type\":\"string\"},{\"name\":\"TaskType\",\"type\":\"string\"},{\"name\":\"StorerKey\",\"type\":\"string\"},{\"name\":\"Sku\",\"type\":\"string\"},{\"name\":\"Lot\",\"type\":\"string\"},{\"name\":\"UOM\",\"type\":\"string\"},{\"name\":\"UOMQty\",\"type\":\"int\"},{\"name\":\"Qty\",\"type\":\"int\"},{\"name\":\"FromLoc\",\"type\":\"string\"},{\"name\":\"LogicalFromLoc\",\"type\":\"string\"},{\"name\":\"FromID\",\"type\":\"string\"},{\"name\":\"ToLoc\",\"type\":\"string\"},{\"name\":\"LogicalToLoc\",\"type\":\"string\"},{\"name\":\"ToID\",\"type\":\"string\"},{\"name\":\"CaseID\",\"type\":\"string\"},{\"name\":\"PickMethod\",\"type\":\"string\"},{\"name\":\"Status\",\"type\":\"string\"},{\"name\":\"StatusMsg\",\"type\":\"string\"},{\"name\":\"Priority\",\"type\":\"string\"},{\"name\":\"SourcePriority\",\"type\":\"string\"},{\"name\":\"HoldKey\",\"type\":\"string\"},{\"name\":\"UserKey\",\"type\":\"string\"},{\"name\":\"UserPosition\",\"type\":\"string\"},{\"name\":\"UserKeyOverride\",\"type\":\"string\"},{\"name\":\"StartTime\",\"type\":\"string\"},{\"name\":\"EndTime\",\"type\":\"string\"},{\"name\":\"SourceType\",\"type\":\"string\"},{\"name\":\"SourceKey\",\"type\":\"string\"},{\"name\":\"PickDetailKey\",\"type\":\"string\"},{\"name\":\"OrderKey\",\"type\":\"string\"},{\"name\":\"OrderLineNumber\",\"type\":\"string\"},{\"name\":\"ListKey\",\"type\":\"string\"},{\"name\":\"WaveKey\",\"type\":\"string\"},{\"name\":\"ReasonKey\",\"type\":\"string\"},{\"name\":\"Message01\",\"type\":\"string\"},{\"name\":\"Message02\",\"type\":\"string\"},{\"name\":\"Message03\",\"type\":\"string\"},{\"name\":\"Door\",\"type\":\"string\"},{\"name\":\"Route\",\"type\":\"string\"},{\"name\":\"Stop\",\"type\":\"string\"},{\"name\":\"PutawayZone\",\"type\":\"string\"},{\"name\":\"Altsku\",\"type\":\"string\"},{\"name\":\"EXT_UDF_STR1\",\"type\":\"string\"},{\"name\":\"EXT_UDF_LKUP5\",\"type\":\"string\"},{\"name\":\"LOTTABLE01\",\"type\":\"string\"},{\"name\":\"LOTTABLE02\",\"type\":\"string\"},{\"name\":\"LOTTABLE03\",\"type\":\"string\"},{\"name\":\"LOTTABLE04\",\"type\":\"string\"},{\"name\":\"LOTTABLE05\",\"type\":\"string\"},{\"name\":\"LOTTABLE06\",\"type\":\"string\"},{\"name\":\"LOTTABLE07\",\"type\":\"string\"},{\"name\":\"LOTTABLE08\",\"type\":\"string\"},{\"name\":\"LOTTABLE09\",\"type\":\"string\"},{\"name\":\"LOTTABLE10\",\"type\":\"string\"},{\"name\":\"LOTTABLE11\",\"type\":\"string\"},{\"name\":\"LOTTABLE12\",\"type\":\"string\"},{\"name\":\"CKEXT_UDF_STR1\",\"type\":\"string\"},{\"name\":\"CKEXT_UDF_STR2\",\"type\":\"string\"},{\"name\":\"CKEXT_UDF_STR3\",\"type\":\"string\"},{\"name\":\"CKEXT_UDF_STR4\",\"type\":\"string\"},{\"name\":\"CKEXT_UDF_STR5\",\"type\":\"string\"},{\"name\":\"CKEXT_UDF_STR6\",\"type\":\"string\"},{\"name\":\"CKEXT_UDF_STR7\",\"type\":\"string\"},{\"name\":\"CKEXT_UDF_STR8\",\"type\":\"string\"},{\"name\":\"CKEXT_UDF_STR9\",\"type\":\"string\"},{\"name\":\"CKEXT_UDF_STR10\",\"type\":\"string\"},{\"name\":\"CKEXT_UDF_STR11\",\"type\":\"string\"},{\"name\":\"CKEXT_UDF_STR12\",\"type\":\"string\"},{\"name\":\"ValidaLoteWos\",\"type\":\"string\"},{\"name\":\"ValidaSerieWos\",\"type\":\"string\"},{\"name\":\"OCDFLAG\",\"type\":\"string\"}],\"name\":\"Detalle\",\"namespace\":\"Andreani.EventoWhOla.Events.LanzadaWosPickingCommon\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.EventoWhOla.Events.Record.EventoWhOlaLanzadaWosPicking\",\"type\":\"record\"}"
}

func (r EventoWhOlaLanzadaWosPicking) SchemaName() string {
	return "Andreani.EventoWhOla.Events.Record.EventoWhOlaLanzadaWosPicking"
}

func (_ EventoWhOlaLanzadaWosPicking) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoWhOlaLanzadaWosPicking) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoWhOlaLanzadaWosPicking) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoWhOlaLanzadaWosPicking) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoWhOlaLanzadaWosPicking) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoWhOlaLanzadaWosPicking) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoWhOlaLanzadaWosPicking) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoWhOlaLanzadaWosPicking) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoWhOlaLanzadaWosPicking) Get(i int) types.Field {
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

func (r *EventoWhOlaLanzadaWosPicking) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventoWhOlaLanzadaWosPicking) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EventoWhOlaLanzadaWosPicking) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ EventoWhOlaLanzadaWosPicking) AppendArray() types.Field { panic("Unsupported operation") }
func (_ EventoWhOlaLanzadaWosPicking) HintSize(int)             { panic("Unsupported operation") }
func (_ EventoWhOlaLanzadaWosPicking) Finalize()                {}

func (_ EventoWhOlaLanzadaWosPicking) AvroCRC64Fingerprint() []byte {
	return []byte(EventoWhOlaLanzadaWosPickingAvroCRC64Fingerprint)
}

func (r EventoWhOlaLanzadaWosPicking) MarshalJSON() ([]byte, error) {
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

func (r *EventoWhOlaLanzadaWosPicking) UnmarshalJSON(data []byte) error {
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
