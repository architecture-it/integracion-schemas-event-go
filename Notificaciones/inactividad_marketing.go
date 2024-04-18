// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     InactividadMarketing.avsc
 */
package NotificacionesEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type InactividadMarketing struct {
	ClienteId string `json:"ClienteId"`

	Email string `json:"Email"`

	Nombre *UnionNullString `json:"Nombre"`

	TelefonoCodigoArea *UnionNullString `json:"TelefonoCodigoArea"`

	TelefonoNumero *UnionNullString `json:"TelefonoNumero"`

	Segmento string `json:"segmento"`

	Mes int64 `json:"Mes"`
}

const InactividadMarketingAvroCRC64Fingerprint = "\n\xec\x1fq6\x169\x1d"

func NewInactividadMarketing() InactividadMarketing {
	r := InactividadMarketing{}
	r.Nombre = nil
	r.TelefonoCodigoArea = nil
	r.TelefonoNumero = nil
	return r
}

func DeserializeInactividadMarketing(r io.Reader) (InactividadMarketing, error) {
	t := NewInactividadMarketing()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeInactividadMarketingFromSchema(r io.Reader, schema string) (InactividadMarketing, error) {
	t := NewInactividadMarketing()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeInactividadMarketing(r InactividadMarketing, w io.Writer) error {
	var err error
	err = vm.WriteString(r.ClienteId, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Email, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Nombre, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TelefonoCodigoArea, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TelefonoNumero, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Segmento, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Mes, w)
	if err != nil {
		return err
	}
	return err
}

func (r InactividadMarketing) Serialize(w io.Writer) error {
	return writeInactividadMarketing(r, w)
}

func (r InactividadMarketing) Schema() string {
	return "{\"fields\":[{\"name\":\"ClienteId\",\"type\":\"string\"},{\"name\":\"Email\",\"type\":\"string\"},{\"default\":null,\"name\":\"Nombre\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TelefonoCodigoArea\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TelefonoNumero\",\"type\":[\"null\",\"string\"]},{\"name\":\"segmento\",\"type\":\"string\"},{\"name\":\"Mes\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}}],\"name\":\"Andreani.Notificaciones.Events.Records.InactividadMarketing\",\"type\":\"record\"}"
}

func (r InactividadMarketing) SchemaName() string {
	return "Andreani.Notificaciones.Events.Records.InactividadMarketing"
}

func (_ InactividadMarketing) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ InactividadMarketing) SetInt(v int32)       { panic("Unsupported operation") }
func (_ InactividadMarketing) SetLong(v int64)      { panic("Unsupported operation") }
func (_ InactividadMarketing) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ InactividadMarketing) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ InactividadMarketing) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ InactividadMarketing) SetString(v string)   { panic("Unsupported operation") }
func (_ InactividadMarketing) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *InactividadMarketing) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.ClienteId}

		return w

	case 1:
		w := types.String{Target: &r.Email}

		return w

	case 2:
		r.Nombre = NewUnionNullString()

		return r.Nombre
	case 3:
		r.TelefonoCodigoArea = NewUnionNullString()

		return r.TelefonoCodigoArea
	case 4:
		r.TelefonoNumero = NewUnionNullString()

		return r.TelefonoNumero
	case 5:
		w := types.String{Target: &r.Segmento}

		return w

	case 6:
		w := types.Long{Target: &r.Mes}

		return w

	}
	panic("Unknown field index")
}

func (r *InactividadMarketing) SetDefault(i int) {
	switch i {
	case 2:
		r.Nombre = nil
		return
	case 3:
		r.TelefonoCodigoArea = nil
		return
	case 4:
		r.TelefonoNumero = nil
		return
	}
	panic("Unknown field index")
}

func (r *InactividadMarketing) NullField(i int) {
	switch i {
	case 2:
		r.Nombre = nil
		return
	case 3:
		r.TelefonoCodigoArea = nil
		return
	case 4:
		r.TelefonoNumero = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ InactividadMarketing) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ InactividadMarketing) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ InactividadMarketing) HintSize(int)                     { panic("Unsupported operation") }
func (_ InactividadMarketing) Finalize()                        {}

func (_ InactividadMarketing) AvroCRC64Fingerprint() []byte {
	return []byte(InactividadMarketingAvroCRC64Fingerprint)
}

func (r InactividadMarketing) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["ClienteId"], err = json.Marshal(r.ClienteId)
	if err != nil {
		return nil, err
	}
	output["Email"], err = json.Marshal(r.Email)
	if err != nil {
		return nil, err
	}
	output["Nombre"], err = json.Marshal(r.Nombre)
	if err != nil {
		return nil, err
	}
	output["TelefonoCodigoArea"], err = json.Marshal(r.TelefonoCodigoArea)
	if err != nil {
		return nil, err
	}
	output["TelefonoNumero"], err = json.Marshal(r.TelefonoNumero)
	if err != nil {
		return nil, err
	}
	output["segmento"], err = json.Marshal(r.Segmento)
	if err != nil {
		return nil, err
	}
	output["Mes"], err = json.Marshal(r.Mes)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *InactividadMarketing) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["ClienteId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ClienteId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ClienteId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Email"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Email); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Email")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Nombre"]; ok {
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
		if v, ok := fields["TelefonoCodigoArea"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TelefonoCodigoArea); err != nil {
			return err
		}
	} else {
		r.TelefonoCodigoArea = NewUnionNullString()

		r.TelefonoCodigoArea = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["TelefonoNumero"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TelefonoNumero); err != nil {
			return err
		}
	} else {
		r.TelefonoNumero = NewUnionNullString()

		r.TelefonoNumero = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["segmento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Segmento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for segmento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Mes"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Mes); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Mes")
	}
	return nil
}
