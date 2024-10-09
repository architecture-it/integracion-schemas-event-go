// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ListInterface640Bienes.avsc
 */
package HCMInterface640MdaBienesEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Interface640BienesData struct {
	Legajo string `json:"Legajo"`

	Dni string `json:"Dni"`

	Nombre string `json:"Nombre"`

	Usuario string `json:"Usuario"`

	FechaBaja string `json:"FechaBaja"`
}

const Interface640BienesDataAvroCRC64Fingerprint = "\xeb\xb7\xc2\xea9\x18\xfbx"

func NewInterface640BienesData() Interface640BienesData {
	r := Interface640BienesData{}
	return r
}

func DeserializeInterface640BienesData(r io.Reader) (Interface640BienesData, error) {
	t := NewInterface640BienesData()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeInterface640BienesDataFromSchema(r io.Reader, schema string) (Interface640BienesData, error) {
	t := NewInterface640BienesData()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeInterface640BienesData(r Interface640BienesData, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Legajo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Dni, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Nombre, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Usuario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaBaja, w)
	if err != nil {
		return err
	}
	return err
}

func (r Interface640BienesData) Serialize(w io.Writer) error {
	return writeInterface640BienesData(r, w)
}

func (r Interface640BienesData) Schema() string {
	return "{\"fields\":[{\"name\":\"Legajo\",\"type\":\"string\"},{\"name\":\"Dni\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Usuario\",\"type\":\"string\"},{\"name\":\"FechaBaja\",\"type\":\"string\"}],\"name\":\"Andreani.HCMInterface640MdaBienes.Events.Record.Interface640BienesData\",\"type\":\"record\"}"
}

func (r Interface640BienesData) SchemaName() string {
	return "Andreani.HCMInterface640MdaBienes.Events.Record.Interface640BienesData"
}

func (_ Interface640BienesData) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Interface640BienesData) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Interface640BienesData) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Interface640BienesData) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Interface640BienesData) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Interface640BienesData) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Interface640BienesData) SetString(v string)   { panic("Unsupported operation") }
func (_ Interface640BienesData) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Interface640BienesData) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Legajo}

		return w

	case 1:
		w := types.String{Target: &r.Dni}

		return w

	case 2:
		w := types.String{Target: &r.Nombre}

		return w

	case 3:
		w := types.String{Target: &r.Usuario}

		return w

	case 4:
		w := types.String{Target: &r.FechaBaja}

		return w

	}
	panic("Unknown field index")
}

func (r *Interface640BienesData) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Interface640BienesData) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Interface640BienesData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Interface640BienesData) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Interface640BienesData) HintSize(int)                     { panic("Unsupported operation") }
func (_ Interface640BienesData) Finalize()                        {}

func (_ Interface640BienesData) AvroCRC64Fingerprint() []byte {
	return []byte(Interface640BienesDataAvroCRC64Fingerprint)
}

func (r Interface640BienesData) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Legajo"], err = json.Marshal(r.Legajo)
	if err != nil {
		return nil, err
	}
	output["Dni"], err = json.Marshal(r.Dni)
	if err != nil {
		return nil, err
	}
	output["Nombre"], err = json.Marshal(r.Nombre)
	if err != nil {
		return nil, err
	}
	output["Usuario"], err = json.Marshal(r.Usuario)
	if err != nil {
		return nil, err
	}
	output["FechaBaja"], err = json.Marshal(r.FechaBaja)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Interface640BienesData) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Legajo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Legajo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Legajo")
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
		if v, ok := fields["Usuario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Usuario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Usuario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaBaja"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaBaja); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaBaja")
	}
	return nil
}
