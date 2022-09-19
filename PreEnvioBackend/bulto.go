// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ResponseGenerarPreenvios.avsc
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

type Bulto struct {
	NumeroDeBulto *UnionNullString `json:"NumeroDeBulto"`

	NumeroDeEnvio *UnionNullString `json:"NumeroDeEnvio"`

	Totalizador *UnionNullString `json:"Totalizador"`

	Linking *UnionNullArrayComponentesDeDireccion `json:"Linking"`
}

const BultoAvroCRC64Fingerprint = "\\\xf4P\f&,(1"

func NewBulto() Bulto {
	r := Bulto{}
	r.NumeroDeBulto = nil
	r.NumeroDeEnvio = nil
	r.Totalizador = nil
	r.Linking = nil
	return r
}

func DeserializeBulto(r io.Reader) (Bulto, error) {
	t := NewBulto()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBultoFromSchema(r io.Reader, schema string) (Bulto, error) {
	t := NewBulto()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBulto(r Bulto, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.NumeroDeBulto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeEnvio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Totalizador, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayComponentesDeDireccion(r.Linking, w)
	if err != nil {
		return err
	}
	return err
}

func (r Bulto) Serialize(w io.Writer) error {
	return writeBulto(r, w)
}

func (r Bulto) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"NumeroDeBulto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Totalizador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Linking\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"Meta\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Contenido\",\"type\":[\"null\",\"string\"]}],\"name\":\"ComponentesDeDireccion\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"Andreani.PreEnvioBackend.Events.Preenvio.Common.Bulto\",\"type\":\"record\"}"
}

func (r Bulto) SchemaName() string {
	return "Andreani.PreEnvioBackend.Events.Preenvio.Common.Bulto"
}

func (_ Bulto) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Bulto) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Bulto) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Bulto) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Bulto) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Bulto) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Bulto) SetString(v string)   { panic("Unsupported operation") }
func (_ Bulto) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Bulto) Get(i int) types.Field {
	switch i {
	case 0:
		r.NumeroDeBulto = NewUnionNullString()

		return r.NumeroDeBulto
	case 1:
		r.NumeroDeEnvio = NewUnionNullString()

		return r.NumeroDeEnvio
	case 2:
		r.Totalizador = NewUnionNullString()

		return r.Totalizador
	case 3:
		r.Linking = NewUnionNullArrayComponentesDeDireccion()

		return r.Linking
	}
	panic("Unknown field index")
}

func (r *Bulto) SetDefault(i int) {
	switch i {
	case 0:
		r.NumeroDeBulto = nil
		return
	case 1:
		r.NumeroDeEnvio = nil
		return
	case 2:
		r.Totalizador = nil
		return
	case 3:
		r.Linking = nil
		return
	}
	panic("Unknown field index")
}

func (r *Bulto) NullField(i int) {
	switch i {
	case 0:
		r.NumeroDeBulto = nil
		return
	case 1:
		r.NumeroDeEnvio = nil
		return
	case 2:
		r.Totalizador = nil
		return
	case 3:
		r.Linking = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Bulto) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Bulto) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Bulto) HintSize(int)                     { panic("Unsupported operation") }
func (_ Bulto) Finalize()                        {}

func (_ Bulto) AvroCRC64Fingerprint() []byte {
	return []byte(BultoAvroCRC64Fingerprint)
}

func (r Bulto) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["NumeroDeBulto"], err = json.Marshal(r.NumeroDeBulto)
	if err != nil {
		return nil, err
	}
	output["NumeroDeEnvio"], err = json.Marshal(r.NumeroDeEnvio)
	if err != nil {
		return nil, err
	}
	output["Totalizador"], err = json.Marshal(r.Totalizador)
	if err != nil {
		return nil, err
	}
	output["Linking"], err = json.Marshal(r.Linking)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Bulto) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeBulto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeBulto); err != nil {
			return err
		}
	} else {
		r.NumeroDeBulto = NewUnionNullString()

		r.NumeroDeBulto = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeEnvio); err != nil {
			return err
		}
	} else {
		r.NumeroDeEnvio = NewUnionNullString()

		r.NumeroDeEnvio = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Totalizador"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Totalizador); err != nil {
			return err
		}
	} else {
		r.Totalizador = NewUnionNullString()

		r.Totalizador = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Linking"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Linking); err != nil {
			return err
		}
	} else {
		r.Linking = NewUnionNullArrayComponentesDeDireccion()

		r.Linking = nil
	}
	return nil
}
