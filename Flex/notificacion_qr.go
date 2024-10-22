// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NotificacionQR.avsc
 */
package FlexEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type NotificacionQR struct {
	OrdenId string `json:"ordenId"`

	Status *UnionNullString `json:"status"`
}

const NotificacionQRAvroCRC64Fingerprint = "\xd8l\xd5\x16G\xaa\xad\xa5"

func NewNotificacionQR() NotificacionQR {
	r := NotificacionQR{}
	return r
}

func DeserializeNotificacionQR(r io.Reader) (NotificacionQR, error) {
	t := NewNotificacionQR()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNotificacionQRFromSchema(r io.Reader, schema string) (NotificacionQR, error) {
	t := NewNotificacionQR()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNotificacionQR(r NotificacionQR, w io.Writer) error {
	var err error
	err = vm.WriteString(r.OrdenId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Status, w)
	if err != nil {
		return err
	}
	return err
}

func (r NotificacionQR) Serialize(w io.Writer) error {
	return writeNotificacionQR(r, w)
}

func (r NotificacionQR) Schema() string {
	return "{\"fields\":[{\"name\":\"ordenId\",\"type\":\"string\"},{\"name\":\"status\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.Flex.Events.Record.NotificacionQR\",\"type\":\"record\"}"
}

func (r NotificacionQR) SchemaName() string {
	return "Andreani.Flex.Events.Record.NotificacionQR"
}

func (_ NotificacionQR) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NotificacionQR) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NotificacionQR) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NotificacionQR) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NotificacionQR) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NotificacionQR) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NotificacionQR) SetString(v string)   { panic("Unsupported operation") }
func (_ NotificacionQR) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NotificacionQR) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.OrdenId}

		return w

	case 1:
		r.Status = NewUnionNullString()

		return r.Status
	}
	panic("Unknown field index")
}

func (r *NotificacionQR) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NotificacionQR) NullField(i int) {
	switch i {
	case 1:
		r.Status = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ NotificacionQR) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NotificacionQR) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NotificacionQR) HintSize(int)                     { panic("Unsupported operation") }
func (_ NotificacionQR) Finalize()                        {}

func (_ NotificacionQR) AvroCRC64Fingerprint() []byte {
	return []byte(NotificacionQRAvroCRC64Fingerprint)
}

func (r NotificacionQR) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["ordenId"], err = json.Marshal(r.OrdenId)
	if err != nil {
		return nil, err
	}
	output["status"], err = json.Marshal(r.Status)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NotificacionQR) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["ordenId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ordenId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["status"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Status); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for status")
	}
	return nil
}