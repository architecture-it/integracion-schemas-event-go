// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Repository.avsc
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

type GithubUser struct {
	Id int32 `json:"Id"`

	LoginName string `json:"LoginName"`

	UserName string `json:"UserName"`

	Email string `json:"Email"`
}

const GithubUserAvroCRC64Fingerprint = "Ky\xd2x\xf7S&t"

func NewGithubUser() GithubUser {
	r := GithubUser{}
	return r
}

func DeserializeGithubUser(r io.Reader) (GithubUser, error) {
	t := NewGithubUser()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeGithubUserFromSchema(r io.Reader, schema string) (GithubUser, error) {
	t := NewGithubUser()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeGithubUser(r GithubUser, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.LoginName, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.UserName, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Email, w)
	if err != nil {
		return err
	}
	return err
}

func (r GithubUser) Serialize(w io.Writer) error {
	return writeGithubUser(r, w)
}

func (r GithubUser) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"LoginName\",\"type\":\"string\"},{\"name\":\"UserName\",\"type\":\"string\"},{\"name\":\"Email\",\"type\":\"string\"}],\"name\":\"Andreani.GithubIntegration.Events.Common.GithubUser\",\"type\":\"record\"}"
}

func (r GithubUser) SchemaName() string {
	return "Andreani.GithubIntegration.Events.Common.GithubUser"
}

func (_ GithubUser) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ GithubUser) SetInt(v int32)       { panic("Unsupported operation") }
func (_ GithubUser) SetLong(v int64)      { panic("Unsupported operation") }
func (_ GithubUser) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ GithubUser) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ GithubUser) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ GithubUser) SetString(v string)   { panic("Unsupported operation") }
func (_ GithubUser) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *GithubUser) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.LoginName}

		return w

	case 2:
		w := types.String{Target: &r.UserName}

		return w

	case 3:
		w := types.String{Target: &r.Email}

		return w

	}
	panic("Unknown field index")
}

func (r *GithubUser) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *GithubUser) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ GithubUser) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ GithubUser) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ GithubUser) HintSize(int)                     { panic("Unsupported operation") }
func (_ GithubUser) Finalize()                        {}

func (_ GithubUser) AvroCRC64Fingerprint() []byte {
	return []byte(GithubUserAvroCRC64Fingerprint)
}

func (r GithubUser) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["LoginName"], err = json.Marshal(r.LoginName)
	if err != nil {
		return nil, err
	}
	output["UserName"], err = json.Marshal(r.UserName)
	if err != nil {
		return nil, err
	}
	output["Email"], err = json.Marshal(r.Email)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *GithubUser) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["LoginName"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoginName); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LoginName")
	}
	val = func() json.RawMessage {
		if v, ok := fields["UserName"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UserName); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for UserName")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Email"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Email); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Email")
	}
	return nil
}
