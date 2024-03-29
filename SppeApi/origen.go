// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Origen.avsc
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

type Origen struct {
	OrigenCiudad *UnionNullString `json:"OrigenCiudad"`

	OrigenCodigoPostal *UnionNullString `json:"OrigenCodigoPostal"`

	OrigenCalle *UnionNullString `json:"OrigenCalle"`

	OrigenNumero *UnionNullString `json:"OrigenNumero"`

	OrigenPiso *UnionNullString `json:"OrigenPiso"`

	OrigenDepartamento *UnionNullString `json:"OrigenDepartamento"`

	OrigenEmail *UnionNullString `json:"OrigenEmail"`

	OrigenTelefono *UnionNullString `json:"OrigenTelefono"`

	OrigenRegion *UnionNullString `json:"OrigenRegion"`
}

const OrigenAvroCRC64Fingerprint = "\xbby\x19\x1c\"2I\x19"

func NewOrigen() Origen {
	r := Origen{}
	return r
}

func DeserializeOrigen(r io.Reader) (Origen, error) {
	t := NewOrigen()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeOrigenFromSchema(r io.Reader, schema string) (Origen, error) {
	t := NewOrigen()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeOrigen(r Origen, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.OrigenCiudad, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.OrigenCodigoPostal, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.OrigenCalle, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.OrigenNumero, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.OrigenPiso, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.OrigenDepartamento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.OrigenEmail, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.OrigenTelefono, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.OrigenRegion, w)
	if err != nil {
		return err
	}
	return err
}

func (r Origen) Serialize(w io.Writer) error {
	return writeOrigen(r, w)
}

func (r Origen) Schema() string {
	return "{\"fields\":[{\"name\":\"OrigenCiudad\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenCodigoPostal\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenCalle\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenNumero\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenPiso\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenDepartamento\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenEmail\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenTelefono\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenRegion\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.SppeApi.Events.OrdenDeRetiroRechazadaCommon.Origen\",\"type\":\"record\"}"
}

func (r Origen) SchemaName() string {
	return "Andreani.SppeApi.Events.OrdenDeRetiroRechazadaCommon.Origen"
}

func (_ Origen) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Origen) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Origen) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Origen) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Origen) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Origen) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Origen) SetString(v string)   { panic("Unsupported operation") }
func (_ Origen) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Origen) Get(i int) types.Field {
	switch i {
	case 0:
		r.OrigenCiudad = NewUnionNullString()

		return r.OrigenCiudad
	case 1:
		r.OrigenCodigoPostal = NewUnionNullString()

		return r.OrigenCodigoPostal
	case 2:
		r.OrigenCalle = NewUnionNullString()

		return r.OrigenCalle
	case 3:
		r.OrigenNumero = NewUnionNullString()

		return r.OrigenNumero
	case 4:
		r.OrigenPiso = NewUnionNullString()

		return r.OrigenPiso
	case 5:
		r.OrigenDepartamento = NewUnionNullString()

		return r.OrigenDepartamento
	case 6:
		r.OrigenEmail = NewUnionNullString()

		return r.OrigenEmail
	case 7:
		r.OrigenTelefono = NewUnionNullString()

		return r.OrigenTelefono
	case 8:
		r.OrigenRegion = NewUnionNullString()

		return r.OrigenRegion
	}
	panic("Unknown field index")
}

func (r *Origen) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Origen) NullField(i int) {
	switch i {
	case 0:
		r.OrigenCiudad = nil
		return
	case 1:
		r.OrigenCodigoPostal = nil
		return
	case 2:
		r.OrigenCalle = nil
		return
	case 3:
		r.OrigenNumero = nil
		return
	case 4:
		r.OrigenPiso = nil
		return
	case 5:
		r.OrigenDepartamento = nil
		return
	case 6:
		r.OrigenEmail = nil
		return
	case 7:
		r.OrigenTelefono = nil
		return
	case 8:
		r.OrigenRegion = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Origen) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Origen) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Origen) HintSize(int)                     { panic("Unsupported operation") }
func (_ Origen) Finalize()                        {}

func (_ Origen) AvroCRC64Fingerprint() []byte {
	return []byte(OrigenAvroCRC64Fingerprint)
}

func (r Origen) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["OrigenCiudad"], err = json.Marshal(r.OrigenCiudad)
	if err != nil {
		return nil, err
	}
	output["OrigenCodigoPostal"], err = json.Marshal(r.OrigenCodigoPostal)
	if err != nil {
		return nil, err
	}
	output["OrigenCalle"], err = json.Marshal(r.OrigenCalle)
	if err != nil {
		return nil, err
	}
	output["OrigenNumero"], err = json.Marshal(r.OrigenNumero)
	if err != nil {
		return nil, err
	}
	output["OrigenPiso"], err = json.Marshal(r.OrigenPiso)
	if err != nil {
		return nil, err
	}
	output["OrigenDepartamento"], err = json.Marshal(r.OrigenDepartamento)
	if err != nil {
		return nil, err
	}
	output["OrigenEmail"], err = json.Marshal(r.OrigenEmail)
	if err != nil {
		return nil, err
	}
	output["OrigenTelefono"], err = json.Marshal(r.OrigenTelefono)
	if err != nil {
		return nil, err
	}
	output["OrigenRegion"], err = json.Marshal(r.OrigenRegion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Origen) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["OrigenCiudad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrigenCiudad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrigenCiudad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OrigenCodigoPostal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrigenCodigoPostal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrigenCodigoPostal")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OrigenCalle"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrigenCalle); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrigenCalle")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OrigenNumero"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrigenNumero); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrigenNumero")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OrigenPiso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrigenPiso); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrigenPiso")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OrigenDepartamento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrigenDepartamento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrigenDepartamento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OrigenEmail"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrigenEmail); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrigenEmail")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OrigenTelefono"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrigenTelefono); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrigenTelefono")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OrigenRegion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrigenRegion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrigenRegion")
	}
	return nil
}
