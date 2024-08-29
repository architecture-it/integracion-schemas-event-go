// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     MDA.avsc
 */
package EAMEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type MDA struct {
	Ticket Ticket `json:"Ticket"`
}

const MDAAvroCRC64Fingerprint = "\xe3\x0fY&\xe4<,\xca"

func NewMDA() MDA {
	r := MDA{}
	r.Ticket = NewTicket()

	return r
}

func DeserializeMDA(r io.Reader) (MDA, error) {
	t := NewMDA()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeMDAFromSchema(r io.Reader, schema string) (MDA, error) {
	t := NewMDA()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeMDA(r MDA, w io.Writer) error {
	var err error
	err = writeTicket(r.Ticket, w)
	if err != nil {
		return err
	}
	return err
}

func (r MDA) Serialize(w io.Writer) error {
	return writeMDA(r, w)
}

func (r MDA) Schema() string {
	return "{\"fields\":[{\"name\":\"Ticket\",\"type\":{\"fields\":[{\"name\":\"IdTicket\",\"type\":\"int\"},{\"name\":\"Grupo\",\"type\":\"string\"},{\"name\":\"Categoria\",\"type\":\"string\"},{\"name\":\"Articulo\",\"type\":\"string\"},{\"name\":\"Solicitante\",\"type\":\"string\"},{\"name\":\"Asunto\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"name\":\"Tecnico\",\"type\":\"string\"},{\"name\":\"Cliente\",\"type\":\"string\"},{\"name\":\"Planta\",\"type\":\"string\"},{\"name\":\"Nave\",\"type\":\"string\"},{\"name\":\"Sector\",\"type\":\"string\"},{\"name\":\"Sectores\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"Procesos\",\"type\":\"string\"},{\"name\":\"Etiqueta\",\"type\":\"string\"},{\"name\":\"Sucursal\",\"type\":\"string\"},{\"name\":\"Sucursales\",\"type\":\"string\"},{\"name\":\"FechaCreacion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"FechaVencimiento\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"FechaResolucion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"FechaFinalizado\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"Estado\",\"type\":\"string\"}],\"name\":\"Ticket\",\"namespace\":\"Andreani.EAM.Events.MDA\",\"type\":\"record\"}}],\"name\":\"Andreani.EAM.Events.Record.MDA\",\"type\":\"record\"}"
}

func (r MDA) SchemaName() string {
	return "Andreani.EAM.Events.Record.MDA"
}

func (_ MDA) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ MDA) SetInt(v int32)       { panic("Unsupported operation") }
func (_ MDA) SetLong(v int64)      { panic("Unsupported operation") }
func (_ MDA) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ MDA) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ MDA) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ MDA) SetString(v string)   { panic("Unsupported operation") }
func (_ MDA) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MDA) Get(i int) types.Field {
	switch i {
	case 0:
		r.Ticket = NewTicket()

		w := types.Record{Target: &r.Ticket}

		return w

	}
	panic("Unknown field index")
}

func (r *MDA) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *MDA) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ MDA) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ MDA) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ MDA) HintSize(int)                     { panic("Unsupported operation") }
func (_ MDA) Finalize()                        {}

func (_ MDA) AvroCRC64Fingerprint() []byte {
	return []byte(MDAAvroCRC64Fingerprint)
}

func (r MDA) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Ticket"], err = json.Marshal(r.Ticket)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *MDA) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Ticket"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ticket); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Ticket")
	}
	return nil
}