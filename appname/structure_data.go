// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     StructureData.avsc
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

type StructureData struct {
	StructureId int32 `json:"StructureId"`

	Id int32 `json:"Id"`

	DateFrom int64 `json:"DateFrom"`

	DateTo int64 `json:"DateTo"`

	Description string `json:"Description"`

	ExternalId *UnionNullString `json:"ExternalId"`

	Type *UnionNullStructureType `json:"Type"`

	Reason *UnionNullStructureReason `json:"Reason"`

	Model *UnionNullStructureModel `json:"Model"`
}

const StructureDataAvroCRC64Fingerprint = "g7~\xe7\xb4'\xab$"

func NewStructureData() StructureData {
	r := StructureData{}
	r.ExternalId = nil
	r.Type = nil
	r.Reason = nil
	r.Model = nil
	return r
}

func DeserializeStructureData(r io.Reader) (StructureData, error) {
	t := NewStructureData()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeStructureDataFromSchema(r io.Reader, schema string) (StructureData, error) {
	t := NewStructureData()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeStructureData(r StructureData, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.StructureId, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.DateFrom, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.DateTo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Description, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ExternalId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullStructureType(r.Type, w)
	if err != nil {
		return err
	}
	err = writeUnionNullStructureReason(r.Reason, w)
	if err != nil {
		return err
	}
	err = writeUnionNullStructureModel(r.Model, w)
	if err != nil {
		return err
	}
	return err
}

func (r StructureData) Serialize(w io.Writer) error {
	return writeStructureData(r, w)
}

func (r StructureData) Schema() string {
	return "{\"fields\":[{\"name\":\"StructureId\",\"type\":\"int\"},{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"DateFrom\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"DateTo\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"Description\",\"type\":\"string\"},{\"default\":null,\"name\":\"ExternalId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Type\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"}],\"name\":\"StructureType\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Reason\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"},{\"name\":\"Reason\",\"type\":\"string\"}],\"name\":\"StructureReason\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Model\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Description\",\"type\":\"string\"}],\"name\":\"StructureModel\",\"type\":\"record\"}]}],\"name\":\"Andreani.RHpro.Events.Common.StructureData\",\"type\":\"record\"}"
}

func (r StructureData) SchemaName() string {
	return "Andreani.RHpro.Events.Common.StructureData"
}

func (_ StructureData) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ StructureData) SetInt(v int32)       { panic("Unsupported operation") }
func (_ StructureData) SetLong(v int64)      { panic("Unsupported operation") }
func (_ StructureData) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ StructureData) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ StructureData) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ StructureData) SetString(v string)   { panic("Unsupported operation") }
func (_ StructureData) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *StructureData) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.StructureId}

		return w

	case 1:
		w := types.Int{Target: &r.Id}

		return w

	case 2:
		w := types.Long{Target: &r.DateFrom}

		return w

	case 3:
		w := types.Long{Target: &r.DateTo}

		return w

	case 4:
		w := types.String{Target: &r.Description}

		return w

	case 5:
		r.ExternalId = NewUnionNullString()

		return r.ExternalId
	case 6:
		r.Type = NewUnionNullStructureType()

		return r.Type
	case 7:
		r.Reason = NewUnionNullStructureReason()

		return r.Reason
	case 8:
		r.Model = NewUnionNullStructureModel()

		return r.Model
	}
	panic("Unknown field index")
}

func (r *StructureData) SetDefault(i int) {
	switch i {
	case 5:
		r.ExternalId = nil
		return
	case 6:
		r.Type = nil
		return
	case 7:
		r.Reason = nil
		return
	case 8:
		r.Model = nil
		return
	}
	panic("Unknown field index")
}

func (r *StructureData) NullField(i int) {
	switch i {
	case 5:
		r.ExternalId = nil
		return
	case 6:
		r.Type = nil
		return
	case 7:
		r.Reason = nil
		return
	case 8:
		r.Model = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ StructureData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ StructureData) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ StructureData) HintSize(int)                     { panic("Unsupported operation") }
func (_ StructureData) Finalize()                        {}

func (_ StructureData) AvroCRC64Fingerprint() []byte {
	return []byte(StructureDataAvroCRC64Fingerprint)
}

func (r StructureData) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["StructureId"], err = json.Marshal(r.StructureId)
	if err != nil {
		return nil, err
	}
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["DateFrom"], err = json.Marshal(r.DateFrom)
	if err != nil {
		return nil, err
	}
	output["DateTo"], err = json.Marshal(r.DateTo)
	if err != nil {
		return nil, err
	}
	output["Description"], err = json.Marshal(r.Description)
	if err != nil {
		return nil, err
	}
	output["ExternalId"], err = json.Marshal(r.ExternalId)
	if err != nil {
		return nil, err
	}
	output["Type"], err = json.Marshal(r.Type)
	if err != nil {
		return nil, err
	}
	output["Reason"], err = json.Marshal(r.Reason)
	if err != nil {
		return nil, err
	}
	output["Model"], err = json.Marshal(r.Model)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *StructureData) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["StructureId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StructureId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for StructureId")
	}
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
		if v, ok := fields["DateFrom"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DateFrom); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DateFrom")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DateTo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DateTo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DateTo")
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
		if v, ok := fields["ExternalId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ExternalId); err != nil {
			return err
		}
	} else {
		r.ExternalId = NewUnionNullString()

		r.ExternalId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Type"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Type); err != nil {
			return err
		}
	} else {
		r.Type = NewUnionNullStructureType()

		r.Type = nil
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
		r.Reason = NewUnionNullStructureReason()

		r.Reason = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Model"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Model); err != nil {
			return err
		}
	} else {
		r.Model = NewUnionNullStructureModel()

		r.Model = nil
	}
	return nil
}
