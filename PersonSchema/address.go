// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Address.avsc
 */
package PersonSchemaEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Address struct {
	StreetName string `json:"StreetName"`

	City string `json:"City"`
}

const AddressAvroCRC64Fingerprint = "&\x83p(\xd1\x00u\xd0"

func NewAddress() Address {
	r := Address{}
	return r
}

func DeserializeAddress(r io.Reader) (Address, error) {
	t := NewAddress()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAddressFromSchema(r io.Reader, schema string) (Address, error) {
	t := NewAddress()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAddress(r Address, w io.Writer) error {
	var err error
	err = vm.WriteString(r.StreetName, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.City, w)
	if err != nil {
		return err
	}
	return err
}

func (r Address) Serialize(w io.Writer) error {
	return writeAddress(r, w)
}

func (r Address) Schema() string {
	return "{\"fields\":[{\"name\":\"StreetName\",\"type\":\"string\"},{\"name\":\"City\",\"type\":\"string\"}],\"name\":\"Andreani.RHpro.Events.Record.Address\",\"type\":\"record\"}"
}

func (r Address) SchemaName() string {
	return "Andreani.RHpro.Events.Record.Address"
}

func (_ Address) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Address) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Address) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Address) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Address) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Address) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Address) SetString(v string)   { panic("Unsupported operation") }
func (_ Address) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Address) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.StreetName}

		return w

	case 1:
		w := types.String{Target: &r.City}

		return w

	}
	panic("Unknown field index")
}

func (r *Address) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Address) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Address) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Address) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Address) HintSize(int)                     { panic("Unsupported operation") }
func (_ Address) Finalize()                        {}

func (_ Address) AvroCRC64Fingerprint() []byte {
	return []byte(AddressAvroCRC64Fingerprint)
}

func (r Address) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["StreetName"], err = json.Marshal(r.StreetName)
	if err != nil {
		return nil, err
	}
	output["City"], err = json.Marshal(r.City)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Address) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["StreetName"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StreetName); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for StreetName")
	}
	val = func() json.RawMessage {
		if v, ok := fields["City"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.City); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for City")
	}
	return nil
}
