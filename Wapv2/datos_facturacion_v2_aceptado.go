// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     DatosFacturacionV2Aceptado.avsc
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

type DatosFacturacionV2Aceptado struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	Pedido Pedido `json:"pedido"`

	Topic string `json:"Topic"`
}

const DatosFacturacionV2AceptadoAvroCRC64Fingerprint = "|\xdb\xf9\x9b\xa6\xfd^\x94"

func NewDatosFacturacionV2Aceptado() DatosFacturacionV2Aceptado {
	r := DatosFacturacionV2Aceptado{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.Pedido = NewPedido()

	r.Topic = "Almacen/Solicitudes/DatosFacturacionV2Aceptado"
	return r
}

func DeserializeDatosFacturacionV2Aceptado(r io.Reader) (DatosFacturacionV2Aceptado, error) {
	t := NewDatosFacturacionV2Aceptado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDatosFacturacionV2AceptadoFromSchema(r io.Reader, schema string) (DatosFacturacionV2Aceptado, error) {
	t := NewDatosFacturacionV2Aceptado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDatosFacturacionV2Aceptado(r DatosFacturacionV2Aceptado, w io.Writer) error {
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
	return err
}

func (r DatosFacturacionV2Aceptado) Serialize(w io.Writer) error {
	return writeDatosFacturacionV2Aceptado(r, w)
}

func (r DatosFacturacionV2Aceptado) Schema() string {
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"eventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"default\":null,\"name\":\"instancia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoDistribucion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"pedido\",\"type\":{\"fields\":[{\"name\":\"numeroOrdenExterna\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"default\":null,\"name\":\"valorDeclarado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"linkImpresionRemito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"archivoImpresionRemito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tieneGestionCobranza\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"facturaLegal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaDeFacturacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"precioValorFC\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cot\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaEntrega\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"datosAdicionales\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]}],\"name\":\"Pedido\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/DatosFacturacionV2Aceptado\",\"name\":\"Topic\",\"type\":\"string\"}],\"name\":\"Andreani.Wapv2.Events.Record.DatosFacturacionV2Aceptado\",\"type\":\"record\"}"
}

func (r DatosFacturacionV2Aceptado) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.DatosFacturacionV2Aceptado"
}

func (_ DatosFacturacionV2Aceptado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DatosFacturacionV2Aceptado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DatosFacturacionV2Aceptado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DatosFacturacionV2Aceptado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DatosFacturacionV2Aceptado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DatosFacturacionV2Aceptado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DatosFacturacionV2Aceptado) SetString(v string)   { panic("Unsupported operation") }
func (_ DatosFacturacionV2Aceptado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DatosFacturacionV2Aceptado) Get(i int) types.Field {
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

	}
	panic("Unknown field index")
}

func (r *DatosFacturacionV2Aceptado) SetDefault(i int) {
	switch i {
	case 2:
		r.Topic = "Almacen/Solicitudes/DatosFacturacionV2Aceptado"
		return
	}
	panic("Unknown field index")
}

func (r *DatosFacturacionV2Aceptado) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ DatosFacturacionV2Aceptado) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DatosFacturacionV2Aceptado) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DatosFacturacionV2Aceptado) HintSize(int)                     { panic("Unsupported operation") }
func (_ DatosFacturacionV2Aceptado) Finalize()                        {}

func (_ DatosFacturacionV2Aceptado) AvroCRC64Fingerprint() []byte {
	return []byte(DatosFacturacionV2AceptadoAvroCRC64Fingerprint)
}

func (r DatosFacturacionV2Aceptado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["solicitudDeAccionAlmacen"], err = json.Marshal(r.SolicitudDeAccionAlmacen)
	if err != nil {
		return nil, err
	}
	output["pedido"], err = json.Marshal(r.Pedido)
	if err != nil {
		return nil, err
	}
	output["Topic"], err = json.Marshal(r.Topic)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DatosFacturacionV2Aceptado) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["pedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Pedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for pedido")
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
		r.Topic = "Almacen/Solicitudes/DatosFacturacionV2Aceptado"
	}
	return nil
}
