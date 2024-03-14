// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ListVismaInterfaceError.avsc
 */
package VismaInterfaceErrorEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ListVismaInterfaceError struct {
	Interface *UnionNullString `json:"Interface"`

	Mensajes []VismaInterfaceErrorData `json:"Mensajes"`

	VismaErrorMessages []VismaErrorMessageData `json:"VismaErrorMessages"`
}

const ListVismaInterfaceErrorAvroCRC64Fingerprint = "\xc7\xeb^4\x90s_r"

func NewListVismaInterfaceError() ListVismaInterfaceError {
	r := ListVismaInterfaceError{}
	r.Interface = nil
	r.Mensajes = make([]VismaInterfaceErrorData, 0)

	r.VismaErrorMessages = make([]VismaErrorMessageData, 0)

	return r
}

func DeserializeListVismaInterfaceError(r io.Reader) (ListVismaInterfaceError, error) {
	t := NewListVismaInterfaceError()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeListVismaInterfaceErrorFromSchema(r io.Reader, schema string) (ListVismaInterfaceError, error) {
	t := NewListVismaInterfaceError()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeListVismaInterfaceError(r ListVismaInterfaceError, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Interface, w)
	if err != nil {
		return err
	}
	err = writeArrayVismaInterfaceErrorData(r.Mensajes, w)
	if err != nil {
		return err
	}
	err = writeArrayVismaErrorMessageData(r.VismaErrorMessages, w)
	if err != nil {
		return err
	}
	return err
}

func (r ListVismaInterfaceError) Serialize(w io.Writer) error {
	return writeListVismaInterfaceError(r, w)
}

func (r ListVismaInterfaceError) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Interface\",\"type\":[\"null\",\"string\"]},{\"name\":\"Mensajes\",\"type\":{\"items\":{\"fields\":[{\"default\":null,\"name\":\"Mensaje\",\"type\":[\"null\",\"string\"]}],\"name\":\"VismaInterfaceErrorData\",\"type\":\"record\"},\"type\":\"array\"}},{\"name\":\"VismaErrorMessages\",\"type\":{\"items\":{\"fields\":[{\"default\":null,\"name\":\"EmployeeId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ErrorMessage\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"RowCsv\",\"type\":[\"null\",\"string\"]}],\"name\":\"VismaErrorMessageData\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.VismaInterfaceError.Events.Record.ListVismaInterfaceError\",\"type\":\"record\"}"
}

func (r ListVismaInterfaceError) SchemaName() string {
	return "Andreani.VismaInterfaceError.Events.Record.ListVismaInterfaceError"
}

func (_ ListVismaInterfaceError) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ListVismaInterfaceError) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ListVismaInterfaceError) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ListVismaInterfaceError) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ListVismaInterfaceError) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ListVismaInterfaceError) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ListVismaInterfaceError) SetString(v string)   { panic("Unsupported operation") }
func (_ ListVismaInterfaceError) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ListVismaInterfaceError) Get(i int) types.Field {
	switch i {
	case 0:
		r.Interface = NewUnionNullString()

		return r.Interface
	case 1:
		r.Mensajes = make([]VismaInterfaceErrorData, 0)

		w := ArrayVismaInterfaceErrorDataWrapper{Target: &r.Mensajes}

		return w

	case 2:
		r.VismaErrorMessages = make([]VismaErrorMessageData, 0)

		w := ArrayVismaErrorMessageDataWrapper{Target: &r.VismaErrorMessages}

		return w

	}
	panic("Unknown field index")
}

func (r *ListVismaInterfaceError) SetDefault(i int) {
	switch i {
	case 0:
		r.Interface = nil
		return
	}
	panic("Unknown field index")
}

func (r *ListVismaInterfaceError) NullField(i int) {
	switch i {
	case 0:
		r.Interface = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ListVismaInterfaceError) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ListVismaInterfaceError) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ListVismaInterfaceError) HintSize(int)                     { panic("Unsupported operation") }
func (_ ListVismaInterfaceError) Finalize()                        {}

func (_ ListVismaInterfaceError) AvroCRC64Fingerprint() []byte {
	return []byte(ListVismaInterfaceErrorAvroCRC64Fingerprint)
}

func (r ListVismaInterfaceError) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Interface"], err = json.Marshal(r.Interface)
	if err != nil {
		return nil, err
	}
	output["Mensajes"], err = json.Marshal(r.Mensajes)
	if err != nil {
		return nil, err
	}
	output["VismaErrorMessages"], err = json.Marshal(r.VismaErrorMessages)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ListVismaInterfaceError) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Interface"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Interface); err != nil {
			return err
		}
	} else {
		r.Interface = NewUnionNullString()

		r.Interface = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Mensajes"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Mensajes); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Mensajes")
	}
	val = func() json.RawMessage {
		if v, ok := fields["VismaErrorMessages"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.VismaErrorMessages); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for VismaErrorMessages")
	}
	return nil
}
