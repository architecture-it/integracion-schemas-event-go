// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoDeAlmacenSCE.avsc
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

type PedidoDeAlmacenSCE struct {
	Destinatario Destinatario `json:"destinatario"`

	Propietario string `json:"propietario"`

	NumeroOrdenExterna string `json:"numeroOrdenExterna"`

	Almacencliente string `json:"almacencliente"`

	FechaPedido int64 `json:"fechaPedido"`

	FechaEntrega *UnionNullLong `json:"fechaEntrega"`

	Franjahoraria *UnionNullString `json:"franjahoraria"`

	FechadeexpedicionSolicitada *UnionNullLong `json:"fechadeexpedicionSolicitada"`

	CambioLoteDirigido *UnionNullString `json:"cambioLoteDirigido"`

	ValorDeclarado *UnionNullString `json:"valorDeclarado"`

	Remito *UnionNullString `json:"remito"`

	FacturaLegal *UnionNullString `json:"facturaLegal"`

	FechadeFacturacion *UnionNullLong `json:"fechadeFacturacion"`

	PreciovalorFC *UnionNullString `json:"preciovalorFC"`

	OrdenDeCompra *UnionNullString `json:"ordenDeCompra"`

	TieneGestionCobranza *UnionNullString `json:"tieneGestionCobranza"`

	Cot *UnionNullString `json:"cot"`

	Notas *UnionNullString `json:"notas"`

	EstatusOTdeTraza *UnionNullString `json:"estatusOTdeTraza"`

	MarketPlace *UnionNullString `json:"marketPlace"`

	Datosadicionales *UnionNullListaDePropiedades `json:"datosadicionales"`

	ListaDetalles ListaDeDetalleDePedido `json:"listaDetalles"`

	Distribuidor Distribuidor `json:"distribuidor"`

	LinkImpresionRemito *UnionNullString `json:"linkImpresionRemito"`

	TiposDeAcondicionamientoSecundario *UnionNullListaDeTiposDeAcondicionamientoSecundario `json:"tiposDeAcondicionamientoSecundario"`
}

const PedidoDeAlmacenSCEAvroCRC64Fingerprint = "\xabP\t\xa8\x0eH\xfa\a"

func NewPedidoDeAlmacenSCE() PedidoDeAlmacenSCE {
	r := PedidoDeAlmacenSCE{}
	r.Destinatario = NewDestinatario()

	r.FechaEntrega = nil
	r.Franjahoraria = nil
	r.FechadeexpedicionSolicitada = nil
	r.CambioLoteDirigido = nil
	r.ValorDeclarado = nil
	r.Remito = nil
	r.FacturaLegal = nil
	r.FechadeFacturacion = nil
	r.PreciovalorFC = nil
	r.OrdenDeCompra = nil
	r.TieneGestionCobranza = nil
	r.Cot = nil
	r.Notas = nil
	r.EstatusOTdeTraza = nil
	r.MarketPlace = nil
	r.Datosadicionales = nil
	r.ListaDetalles = NewListaDeDetalleDePedido()

	r.Distribuidor = NewDistribuidor()

	r.LinkImpresionRemito = nil
	r.TiposDeAcondicionamientoSecundario = nil
	return r
}

func DeserializePedidoDeAlmacenSCE(r io.Reader) (PedidoDeAlmacenSCE, error) {
	t := NewPedidoDeAlmacenSCE()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePedidoDeAlmacenSCEFromSchema(r io.Reader, schema string) (PedidoDeAlmacenSCE, error) {
	t := NewPedidoDeAlmacenSCE()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePedidoDeAlmacenSCE(r PedidoDeAlmacenSCE, w io.Writer) error {
	var err error
	err = writeDestinatario(r.Destinatario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroOrdenExterna, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Almacencliente, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.FechaPedido, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaEntrega, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Franjahoraria, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechadeexpedicionSolicitada, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CambioLoteDirigido, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ValorDeclarado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Remito, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FacturaLegal, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechadeFacturacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.PreciovalorFC, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.OrdenDeCompra, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TieneGestionCobranza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Cot, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Notas, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstatusOTdeTraza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.MarketPlace, w)
	if err != nil {
		return err
	}
	err = writeUnionNullListaDePropiedades(r.Datosadicionales, w)
	if err != nil {
		return err
	}
	err = writeListaDeDetalleDePedido(r.ListaDetalles, w)
	if err != nil {
		return err
	}
	err = writeDistribuidor(r.Distribuidor, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LinkImpresionRemito, w)
	if err != nil {
		return err
	}
	err = writeUnionNullListaDeTiposDeAcondicionamientoSecundario(r.TiposDeAcondicionamientoSecundario, w)
	if err != nil {
		return err
	}
	return err
}

func (r PedidoDeAlmacenSCE) Serialize(w io.Writer) error {
	return writePedidoDeAlmacenSCE(r, w)
}

func (r PedidoDeAlmacenSCE) Schema() string {
	return "{\"fields\":[{\"name\":\"destinatario\",\"type\":{\"fields\":[{\"name\":\"idDestinatario\",\"type\":\"string\"},{\"name\":\"datosPersonales\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idinternocliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"eMail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"tipo\",\"type\":[\"null\",\"string\"]},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDeTelefonos\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"contacto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoDeDocumento\",\"type\":[\"null\",{\"name\":\"TipoDeDocumento\",\"symbols\":[\"undefined\",\"DNI\",\"CUIT\",\"CUIL\"],\"type\":\"enum\"}]}],\"name\":\"DatosPersonales\",\"type\":\"record\"}},{\"name\":\"direccion\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"referenciadedomicilio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"provincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]}],\"name\":\"Direccion\",\"type\":\"record\"}},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"contacto\",\"type\":[\"null\",\"string\"]}],\"name\":\"Destinatario\",\"type\":\"record\"}},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"numeroOrdenExterna\",\"type\":\"string\"},{\"name\":\"almacencliente\",\"type\":\"string\"},{\"name\":\"fechaPedido\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"fechaEntrega\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"franjahoraria\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechadeexpedicionSolicitada\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"cambioLoteDirigido\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"valorDeclarado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"facturaLegal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechadeFacturacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"preciovalorFC\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ordenDeCompra\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tieneGestionCobranza\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cot\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"notas\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estatusOTdeTraza\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"marketPlace\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"datosadicionales\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.ListaDePropiedades\"]},{\"name\":\"listaDetalles\",\"type\":{\"fields\":[{\"name\":\"detallePedido\",\"type\":{\"items\":{\"fields\":[{\"name\":\"articulo\",\"type\":{\"fields\":[{\"name\":\"contratoWarehouse\",\"type\":\"string\"},{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"cantidad\",\"type\":\"double\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"numeropedido\",\"type\":[\"null\",\"string\"]},{\"name\":\"zonaConsumo\",\"type\":[\"null\",\"string\"]},{\"name\":\"unidadmedida\",\"type\":\"string\"},{\"default\":null,\"name\":\"datosadicionales\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.ListaDePropiedades\"]},{\"name\":\"lote\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaDeVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",\"string\"]},{\"name\":\"estado\",\"type\":[\"null\",\"string\"]}],\"name\":\"LoteArticulo\",\"type\":\"record\"}}],\"name\":\"Articulo\",\"type\":\"record\"}},{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]},{\"name\":\"numerodelinea\",\"type\":[\"null\",\"string\"]},{\"name\":\"tipoacondicionamientoescundario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"admitepickingparcial\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tiposDeAcondicionamientoSecundario\",\"type\":[\"null\",{\"fields\":[{\"name\":\"tiposDeAcondicionamientoSecundario\",\"type\":{\"items\":{\"fields\":[{\"name\":\"tipoDeAcondi\",\"type\":\"int\"}],\"name\":\"TiposDeAcondicionamientoSecundario\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDeTiposDeAcondicionamientoSecundario\",\"type\":\"record\"}]}],\"name\":\"DetallePedido\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDeDetalleDePedido\",\"type\":\"record\"}},{\"name\":\"distribuidor\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"etiqueta\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numerodeenvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"datosadicionales\",\"type\":[\"null\",\"string\"]}],\"name\":\"Distribuidor\",\"type\":\"record\"}},{\"default\":null,\"name\":\"linkImpresionRemito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tiposDeAcondicionamientoSecundario\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.ListaDeTiposDeAcondicionamientoSecundario\"]}],\"name\":\"Andreani.Wapv2.Events.Record.PedidoDeAlmacenSCE\",\"type\":\"record\"}"
}

func (r PedidoDeAlmacenSCE) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.PedidoDeAlmacenSCE"
}

func (_ PedidoDeAlmacenSCE) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ PedidoDeAlmacenSCE) SetInt(v int32)       { panic("Unsupported operation") }
func (_ PedidoDeAlmacenSCE) SetLong(v int64)      { panic("Unsupported operation") }
func (_ PedidoDeAlmacenSCE) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ PedidoDeAlmacenSCE) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ PedidoDeAlmacenSCE) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ PedidoDeAlmacenSCE) SetString(v string)   { panic("Unsupported operation") }
func (_ PedidoDeAlmacenSCE) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *PedidoDeAlmacenSCE) Get(i int) types.Field {
	switch i {
	case 0:
		r.Destinatario = NewDestinatario()

		w := types.Record{Target: &r.Destinatario}

		return w

	case 1:
		w := types.String{Target: &r.Propietario}

		return w

	case 2:
		w := types.String{Target: &r.NumeroOrdenExterna}

		return w

	case 3:
		w := types.String{Target: &r.Almacencliente}

		return w

	case 4:
		w := types.Long{Target: &r.FechaPedido}

		return w

	case 5:
		r.FechaEntrega = NewUnionNullLong()

		return r.FechaEntrega
	case 6:
		r.Franjahoraria = NewUnionNullString()

		return r.Franjahoraria
	case 7:
		r.FechadeexpedicionSolicitada = NewUnionNullLong()

		return r.FechadeexpedicionSolicitada
	case 8:
		r.CambioLoteDirigido = NewUnionNullString()

		return r.CambioLoteDirigido
	case 9:
		r.ValorDeclarado = NewUnionNullString()

		return r.ValorDeclarado
	case 10:
		r.Remito = NewUnionNullString()

		return r.Remito
	case 11:
		r.FacturaLegal = NewUnionNullString()

		return r.FacturaLegal
	case 12:
		r.FechadeFacturacion = NewUnionNullLong()

		return r.FechadeFacturacion
	case 13:
		r.PreciovalorFC = NewUnionNullString()

		return r.PreciovalorFC
	case 14:
		r.OrdenDeCompra = NewUnionNullString()

		return r.OrdenDeCompra
	case 15:
		r.TieneGestionCobranza = NewUnionNullString()

		return r.TieneGestionCobranza
	case 16:
		r.Cot = NewUnionNullString()

		return r.Cot
	case 17:
		r.Notas = NewUnionNullString()

		return r.Notas
	case 18:
		r.EstatusOTdeTraza = NewUnionNullString()

		return r.EstatusOTdeTraza
	case 19:
		r.MarketPlace = NewUnionNullString()

		return r.MarketPlace
	case 20:
		r.Datosadicionales = NewUnionNullListaDePropiedades()

		return r.Datosadicionales
	case 21:
		r.ListaDetalles = NewListaDeDetalleDePedido()

		w := types.Record{Target: &r.ListaDetalles}

		return w

	case 22:
		r.Distribuidor = NewDistribuidor()

		w := types.Record{Target: &r.Distribuidor}

		return w

	case 23:
		r.LinkImpresionRemito = NewUnionNullString()

		return r.LinkImpresionRemito
	case 24:
		r.TiposDeAcondicionamientoSecundario = NewUnionNullListaDeTiposDeAcondicionamientoSecundario()

		return r.TiposDeAcondicionamientoSecundario
	}
	panic("Unknown field index")
}

func (r *PedidoDeAlmacenSCE) SetDefault(i int) {
	switch i {
	case 5:
		r.FechaEntrega = nil
		return
	case 6:
		r.Franjahoraria = nil
		return
	case 7:
		r.FechadeexpedicionSolicitada = nil
		return
	case 8:
		r.CambioLoteDirigido = nil
		return
	case 9:
		r.ValorDeclarado = nil
		return
	case 10:
		r.Remito = nil
		return
	case 11:
		r.FacturaLegal = nil
		return
	case 12:
		r.FechadeFacturacion = nil
		return
	case 13:
		r.PreciovalorFC = nil
		return
	case 14:
		r.OrdenDeCompra = nil
		return
	case 15:
		r.TieneGestionCobranza = nil
		return
	case 16:
		r.Cot = nil
		return
	case 17:
		r.Notas = nil
		return
	case 18:
		r.EstatusOTdeTraza = nil
		return
	case 19:
		r.MarketPlace = nil
		return
	case 20:
		r.Datosadicionales = nil
		return
	case 23:
		r.LinkImpresionRemito = nil
		return
	case 24:
		r.TiposDeAcondicionamientoSecundario = nil
		return
	}
	panic("Unknown field index")
}

func (r *PedidoDeAlmacenSCE) NullField(i int) {
	switch i {
	case 5:
		r.FechaEntrega = nil
		return
	case 6:
		r.Franjahoraria = nil
		return
	case 7:
		r.FechadeexpedicionSolicitada = nil
		return
	case 8:
		r.CambioLoteDirigido = nil
		return
	case 9:
		r.ValorDeclarado = nil
		return
	case 10:
		r.Remito = nil
		return
	case 11:
		r.FacturaLegal = nil
		return
	case 12:
		r.FechadeFacturacion = nil
		return
	case 13:
		r.PreciovalorFC = nil
		return
	case 14:
		r.OrdenDeCompra = nil
		return
	case 15:
		r.TieneGestionCobranza = nil
		return
	case 16:
		r.Cot = nil
		return
	case 17:
		r.Notas = nil
		return
	case 18:
		r.EstatusOTdeTraza = nil
		return
	case 19:
		r.MarketPlace = nil
		return
	case 20:
		r.Datosadicionales = nil
		return
	case 23:
		r.LinkImpresionRemito = nil
		return
	case 24:
		r.TiposDeAcondicionamientoSecundario = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ PedidoDeAlmacenSCE) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ PedidoDeAlmacenSCE) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ PedidoDeAlmacenSCE) HintSize(int)                     { panic("Unsupported operation") }
func (_ PedidoDeAlmacenSCE) Finalize()                        {}

func (_ PedidoDeAlmacenSCE) AvroCRC64Fingerprint() []byte {
	return []byte(PedidoDeAlmacenSCEAvroCRC64Fingerprint)
}

func (r PedidoDeAlmacenSCE) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["destinatario"], err = json.Marshal(r.Destinatario)
	if err != nil {
		return nil, err
	}
	output["propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["numeroOrdenExterna"], err = json.Marshal(r.NumeroOrdenExterna)
	if err != nil {
		return nil, err
	}
	output["almacencliente"], err = json.Marshal(r.Almacencliente)
	if err != nil {
		return nil, err
	}
	output["fechaPedido"], err = json.Marshal(r.FechaPedido)
	if err != nil {
		return nil, err
	}
	output["fechaEntrega"], err = json.Marshal(r.FechaEntrega)
	if err != nil {
		return nil, err
	}
	output["franjahoraria"], err = json.Marshal(r.Franjahoraria)
	if err != nil {
		return nil, err
	}
	output["fechadeexpedicionSolicitada"], err = json.Marshal(r.FechadeexpedicionSolicitada)
	if err != nil {
		return nil, err
	}
	output["cambioLoteDirigido"], err = json.Marshal(r.CambioLoteDirigido)
	if err != nil {
		return nil, err
	}
	output["valorDeclarado"], err = json.Marshal(r.ValorDeclarado)
	if err != nil {
		return nil, err
	}
	output["remito"], err = json.Marshal(r.Remito)
	if err != nil {
		return nil, err
	}
	output["facturaLegal"], err = json.Marshal(r.FacturaLegal)
	if err != nil {
		return nil, err
	}
	output["fechadeFacturacion"], err = json.Marshal(r.FechadeFacturacion)
	if err != nil {
		return nil, err
	}
	output["preciovalorFC"], err = json.Marshal(r.PreciovalorFC)
	if err != nil {
		return nil, err
	}
	output["ordenDeCompra"], err = json.Marshal(r.OrdenDeCompra)
	if err != nil {
		return nil, err
	}
	output["tieneGestionCobranza"], err = json.Marshal(r.TieneGestionCobranza)
	if err != nil {
		return nil, err
	}
	output["cot"], err = json.Marshal(r.Cot)
	if err != nil {
		return nil, err
	}
	output["notas"], err = json.Marshal(r.Notas)
	if err != nil {
		return nil, err
	}
	output["estatusOTdeTraza"], err = json.Marshal(r.EstatusOTdeTraza)
	if err != nil {
		return nil, err
	}
	output["marketPlace"], err = json.Marshal(r.MarketPlace)
	if err != nil {
		return nil, err
	}
	output["datosadicionales"], err = json.Marshal(r.Datosadicionales)
	if err != nil {
		return nil, err
	}
	output["listaDetalles"], err = json.Marshal(r.ListaDetalles)
	if err != nil {
		return nil, err
	}
	output["distribuidor"], err = json.Marshal(r.Distribuidor)
	if err != nil {
		return nil, err
	}
	output["linkImpresionRemito"], err = json.Marshal(r.LinkImpresionRemito)
	if err != nil {
		return nil, err
	}
	output["tiposDeAcondicionamientoSecundario"], err = json.Marshal(r.TiposDeAcondicionamientoSecundario)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *PedidoDeAlmacenSCE) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
		return fmt.Errorf("no value specified for destinatario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["propietario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Propietario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for propietario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroOrdenExterna"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroOrdenExterna); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroOrdenExterna")
	}
	val = func() json.RawMessage {
		if v, ok := fields["almacencliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Almacencliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for almacencliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaPedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaPedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fechaPedido")
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaEntrega"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaEntrega); err != nil {
			return err
		}
	} else {
		r.FechaEntrega = NewUnionNullLong()

		r.FechaEntrega = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["franjahoraria"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Franjahoraria); err != nil {
			return err
		}
	} else {
		r.Franjahoraria = NewUnionNullString()

		r.Franjahoraria = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechadeexpedicionSolicitada"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechadeexpedicionSolicitada); err != nil {
			return err
		}
	} else {
		r.FechadeexpedicionSolicitada = NewUnionNullLong()

		r.FechadeexpedicionSolicitada = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["cambioLoteDirigido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CambioLoteDirigido); err != nil {
			return err
		}
	} else {
		r.CambioLoteDirigido = NewUnionNullString()

		r.CambioLoteDirigido = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["valorDeclarado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ValorDeclarado); err != nil {
			return err
		}
	} else {
		r.ValorDeclarado = NewUnionNullString()

		r.ValorDeclarado = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["remito"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Remito); err != nil {
			return err
		}
	} else {
		r.Remito = NewUnionNullString()

		r.Remito = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["facturaLegal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FacturaLegal); err != nil {
			return err
		}
	} else {
		r.FacturaLegal = NewUnionNullString()

		r.FacturaLegal = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechadeFacturacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechadeFacturacion); err != nil {
			return err
		}
	} else {
		r.FechadeFacturacion = NewUnionNullLong()

		r.FechadeFacturacion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["preciovalorFC"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PreciovalorFC); err != nil {
			return err
		}
	} else {
		r.PreciovalorFC = NewUnionNullString()

		r.PreciovalorFC = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ordenDeCompra"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenDeCompra); err != nil {
			return err
		}
	} else {
		r.OrdenDeCompra = NewUnionNullString()

		r.OrdenDeCompra = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["tieneGestionCobranza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TieneGestionCobranza); err != nil {
			return err
		}
	} else {
		r.TieneGestionCobranza = NewUnionNullString()

		r.TieneGestionCobranza = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["cot"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cot); err != nil {
			return err
		}
	} else {
		r.Cot = NewUnionNullString()

		r.Cot = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["notas"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Notas); err != nil {
			return err
		}
	} else {
		r.Notas = NewUnionNullString()

		r.Notas = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["estatusOTdeTraza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstatusOTdeTraza); err != nil {
			return err
		}
	} else {
		r.EstatusOTdeTraza = NewUnionNullString()

		r.EstatusOTdeTraza = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["marketPlace"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MarketPlace); err != nil {
			return err
		}
	} else {
		r.MarketPlace = NewUnionNullString()

		r.MarketPlace = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["datosadicionales"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Datosadicionales); err != nil {
			return err
		}
	} else {
		r.Datosadicionales = NewUnionNullListaDePropiedades()

		r.Datosadicionales = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["listaDetalles"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ListaDetalles); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for listaDetalles")
	}
	val = func() json.RawMessage {
		if v, ok := fields["distribuidor"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Distribuidor); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for distribuidor")
	}
	val = func() json.RawMessage {
		if v, ok := fields["linkImpresionRemito"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LinkImpresionRemito); err != nil {
			return err
		}
	} else {
		r.LinkImpresionRemito = NewUnionNullString()

		r.LinkImpresionRemito = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["tiposDeAcondicionamientoSecundario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TiposDeAcondicionamientoSecundario); err != nil {
			return err
		}
	} else {
		r.TiposDeAcondicionamientoSecundario = NewUnionNullListaDeTiposDeAcondicionamientoSecundario()

		r.TiposDeAcondicionamientoSecundario = nil
	}
	return nil
}
