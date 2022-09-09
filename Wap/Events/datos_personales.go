// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CreacionDeOrdenDeCompraRechazada.avsc
 */
package Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type DatosPersonales struct {
	NumeroDeDocumento *UnionNullString `json:"numeroDeDocumento"`

	NombreCompleto *UnionNullString `json:"nombreCompleto"`

	IdInternoDelCliente *UnionNullString `json:"idInternoDelCliente"`

	EMail *UnionNullString `json:"eMail"`

	Telefonos *UnionNullListaDeTelefonos `json:"telefonos"`

	Agrupador *UnionNullString `json:"agrupador"`

	TipoDeDocumento TipoDeDocumento `json:"tipoDeDocumento"`
}

const DatosPersonalesAvroCRC64Fingerprint = "\xd3\xedHkI\xa3\xca\x06"

func NewDatosPersonales() DatosPersonales {
	r := DatosPersonales{}
	r.NumeroDeDocumento = nil
	r.NombreCompleto = nil
	r.IdInternoDelCliente = nil
	r.EMail = nil
	r.Telefonos = nil
	r.Agrupador = nil
	return r
}

func DeserializeDatosPersonales(r io.Reader) (DatosPersonales, error) {
	t := NewDatosPersonales()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDatosPersonalesFromSchema(r io.Reader, schema string) (DatosPersonales, error) {
	t := NewDatosPersonales()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDatosPersonales(r DatosPersonales, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.NumeroDeDocumento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NombreCompleto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.IdInternoDelCliente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EMail, w)
	if err != nil {
		return err
	}
	err = writeUnionNullListaDeTelefonos(r.Telefonos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Agrupador, w)
	if err != nil {
		return err
	}
	err = writeTipoDeDocumento(r.TipoDeDocumento, w)
	if err != nil {
		return err
	}
	return err
}

func (r DatosPersonales) Serialize(w io.Writer) error {
	return writeDatosPersonales(r, w)
}

func (r DatosPersonales) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idInternoDelCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"eMail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"tipo\",\"type\":\"string\"},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDeTelefonos\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"agrupador\",\"type\":[\"null\",\"string\"]},{\"name\":\"tipoDeDocumento\",\"type\":{\"name\":\"TipoDeDocumento\",\"symbols\":[\"undefined\",\"DNI\",\"CUIT\",\"CUIL\"],\"type\":\"enum\"}}],\"name\":\"Andreani.Wap.DatosDeReferencia.DatosPersonales\",\"type\":\"record\"}"
}

func (r DatosPersonales) SchemaName() string {
	return "Andreani.Wap.DatosDeReferencia.DatosPersonales"
}

func (_ DatosPersonales) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DatosPersonales) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DatosPersonales) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DatosPersonales) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DatosPersonales) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DatosPersonales) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DatosPersonales) SetString(v string)   { panic("Unsupported operation") }
func (_ DatosPersonales) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DatosPersonales) Get(i int) types.Field {
	switch i {
	case 0:
		r.NumeroDeDocumento = NewUnionNullString()

		return r.NumeroDeDocumento
	case 1:
		r.NombreCompleto = NewUnionNullString()

		return r.NombreCompleto
	case 2:
		r.IdInternoDelCliente = NewUnionNullString()

		return r.IdInternoDelCliente
	case 3:
		r.EMail = NewUnionNullString()

		return r.EMail
	case 4:
		r.Telefonos = NewUnionNullListaDeTelefonos()

		return r.Telefonos
	case 5:
		r.Agrupador = NewUnionNullString()

		return r.Agrupador
	case 6:
		w := TipoDeDocumentoWrapper{Target: &r.TipoDeDocumento}

		return w

	}
	panic("Unknown field index")
}

func (r *DatosPersonales) SetDefault(i int) {
	switch i {
	case 0:
		r.NumeroDeDocumento = nil
		return
	case 1:
		r.NombreCompleto = nil
		return
	case 2:
		r.IdInternoDelCliente = nil
		return
	case 3:
		r.EMail = nil
		return
	case 4:
		r.Telefonos = nil
		return
	case 5:
		r.Agrupador = nil
		return
	}
	panic("Unknown field index")
}

func (r *DatosPersonales) NullField(i int) {
	switch i {
	case 0:
		r.NumeroDeDocumento = nil
		return
	case 1:
		r.NombreCompleto = nil
		return
	case 2:
		r.IdInternoDelCliente = nil
		return
	case 3:
		r.EMail = nil
		return
	case 4:
		r.Telefonos = nil
		return
	case 5:
		r.Agrupador = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ DatosPersonales) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DatosPersonales) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DatosPersonales) HintSize(int)                     { panic("Unsupported operation") }
func (_ DatosPersonales) Finalize()                        {}

func (_ DatosPersonales) AvroCRC64Fingerprint() []byte {
	return []byte(DatosPersonalesAvroCRC64Fingerprint)
}

func (r DatosPersonales) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["numeroDeDocumento"], err = json.Marshal(r.NumeroDeDocumento)
	if err != nil {
		return nil, err
	}
	output["nombreCompleto"], err = json.Marshal(r.NombreCompleto)
	if err != nil {
		return nil, err
	}
	output["idInternoDelCliente"], err = json.Marshal(r.IdInternoDelCliente)
	if err != nil {
		return nil, err
	}
	output["eMail"], err = json.Marshal(r.EMail)
	if err != nil {
		return nil, err
	}
	output["telefonos"], err = json.Marshal(r.Telefonos)
	if err != nil {
		return nil, err
	}
	output["agrupador"], err = json.Marshal(r.Agrupador)
	if err != nil {
		return nil, err
	}
	output["tipoDeDocumento"], err = json.Marshal(r.TipoDeDocumento)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DatosPersonales) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["numeroDeDocumento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeDocumento); err != nil {
			return err
		}
	} else {
		r.NumeroDeDocumento = NewUnionNullString()

		r.NumeroDeDocumento = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["nombreCompleto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NombreCompleto); err != nil {
			return err
		}
	} else {
		r.NombreCompleto = NewUnionNullString()

		r.NombreCompleto = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["idInternoDelCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdInternoDelCliente); err != nil {
			return err
		}
	} else {
		r.IdInternoDelCliente = NewUnionNullString()

		r.IdInternoDelCliente = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["eMail"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EMail); err != nil {
			return err
		}
	} else {
		r.EMail = NewUnionNullString()

		r.EMail = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["telefonos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Telefonos); err != nil {
			return err
		}
	} else {
		r.Telefonos = NewUnionNullListaDeTelefonos()

		r.Telefonos = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["agrupador"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Agrupador); err != nil {
			return err
		}
	} else {
		r.Agrupador = NewUnionNullString()

		r.Agrupador = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoDeDocumento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeDocumento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tipoDeDocumento")
	}
	return nil
}
