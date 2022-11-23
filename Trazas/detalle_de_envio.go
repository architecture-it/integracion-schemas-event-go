// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TrazaAltaDeEnvio.avsc
 */
package TrazasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type DetalleDeEnvio struct {
	FechaDeRecibo *UnionNullString `json:"fechaDeRecibo"`

	InformaTarifa bool `json:"informaTarifa"`

	Lote *UnionNullString `json:"lote"`

	NumeroDeEnvio string `json:"numeroDeEnvio"`

	NumeroDeRecibo *UnionNullString `json:"numeroDeRecibo"`

	NumeroInformeDeRendicion *UnionNullString `json:"numeroInformeDeRendicion"`

	Peso *UnionNullFloat `json:"peso"`

	PesoAforado *UnionNullFloat `json:"pesoAforado"`

	ValorDeclarado *UnionNullFloat `json:"valorDeclarado"`

	ValorACobrar *UnionNullFloat `json:"valorACobrar"`

	Volumen *UnionNullFloat `json:"volumen"`

	Origen *UnionNullLugarDeRetiroEntrega `json:"origen"`

	SucursalOrigen *UnionNullLugarDeRetiroEntrega `json:"sucursalOrigen"`

	Destino LugarDeRetiroEntrega `json:"destino"`

	SucursalDestino *UnionNullLugarDeRetiroEntrega `json:"sucursalDestino"`

	CategoriaDeFacturacion *UnionNullString `json:"categoriaDeFacturacion"`

	CentroDeCostosDelCliente *UnionNullString `json:"centroDeCostosDelCliente"`

	Estado string `json:"estado"`

	CodigoDeContratoInterno string `json:"codigoDeContratoInterno"`

	FechaInformeDeRendicion *UnionNullString `json:"fechaInformeDeRendicion"`

	FechaAltaEnvio *UnionNullString `json:"fechaAltaEnvio"`

	OrdenDeCompraCliente *UnionNullString `json:"ordenDeCompraCliente"`

	Tarifa *UnionNullDouble `json:"tarifa"`

	Componentes []DatosPersonales `json:"componentes"`

	Destinatario *UnionNullDatosPersonales `json:"destinatario"`

	Remitente *UnionNullDatosPersonales `json:"remitente"`

	CantidadDeBultos *UnionNullInt `json:"cantidadDeBultos"`

	FechaProbableDeEntrega *UnionNullString `json:"fechaProbableDeEntrega"`

	Observaciones *UnionNullString `json:"observaciones"`

	NumeroDeRemito *UnionNullString `json:"numeroDeRemito"`

	NumeroDeSeguimientoDelCliente *UnionNullString `json:"numeroDeSeguimientoDelCliente"`

	DatosAdicionales *UnionNullMapString `json:"datosAdicionales"`

	FechaPactadaDeEntrega *UnionNullString `json:"fechaPactadaDeEntrega"`

	AltoEnCm *UnionNullDouble `json:"altoEnCm"`

	AnchoEnCm *UnionNullDouble `json:"anchoEnCm"`

	LargoEnCm *UnionNullDouble `json:"largoEnCm"`
}

const DetalleDeEnvioAvroCRC64Fingerprint = "\x0f笁\x19\xa2\xc4\x1c"

func NewDetalleDeEnvio() DetalleDeEnvio {
	r := DetalleDeEnvio{}
	r.FechaDeRecibo = nil
	r.Lote = nil
	r.NumeroDeRecibo = nil
	r.NumeroInformeDeRendicion = nil
	r.Peso = nil
	r.PesoAforado = nil
	r.ValorDeclarado = nil
	r.ValorACobrar = nil
	r.Volumen = nil
	r.Origen = nil
	r.SucursalOrigen = nil
	r.Destino = NewLugarDeRetiroEntrega()

	r.SucursalDestino = nil
	r.CategoriaDeFacturacion = nil
	r.CentroDeCostosDelCliente = nil
	r.FechaInformeDeRendicion = nil
	r.FechaAltaEnvio = nil
	r.OrdenDeCompraCliente = nil
	r.Tarifa = nil
	r.Componentes = make([]DatosPersonales, 0)

	r.Destinatario = nil
	r.Remitente = nil
	r.CantidadDeBultos = nil
	r.FechaProbableDeEntrega = nil
	r.Observaciones = nil
	r.NumeroDeRemito = nil
	r.NumeroDeSeguimientoDelCliente = nil
	r.DatosAdicionales = nil
	r.FechaPactadaDeEntrega = nil
	r.AltoEnCm = nil
	r.AnchoEnCm = nil
	r.LargoEnCm = nil
	return r
}

func DeserializeDetalleDeEnvio(r io.Reader) (DetalleDeEnvio, error) {
	t := NewDetalleDeEnvio()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDetalleDeEnvioFromSchema(r io.Reader, schema string) (DetalleDeEnvio, error) {
	t := NewDetalleDeEnvio()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDetalleDeEnvio(r DetalleDeEnvio, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.FechaDeRecibo, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.InformaTarifa, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Lote, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroDeEnvio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeRecibo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroInformeDeRendicion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullFloat(r.Peso, w)
	if err != nil {
		return err
	}
	err = writeUnionNullFloat(r.PesoAforado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullFloat(r.ValorDeclarado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullFloat(r.ValorACobrar, w)
	if err != nil {
		return err
	}
	err = writeUnionNullFloat(r.Volumen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLugarDeRetiroEntrega(r.Origen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLugarDeRetiroEntrega(r.SucursalOrigen, w)
	if err != nil {
		return err
	}
	err = writeLugarDeRetiroEntrega(r.Destino, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLugarDeRetiroEntrega(r.SucursalDestino, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CategoriaDeFacturacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CentroDeCostosDelCliente, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Estado, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoDeContratoInterno, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaInformeDeRendicion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaAltaEnvio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.OrdenDeCompraCliente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDouble(r.Tarifa, w)
	if err != nil {
		return err
	}
	err = writeArrayDatosPersonales(r.Componentes, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDatosPersonales(r.Destinatario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDatosPersonales(r.Remitente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.CantidadDeBultos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaProbableDeEntrega, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Observaciones, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeRemito, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeSeguimientoDelCliente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullMapString(r.DatosAdicionales, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaPactadaDeEntrega, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDouble(r.AltoEnCm, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDouble(r.AnchoEnCm, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDouble(r.LargoEnCm, w)
	if err != nil {
		return err
	}
	return err
}

func (r DetalleDeEnvio) Serialize(w io.Writer) error {
	return writeDetalleDeEnvio(r, w)
}

func (r DetalleDeEnvio) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"fechaDeRecibo\",\"type\":[\"null\",\"string\"]},{\"name\":\"informaTarifa\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"lote\",\"type\":[\"null\",\"string\"]},{\"name\":\"numeroDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"numeroDeRecibo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroInformeDeRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"peso\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"pesoAforado\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"valorDeclarado\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"valorACobrar\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"volumen\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"origen\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"datosSucursal\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"domicilio\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"abreviaturaProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoDeDireccion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoCiudad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefono\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"segun ISO -3166-2 (https://es.wikipedia.org/wiki/ISO_3166-2:AR)\",\"name\":\"codigoISOProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"segun ISO 3166-1 alpha-2 (https://es.wikipedia.org/wiki/ISO_3166-1#Tabla_de_c.C3.B3digos_alfa-2_asignados_o_reservados)\",\"name\":\"codigoISOPais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",{\"type\":\"map\",\"values\":\"string\"}]},{\"default\":null,\"name\":\"coordenadas\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"latitud\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"longitud\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"altura\",\"type\":[\"null\",\"double\"]}],\"name\":\"GeoReferencia\",\"type\":\"record\"}]}],\"name\":\"Direccion\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"esLocal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"region\",\"type\":[\"null\",\"string\"]}],\"name\":\"LugarDeRetiroEntrega\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"sucursalOrigen\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.LugarDeRetiroEntrega\"]},{\"name\":\"destino\",\"type\":\"Integracion.Esquemas.Referencias.LugarDeRetiroEntrega\"},{\"default\":null,\"name\":\"sucursalDestino\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.LugarDeRetiroEntrega\"]},{\"default\":null,\"name\":\"categoriaDeFacturacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"centroDeCostosDelCliente\",\"type\":[\"null\",\"string\"]},{\"name\":\"estado\",\"type\":\"string\"},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"fechaInformeDeRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaAltaEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ordenDeCompraCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tarifa\",\"type\":[\"null\",\"double\"]},{\"name\":\"componentes\",\"type\":{\"items\":{\"fields\":[{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idInternoDelCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"eMail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"tipo\",\"type\":{\"name\":\"TipoDeTelefono\",\"symbols\":[\"trabajo\",\"celular\",\"casa\",\"otro\"],\"type\":\"enum\"}},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"agrupador\",\"type\":[\"null\",\"string\"]},{\"name\":\"tipoDeDocumento\",\"type\":{\"name\":\"TipoDeDocumento\",\"symbols\":[\"undefined\",\"DNI\",\"CUIT\",\"CUIL\"],\"type\":\"enum\"}}],\"name\":\"DatosPersonales\",\"type\":\"record\"},\"type\":\"array\"}},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosPersonales\"]},{\"default\":null,\"name\":\"remitente\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosPersonales\"]},{\"default\":null,\"name\":\"cantidadDeBultos\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"fechaProbableDeEntrega\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"observaciones\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeRemito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeSeguimientoDelCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"datosAdicionales\",\"type\":[\"null\",{\"type\":\"map\",\"values\":\"string\"}]},{\"default\":null,\"name\":\"fechaPactadaDeEntrega\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"altoEnCm\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"anchoEnCm\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"largoEnCm\",\"type\":[\"null\",\"double\"]}],\"name\":\"Integracion.Esquemas.Referencias.DetalleDeEnvio\",\"type\":\"record\"}"
}

func (r DetalleDeEnvio) SchemaName() string {
	return "Integracion.Esquemas.Referencias.DetalleDeEnvio"
}

func (_ DetalleDeEnvio) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DetalleDeEnvio) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DetalleDeEnvio) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DetalleDeEnvio) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DetalleDeEnvio) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DetalleDeEnvio) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DetalleDeEnvio) SetString(v string)   { panic("Unsupported operation") }
func (_ DetalleDeEnvio) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DetalleDeEnvio) Get(i int) types.Field {
	switch i {
	case 0:
		r.FechaDeRecibo = NewUnionNullString()

		return r.FechaDeRecibo
	case 1:
		w := types.Boolean{Target: &r.InformaTarifa}

		return w

	case 2:
		r.Lote = NewUnionNullString()

		return r.Lote
	case 3:
		w := types.String{Target: &r.NumeroDeEnvio}

		return w

	case 4:
		r.NumeroDeRecibo = NewUnionNullString()

		return r.NumeroDeRecibo
	case 5:
		r.NumeroInformeDeRendicion = NewUnionNullString()

		return r.NumeroInformeDeRendicion
	case 6:
		r.Peso = NewUnionNullFloat()

		return r.Peso
	case 7:
		r.PesoAforado = NewUnionNullFloat()

		return r.PesoAforado
	case 8:
		r.ValorDeclarado = NewUnionNullFloat()

		return r.ValorDeclarado
	case 9:
		r.ValorACobrar = NewUnionNullFloat()

		return r.ValorACobrar
	case 10:
		r.Volumen = NewUnionNullFloat()

		return r.Volumen
	case 11:
		r.Origen = NewUnionNullLugarDeRetiroEntrega()

		return r.Origen
	case 12:
		r.SucursalOrigen = NewUnionNullLugarDeRetiroEntrega()

		return r.SucursalOrigen
	case 13:
		r.Destino = NewLugarDeRetiroEntrega()

		w := types.Record{Target: &r.Destino}

		return w

	case 14:
		r.SucursalDestino = NewUnionNullLugarDeRetiroEntrega()

		return r.SucursalDestino
	case 15:
		r.CategoriaDeFacturacion = NewUnionNullString()

		return r.CategoriaDeFacturacion
	case 16:
		r.CentroDeCostosDelCliente = NewUnionNullString()

		return r.CentroDeCostosDelCliente
	case 17:
		w := types.String{Target: &r.Estado}

		return w

	case 18:
		w := types.String{Target: &r.CodigoDeContratoInterno}

		return w

	case 19:
		r.FechaInformeDeRendicion = NewUnionNullString()

		return r.FechaInformeDeRendicion
	case 20:
		r.FechaAltaEnvio = NewUnionNullString()

		return r.FechaAltaEnvio
	case 21:
		r.OrdenDeCompraCliente = NewUnionNullString()

		return r.OrdenDeCompraCliente
	case 22:
		r.Tarifa = NewUnionNullDouble()

		return r.Tarifa
	case 23:
		r.Componentes = make([]DatosPersonales, 0)

		w := ArrayDatosPersonalesWrapper{Target: &r.Componentes}

		return w

	case 24:
		r.Destinatario = NewUnionNullDatosPersonales()

		return r.Destinatario
	case 25:
		r.Remitente = NewUnionNullDatosPersonales()

		return r.Remitente
	case 26:
		r.CantidadDeBultos = NewUnionNullInt()

		return r.CantidadDeBultos
	case 27:
		r.FechaProbableDeEntrega = NewUnionNullString()

		return r.FechaProbableDeEntrega
	case 28:
		r.Observaciones = NewUnionNullString()

		return r.Observaciones
	case 29:
		r.NumeroDeRemito = NewUnionNullString()

		return r.NumeroDeRemito
	case 30:
		r.NumeroDeSeguimientoDelCliente = NewUnionNullString()

		return r.NumeroDeSeguimientoDelCliente
	case 31:
		r.DatosAdicionales = NewUnionNullMapString()

		return r.DatosAdicionales
	case 32:
		r.FechaPactadaDeEntrega = NewUnionNullString()

		return r.FechaPactadaDeEntrega
	case 33:
		r.AltoEnCm = NewUnionNullDouble()

		return r.AltoEnCm
	case 34:
		r.AnchoEnCm = NewUnionNullDouble()

		return r.AnchoEnCm
	case 35:
		r.LargoEnCm = NewUnionNullDouble()

		return r.LargoEnCm
	}
	panic("Unknown field index")
}

func (r *DetalleDeEnvio) SetDefault(i int) {
	switch i {
	case 0:
		r.FechaDeRecibo = nil
		return
	case 2:
		r.Lote = nil
		return
	case 4:
		r.NumeroDeRecibo = nil
		return
	case 5:
		r.NumeroInformeDeRendicion = nil
		return
	case 6:
		r.Peso = nil
		return
	case 7:
		r.PesoAforado = nil
		return
	case 8:
		r.ValorDeclarado = nil
		return
	case 9:
		r.ValorACobrar = nil
		return
	case 10:
		r.Volumen = nil
		return
	case 11:
		r.Origen = nil
		return
	case 12:
		r.SucursalOrigen = nil
		return
	case 14:
		r.SucursalDestino = nil
		return
	case 15:
		r.CategoriaDeFacturacion = nil
		return
	case 16:
		r.CentroDeCostosDelCliente = nil
		return
	case 19:
		r.FechaInformeDeRendicion = nil
		return
	case 20:
		r.FechaAltaEnvio = nil
		return
	case 21:
		r.OrdenDeCompraCliente = nil
		return
	case 22:
		r.Tarifa = nil
		return
	case 24:
		r.Destinatario = nil
		return
	case 25:
		r.Remitente = nil
		return
	case 26:
		r.CantidadDeBultos = nil
		return
	case 27:
		r.FechaProbableDeEntrega = nil
		return
	case 28:
		r.Observaciones = nil
		return
	case 29:
		r.NumeroDeRemito = nil
		return
	case 30:
		r.NumeroDeSeguimientoDelCliente = nil
		return
	case 31:
		r.DatosAdicionales = nil
		return
	case 32:
		r.FechaPactadaDeEntrega = nil
		return
	case 33:
		r.AltoEnCm = nil
		return
	case 34:
		r.AnchoEnCm = nil
		return
	case 35:
		r.LargoEnCm = nil
		return
	}
	panic("Unknown field index")
}

func (r *DetalleDeEnvio) NullField(i int) {
	switch i {
	case 0:
		r.FechaDeRecibo = nil
		return
	case 2:
		r.Lote = nil
		return
	case 4:
		r.NumeroDeRecibo = nil
		return
	case 5:
		r.NumeroInformeDeRendicion = nil
		return
	case 6:
		r.Peso = nil
		return
	case 7:
		r.PesoAforado = nil
		return
	case 8:
		r.ValorDeclarado = nil
		return
	case 9:
		r.ValorACobrar = nil
		return
	case 10:
		r.Volumen = nil
		return
	case 11:
		r.Origen = nil
		return
	case 12:
		r.SucursalOrigen = nil
		return
	case 14:
		r.SucursalDestino = nil
		return
	case 15:
		r.CategoriaDeFacturacion = nil
		return
	case 16:
		r.CentroDeCostosDelCliente = nil
		return
	case 19:
		r.FechaInformeDeRendicion = nil
		return
	case 20:
		r.FechaAltaEnvio = nil
		return
	case 21:
		r.OrdenDeCompraCliente = nil
		return
	case 22:
		r.Tarifa = nil
		return
	case 24:
		r.Destinatario = nil
		return
	case 25:
		r.Remitente = nil
		return
	case 26:
		r.CantidadDeBultos = nil
		return
	case 27:
		r.FechaProbableDeEntrega = nil
		return
	case 28:
		r.Observaciones = nil
		return
	case 29:
		r.NumeroDeRemito = nil
		return
	case 30:
		r.NumeroDeSeguimientoDelCliente = nil
		return
	case 31:
		r.DatosAdicionales = nil
		return
	case 32:
		r.FechaPactadaDeEntrega = nil
		return
	case 33:
		r.AltoEnCm = nil
		return
	case 34:
		r.AnchoEnCm = nil
		return
	case 35:
		r.LargoEnCm = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ DetalleDeEnvio) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DetalleDeEnvio) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DetalleDeEnvio) HintSize(int)                     { panic("Unsupported operation") }
func (_ DetalleDeEnvio) Finalize()                        {}

func (_ DetalleDeEnvio) AvroCRC64Fingerprint() []byte {
	return []byte(DetalleDeEnvioAvroCRC64Fingerprint)
}

func (r DetalleDeEnvio) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["fechaDeRecibo"], err = json.Marshal(r.FechaDeRecibo)
	if err != nil {
		return nil, err
	}
	output["informaTarifa"], err = json.Marshal(r.InformaTarifa)
	if err != nil {
		return nil, err
	}
	output["lote"], err = json.Marshal(r.Lote)
	if err != nil {
		return nil, err
	}
	output["numeroDeEnvio"], err = json.Marshal(r.NumeroDeEnvio)
	if err != nil {
		return nil, err
	}
	output["numeroDeRecibo"], err = json.Marshal(r.NumeroDeRecibo)
	if err != nil {
		return nil, err
	}
	output["numeroInformeDeRendicion"], err = json.Marshal(r.NumeroInformeDeRendicion)
	if err != nil {
		return nil, err
	}
	output["peso"], err = json.Marshal(r.Peso)
	if err != nil {
		return nil, err
	}
	output["pesoAforado"], err = json.Marshal(r.PesoAforado)
	if err != nil {
		return nil, err
	}
	output["valorDeclarado"], err = json.Marshal(r.ValorDeclarado)
	if err != nil {
		return nil, err
	}
	output["valorACobrar"], err = json.Marshal(r.ValorACobrar)
	if err != nil {
		return nil, err
	}
	output["volumen"], err = json.Marshal(r.Volumen)
	if err != nil {
		return nil, err
	}
	output["origen"], err = json.Marshal(r.Origen)
	if err != nil {
		return nil, err
	}
	output["sucursalOrigen"], err = json.Marshal(r.SucursalOrigen)
	if err != nil {
		return nil, err
	}
	output["destino"], err = json.Marshal(r.Destino)
	if err != nil {
		return nil, err
	}
	output["sucursalDestino"], err = json.Marshal(r.SucursalDestino)
	if err != nil {
		return nil, err
	}
	output["categoriaDeFacturacion"], err = json.Marshal(r.CategoriaDeFacturacion)
	if err != nil {
		return nil, err
	}
	output["centroDeCostosDelCliente"], err = json.Marshal(r.CentroDeCostosDelCliente)
	if err != nil {
		return nil, err
	}
	output["estado"], err = json.Marshal(r.Estado)
	if err != nil {
		return nil, err
	}
	output["codigoDeContratoInterno"], err = json.Marshal(r.CodigoDeContratoInterno)
	if err != nil {
		return nil, err
	}
	output["fechaInformeDeRendicion"], err = json.Marshal(r.FechaInformeDeRendicion)
	if err != nil {
		return nil, err
	}
	output["fechaAltaEnvio"], err = json.Marshal(r.FechaAltaEnvio)
	if err != nil {
		return nil, err
	}
	output["ordenDeCompraCliente"], err = json.Marshal(r.OrdenDeCompraCliente)
	if err != nil {
		return nil, err
	}
	output["tarifa"], err = json.Marshal(r.Tarifa)
	if err != nil {
		return nil, err
	}
	output["componentes"], err = json.Marshal(r.Componentes)
	if err != nil {
		return nil, err
	}
	output["destinatario"], err = json.Marshal(r.Destinatario)
	if err != nil {
		return nil, err
	}
	output["remitente"], err = json.Marshal(r.Remitente)
	if err != nil {
		return nil, err
	}
	output["cantidadDeBultos"], err = json.Marshal(r.CantidadDeBultos)
	if err != nil {
		return nil, err
	}
	output["fechaProbableDeEntrega"], err = json.Marshal(r.FechaProbableDeEntrega)
	if err != nil {
		return nil, err
	}
	output["observaciones"], err = json.Marshal(r.Observaciones)
	if err != nil {
		return nil, err
	}
	output["numeroDeRemito"], err = json.Marshal(r.NumeroDeRemito)
	if err != nil {
		return nil, err
	}
	output["numeroDeSeguimientoDelCliente"], err = json.Marshal(r.NumeroDeSeguimientoDelCliente)
	if err != nil {
		return nil, err
	}
	output["datosAdicionales"], err = json.Marshal(r.DatosAdicionales)
	if err != nil {
		return nil, err
	}
	output["fechaPactadaDeEntrega"], err = json.Marshal(r.FechaPactadaDeEntrega)
	if err != nil {
		return nil, err
	}
	output["altoEnCm"], err = json.Marshal(r.AltoEnCm)
	if err != nil {
		return nil, err
	}
	output["anchoEnCm"], err = json.Marshal(r.AnchoEnCm)
	if err != nil {
		return nil, err
	}
	output["largoEnCm"], err = json.Marshal(r.LargoEnCm)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DetalleDeEnvio) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["fechaDeRecibo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaDeRecibo); err != nil {
			return err
		}
	} else {
		r.FechaDeRecibo = NewUnionNullString()

		r.FechaDeRecibo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["informaTarifa"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.InformaTarifa); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for informaTarifa")
	}
	val = func() json.RawMessage {
		if v, ok := fields["lote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Lote); err != nil {
			return err
		}
	} else {
		r.Lote = NewUnionNullString()

		r.Lote = nil
	}
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
		if v, ok := fields["numeroDeRecibo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeRecibo); err != nil {
			return err
		}
	} else {
		r.NumeroDeRecibo = NewUnionNullString()

		r.NumeroDeRecibo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroInformeDeRendicion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroInformeDeRendicion); err != nil {
			return err
		}
	} else {
		r.NumeroInformeDeRendicion = NewUnionNullString()

		r.NumeroInformeDeRendicion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["peso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Peso); err != nil {
			return err
		}
	} else {
		r.Peso = NewUnionNullFloat()

		r.Peso = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["pesoAforado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PesoAforado); err != nil {
			return err
		}
	} else {
		r.PesoAforado = NewUnionNullFloat()

		r.PesoAforado = nil
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
		r.ValorDeclarado = NewUnionNullFloat()

		r.ValorDeclarado = nil
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
		if v, ok := fields["volumen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Volumen); err != nil {
			return err
		}
	} else {
		r.Volumen = NewUnionNullFloat()

		r.Volumen = nil
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
		r.Origen = NewUnionNullLugarDeRetiroEntrega()

		r.Origen = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["sucursalOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SucursalOrigen); err != nil {
			return err
		}
	} else {
		r.SucursalOrigen = NewUnionNullLugarDeRetiroEntrega()

		r.SucursalOrigen = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["destino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Destino); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for destino")
	}
	val = func() json.RawMessage {
		if v, ok := fields["sucursalDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SucursalDestino); err != nil {
			return err
		}
	} else {
		r.SucursalDestino = NewUnionNullLugarDeRetiroEntrega()

		r.SucursalDestino = nil
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
		if v, ok := fields["estado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for estado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoDeContratoInterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeContratoInterno); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigoDeContratoInterno")
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaInformeDeRendicion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaInformeDeRendicion); err != nil {
			return err
		}
	} else {
		r.FechaInformeDeRendicion = NewUnionNullString()

		r.FechaInformeDeRendicion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaAltaEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaAltaEnvio); err != nil {
			return err
		}
	} else {
		r.FechaAltaEnvio = NewUnionNullString()

		r.FechaAltaEnvio = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ordenDeCompraCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenDeCompraCliente); err != nil {
			return err
		}
	} else {
		r.OrdenDeCompraCliente = NewUnionNullString()

		r.OrdenDeCompraCliente = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["tarifa"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Tarifa); err != nil {
			return err
		}
	} else {
		r.Tarifa = NewUnionNullDouble()

		r.Tarifa = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["componentes"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Componentes); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for componentes")
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
		r.Destinatario = NewUnionNullDatosPersonales()

		r.Destinatario = nil
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
		if v, ok := fields["fechaProbableDeEntrega"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaProbableDeEntrega); err != nil {
			return err
		}
	} else {
		r.FechaProbableDeEntrega = NewUnionNullString()

		r.FechaProbableDeEntrega = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["observaciones"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Observaciones); err != nil {
			return err
		}
	} else {
		r.Observaciones = NewUnionNullString()

		r.Observaciones = nil
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
		if v, ok := fields["numeroDeSeguimientoDelCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeSeguimientoDelCliente); err != nil {
			return err
		}
	} else {
		r.NumeroDeSeguimientoDelCliente = NewUnionNullString()

		r.NumeroDeSeguimientoDelCliente = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["datosAdicionales"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DatosAdicionales); err != nil {
			return err
		}
	} else {
		r.DatosAdicionales = NewUnionNullMapString()

		r.DatosAdicionales = nil
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
		r.FechaPactadaDeEntrega = NewUnionNullString()

		r.FechaPactadaDeEntrega = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["altoEnCm"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AltoEnCm); err != nil {
			return err
		}
	} else {
		r.AltoEnCm = NewUnionNullDouble()

		r.AltoEnCm = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["anchoEnCm"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AnchoEnCm); err != nil {
			return err
		}
	} else {
		r.AnchoEnCm = NewUnionNullDouble()

		r.AnchoEnCm = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["largoEnCm"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LargoEnCm); err != nil {
			return err
		}
	} else {
		r.LargoEnCm = NewUnionNullDouble()

		r.LargoEnCm = nil
	}
	return nil
}
