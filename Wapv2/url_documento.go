// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     UrlDocumento.avsc
 */
package Wapv2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type UrlDocumento struct {
	Url string `json:"url"`
}

const UrlDocumentoAvroCRC64Fingerprint = "\xc8D\x88\xe4\xd1\x1a\x94."

func NewUrlDocumento() UrlDocumento {
	r := UrlDocumento{}
	return r
}

func DeserializeUrlDocumento(r io.Reader) (UrlDocumento, error) {
	t := NewUrlDocumento()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeUrlDocumentoFromSchema(r io.Reader, schema string) (UrlDocumento, error) {
	t := NewUrlDocumento()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeUrlDocumento(r UrlDocumento, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Url, w)
	if err != nil {
		return err
	}
	return err
}

func (r UrlDocumento) Serialize(w io.Writer) error {
	return writeUrlDocumento(r, w)
}

func (r UrlDocumento) Schema() string {
	return "{\"fields\":[{\"name\":\"url\",\"type\":\"string\"}],\"name\":\"Andreani.Wapv2.Events.Record.UrlDocumento\",\"type\":\"record\"}"
}

func (r UrlDocumento) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.UrlDocumento"
}

func (_ UrlDocumento) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ UrlDocumento) SetInt(v int32)       { panic("Unsupported operation") }
func (_ UrlDocumento) SetLong(v int64)      { panic("Unsupported operation") }
func (_ UrlDocumento) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ UrlDocumento) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ UrlDocumento) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ UrlDocumento) SetString(v string)   { panic("Unsupported operation") }
func (_ UrlDocumento) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *UrlDocumento) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Url}

		return w

	}
	panic("Unknown field index")
}

func (r *UrlDocumento) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *UrlDocumento) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ UrlDocumento) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ UrlDocumento) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ UrlDocumento) HintSize(int)                     { panic("Unsupported operation") }
func (_ UrlDocumento) Finalize()                        {}

func (_ UrlDocumento) AvroCRC64Fingerprint() []byte {
	return []byte(UrlDocumentoAvroCRC64Fingerprint)
}

func (r UrlDocumento) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["url"], err = json.Marshal(r.Url)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *UrlDocumento) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["url"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Url); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for url")
	}
	return nil
}
