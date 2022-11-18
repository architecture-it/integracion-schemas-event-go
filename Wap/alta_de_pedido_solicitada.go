// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     AltaDePedidoSolicitada.avsc
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

type AltaDePedidoSolicitada struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	PedidoDeAlmacen PedidoDeAlmacenSCE `json:"pedidoDeAlmacen"`

	Topic string `json:"Topic"`
}

const AltaDePedidoSolicitadaAvroCRC64Fingerprint = "\xcfX-\xab..\xa4\xce"

func NewAltaDePedidoSolicitada() AltaDePedidoSolicitada {
	r := AltaDePedidoSolicitada{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.PedidoDeAlmacen = NewPedidoDeAlmacenSCE()

	r.Topic = "Almacen/Solicitudes/AltaDePedidoSolicitada"
	return r
}

func DeserializeAltaDePedidoSolicitada(r io.Reader) (AltaDePedidoSolicitada, error) {
	t := NewAltaDePedidoSolicitada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAltaDePedidoSolicitadaFromSchema(r io.Reader, schema string) (AltaDePedidoSolicitada, error) {
	t := NewAltaDePedidoSolicitada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAltaDePedidoSolicitada(r AltaDePedidoSolicitada, w io.Writer) error {
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

func (r AltaDePedidoSolicitada) Serialize(w io.Writer) error {
	return writeAltaDePedidoSolicitada(r, w)
}

func (r AltaDePedidoSolicitada) Schema() string {
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"eventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"default\":null,\"name\":\"instancia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoDistribucion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"pedidoDeAlmacen\",\"type\":{\"fields\":[{\"name\":\"destinatario\",\"type\":{\"fields\":[{\"name\":\"datosPersonales\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idInternoDelCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"eMail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"tipo\",\"type\":[\"null\",\"string\"]},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDeTelefonos\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"agrupador\",\"type\":[\"null\",\"string\"]},{\"name\":\"tipoDeDocumento\",\"type\":{\"name\":\"TipoDeDocumento\",\"symbols\":[\"undefined\",\"DNI\",\"CUIT\",\"CUIL\"],\"type\":\"enum\"}}],\"name\":\"DatosPersonales\",\"type\":\"record\"}},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"contacto\",\"type\":[\"null\",\"string\"]}],\"name\":\"Destinatario\",\"type\":\"record\"}},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"numeroPedido\",\"type\":\"string\"},{\"default\":null,\"name\":\"fechaPedido\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"fechaEntrega\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"remito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"referenciaCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoTransportista\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"notas\",\"type\":[\"null\",\"Andreani.Wap.Events.Record.ListaDePropiedades\"]},{\"default\":null,\"name\":\"camposLibres\",\"type\":[\"null\",\"Andreani.Wap.Events.Record.ListaDePropiedades\"]},{\"default\":null,\"name\":\"socioComercial\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"modoTransporte\",\"type\":[\"null\",\"string\"]},{\"name\":\"direccion\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"abreviaturaProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoDeDireccion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoCiudad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefono\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoISOProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoISOPais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",\"Andreani.Wap.Events.Record.ListaDePropiedades\"]}],\"name\":\"Direccion\",\"type\":\"record\"}},{\"name\":\"listaDetalles\",\"type\":{\"fields\":[{\"name\":\"detallePedido\",\"type\":{\"items\":{\"fields\":[{\"name\":\"articulo\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"default\":null,\"name\":\"lote\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaDeVencimiento\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",\"Andreani.Wap.Events.Record.ListaDePropiedades\"]}],\"name\":\"LoteArticulo\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",\"Andreani.Wap.Events.Record.ListaDePropiedades\"]}],\"name\":\"Articulo\",\"type\":\"record\"}},{\"name\":\"numeroPedido\",\"type\":\"string\"},{\"name\":\"unidadMedida\",\"type\":\"string\"},{\"default\":null,\"name\":\"lineaExterna\",\"type\":[\"null\",\"string\"]},{\"name\":\"unidades\",\"type\":\"double\"},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",\"Andreani.Wap.Events.Record.ListaDePropiedades\"]}],\"name\":\"DetallePedido\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDeDetalleDePedido\",\"type\":\"record\"}},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",\"Andreani.Wap.Events.Record.ListaDePropiedades\"]},{\"default\":null,\"name\":\"zonaConsumo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"facturaLegal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"facturaInterna\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tieneGestionCobranza\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"valorSeguro\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ordenDeCompra\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"clientePadre\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"acondicionamientoSecundario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"imprimeDocumentacion\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"detalleDeOrdenDeCompra\",\"type\":[\"null\",{\"fields\":[{\"name\":\"numeroDeLineaDeCliente\",\"type\":\"string\"},{\"name\":\"ordenDeCompraDeCliente\",\"type\":\"string\"},{\"default\":null,\"name\":\"notasDeLinea\",\"type\":[\"null\",\"string\"]},{\"name\":\"numeroDeLinea\",\"type\":\"string\"},{\"name\":\"cantidadPedida\",\"type\":\"double\"},{\"name\":\"articulo\",\"type\":\"Andreani.Wap.Events.Record.Articulo\"},{\"default\":null,\"name\":\"valorDeclarado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"UOM\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"camposLibres\",\"type\":[\"null\",\"Andreani.Wap.Events.Record.ListaDePropiedades\"]},{\"default\":null,\"name\":\"fechaOrdenDeCompra\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoMaterial\",\"type\":[\"null\",\"string\"]}],\"name\":\"DetalleDeOrdenDeCompra\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"linkImpresionRemito\",\"type\":[\"null\",\"string\"]}],\"name\":\"PedidoDeAlmacenSCE\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/AltaDePedidoSolicitada\",\"name\":\"Topic\",\"type\":\"string\"}],\"name\":\"Andreani.Wap.Events.Record.AltaDePedidoSolicitada\",\"type\":\"record\"}"
}

func (r AltaDePedidoSolicitada) SchemaName() string {
	return "Andreani.Wap.Events.Record.AltaDePedidoSolicitada"
}

func (_ AltaDePedidoSolicitada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ AltaDePedidoSolicitada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ AltaDePedidoSolicitada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ AltaDePedidoSolicitada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ AltaDePedidoSolicitada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ AltaDePedidoSolicitada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ AltaDePedidoSolicitada) SetString(v string)   { panic("Unsupported operation") }
func (_ AltaDePedidoSolicitada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AltaDePedidoSolicitada) Get(i int) types.Field {
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

func (r *AltaDePedidoSolicitada) SetDefault(i int) {
	switch i {
	case 2:
		r.Topic = "Almacen/Solicitudes/AltaDePedidoSolicitada"
		return
	}
	panic("Unknown field index")
}

func (r *AltaDePedidoSolicitada) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ AltaDePedidoSolicitada) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ AltaDePedidoSolicitada) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ AltaDePedidoSolicitada) HintSize(int)                     { panic("Unsupported operation") }
func (_ AltaDePedidoSolicitada) Finalize()                        {}

func (_ AltaDePedidoSolicitada) AvroCRC64Fingerprint() []byte {
	return []byte(AltaDePedidoSolicitadaAvroCRC64Fingerprint)
}

func (r AltaDePedidoSolicitada) MarshalJSON() ([]byte, error) {
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

func (r *AltaDePedidoSolicitada) UnmarshalJSON(data []byte) error {
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
		r.Topic = "Almacen/Solicitudes/AltaDePedidoSolicitada"
	}
	return nil
}
