// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     DatosFacturacionAceptadoV2.avsc
 */
package ApiDatosFacturacionv2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type DatosFacturacionAceptadoV2 struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	ActualizacionDatosFacturacion ActualizacionDatosFacturacion `json:"ActualizacionDatosFacturacion"`

	Topic string `json:"Topic"`
}

const DatosFacturacionAceptadoV2AvroCRC64Fingerprint = "-\x06vܐ\t\xd7{"

func NewDatosFacturacionAceptadoV2() DatosFacturacionAceptadoV2 {
	r := DatosFacturacionAceptadoV2{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.ActualizacionDatosFacturacion = NewActualizacionDatosFacturacion()

	r.Topic = "Almacen/Solicitudes/DatosFacturacionAceptadoV2"
	return r
}

func DeserializeDatosFacturacionAceptadoV2(r io.Reader) (DatosFacturacionAceptadoV2, error) {
	t := NewDatosFacturacionAceptadoV2()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDatosFacturacionAceptadoV2FromSchema(r io.Reader, schema string) (DatosFacturacionAceptadoV2, error) {
	t := NewDatosFacturacionAceptadoV2()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDatosFacturacionAceptadoV2(r DatosFacturacionAceptadoV2, w io.Writer) error {
	var err error
	err = writeSolicitudDeAccionAlmacen(r.SolicitudDeAccionAlmacen, w)
	if err != nil {
		return err
	}
	err = writeActualizacionDatosFacturacion(r.ActualizacionDatosFacturacion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Topic, w)
	if err != nil {
		return err
	}
	return err
}

func (r DatosFacturacionAceptadoV2) Serialize(w io.Writer) error {
	return writeDatosFacturacionAceptadoV2(r, w)
}

func (r DatosFacturacionAceptadoV2) Schema() string {
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"IdTransaccion\",\"type\":\"string\"},{\"name\":\"Contrato\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Planta\",\"type\":\"string\"},{\"name\":\"EventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"Timestamp\",\"type\":\"int\"},{\"name\":\"Remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"Destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Vencimiento\",\"type\":[\"null\",\"string\"]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"ActualizacionDatosFacturacion\",\"type\":{\"fields\":[{\"name\":\"OrdenExterna\",\"type\":\"string\"},{\"name\":\"Remito\",\"type\":\"string\"},{\"name\":\"Factura\",\"type\":\"string\"},{\"default\":null,\"name\":\"ValorDeclarado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ValorSeguro\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaFactura\",\"type\":[\"null\",\"long\"]}],\"name\":\"ActualizacionDatosFacturacion\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/DatosFacturacionAceptadoV2\",\"name\":\"Topic\",\"type\":\"string\"}],\"name\":\"Andreani.WapDatosFacturacionv2.Events.Record.DatosFacturacionAceptadoV2\",\"type\":\"record\"}"
}

func (r DatosFacturacionAceptadoV2) SchemaName() string {
	return "Andreani.WapDatosFacturacionv2.Events.Record.DatosFacturacionAceptadoV2"
}

func (_ DatosFacturacionAceptadoV2) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DatosFacturacionAceptadoV2) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DatosFacturacionAceptadoV2) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DatosFacturacionAceptadoV2) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DatosFacturacionAceptadoV2) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DatosFacturacionAceptadoV2) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DatosFacturacionAceptadoV2) SetString(v string)   { panic("Unsupported operation") }
func (_ DatosFacturacionAceptadoV2) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DatosFacturacionAceptadoV2) Get(i int) types.Field {
	switch i {
	case 0:
		r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

		w := types.Record{Target: &r.SolicitudDeAccionAlmacen}

		return w

	case 1:
		r.ActualizacionDatosFacturacion = NewActualizacionDatosFacturacion()

		w := types.Record{Target: &r.ActualizacionDatosFacturacion}

		return w

	case 2:
		w := types.String{Target: &r.Topic}

		return w

	}
	panic("Unknown field index")
}

func (r *DatosFacturacionAceptadoV2) SetDefault(i int) {
	switch i {
	case 2:
		r.Topic = "Almacen/Solicitudes/DatosFacturacionAceptadoV2"
		return
	}
	panic("Unknown field index")
}

func (r *DatosFacturacionAceptadoV2) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ DatosFacturacionAceptadoV2) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DatosFacturacionAceptadoV2) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DatosFacturacionAceptadoV2) HintSize(int)                     { panic("Unsupported operation") }
func (_ DatosFacturacionAceptadoV2) Finalize()                        {}

func (_ DatosFacturacionAceptadoV2) AvroCRC64Fingerprint() []byte {
	return []byte(DatosFacturacionAceptadoV2AvroCRC64Fingerprint)
}

func (r DatosFacturacionAceptadoV2) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["solicitudDeAccionAlmacen"], err = json.Marshal(r.SolicitudDeAccionAlmacen)
	if err != nil {
		return nil, err
	}
	output["ActualizacionDatosFacturacion"], err = json.Marshal(r.ActualizacionDatosFacturacion)
	if err != nil {
		return nil, err
	}
	output["Topic"], err = json.Marshal(r.Topic)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DatosFacturacionAceptadoV2) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["ActualizacionDatosFacturacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ActualizacionDatosFacturacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ActualizacionDatosFacturacion")
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
		r.Topic = "Almacen/Solicitudes/DatosFacturacionAceptadoV2"
	}
	return nil
}
