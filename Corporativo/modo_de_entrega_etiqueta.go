// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ModoDeEntregaEtiqueta.avsc
 */
package CorporativoEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ModoDeEntregaEtiqueta struct {
	Nombre string `json:"Nombre"`
}

const ModoDeEntregaEtiquetaAvroCRC64Fingerprint = "\xf5\x86N)}\xb4\x03\x82"

func NewModoDeEntregaEtiqueta() ModoDeEntregaEtiqueta {
	r := ModoDeEntregaEtiqueta{}
	return r
}

func DeserializeModoDeEntregaEtiqueta(r io.Reader) (ModoDeEntregaEtiqueta, error) {
	t := NewModoDeEntregaEtiqueta()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeModoDeEntregaEtiquetaFromSchema(r io.Reader, schema string) (ModoDeEntregaEtiqueta, error) {
	t := NewModoDeEntregaEtiqueta()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeModoDeEntregaEtiqueta(r ModoDeEntregaEtiqueta, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Nombre, w)
	if err != nil {
		return err
	}
	return err
}

func (r ModoDeEntregaEtiqueta) Serialize(w io.Writer) error {
	return writeModoDeEntregaEtiqueta(r, w)
}

func (r ModoDeEntregaEtiqueta) Schema() string {
	return "{\"fields\":[{\"name\":\"Nombre\",\"type\":\"string\"}],\"name\":\"Andreani.Corporativo.Events.Record.ModoDeEntregaEtiqueta\",\"type\":\"record\"}"
}

func (r ModoDeEntregaEtiqueta) SchemaName() string {
	return "Andreani.Corporativo.Events.Record.ModoDeEntregaEtiqueta"
}

func (_ ModoDeEntregaEtiqueta) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ModoDeEntregaEtiqueta) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ModoDeEntregaEtiqueta) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ModoDeEntregaEtiqueta) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ModoDeEntregaEtiqueta) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ModoDeEntregaEtiqueta) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ModoDeEntregaEtiqueta) SetString(v string)   { panic("Unsupported operation") }
func (_ ModoDeEntregaEtiqueta) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ModoDeEntregaEtiqueta) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Nombre}

		return w

	}
	panic("Unknown field index")
}

func (r *ModoDeEntregaEtiqueta) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ModoDeEntregaEtiqueta) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ModoDeEntregaEtiqueta) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ModoDeEntregaEtiqueta) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ModoDeEntregaEtiqueta) HintSize(int)                     { panic("Unsupported operation") }
func (_ ModoDeEntregaEtiqueta) Finalize()                        {}

func (_ ModoDeEntregaEtiqueta) AvroCRC64Fingerprint() []byte {
	return []byte(ModoDeEntregaEtiquetaAvroCRC64Fingerprint)
}

func (r ModoDeEntregaEtiqueta) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Nombre"], err = json.Marshal(r.Nombre)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ModoDeEntregaEtiqueta) UnmarshalJSON(data []byte) error {
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
	return nil
}
