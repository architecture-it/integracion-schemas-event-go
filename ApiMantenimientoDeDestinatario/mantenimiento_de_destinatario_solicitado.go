// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     MantenimientoDeDestinatarioSolicitado.avsc
 */
package ApiMantenimientoDeDestinatarioEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type MantenimientoDeDestinatarioSolicitado struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	DestinatarioAlta DestinatarioAlta `json:"destinatarioAlta"`

	Topic string `json:"Topic"`
}

const MantenimientoDeDestinatarioSolicitadoAvroCRC64Fingerprint = "\x0e\xbe\xf63B\x8d\x89+"

func NewMantenimientoDeDestinatarioSolicitado() MantenimientoDeDestinatarioSolicitado {
	r := MantenimientoDeDestinatarioSolicitado{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.DestinatarioAlta = NewDestinatarioAlta()

	r.Topic = "Almacen/Solicitudes/MantenimientoDeDestinatarioSolicitado"
	return r
}

func DeserializeMantenimientoDeDestinatarioSolicitado(r io.Reader) (MantenimientoDeDestinatarioSolicitado, error) {
	t := NewMantenimientoDeDestinatarioSolicitado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeMantenimientoDeDestinatarioSolicitadoFromSchema(r io.Reader, schema string) (MantenimientoDeDestinatarioSolicitado, error) {
	t := NewMantenimientoDeDestinatarioSolicitado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeMantenimientoDeDestinatarioSolicitado(r MantenimientoDeDestinatarioSolicitado, w io.Writer) error {
	var err error
	err = writeSolicitudDeAccionAlmacen(r.SolicitudDeAccionAlmacen, w)
	if err != nil {
		return err
	}
	err = writeDestinatarioAlta(r.DestinatarioAlta, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Topic, w)
	if err != nil {
		return err
	}
	return err
}

func (r MantenimientoDeDestinatarioSolicitado) Serialize(w io.Writer) error {
	return writeMantenimientoDeDestinatarioSolicitado(r, w)
}

func (r MantenimientoDeDestinatarioSolicitado) Schema() string {
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"eventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"name\":\"contrato\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"name\":\"almacenSAP\",\"type\":\"string\"},{\"name\":\"planta\",\"type\":\"string\"}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"destinatarioAlta\",\"type\":{\"fields\":[{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"tipo\",\"type\":\"string\"},{\"name\":\"direccion\",\"type\":\"string\"},{\"name\":\"ciudad\",\"type\":\"string\"},{\"name\":\"compania\",\"type\":\"string\"},{\"name\":\"pais\",\"type\":\"string\"},{\"name\":\"descripcion\",\"type\":\"string\"},{\"name\":\"lugar\",\"type\":\"string\"},{\"name\":\"telefono\",\"type\":\"string\"},{\"name\":\"provincia\",\"type\":\"string\"},{\"name\":\"identificadorFiscal\",\"type\":\"string\"},{\"name\":\"codigoPostal\",\"type\":\"string\"},{\"name\":\"cuit\",\"type\":\"string\"},{\"name\":\"ramo\",\"type\":\"string\"},{\"name\":\"gln\",\"type\":\"string\"},{\"name\":\"agente\",\"type\":\"string\"}],\"name\":\"DestinatarioAlta\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/MantenimientoDeDestinatarioSolicitado\",\"name\":\"Topic\",\"type\":\"string\"}],\"name\":\"Andreani.ApiMantenimientoDeDestinatario.Events.Record.MantenimientoDeDestinatarioSolicitado\",\"type\":\"record\"}"
}

func (r MantenimientoDeDestinatarioSolicitado) SchemaName() string {
	return "Andreani.ApiMantenimientoDeDestinatario.Events.Record.MantenimientoDeDestinatarioSolicitado"
}

func (_ MantenimientoDeDestinatarioSolicitado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ MantenimientoDeDestinatarioSolicitado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ MantenimientoDeDestinatarioSolicitado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ MantenimientoDeDestinatarioSolicitado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ MantenimientoDeDestinatarioSolicitado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ MantenimientoDeDestinatarioSolicitado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ MantenimientoDeDestinatarioSolicitado) SetString(v string)   { panic("Unsupported operation") }
func (_ MantenimientoDeDestinatarioSolicitado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MantenimientoDeDestinatarioSolicitado) Get(i int) types.Field {
	switch i {
	case 0:
		r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

		w := types.Record{Target: &r.SolicitudDeAccionAlmacen}

		return w

	case 1:
		r.DestinatarioAlta = NewDestinatarioAlta()

		w := types.Record{Target: &r.DestinatarioAlta}

		return w

	case 2:
		w := types.String{Target: &r.Topic}

		return w

	}
	panic("Unknown field index")
}

func (r *MantenimientoDeDestinatarioSolicitado) SetDefault(i int) {
	switch i {
	case 2:
		r.Topic = "Almacen/Solicitudes/MantenimientoDeDestinatarioSolicitado"
		return
	}
	panic("Unknown field index")
}

func (r *MantenimientoDeDestinatarioSolicitado) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ MantenimientoDeDestinatarioSolicitado) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ MantenimientoDeDestinatarioSolicitado) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ MantenimientoDeDestinatarioSolicitado) HintSize(int) { panic("Unsupported operation") }
func (_ MantenimientoDeDestinatarioSolicitado) Finalize()    {}

func (_ MantenimientoDeDestinatarioSolicitado) AvroCRC64Fingerprint() []byte {
	return []byte(MantenimientoDeDestinatarioSolicitadoAvroCRC64Fingerprint)
}

func (r MantenimientoDeDestinatarioSolicitado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["solicitudDeAccionAlmacen"], err = json.Marshal(r.SolicitudDeAccionAlmacen)
	if err != nil {
		return nil, err
	}
	output["destinatarioAlta"], err = json.Marshal(r.DestinatarioAlta)
	if err != nil {
		return nil, err
	}
	output["Topic"], err = json.Marshal(r.Topic)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *MantenimientoDeDestinatarioSolicitado) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["destinatarioAlta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioAlta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for destinatarioAlta")
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
		r.Topic = "Almacen/Solicitudes/MantenimientoDeDestinatarioSolicitado"
	}
	return nil
}
