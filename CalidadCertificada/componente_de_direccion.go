// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Postal.avsc
 */
package CalidadCertificadaEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ComponenteDeDireccion struct {
	Meta string `json:"Meta"`

	Contenido string `json:"Contenido"`
}

const ComponenteDeDireccionAvroCRC64Fingerprint = "G\"H\x17\x94PY\xca"

func NewComponenteDeDireccion() ComponenteDeDireccion {
	r := ComponenteDeDireccion{}
	return r
}

func DeserializeComponenteDeDireccion(r io.Reader) (ComponenteDeDireccion, error) {
	t := NewComponenteDeDireccion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeComponenteDeDireccionFromSchema(r io.Reader, schema string) (ComponenteDeDireccion, error) {
	t := NewComponenteDeDireccion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeComponenteDeDireccion(r ComponenteDeDireccion, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Meta, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Contenido, w)
	if err != nil {
		return err
	}
	return err
}

func (r ComponenteDeDireccion) Serialize(w io.Writer) error {
	return writeComponenteDeDireccion(r, w)
}

func (r ComponenteDeDireccion) Schema() string {
	return "{\"fields\":[{\"name\":\"Meta\",\"type\":\"string\"},{\"name\":\"Contenido\",\"type\":\"string\"}],\"name\":\"Andreani.CalidadCertificada.Events.Record.ComponenteDeDireccion\",\"type\":\"record\"}"
}

func (r ComponenteDeDireccion) SchemaName() string {
	return "Andreani.CalidadCertificada.Events.Record.ComponenteDeDireccion"
}

func (_ ComponenteDeDireccion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ComponenteDeDireccion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ComponenteDeDireccion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ComponenteDeDireccion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ComponenteDeDireccion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ComponenteDeDireccion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ComponenteDeDireccion) SetString(v string)   { panic("Unsupported operation") }
func (_ ComponenteDeDireccion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ComponenteDeDireccion) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Meta}

		return w

	case 1:
		w := types.String{Target: &r.Contenido}

		return w

	}
	panic("Unknown field index")
}

func (r *ComponenteDeDireccion) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ComponenteDeDireccion) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ComponenteDeDireccion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ComponenteDeDireccion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ComponenteDeDireccion) HintSize(int)                     { panic("Unsupported operation") }
func (_ ComponenteDeDireccion) Finalize()                        {}

func (_ ComponenteDeDireccion) AvroCRC64Fingerprint() []byte {
	return []byte(ComponenteDeDireccionAvroCRC64Fingerprint)
}

func (r ComponenteDeDireccion) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Meta"], err = json.Marshal(r.Meta)
	if err != nil {
		return nil, err
	}
	output["Contenido"], err = json.Marshal(r.Contenido)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ComponenteDeDireccion) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Meta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Meta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Meta")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Contenido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contenido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Contenido")
	}
	return nil
}
