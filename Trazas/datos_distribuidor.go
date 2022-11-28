// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Visita.avsc
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

type DatosDistribuidor struct {
	DatosPersonales DatosPersonales `json:"datosPersonales"`

	MedioDeLocomocion *UnionNullString `json:"medioDeLocomocion"`

	TipoDeDistribuidor *UnionNullString `json:"tipoDeDistribuidor"`

	Legajo *UnionNullString `json:"legajo"`

	SucursalDondeTrabaja *UnionNullDatosSucursal `json:"sucursalDondeTrabaja"`

	EsEventual *UnionNullBool `json:"esEventual"`

	Login *UnionNullString `json:"login"`

	Idgla *UnionNullString `json:"idgla"`

	Cuit *UnionNullString `json:"cuit"`
}

const DatosDistribuidorAvroCRC64Fingerprint = "\x18\x00\xfe\a\x06\x82'I"

func NewDatosDistribuidor() DatosDistribuidor {
	r := DatosDistribuidor{}
	r.DatosPersonales = NewDatosPersonales()

	r.MedioDeLocomocion = nil
	r.TipoDeDistribuidor = nil
	r.Legajo = nil
	r.SucursalDondeTrabaja = nil
	r.EsEventual = nil
	r.Login = nil
	r.Idgla = nil
	r.Cuit = nil
	return r
}

func DeserializeDatosDistribuidor(r io.Reader) (DatosDistribuidor, error) {
	t := NewDatosDistribuidor()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDatosDistribuidorFromSchema(r io.Reader, schema string) (DatosDistribuidor, error) {
	t := NewDatosDistribuidor()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDatosDistribuidor(r DatosDistribuidor, w io.Writer) error {
	var err error
	err = writeDatosPersonales(r.DatosPersonales, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.MedioDeLocomocion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoDeDistribuidor, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Legajo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDatosSucursal(r.SucursalDondeTrabaja, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.EsEventual, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Login, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Idgla, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Cuit, w)
	if err != nil {
		return err
	}
	return err
}

func (r DatosDistribuidor) Serialize(w io.Writer) error {
	return writeDatosDistribuidor(r, w)
}

func (r DatosDistribuidor) Schema() string {
	return "{\"fields\":[{\"name\":\"datosPersonales\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idInternoDelCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"eMail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"tipo\",\"type\":{\"name\":\"TipoDeTelefono\",\"symbols\":[\"trabajo\",\"celular\",\"casa\",\"otro\"],\"type\":\"enum\"}},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"agrupador\",\"type\":[\"null\",\"string\"]},{\"name\":\"tipoDeDocumento\",\"type\":{\"name\":\"TipoDeDocumento\",\"symbols\":[\"undefined\",\"DNI\",\"CUIT\",\"CUIL\"],\"type\":\"enum\"}}],\"name\":\"DatosPersonales\",\"type\":\"record\"}},{\"default\":null,\"name\":\"medioDeLocomocion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoDeDistribuidor\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"legajo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalDondeTrabaja\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"esEventual\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"login\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idgla\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cuit\",\"type\":[\"null\",\"string\"]}],\"name\":\"Integracion.Esquemas.Referencias.DatosDistribuidor\",\"type\":\"record\"}"
}

func (r DatosDistribuidor) SchemaName() string {
	return "Integracion.Esquemas.Referencias.DatosDistribuidor"
}

func (_ DatosDistribuidor) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DatosDistribuidor) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DatosDistribuidor) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DatosDistribuidor) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DatosDistribuidor) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DatosDistribuidor) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DatosDistribuidor) SetString(v string)   { panic("Unsupported operation") }
func (_ DatosDistribuidor) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DatosDistribuidor) Get(i int) types.Field {
	switch i {
	case 0:
		r.DatosPersonales = NewDatosPersonales()

		w := types.Record{Target: &r.DatosPersonales}

		return w

	case 1:
		r.MedioDeLocomocion = NewUnionNullString()

		return r.MedioDeLocomocion
	case 2:
		r.TipoDeDistribuidor = NewUnionNullString()

		return r.TipoDeDistribuidor
	case 3:
		r.Legajo = NewUnionNullString()

		return r.Legajo
	case 4:
		r.SucursalDondeTrabaja = NewUnionNullDatosSucursal()

		return r.SucursalDondeTrabaja
	case 5:
		r.EsEventual = NewUnionNullBool()

		return r.EsEventual
	case 6:
		r.Login = NewUnionNullString()

		return r.Login
	case 7:
		r.Idgla = NewUnionNullString()

		return r.Idgla
	case 8:
		r.Cuit = NewUnionNullString()

		return r.Cuit
	}
	panic("Unknown field index")
}

func (r *DatosDistribuidor) SetDefault(i int) {
	switch i {
	case 1:
		r.MedioDeLocomocion = nil
		return
	case 2:
		r.TipoDeDistribuidor = nil
		return
	case 3:
		r.Legajo = nil
		return
	case 4:
		r.SucursalDondeTrabaja = nil
		return
	case 5:
		r.EsEventual = nil
		return
	case 6:
		r.Login = nil
		return
	case 7:
		r.Idgla = nil
		return
	case 8:
		r.Cuit = nil
		return
	}
	panic("Unknown field index")
}

func (r *DatosDistribuidor) NullField(i int) {
	switch i {
	case 1:
		r.MedioDeLocomocion = nil
		return
	case 2:
		r.TipoDeDistribuidor = nil
		return
	case 3:
		r.Legajo = nil
		return
	case 4:
		r.SucursalDondeTrabaja = nil
		return
	case 5:
		r.EsEventual = nil
		return
	case 6:
		r.Login = nil
		return
	case 7:
		r.Idgla = nil
		return
	case 8:
		r.Cuit = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ DatosDistribuidor) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DatosDistribuidor) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DatosDistribuidor) HintSize(int)                     { panic("Unsupported operation") }
func (_ DatosDistribuidor) Finalize()                        {}

func (_ DatosDistribuidor) AvroCRC64Fingerprint() []byte {
	return []byte(DatosDistribuidorAvroCRC64Fingerprint)
}

func (r DatosDistribuidor) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["datosPersonales"], err = json.Marshal(r.DatosPersonales)
	if err != nil {
		return nil, err
	}
	output["medioDeLocomocion"], err = json.Marshal(r.MedioDeLocomocion)
	if err != nil {
		return nil, err
	}
	output["tipoDeDistribuidor"], err = json.Marshal(r.TipoDeDistribuidor)
	if err != nil {
		return nil, err
	}
	output["legajo"], err = json.Marshal(r.Legajo)
	if err != nil {
		return nil, err
	}
	output["sucursalDondeTrabaja"], err = json.Marshal(r.SucursalDondeTrabaja)
	if err != nil {
		return nil, err
	}
	output["esEventual"], err = json.Marshal(r.EsEventual)
	if err != nil {
		return nil, err
	}
	output["login"], err = json.Marshal(r.Login)
	if err != nil {
		return nil, err
	}
	output["idgla"], err = json.Marshal(r.Idgla)
	if err != nil {
		return nil, err
	}
	output["cuit"], err = json.Marshal(r.Cuit)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DatosDistribuidor) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["datosPersonales"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DatosPersonales); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for datosPersonales")
	}
	val = func() json.RawMessage {
		if v, ok := fields["medioDeLocomocion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MedioDeLocomocion); err != nil {
			return err
		}
	} else {
		r.MedioDeLocomocion = NewUnionNullString()

		r.MedioDeLocomocion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoDeDistribuidor"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeDistribuidor); err != nil {
			return err
		}
	} else {
		r.TipoDeDistribuidor = NewUnionNullString()

		r.TipoDeDistribuidor = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["legajo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Legajo); err != nil {
			return err
		}
	} else {
		r.Legajo = NewUnionNullString()

		r.Legajo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["sucursalDondeTrabaja"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SucursalDondeTrabaja); err != nil {
			return err
		}
	} else {
		r.SucursalDondeTrabaja = NewUnionNullDatosSucursal()

		r.SucursalDondeTrabaja = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["esEventual"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EsEventual); err != nil {
			return err
		}
	} else {
		r.EsEventual = NewUnionNullBool()

		r.EsEventual = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["login"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Login); err != nil {
			return err
		}
	} else {
		r.Login = NewUnionNullString()

		r.Login = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["idgla"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Idgla); err != nil {
			return err
		}
	} else {
		r.Idgla = NewUnionNullString()

		r.Idgla = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["cuit"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cuit); err != nil {
			return err
		}
	} else {
		r.Cuit = NewUnionNullString()

		r.Cuit = nil
	}
	return nil
}