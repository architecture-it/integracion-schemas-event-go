// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     OrdenDeEnvioSolicitada.avsc
 */
package AltaOrdenEnvioEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type DetalleDeOrdenDeEnvio struct {
	NumeroDeEnvio string `json:"numeroDeEnvio"`

	Remitente *UnionNullDatosPersonales `json:"remitente"`

	Destinatario DatosPersonales `json:"destinatario"`

	DestinatarioAlternativo *UnionNullArrayDatosPersonales `json:"destinatarioAlternativo"`

	DestinoInformado LugarDeRetiroOEntrega `json:"destinoInformado"`

	Origen *UnionNullLugarDeRetiroOEntrega `json:"origen"`

	Contrato string `json:"contrato"`

	TipoDeServicio *UnionNullString `json:"tipoDeServicio"`

	NumeroDeRemito *UnionNullString `json:"numeroDeRemito"`

	RemitosComplementarios *UnionNullArrayString `json:"remitosComplementarios"`

	CentroDeCostosDelCliente *UnionNullString `json:"centroDeCostosDelCliente"`

	CategoriaDeProducto *UnionNullString `json:"categoriaDeProducto"`

	CategoriaDeFacturacion *UnionNullString `json:"categoriaDeFacturacion"`

	ValorACobrar *UnionNullFloat `json:"valorACobrar"`

	FechaPactadaDeEntrega *UnionNullFechaPactadaDeEntrega `json:"fechaPactadaDeEntrega"`

	ProductoAEntregar *UnionNullString `json:"productoAEntregar"`

	ProductoARetirar *UnionNullString `json:"productoARetirar"`

	PagoDestino *UnionNullTipoPagoDestino `json:"pagoDestino"`

	SucursalDeDistribucion *UnionNullDatosSucursal `json:"sucursalDeDistribucion"`

	SucursalCliente *UnionNullString `json:"sucursalCliente"`

	ListaPaquetes *UnionNullListaPaquetes `json:"listaPaquetes"`

	CodigoVerificadorDeEntrega *UnionNullString `json:"codigoVerificadorDeEntrega"`

	CantidadDeBultos *UnionNullInt `json:"cantidadDeBultos"`

	AgrupadorDeBulto *UnionNullString `json:"agrupadorDeBulto"`

	IdPedido *UnionNullString `json:"idPedido"`

	PagoPendienteEnMostrador *UnionNullBool `json:"pagoPendienteEnMostrador"`
}

const DetalleDeOrdenDeEnvioAvroCRC64Fingerprint = "i\xee\xe0\x91\xeb\xb3\x1a?"

func NewDetalleDeOrdenDeEnvio() DetalleDeOrdenDeEnvio {
	r := DetalleDeOrdenDeEnvio{}
	r.Remitente = nil
	r.Destinatario = NewDatosPersonales()

	r.DestinatarioAlternativo = nil
	r.DestinoInformado = NewLugarDeRetiroOEntrega()

	r.Origen = nil
	r.TipoDeServicio = nil
	r.NumeroDeRemito = nil
	r.RemitosComplementarios = nil
	r.CentroDeCostosDelCliente = nil
	r.CategoriaDeProducto = nil
	r.CategoriaDeFacturacion = nil
	r.ValorACobrar = nil
	r.FechaPactadaDeEntrega = nil
	r.ProductoAEntregar = nil
	r.ProductoARetirar = nil
	r.PagoDestino = nil
	r.SucursalDeDistribucion = nil
	r.SucursalCliente = nil
	r.ListaPaquetes = nil
	r.CodigoVerificadorDeEntrega = nil
	r.CantidadDeBultos = nil
	r.AgrupadorDeBulto = nil
	r.IdPedido = nil
	return r
}

func DeserializeDetalleDeOrdenDeEnvio(r io.Reader) (DetalleDeOrdenDeEnvio, error) {
	t := NewDetalleDeOrdenDeEnvio()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDetalleDeOrdenDeEnvioFromSchema(r io.Reader, schema string) (DetalleDeOrdenDeEnvio, error) {
	t := NewDetalleDeOrdenDeEnvio()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDetalleDeOrdenDeEnvio(r DetalleDeOrdenDeEnvio, w io.Writer) error {
	var err error
	err = vm.WriteString(r.NumeroDeEnvio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDatosPersonales(r.Remitente, w)
	if err != nil {
		return err
	}
	err = writeDatosPersonales(r.Destinatario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayDatosPersonales(r.DestinatarioAlternativo, w)
	if err != nil {
		return err
	}
	err = writeLugarDeRetiroOEntrega(r.DestinoInformado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLugarDeRetiroOEntrega(r.Origen, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Contrato, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoDeServicio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeRemito, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayString(r.RemitosComplementarios, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CentroDeCostosDelCliente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CategoriaDeProducto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CategoriaDeFacturacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullFloat(r.ValorACobrar, w)
	if err != nil {
		return err
	}
	err = writeUnionNullFechaPactadaDeEntrega(r.FechaPactadaDeEntrega, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ProductoAEntregar, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ProductoARetirar, w)
	if err != nil {
		return err
	}
	err = writeUnionNullTipoPagoDestino(r.PagoDestino, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDatosSucursal(r.SucursalDeDistribucion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.SucursalCliente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullListaPaquetes(r.ListaPaquetes, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoVerificadorDeEntrega, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.CantidadDeBultos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.AgrupadorDeBulto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.IdPedido, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.PagoPendienteEnMostrador, w)
	if err != nil {
		return err
	}
	return err
}

func (r DetalleDeOrdenDeEnvio) Serialize(w io.Writer) error {
	return writeDetalleDeOrdenDeEnvio(r, w)
}

func (r DetalleDeOrdenDeEnvio) Schema() string {
	return "{\"fields\":[{\"name\":\"numeroDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"remitente\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idInternoDelCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"eMail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"tipo\",\"type\":\"string\"},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDeTelefonos\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"agrupador\",\"type\":[\"null\",\"string\"]},{\"name\":\"tipoDeDocumento\",\"type\":{\"name\":\"TipoDeDocumento\",\"symbols\":[\"undefined\",\"DNI\",\"CUIT\",\"CUIL\"],\"type\":\"enum\"}}],\"name\":\"DatosPersonales\",\"type\":\"record\"}]},{\"name\":\"destinatario\",\"type\":\"Andreani.AltaOrdenEnvio.Events.Common.DatosPersonales\"},{\"default\":null,\"name\":\"destinatarioAlternativo\",\"type\":[\"null\",{\"items\":\"Andreani.AltaOrdenEnvio.Events.Common.DatosPersonales\",\"type\":\"array\"}]},{\"name\":\"destinoInformado\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"postal\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"abreviaturaProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoDeDireccion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoCiudad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefono\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"segun ISO -3166-2 (https://es.wikipedia.org/wiki/ISO_3166-2:AR)\",\"name\":\"codigoISOProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"segun ISO 3166-1 alpha-2 (https://es.wikipedia.org/wiki/ISO_3166-1#Tabla_de_c.C3.B3digos_alfa-2_asignados_o_reservados)\",\"name\":\"codigoISOPais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"casillaDeCorreo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"region\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]}],\"name\":\"Direccion\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"datosSucursal\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"ID\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nomenclatura\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"direccion\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDePropiedades\"]},{\"default\":null,\"name\":\"casillaDeCorreo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"region\",\"type\":[\"null\",\"string\"]}],\"name\":\"DireccionPostal\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"datosAdicionales\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDePropiedades\"]},{\"default\":null,\"name\":\"listaDeTelefonos\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDeTelefonos\"]}],\"name\":\"Sucursal\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"GeoCoordenada\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"latitud\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"longitud\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"elevacion\",\"type\":[\"null\",\"double\"]}],\"name\":\"GeoCoordenada\",\"type\":\"record\"}]}],\"name\":\"LugarDeRetiroOEntrega\",\"type\":\"record\"}},{\"default\":null,\"name\":\"origen\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.LugarDeRetiroOEntrega\"]},{\"name\":\"contrato\",\"type\":\"string\"},{\"default\":null,\"name\":\"tipoDeServicio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeRemito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remitosComplementarios\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"centroDeCostosDelCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"categoriaDeProducto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"categoriaDeFacturacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"valorACobrar\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"fechaPactadaDeEntrega\",\"type\":[\"null\",{\"fields\":[{\"name\":\"fecha\",\"type\":[\"null\",\"string\"]},{\"name\":\"horaDesde\",\"type\":[\"null\",\"string\"]},{\"name\":\"horaHasta\",\"type\":[\"null\",\"string\"]}],\"name\":\"FechaPactadaDeEntrega\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"productoAEntregar\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"productoARetirar\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pagoDestino\",\"type\":[\"null\",{\"name\":\"TipoPagoDestino\",\"symbols\":[\"undefined\",\"P\",\"D\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"sucursalDeDistribucion\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"sucursalCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"listaPaquetes\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaPaquetes\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"pesoEnKg\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"altoEnCm\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"anchoEnCm\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"largoEnCm\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"referenciasDelCliente\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDePropiedades\"]},{\"default\":null,\"name\":\"volumenEnCm3\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"valorDeclaradoSinImpuesto\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"valorDeclaradoConImpuesto\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"numeroDeBulto\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"valorDeclarado\",\"type\":[\"null\",\"double\"]}],\"name\":\"DetalleDePaquete\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaPaquetes\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"codigoVerificadorDeEntrega\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cantidadDeBultos\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"agrupadorDeBulto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idPedido\",\"type\":[\"null\",\"string\"]},{\"name\":\"pagoPendienteEnMostrador\",\"type\":[\"null\",\"boolean\"]}],\"name\":\"Andreani.AltaOrdenEnvio.Events.Common.DetalleDeOrdenDeEnvio\",\"type\":\"record\"}"
}

func (r DetalleDeOrdenDeEnvio) SchemaName() string {
	return "Andreani.AltaOrdenEnvio.Events.Common.DetalleDeOrdenDeEnvio"
}

func (_ DetalleDeOrdenDeEnvio) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeEnvio) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeEnvio) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeEnvio) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeEnvio) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeEnvio) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeEnvio) SetString(v string)   { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeEnvio) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DetalleDeOrdenDeEnvio) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.NumeroDeEnvio}

		return w

	case 1:
		r.Remitente = NewUnionNullDatosPersonales()

		return r.Remitente
	case 2:
		r.Destinatario = NewDatosPersonales()

		w := types.Record{Target: &r.Destinatario}

		return w

	case 3:
		r.DestinatarioAlternativo = NewUnionNullArrayDatosPersonales()

		return r.DestinatarioAlternativo
	case 4:
		r.DestinoInformado = NewLugarDeRetiroOEntrega()

		w := types.Record{Target: &r.DestinoInformado}

		return w

	case 5:
		r.Origen = NewUnionNullLugarDeRetiroOEntrega()

		return r.Origen
	case 6:
		w := types.String{Target: &r.Contrato}

		return w

	case 7:
		r.TipoDeServicio = NewUnionNullString()

		return r.TipoDeServicio
	case 8:
		r.NumeroDeRemito = NewUnionNullString()

		return r.NumeroDeRemito
	case 9:
		r.RemitosComplementarios = NewUnionNullArrayString()

		return r.RemitosComplementarios
	case 10:
		r.CentroDeCostosDelCliente = NewUnionNullString()

		return r.CentroDeCostosDelCliente
	case 11:
		r.CategoriaDeProducto = NewUnionNullString()

		return r.CategoriaDeProducto
	case 12:
		r.CategoriaDeFacturacion = NewUnionNullString()

		return r.CategoriaDeFacturacion
	case 13:
		r.ValorACobrar = NewUnionNullFloat()

		return r.ValorACobrar
	case 14:
		r.FechaPactadaDeEntrega = NewUnionNullFechaPactadaDeEntrega()

		return r.FechaPactadaDeEntrega
	case 15:
		r.ProductoAEntregar = NewUnionNullString()

		return r.ProductoAEntregar
	case 16:
		r.ProductoARetirar = NewUnionNullString()

		return r.ProductoARetirar
	case 17:
		r.PagoDestino = NewUnionNullTipoPagoDestino()

		return r.PagoDestino
	case 18:
		r.SucursalDeDistribucion = NewUnionNullDatosSucursal()

		return r.SucursalDeDistribucion
	case 19:
		r.SucursalCliente = NewUnionNullString()

		return r.SucursalCliente
	case 20:
		r.ListaPaquetes = NewUnionNullListaPaquetes()

		return r.ListaPaquetes
	case 21:
		r.CodigoVerificadorDeEntrega = NewUnionNullString()

		return r.CodigoVerificadorDeEntrega
	case 22:
		r.CantidadDeBultos = NewUnionNullInt()

		return r.CantidadDeBultos
	case 23:
		r.AgrupadorDeBulto = NewUnionNullString()

		return r.AgrupadorDeBulto
	case 24:
		r.IdPedido = NewUnionNullString()

		return r.IdPedido
	case 25:
		r.PagoPendienteEnMostrador = NewUnionNullBool()

		return r.PagoPendienteEnMostrador
	}
	panic("Unknown field index")
}

func (r *DetalleDeOrdenDeEnvio) SetDefault(i int) {
	switch i {
	case 1:
		r.Remitente = nil
		return
	case 3:
		r.DestinatarioAlternativo = nil
		return
	case 5:
		r.Origen = nil
		return
	case 7:
		r.TipoDeServicio = nil
		return
	case 8:
		r.NumeroDeRemito = nil
		return
	case 9:
		r.RemitosComplementarios = nil
		return
	case 10:
		r.CentroDeCostosDelCliente = nil
		return
	case 11:
		r.CategoriaDeProducto = nil
		return
	case 12:
		r.CategoriaDeFacturacion = nil
		return
	case 13:
		r.ValorACobrar = nil
		return
	case 14:
		r.FechaPactadaDeEntrega = nil
		return
	case 15:
		r.ProductoAEntregar = nil
		return
	case 16:
		r.ProductoARetirar = nil
		return
	case 17:
		r.PagoDestino = nil
		return
	case 18:
		r.SucursalDeDistribucion = nil
		return
	case 19:
		r.SucursalCliente = nil
		return
	case 20:
		r.ListaPaquetes = nil
		return
	case 21:
		r.CodigoVerificadorDeEntrega = nil
		return
	case 22:
		r.CantidadDeBultos = nil
		return
	case 23:
		r.AgrupadorDeBulto = nil
		return
	case 24:
		r.IdPedido = nil
		return
	}
	panic("Unknown field index")
}

func (r *DetalleDeOrdenDeEnvio) NullField(i int) {
	switch i {
	case 1:
		r.Remitente = nil
		return
	case 3:
		r.DestinatarioAlternativo = nil
		return
	case 5:
		r.Origen = nil
		return
	case 7:
		r.TipoDeServicio = nil
		return
	case 8:
		r.NumeroDeRemito = nil
		return
	case 9:
		r.RemitosComplementarios = nil
		return
	case 10:
		r.CentroDeCostosDelCliente = nil
		return
	case 11:
		r.CategoriaDeProducto = nil
		return
	case 12:
		r.CategoriaDeFacturacion = nil
		return
	case 13:
		r.ValorACobrar = nil
		return
	case 14:
		r.FechaPactadaDeEntrega = nil
		return
	case 15:
		r.ProductoAEntregar = nil
		return
	case 16:
		r.ProductoARetirar = nil
		return
	case 17:
		r.PagoDestino = nil
		return
	case 18:
		r.SucursalDeDistribucion = nil
		return
	case 19:
		r.SucursalCliente = nil
		return
	case 20:
		r.ListaPaquetes = nil
		return
	case 21:
		r.CodigoVerificadorDeEntrega = nil
		return
	case 22:
		r.CantidadDeBultos = nil
		return
	case 23:
		r.AgrupadorDeBulto = nil
		return
	case 24:
		r.IdPedido = nil
		return
	case 25:
		r.PagoPendienteEnMostrador = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ DetalleDeOrdenDeEnvio) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeEnvio) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeEnvio) HintSize(int)                     { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeEnvio) Finalize()                        {}

func (_ DetalleDeOrdenDeEnvio) AvroCRC64Fingerprint() []byte {
	return []byte(DetalleDeOrdenDeEnvioAvroCRC64Fingerprint)
}

func (r DetalleDeOrdenDeEnvio) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["numeroDeEnvio"], err = json.Marshal(r.NumeroDeEnvio)
	if err != nil {
		return nil, err
	}
	output["remitente"], err = json.Marshal(r.Remitente)
	if err != nil {
		return nil, err
	}
	output["destinatario"], err = json.Marshal(r.Destinatario)
	if err != nil {
		return nil, err
	}
	output["destinatarioAlternativo"], err = json.Marshal(r.DestinatarioAlternativo)
	if err != nil {
		return nil, err
	}
	output["destinoInformado"], err = json.Marshal(r.DestinoInformado)
	if err != nil {
		return nil, err
	}
	output["origen"], err = json.Marshal(r.Origen)
	if err != nil {
		return nil, err
	}
	output["contrato"], err = json.Marshal(r.Contrato)
	if err != nil {
		return nil, err
	}
	output["tipoDeServicio"], err = json.Marshal(r.TipoDeServicio)
	if err != nil {
		return nil, err
	}
	output["numeroDeRemito"], err = json.Marshal(r.NumeroDeRemito)
	if err != nil {
		return nil, err
	}
	output["remitosComplementarios"], err = json.Marshal(r.RemitosComplementarios)
	if err != nil {
		return nil, err
	}
	output["centroDeCostosDelCliente"], err = json.Marshal(r.CentroDeCostosDelCliente)
	if err != nil {
		return nil, err
	}
	output["categoriaDeProducto"], err = json.Marshal(r.CategoriaDeProducto)
	if err != nil {
		return nil, err
	}
	output["categoriaDeFacturacion"], err = json.Marshal(r.CategoriaDeFacturacion)
	if err != nil {
		return nil, err
	}
	output["valorACobrar"], err = json.Marshal(r.ValorACobrar)
	if err != nil {
		return nil, err
	}
	output["fechaPactadaDeEntrega"], err = json.Marshal(r.FechaPactadaDeEntrega)
	if err != nil {
		return nil, err
	}
	output["productoAEntregar"], err = json.Marshal(r.ProductoAEntregar)
	if err != nil {
		return nil, err
	}
	output["productoARetirar"], err = json.Marshal(r.ProductoARetirar)
	if err != nil {
		return nil, err
	}
	output["pagoDestino"], err = json.Marshal(r.PagoDestino)
	if err != nil {
		return nil, err
	}
	output["sucursalDeDistribucion"], err = json.Marshal(r.SucursalDeDistribucion)
	if err != nil {
		return nil, err
	}
	output["sucursalCliente"], err = json.Marshal(r.SucursalCliente)
	if err != nil {
		return nil, err
	}
	output["listaPaquetes"], err = json.Marshal(r.ListaPaquetes)
	if err != nil {
		return nil, err
	}
	output["codigoVerificadorDeEntrega"], err = json.Marshal(r.CodigoVerificadorDeEntrega)
	if err != nil {
		return nil, err
	}
	output["cantidadDeBultos"], err = json.Marshal(r.CantidadDeBultos)
	if err != nil {
		return nil, err
	}
	output["agrupadorDeBulto"], err = json.Marshal(r.AgrupadorDeBulto)
	if err != nil {
		return nil, err
	}
	output["idPedido"], err = json.Marshal(r.IdPedido)
	if err != nil {
		return nil, err
	}
	output["pagoPendienteEnMostrador"], err = json.Marshal(r.PagoPendienteEnMostrador)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DetalleDeOrdenDeEnvio) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["numeroDeEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeEnvio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroDeEnvio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["remitente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Remitente); err != nil {
			return err
		}
	} else {
		r.Remitente = NewUnionNullDatosPersonales()

		r.Remitente = nil
	}
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
		if v, ok := fields["destinatarioAlternativo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioAlternativo); err != nil {
			return err
		}
	} else {
		r.DestinatarioAlternativo = NewUnionNullArrayDatosPersonales()

		r.DestinatarioAlternativo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["destinoInformado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinoInformado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for destinoInformado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["origen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Origen); err != nil {
			return err
		}
	} else {
		r.Origen = NewUnionNullLugarDeRetiroOEntrega()

		r.Origen = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["contrato"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contrato); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for contrato")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoDeServicio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeServicio); err != nil {
			return err
		}
	} else {
		r.TipoDeServicio = NewUnionNullString()

		r.TipoDeServicio = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroDeRemito"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeRemito); err != nil {
			return err
		}
	} else {
		r.NumeroDeRemito = NewUnionNullString()

		r.NumeroDeRemito = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["remitosComplementarios"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RemitosComplementarios); err != nil {
			return err
		}
	} else {
		r.RemitosComplementarios = NewUnionNullArrayString()

		r.RemitosComplementarios = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["centroDeCostosDelCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CentroDeCostosDelCliente); err != nil {
			return err
		}
	} else {
		r.CentroDeCostosDelCliente = NewUnionNullString()

		r.CentroDeCostosDelCliente = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["categoriaDeProducto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CategoriaDeProducto); err != nil {
			return err
		}
	} else {
		r.CategoriaDeProducto = NewUnionNullString()

		r.CategoriaDeProducto = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["categoriaDeFacturacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CategoriaDeFacturacion); err != nil {
			return err
		}
	} else {
		r.CategoriaDeFacturacion = NewUnionNullString()

		r.CategoriaDeFacturacion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["valorACobrar"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ValorACobrar); err != nil {
			return err
		}
	} else {
		r.ValorACobrar = NewUnionNullFloat()

		r.ValorACobrar = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaPactadaDeEntrega"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaPactadaDeEntrega); err != nil {
			return err
		}
	} else {
		r.FechaPactadaDeEntrega = NewUnionNullFechaPactadaDeEntrega()

		r.FechaPactadaDeEntrega = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["productoAEntregar"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ProductoAEntregar); err != nil {
			return err
		}
	} else {
		r.ProductoAEntregar = NewUnionNullString()

		r.ProductoAEntregar = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["productoARetirar"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ProductoARetirar); err != nil {
			return err
		}
	} else {
		r.ProductoARetirar = NewUnionNullString()

		r.ProductoARetirar = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["pagoDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PagoDestino); err != nil {
			return err
		}
	} else {
		r.PagoDestino = NewUnionNullTipoPagoDestino()

		r.PagoDestino = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["sucursalDeDistribucion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SucursalDeDistribucion); err != nil {
			return err
		}
	} else {
		r.SucursalDeDistribucion = NewUnionNullDatosSucursal()

		r.SucursalDeDistribucion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["sucursalCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SucursalCliente); err != nil {
			return err
		}
	} else {
		r.SucursalCliente = NewUnionNullString()

		r.SucursalCliente = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["listaPaquetes"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ListaPaquetes); err != nil {
			return err
		}
	} else {
		r.ListaPaquetes = NewUnionNullListaPaquetes()

		r.ListaPaquetes = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoVerificadorDeEntrega"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoVerificadorDeEntrega); err != nil {
			return err
		}
	} else {
		r.CodigoVerificadorDeEntrega = NewUnionNullString()

		r.CodigoVerificadorDeEntrega = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["cantidadDeBultos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadDeBultos); err != nil {
			return err
		}
	} else {
		r.CantidadDeBultos = NewUnionNullInt()

		r.CantidadDeBultos = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["agrupadorDeBulto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AgrupadorDeBulto); err != nil {
			return err
		}
	} else {
		r.AgrupadorDeBulto = NewUnionNullString()

		r.AgrupadorDeBulto = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["idPedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdPedido); err != nil {
			return err
		}
	} else {
		r.IdPedido = NewUnionNullString()

		r.IdPedido = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["pagoPendienteEnMostrador"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PagoPendienteEnMostrador); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for pagoPendienteEnMostrador")
	}
	return nil
}
