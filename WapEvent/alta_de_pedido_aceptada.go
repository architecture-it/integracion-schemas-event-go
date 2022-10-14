// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     AltaDePedidoAceptada.avsc
 */
package WapEventEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type AltaDePedidoAceptada struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	PedidoDeAlmacen PedidoDeAlmacenSCE `json:"pedidoDeAlmacen"`

	Topic string `json:"Topic"`
}

const AltaDePedidoAceptadaAvroCRC64Fingerprint = "\xd9L8\xbc\xf8U\xef\x17"

func NewAltaDePedidoAceptada() AltaDePedidoAceptada {
	r := AltaDePedidoAceptada{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.PedidoDeAlmacen = NewPedidoDeAlmacenSCE()

	r.Topic = "Almacen/Solicitudes/AltaDePedidoAceptada"
	return r
}

func DeserializeAltaDePedidoAceptada(r io.Reader) (AltaDePedidoAceptada, error) {
	t := NewAltaDePedidoAceptada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAltaDePedidoAceptadaFromSchema(r io.Reader, schema string) (AltaDePedidoAceptada, error) {
	t := NewAltaDePedidoAceptada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAltaDePedidoAceptada(r AltaDePedidoAceptada, w io.Writer) error {
	var err error
	err = writeSolicitudDeAccionAlmacen(r.SolicitudDeAccionAlmacen, w)
	if err != nil {
		return err
	}
	err = writePedidoDeAlmacenSCE(r.PedidoDeAlmacen, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Topic, w)
	if err != nil {
		return err
	}
	return err
}

func (r AltaDePedidoAceptada) Serialize(w io.Writer) error {
	return writeAltaDePedidoAceptada(r, w)
}

func (r AltaDePedidoAceptada) Schema() string {
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"eventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"name\":\"contrato\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"name\":\"planta\",\"type\":\"string\"}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"pedidoDeAlmacen\",\"type\":{\"fields\":[{\"name\":\"destinatario\",\"type\":{\"fields\":[{\"name\":\"idDestinatario\",\"type\":\"string\"},{\"name\":\"datosPersonales\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idinternocliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"eMail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"tipo\",\"type\":[\"null\",\"string\"]},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDeTelefonos\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"contacto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoDeDocumento\",\"type\":[\"null\",{\"name\":\"TipoDeDocumento\",\"symbols\":[\"undefined\",\"DNI\",\"CUIT\",\"CUIL\"],\"type\":\"enum\"}]}],\"name\":\"DatosPersonales\",\"type\":\"record\"}},{\"name\":\"direccion\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"referenciadedomicilio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"provincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]}],\"name\":\"Direccion\",\"type\":\"record\"}},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"contacto\",\"type\":[\"null\",\"string\"]}],\"name\":\"Destinatario\",\"type\":\"record\"}},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"numeroOrdenExterna\",\"type\":\"string\"},{\"name\":\"contratowarehouse\",\"type\":\"string\"},{\"name\":\"almacencliente\",\"type\":\"string\"},{\"default\":null,\"name\":\"prioridadPreparacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaPedido\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"fechaEntrega\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"franjahoraria\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"fechadeexpedicionSolicitada\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"cambioLoteDirigido\",\"type\":[\"null\",\"string\"]},{\"name\":\"tipo\",\"type\":\"string\"},{\"default\":null,\"name\":\"valorDeclarado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idAndreani\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"facturaLegal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechadeFacturacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"preciovalorFC\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ordenDeCompra\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tieneGestionCobranza\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"socioComercial\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cot\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"notas\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"datosadicionales\",\"type\":[\"null\",\"Wap.Events.Record.ListaDePropiedades\"]},{\"name\":\"listaDetalles\",\"type\":{\"fields\":[{\"name\":\"detallePedido\",\"type\":{\"items\":{\"fields\":[{\"name\":\"articulo\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"cantidad\",\"type\":\"double\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"numeropedido\",\"type\":[\"null\",\"string\"]},{\"name\":\"unidadmedida\",\"type\":\"string\"},{\"name\":\"lineaexterna\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"datosadicionales\",\"type\":[\"null\",\"Wap.Events.Record.ListaDePropiedades\"]},{\"name\":\"lote\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaDeVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",\"string\"]}],\"name\":\"LoteArticulo\",\"type\":\"record\"}}],\"name\":\"Articulo\",\"type\":\"record\"}},{\"name\":\"numerodelinea\",\"type\":[\"null\",\"string\"]},{\"name\":\"tipoacondicionamientoescundario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"admitepickingparcial\",\"type\":[\"null\",\"string\"]}],\"name\":\"DetallePedido\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDeDetalleDePedido\",\"type\":\"record\"}},{\"name\":\"distribuidor\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"etiqueta\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numerodeenvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"datosadicionales\",\"type\":[\"null\",\"string\"]}],\"name\":\"Distribuidor\",\"type\":\"record\"}}],\"name\":\"PedidoDeAlmacenSCE\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/AltaDePedidoAceptada\",\"name\":\"Topic\",\"type\":\"string\"}],\"name\":\"Wap.Events.Record.AltaDePedidoAceptada\",\"type\":\"record\"}"
}

func (r AltaDePedidoAceptada) SchemaName() string {
	return "Wap.Events.Record.AltaDePedidoAceptada"
}

func (_ AltaDePedidoAceptada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ AltaDePedidoAceptada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ AltaDePedidoAceptada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ AltaDePedidoAceptada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ AltaDePedidoAceptada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ AltaDePedidoAceptada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ AltaDePedidoAceptada) SetString(v string)   { panic("Unsupported operation") }
func (_ AltaDePedidoAceptada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AltaDePedidoAceptada) Get(i int) types.Field {
	switch i {
	case 0:
		r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

		w := types.Record{Target: &r.SolicitudDeAccionAlmacen}

		return w

	case 1:
		r.PedidoDeAlmacen = NewPedidoDeAlmacenSCE()

		w := types.Record{Target: &r.PedidoDeAlmacen}

		return w

	case 2:
		w := types.String{Target: &r.Topic}

		return w

	}
	panic("Unknown field index")
}

func (r *AltaDePedidoAceptada) SetDefault(i int) {
	switch i {
	case 2:
		r.Topic = "Almacen/Solicitudes/AltaDePedidoAceptada"
		return
	}
	panic("Unknown field index")
}

func (r *AltaDePedidoAceptada) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ AltaDePedidoAceptada) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ AltaDePedidoAceptada) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ AltaDePedidoAceptada) HintSize(int)                     { panic("Unsupported operation") }
func (_ AltaDePedidoAceptada) Finalize()                        {}

func (_ AltaDePedidoAceptada) AvroCRC64Fingerprint() []byte {
	return []byte(AltaDePedidoAceptadaAvroCRC64Fingerprint)
}

func (r AltaDePedidoAceptada) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["solicitudDeAccionAlmacen"], err = json.Marshal(r.SolicitudDeAccionAlmacen)
	if err != nil {
		return nil, err
	}
	output["pedidoDeAlmacen"], err = json.Marshal(r.PedidoDeAlmacen)
	if err != nil {
		return nil, err
	}
	output["Topic"], err = json.Marshal(r.Topic)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *AltaDePedidoAceptada) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["pedidoDeAlmacen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PedidoDeAlmacen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for pedidoDeAlmacen")
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
		r.Topic = "Almacen/Solicitudes/AltaDePedidoAceptada"
	}
	return nil
}
