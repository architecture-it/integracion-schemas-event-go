// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NotificacionWh.avsc
 */
package NotificacionesWhEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Notificacion struct {
	Destinatarios string `json:"Destinatarios"`

	Copia *UnionNullString `json:"Copia"`

	CopiaOculta *UnionNullString `json:"CopiaOculta"`

	Asunto string `json:"Asunto"`

	Encabezado *UnionNullString `json:"Encabezado"`

	Cuerpo string `json:"Cuerpo"`

	Pie *UnionNullString `json:"Pie"`
}

const NotificacionAvroCRC64Fingerprint = "\b裪P\x88^t"

func NewNotificacion() Notificacion {
	r := Notificacion{}
	return r
}

func DeserializeNotificacion(r io.Reader) (Notificacion, error) {
	t := NewNotificacion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNotificacionFromSchema(r io.Reader, schema string) (Notificacion, error) {
	t := NewNotificacion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNotificacion(r Notificacion, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Destinatarios, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Copia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CopiaOculta, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Asunto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Encabezado, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Cuerpo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Pie, w)
	if err != nil {
		return err
	}
	return err
}

func (r Notificacion) Serialize(w io.Writer) error {
	return writeNotificacion(r, w)
}

func (r Notificacion) Schema() string {
	return "{\"fields\":[{\"name\":\"Destinatarios\",\"type\":\"string\"},{\"name\":\"Copia\",\"type\":[\"null\",\"string\"]},{\"name\":\"CopiaOculta\",\"type\":[\"null\",\"string\"]},{\"name\":\"Asunto\",\"type\":\"string\"},{\"name\":\"Encabezado\",\"type\":[\"null\",\"string\"]},{\"name\":\"Cuerpo\",\"type\":\"string\"},{\"name\":\"Pie\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.NotificacionesWh.Events.Common.Notificacion\",\"type\":\"record\"}"
}

func (r Notificacion) SchemaName() string {
	return "Andreani.NotificacionesWh.Events.Common.Notificacion"
}

func (_ Notificacion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Notificacion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Notificacion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Notificacion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Notificacion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Notificacion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Notificacion) SetString(v string)   { panic("Unsupported operation") }
func (_ Notificacion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Notificacion) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Destinatarios}

		return w

	case 1:
		r.Copia = NewUnionNullString()

		return r.Copia
	case 2:
		r.CopiaOculta = NewUnionNullString()

		return r.CopiaOculta
	case 3:
		w := types.String{Target: &r.Asunto}

		return w

	case 4:
		r.Encabezado = NewUnionNullString()

		return r.Encabezado
	case 5:
		w := types.String{Target: &r.Cuerpo}

		return w

	case 6:
		r.Pie = NewUnionNullString()

		return r.Pie
	}
	panic("Unknown field index")
}

func (r *Notificacion) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Notificacion) NullField(i int) {
	switch i {
	case 1:
		r.Copia = nil
		return
	case 2:
		r.CopiaOculta = nil
		return
	case 4:
		r.Encabezado = nil
		return
	case 6:
		r.Pie = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Notificacion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Notificacion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Notificacion) HintSize(int)                     { panic("Unsupported operation") }
func (_ Notificacion) Finalize()                        {}

func (_ Notificacion) AvroCRC64Fingerprint() []byte {
	return []byte(NotificacionAvroCRC64Fingerprint)
}

func (r Notificacion) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Destinatarios"], err = json.Marshal(r.Destinatarios)
	if err != nil {
		return nil, err
	}
	output["Copia"], err = json.Marshal(r.Copia)
	if err != nil {
		return nil, err
	}
	output["CopiaOculta"], err = json.Marshal(r.CopiaOculta)
	if err != nil {
		return nil, err
	}
	output["Asunto"], err = json.Marshal(r.Asunto)
	if err != nil {
		return nil, err
	}
	output["Encabezado"], err = json.Marshal(r.Encabezado)
	if err != nil {
		return nil, err
	}
	output["Cuerpo"], err = json.Marshal(r.Cuerpo)
	if err != nil {
		return nil, err
	}
	output["Pie"], err = json.Marshal(r.Pie)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Notificacion) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Destinatarios"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Destinatarios); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Destinatarios")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Copia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Copia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Copia")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CopiaOculta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CopiaOculta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CopiaOculta")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Asunto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Asunto); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Asunto")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Encabezado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Encabezado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Encabezado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Cuerpo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cuerpo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Cuerpo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Pie"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Pie); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Pie")
	}
	return nil
}
