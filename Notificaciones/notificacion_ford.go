// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NotificacionFord.avsc
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

type NotificacionFord struct {
	Evento *UnionNullString `json:"evento"`

	CodigoDestinatario *UnionNullString `json:"codigoDestinatario"`

	InternalCliente *UnionNullString `json:"internalCliente"`

	Contrato *UnionNullString `json:"contrato"`

	MailDestino *UnionNullString `json:"mailDestino"`

	NombreDestinatario *UnionNullString `json:"nombreDestinatario"`

	CodigoPostalDestino *UnionNullLong `json:"codigoPostalDestino"`

	Eda *UnionNullLong `json:"eda"`

	Concesionario *UnionNullMaestroConcesionario `json:"concesionario"`

	Detalles []DetalleFord `json:"detalles"`
}

const NotificacionFordAvroCRC64Fingerprint = "ؐ\xa9F\xbc\xact\xae"

func NewNotificacionFord() NotificacionFord {
	r := NotificacionFord{}
	r.Evento = nil
	r.CodigoDestinatario = nil
	r.InternalCliente = nil
	r.Contrato = nil
	r.MailDestino = nil
	r.NombreDestinatario = nil
	r.CodigoPostalDestino = nil
	r.Eda = nil
	r.Concesionario = nil
	r.Detalles = make([]DetalleFord, 0)

	return r
}

func DeserializeNotificacionFord(r io.Reader) (NotificacionFord, error) {
	t := NewNotificacionFord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNotificacionFordFromSchema(r io.Reader, schema string) (NotificacionFord, error) {
	t := NewNotificacionFord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNotificacionFord(r NotificacionFord, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Evento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoDestinatario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.InternalCliente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Contrato, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.MailDestino, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NombreDestinatario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.CodigoPostalDestino, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.Eda, w)
	if err != nil {
		return err
	}
	err = writeUnionNullMaestroConcesionario(r.Concesionario, w)
	if err != nil {
		return err
	}
	err = writeArrayDetalleFord(r.Detalles, w)
	if err != nil {
		return err
	}
	return err
}

func (r NotificacionFord) Serialize(w io.Writer) error {
	return writeNotificacionFord(r, w)
}

func (r NotificacionFord) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"evento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoDestinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"internalCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"mailDestino\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreDestinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostalDestino\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"eda\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"concesionario\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"id_concesionario\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"razon_social\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"direccion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"provincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cp_origen\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"cp_destino\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"codigo_cliente\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"contrato\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"sla\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"corte_horario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"email_1\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"email_2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"email_3\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"email_4\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"email_5\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"email_6\",\"type\":[\"null\",\"string\"]}],\"name\":\"MaestroConcesionario\",\"type\":\"record\"}]},{\"name\":\"detalles\",\"type\":{\"items\":{\"fields\":[{\"default\":null,\"name\":\"codigoDeEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"notaDespacho\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"bultos\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"fechaAdmision\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"fechaInsercion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"DetalleFord\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.Notificaciones.Events.Records.NotificacionFord\",\"type\":\"record\"}"
}

func (r NotificacionFord) SchemaName() string {
	return "Andreani.Notificaciones.Events.Records.NotificacionFord"
}

func (_ NotificacionFord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NotificacionFord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NotificacionFord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NotificacionFord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NotificacionFord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NotificacionFord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NotificacionFord) SetString(v string)   { panic("Unsupported operation") }
func (_ NotificacionFord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NotificacionFord) Get(i int) types.Field {
	switch i {
	case 0:
		r.Evento = NewUnionNullString()

		return r.Evento
	case 1:
		r.CodigoDestinatario = NewUnionNullString()

		return r.CodigoDestinatario
	case 2:
		r.InternalCliente = NewUnionNullString()

		return r.InternalCliente
	case 3:
		r.Contrato = NewUnionNullString()

		return r.Contrato
	case 4:
		r.MailDestino = NewUnionNullString()

		return r.MailDestino
	case 5:
		r.NombreDestinatario = NewUnionNullString()

		return r.NombreDestinatario
	case 6:
		r.CodigoPostalDestino = NewUnionNullLong()

		return r.CodigoPostalDestino
	case 7:
		r.Eda = NewUnionNullLong()

		return r.Eda
	case 8:
		r.Concesionario = NewUnionNullMaestroConcesionario()

		return r.Concesionario
	case 9:
		r.Detalles = make([]DetalleFord, 0)

		w := ArrayDetalleFordWrapper{Target: &r.Detalles}

		return w

	}
	panic("Unknown field index")
}

func (r *NotificacionFord) SetDefault(i int) {
	switch i {
	case 0:
		r.Evento = nil
		return
	case 1:
		r.CodigoDestinatario = nil
		return
	case 2:
		r.InternalCliente = nil
		return
	case 3:
		r.Contrato = nil
		return
	case 4:
		r.MailDestino = nil
		return
	case 5:
		r.NombreDestinatario = nil
		return
	case 6:
		r.CodigoPostalDestino = nil
		return
	case 7:
		r.Eda = nil
		return
	case 8:
		r.Concesionario = nil
		return
	}
	panic("Unknown field index")
}

func (r *NotificacionFord) NullField(i int) {
	switch i {
	case 0:
		r.Evento = nil
		return
	case 1:
		r.CodigoDestinatario = nil
		return
	case 2:
		r.InternalCliente = nil
		return
	case 3:
		r.Contrato = nil
		return
	case 4:
		r.MailDestino = nil
		return
	case 5:
		r.NombreDestinatario = nil
		return
	case 6:
		r.CodigoPostalDestino = nil
		return
	case 7:
		r.Eda = nil
		return
	case 8:
		r.Concesionario = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ NotificacionFord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NotificacionFord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NotificacionFord) HintSize(int)                     { panic("Unsupported operation") }
func (_ NotificacionFord) Finalize()                        {}

func (_ NotificacionFord) AvroCRC64Fingerprint() []byte {
	return []byte(NotificacionFordAvroCRC64Fingerprint)
}

func (r NotificacionFord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["evento"], err = json.Marshal(r.Evento)
	if err != nil {
		return nil, err
	}
	output["codigoDestinatario"], err = json.Marshal(r.CodigoDestinatario)
	if err != nil {
		return nil, err
	}
	output["internalCliente"], err = json.Marshal(r.InternalCliente)
	if err != nil {
		return nil, err
	}
	output["contrato"], err = json.Marshal(r.Contrato)
	if err != nil {
		return nil, err
	}
	output["mailDestino"], err = json.Marshal(r.MailDestino)
	if err != nil {
		return nil, err
	}
	output["nombreDestinatario"], err = json.Marshal(r.NombreDestinatario)
	if err != nil {
		return nil, err
	}
	output["codigoPostalDestino"], err = json.Marshal(r.CodigoPostalDestino)
	if err != nil {
		return nil, err
	}
	output["eda"], err = json.Marshal(r.Eda)
	if err != nil {
		return nil, err
	}
	output["concesionario"], err = json.Marshal(r.Concesionario)
	if err != nil {
		return nil, err
	}
	output["detalles"], err = json.Marshal(r.Detalles)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NotificacionFord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["evento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Evento); err != nil {
			return err
		}
	} else {
		r.Evento = NewUnionNullString()

		r.Evento = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoDestinatario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDestinatario); err != nil {
			return err
		}
	} else {
		r.CodigoDestinatario = NewUnionNullString()

		r.CodigoDestinatario = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["internalCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.InternalCliente); err != nil {
			return err
		}
	} else {
		r.InternalCliente = NewUnionNullString()

		r.InternalCliente = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["contrato"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contrato); err != nil {
			return err
		}
	} else {
		r.Contrato = NewUnionNullString()

		r.Contrato = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["mailDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MailDestino); err != nil {
			return err
		}
	} else {
		r.MailDestino = NewUnionNullString()

		r.MailDestino = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["nombreDestinatario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NombreDestinatario); err != nil {
			return err
		}
	} else {
		r.NombreDestinatario = NewUnionNullString()

		r.NombreDestinatario = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoPostalDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoPostalDestino); err != nil {
			return err
		}
	} else {
		r.CodigoPostalDestino = NewUnionNullLong()

		r.CodigoPostalDestino = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["eda"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Eda); err != nil {
			return err
		}
	} else {
		r.Eda = NewUnionNullLong()

		r.Eda = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["concesionario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Concesionario); err != nil {
			return err
		}
	} else {
		r.Concesionario = NewUnionNullMaestroConcesionario()

		r.Concesionario = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["detalles"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Detalles); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for detalles")
	}
	return nil
}
