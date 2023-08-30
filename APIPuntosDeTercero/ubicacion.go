// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Ubicacion.avsc
 */
package APIPuntosDeTerceroEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Ubicacion struct {
	Calle *UnionNullString `json:"calle"`

	Numero *UnionNullString `json:"numero"`

	Piso *UnionNullString `json:"piso"`

	Departamento *UnionNullString `json:"departamento"`

	CodigoPostal *UnionNullString `json:"codigoPostal"`

	Latitud *UnionNullString `json:"latitud"`

	Longitud *UnionNullString `json:"longitud"`
}

const UbicacionAvroCRC64Fingerprint = "\xf8\x9a\xfdŔ\xe9b\xa6"

func NewUbicacion() Ubicacion {
	r := Ubicacion{}
	return r
}

func DeserializeUbicacion(r io.Reader) (Ubicacion, error) {
	t := NewUbicacion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeUbicacionFromSchema(r io.Reader, schema string) (Ubicacion, error) {
	t := NewUbicacion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeUbicacion(r Ubicacion, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Calle, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Numero, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Piso, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Departamento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoPostal, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Latitud, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Longitud, w)
	if err != nil {
		return err
	}
	return err
}

func (r Ubicacion) Serialize(w io.Writer) error {
	return writeUbicacion(r, w)
}

func (r Ubicacion) Schema() string {
	return "{\"fields\":[{\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"name\":\"latitud\",\"type\":[\"null\",\"string\"]},{\"name\":\"longitud\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.PuntoDeTercero.Events.Common.Ubicacion\",\"type\":\"record\"}"
}

func (r Ubicacion) SchemaName() string {
	return "Andreani.PuntoDeTercero.Events.Common.Ubicacion"
}

func (_ Ubicacion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Ubicacion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Ubicacion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Ubicacion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Ubicacion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Ubicacion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Ubicacion) SetString(v string)   { panic("Unsupported operation") }
func (_ Ubicacion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Ubicacion) Get(i int) types.Field {
	switch i {
	case 0:
		r.Calle = NewUnionNullString()

		return r.Calle
	case 1:
		r.Numero = NewUnionNullString()

		return r.Numero
	case 2:
		r.Piso = NewUnionNullString()

		return r.Piso
	case 3:
		r.Departamento = NewUnionNullString()

		return r.Departamento
	case 4:
		r.CodigoPostal = NewUnionNullString()

		return r.CodigoPostal
	case 5:
		r.Latitud = NewUnionNullString()

		return r.Latitud
	case 6:
		r.Longitud = NewUnionNullString()

		return r.Longitud
	}
	panic("Unknown field index")
}

func (r *Ubicacion) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Ubicacion) NullField(i int) {
	switch i {
	case 0:
		r.Calle = nil
		return
	case 1:
		r.Numero = nil
		return
	case 2:
		r.Piso = nil
		return
	case 3:
		r.Departamento = nil
		return
	case 4:
		r.CodigoPostal = nil
		return
	case 5:
		r.Latitud = nil
		return
	case 6:
		r.Longitud = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Ubicacion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Ubicacion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Ubicacion) HintSize(int)                     { panic("Unsupported operation") }
func (_ Ubicacion) Finalize()                        {}

func (_ Ubicacion) AvroCRC64Fingerprint() []byte {
	return []byte(UbicacionAvroCRC64Fingerprint)
}

func (r Ubicacion) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["calle"], err = json.Marshal(r.Calle)
	if err != nil {
		return nil, err
	}
	output["numero"], err = json.Marshal(r.Numero)
	if err != nil {
		return nil, err
	}
	output["piso"], err = json.Marshal(r.Piso)
	if err != nil {
		return nil, err
	}
	output["departamento"], err = json.Marshal(r.Departamento)
	if err != nil {
		return nil, err
	}
	output["codigoPostal"], err = json.Marshal(r.CodigoPostal)
	if err != nil {
		return nil, err
	}
	output["latitud"], err = json.Marshal(r.Latitud)
	if err != nil {
		return nil, err
	}
	output["longitud"], err = json.Marshal(r.Longitud)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Ubicacion) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["calle"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Calle); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for calle")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numero"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Numero); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numero")
	}
	val = func() json.RawMessage {
		if v, ok := fields["piso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Piso); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for piso")
	}
	val = func() json.RawMessage {
		if v, ok := fields["departamento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Departamento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for departamento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoPostal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoPostal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigoPostal")
	}
	val = func() json.RawMessage {
		if v, ok := fields["latitud"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Latitud); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for latitud")
	}
	val = func() json.RawMessage {
		if v, ok := fields["longitud"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Longitud); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for longitud")
	}
	return nil
}
