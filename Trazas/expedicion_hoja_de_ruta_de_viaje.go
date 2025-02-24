// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ExpedicionHojaDeRutaDeViaje.avsc
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

type ExpedicionHojaDeRutaDeViaje struct {
	Traza Traza `json:"traza"`

	HojaDeRuta string `json:"hojaDeRuta"`

	Distribuidor *UnionNullDatosDistribuidor `json:"distribuidor"`

	SucursalQueExpide *UnionNullDatosSucursal `json:"sucursalQueExpide"`

	SucursalQueRecibe *UnionNullDatosSucursal `json:"sucursalQueRecibe"`

	MedioDeExpedicion *UnionNullString `json:"medioDeExpedicion"`
}

const ExpedicionHojaDeRutaDeViajeAvroCRC64Fingerprint = "=\xc1\x82ԍ8El"

func NewExpedicionHojaDeRutaDeViaje() ExpedicionHojaDeRutaDeViaje {
	r := ExpedicionHojaDeRutaDeViaje{}
	r.Traza = NewTraza()

	r.Distribuidor = nil
	r.SucursalQueExpide = nil
	r.SucursalQueRecibe = nil
	r.MedioDeExpedicion = nil
	return r
}

func DeserializeExpedicionHojaDeRutaDeViaje(r io.Reader) (ExpedicionHojaDeRutaDeViaje, error) {
	t := NewExpedicionHojaDeRutaDeViaje()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeExpedicionHojaDeRutaDeViajeFromSchema(r io.Reader, schema string) (ExpedicionHojaDeRutaDeViaje, error) {
	t := NewExpedicionHojaDeRutaDeViaje()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeExpedicionHojaDeRutaDeViaje(r ExpedicionHojaDeRutaDeViaje, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
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
	err = writeUnionNullDatosSucursal(r.SucursalQueExpide, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDatosSucursal(r.SucursalQueRecibe, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.MedioDeExpedicion, w)
	if err != nil {
		return err
	}
	return err
}

func (r ExpedicionHojaDeRutaDeViaje) Serialize(w io.Writer) error {
	return writeExpedicionHojaDeRutaDeViaje(r, w)
}

func (r ExpedicionHojaDeRutaDeViaje) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"LogisticShipmentType\",\"type\":[\"null\",\"string\"]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"name\":\"hojaDeRuta\",\"type\":\"string\"},{\"default\":null,\"name\":\"distribuidor\",\"type\":[\"null\",{\"fields\":[{\"name\":\"datosPersonales\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idInternoDelCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"eMail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"tipo\",\"type\":{\"name\":\"TipoDeTelefono\",\"symbols\":[\"trabajo\",\"celular\",\"casa\",\"otro\"],\"type\":\"enum\"}},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"agrupador\",\"type\":[\"null\",\"string\"]},{\"name\":\"tipoDeDocumento\",\"type\":{\"name\":\"TipoDeDocumento\",\"symbols\":[\"undefined\",\"DNI\",\"CUIT\",\"CUIL\"],\"type\":\"enum\"}}],\"name\":\"DatosPersonales\",\"type\":\"record\"}},{\"default\":null,\"name\":\"medioDeLocomocion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoDeDistribuidor\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"legajo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalDondeTrabaja\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosSucursal\"]},{\"default\":null,\"name\":\"esEventual\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"login\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idgla\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cuit\",\"type\":[\"null\",\"string\"]}],\"name\":\"DatosDistribuidor\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"sucursalQueExpide\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosSucursal\"]},{\"default\":null,\"name\":\"sucursalQueRecibe\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosSucursal\"]},{\"default\":null,\"name\":\"medioDeExpedicion\",\"type\":[\"null\",\"string\"]}],\"name\":\"Integracion.Esquemas.Trazas.ExpedicionHojaDeRutaDeViaje\",\"type\":\"record\"}"
}

func (r ExpedicionHojaDeRutaDeViaje) SchemaName() string {
	return "Integracion.Esquemas.Trazas.ExpedicionHojaDeRutaDeViaje"
}

func (_ ExpedicionHojaDeRutaDeViaje) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ExpedicionHojaDeRutaDeViaje) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ExpedicionHojaDeRutaDeViaje) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ExpedicionHojaDeRutaDeViaje) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ExpedicionHojaDeRutaDeViaje) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ExpedicionHojaDeRutaDeViaje) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ExpedicionHojaDeRutaDeViaje) SetString(v string)   { panic("Unsupported operation") }
func (_ ExpedicionHojaDeRutaDeViaje) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ExpedicionHojaDeRutaDeViaje) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		w := types.String{Target: &r.HojaDeRuta}

		return w

	case 2:
		r.Distribuidor = NewUnionNullDatosDistribuidor()

		return r.Distribuidor
	case 3:
		r.SucursalQueExpide = NewUnionNullDatosSucursal()

		return r.SucursalQueExpide
	case 4:
		r.SucursalQueRecibe = NewUnionNullDatosSucursal()

		return r.SucursalQueRecibe
	case 5:
		r.MedioDeExpedicion = NewUnionNullString()

		return r.MedioDeExpedicion
	}
	panic("Unknown field index")
}

func (r *ExpedicionHojaDeRutaDeViaje) SetDefault(i int) {
	switch i {
	case 2:
		r.Distribuidor = nil
		return
	case 3:
		r.SucursalQueExpide = nil
		return
	case 4:
		r.SucursalQueRecibe = nil
		return
	case 5:
		r.MedioDeExpedicion = nil
		return
	}
	panic("Unknown field index")
}

func (r *ExpedicionHojaDeRutaDeViaje) NullField(i int) {
	switch i {
	case 2:
		r.Distribuidor = nil
		return
	case 3:
		r.SucursalQueExpide = nil
		return
	case 4:
		r.SucursalQueRecibe = nil
		return
	case 5:
		r.MedioDeExpedicion = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ExpedicionHojaDeRutaDeViaje) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ExpedicionHojaDeRutaDeViaje) AppendArray() types.Field { panic("Unsupported operation") }
func (_ ExpedicionHojaDeRutaDeViaje) HintSize(int)             { panic("Unsupported operation") }
func (_ ExpedicionHojaDeRutaDeViaje) Finalize()                {}

func (_ ExpedicionHojaDeRutaDeViaje) AvroCRC64Fingerprint() []byte {
	return []byte(ExpedicionHojaDeRutaDeViajeAvroCRC64Fingerprint)
}

func (r ExpedicionHojaDeRutaDeViaje) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
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
	output["sucursalQueExpide"], err = json.Marshal(r.SucursalQueExpide)
	if err != nil {
		return nil, err
	}
	output["sucursalQueRecibe"], err = json.Marshal(r.SucursalQueRecibe)
	if err != nil {
		return nil, err
	}
	output["medioDeExpedicion"], err = json.Marshal(r.MedioDeExpedicion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ExpedicionHojaDeRutaDeViaje) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["sucursalQueExpide"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SucursalQueExpide); err != nil {
			return err
		}
	} else {
		r.SucursalQueExpide = NewUnionNullDatosSucursal()

		r.SucursalQueExpide = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["sucursalQueRecibe"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SucursalQueRecibe); err != nil {
			return err
		}
	} else {
		r.SucursalQueRecibe = NewUnionNullDatosSucursal()

		r.SucursalQueRecibe = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["medioDeExpedicion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MedioDeExpedicion); err != nil {
			return err
		}
	} else {
		r.MedioDeExpedicion = NewUnionNullString()

		r.MedioDeExpedicion = nil
	}
	return nil
}
