// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CambioDeAlmacenConRechazadoV2.avsc
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

type CambioDeAlmacenConRechazadoV2 struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	CambioAlmacenConLpn CambioAlmacenConLpn `json:"cambioAlmacenConLpn"`

	Topic string `json:"Topic"`

	Razon string `json:"razon"`
}

const CambioDeAlmacenConRechazadoV2AvroCRC64Fingerprint = "\xb4\xf4<\x0f\xa1K\x0fo"

func NewCambioDeAlmacenConRechazadoV2() CambioDeAlmacenConRechazadoV2 {
	r := CambioDeAlmacenConRechazadoV2{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.CambioAlmacenConLpn = NewCambioAlmacenConLpn()

	r.Topic = "Almacen/Solicitudes/CambioDeAlmacenConLpnRechazadoV2"
	return r
}

func DeserializeCambioDeAlmacenConRechazadoV2(r io.Reader) (CambioDeAlmacenConRechazadoV2, error) {
	t := NewCambioDeAlmacenConRechazadoV2()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCambioDeAlmacenConRechazadoV2FromSchema(r io.Reader, schema string) (CambioDeAlmacenConRechazadoV2, error) {
	t := NewCambioDeAlmacenConRechazadoV2()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCambioDeAlmacenConRechazadoV2(r CambioDeAlmacenConRechazadoV2, w io.Writer) error {
	var err error
	err = writeSolicitudDeAccionAlmacen(r.SolicitudDeAccionAlmacen, w)
	if err != nil {
		return err
	}
	err = writeCambioAlmacenConLpn(r.CambioAlmacenConLpn, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Topic, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Razon, w)
	if err != nil {
		return err
	}
	return err
}

func (r CambioDeAlmacenConRechazadoV2) Serialize(w io.Writer) error {
	return writeCambioDeAlmacenConRechazadoV2(r, w)
}

func (r CambioDeAlmacenConRechazadoV2) Schema() string {
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"eventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"default\":null,\"name\":\"instancia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoDistribucion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"cambioAlmacenConLpn\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"propietario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"articulo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"lote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cantidad\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"almacenOrigen\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"almacenDestino\",\"type\":[\"null\",\"string\"]}],\"name\":\"CambioAlmacenConLpn\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/CambioDeAlmacenConLpnRechazadoV2\",\"name\":\"Topic\",\"type\":\"string\"},{\"name\":\"razon\",\"type\":\"string\"}],\"name\":\"Andreani.Wapv2.Events.Record.CambioDeAlmacenConRechazadoV2\",\"type\":\"record\"}"
}

func (r CambioDeAlmacenConRechazadoV2) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.CambioDeAlmacenConRechazadoV2"
}

func (_ CambioDeAlmacenConRechazadoV2) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CambioDeAlmacenConRechazadoV2) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CambioDeAlmacenConRechazadoV2) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CambioDeAlmacenConRechazadoV2) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CambioDeAlmacenConRechazadoV2) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CambioDeAlmacenConRechazadoV2) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CambioDeAlmacenConRechazadoV2) SetString(v string)   { panic("Unsupported operation") }
func (_ CambioDeAlmacenConRechazadoV2) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CambioDeAlmacenConRechazadoV2) Get(i int) types.Field {
	switch i {
	case 0:
		r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

		w := types.Record{Target: &r.SolicitudDeAccionAlmacen}

		return w

	case 1:
		r.CambioAlmacenConLpn = NewCambioAlmacenConLpn()

		w := types.Record{Target: &r.CambioAlmacenConLpn}

		return w

	case 2:
		w := types.String{Target: &r.Topic}

		return w

	case 3:
		w := types.String{Target: &r.Razon}

		return w

	}
	panic("Unknown field index")
}

func (r *CambioDeAlmacenConRechazadoV2) SetDefault(i int) {
	switch i {
	case 2:
		r.Topic = "Almacen/Solicitudes/CambioDeAlmacenConLpnRechazadoV2"
		return
	}
	panic("Unknown field index")
}

func (r *CambioDeAlmacenConRechazadoV2) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ CambioDeAlmacenConRechazadoV2) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ CambioDeAlmacenConRechazadoV2) AppendArray() types.Field { panic("Unsupported operation") }
func (_ CambioDeAlmacenConRechazadoV2) HintSize(int)             { panic("Unsupported operation") }
func (_ CambioDeAlmacenConRechazadoV2) Finalize()                {}

func (_ CambioDeAlmacenConRechazadoV2) AvroCRC64Fingerprint() []byte {
	return []byte(CambioDeAlmacenConRechazadoV2AvroCRC64Fingerprint)
}

func (r CambioDeAlmacenConRechazadoV2) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["solicitudDeAccionAlmacen"], err = json.Marshal(r.SolicitudDeAccionAlmacen)
	if err != nil {
		return nil, err
	}
	output["cambioAlmacenConLpn"], err = json.Marshal(r.CambioAlmacenConLpn)
	if err != nil {
		return nil, err
	}
	output["Topic"], err = json.Marshal(r.Topic)
	if err != nil {
		return nil, err
	}
	output["razon"], err = json.Marshal(r.Razon)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CambioDeAlmacenConRechazadoV2) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["cambioAlmacenConLpn"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CambioAlmacenConLpn); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cambioAlmacenConLpn")
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
		r.Topic = "Almacen/Solicitudes/CambioDeAlmacenConLpnRechazadoV2"
	}
	val = func() json.RawMessage {
		if v, ok := fields["razon"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Razon); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for razon")
	}
	return nil
}