// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     BultoDistribucion.avsc
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

type BultoDistribucion struct {
	Traza TrazaDeBulto `json:"traza"`

	Bulto *UnionNullDetalleDelBulto `json:"bulto"`

	HojaDeRuta string `json:"hojaDeRuta"`

	Distribuidor *UnionNullDatosDistribuidor `json:"distribuidor"`

	DesdeDonde *UnionNullDatosSucursal `json:"desdeDonde"`
}

const BultoDistribucionAvroCRC64Fingerprint = "\xb4\xf04\xed%\x12\x9d\xfe"

func NewBultoDistribucion() BultoDistribucion {
	r := BultoDistribucion{}
	r.Traza = NewTrazaDeBulto()

	r.Bulto = nil
	r.Distribuidor = nil
	r.DesdeDonde = nil
	return r
}

func DeserializeBultoDistribucion(r io.Reader) (BultoDistribucion, error) {
	t := NewBultoDistribucion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBultoDistribucionFromSchema(r io.Reader, schema string) (BultoDistribucion, error) {
	t := NewBultoDistribucion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBultoDistribucion(r BultoDistribucion, w io.Writer) error {
	var err error
	err = writeTrazaDeBulto(r.Traza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDetalleDelBulto(r.Bulto, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.HojaDeRuta, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDatosDistribuidor(r.Distribuidor, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDatosSucursal(r.DesdeDonde, w)
	if err != nil {
		return err
	}
	return err
}

func (r BultoDistribucion) Serialize(w io.Writer) error {
	return writeBultoDistribucion(r, w)
}

func (r BultoDistribucion) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"name\":\"numeroDeBulto\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"tipoDeBulto\",\"type\":[\"null\",{\"name\":\"TipoDeBulto\",\"symbols\":[\"Paquete\",\"Contenedor\"],\"type\":\"enum\"}]}],\"name\":\"TrazaDeBulto\",\"namespace\":\"Integracion.Esquemas.Bulto.Referencias\",\"type\":\"record\"}},{\"default\":null,\"name\":\"bulto\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"pesoEnKg\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"altoEnCm\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"anchoEnCm\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"largoEnCm\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"referenciasDelCliente\",\"type\":[\"null\",{\"type\":\"map\",\"values\":\"string\"}]},{\"default\":null,\"name\":\"volumenEnCm3\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"valorDeclaradoSinImpuesto\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"valorDeclaradoConImpuesto\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"numeroDeBulto\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"valorDeclarado\",\"type\":[\"null\",\"float\"]}],\"name\":\"DetalleDelBulto\",\"namespace\":\"Integracion.Esquemas.Bulto.Referencias\",\"type\":\"record\"}]},{\"name\":\"hojaDeRuta\",\"type\":\"string\"},{\"default\":null,\"name\":\"distribuidor\",\"type\":[\"null\",{\"fields\":[{\"name\":\"datosPersonales\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idInternoDelCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"eMail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"tipo\",\"type\":{\"name\":\"TipoDeTelefono\",\"symbols\":[\"trabajo\",\"celular\",\"casa\",\"otro\"],\"type\":\"enum\"}},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"agrupador\",\"type\":[\"null\",\"string\"]},{\"name\":\"tipoDeDocumento\",\"type\":{\"name\":\"TipoDeDocumento\",\"symbols\":[\"undefined\",\"DNI\",\"CUIT\",\"CUIL\"],\"type\":\"enum\"}}],\"name\":\"DatosPersonales\",\"type\":\"record\"}},{\"default\":null,\"name\":\"medioDeLocomocion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoDeDistribuidor\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"legajo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalDondeTrabaja\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosSucursal\"]},{\"default\":null,\"name\":\"esEventual\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"login\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idgla\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cuit\",\"type\":[\"null\",\"string\"]}],\"name\":\"DatosDistribuidor\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"desdeDonde\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosSucursal\"]}],\"name\":\"Integracion.Esquemas.Bulto.Trazas.BultoDistribucion\",\"type\":\"record\"}"
}

func (r BultoDistribucion) SchemaName() string {
	return "Integracion.Esquemas.Bulto.Trazas.BultoDistribucion"
}

func (_ BultoDistribucion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BultoDistribucion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BultoDistribucion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BultoDistribucion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BultoDistribucion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BultoDistribucion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BultoDistribucion) SetString(v string)   { panic("Unsupported operation") }
func (_ BultoDistribucion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BultoDistribucion) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTrazaDeBulto()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		r.Bulto = NewUnionNullDetalleDelBulto()

		return r.Bulto
	case 2:
		w := types.String{Target: &r.HojaDeRuta}

		return w

	case 3:
		r.Distribuidor = NewUnionNullDatosDistribuidor()

		return r.Distribuidor
	case 4:
		r.DesdeDonde = NewUnionNullDatosSucursal()

		return r.DesdeDonde
	}
	panic("Unknown field index")
}

func (r *BultoDistribucion) SetDefault(i int) {
	switch i {
	case 1:
		r.Bulto = nil
		return
	case 3:
		r.Distribuidor = nil
		return
	case 4:
		r.DesdeDonde = nil
		return
	}
	panic("Unknown field index")
}

func (r *BultoDistribucion) NullField(i int) {
	switch i {
	case 1:
		r.Bulto = nil
		return
	case 3:
		r.Distribuidor = nil
		return
	case 4:
		r.DesdeDonde = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ BultoDistribucion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ BultoDistribucion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ BultoDistribucion) HintSize(int)                     { panic("Unsupported operation") }
func (_ BultoDistribucion) Finalize()                        {}

func (_ BultoDistribucion) AvroCRC64Fingerprint() []byte {
	return []byte(BultoDistribucionAvroCRC64Fingerprint)
}

func (r BultoDistribucion) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["bulto"], err = json.Marshal(r.Bulto)
	if err != nil {
		return nil, err
	}
	output["hojaDeRuta"], err = json.Marshal(r.HojaDeRuta)
	if err != nil {
		return nil, err
	}
	output["distribuidor"], err = json.Marshal(r.Distribuidor)
	if err != nil {
		return nil, err
	}
	output["desdeDonde"], err = json.Marshal(r.DesdeDonde)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *BultoDistribucion) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["bulto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Bulto); err != nil {
			return err
		}
	} else {
		r.Bulto = NewUnionNullDetalleDelBulto()

		r.Bulto = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["hojaDeRuta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.HojaDeRuta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for hojaDeRuta")
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
		r.Distribuidor = NewUnionNullDatosDistribuidor()

		r.Distribuidor = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["desdeDonde"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DesdeDonde); err != nil {
			return err
		}
	} else {
		r.DesdeDonde = NewUnionNullDatosSucursal()

		r.DesdeDonde = nil
	}
	return nil
}
