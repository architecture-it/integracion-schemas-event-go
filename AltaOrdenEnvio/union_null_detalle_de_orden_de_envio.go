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

type UnionNullDetalleDeOrdenDeEnvioTypeEnum int

const (
	UnionNullDetalleDeOrdenDeEnvioTypeEnumDetalleDeOrdenDeEnvio UnionNullDetalleDeOrdenDeEnvioTypeEnum = 1
)

type UnionNullDetalleDeOrdenDeEnvio struct {
	Null                  *types.NullVal
	DetalleDeOrdenDeEnvio DetalleDeOrdenDeEnvio
	UnionType             UnionNullDetalleDeOrdenDeEnvioTypeEnum
}

func writeUnionNullDetalleDeOrdenDeEnvio(r *UnionNullDetalleDeOrdenDeEnvio, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullDetalleDeOrdenDeEnvioTypeEnumDetalleDeOrdenDeEnvio:
		return writeDetalleDeOrdenDeEnvio(r.DetalleDeOrdenDeEnvio, w)
	}
	return fmt.Errorf("invalid value for *UnionNullDetalleDeOrdenDeEnvio")
}

func NewUnionNullDetalleDeOrdenDeEnvio() *UnionNullDetalleDeOrdenDeEnvio {
	return &UnionNullDetalleDeOrdenDeEnvio{}
}

func (r *UnionNullDetalleDeOrdenDeEnvio) Serialize(w io.Writer) error {
	return writeUnionNullDetalleDeOrdenDeEnvio(r, w)
}

func DeserializeUnionNullDetalleDeOrdenDeEnvio(r io.Reader) (*UnionNullDetalleDeOrdenDeEnvio, error) {
	t := NewUnionNullDetalleDeOrdenDeEnvio()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionNullDetalleDeOrdenDeEnvioFromSchema(r io.Reader, schema string) (*UnionNullDetalleDeOrdenDeEnvio, error) {
	t := NewUnionNullDetalleDeOrdenDeEnvio()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r *UnionNullDetalleDeOrdenDeEnvio) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"numeroDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"remitente\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idInternoDelCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"eMail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"tipo\",\"type\":\"string\"},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDeTelefonos\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"agrupador\",\"type\":[\"null\",\"string\"]},{\"name\":\"tipoDeDocumento\",\"type\":{\"name\":\"TipoDeDocumento\",\"symbols\":[\"undefined\",\"DNI\",\"CUIT\",\"CUIL\"],\"type\":\"enum\"}}],\"name\":\"DatosPersonales\",\"type\":\"record\"}]},{\"name\":\"destinatario\",\"type\":\"Andreani.AltaOrdenEnvio.Events.Common.DatosPersonales\"},{\"default\":null,\"name\":\"destinatarioAlternativo\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.DatosPersonales\"]},{\"name\":\"destino\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"postal\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"abreviaturaProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoDeDireccion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoCiudad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefono\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"segun ISO -3166-2 (https://es.wikipedia.org/wiki/ISO_3166-2:AR)\",\"name\":\"codigoISOProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"segun ISO 3166-1 alpha-2 (https://es.wikipedia.org/wiki/ISO_3166-1#Tabla_de_c.C3.B3digos_alfa-2_asignados_o_reservados)\",\"name\":\"codigoISOPais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"casillaDeCorreo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"region\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]}],\"name\":\"Direccion\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"datosSucursal\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"ID\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nomenclatura\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"direccion\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Region\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDePropiedades\"]},{\"default\":null,\"name\":\"casillaDeCorreo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"region\",\"type\":[\"null\",\"string\"]}],\"name\":\"DireccionPostal\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"datosAdicionales\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDePropiedades\"]},{\"default\":null,\"name\":\"listaDeTelefonos\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDeTelefonos\"]}],\"name\":\"Sucursal\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"GeoCoordenada\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"latitud\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"longitud\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"elevacion\",\"type\":[\"null\",\"double\"]}],\"name\":\"GeoCoordenada\",\"type\":\"record\"}]}],\"name\":\"LugarDeRetiroOEntrega\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"origen\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.LugarDeRetiroOEntrega\"]},{\"name\":\"contrato\",\"type\":\"string\"},{\"default\":null,\"name\":\"tipoDeServicio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeRemito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remitosComplementarios\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"centroDeCostosDelCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"categoriaDeProducto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"categoriaDeFacturacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"valorACobrar\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"fechaPactadaDeEntrega\",\"type\":[\"null\",{\"fields\":[{\"name\":\"fecha\",\"type\":[\"null\",\"string\"]},{\"name\":\"horaDesde\",\"type\":[\"null\",\"string\"]},{\"name\":\"horaHasta\",\"type\":[\"null\",\"string\"]}],\"name\":\"FechaPactadaDeEntrega\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"productoAEntregar\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"productoARetirar\",\"type\":[\"null\",\"string\"]},{\"name\":\"pagoDestino\",\"type\":{\"name\":\"TipoPagoDestino\",\"symbols\":[\"undefined\",\"P\",\"D\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"sucursalDeDistribucion\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"sucursalCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"listaPaquetes\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaPaquetes\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"pesoEnKg\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"altoEnCm\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"anchoEnCm\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"largoEnCm\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"referenciasDelCliente\",\"type\":[\"null\",{\"type\":\"map\",\"values\":\"string\"}]},{\"default\":null,\"name\":\"volumenEnCm3\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"valorDeclaradoSinImpuesto\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"valorDeclaradoConImpuesto\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"numeroDeBulto\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"valorDeclarado\",\"type\":[\"null\",\"double\"]}],\"name\":\"DetalleDePaquete\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaPaquetes\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"codigoVerificadorDeEntrega\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cantidadDeBultos\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"agrupadorDeBulto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idPedido\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pagoEnMostrador\",\"type\":[\"null\",{\"fields\":[{\"name\":\"montoACobrar\",\"type\":[\"null\",\"int\"]},{\"name\":\"documentoTipo\",\"type\":[\"null\",\"string\"]},{\"name\":\"documentoNumero\",\"type\":[\"null\",\"string\"]}],\"name\":\"DatosPagoEnMostrador\",\"type\":\"record\"}]}],\"name\":\"DetalleDeOrdenDeEnvio\",\"namespace\":\"Andreani.AltaOrdenEnvio.Events.Common\",\"type\":\"record\"}]"
}

func (_ *UnionNullDetalleDeOrdenDeEnvio) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullDetalleDeOrdenDeEnvio) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullDetalleDeOrdenDeEnvio) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullDetalleDeOrdenDeEnvio) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullDetalleDeOrdenDeEnvio) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullDetalleDeOrdenDeEnvio) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullDetalleDeOrdenDeEnvio) SetLong(v int64) {

	r.UnionType = (UnionNullDetalleDeOrdenDeEnvioTypeEnum)(v)
}

func (r *UnionNullDetalleDeOrdenDeEnvio) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.DetalleDeOrdenDeEnvio = NewDetalleDeOrdenDeEnvio()
		return &types.Record{Target: (&r.DetalleDeOrdenDeEnvio)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullDetalleDeOrdenDeEnvio) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullDetalleDeOrdenDeEnvio) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullDetalleDeOrdenDeEnvio) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullDetalleDeOrdenDeEnvio) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullDetalleDeOrdenDeEnvio) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullDetalleDeOrdenDeEnvio) Finalize()                {}

func (r *UnionNullDetalleDeOrdenDeEnvio) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullDetalleDeOrdenDeEnvioTypeEnumDetalleDeOrdenDeEnvio:
		return json.Marshal(map[string]interface{}{"Andreani.AltaOrdenEnvio.Events.Common.DetalleDeOrdenDeEnvio": r.DetalleDeOrdenDeEnvio})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullDetalleDeOrdenDeEnvio")
}

func (r *UnionNullDetalleDeOrdenDeEnvio) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.AltaOrdenEnvio.Events.Common.DetalleDeOrdenDeEnvio"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.DetalleDeOrdenDeEnvio)
	}
	return fmt.Errorf("invalid value for *UnionNullDetalleDeOrdenDeEnvio")
}
