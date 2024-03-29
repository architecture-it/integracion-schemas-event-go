// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     WapTrasladoInternoDeMercaderiaSolicitadaV2.avsc
 */
package WapTrasladoInternoDeMercaderiaV2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Propietario struct {
	Destino string `json:"Destino"`

	Origen string `json:"Origen"`
}

const PropietarioAvroCRC64Fingerprint = "\x0e\xbc\xa8\xae\x1fս+"

func NewPropietario() Propietario {
	r := Propietario{}
	return r
}

func DeserializePropietario(r io.Reader) (Propietario, error) {
	t := NewPropietario()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePropietarioFromSchema(r io.Reader, schema string) (Propietario, error) {
	t := NewPropietario()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePropietario(r Propietario, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Destino, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Origen, w)
	if err != nil {
		return err
	}
	return err
}

func (r Propietario) Serialize(w io.Writer) error {
	return writePropietario(r, w)
}

func (r Propietario) Schema() string {
	return "{\"fields\":[{\"name\":\"Destino\",\"type\":\"string\"},{\"name\":\"Origen\",\"type\":\"string\"}],\"name\":\"Andreani.WapTrasladoInternoDeMercaderiaV2.Events.Record.Propietario\",\"type\":\"record\"}"
}

func (r Propietario) SchemaName() string {
	return "Andreani.WapTrasladoInternoDeMercaderiaV2.Events.Record.Propietario"
}

func (_ Propietario) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Propietario) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Propietario) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Propietario) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Propietario) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Propietario) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Propietario) SetString(v string)   { panic("Unsupported operation") }
func (_ Propietario) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Propietario) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Destino}

		return w

	case 1:
		w := types.String{Target: &r.Origen}

		return w

	}
	panic("Unknown field index")
}

func (r *Propietario) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Propietario) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Propietario) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Propietario) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Propietario) HintSize(int)                     { panic("Unsupported operation") }
func (_ Propietario) Finalize()                        {}

func (_ Propietario) AvroCRC64Fingerprint() []byte {
	return []byte(PropietarioAvroCRC64Fingerprint)
}

func (r Propietario) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Destino"], err = json.Marshal(r.Destino)
	if err != nil {
		return nil, err
	}
	output["Origen"], err = json.Marshal(r.Origen)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Propietario) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Destino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Destino); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Destino")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Origen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Origen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Origen")
	}
	return nil
}
