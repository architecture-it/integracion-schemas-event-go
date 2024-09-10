// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     StructureModel.avsc
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

type StructureModel struct {
	Id int32 `json:"Id"`

	Description string `json:"Description"`
}

const StructureModelAvroCRC64Fingerprint = "FKw\xf6\x81lPF"

func NewStructureModel() StructureModel {
	r := StructureModel{}
	return r
}

func DeserializeStructureModel(r io.Reader) (StructureModel, error) {
	t := NewStructureModel()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeStructureModelFromSchema(r io.Reader, schema string) (StructureModel, error) {
	t := NewStructureModel()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeStructureModel(r StructureModel, w io.Writer) error {
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

func (r StructureModel) Serialize(w io.Writer) error {
	return writeStructureModel(r, w)
}

func (r StructureModel) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"}],\"name\":\"Andreani.RHpro.Events.Common.StructureModel\",\"type\":\"record\"}"
}

func (r StructureModel) SchemaName() string {
	return "Andreani.RHpro.Events.Common.StructureModel"
}

func (_ StructureModel) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ StructureModel) SetInt(v int32)       { panic("Unsupported operation") }
func (_ StructureModel) SetLong(v int64)      { panic("Unsupported operation") }
func (_ StructureModel) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ StructureModel) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ StructureModel) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ StructureModel) SetString(v string)   { panic("Unsupported operation") }
func (_ StructureModel) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *StructureModel) Get(i int) types.Field {
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

func (r *StructureModel) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *StructureModel) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ StructureModel) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ StructureModel) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ StructureModel) HintSize(int)                     { panic("Unsupported operation") }
func (_ StructureModel) Finalize()                        {}

func (_ StructureModel) AvroCRC64Fingerprint() []byte {
	return []byte(StructureModelAvroCRC64Fingerprint)
}

func (r StructureModel) MarshalJSON() ([]byte, error) {
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

func (r *StructureModel) UnmarshalJSON(data []byte) error {
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