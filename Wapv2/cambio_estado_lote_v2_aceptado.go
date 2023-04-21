// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CambioEstadoLoteV2Aceptado.avsc
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

type CambioEstadoLoteV2Aceptado struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	CambioEstadoDeLoteV2 CambioEstadoDeLoteV2 `json:"cambioEstadoDeLoteV2"`

	Topic string `json:"Topic"`

	Razon string `json:"razon"`
}

const CambioEstadoLoteV2AceptadoAvroCRC64Fingerprint = "\xd0\xf2\x89\x85ԉ\xc23"

func NewCambioEstadoLoteV2Aceptado() CambioEstadoLoteV2Aceptado {
	r := CambioEstadoLoteV2Aceptado{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.CambioEstadoDeLoteV2 = NewCambioEstadoDeLoteV2()

	r.Topic = "Almacen/Solicitudes/CambioEstadoLoteV2Aceptado"
	return r
}

func DeserializeCambioEstadoLoteV2Aceptado(r io.Reader) (CambioEstadoLoteV2Aceptado, error) {
	t := NewCambioEstadoLoteV2Aceptado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCambioEstadoLoteV2AceptadoFromSchema(r io.Reader, schema string) (CambioEstadoLoteV2Aceptado, error) {
	t := NewCambioEstadoLoteV2Aceptado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCambioEstadoLoteV2Aceptado(r CambioEstadoLoteV2Aceptado, w io.Writer) error {
	var err error
	err = writeSolicitudDeAccionAlmacen(r.SolicitudDeAccionAlmacen, w)
	if err != nil {
		return err
	}
	err = writeCambioEstadoDeLoteV2(r.CambioEstadoDeLoteV2, w)
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

func (r CambioEstadoLoteV2Aceptado) Serialize(w io.Writer) error {
	return writeCambioEstadoLoteV2Aceptado(r, w)
}

func (r CambioEstadoLoteV2Aceptado) Schema() string {
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"eventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"default\":null,\"name\":\"instancia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoDistribucion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"cambioEstadoDeLoteV2\",\"type\":{\"fields\":[{\"name\":\"articulo\",\"type\":\"string\"},{\"name\":\"estadoLote\",\"type\":\"string\"},{\"default\":null,\"name\":\"fechaVencimiento\",\"type\":[\"null\",\"string\"]},{\"name\":\"loteCaja\",\"type\":\"string\"},{\"name\":\"loteSecundario\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"}],\"name\":\"CambioEstadoDeLoteV2\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/CambioEstadoLoteV2Aceptado\",\"name\":\"Topic\",\"type\":\"string\"},{\"name\":\"razon\",\"type\":\"string\"}],\"name\":\"Andreani.Wapv2.Events.Record.CambioEstadoLoteV2Aceptado\",\"type\":\"record\"}"
}

func (r CambioEstadoLoteV2Aceptado) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.CambioEstadoLoteV2Aceptado"
}

func (_ CambioEstadoLoteV2Aceptado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CambioEstadoLoteV2Aceptado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CambioEstadoLoteV2Aceptado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CambioEstadoLoteV2Aceptado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CambioEstadoLoteV2Aceptado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CambioEstadoLoteV2Aceptado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CambioEstadoLoteV2Aceptado) SetString(v string)   { panic("Unsupported operation") }
func (_ CambioEstadoLoteV2Aceptado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CambioEstadoLoteV2Aceptado) Get(i int) types.Field {
	switch i {
	case 0:
		r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

		w := types.Record{Target: &r.SolicitudDeAccionAlmacen}

		return w

	case 1:
		r.CambioEstadoDeLoteV2 = NewCambioEstadoDeLoteV2()

		w := types.Record{Target: &r.CambioEstadoDeLoteV2}

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

func (r *CambioEstadoLoteV2Aceptado) SetDefault(i int) {
	switch i {
	case 2:
		r.Topic = "Almacen/Solicitudes/CambioEstadoLoteV2Aceptado"
		return
	}
	panic("Unknown field index")
}

func (r *CambioEstadoLoteV2Aceptado) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ CambioEstadoLoteV2Aceptado) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CambioEstadoLoteV2Aceptado) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CambioEstadoLoteV2Aceptado) HintSize(int)                     { panic("Unsupported operation") }
func (_ CambioEstadoLoteV2Aceptado) Finalize()                        {}

func (_ CambioEstadoLoteV2Aceptado) AvroCRC64Fingerprint() []byte {
	return []byte(CambioEstadoLoteV2AceptadoAvroCRC64Fingerprint)
}

func (r CambioEstadoLoteV2Aceptado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["solicitudDeAccionAlmacen"], err = json.Marshal(r.SolicitudDeAccionAlmacen)
	if err != nil {
		return nil, err
	}
	output["cambioEstadoDeLoteV2"], err = json.Marshal(r.CambioEstadoDeLoteV2)
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

func (r *CambioEstadoLoteV2Aceptado) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["cambioEstadoDeLoteV2"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CambioEstadoDeLoteV2); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cambioEstadoDeLoteV2")
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
		r.Topic = "Almacen/Solicitudes/CambioEstadoLoteV2Aceptado"
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
