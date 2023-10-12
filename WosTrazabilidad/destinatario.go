// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     WosTrazaAnmat.avsc
 */
package WosTrazabilidadEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Destinatario struct {
	Nombre string `json:"Nombre"`

	Apellido string `json:"Apellido"`

	TipoDocumento int32 `json:"TipoDocumento"`
}

const DestinatarioAvroCRC64Fingerprint = "gJG\x17\xed\xc8RM"

func NewDestinatario() Destinatario {
	r := Destinatario{}
	return r
}

func DeserializeDestinatario(r io.Reader) (Destinatario, error) {
	t := NewDestinatario()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDestinatarioFromSchema(r io.Reader, schema string) (Destinatario, error) {
	t := NewDestinatario()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDestinatario(r Destinatario, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Nombre, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Apellido, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.TipoDocumento, w)
	if err != nil {
		return err
	}
	return err
}

func (r Destinatario) Serialize(w io.Writer) error {
	return writeDestinatario(r, w)
}

func (r Destinatario) Schema() string {
	return "{\"fields\":[{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Apellido\",\"type\":\"string\"},{\"name\":\"TipoDocumento\",\"type\":\"int\"}],\"name\":\"Andreani.WosTrazabilidad.Events.AnmatCommon.Destinatario\",\"type\":\"record\"}"
}

func (r Destinatario) SchemaName() string {
	return "Andreani.WosTrazabilidad.Events.AnmatCommon.Destinatario"
}

func (_ Destinatario) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Destinatario) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Destinatario) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Destinatario) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Destinatario) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Destinatario) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Destinatario) SetString(v string)   { panic("Unsupported operation") }
func (_ Destinatario) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Destinatario) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Nombre}

		return w

	case 1:
		w := types.String{Target: &r.Apellido}

		return w

	case 2:
		w := types.Int{Target: &r.TipoDocumento}

		return w

	}
	panic("Unknown field index")
}

func (r *Destinatario) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Destinatario) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Destinatario) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Destinatario) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Destinatario) HintSize(int)                     { panic("Unsupported operation") }
func (_ Destinatario) Finalize()                        {}

func (_ Destinatario) AvroCRC64Fingerprint() []byte {
	return []byte(DestinatarioAvroCRC64Fingerprint)
}

func (r Destinatario) MarshalJSON() ([]byte, error) {
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
	output["TipoDocumento"], err = json.Marshal(r.TipoDocumento)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Destinatario) UnmarshalJSON(data []byte) error {
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
		return fmt.Errorf("no value specified for Apellido")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoDocumento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDocumento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoDocumento")
	}
	return nil
}
