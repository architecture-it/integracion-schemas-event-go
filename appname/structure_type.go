// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     StructureType.avsc
 */
package appnameEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type StructureType struct {
	Id int32 `json:"Id"`

	Description string `json:"Description"`
}

const StructureTypeAvroCRC64Fingerprint = "\xad\xe1}\xa8\x9e\xee\xbb&"

func NewStructureType() StructureType {
	r := StructureType{}
	return r
}

func DeserializeStructureType(r io.Reader) (StructureType, error) {
	t := NewStructureType()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeStructureTypeFromSchema(r io.Reader, schema string) (StructureType, error) {
	t := NewStructureType()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeStructureType(r StructureType, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Description, w)
	if err != nil {
		return err
	}
	return err
}

func (r StructureType) Serialize(w io.Writer) error {
	return writeStructureType(r, w)
}

func (r StructureType) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"}],\"name\":\"Andreani.RHpro.Events.Common.StructureType\",\"type\":\"record\"}"
}

func (r StructureType) SchemaName() string {
	return "Andreani.RHpro.Events.Common.StructureType"
}

func (_ StructureType) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ StructureType) SetInt(v int32)       { panic("Unsupported operation") }
func (_ StructureType) SetLong(v int64)      { panic("Unsupported operation") }
func (_ StructureType) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ StructureType) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ StructureType) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ StructureType) SetString(v string)   { panic("Unsupported operation") }
func (_ StructureType) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *StructureType) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Description}

		return w

	}
	panic("Unknown field index")
}

func (r *StructureType) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *StructureType) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ StructureType) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ StructureType) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ StructureType) HintSize(int)                     { panic("Unsupported operation") }
func (_ StructureType) Finalize()                        {}

func (_ StructureType) AvroCRC64Fingerprint() []byte {
	return []byte(StructureTypeAvroCRC64Fingerprint)
}

func (r StructureType) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["Description"], err = json.Marshal(r.Description)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *StructureType) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Description"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Description); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Description")
	}
	return nil
}
