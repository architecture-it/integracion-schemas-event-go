// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     MantenimientoDeLoteSolicitado.avsc
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

type MantenimientoDeLoteSolicitado struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	DetalleDeLote DetalleDeLote `json:"detalleDeLote"`

	Topic string `json:"Topic"`
}

const MantenimientoDeLoteSolicitadoAvroCRC64Fingerprint = "\x1f\xd1ə\xde\x1eӶ"

func NewMantenimientoDeLoteSolicitado() MantenimientoDeLoteSolicitado {
	r := MantenimientoDeLoteSolicitado{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.DetalleDeLote = NewDetalleDeLote()

	r.Topic = "Almacen/Solicitudes/MantenimientoDeLoteSolicitado"
	return r
}

func DeserializeMantenimientoDeLoteSolicitado(r io.Reader) (MantenimientoDeLoteSolicitado, error) {
	t := NewMantenimientoDeLoteSolicitado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeMantenimientoDeLoteSolicitadoFromSchema(r io.Reader, schema string) (MantenimientoDeLoteSolicitado, error) {
	t := NewMantenimientoDeLoteSolicitado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeMantenimientoDeLoteSolicitado(r MantenimientoDeLoteSolicitado, w io.Writer) error {
	var err error
	err = writeSolicitudDeAccionAlmacen(r.SolicitudDeAccionAlmacen, w)
	if err != nil {
		return err
	}
	err = writeDetalleDeLote(r.DetalleDeLote, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Topic, w)
	if err != nil {
		return err
	}
	return err
}

func (r MantenimientoDeLoteSolicitado) Serialize(w io.Writer) error {
	return writeMantenimientoDeLoteSolicitado(r, w)
}

func (r MantenimientoDeLoteSolicitado) Schema() string {
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"eventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"name\":\"contrato\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"name\":\"planta\",\"type\":\"string\"},{\"default\":null,\"name\":\"planta_instancia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoDistribucion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"detalleDeLote\",\"type\":{\"fields\":[{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"articulo\",\"type\":\"string\"},{\"name\":\"paquete\",\"type\":\"string\"},{\"name\":\"loteCaja\",\"type\":\"string\"},{\"name\":\"loteSecundario\",\"type\":\"string\"},{\"name\":\"loteSap\",\"type\":\"string\"},{\"name\":\"fechaFabricacion\",\"type\":\"string\"},{\"name\":\"fechaVencimiento\",\"type\":\"string\"},{\"name\":\"trazable\",\"type\":\"string\"},{\"name\":\"estado\",\"type\":\"string\"},{\"name\":\"procedencia\",\"type\":\"string\"},{\"name\":\"campoLibre1\",\"type\":\"string\"},{\"name\":\"campoLibre2\",\"type\":\"string\"},{\"name\":\"campoLibre3\",\"type\":\"string\"},{\"name\":\"campoLibre4\",\"type\":\"string\"},{\"name\":\"campoLibre5\",\"type\":\"string\"},{\"name\":\"loteExternoCliente\",\"type\":\"string\"},{\"name\":\"deliverByDate\",\"type\":\"string\"},{\"name\":\"bestByDate\",\"type\":\"string\"},{\"name\":\"fechaCreacion\",\"type\":\"string\"},{\"name\":\"usuarioCreacion\",\"type\":\"string\"},{\"name\":\"fechaEdicion\",\"type\":\"string\"},{\"name\":\"usuarioEdicion\",\"type\":\"string\"}],\"name\":\"DetalleDeLote\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/MantenimientoDeLoteSolicitado\",\"name\":\"Topic\",\"type\":\"string\"}],\"name\":\"Andreani.Wap.Events.Record.MantenimientoDeLoteSolicitado\",\"type\":\"record\"}"
}

func (r MantenimientoDeLoteSolicitado) SchemaName() string {
	return "Andreani.Wap.Events.Record.MantenimientoDeLoteSolicitado"
}

func (_ MantenimientoDeLoteSolicitado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ MantenimientoDeLoteSolicitado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ MantenimientoDeLoteSolicitado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ MantenimientoDeLoteSolicitado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ MantenimientoDeLoteSolicitado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ MantenimientoDeLoteSolicitado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ MantenimientoDeLoteSolicitado) SetString(v string)   { panic("Unsupported operation") }
func (_ MantenimientoDeLoteSolicitado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MantenimientoDeLoteSolicitado) Get(i int) types.Field {
	switch i {
	case 0:
		r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

		w := types.Record{Target: &r.SolicitudDeAccionAlmacen}

		return w

	case 1:
		r.DetalleDeLote = NewDetalleDeLote()

		w := types.Record{Target: &r.DetalleDeLote}

		return w

	case 2:
		w := types.String{Target: &r.Topic}

		return w

	}
	panic("Unknown field index")
}

func (r *MantenimientoDeLoteSolicitado) SetDefault(i int) {
	switch i {
	case 2:
		r.Topic = "Almacen/Solicitudes/MantenimientoDeLoteSolicitado"
		return
	}
	panic("Unknown field index")
}

func (r *MantenimientoDeLoteSolicitado) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ MantenimientoDeLoteSolicitado) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ MantenimientoDeLoteSolicitado) AppendArray() types.Field { panic("Unsupported operation") }
func (_ MantenimientoDeLoteSolicitado) HintSize(int)             { panic("Unsupported operation") }
func (_ MantenimientoDeLoteSolicitado) Finalize()                {}

func (_ MantenimientoDeLoteSolicitado) AvroCRC64Fingerprint() []byte {
	return []byte(MantenimientoDeLoteSolicitadoAvroCRC64Fingerprint)
}

func (r MantenimientoDeLoteSolicitado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["solicitudDeAccionAlmacen"], err = json.Marshal(r.SolicitudDeAccionAlmacen)
	if err != nil {
		return nil, err
	}
	output["detalleDeLote"], err = json.Marshal(r.DetalleDeLote)
	if err != nil {
		return nil, err
	}
	output["Topic"], err = json.Marshal(r.Topic)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *MantenimientoDeLoteSolicitado) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["detalleDeLote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DetalleDeLote); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for detalleDeLote")
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
		r.Topic = "Almacen/Solicitudes/MantenimientoDeLoteSolicitado"
	}
	return nil
}
