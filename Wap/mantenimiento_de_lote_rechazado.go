// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     MantenimientoDeLoteRechazado.avsc
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

type MantenimientoDeLoteRechazado struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	DetalleDeLote DetalleDeLote `json:"detalleDeLote"`

	Topic string `json:"Topic"`

	Razon string `json:"razon"`
}

const MantenimientoDeLoteRechazadoAvroCRC64Fingerprint = "\x1e\a.\n\xbf\x9e^\xc9"

func NewMantenimientoDeLoteRechazado() MantenimientoDeLoteRechazado {
	r := MantenimientoDeLoteRechazado{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.DetalleDeLote = NewDetalleDeLote()

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
	err = writeDetalleDeLote(r.DetalleDeLote, w)
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
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"eventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"name\":\"contrato\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"name\":\"planta\",\"type\":\"string\"},{\"default\":null,\"name\":\"planta_instancia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoDistribucion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoWarehousePedido\",\"type\":[\"null\",\"string\"]}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"detalleDeLote\",\"type\":{\"fields\":[{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"articulo\",\"type\":\"string\"},{\"name\":\"paquete\",\"type\":\"string\"},{\"name\":\"loteCaja\",\"type\":\"string\"},{\"name\":\"loteSecundario\",\"type\":\"string\"},{\"name\":\"loteSap\",\"type\":\"string\"},{\"name\":\"fechaFabricacion\",\"type\":\"string\"},{\"name\":\"fechaVencimiento\",\"type\":\"string\"},{\"name\":\"trazable\",\"type\":\"string\"},{\"name\":\"estado\",\"type\":\"string\"},{\"name\":\"procedencia\",\"type\":\"string\"},{\"name\":\"campoLibre1\",\"type\":\"string\"},{\"name\":\"campoLibre2\",\"type\":\"string\"},{\"name\":\"campoLibre3\",\"type\":\"string\"},{\"name\":\"campoLibre4\",\"type\":\"string\"},{\"name\":\"campoLibre5\",\"type\":\"string\"},{\"name\":\"loteExternoCliente\",\"type\":\"string\"},{\"name\":\"deliverByDate\",\"type\":\"string\"},{\"name\":\"bestByDate\",\"type\":\"string\"},{\"name\":\"fechaCreacion\",\"type\":\"string\"},{\"name\":\"usuarioCreacion\",\"type\":\"string\"},{\"name\":\"fechaEdicion\",\"type\":\"string\"},{\"name\":\"usuarioEdicion\",\"type\":\"string\"}],\"name\":\"DetalleDeLote\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/MantenimientoDeLoteRechazado\",\"name\":\"Topic\",\"type\":\"string\"},{\"name\":\"razon\",\"type\":\"string\"}],\"name\":\"Andreani.Wap.Events.Record.MantenimientoDeLoteRechazado\",\"type\":\"record\"}"
}

func (r MantenimientoDeLoteRechazado) SchemaName() string {
	return "Andreani.Wap.Events.Record.MantenimientoDeLoteRechazado"
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
		r.DetalleDeLote = NewDetalleDeLote()

		w := types.Record{Target: &r.DetalleDeLote}

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
	output["detalleDeLote"], err = json.Marshal(r.DetalleDeLote)
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
