// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     GenerarPreenvios.avsc
 */
package PreEnvioBackendEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type DatosPersona struct {
	NombreCompleto *UnionNullString `json:"NombreCompleto"`

	Email *UnionNullString `json:"Email"`

	DocumentoTipo *UnionNullString `json:"DocumentoTipo"`

	DocumentoNumero *UnionNullString `json:"DocumentoNumero"`

	Telefonos *UnionNullArrayTelefonos `json:"Telefonos"`
}

const DatosPersonaAvroCRC64Fingerprint = "\x1a\xea\xb67\xf8u~\x01"

func NewDatosPersona() DatosPersona {
	r := DatosPersona{}
	r.NombreCompleto = nil
	r.Email = nil
	r.DocumentoTipo = nil
	r.DocumentoNumero = nil
	r.Telefonos = nil
	return r
}

func DeserializeDatosPersona(r io.Reader) (DatosPersona, error) {
	t := NewDatosPersona()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDatosPersonaFromSchema(r io.Reader, schema string) (DatosPersona, error) {
	t := NewDatosPersona()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDatosPersona(r DatosPersona, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.NombreCompleto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Email, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DocumentoTipo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DocumentoNumero, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayTelefonos(r.Telefonos, w)
	if err != nil {
		return err
	}
	return err
}

func (r DatosPersona) Serialize(w io.Writer) error {
	return writeDatosPersona(r, w)
}

func (r DatosPersona) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"NombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Email\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DocumentoTipo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DocumentoNumero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Telefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"Tipo\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"Numero\",\"type\":[\"null\",\"string\"]}],\"name\":\"Telefonos\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"Andreani.PreEnvioBackend.Events.Preenvio.Common.DatosPersona\",\"type\":\"record\"}"
}

func (r DatosPersona) SchemaName() string {
	return "Andreani.PreEnvioBackend.Events.Preenvio.Common.DatosPersona"
}

func (_ DatosPersona) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DatosPersona) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DatosPersona) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DatosPersona) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DatosPersona) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DatosPersona) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DatosPersona) SetString(v string)   { panic("Unsupported operation") }
func (_ DatosPersona) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DatosPersona) Get(i int) types.Field {
	switch i {
	case 0:
		r.NombreCompleto = NewUnionNullString()

		return r.NombreCompleto
	case 1:
		r.Email = NewUnionNullString()

		return r.Email
	case 2:
		r.DocumentoTipo = NewUnionNullString()

		return r.DocumentoTipo
	case 3:
		r.DocumentoNumero = NewUnionNullString()

		return r.DocumentoNumero
	case 4:
		r.Telefonos = NewUnionNullArrayTelefonos()

		return r.Telefonos
	}
	panic("Unknown field index")
}

func (r *DatosPersona) SetDefault(i int) {
	switch i {
	case 0:
		r.NombreCompleto = nil
		return
	case 1:
		r.Email = nil
		return
	case 2:
		r.DocumentoTipo = nil
		return
	case 3:
		r.DocumentoNumero = nil
		return
	case 4:
		r.Telefonos = nil
		return
	}
	panic("Unknown field index")
}

func (r *DatosPersona) NullField(i int) {
	switch i {
	case 0:
		r.NombreCompleto = nil
		return
	case 1:
		r.Email = nil
		return
	case 2:
		r.DocumentoTipo = nil
		return
	case 3:
		r.DocumentoNumero = nil
		return
	case 4:
		r.Telefonos = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ DatosPersona) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DatosPersona) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DatosPersona) HintSize(int)                     { panic("Unsupported operation") }
func (_ DatosPersona) Finalize()                        {}

func (_ DatosPersona) AvroCRC64Fingerprint() []byte {
	return []byte(DatosPersonaAvroCRC64Fingerprint)
}

func (r DatosPersona) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["NombreCompleto"], err = json.Marshal(r.NombreCompleto)
	if err != nil {
		return nil, err
	}
	output["Email"], err = json.Marshal(r.Email)
	if err != nil {
		return nil, err
	}
	output["DocumentoTipo"], err = json.Marshal(r.DocumentoTipo)
	if err != nil {
		return nil, err
	}
	output["DocumentoNumero"], err = json.Marshal(r.DocumentoNumero)
	if err != nil {
		return nil, err
	}
	output["Telefonos"], err = json.Marshal(r.Telefonos)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DatosPersona) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["NombreCompleto"]; ok {
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
		if v, ok := fields["Email"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Email); err != nil {
			return err
		}
	} else {
		r.Email = NewUnionNullString()

		r.Email = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DocumentoTipo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DocumentoTipo); err != nil {
			return err
		}
	} else {
		r.DocumentoTipo = NewUnionNullString()

		r.DocumentoTipo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DocumentoNumero"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DocumentoNumero); err != nil {
			return err
		}
	} else {
		r.DocumentoNumero = NewUnionNullString()

		r.DocumentoNumero = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Telefonos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Telefonos); err != nil {
			return err
		}
	} else {
		r.Telefonos = NewUnionNullArrayTelefonos()

		r.Telefonos = nil
	}
	return nil
}