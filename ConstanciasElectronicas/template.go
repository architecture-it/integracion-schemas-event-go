// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Template.avsc
 */
package ConstanciasElectronicasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Template struct {
	Type TemplateType `json:"type"`

	Modules string `json:"modules"`

	Status EntityStatus `json:"status"`
}

const TemplateAvroCRC64Fingerprint = "\ff\x98Kxɯ\xaa"

func NewTemplate() Template {
	r := Template{}
	return r
}

func DeserializeTemplate(r io.Reader) (Template, error) {
	t := NewTemplate()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeTemplateFromSchema(r io.Reader, schema string) (Template, error) {
	t := NewTemplate()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeTemplate(r Template, w io.Writer) error {
	var err error
	err = writeTemplateType(r.Type, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Modules, w)
	if err != nil {
		return err
	}
	err = writeEntityStatus(r.Status, w)
	if err != nil {
		return err
	}
	return err
}

func (r Template) Serialize(w io.Writer) error {
	return writeTemplate(r, w)
}

func (r Template) Schema() string {
	return "{\"fields\":[{\"name\":\"type\",\"type\":{\"name\":\"TemplateType\",\"symbols\":[\"Undefined\",\"Delivery\",\"NoDelivered\"],\"type\":\"enum\"}},{\"name\":\"modules\",\"type\":\"string\"},{\"name\":\"status\",\"type\":{\"name\":\"EntityStatus\",\"symbols\":[\"Inactive\",\"Active\",\"Expired\",\"Deleted\"],\"type\":\"enum\"}}],\"name\":\"Andreani.ConstanciasElectronicas.Common.Template\",\"type\":\"record\"}"
}

func (r Template) SchemaName() string {
	return "Andreani.ConstanciasElectronicas.Common.Template"
}

func (_ Template) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Template) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Template) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Template) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Template) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Template) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Template) SetString(v string)   { panic("Unsupported operation") }
func (_ Template) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Template) Get(i int) types.Field {
	switch i {
	case 0:
		w := TemplateTypeWrapper{Target: &r.Type}

		return w

	case 1:
		w := types.String{Target: &r.Modules}

		return w

	case 2:
		w := EntityStatusWrapper{Target: &r.Status}

		return w

	}
	panic("Unknown field index")
}

func (r *Template) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Template) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Template) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Template) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Template) HintSize(int)                     { panic("Unsupported operation") }
func (_ Template) Finalize()                        {}

func (_ Template) AvroCRC64Fingerprint() []byte {
	return []byte(TemplateAvroCRC64Fingerprint)
}

func (r Template) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["type"], err = json.Marshal(r.Type)
	if err != nil {
		return nil, err
	}
	output["modules"], err = json.Marshal(r.Modules)
	if err != nil {
		return nil, err
	}
	output["status"], err = json.Marshal(r.Status)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Template) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["type"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Type); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for type")
	}
	val = func() json.RawMessage {
		if v, ok := fields["modules"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Modules); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for modules")
	}
	val = func() json.RawMessage {
		if v, ok := fields["status"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Status); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for status")
	}
	return nil
}