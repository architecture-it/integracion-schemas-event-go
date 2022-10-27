// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     AltaDePedidosSolicitada.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Direccion struct {
	AbreviaturaProvincia *UnionNullString `json:"abreviaturaProvincia"`

	Calle *UnionNullString `json:"calle"`

	CodigoDeDireccion *UnionNullString `json:"codigoDeDireccion"`

	CodigoCiudad *UnionNullString `json:"codigoCiudad"`

	CodigoPostal *UnionNullString `json:"codigoPostal"`

	NombreProvincia *UnionNullString `json:"nombreProvincia"`

	Numero *UnionNullString `json:"numero"`

	Telefono *UnionNullString `json:"telefono"`

	CodigoISOProvincia *UnionNullString `json:"codigoISOProvincia"`

	CodigoISOPais *UnionNullString `json:"codigoISOPais"`

	Localidad *UnionNullString `json:"localidad"`

	ComponentesDeDireccion *UnionNullListaDePropiedades `json:"componentesDeDireccion"`
}

const DireccionAvroCRC64Fingerprint = "\x8c\xa8mOޚ\x00j"

func NewDireccion() Direccion {
	r := Direccion{}
	r.AbreviaturaProvincia = nil
	r.Calle = nil
	r.CodigoDeDireccion = nil
	r.CodigoCiudad = nil
	r.CodigoPostal = nil
	r.NombreProvincia = nil
	r.Numero = nil
	r.Telefono = nil
	r.CodigoISOProvincia = nil
	r.CodigoISOPais = nil
	r.Localidad = nil
	r.ComponentesDeDireccion = nil
	return r
}

func DeserializeDireccion(r io.Reader) (Direccion, error) {
	t := NewDireccion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDireccionFromSchema(r io.Reader, schema string) (Direccion, error) {
	t := NewDireccion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDireccion(r Direccion, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.AbreviaturaProvincia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Calle, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoDeDireccion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoCiudad, w)
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
	err = writeUnionNullString(r.Numero, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Telefono, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoISOProvincia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoISOPais, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Localidad, w)
	if err != nil {
		return err
	}
	err = writeUnionNullListaDePropiedades(r.ComponentesDeDireccion, w)
	if err != nil {
		return err
	}
	return err
}

func (r Direccion) Serialize(w io.Writer) error {
	return writeDireccion(r, w)
}

func (r Direccion) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"abreviaturaProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoDeDireccion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoCiudad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefono\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoISOProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoISOPais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"namespace\":\"Andreani.Wap.Events.Record\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"namespace\":\"Andreani.Wap.Events.Record\",\"type\":\"record\"}]}],\"name\":\"Andreani.Wap.Events.Record.Direccion\",\"type\":\"record\"}"
}

func (r Direccion) SchemaName() string {
	return "Andreani.Wap.Events.Record.Direccion"
}

func (_ Direccion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Direccion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Direccion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Direccion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Direccion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Direccion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Direccion) SetString(v string)   { panic("Unsupported operation") }
func (_ Direccion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Direccion) Get(i int) types.Field {
	switch i {
	case 0:
		r.AbreviaturaProvincia = NewUnionNullString()

		return r.AbreviaturaProvincia
	case 1:
		r.Calle = NewUnionNullString()

		return r.Calle
	case 2:
		r.CodigoDeDireccion = NewUnionNullString()

		return r.CodigoDeDireccion
	case 3:
		r.CodigoCiudad = NewUnionNullString()

		return r.CodigoCiudad
	case 4:
		r.CodigoPostal = NewUnionNullString()

		return r.CodigoPostal
	case 5:
		r.NombreProvincia = NewUnionNullString()

		return r.NombreProvincia
	case 6:
		r.Numero = NewUnionNullString()

		return r.Numero
	case 7:
		r.Telefono = NewUnionNullString()

		return r.Telefono
	case 8:
		r.CodigoISOProvincia = NewUnionNullString()

		return r.CodigoISOProvincia
	case 9:
		r.CodigoISOPais = NewUnionNullString()

		return r.CodigoISOPais
	case 10:
		r.Localidad = NewUnionNullString()

		return r.Localidad
	case 11:
		r.ComponentesDeDireccion = NewUnionNullListaDePropiedades()

		return r.ComponentesDeDireccion
	}
	panic("Unknown field index")
}

func (r *Direccion) SetDefault(i int) {
	switch i {
	case 0:
		r.AbreviaturaProvincia = nil
		return
	case 1:
		r.Calle = nil
		return
	case 2:
		r.CodigoDeDireccion = nil
		return
	case 3:
		r.CodigoCiudad = nil
		return
	case 4:
		r.CodigoPostal = nil
		return
	case 5:
		r.NombreProvincia = nil
		return
	case 6:
		r.Numero = nil
		return
	case 7:
		r.Telefono = nil
		return
	case 8:
		r.CodigoISOProvincia = nil
		return
	case 9:
		r.CodigoISOPais = nil
		return
	case 10:
		r.Localidad = nil
		return
	case 11:
		r.ComponentesDeDireccion = nil
		return
	}
	panic("Unknown field index")
}

func (r *Direccion) NullField(i int) {
	switch i {
	case 0:
		r.AbreviaturaProvincia = nil
		return
	case 1:
		r.Calle = nil
		return
	case 2:
		r.CodigoDeDireccion = nil
		return
	case 3:
		r.CodigoCiudad = nil
		return
	case 4:
		r.CodigoPostal = nil
		return
	case 5:
		r.NombreProvincia = nil
		return
	case 6:
		r.Numero = nil
		return
	case 7:
		r.Telefono = nil
		return
	case 8:
		r.CodigoISOProvincia = nil
		return
	case 9:
		r.CodigoISOPais = nil
		return
	case 10:
		r.Localidad = nil
		return
	case 11:
		r.ComponentesDeDireccion = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Direccion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Direccion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Direccion) HintSize(int)                     { panic("Unsupported operation") }
func (_ Direccion) Finalize()                        {}

func (_ Direccion) AvroCRC64Fingerprint() []byte {
	return []byte(DireccionAvroCRC64Fingerprint)
}

func (r Direccion) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["abreviaturaProvincia"], err = json.Marshal(r.AbreviaturaProvincia)
	if err != nil {
		return nil, err
	}
	output["calle"], err = json.Marshal(r.Calle)
	if err != nil {
		return nil, err
	}
	output["codigoDeDireccion"], err = json.Marshal(r.CodigoDeDireccion)
	if err != nil {
		return nil, err
	}
	output["codigoCiudad"], err = json.Marshal(r.CodigoCiudad)
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
	output["numero"], err = json.Marshal(r.Numero)
	if err != nil {
		return nil, err
	}
	output["telefono"], err = json.Marshal(r.Telefono)
	if err != nil {
		return nil, err
	}
	output["codigoISOProvincia"], err = json.Marshal(r.CodigoISOProvincia)
	if err != nil {
		return nil, err
	}
	output["codigoISOPais"], err = json.Marshal(r.CodigoISOPais)
	if err != nil {
		return nil, err
	}
	output["localidad"], err = json.Marshal(r.Localidad)
	if err != nil {
		return nil, err
	}
	output["componentesDeDireccion"], err = json.Marshal(r.ComponentesDeDireccion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Direccion) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["abreviaturaProvincia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AbreviaturaProvincia); err != nil {
			return err
		}
	} else {
		r.AbreviaturaProvincia = NewUnionNullString()

		r.AbreviaturaProvincia = nil
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
		if v, ok := fields["codigoDeDireccion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeDireccion); err != nil {
			return err
		}
	} else {
		r.CodigoDeDireccion = NewUnionNullString()

		r.CodigoDeDireccion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoCiudad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoCiudad); err != nil {
			return err
		}
	} else {
		r.CodigoCiudad = NewUnionNullString()

		r.CodigoCiudad = nil
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
		r.Numero = NewUnionNullString()

		r.Numero = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["telefono"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Telefono); err != nil {
			return err
		}
	} else {
		r.Telefono = NewUnionNullString()

		r.Telefono = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoISOProvincia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoISOProvincia); err != nil {
			return err
		}
	} else {
		r.CodigoISOProvincia = NewUnionNullString()

		r.CodigoISOProvincia = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoISOPais"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoISOPais); err != nil {
			return err
		}
	} else {
		r.CodigoISOPais = NewUnionNullString()

		r.CodigoISOPais = nil
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
		if v, ok := fields["componentesDeDireccion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ComponentesDeDireccion); err != nil {
			return err
		}
	} else {
		r.ComponentesDeDireccion = NewUnionNullListaDePropiedades()

		r.ComponentesDeDireccion = nil
	}
	return nil
}
