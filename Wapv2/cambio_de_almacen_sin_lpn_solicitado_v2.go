// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CambioDeAlmacenSinLpnSolicitadoV2.avsc
 */
package Wapv2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type CambioDeAlmacenSinLpnSolicitadoV2 struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	CambioAlmacenSinLpn CambioAlmacenSinLpn `json:"cambioAlmacenSinLpn"`

	Topic string `json:"Topic"`
}

const CambioDeAlmacenSinLpnSolicitadoV2AvroCRC64Fingerprint = "\xc0^\x00\xc4?Y\xf1\xc9"

func NewCambioDeAlmacenSinLpnSolicitadoV2() CambioDeAlmacenSinLpnSolicitadoV2 {
	r := CambioDeAlmacenSinLpnSolicitadoV2{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.CambioAlmacenSinLpn = NewCambioAlmacenSinLpn()

	r.Topic = "Almacen/Solicitudes/CambioDeAlmacenSinLpnSolicitadoV2"
	return r
}

func DeserializeCambioDeAlmacenSinLpnSolicitadoV2(r io.Reader) (CambioDeAlmacenSinLpnSolicitadoV2, error) {
	t := NewCambioDeAlmacenSinLpnSolicitadoV2()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCambioDeAlmacenSinLpnSolicitadoV2FromSchema(r io.Reader, schema string) (CambioDeAlmacenSinLpnSolicitadoV2, error) {
	t := NewCambioDeAlmacenSinLpnSolicitadoV2()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCambioDeAlmacenSinLpnSolicitadoV2(r CambioDeAlmacenSinLpnSolicitadoV2, w io.Writer) error {
	var err error
	err = writeSolicitudDeAccionAlmacen(r.SolicitudDeAccionAlmacen, w)
	if err != nil {
		return err
	}
	err = writeCambioAlmacenSinLpn(r.CambioAlmacenSinLpn, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Topic, w)
	if err != nil {
		return err
	}
	return err
}

func (r CambioDeAlmacenSinLpnSolicitadoV2) Serialize(w io.Writer) error {
	return writeCambioDeAlmacenSinLpnSolicitadoV2(r, w)
}

func (r CambioDeAlmacenSinLpnSolicitadoV2) Schema() string {
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"eventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"default\":null,\"name\":\"instancia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoDistribucion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"cambioAlmacenSinLpn\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"propietario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"articulo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"lote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cantidad\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"almacenOrigen\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"almacenDestino\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"serie\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteEstado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ordenExterna\",\"type\":[\"null\",\"string\"]}],\"name\":\"CambioAlmacenSinLpn\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/CambioDeAlmacenSinLpnSolicitadoV2\",\"name\":\"Topic\",\"type\":\"string\"}],\"name\":\"Andreani.Wapv2.Events.Record.CambioDeAlmacenSinLpnSolicitadoV2\",\"type\":\"record\"}"
}

func (r CambioDeAlmacenSinLpnSolicitadoV2) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.CambioDeAlmacenSinLpnSolicitadoV2"
}

func (_ CambioDeAlmacenSinLpnSolicitadoV2) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CambioDeAlmacenSinLpnSolicitadoV2) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CambioDeAlmacenSinLpnSolicitadoV2) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CambioDeAlmacenSinLpnSolicitadoV2) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CambioDeAlmacenSinLpnSolicitadoV2) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CambioDeAlmacenSinLpnSolicitadoV2) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CambioDeAlmacenSinLpnSolicitadoV2) SetString(v string)   { panic("Unsupported operation") }
func (_ CambioDeAlmacenSinLpnSolicitadoV2) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CambioDeAlmacenSinLpnSolicitadoV2) Get(i int) types.Field {
	switch i {
	case 0:
		r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

		w := types.Record{Target: &r.SolicitudDeAccionAlmacen}

		return w

	case 1:
		r.CambioAlmacenSinLpn = NewCambioAlmacenSinLpn()

		w := types.Record{Target: &r.CambioAlmacenSinLpn}

		return w

	case 2:
		w := types.String{Target: &r.Topic}

		return w

	}
	panic("Unknown field index")
}

func (r *CambioDeAlmacenSinLpnSolicitadoV2) SetDefault(i int) {
	switch i {
	case 2:
		r.Topic = "Almacen/Solicitudes/CambioDeAlmacenSinLpnSolicitadoV2"
		return
	}
	panic("Unknown field index")
}

func (r *CambioDeAlmacenSinLpnSolicitadoV2) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ CambioDeAlmacenSinLpnSolicitadoV2) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ CambioDeAlmacenSinLpnSolicitadoV2) AppendArray() types.Field { panic("Unsupported operation") }
func (_ CambioDeAlmacenSinLpnSolicitadoV2) HintSize(int)             { panic("Unsupported operation") }
func (_ CambioDeAlmacenSinLpnSolicitadoV2) Finalize()                {}

func (_ CambioDeAlmacenSinLpnSolicitadoV2) AvroCRC64Fingerprint() []byte {
	return []byte(CambioDeAlmacenSinLpnSolicitadoV2AvroCRC64Fingerprint)
}

func (r CambioDeAlmacenSinLpnSolicitadoV2) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["solicitudDeAccionAlmacen"], err = json.Marshal(r.SolicitudDeAccionAlmacen)
	if err != nil {
		return nil, err
	}
	output["cambioAlmacenSinLpn"], err = json.Marshal(r.CambioAlmacenSinLpn)
	if err != nil {
		return nil, err
	}
	output["Topic"], err = json.Marshal(r.Topic)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CambioDeAlmacenSinLpnSolicitadoV2) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["solicitudDeAccionAlmacen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SolicitudDeAccionAlmacen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for solicitudDeAccionAlmacen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["cambioAlmacenSinLpn"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CambioAlmacenSinLpn); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cambioAlmacenSinLpn")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Topic"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Topic); err != nil {
			return err
		}
	} else {
		r.Topic = "Almacen/Solicitudes/CambioDeAlmacenSinLpnSolicitadoV2"
	}
	return nil
}
