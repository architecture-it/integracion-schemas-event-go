// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Origen.avsc
 */
package Wapv2Events

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
	IdOrigen string `json:"idOrigen"`

	Calle *UnionNullString `json:"calle"`

	CodigoPostal *UnionNullString `json:"codigoPostal"`

	NombreProvincia *UnionNullString `json:"nombreProvincia"`

	Localidad *UnionNullString `json:"localidad"`

	Piso *UnionNullString `json:"piso"`

	Departamento *UnionNullString `json:"departamento"`

	Contacto *UnionNullString `json:"contacto"`

	DatosAdicionales *UnionNullListaDePropiedades `json:"datosAdicionales"`
}

const OrigenAvroCRC64Fingerprint = "陣\xc0@N\x95'"

func NewOrigen() Origen {
	r := Origen{}
	r.Calle = nil
	r.CodigoPostal = nil
	r.NombreProvincia = nil
	r.Localidad = nil
	r.Piso = nil
	r.Departamento = nil
	r.Contacto = nil
	r.DatosAdicionales = nil
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
	err = vm.WriteString(r.IdOrigen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Calle, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoPostal, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NombreProvincia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Localidad, w)
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
	err = writeUnionNullString(r.Contacto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullListaDePropiedades(r.DatosAdicionales, w)
	if err != nil {
		return err
	}
	return err
}

func (r Origen) Serialize(w io.Writer) error {
	return writeOrigen(r, w)
}

func (r Origen) Schema() string {
	return "{\"fields\":[{\"name\":\"idOrigen\",\"type\":\"string\"},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contacto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"datosAdicionales\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]}],\"name\":\"Andreani.Wapv2.Events.Record.Origen\",\"type\":\"record\"}"
}

func (r Origen) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.Origen"
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
		w := types.String{Target: &r.IdOrigen}

		return w

	case 1:
		r.Calle = NewUnionNullString()

		return r.Calle
	case 2:
		r.CodigoPostal = NewUnionNullString()

		return r.CodigoPostal
	case 3:
		r.NombreProvincia = NewUnionNullString()

		return r.NombreProvincia
	case 4:
		r.Localidad = NewUnionNullString()

		return r.Localidad
	case 5:
		r.Piso = NewUnionNullString()

		return r.Piso
	case 6:
		r.Departamento = NewUnionNullString()

		return r.Departamento
	case 7:
		r.Contacto = NewUnionNullString()

		return r.Contacto
	case 8:
		r.DatosAdicionales = NewUnionNullListaDePropiedades()

		return r.DatosAdicionales
	}
	panic("Unknown field index")
}

func (r *Origen) SetDefault(i int) {
	switch i {
	case 1:
		r.Calle = nil
		return
	case 2:
		r.CodigoPostal = nil
		return
	case 3:
		r.NombreProvincia = nil
		return
	case 4:
		r.Localidad = nil
		return
	case 5:
		r.Piso = nil
		return
	case 6:
		r.Departamento = nil
		return
	case 7:
		r.Contacto = nil
		return
	case 8:
		r.DatosAdicionales = nil
		return
	}
	panic("Unknown field index")
}

func (r *Origen) NullField(i int) {
	switch i {
	case 1:
		r.Calle = nil
		return
	case 2:
		r.CodigoPostal = nil
		return
	case 3:
		r.NombreProvincia = nil
		return
	case 4:
		r.Localidad = nil
		return
	case 5:
		r.Piso = nil
		return
	case 6:
		r.Departamento = nil
		return
	case 7:
		r.Contacto = nil
		return
	case 8:
		r.DatosAdicionales = nil
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
	output["idOrigen"], err = json.Marshal(r.IdOrigen)
	if err != nil {
		return nil, err
	}
	output["calle"], err = json.Marshal(r.Calle)
	if err != nil {
		return nil, err
	}
	output["codigoPostal"], err = json.Marshal(r.CodigoPostal)
	if err != nil {
		return nil, err
	}
	output["nombreProvincia"], err = json.Marshal(r.NombreProvincia)
	if err != nil {
		return nil, err
	}
	output["localidad"], err = json.Marshal(r.Localidad)
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
	output["contacto"], err = json.Marshal(r.Contacto)
	if err != nil {
		return nil, err
	}
	output["datosAdicionales"], err = json.Marshal(r.DatosAdicionales)
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
		if v, ok := fields["idOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdOrigen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idOrigen")
	}
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
		r.Calle = NewUnionNullString()

		r.Calle = nil
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
		r.CodigoPostal = NewUnionNullString()

		r.CodigoPostal = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["nombreProvincia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NombreProvincia); err != nil {
			return err
		}
	} else {
		r.NombreProvincia = NewUnionNullString()

		r.NombreProvincia = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["localidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Localidad); err != nil {
			return err
		}
	} else {
		r.Localidad = NewUnionNullString()

		r.Localidad = nil
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
		r.Piso = NewUnionNullString()

		r.Piso = nil
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
		r.Departamento = NewUnionNullString()

		r.Departamento = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["contacto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contacto); err != nil {
			return err
		}
	} else {
		r.Contacto = NewUnionNullString()

		r.Contacto = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["datosAdicionales"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DatosAdicionales); err != nil {
			return err
		}
	} else {
		r.DatosAdicionales = NewUnionNullListaDePropiedades()

		r.DatosAdicionales = nil
	}
	return nil
}
