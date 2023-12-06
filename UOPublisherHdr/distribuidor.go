// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     UOHdrCreada.avsc
 */
package UOPublisherHdrEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Distribuidor struct {
	Dni string `json:"Dni"`

	Nombre *UnionNullString `json:"Nombre"`

	Apellido *UnionNullString `json:"Apellido"`

	Movilidad *UnionNullInt `json:"Movilidad"`
}

const DistribuidorAvroCRC64Fingerprint = "\xdaNB\xd0u\xc1\x0e\x7f"

func NewDistribuidor() Distribuidor {
	r := Distribuidor{}
	r.Nombre = nil
	r.Apellido = nil
	r.Movilidad = nil
	return r
}

func DeserializeDistribuidor(r io.Reader) (Distribuidor, error) {
	t := NewDistribuidor()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDistribuidorFromSchema(r io.Reader, schema string) (Distribuidor, error) {
	t := NewDistribuidor()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDistribuidor(r Distribuidor, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Dni, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Nombre, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Apellido, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.Movilidad, w)
	if err != nil {
		return err
	}
	return err
}

func (r Distribuidor) Serialize(w io.Writer) error {
	return writeDistribuidor(r, w)
}

func (r Distribuidor) Schema() string {
	return "{\"fields\":[{\"name\":\"Dni\",\"type\":\"string\"},{\"default\":null,\"name\":\"Nombre\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Apellido\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Movilidad\",\"type\":[\"null\",\"int\"]}],\"name\":\"Andreani.UOPublisherHdr.Events.Common.Distribuidor\",\"type\":\"record\"}"
}

func (r Distribuidor) SchemaName() string {
	return "Andreani.UOPublisherHdr.Events.Common.Distribuidor"
}

func (_ Distribuidor) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Distribuidor) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Distribuidor) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Distribuidor) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Distribuidor) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Distribuidor) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Distribuidor) SetString(v string)   { panic("Unsupported operation") }
func (_ Distribuidor) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Distribuidor) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Dni}

		return w

	case 1:
		r.Nombre = NewUnionNullString()

		return r.Nombre
	case 2:
		r.Apellido = NewUnionNullString()

		return r.Apellido
	case 3:
		r.Movilidad = NewUnionNullInt()

		return r.Movilidad
	}
	panic("Unknown field index")
}

func (r *Distribuidor) SetDefault(i int) {
	switch i {
	case 1:
		r.Nombre = nil
		return
	case 2:
		r.Apellido = nil
		return
	case 3:
		r.Movilidad = nil
		return
	}
	panic("Unknown field index")
}

func (r *Distribuidor) NullField(i int) {
	switch i {
	case 1:
		r.Nombre = nil
		return
	case 2:
		r.Apellido = nil
		return
	case 3:
		r.Movilidad = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Distribuidor) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Distribuidor) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Distribuidor) HintSize(int)                     { panic("Unsupported operation") }
func (_ Distribuidor) Finalize()                        {}

func (_ Distribuidor) AvroCRC64Fingerprint() []byte {
	return []byte(DistribuidorAvroCRC64Fingerprint)
}

func (r Distribuidor) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Dni"], err = json.Marshal(r.Dni)
	if err != nil {
		return nil, err
	}
	output["Nombre"], err = json.Marshal(r.Nombre)
	if err != nil {
		return nil, err
	}
	output["Apellido"], err = json.Marshal(r.Apellido)
	if err != nil {
		return nil, err
	}
	output["Movilidad"], err = json.Marshal(r.Movilidad)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Distribuidor) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
		r.Nombre = NewUnionNullString()

		r.Nombre = nil
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
		if v, ok := fields["Movilidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Movilidad); err != nil {
			return err
		}
	} else {
		r.Movilidad = NewUnionNullInt()

		r.Movilidad = nil
	}
	return nil
}
