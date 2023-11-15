// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ListInterface282.avsc
 */
package HCMInterface282Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ListInterface282 struct {
	Interfaces []Interface282Data `json:"interfaces"`
}

const ListInterface282AvroCRC64Fingerprint = "\xe2\xc8\xf8\xaazļ'"

func NewListInterface282() ListInterface282 {
	r := ListInterface282{}
	r.Interfaces = make([]Interface282Data, 0)

	return r
}

func DeserializeListInterface282(r io.Reader) (ListInterface282, error) {
	t := NewListInterface282()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeListInterface282FromSchema(r io.Reader, schema string) (ListInterface282, error) {
	t := NewListInterface282()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeListInterface282(r ListInterface282, w io.Writer) error {
	var err error
	err = writeArrayInterface282Data(r.Interfaces, w)
	if err != nil {
		return err
	}
	return err
}

func (r ListInterface282) Serialize(w io.Writer) error {
	return writeListInterface282(r, w)
}

func (r ListInterface282) Schema() string {
	return "{\"fields\":[{\"name\":\"interfaces\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Legajo\",\"type\":\"long\"},{\"name\":\"FechaAcuerdo\",\"type\":\"string\"},{\"name\":\"SueldoJournal\",\"type\":\"long\"},{\"name\":\"ZonaBandaSalarial\",\"type\":\"string\"}],\"name\":\"Interface282Data\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.HCMInterface282.Events.Record.ListInterface282\",\"type\":\"record\"}"
}

func (r ListInterface282) SchemaName() string {
	return "Andreani.HCMInterface282.Events.Record.ListInterface282"
}

func (_ ListInterface282) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ListInterface282) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ListInterface282) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ListInterface282) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ListInterface282) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ListInterface282) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ListInterface282) SetString(v string)   { panic("Unsupported operation") }
func (_ ListInterface282) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ListInterface282) Get(i int) types.Field {
	switch i {
	case 0:
		r.Interfaces = make([]Interface282Data, 0)

		w := ArrayInterface282DataWrapper{Target: &r.Interfaces}

		return w

	}
	panic("Unknown field index")
}

func (r *ListInterface282) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ListInterface282) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ListInterface282) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ListInterface282) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ListInterface282) HintSize(int)                     { panic("Unsupported operation") }
func (_ ListInterface282) Finalize()                        {}

func (_ ListInterface282) AvroCRC64Fingerprint() []byte {
	return []byte(ListInterface282AvroCRC64Fingerprint)
}

func (r ListInterface282) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["interfaces"], err = json.Marshal(r.Interfaces)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ListInterface282) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["interfaces"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Interfaces); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for interfaces")
	}
	return nil
}
