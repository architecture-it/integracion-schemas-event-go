// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     GestionEntrega.avsc
 */
package AndiExternoEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type GestionEntrega struct {
	NroEnvio string `json:"NroEnvio"`
}

const GestionEntregaAvroCRC64Fingerprint = "Z\xc3䊱\xc7\xd5%"

func NewGestionEntrega() GestionEntrega {
	r := GestionEntrega{}
	return r
}

func DeserializeGestionEntrega(r io.Reader) (GestionEntrega, error) {
	t := NewGestionEntrega()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeGestionEntregaFromSchema(r io.Reader, schema string) (GestionEntrega, error) {
	t := NewGestionEntrega()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeGestionEntrega(r GestionEntrega, w io.Writer) error {
	var err error
	err = vm.WriteString(r.NroEnvio, w)
	if err != nil {
		return err
	}
	return err
}

func (r GestionEntrega) Serialize(w io.Writer) error {
	return writeGestionEntrega(r, w)
}

func (r GestionEntrega) Schema() string {
	return "{\"fields\":[{\"name\":\"NroEnvio\",\"type\":\"string\"}],\"name\":\"Andreani.AndiExterno.Events.Record.GestionEntrega\",\"type\":\"record\"}"
}

func (r GestionEntrega) SchemaName() string {
	return "Andreani.AndiExterno.Events.Record.GestionEntrega"
}

func (_ GestionEntrega) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ GestionEntrega) SetInt(v int32)       { panic("Unsupported operation") }
func (_ GestionEntrega) SetLong(v int64)      { panic("Unsupported operation") }
func (_ GestionEntrega) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ GestionEntrega) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ GestionEntrega) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ GestionEntrega) SetString(v string)   { panic("Unsupported operation") }
func (_ GestionEntrega) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *GestionEntrega) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.NroEnvio}

		return w

	}
	panic("Unknown field index")
}

func (r *GestionEntrega) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *GestionEntrega) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ GestionEntrega) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ GestionEntrega) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ GestionEntrega) HintSize(int)                     { panic("Unsupported operation") }
func (_ GestionEntrega) Finalize()                        {}

func (_ GestionEntrega) AvroCRC64Fingerprint() []byte {
	return []byte(GestionEntregaAvroCRC64Fingerprint)
}

func (r GestionEntrega) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["NroEnvio"], err = json.Marshal(r.NroEnvio)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *GestionEntrega) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["NroEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NroEnvio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NroEnvio")
	}
	return nil
}
