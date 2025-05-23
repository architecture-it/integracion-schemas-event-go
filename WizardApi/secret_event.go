// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     SecretEvent.avsc
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

type SecretEvent struct {
	Id int32 `json:"Id"`

	Key string `json:"Key"`

	AuditInfo AuditEvent `json:"AuditInfo"`
}

const SecretEventAvroCRC64Fingerprint = "\xe0\f\xa4e\x7f\xd8\x1c\x02"

func NewSecretEvent() SecretEvent {
	r := SecretEvent{}
	r.AuditInfo = NewAuditEvent()

	return r
}

func DeserializeSecretEvent(r io.Reader) (SecretEvent, error) {
	t := NewSecretEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSecretEventFromSchema(r io.Reader, schema string) (SecretEvent, error) {
	t := NewSecretEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSecretEvent(r SecretEvent, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Key, w)
	if err != nil {
		return err
	}
	err = writeAuditEvent(r.AuditInfo, w)
	if err != nil {
		return err
	}
	return err
}

func (r SecretEvent) Serialize(w io.Writer) error {
	return writeSecretEvent(r, w)
}

func (r SecretEvent) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Key\",\"type\":\"string\"},{\"name\":\"AuditInfo\",\"type\":{\"fields\":[{\"name\":\"CreateBy\",\"type\":\"string\"},{\"name\":\"CreateDate\",\"type\":{\"logicalType\":\"date\",\"type\":\"int\"}},{\"default\":null,\"name\":\"UpdateBy\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"UpdateDate\",\"type\":[\"null\",{\"logicalType\":\"date\",\"type\":\"int\"}]},{\"name\":\"Deleted\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"DeletedBy\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DeletedDate\",\"type\":[\"null\",{\"logicalType\":\"date\",\"type\":\"int\"}]}],\"name\":\"AuditEvent\",\"namespace\":\"Andreani.WizardApi.Events.Common\",\"type\":\"record\"}}],\"name\":\"Andreani.WizardApi.Events.Record.SecretEvent\",\"type\":\"record\"}"
}

func (r SecretEvent) SchemaName() string {
	return "Andreani.WizardApi.Events.Record.SecretEvent"
}

func (_ SecretEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ SecretEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ SecretEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ SecretEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ SecretEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ SecretEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ SecretEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ SecretEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *SecretEvent) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Key}

		return w

	case 2:
		r.AuditInfo = NewAuditEvent()

		w := types.Record{Target: &r.AuditInfo}

		return w

	}
	panic("Unknown field index")
}

func (r *SecretEvent) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *SecretEvent) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ SecretEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ SecretEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ SecretEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ SecretEvent) Finalize()                        {}

func (_ SecretEvent) AvroCRC64Fingerprint() []byte {
	return []byte(SecretEventAvroCRC64Fingerprint)
}

func (r SecretEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["Key"], err = json.Marshal(r.Key)
	if err != nil {
		return nil, err
	}
	output["AuditInfo"], err = json.Marshal(r.AuditInfo)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *SecretEvent) UnmarshalJSON(data []byte) error {
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
		return fmt.Errorf("no value specified for Key")
	}
	val = func() json.RawMessage {
		if v, ok := fields["AuditInfo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AuditInfo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for AuditInfo")
	}
	return nil
}
