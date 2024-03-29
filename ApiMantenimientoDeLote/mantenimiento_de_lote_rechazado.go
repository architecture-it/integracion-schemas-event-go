// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     MantenimientoDeLoteRechazado.avsc
 */
package ApiMantenimientoDeLoteEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type MantenimientoDeLoteRechazado struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	MantenimientoDeLoteRequest MantenimientoDeLoteRequest `json:"mantenimientoDeLoteRequest"`

	Topic string `json:"Topic"`

	Razon string `json:"razon"`
}

const MantenimientoDeLoteRechazadoAvroCRC64Fingerprint = "~q\a\n\xc61K\xbc"

func NewMantenimientoDeLoteRechazado() MantenimientoDeLoteRechazado {
	r := MantenimientoDeLoteRechazado{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.MantenimientoDeLoteRequest = NewMantenimientoDeLoteRequest()

	r.Topic = "Almacen/Solicitudes/MantenimientoDeLoteRechazado"
	return r
}

func DeserializeMantenimientoDeLoteRechazado(r io.Reader) (MantenimientoDeLoteRechazado, error) {
	t := NewMantenimientoDeLoteRechazado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeMantenimientoDeLoteRechazadoFromSchema(r io.Reader, schema string) (MantenimientoDeLoteRechazado, error) {
	t := NewMantenimientoDeLoteRechazado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeMantenimientoDeLoteRechazado(r MantenimientoDeLoteRechazado, w io.Writer) error {
	var err error
	err = writeSolicitudDeAccionAlmacen(r.SolicitudDeAccionAlmacen, w)
	if err != nil {
		return err
	}
	err = writeMantenimientoDeLoteRequest(r.MantenimientoDeLoteRequest, w)
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

func (r MantenimientoDeLoteRechazado) Serialize(w io.Writer) error {
	return writeMantenimientoDeLoteRechazado(r, w)
}

func (r MantenimientoDeLoteRechazado) Schema() string {
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"eventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"name\":\"numeroDeOrden\",\"type\":\"string\"},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"name\":\"contrato\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"name\":\"planta\",\"type\":\"string\"}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"mantenimientoDeLoteRequest\",\"type\":{\"fields\":[{\"name\":\"contrato\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":[\"null\",\"string\"]},{\"name\":\"almacenSap\",\"type\":[\"null\",\"string\"]},{\"name\":\"planta\",\"type\":\"string\"},{\"name\":\"uriConsulta\",\"type\":[\"null\",\"string\"]},{\"name\":\"mantenimientoDeLote\",\"type\":{\"fields\":[{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"articulo\",\"type\":\"string\"},{\"name\":\"paquete\",\"type\":[\"null\",\"string\"]},{\"name\":\"loteCaja\",\"type\":[\"null\",\"string\"]},{\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"name\":\"loteSap\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaFabricacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaVencimiento\",\"type\":[\"null\",\"string\"]},{\"name\":\"trazable\",\"type\":[\"null\",\"string\"]},{\"name\":\"estado\",\"type\":[\"null\",\"string\"]},{\"name\":\"procedencia\",\"type\":[\"null\",\"string\"]},{\"name\":\"campoLibre1\",\"type\":[\"null\",\"string\"]},{\"name\":\"campoLibre2\",\"type\":[\"null\",\"string\"]},{\"name\":\"campoLibre3\",\"type\":[\"null\",\"string\"]},{\"name\":\"campoLibre4\",\"type\":[\"null\",\"string\"]},{\"name\":\"campoLibre5\",\"type\":[\"null\",\"string\"]},{\"name\":\"loteExternoCliente\",\"type\":[\"null\",\"string\"]},{\"name\":\"deliverByDate\",\"type\":[\"null\",\"string\"]},{\"name\":\"bestByDate\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaCreacion\",\"type\":\"string\"},{\"name\":\"usuarioCreacion\",\"type\":\"string\"},{\"name\":\"fechaEdicion\",\"type\":\"string\"},{\"name\":\"usuarioEdicion\",\"type\":\"string\"}],\"name\":\"MantenimientoDeLote\",\"type\":\"record\"}}],\"name\":\"MantenimientoDeLoteRequest\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/MantenimientoDeLoteRechazado\",\"name\":\"Topic\",\"type\":\"string\"},{\"name\":\"razon\",\"type\":\"string\"}],\"name\":\"Andreani.ApiMantenimientoDeLote.Events.Record.MantenimientoDeLoteRechazado\",\"type\":\"record\"}"
}

func (r MantenimientoDeLoteRechazado) SchemaName() string {
	return "Andreani.ApiMantenimientoDeLote.Events.Record.MantenimientoDeLoteRechazado"
}

func (_ MantenimientoDeLoteRechazado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRechazado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRechazado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRechazado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRechazado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRechazado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRechazado) SetString(v string)   { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRechazado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MantenimientoDeLoteRechazado) Get(i int) types.Field {
	switch i {
	case 0:
		r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

		w := types.Record{Target: &r.SolicitudDeAccionAlmacen}

		return w

	case 1:
		r.MantenimientoDeLoteRequest = NewMantenimientoDeLoteRequest()

		w := types.Record{Target: &r.MantenimientoDeLoteRequest}

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

func (r *MantenimientoDeLoteRechazado) SetDefault(i int) {
	switch i {
	case 2:
		r.Topic = "Almacen/Solicitudes/MantenimientoDeLoteRechazado"
		return
	}
	panic("Unknown field index")
}

func (r *MantenimientoDeLoteRechazado) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ MantenimientoDeLoteRechazado) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ MantenimientoDeLoteRechazado) AppendArray() types.Field { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRechazado) HintSize(int)             { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRechazado) Finalize()                {}

func (_ MantenimientoDeLoteRechazado) AvroCRC64Fingerprint() []byte {
	return []byte(MantenimientoDeLoteRechazadoAvroCRC64Fingerprint)
}

func (r MantenimientoDeLoteRechazado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["solicitudDeAccionAlmacen"], err = json.Marshal(r.SolicitudDeAccionAlmacen)
	if err != nil {
		return nil, err
	}
	output["mantenimientoDeLoteRequest"], err = json.Marshal(r.MantenimientoDeLoteRequest)
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

func (r *MantenimientoDeLoteRechazado) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["mantenimientoDeLoteRequest"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MantenimientoDeLoteRequest); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for mantenimientoDeLoteRequest")
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
		r.Topic = "Almacen/Solicitudes/MantenimientoDeLoteRechazado"
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
