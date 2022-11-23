// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Alta.avsc
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

type Alta struct {
	Traza Traza `json:"traza"`

	Envio DetalleDeEnvio `json:"envio"`

	Contrato Contrato `json:"contrato"`

	CantidadDeBultos *UnionNullInt `json:"cantidadDeBultos"`

	AgrupadorDeBultos *UnionNullString `json:"agrupadorDeBultos"`

	NumeroDeBulto *UnionNullInt `json:"numeroDeBulto"`
}

const AltaAvroCRC64Fingerprint = "B\x8f3M_\xb0\t\xb2"

func NewAlta() Alta {
	r := Alta{}
	r.Traza = NewTraza()

	r.Envio = NewDetalleDeEnvio()

	r.Contrato = NewContrato()

	r.CantidadDeBultos = nil
	r.AgrupadorDeBultos = nil
	r.NumeroDeBulto = nil
	return r
}

func DeserializeAlta(r io.Reader) (Alta, error) {
	t := NewAlta()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAltaFromSchema(r io.Reader, schema string) (Alta, error) {
	t := NewAlta()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAlta(r Alta, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = writeDetalleDeEnvio(r.Envio, w)
	if err != nil {
		return err
	}
	err = writeContrato(r.Contrato, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.CantidadDeBultos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.AgrupadorDeBultos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.NumeroDeBulto, w)
	if err != nil {
		return err
	}
	return err
}

func (r Alta) Serialize(w io.Writer) error {
	return writeAlta(r, w)
}

func (r Alta) Schema() string {
	return "{\"aliases\":[\"AltaManual\",\"AltaAutomatica\",\"AltaRemota\",\"AltaConAnomalias\",\"AltaDeExcedente\"],\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"name\":\"envio\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"fechaDeRecibo\",\"type\":[\"null\",\"string\"]},{\"name\":\"informaTarifa\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"lote\",\"type\":[\"null\",\"string\"]},{\"name\":\"numeroDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"numeroDeRecibo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroInformeDeRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"peso\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"pesoAforado\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"valorDeclarado\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"valorACobrar\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"volumen\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"origen\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"datosSucursal\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosSucursal\"]},{\"default\":null,\"name\":\"domicilio\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"abreviaturaProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoDeDireccion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoCiudad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefono\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"segun ISO -3166-2 (https://es.wikipedia.org/wiki/ISO_3166-2:AR)\",\"name\":\"codigoISOProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"segun ISO 3166-1 alpha-2 (https://es.wikipedia.org/wiki/ISO_3166-1#Tabla_de_c.C3.B3digos_alfa-2_asignados_o_reservados)\",\"name\":\"codigoISOPais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",{\"type\":\"map\",\"values\":\"string\"}]},{\"default\":null,\"name\":\"coordenadas\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"latitud\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"longitud\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"altura\",\"type\":[\"null\",\"double\"]}],\"name\":\"GeoReferencia\",\"type\":\"record\"}]}],\"name\":\"Direccion\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"esLocal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"region\",\"type\":[\"null\",\"string\"]}],\"name\":\"LugarDeRetiroEntrega\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"sucursalOrigen\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.LugarDeRetiroEntrega\"]},{\"name\":\"destino\",\"type\":\"Integracion.Esquemas.Referencias.LugarDeRetiroEntrega\"},{\"default\":null,\"name\":\"sucursalDestino\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.LugarDeRetiroEntrega\"]},{\"default\":null,\"name\":\"categoriaDeFacturacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"centroDeCostosDelCliente\",\"type\":[\"null\",\"string\"]},{\"name\":\"estado\",\"type\":\"string\"},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"fechaInformeDeRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaAltaEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ordenDeCompraCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tarifa\",\"type\":[\"null\",\"double\"]},{\"name\":\"componentes\",\"type\":{\"items\":{\"fields\":[{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idInternoDelCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"eMail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"tipo\",\"type\":{\"name\":\"TipoDeTelefono\",\"symbols\":[\"trabajo\",\"celular\",\"casa\",\"otro\"],\"type\":\"enum\"}},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"agrupador\",\"type\":[\"null\",\"string\"]},{\"name\":\"tipoDeDocumento\",\"type\":{\"name\":\"TipoDeDocumento\",\"symbols\":[\"undefined\",\"DNI\",\"CUIT\",\"CUIL\"],\"type\":\"enum\"}}],\"name\":\"DatosPersonales\",\"type\":\"record\"},\"type\":\"array\"}},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosPersonales\"]},{\"default\":null,\"name\":\"remitente\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosPersonales\"]},{\"default\":null,\"name\":\"cantidadDeBultos\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"fechaProbableDeEntrega\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"observaciones\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeRemito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeSeguimientoDelCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"datosAdicionales\",\"type\":[\"null\",{\"type\":\"map\",\"values\":\"string\"}]},{\"default\":null,\"name\":\"fechaPactadaDeEntrega\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"altoEnCm\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"anchoEnCm\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"largoEnCm\",\"type\":[\"null\",\"double\"]}],\"name\":\"DetalleDeEnvio\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}},{\"name\":\"contrato\",\"type\":{\"fields\":[{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]},{\"name\":\"estaActivo\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"tipoDeServicio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoDeDireccion\",\"type\":[\"null\",\"string\"]},{\"name\":\"codigoDeClienteInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"vigenciaDesde\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vigenciaHasta\",\"type\":[\"null\",\"string\"]}],\"name\":\"Contrato\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}},{\"default\":null,\"name\":\"cantidadDeBultos\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"agrupadorDeBultos\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeBulto\",\"type\":[\"null\",\"int\"]}],\"name\":\"Integracion.Esquemas.Trazas.Alta\",\"type\":\"record\"}"
}

func (r Alta) SchemaName() string {
	return "Integracion.Esquemas.Trazas.Alta"
}

func (_ Alta) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Alta) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Alta) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Alta) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Alta) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Alta) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Alta) SetString(v string)   { panic("Unsupported operation") }
func (_ Alta) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Alta) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		r.Envio = NewDetalleDeEnvio()

		w := types.Record{Target: &r.Envio}

		return w

	case 2:
		r.Contrato = NewContrato()

		w := types.Record{Target: &r.Contrato}

		return w

	case 3:
		r.CantidadDeBultos = NewUnionNullInt()

		return r.CantidadDeBultos
	case 4:
		r.AgrupadorDeBultos = NewUnionNullString()

		return r.AgrupadorDeBultos
	case 5:
		r.NumeroDeBulto = NewUnionNullInt()

		return r.NumeroDeBulto
	}
	panic("Unknown field index")
}

func (r *Alta) SetDefault(i int) {
	switch i {
	case 3:
		r.CantidadDeBultos = nil
		return
	case 4:
		r.AgrupadorDeBultos = nil
		return
	case 5:
		r.NumeroDeBulto = nil
		return
	}
	panic("Unknown field index")
}

func (r *Alta) NullField(i int) {
	switch i {
	case 3:
		r.CantidadDeBultos = nil
		return
	case 4:
		r.AgrupadorDeBultos = nil
		return
	case 5:
		r.NumeroDeBulto = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Alta) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Alta) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Alta) HintSize(int)                     { panic("Unsupported operation") }
func (_ Alta) Finalize()                        {}

func (_ Alta) AvroCRC64Fingerprint() []byte {
	return []byte(AltaAvroCRC64Fingerprint)
}

func (r Alta) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["envio"], err = json.Marshal(r.Envio)
	if err != nil {
		return nil, err
	}
	output["contrato"], err = json.Marshal(r.Contrato)
	if err != nil {
		return nil, err
	}
	output["cantidadDeBultos"], err = json.Marshal(r.CantidadDeBultos)
	if err != nil {
		return nil, err
	}
	output["agrupadorDeBultos"], err = json.Marshal(r.AgrupadorDeBultos)
	if err != nil {
		return nil, err
	}
	output["numeroDeBulto"], err = json.Marshal(r.NumeroDeBulto)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Alta) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["envio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Envio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for envio")
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
		if v, ok := fields["agrupadorDeBultos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AgrupadorDeBultos); err != nil {
			return err
		}
	} else {
		r.AgrupadorDeBultos = NewUnionNullString()

		r.AgrupadorDeBultos = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroDeBulto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeBulto); err != nil {
			return err
		}
	} else {
		r.NumeroDeBulto = NewUnionNullInt()

		r.NumeroDeBulto = nil
	}
	return nil
}
