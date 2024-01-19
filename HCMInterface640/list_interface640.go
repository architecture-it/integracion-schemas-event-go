// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ListInterface640.avsc
 */
package HCMInterface640Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ListInterface640 struct {
	Interfaces []Interface640Data `json:"interfaces"`
}

const ListInterface640AvroCRC64Fingerprint = "HuȢ\x80\xefG\t"

func NewListInterface640() ListInterface640 {
	r := ListInterface640{}
	r.Interfaces = make([]Interface640Data, 0)

	return r
}

func DeserializeListInterface640(r io.Reader) (ListInterface640, error) {
	t := NewListInterface640()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeListInterface640FromSchema(r io.Reader, schema string) (ListInterface640, error) {
	t := NewListInterface640()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeListInterface640(r ListInterface640, w io.Writer) error {
	var err error
	err = writeArrayInterface640Data(r.Interfaces, w)
	if err != nil {
		return err
	}
	return err
}

func (r ListInterface640) Serialize(w io.Writer) error {
	return writeListInterface640(r, w)
}

func (r ListInterface640) Schema() string {
	return "{\"fields\":[{\"name\":\"interfaces\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Legajo\",\"type\":\"string\"},{\"name\":\"CausaBaja\",\"type\":\"string\"},{\"name\":\"FechaAlta\",\"type\":\"string\"},{\"name\":\"FechaBaja\",\"type\":\"string\"},{\"name\":\"Estado\",\"type\":\"string\"},{\"name\":\"AntigPSueldo\",\"type\":\"string\"},{\"name\":\"AntigPVacaciones\",\"type\":\"string\"},{\"name\":\"AntigPIndemnizacion\",\"type\":\"string\"},{\"name\":\"AntigPReal\",\"type\":\"string\"},{\"name\":\"FechaAltaReconocida\",\"type\":\"string\"}],\"name\":\"Interface640Data\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.HCMInterface640.Events.Record.ListInterface640\",\"type\":\"record\"}"
}

func (r ListInterface640) SchemaName() string {
	return "Andreani.HCMInterface640.Events.Record.ListInterface640"
}

func (_ ListInterface640) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ListInterface640) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ListInterface640) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ListInterface640) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ListInterface640) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ListInterface640) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ListInterface640) SetString(v string)   { panic("Unsupported operation") }
func (_ ListInterface640) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ListInterface640) Get(i int) types.Field {
	switch i {
	case 0:
		r.Interfaces = make([]Interface640Data, 0)

		w := ArrayInterface640DataWrapper{Target: &r.Interfaces}

		return w

	}
	panic("Unknown field index")
}

func (r *ListInterface640) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ListInterface640) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ListInterface640) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ListInterface640) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ListInterface640) HintSize(int)                     { panic("Unsupported operation") }
func (_ ListInterface640) Finalize()                        {}

func (_ ListInterface640) AvroCRC64Fingerprint() []byte {
	return []byte(ListInterface640AvroCRC64Fingerprint)
}

func (r ListInterface640) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["interfaces"], err = json.Marshal(r.Interfaces)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ListInterface640) UnmarshalJSON(data []byte) error {
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
