// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     UsuarioEtiqueta.avsc
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

type ClienteEtiqueta struct {
	CodigoAndreani string `json:"CodigoAndreani"`

	Nombre string `json:"Nombre"`

	Logo string `json:"Logo"`
}

const ClienteEtiquetaAvroCRC64Fingerprint = "h\x81\x100~#\xde\xff"

func NewClienteEtiqueta() ClienteEtiqueta {
	r := ClienteEtiqueta{}
	return r
}

func DeserializeClienteEtiqueta(r io.Reader) (ClienteEtiqueta, error) {
	t := NewClienteEtiqueta()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeClienteEtiquetaFromSchema(r io.Reader, schema string) (ClienteEtiqueta, error) {
	t := NewClienteEtiqueta()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeClienteEtiqueta(r ClienteEtiqueta, w io.Writer) error {
	var err error
	err = vm.WriteString(r.CodigoAndreani, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Nombre, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Logo, w)
	if err != nil {
		return err
	}
	return err
}

func (r ClienteEtiqueta) Serialize(w io.Writer) error {
	return writeClienteEtiqueta(r, w)
}

func (r ClienteEtiqueta) Schema() string {
	return "{\"fields\":[{\"name\":\"CodigoAndreani\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Logo\",\"type\":\"string\"}],\"name\":\"Andreani.Corporativo.Events.Record.ClienteEtiqueta\",\"type\":\"record\"}"
}

func (r ClienteEtiqueta) SchemaName() string {
	return "Andreani.Corporativo.Events.Record.ClienteEtiqueta"
}

func (_ ClienteEtiqueta) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ClienteEtiqueta) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ClienteEtiqueta) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ClienteEtiqueta) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ClienteEtiqueta) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ClienteEtiqueta) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ClienteEtiqueta) SetString(v string)   { panic("Unsupported operation") }
func (_ ClienteEtiqueta) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ClienteEtiqueta) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.CodigoAndreani}

		return w

	case 1:
		w := types.String{Target: &r.Nombre}

		return w

	case 2:
		w := types.String{Target: &r.Logo}

		return w

	}
	panic("Unknown field index")
}

func (r *ClienteEtiqueta) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ClienteEtiqueta) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ClienteEtiqueta) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ClienteEtiqueta) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ClienteEtiqueta) HintSize(int)                     { panic("Unsupported operation") }
func (_ ClienteEtiqueta) Finalize()                        {}

func (_ ClienteEtiqueta) AvroCRC64Fingerprint() []byte {
	return []byte(ClienteEtiquetaAvroCRC64Fingerprint)
}

func (r ClienteEtiqueta) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["CodigoAndreani"], err = json.Marshal(r.CodigoAndreani)
	if err != nil {
		return nil, err
	}
	output["Nombre"], err = json.Marshal(r.Nombre)
	if err != nil {
		return nil, err
	}
	output["Logo"], err = json.Marshal(r.Logo)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ClienteEtiqueta) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["CodigoAndreani"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoAndreani); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodigoAndreani")
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
		return fmt.Errorf("no value specified for Nombre")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Logo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Logo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Logo")
	}
	return nil
}
