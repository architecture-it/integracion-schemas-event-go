// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Notificaciones.avsc
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

type Notificaciones struct {
	IdModelo int64 `json:"idModelo"`

	FechaNotificacion int64 `json:"fechaNotificacion"`

	FechaEvento int64 `json:"fechaEvento"`

	IdSistema int32 `json:"idSistema"`

	Sistema string `json:"sistema"`

	Contrato string `json:"contrato"`

	Envio string `json:"envio"`

	IdEvento int32 `json:"idEvento"`

	Evento string `json:"evento"`

	IdMotivo *UnionNullInt `json:"idMotivo"`

	Motivo *UnionNullString `json:"motivo"`

	IdSalida int32 `json:"idSalida"`

	Salida string `json:"salida"`

	UrlSalida string `json:"urlSalida"`

	IdRegla *UnionNullInt `json:"idRegla"`

	Regla *UnionNullString `json:"Regla"`

	Destinatario *UnionNullString `json:"destinatario"`

	DestinatarioNotificacion *UnionNullString `json:"destinatarioNotificacion"`

	DestinatarioEstado *UnionNullString `json:"destinatarioEstado"`

	DestinatarioObservacion *UnionNullString `json:"destinatarioObservacion"`

	DestinatarioNotificacionCaracteres *UnionNullInt `json:"destinatarioNotificacionCaracteres"`

	Remitente *UnionNullString `json:"remitente"`

	RemitenteNotificacion *UnionNullString `json:"remitenteNotificacion"`

	RemitenteEstado *UnionNullString `json:"remitenteEstado"`

	RemitenteObservacion *UnionNullString `json:"remitenteObservacion"`

	RemitenteNotificacionCaracteres *UnionNullInt `json:"remitenteNotificacionCaracteres"`

	ProveedorSMS *UnionNullString `json:"proveedorSMS"`
}

const NotificacionesAvroCRC64Fingerprint = "\xb9\xadqFf\x04\xfc\xf4"

func NewNotificaciones() Notificaciones {
	r := Notificaciones{}
	r.IdMotivo = nil
	r.Motivo = nil
	r.IdRegla = nil
	r.Regla = nil
	r.Destinatario = nil
	r.DestinatarioNotificacion = nil
	r.DestinatarioEstado = nil
	r.DestinatarioObservacion = nil
	r.DestinatarioNotificacionCaracteres = nil
	r.Remitente = nil
	r.RemitenteNotificacion = nil
	r.RemitenteEstado = nil
	r.RemitenteObservacion = nil
	r.RemitenteNotificacionCaracteres = nil
	r.ProveedorSMS = nil
	return r
}

func DeserializeNotificaciones(r io.Reader) (Notificaciones, error) {
	t := NewNotificaciones()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNotificacionesFromSchema(r io.Reader, schema string) (Notificaciones, error) {
	t := NewNotificaciones()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNotificaciones(r Notificaciones, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.IdModelo, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.FechaNotificacion, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.FechaEvento, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.IdSistema, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Sistema, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Contrato, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Envio, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.IdEvento, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Evento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.IdMotivo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Motivo, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.IdSalida, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Salida, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.UrlSalida, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.IdRegla, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Regla, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Destinatario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioNotificacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioEstado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioObservacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.DestinatarioNotificacionCaracteres, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Remitente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.RemitenteNotificacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.RemitenteEstado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.RemitenteObservacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.RemitenteNotificacionCaracteres, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ProveedorSMS, w)
	if err != nil {
		return err
	}
	return err
}

func (r Notificaciones) Serialize(w io.Writer) error {
	return writeNotificaciones(r, w)
}

func (r Notificaciones) Schema() string {
	return "{\"fields\":[{\"name\":\"idModelo\",\"type\":\"long\"},{\"name\":\"fechaNotificacion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"fechaEvento\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"idSistema\",\"type\":\"int\"},{\"name\":\"sistema\",\"type\":\"string\"},{\"name\":\"contrato\",\"type\":\"string\"},{\"name\":\"envio\",\"type\":\"string\"},{\"name\":\"idEvento\",\"type\":\"int\"},{\"name\":\"evento\",\"type\":\"string\"},{\"default\":null,\"name\":\"idMotivo\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"motivo\",\"type\":[\"null\",\"string\"]},{\"name\":\"idSalida\",\"type\":\"int\"},{\"name\":\"salida\",\"type\":\"string\"},{\"name\":\"urlSalida\",\"type\":\"string\"},{\"default\":null,\"name\":\"idRegla\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"Regla\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"destinatarioNotificacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"destinatarioEstado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"destinatarioObservacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"destinatarioNotificacionCaracteres\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"remitente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remitenteNotificacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remitenteEstado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remitenteObservacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remitenteNotificacionCaracteres\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"proveedorSMS\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.Notificaciones.Events.Records.Notificaciones\",\"type\":\"record\"}"
}

func (r Notificaciones) SchemaName() string {
	return "Andreani.Notificaciones.Events.Records.Notificaciones"
}

func (_ Notificaciones) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Notificaciones) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Notificaciones) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Notificaciones) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Notificaciones) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Notificaciones) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Notificaciones) SetString(v string)   { panic("Unsupported operation") }
func (_ Notificaciones) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Notificaciones) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.IdModelo}

		return w

	case 1:
		w := types.Long{Target: &r.FechaNotificacion}

		return w

	case 2:
		w := types.Long{Target: &r.FechaEvento}

		return w

	case 3:
		w := types.Int{Target: &r.IdSistema}

		return w

	case 4:
		w := types.String{Target: &r.Sistema}

		return w

	case 5:
		w := types.String{Target: &r.Contrato}

		return w

	case 6:
		w := types.String{Target: &r.Envio}

		return w

	case 7:
		w := types.Int{Target: &r.IdEvento}

		return w

	case 8:
		w := types.String{Target: &r.Evento}

		return w

	case 9:
		r.IdMotivo = NewUnionNullInt()

		return r.IdMotivo
	case 10:
		r.Motivo = NewUnionNullString()

		return r.Motivo
	case 11:
		w := types.Int{Target: &r.IdSalida}

		return w

	case 12:
		w := types.String{Target: &r.Salida}

		return w

	case 13:
		w := types.String{Target: &r.UrlSalida}

		return w

	case 14:
		r.IdRegla = NewUnionNullInt()

		return r.IdRegla
	case 15:
		r.Regla = NewUnionNullString()

		return r.Regla
	case 16:
		r.Destinatario = NewUnionNullString()

		return r.Destinatario
	case 17:
		r.DestinatarioNotificacion = NewUnionNullString()

		return r.DestinatarioNotificacion
	case 18:
		r.DestinatarioEstado = NewUnionNullString()

		return r.DestinatarioEstado
	case 19:
		r.DestinatarioObservacion = NewUnionNullString()

		return r.DestinatarioObservacion
	case 20:
		r.DestinatarioNotificacionCaracteres = NewUnionNullInt()

		return r.DestinatarioNotificacionCaracteres
	case 21:
		r.Remitente = NewUnionNullString()

		return r.Remitente
	case 22:
		r.RemitenteNotificacion = NewUnionNullString()

		return r.RemitenteNotificacion
	case 23:
		r.RemitenteEstado = NewUnionNullString()

		return r.RemitenteEstado
	case 24:
		r.RemitenteObservacion = NewUnionNullString()

		return r.RemitenteObservacion
	case 25:
		r.RemitenteNotificacionCaracteres = NewUnionNullInt()

		return r.RemitenteNotificacionCaracteres
	case 26:
		r.ProveedorSMS = NewUnionNullString()

		return r.ProveedorSMS
	}
	panic("Unknown field index")
}

func (r *Notificaciones) SetDefault(i int) {
	switch i {
	case 9:
		r.IdMotivo = nil
		return
	case 10:
		r.Motivo = nil
		return
	case 14:
		r.IdRegla = nil
		return
	case 15:
		r.Regla = nil
		return
	case 16:
		r.Destinatario = nil
		return
	case 17:
		r.DestinatarioNotificacion = nil
		return
	case 18:
		r.DestinatarioEstado = nil
		return
	case 19:
		r.DestinatarioObservacion = nil
		return
	case 20:
		r.DestinatarioNotificacionCaracteres = nil
		return
	case 21:
		r.Remitente = nil
		return
	case 22:
		r.RemitenteNotificacion = nil
		return
	case 23:
		r.RemitenteEstado = nil
		return
	case 24:
		r.RemitenteObservacion = nil
		return
	case 25:
		r.RemitenteNotificacionCaracteres = nil
		return
	case 26:
		r.ProveedorSMS = nil
		return
	}
	panic("Unknown field index")
}

func (r *Notificaciones) NullField(i int) {
	switch i {
	case 9:
		r.IdMotivo = nil
		return
	case 10:
		r.Motivo = nil
		return
	case 14:
		r.IdRegla = nil
		return
	case 15:
		r.Regla = nil
		return
	case 16:
		r.Destinatario = nil
		return
	case 17:
		r.DestinatarioNotificacion = nil
		return
	case 18:
		r.DestinatarioEstado = nil
		return
	case 19:
		r.DestinatarioObservacion = nil
		return
	case 20:
		r.DestinatarioNotificacionCaracteres = nil
		return
	case 21:
		r.Remitente = nil
		return
	case 22:
		r.RemitenteNotificacion = nil
		return
	case 23:
		r.RemitenteEstado = nil
		return
	case 24:
		r.RemitenteObservacion = nil
		return
	case 25:
		r.RemitenteNotificacionCaracteres = nil
		return
	case 26:
		r.ProveedorSMS = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Notificaciones) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Notificaciones) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Notificaciones) HintSize(int)                     { panic("Unsupported operation") }
func (_ Notificaciones) Finalize()                        {}

func (_ Notificaciones) AvroCRC64Fingerprint() []byte {
	return []byte(NotificacionesAvroCRC64Fingerprint)
}

func (r Notificaciones) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["idModelo"], err = json.Marshal(r.IdModelo)
	if err != nil {
		return nil, err
	}
	output["fechaNotificacion"], err = json.Marshal(r.FechaNotificacion)
	if err != nil {
		return nil, err
	}
	output["fechaEvento"], err = json.Marshal(r.FechaEvento)
	if err != nil {
		return nil, err
	}
	output["idSistema"], err = json.Marshal(r.IdSistema)
	if err != nil {
		return nil, err
	}
	output["sistema"], err = json.Marshal(r.Sistema)
	if err != nil {
		return nil, err
	}
	output["contrato"], err = json.Marshal(r.Contrato)
	if err != nil {
		return nil, err
	}
	output["envio"], err = json.Marshal(r.Envio)
	if err != nil {
		return nil, err
	}
	output["idEvento"], err = json.Marshal(r.IdEvento)
	if err != nil {
		return nil, err
	}
	output["evento"], err = json.Marshal(r.Evento)
	if err != nil {
		return nil, err
	}
	output["idMotivo"], err = json.Marshal(r.IdMotivo)
	if err != nil {
		return nil, err
	}
	output["motivo"], err = json.Marshal(r.Motivo)
	if err != nil {
		return nil, err
	}
	output["idSalida"], err = json.Marshal(r.IdSalida)
	if err != nil {
		return nil, err
	}
	output["salida"], err = json.Marshal(r.Salida)
	if err != nil {
		return nil, err
	}
	output["urlSalida"], err = json.Marshal(r.UrlSalida)
	if err != nil {
		return nil, err
	}
	output["idRegla"], err = json.Marshal(r.IdRegla)
	if err != nil {
		return nil, err
	}
	output["Regla"], err = json.Marshal(r.Regla)
	if err != nil {
		return nil, err
	}
	output["destinatario"], err = json.Marshal(r.Destinatario)
	if err != nil {
		return nil, err
	}
	output["destinatarioNotificacion"], err = json.Marshal(r.DestinatarioNotificacion)
	if err != nil {
		return nil, err
	}
	output["destinatarioEstado"], err = json.Marshal(r.DestinatarioEstado)
	if err != nil {
		return nil, err
	}
	output["destinatarioObservacion"], err = json.Marshal(r.DestinatarioObservacion)
	if err != nil {
		return nil, err
	}
	output["destinatarioNotificacionCaracteres"], err = json.Marshal(r.DestinatarioNotificacionCaracteres)
	if err != nil {
		return nil, err
	}
	output["remitente"], err = json.Marshal(r.Remitente)
	if err != nil {
		return nil, err
	}
	output["remitenteNotificacion"], err = json.Marshal(r.RemitenteNotificacion)
	if err != nil {
		return nil, err
	}
	output["remitenteEstado"], err = json.Marshal(r.RemitenteEstado)
	if err != nil {
		return nil, err
	}
	output["remitenteObservacion"], err = json.Marshal(r.RemitenteObservacion)
	if err != nil {
		return nil, err
	}
	output["remitenteNotificacionCaracteres"], err = json.Marshal(r.RemitenteNotificacionCaracteres)
	if err != nil {
		return nil, err
	}
	output["proveedorSMS"], err = json.Marshal(r.ProveedorSMS)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Notificaciones) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["idModelo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdModelo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idModelo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaNotificacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaNotificacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fechaNotificacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaEvento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaEvento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fechaEvento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["idSistema"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdSistema); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idSistema")
	}
	val = func() json.RawMessage {
		if v, ok := fields["sistema"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sistema); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for sistema")
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
		return fmt.Errorf("no value specified for contrato")
	}
	val = func() json.RawMessage {
		if v, ok := fields["envio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Envio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for envio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["idEvento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdEvento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idEvento")
	}
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
		return fmt.Errorf("no value specified for evento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["idMotivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdMotivo); err != nil {
			return err
		}
	} else {
		r.IdMotivo = NewUnionNullInt()

		r.IdMotivo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["motivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Motivo); err != nil {
			return err
		}
	} else {
		r.Motivo = NewUnionNullString()

		r.Motivo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["idSalida"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdSalida); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idSalida")
	}
	val = func() json.RawMessage {
		if v, ok := fields["salida"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Salida); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for salida")
	}
	val = func() json.RawMessage {
		if v, ok := fields["urlSalida"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UrlSalida); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for urlSalida")
	}
	val = func() json.RawMessage {
		if v, ok := fields["idRegla"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdRegla); err != nil {
			return err
		}
	} else {
		r.IdRegla = NewUnionNullInt()

		r.IdRegla = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Regla"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Regla); err != nil {
			return err
		}
	} else {
		r.Regla = NewUnionNullString()

		r.Regla = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["destinatario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Destinatario); err != nil {
			return err
		}
	} else {
		r.Destinatario = NewUnionNullString()

		r.Destinatario = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["destinatarioNotificacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioNotificacion); err != nil {
			return err
		}
	} else {
		r.DestinatarioNotificacion = NewUnionNullString()

		r.DestinatarioNotificacion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["destinatarioEstado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioEstado); err != nil {
			return err
		}
	} else {
		r.DestinatarioEstado = NewUnionNullString()

		r.DestinatarioEstado = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["destinatarioObservacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioObservacion); err != nil {
			return err
		}
	} else {
		r.DestinatarioObservacion = NewUnionNullString()

		r.DestinatarioObservacion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["destinatarioNotificacionCaracteres"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioNotificacionCaracteres); err != nil {
			return err
		}
	} else {
		r.DestinatarioNotificacionCaracteres = NewUnionNullInt()

		r.DestinatarioNotificacionCaracteres = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["remitente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Remitente); err != nil {
			return err
		}
	} else {
		r.Remitente = NewUnionNullString()

		r.Remitente = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["remitenteNotificacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RemitenteNotificacion); err != nil {
			return err
		}
	} else {
		r.RemitenteNotificacion = NewUnionNullString()

		r.RemitenteNotificacion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["remitenteEstado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RemitenteEstado); err != nil {
			return err
		}
	} else {
		r.RemitenteEstado = NewUnionNullString()

		r.RemitenteEstado = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["remitenteObservacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RemitenteObservacion); err != nil {
			return err
		}
	} else {
		r.RemitenteObservacion = NewUnionNullString()

		r.RemitenteObservacion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["remitenteNotificacionCaracteres"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RemitenteNotificacionCaracteres); err != nil {
			return err
		}
	} else {
		r.RemitenteNotificacionCaracteres = NewUnionNullInt()

		r.RemitenteNotificacionCaracteres = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["proveedorSMS"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ProveedorSMS); err != nil {
			return err
		}
	} else {
		r.ProveedorSMS = NewUnionNullString()

		r.ProveedorSMS = nil
	}
	return nil
}
