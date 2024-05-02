// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoUno.avsc
 */
package PruebaEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EventoUno struct {
	Cabecera *UnionNullCabecera `json:"Cabecera"`

	Detalle *UnionNullDetalle `json:"Detalle"`
}

const EventoUnoAvroCRC64Fingerprint = "k\xcfN\xea\xb6,:\xb4"

func NewEventoUno() EventoUno {
	r := EventoUno{}
	r.Cabecera = nil
	r.Detalle = nil
	return r
}

func DeserializeEventoUno(r io.Reader) (EventoUno, error) {
	t := NewEventoUno()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoUnoFromSchema(r io.Reader, schema string) (EventoUno, error) {
	t := NewEventoUno()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoUno(r EventoUno, w io.Writer) error {
	var err error
	err = writeUnionNullCabecera(r.Cabecera, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDetalle(r.Detalle, w)
	if err != nil {
		return err
	}
	return err
}

func (r EventoUno) Serialize(w io.Writer) error {
	return writeEventoUno(r, w)
}

func (r EventoUno) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Cabecera\",\"type\":[\"null\",{\"fields\":[{\"name\":\"idObligatorio\",\"type\":\"int\"},{\"name\":\"activoObligatorio\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"descripcionOpcional\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaDeNacimientoOpcional\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"Cabecera\",\"namespace\":\"Andreani.Prueba.Events.Common\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Detalle\",\"type\":[\"null\",{\"fields\":[{\"name\":\"id\",\"type\":\"int\"},{\"name\":\"otroDato\",\"type\":\"string\"}],\"name\":\"Detalle\",\"namespace\":\"Andreani.Prueba.Events.Common\",\"type\":\"record\"}]}],\"name\":\"Andreani.Prueba.Events.Record.EventoUno\",\"type\":\"record\"}"
}

func (r EventoUno) SchemaName() string {
	return "Andreani.Prueba.Events.Record.EventoUno"
}

func (_ EventoUno) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoUno) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoUno) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoUno) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoUno) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoUno) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoUno) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoUno) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoUno) Get(i int) types.Field {
	switch i {
	case 0:
		r.Cabecera = NewUnionNullCabecera()

		return r.Cabecera
	case 1:
		r.Detalle = NewUnionNullDetalle()

		return r.Detalle
	}
	panic("Unknown field index")
}

func (r *EventoUno) SetDefault(i int) {
	switch i {
	case 0:
		r.Cabecera = nil
		return
	case 1:
		r.Detalle = nil
		return
	}
	panic("Unknown field index")
}

func (r *EventoUno) NullField(i int) {
	switch i {
	case 0:
		r.Cabecera = nil
		return
	case 1:
		r.Detalle = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ EventoUno) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EventoUno) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EventoUno) HintSize(int)                     { panic("Unsupported operation") }
func (_ EventoUno) Finalize()                        {}

func (_ EventoUno) AvroCRC64Fingerprint() []byte {
	return []byte(EventoUnoAvroCRC64Fingerprint)
}

func (r EventoUno) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
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

func (r *EventoUno) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
		r.Cabecera = NewUnionNullCabecera()

		r.Cabecera = nil
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
		r.Detalle = NewUnionNullDetalle()

		r.Detalle = nil
	}
	return nil
}
