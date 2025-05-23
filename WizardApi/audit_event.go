// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     VariableEvent.avsc
 */
package WizardApiEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type AuditEvent struct {
	CreateBy string `json:"CreateBy"`

	CreateDate int32 `json:"CreateDate"`

	UpdateBy *UnionNullString `json:"UpdateBy"`

	UpdateDate *UnionNullInt `json:"UpdateDate"`

	Deleted bool `json:"Deleted"`

	DeletedBy *UnionNullString `json:"DeletedBy"`

	DeletedDate *UnionNullInt `json:"DeletedDate"`
}

const AuditEventAvroCRC64Fingerprint = "#3\x13\xb3\xfa\x93\x95\xf7"

func NewAuditEvent() AuditEvent {
	r := AuditEvent{}
	r.UpdateBy = nil
	r.UpdateDate = nil
	r.DeletedBy = nil
	r.DeletedDate = nil
	return r
}

func DeserializeAuditEvent(r io.Reader) (AuditEvent, error) {
	t := NewAuditEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAuditEventFromSchema(r io.Reader, schema string) (AuditEvent, error) {
	t := NewAuditEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAuditEvent(r AuditEvent, w io.Writer) error {
	var err error
	err = vm.WriteString(r.CreateBy, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.CreateDate, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.UpdateBy, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.UpdateDate, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Deleted, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DeletedBy, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.DeletedDate, w)
	if err != nil {
		return err
	}
	return err
}

func (r AuditEvent) Serialize(w io.Writer) error {
	return writeAuditEvent(r, w)
}

func (r AuditEvent) Schema() string {
	return "{\"fields\":[{\"name\":\"CreateBy\",\"type\":\"string\"},{\"name\":\"CreateDate\",\"type\":{\"logicalType\":\"date\",\"type\":\"int\"}},{\"default\":null,\"name\":\"UpdateBy\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"UpdateDate\",\"type\":[\"null\",{\"logicalType\":\"date\",\"type\":\"int\"}]},{\"name\":\"Deleted\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"DeletedBy\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DeletedDate\",\"type\":[\"null\",{\"logicalType\":\"date\",\"type\":\"int\"}]}],\"name\":\"Andreani.WizardApi.Events.Common.AuditEvent\",\"type\":\"record\"}"
}

func (r AuditEvent) SchemaName() string {
	return "Andreani.WizardApi.Events.Common.AuditEvent"
}

func (_ AuditEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ AuditEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ AuditEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ AuditEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ AuditEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ AuditEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ AuditEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ AuditEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AuditEvent) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.CreateBy}

		return w

	case 1:
		w := types.Int{Target: &r.CreateDate}

		return w

	case 2:
		r.UpdateBy = NewUnionNullString()

		return r.UpdateBy
	case 3:
		r.UpdateDate = NewUnionNullInt()

		return r.UpdateDate
	case 4:
		w := types.Boolean{Target: &r.Deleted}

		return w

	case 5:
		r.DeletedBy = NewUnionNullString()

		return r.DeletedBy
	case 6:
		r.DeletedDate = NewUnionNullInt()

		return r.DeletedDate
	}
	panic("Unknown field index")
}

func (r *AuditEvent) SetDefault(i int) {
	switch i {
	case 2:
		r.UpdateBy = nil
		return
	case 3:
		r.UpdateDate = nil
		return
	case 5:
		r.DeletedBy = nil
		return
	case 6:
		r.DeletedDate = nil
		return
	}
	panic("Unknown field index")
}

func (r *AuditEvent) NullField(i int) {
	switch i {
	case 2:
		r.UpdateBy = nil
		return
	case 3:
		r.UpdateDate = nil
		return
	case 5:
		r.DeletedBy = nil
		return
	case 6:
		r.DeletedDate = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ AuditEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ AuditEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ AuditEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ AuditEvent) Finalize()                        {}

func (_ AuditEvent) AvroCRC64Fingerprint() []byte {
	return []byte(AuditEventAvroCRC64Fingerprint)
}

func (r AuditEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["CreateBy"], err = json.Marshal(r.CreateBy)
	if err != nil {
		return nil, err
	}
	output["CreateDate"], err = json.Marshal(r.CreateDate)
	if err != nil {
		return nil, err
	}
	output["UpdateBy"], err = json.Marshal(r.UpdateBy)
	if err != nil {
		return nil, err
	}
	output["UpdateDate"], err = json.Marshal(r.UpdateDate)
	if err != nil {
		return nil, err
	}
	output["Deleted"], err = json.Marshal(r.Deleted)
	if err != nil {
		return nil, err
	}
	output["DeletedBy"], err = json.Marshal(r.DeletedBy)
	if err != nil {
		return nil, err
	}
	output["DeletedDate"], err = json.Marshal(r.DeletedDate)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *AuditEvent) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["CreateBy"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CreateBy); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CreateBy")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CreateDate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CreateDate); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CreateDate")
	}
	val = func() json.RawMessage {
		if v, ok := fields["UpdateBy"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UpdateBy); err != nil {
			return err
		}
	} else {
		r.UpdateBy = NewUnionNullString()

		r.UpdateBy = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["UpdateDate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UpdateDate); err != nil {
			return err
		}
	} else {
		r.UpdateDate = NewUnionNullInt()

		r.UpdateDate = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Deleted"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Deleted); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Deleted")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DeletedBy"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DeletedBy); err != nil {
			return err
		}
	} else {
		r.DeletedBy = NewUnionNullString()

		r.DeletedBy = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DeletedDate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DeletedDate); err != nil {
			return err
		}
	} else {
		r.DeletedDate = NewUnionNullInt()

		r.DeletedDate = nil
	}
	return nil
}
