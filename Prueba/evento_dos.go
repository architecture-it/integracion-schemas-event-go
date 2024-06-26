// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoDos.avsc
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

type EventoDos struct {
	Nombre string `json:"nombre"`

	Edad int32 `json:"edad"`

	Cabecera *UnionNullCabecera `json:"Cabecera"`
}

const EventoDosAvroCRC64Fingerprint = "/\xef\xfc\x1c\xbe\xfe~\x91"

func NewEventoDos() EventoDos {
	r := EventoDos{}
	r.Cabecera = nil
	return r
}

func DeserializeEventoDos(r io.Reader) (EventoDos, error) {
	t := NewEventoDos()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoDosFromSchema(r io.Reader, schema string) (EventoDos, error) {
	t := NewEventoDos()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoDos(r EventoDos, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Nombre, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Edad, w)
	if err != nil {
		return err
	}
	err = writeUnionNullCabecera(r.Cabecera, w)
	if err != nil {
		return err
	}
	return err
}

func (r EventoDos) Serialize(w io.Writer) error {
	return writeEventoDos(r, w)
}

func (r EventoDos) Schema() string {
	return "{\"fields\":[{\"name\":\"nombre\",\"type\":\"string\"},{\"name\":\"edad\",\"type\":\"int\"},{\"default\":null,\"name\":\"Cabecera\",\"type\":[\"null\",{\"fields\":[{\"name\":\"idObligatorio\",\"type\":\"int\"},{\"name\":\"activoObligatorio\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"descripcionOpcional\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaDeNacimientoOpcional\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"Cabecera\",\"namespace\":\"Andreani.Prueba.Events.Common\",\"type\":\"record\"}]}],\"name\":\"Andreani.Prueba.Events.Record.EventoDos\",\"type\":\"record\"}"
}

func (r EventoDos) SchemaName() string {
	return "Andreani.Prueba.Events.Record.EventoDos"
}

func (_ EventoDos) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoDos) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoDos) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoDos) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoDos) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoDos) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoDos) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoDos) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoDos) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Nombre}

		return w

	case 1:
		w := types.Int{Target: &r.Edad}

		return w

	case 2:
		r.Cabecera = NewUnionNullCabecera()

		return r.Cabecera
	}
	panic("Unknown field index")
}

func (r *EventoDos) SetDefault(i int) {
	switch i {
	case 2:
		r.Cabecera = nil
		return
	}
	panic("Unknown field index")
}

func (r *EventoDos) NullField(i int) {
	switch i {
	case 2:
		r.Cabecera = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ EventoDos) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EventoDos) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EventoDos) HintSize(int)                     { panic("Unsupported operation") }
func (_ EventoDos) Finalize()                        {}

func (_ EventoDos) AvroCRC64Fingerprint() []byte {
	return []byte(EventoDosAvroCRC64Fingerprint)
}

func (r EventoDos) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["nombre"], err = json.Marshal(r.Nombre)
	if err != nil {
		return nil, err
	}
	output["edad"], err = json.Marshal(r.Edad)
	if err != nil {
		return nil, err
	}
	output["Cabecera"], err = json.Marshal(r.Cabecera)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EventoDos) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["nombre"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Nombre); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for nombre")
	}
	val = func() json.RawMessage {
		if v, ok := fields["edad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Edad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for edad")
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
		r.Cabecera = NewUnionNullCabecera()

		r.Cabecera = nil
	}
	return nil
}
