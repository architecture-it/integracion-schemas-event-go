// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     OrdenDeRetiroRechazada.avsc
 */
package SppeApiEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Destino struct {
	DestinatarioCalle *UnionNullString `json:"DestinatarioCalle"`

	DestinarioNumero *UnionNullString `json:"DestinarioNumero"`

	DestinatarioPiso *UnionNullString `json:"DestinatarioPiso"`

	DestinatarioDepartamento *UnionNullString `json:"DestinatarioDepartamento"`

	DestinatarioGLNDNI *UnionNullString `json:"DestinatarioGLNDNI"`

	DestinatarioCiudad *UnionNullString `json:"DestinatarioCiudad"`

	DestinatarioProvincia *UnionNullString `json:"DestinatarioProvincia"`

	DestinatarioCodigoPostal *UnionNullString `json:"DestinatarioCodigoPostal"`

	DestinatarioTelefono *UnionNullString `json:"DestinatarioTelefono"`

	DestinatarioEmail *UnionNullString `json:"DestinatarioEmail"`
}

const DestinoAvroCRC64Fingerprint = ",s\xc99M\xe3c\x8a"

func NewDestino() Destino {
	r := Destino{}
	return r
}

func DeserializeDestino(r io.Reader) (Destino, error) {
	t := NewDestino()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDestinoFromSchema(r io.Reader, schema string) (Destino, error) {
	t := NewDestino()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDestino(r Destino, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.DestinatarioCalle, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinarioNumero, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioPiso, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioDepartamento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioGLNDNI, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioCiudad, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioProvincia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioCodigoPostal, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioTelefono, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioEmail, w)
	if err != nil {
		return err
	}
	return err
}

func (r Destino) Serialize(w io.Writer) error {
	return writeDestino(r, w)
}

func (r Destino) Schema() string {
	return "{\"fields\":[{\"name\":\"DestinatarioCalle\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinarioNumero\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioPiso\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioDepartamento\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioGLNDNI\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioCiudad\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioProvincia\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioCodigoPostal\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioTelefono\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioEmail\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.SppeApi.Events.OrdenDeRetiroRechazadaCommon.Destino\",\"type\":\"record\"}"
}

func (r Destino) SchemaName() string {
	return "Andreani.SppeApi.Events.OrdenDeRetiroRechazadaCommon.Destino"
}

func (_ Destino) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Destino) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Destino) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Destino) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Destino) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Destino) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Destino) SetString(v string)   { panic("Unsupported operation") }
func (_ Destino) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Destino) Get(i int) types.Field {
	switch i {
	case 0:
		r.DestinatarioCalle = NewUnionNullString()

		return r.DestinatarioCalle
	case 1:
		r.DestinarioNumero = NewUnionNullString()

		return r.DestinarioNumero
	case 2:
		r.DestinatarioPiso = NewUnionNullString()

		return r.DestinatarioPiso
	case 3:
		r.DestinatarioDepartamento = NewUnionNullString()

		return r.DestinatarioDepartamento
	case 4:
		r.DestinatarioGLNDNI = NewUnionNullString()

		return r.DestinatarioGLNDNI
	case 5:
		r.DestinatarioCiudad = NewUnionNullString()

		return r.DestinatarioCiudad
	case 6:
		r.DestinatarioProvincia = NewUnionNullString()

		return r.DestinatarioProvincia
	case 7:
		r.DestinatarioCodigoPostal = NewUnionNullString()

		return r.DestinatarioCodigoPostal
	case 8:
		r.DestinatarioTelefono = NewUnionNullString()

		return r.DestinatarioTelefono
	case 9:
		r.DestinatarioEmail = NewUnionNullString()

		return r.DestinatarioEmail
	}
	panic("Unknown field index")
}

func (r *Destino) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Destino) NullField(i int) {
	switch i {
	case 0:
		r.DestinatarioCalle = nil
		return
	case 1:
		r.DestinarioNumero = nil
		return
	case 2:
		r.DestinatarioPiso = nil
		return
	case 3:
		r.DestinatarioDepartamento = nil
		return
	case 4:
		r.DestinatarioGLNDNI = nil
		return
	case 5:
		r.DestinatarioCiudad = nil
		return
	case 6:
		r.DestinatarioProvincia = nil
		return
	case 7:
		r.DestinatarioCodigoPostal = nil
		return
	case 8:
		r.DestinatarioTelefono = nil
		return
	case 9:
		r.DestinatarioEmail = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Destino) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Destino) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Destino) HintSize(int)                     { panic("Unsupported operation") }
func (_ Destino) Finalize()                        {}

func (_ Destino) AvroCRC64Fingerprint() []byte {
	return []byte(DestinoAvroCRC64Fingerprint)
}

func (r Destino) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["DestinatarioCalle"], err = json.Marshal(r.DestinatarioCalle)
	if err != nil {
		return nil, err
	}
	output["DestinarioNumero"], err = json.Marshal(r.DestinarioNumero)
	if err != nil {
		return nil, err
	}
	output["DestinatarioPiso"], err = json.Marshal(r.DestinatarioPiso)
	if err != nil {
		return nil, err
	}
	output["DestinatarioDepartamento"], err = json.Marshal(r.DestinatarioDepartamento)
	if err != nil {
		return nil, err
	}
	output["DestinatarioGLNDNI"], err = json.Marshal(r.DestinatarioGLNDNI)
	if err != nil {
		return nil, err
	}
	output["DestinatarioCiudad"], err = json.Marshal(r.DestinatarioCiudad)
	if err != nil {
		return nil, err
	}
	output["DestinatarioProvincia"], err = json.Marshal(r.DestinatarioProvincia)
	if err != nil {
		return nil, err
	}
	output["DestinatarioCodigoPostal"], err = json.Marshal(r.DestinatarioCodigoPostal)
	if err != nil {
		return nil, err
	}
	output["DestinatarioTelefono"], err = json.Marshal(r.DestinatarioTelefono)
	if err != nil {
		return nil, err
	}
	output["DestinatarioEmail"], err = json.Marshal(r.DestinatarioEmail)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Destino) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioCalle"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioCalle); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioCalle")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinarioNumero"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinarioNumero); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinarioNumero")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioPiso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioPiso); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioPiso")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioDepartamento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioDepartamento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioDepartamento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioGLNDNI"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioGLNDNI); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioGLNDNI")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioCiudad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioCiudad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioCiudad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioProvincia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioProvincia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioProvincia")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioCodigoPostal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioCodigoPostal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioCodigoPostal")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioTelefono"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioTelefono); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioTelefono")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioEmail"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioEmail); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioEmail")
	}
	return nil
}
