// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NovedadesCustom.avsc
 */
package NovedadesCustomEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Metadata struct {
	Key *UnionNullString `json:"Key"`

	Value *UnionNullString `json:"Value"`
}

const MetadataAvroCRC64Fingerprint = "\xfet\x00\xbdR\x1d\x04\xaa"

func NewMetadata() Metadata {
	r := Metadata{}
	r.Key = nil
	r.Value = nil
	return r
}

func DeserializeMetadata(r io.Reader) (Metadata, error) {
	t := NewMetadata()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeMetadataFromSchema(r io.Reader, schema string) (Metadata, error) {
	t := NewMetadata()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeMetadata(r Metadata, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Key, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Value, w)
	if err != nil {
		return err
	}
	return err
}

func (r Metadata) Serialize(w io.Writer) error {
	return writeMetadata(r, w)
}

func (r Metadata) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Key\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Value\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.NovedadesCustom.Events.Common.Metadata\",\"type\":\"record\"}"
}

func (r Metadata) SchemaName() string {
	return "Andreani.NovedadesCustom.Events.Common.Metadata"
}

func (_ Metadata) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Metadata) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Metadata) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Metadata) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Metadata) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Metadata) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Metadata) SetString(v string)   { panic("Unsupported operation") }
func (_ Metadata) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Metadata) Get(i int) types.Field {
	switch i {
	case 0:
		r.Key = NewUnionNullString()

		return r.Key
	case 1:
		r.Value = NewUnionNullString()

		return r.Value
	}
	panic("Unknown field index")
}

func (r *Metadata) SetDefault(i int) {
	switch i {
	case 0:
		r.Key = nil
		return
	case 1:
		r.Value = nil
		return
	}
	panic("Unknown field index")
}

func (r *Metadata) NullField(i int) {
	switch i {
	case 0:
		r.Key = nil
		return
	case 1:
		r.Value = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Metadata) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Metadata) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Metadata) HintSize(int)                     { panic("Unsupported operation") }
func (_ Metadata) Finalize()                        {}

func (_ Metadata) AvroCRC64Fingerprint() []byte {
	return []byte(MetadataAvroCRC64Fingerprint)
}

func (r Metadata) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Key"], err = json.Marshal(r.Key)
	if err != nil {
		return nil, err
	}
	output["Value"], err = json.Marshal(r.Value)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Metadata) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Key"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Key); err != nil {
			return err
		}
	} else {
		r.Key = NewUnionNullString()

		r.Key = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Value"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Value); err != nil {
			return err
		}
	} else {
		r.Value = NewUnionNullString()

		r.Value = nil
	}
	return nil
}
