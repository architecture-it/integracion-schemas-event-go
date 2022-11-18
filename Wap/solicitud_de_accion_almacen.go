// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     SolicitudDeAccionAlmacen.avsc
 */
package WapEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type SolicitudDeAccionAlmacen struct {
	EventoDeNegocio EventoDeNegocio `json:"eventoDeNegocio"`

	IdTransaccion string `json:"idTransaccion"`

	Almacen string `json:"almacen"`

	Instancia *UnionNullString `json:"instancia"`

	ContratoDistribucion *UnionNullString `json:"contratoDistribucion"`

	ContratoWarehouse *UnionNullString `json:"contratoWarehouse"`
}

const SolicitudDeAccionAlmacenAvroCRC64Fingerprint = "Q\xa4F:\xf1Pn\xb8"

func NewSolicitudDeAccionAlmacen() SolicitudDeAccionAlmacen {
	r := SolicitudDeAccionAlmacen{}
	r.EventoDeNegocio = NewEventoDeNegocio()

	r.Instancia = nil
	r.ContratoDistribucion = nil
	r.ContratoWarehouse = nil
	return r
}

func DeserializeSolicitudDeAccionAlmacen(r io.Reader) (SolicitudDeAccionAlmacen, error) {
	t := NewSolicitudDeAccionAlmacen()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSolicitudDeAccionAlmacenFromSchema(r io.Reader, schema string) (SolicitudDeAccionAlmacen, error) {
	t := NewSolicitudDeAccionAlmacen()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSolicitudDeAccionAlmacen(r SolicitudDeAccionAlmacen, w io.Writer) error {
	var err error
	err = writeEventoDeNegocio(r.EventoDeNegocio, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.IdTransaccion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Almacen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Instancia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ContratoDistribucion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ContratoWarehouse, w)
	if err != nil {
		return err
	}
	return err
}

func (r SolicitudDeAccionAlmacen) Serialize(w io.Writer) error {
	return writeSolicitudDeAccionAlmacen(r, w)
}

func (r SolicitudDeAccionAlmacen) Schema() string {
	return "{\"fields\":[{\"name\":\"eventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"default\":null,\"name\":\"instancia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoDistribucion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.Wap.Events.Record.SolicitudDeAccionAlmacen\",\"type\":\"record\"}"
}

func (r SolicitudDeAccionAlmacen) SchemaName() string {
	return "Andreani.Wap.Events.Record.SolicitudDeAccionAlmacen"
}

func (_ SolicitudDeAccionAlmacen) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ SolicitudDeAccionAlmacen) SetInt(v int32)       { panic("Unsupported operation") }
func (_ SolicitudDeAccionAlmacen) SetLong(v int64)      { panic("Unsupported operation") }
func (_ SolicitudDeAccionAlmacen) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ SolicitudDeAccionAlmacen) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ SolicitudDeAccionAlmacen) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ SolicitudDeAccionAlmacen) SetString(v string)   { panic("Unsupported operation") }
func (_ SolicitudDeAccionAlmacen) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *SolicitudDeAccionAlmacen) Get(i int) types.Field {
	switch i {
	case 0:
		r.EventoDeNegocio = NewEventoDeNegocio()

		w := types.Record{Target: &r.EventoDeNegocio}

		return w

	case 1:
		w := types.String{Target: &r.IdTransaccion}

		return w

	case 2:
		w := types.String{Target: &r.Almacen}

		return w

	case 3:
		r.Instancia = NewUnionNullString()

		return r.Instancia
	case 4:
		r.ContratoDistribucion = NewUnionNullString()

		return r.ContratoDistribucion
	case 5:
		r.ContratoWarehouse = NewUnionNullString()

		return r.ContratoWarehouse
	}
	panic("Unknown field index")
}

func (r *SolicitudDeAccionAlmacen) SetDefault(i int) {
	switch i {
	case 3:
		r.Instancia = nil
		return
	case 4:
		r.ContratoDistribucion = nil
		return
	case 5:
		r.ContratoWarehouse = nil
		return
	}
	panic("Unknown field index")
}

func (r *SolicitudDeAccionAlmacen) NullField(i int) {
	switch i {
	case 3:
		r.Instancia = nil
		return
	case 4:
		r.ContratoDistribucion = nil
		return
	case 5:
		r.ContratoWarehouse = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ SolicitudDeAccionAlmacen) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ SolicitudDeAccionAlmacen) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ SolicitudDeAccionAlmacen) HintSize(int)                     { panic("Unsupported operation") }
func (_ SolicitudDeAccionAlmacen) Finalize()                        {}

func (_ SolicitudDeAccionAlmacen) AvroCRC64Fingerprint() []byte {
	return []byte(SolicitudDeAccionAlmacenAvroCRC64Fingerprint)
}

func (r SolicitudDeAccionAlmacen) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["eventoDeNegocio"], err = json.Marshal(r.EventoDeNegocio)
	if err != nil {
		return nil, err
	}
	output["idTransaccion"], err = json.Marshal(r.IdTransaccion)
	if err != nil {
		return nil, err
	}
	output["almacen"], err = json.Marshal(r.Almacen)
	if err != nil {
		return nil, err
	}
	output["instancia"], err = json.Marshal(r.Instancia)
	if err != nil {
		return nil, err
	}
	output["contratoDistribucion"], err = json.Marshal(r.ContratoDistribucion)
	if err != nil {
		return nil, err
	}
	output["contratoWarehouse"], err = json.Marshal(r.ContratoWarehouse)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *SolicitudDeAccionAlmacen) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["eventoDeNegocio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EventoDeNegocio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for eventoDeNegocio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["idTransaccion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdTransaccion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idTransaccion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["almacen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Almacen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for almacen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["instancia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Instancia); err != nil {
			return err
		}
	} else {
		r.Instancia = NewUnionNullString()

		r.Instancia = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["contratoDistribucion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ContratoDistribucion); err != nil {
			return err
		}
	} else {
		r.ContratoDistribucion = NewUnionNullString()

		r.ContratoDistribucion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["contratoWarehouse"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ContratoWarehouse); err != nil {
			return err
		}
	} else {
		r.ContratoWarehouse = NewUnionNullString()

		r.ContratoWarehouse = nil
	}
	return nil
}
