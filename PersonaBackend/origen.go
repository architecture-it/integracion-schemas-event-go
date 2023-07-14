// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Origen.avsc
 */
package PersonaBackendEvents

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
	Tipo string `json:"Tipo"`

	Calle string `json:"Calle"`

	Numero *UnionNullString `json:"Numero"`

	Piso *UnionNullString `json:"Piso"`

	Unidad *UnionNullString `json:"Unidad"`

	Localidad string `json:"Localidad"`

	CodigoPostal string `json:"CodigoPostal"`

	Provincia string `json:"Provincia"`

	SucursalId *UnionNullString `json:"SucursalId"`
}

const OrigenAvroCRC64Fingerprint = "/FF\x18H\x7f\xb1\xd0"

func NewOrigen() Origen {
	r := Origen{}
	r.Numero = nil
	r.Piso = nil
	r.Unidad = nil
	r.SucursalId = nil
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
	err = vm.WriteString(r.Tipo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Calle, w)
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
	err = writeUnionNullString(r.Unidad, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Localidad, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoPostal, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Provincia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.SucursalId, w)
	if err != nil {
		return err
	}
	return err
}

func (r Origen) Serialize(w io.Writer) error {
	return writeOrigen(r, w)
}

func (r Origen) Schema() string {
	return "{\"fields\":[{\"name\":\"Tipo\",\"type\":\"string\"},{\"name\":\"Calle\",\"type\":\"string\"},{\"default\":null,\"name\":\"Numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Unidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"CodigoPostal\",\"type\":\"string\"},{\"name\":\"Provincia\",\"type\":\"string\"},{\"default\":null,\"name\":\"SucursalId\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.PersonaBackend.Events.Common.Origen\",\"type\":\"record\"}"
}

func (r Origen) SchemaName() string {
	return "Andreani.PersonaBackend.Events.Common.Origen"
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
		w := types.String{Target: &r.Tipo}

		return w

	case 1:
		w := types.String{Target: &r.Calle}

		return w

	case 2:
		r.Numero = NewUnionNullString()

		return r.Numero
	case 3:
		r.Piso = NewUnionNullString()

		return r.Piso
	case 4:
		r.Unidad = NewUnionNullString()

		return r.Unidad
	case 5:
		w := types.String{Target: &r.Localidad}

		return w

	case 6:
		w := types.String{Target: &r.CodigoPostal}

		return w

	case 7:
		w := types.String{Target: &r.Provincia}

		return w

	case 8:
		r.SucursalId = NewUnionNullString()

		return r.SucursalId
	}
	panic("Unknown field index")
}

func (r *Origen) SetDefault(i int) {
	switch i {
	case 2:
		r.Numero = nil
		return
	case 3:
		r.Piso = nil
		return
	case 4:
		r.Unidad = nil
		return
	case 8:
		r.SucursalId = nil
		return
	}
	panic("Unknown field index")
}

func (r *Origen) NullField(i int) {
	switch i {
	case 2:
		r.Numero = nil
		return
	case 3:
		r.Piso = nil
		return
	case 4:
		r.Unidad = nil
		return
	case 8:
		r.SucursalId = nil
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
	output["Tipo"], err = json.Marshal(r.Tipo)
	if err != nil {
		return nil, err
	}
	output["Calle"], err = json.Marshal(r.Calle)
	if err != nil {
		return nil, err
	}
	output["Numero"], err = json.Marshal(r.Numero)
	if err != nil {
		return nil, err
	}
	output["Piso"], err = json.Marshal(r.Piso)
	if err != nil {
		return nil, err
	}
	output["Unidad"], err = json.Marshal(r.Unidad)
	if err != nil {
		return nil, err
	}
	output["Localidad"], err = json.Marshal(r.Localidad)
	if err != nil {
		return nil, err
	}
	output["CodigoPostal"], err = json.Marshal(r.CodigoPostal)
	if err != nil {
		return nil, err
	}
	output["Provincia"], err = json.Marshal(r.Provincia)
	if err != nil {
		return nil, err
	}
	output["SucursalId"], err = json.Marshal(r.SucursalId)
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
		if v, ok := fields["Tipo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Tipo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Tipo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Calle"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Calle); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Calle")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Numero"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Numero); err != nil {
			return err
		}
	} else {
		r.Numero = NewUnionNullString()

		r.Numero = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Piso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Piso); err != nil {
			return err
		}
	} else {
		r.Piso = NewUnionNullString()

		r.Piso = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Unidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Unidad); err != nil {
			return err
		}
	} else {
		r.Unidad = NewUnionNullString()

		r.Unidad = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Localidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Localidad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Localidad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoPostal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoPostal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodigoPostal")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Provincia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Provincia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Provincia")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SucursalId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SucursalId); err != nil {
			return err
		}
	} else {
		r.SucursalId = NewUnionNullString()

		r.SucursalId = nil
	}
	return nil
}
