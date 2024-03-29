// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     RepublicOrder.avsc
 */
package RepublicOrder_SchemeEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type RepublicOrder struct {
	Ids []string `json:"ids"`

	User string `json:"user"`
}

const RepublicOrderAvroCRC64Fingerprint = "5Z;A\xd6\x16\\\x95"

func NewRepublicOrder() RepublicOrder {
	r := RepublicOrder{}
	r.Ids = make([]string, 0)

	return r
}

func DeserializeRepublicOrder(r io.Reader) (RepublicOrder, error) {
	t := NewRepublicOrder()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRepublicOrderFromSchema(r io.Reader, schema string) (RepublicOrder, error) {
	t := NewRepublicOrder()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRepublicOrder(r RepublicOrder, w io.Writer) error {
	var err error
	err = writeArrayString(r.Ids, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.User, w)
	if err != nil {
		return err
	}
	return err
}

func (r RepublicOrder) Serialize(w io.Writer) error {
	return writeRepublicOrder(r, w)
}

func (r RepublicOrder) Schema() string {
	return "{\"fields\":[{\"name\":\"ids\",\"type\":{\"items\":\"string\",\"type\":\"array\"}},{\"name\":\"user\",\"type\":\"string\"}],\"name\":\"Andreani.Kafkadeadline.RepublicOrder.Scheme.RepublicOrder\",\"type\":\"record\"}"
}

func (r RepublicOrder) SchemaName() string {
	return "Andreani.Kafkadeadline.RepublicOrder.Scheme.RepublicOrder"
}

func (_ RepublicOrder) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ RepublicOrder) SetInt(v int32)       { panic("Unsupported operation") }
func (_ RepublicOrder) SetLong(v int64)      { panic("Unsupported operation") }
func (_ RepublicOrder) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ RepublicOrder) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ RepublicOrder) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ RepublicOrder) SetString(v string)   { panic("Unsupported operation") }
func (_ RepublicOrder) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RepublicOrder) Get(i int) types.Field {
	switch i {
	case 0:
		r.Ids = make([]string, 0)

		w := ArrayStringWrapper{Target: &r.Ids}

		return w

	case 1:
		w := types.String{Target: &r.User}

		return w

	}
	panic("Unknown field index")
}

func (r *RepublicOrder) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *RepublicOrder) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ RepublicOrder) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ RepublicOrder) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ RepublicOrder) HintSize(int)                     { panic("Unsupported operation") }
func (_ RepublicOrder) Finalize()                        {}

func (_ RepublicOrder) AvroCRC64Fingerprint() []byte {
	return []byte(RepublicOrderAvroCRC64Fingerprint)
}

func (r RepublicOrder) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["ids"], err = json.Marshal(r.Ids)
	if err != nil {
		return nil, err
	}
	output["user"], err = json.Marshal(r.User)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *RepublicOrder) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["ids"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ids); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ids")
	}
	val = func() json.RawMessage {
		if v, ok := fields["user"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.User); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for user")
	}
	return nil
}
