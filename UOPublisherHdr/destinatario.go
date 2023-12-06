// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     UOHdrCreada.avsc
 */
package UOPublisherHdrEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Destinatario struct {
	DestinatarioId int32 `json:"DestinatarioId"`

	Nombre *UnionNullString `json:"nombre"`

	Apellido *UnionNullString `json:"apellido"`

	Dni *UnionNullString `json:"dni"`

	Telefono *UnionNullString `json:"telefono"`

	Celular *UnionNullString `json:"celular"`

	Email *UnionNullString `json:"email"`
}

const DestinatarioAvroCRC64Fingerprint = "\x90 (\xcb\xfb\x11Y\xa4"

func NewDestinatario() Destinatario {
	r := Destinatario{}
	r.Nombre = nil
	r.Apellido = nil
	r.Dni = nil
	r.Telefono = nil
	r.Celular = nil
	r.Email = nil
	return r
}

func DeserializeDestinatario(r io.Reader) (Destinatario, error) {
	t := NewDestinatario()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDestinatarioFromSchema(r io.Reader, schema string) (Destinatario, error) {
	t := NewDestinatario()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDestinatario(r Destinatario, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.DestinatarioId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Nombre, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Apellido, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Dni, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Telefono, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Celular, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Email, w)
	if err != nil {
		return err
	}
	return err
}

func (r Destinatario) Serialize(w io.Writer) error {
	return writeDestinatario(r, w)
}

func (r Destinatario) Schema() string {
	return "{\"fields\":[{\"name\":\"DestinatarioId\",\"type\":\"int\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"apellido\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"dni\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefono\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"celular\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"email\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.UOPublisherHdr.Events.Common.Destinatario\",\"type\":\"record\"}"
}

func (r Destinatario) SchemaName() string {
	return "Andreani.UOPublisherHdr.Events.Common.Destinatario"
}

func (_ Destinatario) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Destinatario) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Destinatario) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Destinatario) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Destinatario) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Destinatario) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Destinatario) SetString(v string)   { panic("Unsupported operation") }
func (_ Destinatario) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Destinatario) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.DestinatarioId}

		return w

	case 1:
		r.Nombre = NewUnionNullString()

		return r.Nombre
	case 2:
		r.Apellido = NewUnionNullString()

		return r.Apellido
	case 3:
		r.Dni = NewUnionNullString()

		return r.Dni
	case 4:
		r.Telefono = NewUnionNullString()

		return r.Telefono
	case 5:
		r.Celular = NewUnionNullString()

		return r.Celular
	case 6:
		r.Email = NewUnionNullString()

		return r.Email
	}
	panic("Unknown field index")
}

func (r *Destinatario) SetDefault(i int) {
	switch i {
	case 1:
		r.Nombre = nil
		return
	case 2:
		r.Apellido = nil
		return
	case 3:
		r.Dni = nil
		return
	case 4:
		r.Telefono = nil
		return
	case 5:
		r.Celular = nil
		return
	case 6:
		r.Email = nil
		return
	}
	panic("Unknown field index")
}

func (r *Destinatario) NullField(i int) {
	switch i {
	case 1:
		r.Nombre = nil
		return
	case 2:
		r.Apellido = nil
		return
	case 3:
		r.Dni = nil
		return
	case 4:
		r.Telefono = nil
		return
	case 5:
		r.Celular = nil
		return
	case 6:
		r.Email = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Destinatario) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Destinatario) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Destinatario) HintSize(int)                     { panic("Unsupported operation") }
func (_ Destinatario) Finalize()                        {}

func (_ Destinatario) AvroCRC64Fingerprint() []byte {
	return []byte(DestinatarioAvroCRC64Fingerprint)
}

func (r Destinatario) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["DestinatarioId"], err = json.Marshal(r.DestinatarioId)
	if err != nil {
		return nil, err
	}
	output["nombre"], err = json.Marshal(r.Nombre)
	if err != nil {
		return nil, err
	}
	output["apellido"], err = json.Marshal(r.Apellido)
	if err != nil {
		return nil, err
	}
	output["dni"], err = json.Marshal(r.Dni)
	if err != nil {
		return nil, err
	}
	output["telefono"], err = json.Marshal(r.Telefono)
	if err != nil {
		return nil, err
	}
	output["celular"], err = json.Marshal(r.Celular)
	if err != nil {
		return nil, err
	}
	output["email"], err = json.Marshal(r.Email)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Destinatario) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioId")
	}
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
		r.Nombre = NewUnionNullString()

		r.Nombre = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["apellido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Apellido); err != nil {
			return err
		}
	} else {
		r.Apellido = NewUnionNullString()

		r.Apellido = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["dni"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Dni); err != nil {
			return err
		}
	} else {
		r.Dni = NewUnionNullString()

		r.Dni = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["telefono"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Telefono); err != nil {
			return err
		}
	} else {
		r.Telefono = NewUnionNullString()

		r.Telefono = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["celular"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Celular); err != nil {
			return err
		}
	} else {
		r.Celular = NewUnionNullString()

		r.Celular = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["email"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Email); err != nil {
			return err
		}
	} else {
		r.Email = NewUnionNullString()

		r.Email = nil
	}
	return nil
}