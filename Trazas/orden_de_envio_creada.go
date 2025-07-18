// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     OrdenDeEnvioCreada.avsc
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

type OrdenDeEnvioCreada struct {
	Traza Traza `json:"traza"`

	DatosDeLaOrden DetalleDeOrdenDeEnvio `json:"datosDeLaOrden"`

	Linking *UnionNullMapString `json:"linking"`
}

const OrdenDeEnvioCreadaAvroCRC64Fingerprint = "\xefj\x92\x1c\x16_\xd6\n"

func NewOrdenDeEnvioCreada() OrdenDeEnvioCreada {
	r := OrdenDeEnvioCreada{}
	r.Traza = NewTraza()

	r.DatosDeLaOrden = NewDetalleDeOrdenDeEnvio()

	r.Linking = nil
	return r
}

func DeserializeOrdenDeEnvioCreada(r io.Reader) (OrdenDeEnvioCreada, error) {
	t := NewOrdenDeEnvioCreada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeOrdenDeEnvioCreadaFromSchema(r io.Reader, schema string) (OrdenDeEnvioCreada, error) {
	t := NewOrdenDeEnvioCreada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeOrdenDeEnvioCreada(r OrdenDeEnvioCreada, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = writeDetalleDeOrdenDeEnvio(r.DatosDeLaOrden, w)
	if err != nil {
		return err
	}
	err = writeUnionNullMapString(r.Linking, w)
	if err != nil {
		return err
	}
	return err
}

func (r OrdenDeEnvioCreada) Serialize(w io.Writer) error {
	return writeOrdenDeEnvioCreada(r, w)
}

func (r OrdenDeEnvioCreada) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"name\":\"datosDeLaOrden\",\"type\":{\"fields\":[{\"name\":\"numeroDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"remitenteInformado\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idInternoDelCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"eMail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"tipo\",\"type\":{\"name\":\"TipoDeTelefono\",\"symbols\":[\"trabajo\",\"celular\",\"casa\",\"otro\"],\"type\":\"enum\"}},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"agrupador\",\"type\":[\"null\",\"string\"]},{\"name\":\"tipoDeDocumento\",\"type\":{\"name\":\"TipoDeDocumento\",\"symbols\":[\"undefined\",\"DNI\",\"CUIT\",\"CUIL\"],\"type\":\"enum\"}}],\"name\":\"DatosPersonales\",\"type\":\"record\"}]},{\"name\":\"destinatarioInformado\",\"type\":\"Integracion.Esquemas.Referencias.DatosPersonales\"},{\"default\":null,\"name\":\"destinatarioAlternativos\",\"type\":[\"null\",{\"items\":\"Integracion.Esquemas.Referencias.DatosPersonales\",\"type\":\"array\"}]},{\"name\":\"destinoInformado\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"datosSucursal\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosSucursal\"]},{\"default\":null,\"name\":\"domicilio\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"abreviaturaProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoDeDireccion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoCiudad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefono\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"segun ISO -3166-2 (https://es.wikipedia.org/wiki/ISO_3166-2:AR)\",\"name\":\"codigoISOProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"segun ISO 3166-1 alpha-2 (https://es.wikipedia.org/wiki/ISO_3166-1#Tabla_de_c.C3.B3digos_alfa-2_asignados_o_reservados)\",\"name\":\"codigoISOPais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",{\"type\":\"map\",\"values\":\"string\"}]},{\"default\":null,\"name\":\"coordenadas\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"latitud\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"longitud\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"altura\",\"type\":[\"null\",\"double\"]}],\"name\":\"GeoReferencia\",\"type\":\"record\"}]}],\"name\":\"Direccion\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"esLocal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"region\",\"type\":[\"null\",\"string\"]}],\"name\":\"LugarDeRetiroEntrega\",\"type\":\"record\"}},{\"default\":null,\"name\":\"origenInformado\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.LugarDeRetiroEntrega\"]},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoDeServicio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeRemito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remitosComplementarios\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"centroDeCostosDelCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"categoriaDeProducto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"categoriaDeFacturacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"valorACobrar\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"fechaPactadaDeEntrega\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"fecha\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"horaDesde\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"horaHasta\",\"type\":[\"null\",\"string\"]}],\"name\":\"FechaPactada\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"productoAEntregar\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"productoARetirar\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pagoDestino\",\"type\":[\"null\",{\"name\":\"TipoPagoDestino\",\"symbols\":[\"undefined\",\"P\",\"D\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"sucursalDeDistribucion\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosSucursal\"]},{\"default\":null,\"name\":\"sucursalCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"listaDePaquetes\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"pesoEnKg\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"altoEnCm\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"anchoEnCm\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"largoEnCm\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"referenciasDelCliente\",\"type\":[\"null\",{\"type\":\"map\",\"values\":\"string\"}]},{\"default\":null,\"name\":\"volumenEnCm3\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"valorDeclaradoSinImpuesto\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"valorDeclaradoConImpuesto\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"numeroDeBulto\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"valorDeclarado\",\"type\":[\"null\",\"float\"]}],\"name\":\"DetalleDePaquete\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"codigoVerificadorDeEntrega\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cantidadDeBultos\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"agrupadorDeBulto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idPedido\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"unidadOperativaDeOrigen\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosSucursal\"]}],\"name\":\"DetalleDeOrdenDeEnvio\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}},{\"default\":null,\"name\":\"linking\",\"type\":[\"null\",{\"type\":\"map\",\"values\":\"string\"}]}],\"name\":\"Integracion.Esquemas.Trazas.OrdenDeEnvioCreada\",\"type\":\"record\"}"
}

func (r OrdenDeEnvioCreada) SchemaName() string {
	return "Integracion.Esquemas.Trazas.OrdenDeEnvioCreada"
}

func (_ OrdenDeEnvioCreada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ OrdenDeEnvioCreada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ OrdenDeEnvioCreada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ OrdenDeEnvioCreada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ OrdenDeEnvioCreada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ OrdenDeEnvioCreada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ OrdenDeEnvioCreada) SetString(v string)   { panic("Unsupported operation") }
func (_ OrdenDeEnvioCreada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *OrdenDeEnvioCreada) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		r.DatosDeLaOrden = NewDetalleDeOrdenDeEnvio()

		w := types.Record{Target: &r.DatosDeLaOrden}

		return w

	case 2:
		r.Linking = NewUnionNullMapString()

		return r.Linking
	}
	panic("Unknown field index")
}

func (r *OrdenDeEnvioCreada) SetDefault(i int) {
	switch i {
	case 2:
		r.Linking = nil
		return
	}
	panic("Unknown field index")
}

func (r *OrdenDeEnvioCreada) NullField(i int) {
	switch i {
	case 2:
		r.Linking = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ OrdenDeEnvioCreada) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ OrdenDeEnvioCreada) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ OrdenDeEnvioCreada) HintSize(int)                     { panic("Unsupported operation") }
func (_ OrdenDeEnvioCreada) Finalize()                        {}

func (_ OrdenDeEnvioCreada) AvroCRC64Fingerprint() []byte {
	return []byte(OrdenDeEnvioCreadaAvroCRC64Fingerprint)
}

func (r OrdenDeEnvioCreada) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["datosDeLaOrden"], err = json.Marshal(r.DatosDeLaOrden)
	if err != nil {
		return nil, err
	}
	output["linking"], err = json.Marshal(r.Linking)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *OrdenDeEnvioCreada) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["traza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Traza); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for traza")
	}
	val = func() json.RawMessage {
		if v, ok := fields["datosDeLaOrden"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DatosDeLaOrden); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for datosDeLaOrden")
	}
	val = func() json.RawMessage {
		if v, ok := fields["linking"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Linking); err != nil {
			return err
		}
	} else {
		r.Linking = NewUnionNullMapString()

		r.Linking = nil
	}
	return nil
}
