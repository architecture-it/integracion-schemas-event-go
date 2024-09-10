// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Authentication.avsc
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

type Authentication struct {
	Token string `json:"token"`

	Tenant string `json:"tenant"`
}

const AuthenticationAvroCRC64Fingerprint = "\xec9\xddC\xd3[,\xa3"

func NewAuthentication() Authentication {
	r := Authentication{}
	return r
}

func DeserializeAuthentication(r io.Reader) (Authentication, error) {
	t := NewAuthentication()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAuthenticationFromSchema(r io.Reader, schema string) (Authentication, error) {
	t := NewAuthentication()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAuthentication(r Authentication, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Token, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Tenant, w)
	if err != nil {
		return err
	}
	return err
}

func (r Authentication) Serialize(w io.Writer) error {
	return writeAuthentication(r, w)
}

func (r Authentication) Schema() string {
	return "{\"fields\":[{\"name\":\"token\",\"type\":\"string\"},{\"name\":\"tenant\",\"type\":\"string\"}],\"name\":\"Andreani.RHpro.Events.Record.Authentication\",\"type\":\"record\"}"
}

func (r Authentication) SchemaName() string {
	return "Andreani.RHpro.Events.Record.Authentication"
}

func (_ Authentication) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Authentication) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Authentication) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Authentication) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Authentication) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Authentication) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Authentication) SetString(v string)   { panic("Unsupported operation") }
func (_ Authentication) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Authentication) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Token}

		return w

	case 1:
		w := types.String{Target: &r.Tenant}

		return w

	}
	panic("Unknown field index")
}

func (r *Authentication) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Authentication) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Authentication) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Authentication) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Authentication) HintSize(int)                     { panic("Unsupported operation") }
func (_ Authentication) Finalize()                        {}

func (_ Authentication) AvroCRC64Fingerprint() []byte {
	return []byte(AuthenticationAvroCRC64Fingerprint)
}

func (r Authentication) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["token"], err = json.Marshal(r.Token)
	if err != nil {
		return nil, err
	}
	output["tenant"], err = json.Marshal(r.Tenant)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Authentication) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["token"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Token); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for token")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tenant"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Tenant); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tenant")
	}
	return nil
}