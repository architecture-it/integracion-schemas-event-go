// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     DatosFacturacionRechazadoV2.avsc
 */
package WapDatosFacturacionv2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type DatosFacturacionRechazadoV2 struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	Pedido Pedido `json:"Pedido"`

	Topic string `json:"Topic"`

	Razon string `json:"Razon"`
}

const DatosFacturacionRechazadoV2AvroCRC64Fingerprint = "\xea;\xd2\xfcR\xef\x02,"

func NewDatosFacturacionRechazadoV2() DatosFacturacionRechazadoV2 {
	r := DatosFacturacionRechazadoV2{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.Pedido = NewPedido()

	r.Topic = "Almacen/Solicitudes/DatosFacturacionRechazadoV2"
	return r
}

func DeserializeDatosFacturacionRechazadoV2(r io.Reader) (DatosFacturacionRechazadoV2, error) {
	t := NewDatosFacturacionRechazadoV2()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDatosFacturacionRechazadoV2FromSchema(r io.Reader, schema string) (DatosFacturacionRechazadoV2, error) {
	t := NewDatosFacturacionRechazadoV2()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDatosFacturacionRechazadoV2(r DatosFacturacionRechazadoV2, w io.Writer) error {
	var err error
	err = writeSolicitudDeAccionAlmacen(r.SolicitudDeAccionAlmacen, w)
	if err != nil {
		return err
	}
	err = writePedido(r.Pedido, w)
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

func (r DatosFacturacionRechazadoV2) Serialize(w io.Writer) error {
	return writeDatosFacturacionRechazadoV2(r, w)
}

func (r DatosFacturacionRechazadoV2) Schema() string {
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"IdTransaccion\",\"type\":\"string\"},{\"name\":\"ContratoDistribución\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"},{\"name\":\"EventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"Timestamp\",\"type\":\"int\"},{\"name\":\"Remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"Destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Vencimiento\",\"type\":[\"null\",\"string\"]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"Pedido\",\"type\":{\"fields\":[{\"name\":\"NumeroOrdenExterna\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"default\":null,\"name\":\"ValorDeclarado\",\"type\":[\"null\",\"string\"]},{\"name\":\"Remito\",\"type\":\"string\"},{\"default\":null,\"name\":\"LinkImpresionRemito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ArchivoImpresionRemito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TieneGestionCobranza\",\"type\":[\"null\",\"string\"]},{\"name\":\"FacturaLegal\",\"type\":\"string\"},{\"default\":null,\"name\":\"FechaDeFacturacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"PrecioValorFc\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"COT\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaEntrega\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"DatosAdicionales\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"DatosAdicionales\",\"type\":\"record\"}]}],\"name\":\"Pedido\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/DatosFacturacionRechazadoV2\",\"name\":\"Topic\",\"type\":\"string\"},{\"name\":\"Razon\",\"type\":\"string\"}],\"name\":\"Andreani.WapDatosFacturacionv2.Events.Record.DatosFacturacionRechazadoV2\",\"type\":\"record\"}"
}

func (r DatosFacturacionRechazadoV2) SchemaName() string {
	return "Andreani.WapDatosFacturacionv2.Events.Record.DatosFacturacionRechazadoV2"
}

func (_ DatosFacturacionRechazadoV2) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DatosFacturacionRechazadoV2) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DatosFacturacionRechazadoV2) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DatosFacturacionRechazadoV2) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DatosFacturacionRechazadoV2) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DatosFacturacionRechazadoV2) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DatosFacturacionRechazadoV2) SetString(v string)   { panic("Unsupported operation") }
func (_ DatosFacturacionRechazadoV2) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DatosFacturacionRechazadoV2) Get(i int) types.Field {
	switch i {
	case 0:
		r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

		w := types.Record{Target: &r.SolicitudDeAccionAlmacen}

		return w

	case 1:
		r.Pedido = NewPedido()

		w := types.Record{Target: &r.Pedido}

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

func (r *DatosFacturacionRechazadoV2) SetDefault(i int) {
	switch i {
	case 2:
		r.Topic = "Almacen/Solicitudes/DatosFacturacionRechazadoV2"
		return
	}
	panic("Unknown field index")
}

func (r *DatosFacturacionRechazadoV2) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ DatosFacturacionRechazadoV2) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ DatosFacturacionRechazadoV2) AppendArray() types.Field { panic("Unsupported operation") }
func (_ DatosFacturacionRechazadoV2) HintSize(int)             { panic("Unsupported operation") }
func (_ DatosFacturacionRechazadoV2) Finalize()                {}

func (_ DatosFacturacionRechazadoV2) AvroCRC64Fingerprint() []byte {
	return []byte(DatosFacturacionRechazadoV2AvroCRC64Fingerprint)
}

func (r DatosFacturacionRechazadoV2) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["solicitudDeAccionAlmacen"], err = json.Marshal(r.SolicitudDeAccionAlmacen)
	if err != nil {
		return nil, err
	}
	output["Pedido"], err = json.Marshal(r.Pedido)
	if err != nil {
		return nil, err
	}
	output["Topic"], err = json.Marshal(r.Topic)
	if err != nil {
		return nil, err
	}
	output["Razon"], err = json.Marshal(r.Razon)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DatosFacturacionRechazadoV2) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["Pedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Pedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Pedido")
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
		r.Topic = "Almacen/Solicitudes/DatosFacturacionRechazadoV2"
	}
	val = func() json.RawMessage {
		if v, ok := fields["Razon"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Razon); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Razon")
	}
	return nil
}
