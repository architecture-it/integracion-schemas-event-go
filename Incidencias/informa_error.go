// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     InformaError.avsc
 */
package IncidenciasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type InformaError struct {
	IdIncidencia *UnionNullInt `json:"IdIncidencia"`

	EventoNombre *UnionNullString `json:"EventoNombre"`

	Error *UnionNullString `json:"Error"`

	IdentificadorExterno *UnionNullString `json:"IdentificadorExterno"`
}

const InformaErrorAvroCRC64Fingerprint = "n=LlI\xdb\xd6T"

func NewInformaError() InformaError {
	r := InformaError{}
	r.IdIncidencia = nil
	r.EventoNombre = nil
	r.Error = nil
	r.IdentificadorExterno = nil
	return r
}

func DeserializeInformaError(r io.Reader) (InformaError, error) {
	t := NewInformaError()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeInformaErrorFromSchema(r io.Reader, schema string) (InformaError, error) {
	t := NewInformaError()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeInformaError(r InformaError, w io.Writer) error {
	var err error
	err = writeUnionNullInt(r.IdIncidencia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EventoNombre, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Error, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.IdentificadorExterno, w)
	if err != nil {
		return err
	}
	return err
}

func (r InformaError) Serialize(w io.Writer) error {
	return writeInformaError(r, w)
}

func (r InformaError) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"IdIncidencia\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"EventoNombre\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Error\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"IdentificadorExterno\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.Incidencias.Events.Record.InformaError\",\"type\":\"record\"}"
}

func (r InformaError) SchemaName() string {
	return "Andreani.Incidencias.Events.Record.InformaError"
}

func (_ InformaError) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ InformaError) SetInt(v int32)       { panic("Unsupported operation") }
func (_ InformaError) SetLong(v int64)      { panic("Unsupported operation") }
func (_ InformaError) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ InformaError) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ InformaError) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ InformaError) SetString(v string)   { panic("Unsupported operation") }
func (_ InformaError) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *InformaError) Get(i int) types.Field {
	switch i {
	case 0:
		r.IdIncidencia = NewUnionNullInt()

		return r.IdIncidencia
	case 1:
		r.EventoNombre = NewUnionNullString()

		return r.EventoNombre
	case 2:
		r.Error = NewUnionNullString()

		return r.Error
	case 3:
		r.IdentificadorExterno = NewUnionNullString()

		return r.IdentificadorExterno
	}
	panic("Unknown field index")
}

func (r *InformaError) SetDefault(i int) {
	switch i {
	case 0:
		r.IdIncidencia = nil
		return
	case 1:
		r.EventoNombre = nil
		return
	case 2:
		r.Error = nil
		return
	case 3:
		r.IdentificadorExterno = nil
		return
	}
	panic("Unknown field index")
}

func (r *InformaError) NullField(i int) {
	switch i {
	case 0:
		r.IdIncidencia = nil
		return
	case 1:
		r.EventoNombre = nil
		return
	case 2:
		r.Error = nil
		return
	case 3:
		r.IdentificadorExterno = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ InformaError) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ InformaError) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ InformaError) HintSize(int)                     { panic("Unsupported operation") }
func (_ InformaError) Finalize()                        {}

func (_ InformaError) AvroCRC64Fingerprint() []byte {
	return []byte(InformaErrorAvroCRC64Fingerprint)
}

func (r InformaError) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["IdIncidencia"], err = json.Marshal(r.IdIncidencia)
	if err != nil {
		return nil, err
	}
	output["EventoNombre"], err = json.Marshal(r.EventoNombre)
	if err != nil {
		return nil, err
	}
	output["Error"], err = json.Marshal(r.Error)
	if err != nil {
		return nil, err
	}
	output["IdentificadorExterno"], err = json.Marshal(r.IdentificadorExterno)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *InformaError) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["IdIncidencia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdIncidencia); err != nil {
			return err
		}
	} else {
		r.IdIncidencia = NewUnionNullInt()

		r.IdIncidencia = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["EventoNombre"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EventoNombre); err != nil {
			return err
		}
	} else {
		r.EventoNombre = NewUnionNullString()

		r.EventoNombre = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Error"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Error); err != nil {
			return err
		}
	} else {
		r.Error = NewUnionNullString()

		r.Error = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["IdentificadorExterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdentificadorExterno); err != nil {
			return err
		}
	} else {
		r.IdentificadorExterno = NewUnionNullString()

		r.IdentificadorExterno = nil
	}
	return nil
}