// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Remitente.avsc
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

type Remitente struct {
	Nombre string `json:"Nombre"`

	Apellido *UnionNullString `json:"Apellido"`

	Email string `json:"Email"`

	Telefono string `json:"Telefono"`

	Dni string `json:"Dni"`

	Calle string `json:"Calle"`

	Numero *UnionNullString `json:"Numero"`

	Localidad string `json:"Localidad"`

	CodigoPostal string `json:"CodigoPostal"`

	Piso *UnionNullString `json:"Piso"`

	Unidad *UnionNullString `json:"Unidad"`

	Pais string `json:"Pais"`

	Region string `json:"Region"`
}

const RemitenteAvroCRC64Fingerprint = "\t20\xe45\xdc\xcfW"

func NewRemitente() Remitente {
	r := Remitente{}
	r.Apellido = nil
	r.Numero = nil
	r.Piso = nil
	r.Unidad = nil
	return r
}

func DeserializeRemitente(r io.Reader) (Remitente, error) {
	t := NewRemitente()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRemitenteFromSchema(r io.Reader, schema string) (Remitente, error) {
	t := NewRemitente()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRemitente(r Remitente, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Nombre, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Apellido, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Email, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Telefono, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Dni, w)
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
	err = vm.WriteString(r.Localidad, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoPostal, w)
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
	err = vm.WriteString(r.Pais, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Region, w)
	if err != nil {
		return err
	}
	return err
}

func (r Remitente) Serialize(w io.Writer) error {
	return writeRemitente(r, w)
}

func (r Remitente) Schema() string {
	return "{\"fields\":[{\"name\":\"Nombre\",\"type\":\"string\"},{\"default\":null,\"name\":\"Apellido\",\"type\":[\"null\",\"string\"]},{\"name\":\"Email\",\"type\":\"string\"},{\"name\":\"Telefono\",\"type\":\"string\"},{\"name\":\"Dni\",\"type\":\"string\"},{\"name\":\"Calle\",\"type\":\"string\"},{\"default\":null,\"name\":\"Numero\",\"type\":[\"null\",\"string\"]},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"CodigoPostal\",\"type\":\"string\"},{\"default\":null,\"name\":\"Piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Unidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"Pais\",\"type\":\"string\"},{\"name\":\"Region\",\"type\":\"string\"}],\"name\":\"Andreani.PersonaBackend.Events.Common.Remitente\",\"type\":\"record\"}"
}

func (r Remitente) SchemaName() string {
	return "Andreani.PersonaBackend.Events.Common.Remitente"
}

func (_ Remitente) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Remitente) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Remitente) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Remitente) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Remitente) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Remitente) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Remitente) SetString(v string)   { panic("Unsupported operation") }
func (_ Remitente) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Remitente) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Nombre}

		return w

	case 1:
		r.Apellido = NewUnionNullString()

		return r.Apellido
	case 2:
		w := types.String{Target: &r.Email}

		return w

	case 3:
		w := types.String{Target: &r.Telefono}

		return w

	case 4:
		w := types.String{Target: &r.Dni}

		return w

	case 5:
		w := types.String{Target: &r.Calle}

		return w

	case 6:
		r.Numero = NewUnionNullString()

		return r.Numero
	case 7:
		w := types.String{Target: &r.Localidad}

		return w

	case 8:
		w := types.String{Target: &r.CodigoPostal}

		return w

	case 9:
		r.Piso = NewUnionNullString()

		return r.Piso
	case 10:
		r.Unidad = NewUnionNullString()

		return r.Unidad
	case 11:
		w := types.String{Target: &r.Pais}

		return w

	case 12:
		w := types.String{Target: &r.Region}

		return w

	}
	panic("Unknown field index")
}

func (r *Remitente) SetDefault(i int) {
	switch i {
	case 1:
		r.Apellido = nil
		return
	case 6:
		r.Numero = nil
		return
	case 9:
		r.Piso = nil
		return
	case 10:
		r.Unidad = nil
		return
	}
	panic("Unknown field index")
}

func (r *Remitente) NullField(i int) {
	switch i {
	case 1:
		r.Apellido = nil
		return
	case 6:
		r.Numero = nil
		return
	case 9:
		r.Piso = nil
		return
	case 10:
		r.Unidad = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Remitente) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Remitente) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Remitente) HintSize(int)                     { panic("Unsupported operation") }
func (_ Remitente) Finalize()                        {}

func (_ Remitente) AvroCRC64Fingerprint() []byte {
	return []byte(RemitenteAvroCRC64Fingerprint)
}

func (r Remitente) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Nombre"], err = json.Marshal(r.Nombre)
	if err != nil {
		return nil, err
	}
	output["Apellido"], err = json.Marshal(r.Apellido)
	if err != nil {
		return nil, err
	}
	output["Email"], err = json.Marshal(r.Email)
	if err != nil {
		return nil, err
	}
	output["Telefono"], err = json.Marshal(r.Telefono)
	if err != nil {
		return nil, err
	}
	output["Dni"], err = json.Marshal(r.Dni)
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
	output["Localidad"], err = json.Marshal(r.Localidad)
	if err != nil {
		return nil, err
	}
	output["CodigoPostal"], err = json.Marshal(r.CodigoPostal)
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
	output["Pais"], err = json.Marshal(r.Pais)
	if err != nil {
		return nil, err
	}
	output["Region"], err = json.Marshal(r.Region)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Remitente) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Nombre"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Nombre); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Nombre")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Apellido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Apellido); err != nil {
			return err
		}
	} else {
		r.Apellido = NewUnionNullString()

		r.Apellido = nil
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
		return fmt.Errorf("no value specified for Email")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Telefono"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Telefono); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Telefono")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Dni"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Dni); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Dni")
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
		if v, ok := fields["Pais"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Pais); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Pais")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Region"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Region); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Region")
	}
	return nil
}
