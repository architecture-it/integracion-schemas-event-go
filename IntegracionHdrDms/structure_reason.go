// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     StructureReason.avsc
 */
package IntegracionHdrDmsEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type StructureReason struct {
	Id int32 `json:"Id"`

	Description string `json:"Description"`

	Reason string `json:"Reason"`
}

const StructureReasonAvroCRC64Fingerprint = "~N]-\a\xa9\xc9L"

func NewStructureReason() StructureReason {
	r := StructureReason{}
	return r
}

func DeserializeStructureReason(r io.Reader) (StructureReason, error) {
	t := NewStructureReason()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeStructureReasonFromSchema(r io.Reader, schema string) (StructureReason, error) {
	t := NewStructureReason()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeStructureReason(r StructureReason, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Description, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Reason, w)
	if err != nil {
		return err
	}
	return err
}

func (r StructureReason) Serialize(w io.Writer) error {
	return writeStructureReason(r, w)
}

func (r StructureReason) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"},{\"name\":\"Reason\",\"type\":\"string\"}],\"name\":\"Andreani.IntegracionHdrDms.Events.Common.StructureReason\",\"type\":\"record\"}"
}

func (r StructureReason) SchemaName() string {
	return "Andreani.IntegracionHdrDms.Events.Common.StructureReason"
}

func (_ StructureReason) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ StructureReason) SetInt(v int32)       { panic("Unsupported operation") }
func (_ StructureReason) SetLong(v int64)      { panic("Unsupported operation") }
func (_ StructureReason) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ StructureReason) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ StructureReason) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ StructureReason) SetString(v string)   { panic("Unsupported operation") }
func (_ StructureReason) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *StructureReason) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Description}

		return w

	case 2:
		w := types.String{Target: &r.Reason}

		return w

	}
	panic("Unknown field index")
}

func (r *StructureReason) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *StructureReason) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ StructureReason) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ StructureReason) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ StructureReason) HintSize(int)                     { panic("Unsupported operation") }
func (_ StructureReason) Finalize()                        {}

func (_ StructureReason) AvroCRC64Fingerprint() []byte {
	return []byte(StructureReasonAvroCRC64Fingerprint)
}

func (r StructureReason) MarshalJSON() ([]byte, error) {
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
	output["Reason"], err = json.Marshal(r.Reason)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *StructureReason) UnmarshalJSON(data []byte) error {
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
	val = func() json.RawMessage {
		if v, ok := fields["Reason"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Reason); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Reason")
	}
	return nil
}
