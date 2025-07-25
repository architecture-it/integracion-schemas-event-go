// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EmailMessage.avsc
 */
package IAContactoEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EmailMessage struct {
	Id string `json:"Id"`

	Subject *UnionNullString `json:"Subject"`

	ReceivedDateTime *UnionNullLong `json:"ReceivedDateTime"`

	Sender *UnionNullSender `json:"Sender"`

	Body *UnionNullBody `json:"Body"`
}

const EmailMessageAvroCRC64Fingerprint = "y\xc4\xc5\xd3+\x90\xcd\xff"

func NewEmailMessage() EmailMessage {
	r := EmailMessage{}
	return r
}

func DeserializeEmailMessage(r io.Reader) (EmailMessage, error) {
	t := NewEmailMessage()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEmailMessageFromSchema(r io.Reader, schema string) (EmailMessage, error) {
	t := NewEmailMessage()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEmailMessage(r EmailMessage, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Subject, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.ReceivedDateTime, w)
	if err != nil {
		return err
	}
	err = writeUnionNullSender(r.Sender, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBody(r.Body, w)
	if err != nil {
		return err
	}
	return err
}

func (r EmailMessage) Serialize(w io.Writer) error {
	return writeEmailMessage(r, w)
}

func (r EmailMessage) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Subject\",\"type\":[\"null\",\"string\"]},{\"name\":\"ReceivedDateTime\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"Sender\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"EmailAddress\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Address\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Name\",\"type\":[\"null\",\"string\"]}],\"name\":\"EmailAddress\",\"type\":\"record\"}]}],\"name\":\"Sender\",\"type\":\"record\"}]},{\"name\":\"Body\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"ContentType\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Content\",\"type\":[\"null\",\"string\"]}],\"name\":\"Body\",\"type\":\"record\"}]}],\"name\":\"Andreani.IAContacto.Events.Record.EmailMessage\",\"type\":\"record\"}"
}

func (r EmailMessage) SchemaName() string {
	return "Andreani.IAContacto.Events.Record.EmailMessage"
}

func (_ EmailMessage) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EmailMessage) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EmailMessage) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EmailMessage) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EmailMessage) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EmailMessage) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EmailMessage) SetString(v string)   { panic("Unsupported operation") }
func (_ EmailMessage) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EmailMessage) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Id}

		return w

	case 1:
		r.Subject = NewUnionNullString()

		return r.Subject
	case 2:
		r.ReceivedDateTime = NewUnionNullLong()

		return r.ReceivedDateTime
	case 3:
		r.Sender = NewUnionNullSender()

		return r.Sender
	case 4:
		r.Body = NewUnionNullBody()

		return r.Body
	}
	panic("Unknown field index")
}

func (r *EmailMessage) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EmailMessage) NullField(i int) {
	switch i {
	case 1:
		r.Subject = nil
		return
	case 2:
		r.ReceivedDateTime = nil
		return
	case 3:
		r.Sender = nil
		return
	case 4:
		r.Body = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ EmailMessage) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EmailMessage) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EmailMessage) HintSize(int)                     { panic("Unsupported operation") }
func (_ EmailMessage) Finalize()                        {}

func (_ EmailMessage) AvroCRC64Fingerprint() []byte {
	return []byte(EmailMessageAvroCRC64Fingerprint)
}

func (r EmailMessage) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["Subject"], err = json.Marshal(r.Subject)
	if err != nil {
		return nil, err
	}
	output["ReceivedDateTime"], err = json.Marshal(r.ReceivedDateTime)
	if err != nil {
		return nil, err
	}
	output["Sender"], err = json.Marshal(r.Sender)
	if err != nil {
		return nil, err
	}
	output["Body"], err = json.Marshal(r.Body)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EmailMessage) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["Subject"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Subject); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Subject")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ReceivedDateTime"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ReceivedDateTime); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ReceivedDateTime")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Sender"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sender); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Sender")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Body"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Body); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Body")
	}
	return nil
}
