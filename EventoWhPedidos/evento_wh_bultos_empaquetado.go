// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhBultosEmpaquetado.avsc
 */
package EventoWhPedidosEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EventoWhBultosEmpaquetado struct {
	Identificacion Identificacion `json:"Identificacion"`

	Cabecera Cabecera `json:"Cabecera"`

	ContenedorRetornable *UnionNullContenedorRetornable `json:"ContenedorRetornable"`
}

const EventoWhBultosEmpaquetadoAvroCRC64Fingerprint = "V\xbfr\xf1\xf1\xe8\x83\xc4"

func NewEventoWhBultosEmpaquetado() EventoWhBultosEmpaquetado {
	r := EventoWhBultosEmpaquetado{}
	r.Identificacion = NewIdentificacion()

	r.Cabecera = NewCabecera()

	return r
}

func DeserializeEventoWhBultosEmpaquetado(r io.Reader) (EventoWhBultosEmpaquetado, error) {
	t := NewEventoWhBultosEmpaquetado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoWhBultosEmpaquetadoFromSchema(r io.Reader, schema string) (EventoWhBultosEmpaquetado, error) {
	t := NewEventoWhBultosEmpaquetado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoWhBultosEmpaquetado(r EventoWhBultosEmpaquetado, w io.Writer) error {
	var err error
	err = writeIdentificacion(r.Identificacion, w)
	if err != nil {
		return err
	}
	err = writeCabecera(r.Cabecera, w)
	if err != nil {
		return err
	}
	err = writeUnionNullContenedorRetornable(r.ContenedorRetornable, w)
	if err != nil {
		return err
	}
	return err
}

func (r EventoWhBultosEmpaquetado) Serialize(w io.Writer) error {
	return writeEventoWhBultosEmpaquetado(r, w)
}

func (r EventoWhBultosEmpaquetado) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.EventoWhPedidos.Events.BultosEmpaquetadoCommon\",\"type\":\"record\"}},{\"name\":\"Cabecera\",\"type\":{\"fields\":[{\"name\":\"OrdenWH\",\"type\":\"string\"},{\"name\":\"ContratoTMS\",\"type\":\"string\"},{\"name\":\"Lpn\",\"type\":\"string\"},{\"name\":\"BultosTotal\",\"type\":\"int\"},{\"name\":\"EXT_UDF_STR1\",\"type\":[\"null\",\"string\"]}],\"name\":\"Cabecera\",\"namespace\":\"Andreani.EventoWhPedidos.Events.BultosEmpaquetadoCommon\",\"type\":\"record\"}},{\"name\":\"ContenedorRetornable\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Tara\",\"type\":\"float\"},{\"name\":\"CartonType\",\"type\":\"string\"},{\"name\":\"ContratoRetornable\",\"type\":[\"null\",\"string\"]}],\"name\":\"ContenedorRetornable\",\"namespace\":\"Andreani.EventoWhPedidos.Events.BultosEmpaquetadoCommon\",\"type\":\"record\"}]}],\"name\":\"Andreani.EventoWhPedidos.Events.Record.EventoWhBultosEmpaquetado\",\"type\":\"record\"}"
}

func (r EventoWhBultosEmpaquetado) SchemaName() string {
	return "Andreani.EventoWhPedidos.Events.Record.EventoWhBultosEmpaquetado"
}

func (_ EventoWhBultosEmpaquetado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoWhBultosEmpaquetado) Get(i int) types.Field {
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
		r.ContenedorRetornable = NewUnionNullContenedorRetornable()

		return r.ContenedorRetornable
	}
	panic("Unknown field index")
}

func (r *EventoWhBultosEmpaquetado) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventoWhBultosEmpaquetado) NullField(i int) {
	switch i {
	case 2:
		r.ContenedorRetornable = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ EventoWhBultosEmpaquetado) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) HintSize(int)                     { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) Finalize()                        {}

func (_ EventoWhBultosEmpaquetado) AvroCRC64Fingerprint() []byte {
	return []byte(EventoWhBultosEmpaquetadoAvroCRC64Fingerprint)
}

func (r EventoWhBultosEmpaquetado) MarshalJSON() ([]byte, error) {
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
	output["ContenedorRetornable"], err = json.Marshal(r.ContenedorRetornable)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EventoWhBultosEmpaquetado) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["ContenedorRetornable"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ContenedorRetornable); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ContenedorRetornable")
	}
	return nil
}
