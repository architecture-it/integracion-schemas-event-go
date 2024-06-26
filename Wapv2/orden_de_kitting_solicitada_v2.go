// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     OrdenDeKittingSolicitadaV2.avsc
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

type OrdenDeKittingSolicitadaV2 struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	OrdenDeKitting OrdenDeKitting `json:"ordenDeKitting"`

	Topic string `json:"Topic"`
}

const OrdenDeKittingSolicitadaV2AvroCRC64Fingerprint = "\xe8f;\xb9\xfb\x85\x98\t"

func NewOrdenDeKittingSolicitadaV2() OrdenDeKittingSolicitadaV2 {
	r := OrdenDeKittingSolicitadaV2{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.OrdenDeKitting = NewOrdenDeKitting()

	r.Topic = "Almacen/Solicitudes/OrdenDeKittingSolicitadaV2"
	return r
}

func DeserializeOrdenDeKittingSolicitadaV2(r io.Reader) (OrdenDeKittingSolicitadaV2, error) {
	t := NewOrdenDeKittingSolicitadaV2()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeOrdenDeKittingSolicitadaV2FromSchema(r io.Reader, schema string) (OrdenDeKittingSolicitadaV2, error) {
	t := NewOrdenDeKittingSolicitadaV2()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeOrdenDeKittingSolicitadaV2(r OrdenDeKittingSolicitadaV2, w io.Writer) error {
	var err error
	err = writeSolicitudDeAccionAlmacen(r.SolicitudDeAccionAlmacen, w)
	if err != nil {
		return err
	}
	err = writeOrdenDeKitting(r.OrdenDeKitting, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Topic, w)
	if err != nil {
		return err
	}
	return err
}

func (r OrdenDeKittingSolicitadaV2) Serialize(w io.Writer) error {
	return writeOrdenDeKittingSolicitadaV2(r, w)
}

func (r OrdenDeKittingSolicitadaV2) Schema() string {
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"eventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"default\":null,\"name\":\"instancia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoDistribucion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"ordenDeKitting\",\"type\":{\"fields\":[{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"numeroOrdenExterna\",\"type\":\"string\"},{\"default\":null,\"name\":\"articulo\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"cantidad\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"loteFabricante\",\"type\":\"string\"},{\"name\":\"loteSecundario\",\"type\":\"string\"},{\"name\":\"loteAlmacen\",\"type\":\"string\"},{\"name\":\"loteEstado\",\"type\":\"string\"},{\"default\":null,\"name\":\"acondi\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]}],\"name\":\"ArticuloKitting\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"articuloOrigen\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"codigoArticulo\",\"type\":\"string\"},{\"name\":\"loteAlmacen\",\"type\":\"string\"},{\"name\":\"loteSecundario\",\"type\":\"string\"},{\"name\":\"loteProveedor\",\"type\":\"string\"},{\"name\":\"loteEstado\",\"type\":\"string\"}],\"name\":\"ArticuloKittingOrigen\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"OrdenDeKitting\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/OrdenDeKittingSolicitadaV2\",\"name\":\"Topic\",\"type\":\"string\"}],\"name\":\"Andreani.Wapv2.Events.Record.OrdenDeKittingSolicitadaV2\",\"type\":\"record\"}"
}

func (r OrdenDeKittingSolicitadaV2) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.OrdenDeKittingSolicitadaV2"
}

func (_ OrdenDeKittingSolicitadaV2) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ OrdenDeKittingSolicitadaV2) SetInt(v int32)       { panic("Unsupported operation") }
func (_ OrdenDeKittingSolicitadaV2) SetLong(v int64)      { panic("Unsupported operation") }
func (_ OrdenDeKittingSolicitadaV2) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ OrdenDeKittingSolicitadaV2) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ OrdenDeKittingSolicitadaV2) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ OrdenDeKittingSolicitadaV2) SetString(v string)   { panic("Unsupported operation") }
func (_ OrdenDeKittingSolicitadaV2) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *OrdenDeKittingSolicitadaV2) Get(i int) types.Field {
	switch i {
	case 0:
		r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

		w := types.Record{Target: &r.SolicitudDeAccionAlmacen}

		return w

	case 1:
		r.OrdenDeKitting = NewOrdenDeKitting()

		w := types.Record{Target: &r.OrdenDeKitting}

		return w

	case 2:
		w := types.String{Target: &r.Topic}

		return w

	}
	panic("Unknown field index")
}

func (r *OrdenDeKittingSolicitadaV2) SetDefault(i int) {
	switch i {
	case 2:
		r.Topic = "Almacen/Solicitudes/OrdenDeKittingSolicitadaV2"
		return
	}
	panic("Unknown field index")
}

func (r *OrdenDeKittingSolicitadaV2) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ OrdenDeKittingSolicitadaV2) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ OrdenDeKittingSolicitadaV2) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ OrdenDeKittingSolicitadaV2) HintSize(int)                     { panic("Unsupported operation") }
func (_ OrdenDeKittingSolicitadaV2) Finalize()                        {}

func (_ OrdenDeKittingSolicitadaV2) AvroCRC64Fingerprint() []byte {
	return []byte(OrdenDeKittingSolicitadaV2AvroCRC64Fingerprint)
}

func (r OrdenDeKittingSolicitadaV2) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["solicitudDeAccionAlmacen"], err = json.Marshal(r.SolicitudDeAccionAlmacen)
	if err != nil {
		return nil, err
	}
	output["ordenDeKitting"], err = json.Marshal(r.OrdenDeKitting)
	if err != nil {
		return nil, err
	}
	output["Topic"], err = json.Marshal(r.Topic)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *OrdenDeKittingSolicitadaV2) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["ordenDeKitting"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenDeKitting); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ordenDeKitting")
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
		r.Topic = "Almacen/Solicitudes/OrdenDeKittingSolicitadaV2"
	}
	return nil
}
