// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     KafkaDemo.avsc
 */
package TestEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type KafkaDemo struct {
	Id int32 `json:"id"`

	Title string `json:"title"`

	Attendance int32 `json:"attendance"`

	DidYouUnderstand bool `json:"didYouUnderstand"`

	Me Person `json:"me"`

	Time int64 `json:"time"`
}

const KafkaDemoAvroCRC64Fingerprint = "Y\xb89\xd8S\xc1\xeaS"

func NewKafkaDemo() KafkaDemo {
	r := KafkaDemo{}
	r.Me = NewPerson()

	return r
}

func DeserializeKafkaDemo(r io.Reader) (KafkaDemo, error) {
	t := NewKafkaDemo()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeKafkaDemoFromSchema(r io.Reader, schema string) (KafkaDemo, error) {
	t := NewKafkaDemo()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeKafkaDemo(r KafkaDemo, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Title, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Attendance, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.DidYouUnderstand, w)
	if err != nil {
		return err
	}
	err = writePerson(r.Me, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Time, w)
	if err != nil {
		return err
	}
	return err
}

func (r KafkaDemo) Serialize(w io.Writer) error {
	return writeKafkaDemo(r, w)
}

func (r KafkaDemo) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":\"int\"},{\"name\":\"title\",\"type\":\"string\"},{\"name\":\"attendance\",\"type\":\"int\"},{\"name\":\"didYouUnderstand\",\"type\":\"boolean\"},{\"name\":\"me\",\"type\":{\"fields\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"surname\",\"type\":\"string\"},{\"name\":\"seniority\",\"type\":\"string\"},{\"name\":\"onSite\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"team\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"tl\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"boss\",\"type\":[\"null\",\"string\"]},{\"name\":\"members\",\"type\":{\"items\":\"string\",\"type\":\"array\"}}],\"name\":\"Team\",\"type\":\"record\"}]}],\"name\":\"Person\",\"namespace\":\"Andreani.Test.Events.Record.Common\",\"type\":\"record\"}},{\"name\":\"time\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}}],\"name\":\"Andreani.Test.Events.Record.KafkaDemo\",\"type\":\"record\"}"
}

func (r KafkaDemo) SchemaName() string {
	return "Andreani.Test.Events.Record.KafkaDemo"
}

func (_ KafkaDemo) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ KafkaDemo) SetInt(v int32)       { panic("Unsupported operation") }
func (_ KafkaDemo) SetLong(v int64)      { panic("Unsupported operation") }
func (_ KafkaDemo) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ KafkaDemo) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ KafkaDemo) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ KafkaDemo) SetString(v string)   { panic("Unsupported operation") }
func (_ KafkaDemo) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *KafkaDemo) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Title}

		return w

	case 2:
		w := types.Int{Target: &r.Attendance}

		return w

	case 3:
		w := types.Boolean{Target: &r.DidYouUnderstand}

		return w

	case 4:
		r.Me = NewPerson()

		w := types.Record{Target: &r.Me}

		return w

	case 5:
		w := types.Long{Target: &r.Time}

		return w

	}
	panic("Unknown field index")
}

func (r *KafkaDemo) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *KafkaDemo) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ KafkaDemo) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ KafkaDemo) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ KafkaDemo) HintSize(int)                     { panic("Unsupported operation") }
func (_ KafkaDemo) Finalize()                        {}

func (_ KafkaDemo) AvroCRC64Fingerprint() []byte {
	return []byte(KafkaDemoAvroCRC64Fingerprint)
}

func (r KafkaDemo) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["title"], err = json.Marshal(r.Title)
	if err != nil {
		return nil, err
	}
	output["attendance"], err = json.Marshal(r.Attendance)
	if err != nil {
		return nil, err
	}
	output["didYouUnderstand"], err = json.Marshal(r.DidYouUnderstand)
	if err != nil {
		return nil, err
	}
	output["me"], err = json.Marshal(r.Me)
	if err != nil {
		return nil, err
	}
	output["time"], err = json.Marshal(r.Time)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *KafkaDemo) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["title"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Title); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for title")
	}
	val = func() json.RawMessage {
		if v, ok := fields["attendance"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Attendance); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for attendance")
	}
	val = func() json.RawMessage {
		if v, ok := fields["didYouUnderstand"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DidYouUnderstand); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for didYouUnderstand")
	}
	val = func() json.RawMessage {
		if v, ok := fields["me"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Me); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for me")
	}
	val = func() json.RawMessage {
		if v, ok := fields["time"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Time); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for time")
	}
	return nil
}
