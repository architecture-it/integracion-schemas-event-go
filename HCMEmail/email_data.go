// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ListEmail.avsc
 */
package HCMEmailEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EmailData struct {
	Subject string `json:"Subject"`

	Body string `json:"Body"`

	Attachment string `json:"Attachment"`

	ToAddresses string `json:"ToAddresses"`

	FromAddress string `json:"FromAddress"`
}

const EmailDataAvroCRC64Fingerprint = "\x9d\x17e\x8a\xf6|\xb5c"

func NewEmailData() EmailData {
	r := EmailData{}
	return r
}

func DeserializeEmailData(r io.Reader) (EmailData, error) {
	t := NewEmailData()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEmailDataFromSchema(r io.Reader, schema string) (EmailData, error) {
	t := NewEmailData()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEmailData(r EmailData, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Subject, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Body, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Attachment, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.ToAddresses, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FromAddress, w)
	if err != nil {
		return err
	}
	return err
}

func (r EmailData) Serialize(w io.Writer) error {
	return writeEmailData(r, w)
}

func (r EmailData) Schema() string {
	return "{\"fields\":[{\"name\":\"Subject\",\"type\":\"string\"},{\"name\":\"Body\",\"type\":\"string\"},{\"name\":\"Attachment\",\"type\":\"string\"},{\"name\":\"ToAddresses\",\"type\":\"string\"},{\"name\":\"FromAddress\",\"type\":\"string\"}],\"name\":\"Andreani.HCMEmail.Events.Record.EmailData\",\"type\":\"record\"}"
}

func (r EmailData) SchemaName() string {
	return "Andreani.HCMEmail.Events.Record.EmailData"
}

func (_ EmailData) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EmailData) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EmailData) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EmailData) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EmailData) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EmailData) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EmailData) SetString(v string)   { panic("Unsupported operation") }
func (_ EmailData) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EmailData) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Subject}

		return w

	case 1:
		w := types.String{Target: &r.Body}

		return w

	case 2:
		w := types.String{Target: &r.Attachment}

		return w

	case 3:
		w := types.String{Target: &r.ToAddresses}

		return w

	case 4:
		w := types.String{Target: &r.FromAddress}

		return w

	}
	panic("Unknown field index")
}

func (r *EmailData) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EmailData) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EmailData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EmailData) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EmailData) HintSize(int)                     { panic("Unsupported operation") }
func (_ EmailData) Finalize()                        {}

func (_ EmailData) AvroCRC64Fingerprint() []byte {
	return []byte(EmailDataAvroCRC64Fingerprint)
}

func (r EmailData) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Subject"], err = json.Marshal(r.Subject)
	if err != nil {
		return nil, err
	}
	output["Body"], err = json.Marshal(r.Body)
	if err != nil {
		return nil, err
	}
	output["Attachment"], err = json.Marshal(r.Attachment)
	if err != nil {
		return nil, err
	}
	output["ToAddresses"], err = json.Marshal(r.ToAddresses)
	if err != nil {
		return nil, err
	}
	output["FromAddress"], err = json.Marshal(r.FromAddress)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EmailData) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
	val = func() json.RawMessage {
		if v, ok := fields["Attachment"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Attachment); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Attachment")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ToAddresses"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ToAddresses); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ToAddresses")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FromAddress"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FromAddress); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FromAddress")
	}
	return nil
}
