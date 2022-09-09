// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CreacionDeOrdenDeCompraRechazada.avsc
 */
package Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EventoDeNegocio struct {
	Timestamp int64 `json:"timestamp"`

	Remitente string `json:"remitente"`

	Destinatario *UnionNullString `json:"destinatario"`

	NumeroDeOrden *UnionNullString `json:"numeroDeOrden"`

	Vencimiento *UnionNullLong `json:"vencimiento"`
}

const EventoDeNegocioAvroCRC64Fingerprint = "\xd7\b3\xafI\vW\xe2"

func NewEventoDeNegocio() EventoDeNegocio {
	r := EventoDeNegocio{}
	r.Destinatario = nil
	r.NumeroDeOrden = nil
	r.Vencimiento = nil
	return r
}

func DeserializeEventoDeNegocio(r io.Reader) (EventoDeNegocio, error) {
	t := NewEventoDeNegocio()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoDeNegocioFromSchema(r io.Reader, schema string) (EventoDeNegocio, error) {
	t := NewEventoDeNegocio()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoDeNegocio(r EventoDeNegocio, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.Timestamp, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Remitente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Destinatario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeOrden, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.Vencimiento, w)
	if err != nil {
		return err
	}
	return err
}

func (r EventoDeNegocio) Serialize(w io.Writer) error {
	return writeEventoDeNegocio(r, w)
}

func (r EventoDeNegocio) Schema() string {
	return "{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"Andreani.Wap.EventosDeNegocio.EventoDeNegocio\",\"type\":\"record\"}"
}

func (r EventoDeNegocio) SchemaName() string {
	return "Andreani.Wap.EventosDeNegocio.EventoDeNegocio"
}

func (_ EventoDeNegocio) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoDeNegocio) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoDeNegocio) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoDeNegocio) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoDeNegocio) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoDeNegocio) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoDeNegocio) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoDeNegocio) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoDeNegocio) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.Timestamp}

		return w

	case 1:
		w := types.String{Target: &r.Remitente}

		return w

	case 2:
		r.Destinatario = NewUnionNullString()

		return r.Destinatario
	case 3:
		r.NumeroDeOrden = NewUnionNullString()

		return r.NumeroDeOrden
	case 4:
		r.Vencimiento = NewUnionNullLong()

		return r.Vencimiento
	}
	panic("Unknown field index")
}

func (r *EventoDeNegocio) SetDefault(i int) {
	switch i {
	case 2:
		r.Destinatario = nil
		return
	case 3:
		r.NumeroDeOrden = nil
		return
	case 4:
		r.Vencimiento = nil
		return
	}
	panic("Unknown field index")
}

func (r *EventoDeNegocio) NullField(i int) {
	switch i {
	case 2:
		r.Destinatario = nil
		return
	case 3:
		r.NumeroDeOrden = nil
		return
	case 4:
		r.Vencimiento = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ EventoDeNegocio) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EventoDeNegocio) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EventoDeNegocio) HintSize(int)                     { panic("Unsupported operation") }
func (_ EventoDeNegocio) Finalize()                        {}

func (_ EventoDeNegocio) AvroCRC64Fingerprint() []byte {
	return []byte(EventoDeNegocioAvroCRC64Fingerprint)
}

func (r EventoDeNegocio) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["timestamp"], err = json.Marshal(r.Timestamp)
	if err != nil {
		return nil, err
	}
	output["remitente"], err = json.Marshal(r.Remitente)
	if err != nil {
		return nil, err
	}
	output["destinatario"], err = json.Marshal(r.Destinatario)
	if err != nil {
		return nil, err
	}
	output["numeroDeOrden"], err = json.Marshal(r.NumeroDeOrden)
	if err != nil {
		return nil, err
	}
	output["vencimiento"], err = json.Marshal(r.Vencimiento)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EventoDeNegocio) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["timestamp"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Timestamp); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for timestamp")
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
		return fmt.Errorf("no value specified for remitente")
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
		if v, ok := fields["numeroDeOrden"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeOrden); err != nil {
			return err
		}
	} else {
		r.NumeroDeOrden = NewUnionNullString()

		r.NumeroDeOrden = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["vencimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Vencimiento); err != nil {
			return err
		}
	} else {
		r.Vencimiento = NewUnionNullLong()

		r.Vencimiento = nil
	}
	return nil
}
