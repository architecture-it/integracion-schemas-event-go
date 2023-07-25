// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     AltaDePedidoRechazada.avsc
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

type AltaDePedidoRechazada struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	PedidoDeAlmacen PedidoDeAlmacenSCE `json:"pedidoDeAlmacen"`

	Topic string `json:"Topic"`

	Razon string `json:"razon"`
}

const AltaDePedidoRechazadaAvroCRC64Fingerprint = "nĹ\xd4\xd1\xe3\xa4y"

func NewAltaDePedidoRechazada() AltaDePedidoRechazada {
	r := AltaDePedidoRechazada{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.PedidoDeAlmacen = NewPedidoDeAlmacenSCE()

	r.Topic = "Almacen/Solicitudes/AltaDePedidoRechazada"
	return r
}

func DeserializeAltaDePedidoRechazada(r io.Reader) (AltaDePedidoRechazada, error) {
	t := NewAltaDePedidoRechazada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAltaDePedidoRechazadaFromSchema(r io.Reader, schema string) (AltaDePedidoRechazada, error) {
	t := NewAltaDePedidoRechazada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAltaDePedidoRechazada(r AltaDePedidoRechazada, w io.Writer) error {
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
	err = vm.WriteString(r.Razon, w)
	if err != nil {
		return err
	}
	return err
}

func (r AltaDePedidoRechazada) Serialize(w io.Writer) error {
	return writeAltaDePedidoRechazada(r, w)
}

func (r AltaDePedidoRechazada) Schema() string {
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"eventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"default\":null,\"name\":\"instancia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoDistribucion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"pedidoDeAlmacen\",\"type\":{\"fields\":[{\"name\":\"destinatario\",\"type\":{\"fields\":[{\"name\":\"idDestinatario\",\"type\":\"string\"},{\"name\":\"datosPersonales\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idinternocliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"eMail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"tipo\",\"type\":[\"null\",\"string\"]},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDeTelefonos\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"contacto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoDeDocumento\",\"type\":[\"null\",{\"name\":\"TipoDeDocumento\",\"symbols\":[\"undefined\",\"DNI\",\"CUIT\",\"CUIL\"],\"type\":\"enum\"}]}],\"name\":\"DatosPersonales\",\"type\":\"record\"}},{\"name\":\"direccion\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"referenciadedomicilio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"provincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]}],\"name\":\"Direccion\",\"type\":\"record\"}},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"contacto\",\"type\":[\"null\",\"string\"]}],\"name\":\"Destinatario\",\"type\":\"record\"}},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"numeroOrdenExterna\",\"type\":\"string\"},{\"name\":\"almacencliente\",\"type\":\"string\"},{\"name\":\"fechaPedido\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"fechaEntrega\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"franjahoraria\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechadeexpedicionSolicitada\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"cambioLoteDirigido\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"valorDeclarado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"facturaLegal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechadeFacturacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"preciovalorFC\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ordenDeCompra\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tieneGestionCobranza\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cot\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"notas\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estatusOTdeTraza\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"marketPlace\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"datosadicionales\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.ListaDePropiedades\"]},{\"name\":\"listaDetalles\",\"type\":{\"fields\":[{\"name\":\"detallePedido\",\"type\":{\"items\":{\"fields\":[{\"name\":\"articulo\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"cantidad\",\"type\":\"double\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"numeropedido\",\"type\":[\"null\",\"string\"]},{\"name\":\"zonaConsumo\",\"type\":[\"null\",\"string\"]},{\"name\":\"unidadmedida\",\"type\":\"string\"},{\"default\":null,\"name\":\"datosadicionales\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.ListaDePropiedades\"]},{\"name\":\"lote\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaDeVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",\"string\"]},{\"name\":\"estado\",\"type\":[\"null\",\"string\"]}],\"name\":\"LoteArticulo\",\"type\":\"record\"}}],\"name\":\"Articulo\",\"type\":\"record\"}},{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]},{\"name\":\"numerodelinea\",\"type\":[\"null\",\"string\"]},{\"name\":\"tipoacondicionamientoescundario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"admitepickingparcial\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estatusOTdeAcondi\",\"type\":[\"null\",\"string\"]},{\"name\":\"diasVencimientoMinimo\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"tiposDeAcondicionamientoSecundario\",\"type\":[\"null\",{\"fields\":[{\"name\":\"tiposDeAcondicionamientoSecundario\",\"type\":{\"items\":{\"fields\":[{\"name\":\"tipoDeAcondi\",\"type\":\"string\"}],\"name\":\"TiposDeAcondicionamientoSecundario\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDeTiposDeAcondicionamientoSecundario\",\"type\":\"record\"}]}],\"name\":\"DetallePedido\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDeDetalleDePedido\",\"type\":\"record\"}},{\"name\":\"distribuidor\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"etiqueta\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numerodeenvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"datosadicionales\",\"type\":[\"null\",\"string\"]}],\"name\":\"Distribuidor\",\"type\":\"record\"}},{\"default\":null,\"name\":\"linkImpresionRemito\",\"type\":[\"null\",{\"fields\":[{\"name\":\"documentos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"ImpRemito\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDeImpRemito\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"codigosDeTrazaANMAT\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigosDeTrazaANMAT\",\"type\":{\"items\":{\"fields\":[{\"name\":\"codigoDeTrazaANMAT\",\"type\":\"int\"}],\"name\":\"CodigosDeTrazaANMAT\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDeCodigosDeTrazaANMAT\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"archivos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"documentos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"numeroEntidad\",\"type\":\"string\"},{\"name\":\"documento\",\"type\":\"string\"}],\"name\":\"File\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDeFiles\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"archivoRemito\",\"type\":[\"null\",{\"fields\":[{\"name\":\"documentos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"numeroEntidad\",\"type\":\"string\"},{\"name\":\"documento\",\"type\":\"string\"}],\"name\":\"Remito\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDeRemitos\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"urlsDocumentos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"documentos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"url\",\"type\":\"string\"}],\"name\":\"UrlDocumento\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaUrlsDocumentos\",\"type\":\"record\"}]}],\"name\":\"PedidoDeAlmacenSCE\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/AltaDePedidoRechazada\",\"name\":\"Topic\",\"type\":\"string\"},{\"name\":\"razon\",\"type\":\"string\"}],\"name\":\"Andreani.Wapv2.Events.Record.AltaDePedidoRechazada\",\"type\":\"record\"}"
}

func (r AltaDePedidoRechazada) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.AltaDePedidoRechazada"
}

func (_ AltaDePedidoRechazada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ AltaDePedidoRechazada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ AltaDePedidoRechazada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ AltaDePedidoRechazada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ AltaDePedidoRechazada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ AltaDePedidoRechazada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ AltaDePedidoRechazada) SetString(v string)   { panic("Unsupported operation") }
func (_ AltaDePedidoRechazada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AltaDePedidoRechazada) Get(i int) types.Field {
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

	case 3:
		w := types.String{Target: &r.Razon}

		return w

	}
	panic("Unknown field index")
}

func (r *AltaDePedidoRechazada) SetDefault(i int) {
	switch i {
	case 2:
		r.Topic = "Almacen/Solicitudes/AltaDePedidoRechazada"
		return
	}
	panic("Unknown field index")
}

func (r *AltaDePedidoRechazada) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ AltaDePedidoRechazada) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ AltaDePedidoRechazada) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ AltaDePedidoRechazada) HintSize(int)                     { panic("Unsupported operation") }
func (_ AltaDePedidoRechazada) Finalize()                        {}

func (_ AltaDePedidoRechazada) AvroCRC64Fingerprint() []byte {
	return []byte(AltaDePedidoRechazadaAvroCRC64Fingerprint)
}

func (r AltaDePedidoRechazada) MarshalJSON() ([]byte, error) {
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
	output["razon"], err = json.Marshal(r.Razon)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *AltaDePedidoRechazada) UnmarshalJSON(data []byte) error {
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
		r.Topic = "Almacen/Solicitudes/AltaDePedidoRechazada"
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
