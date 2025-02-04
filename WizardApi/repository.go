// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Repository.avsc
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

type Repository struct {
	RepositoryId int64 `json:"RepositoryId"`

	Name string `json:"Name"`

	User GithubUser `json:"User"`

	Organization Organization `json:"Organization"`
}

const RepositoryAvroCRC64Fingerprint = "\xfc-\x9e\xd7.\xa3\xb7*"

func NewRepository() Repository {
	r := Repository{}
	r.User = NewGithubUser()

	r.Organization = NewOrganization()

	return r
}

func DeserializeRepository(r io.Reader) (Repository, error) {
	t := NewRepository()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRepositoryFromSchema(r io.Reader, schema string) (Repository, error) {
	t := NewRepository()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRepository(r Repository, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.RepositoryId, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	err = writeGithubUser(r.User, w)
	if err != nil {
		return err
	}
	err = writeOrganization(r.Organization, w)
	if err != nil {
		return err
	}
	return err
}

func (r Repository) Serialize(w io.Writer) error {
	return writeRepository(r, w)
}

func (r Repository) Schema() string {
	return "{\"fields\":[{\"name\":\"RepositoryId\",\"type\":\"long\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"User\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"LoginName\",\"type\":\"string\"},{\"name\":\"UserName\",\"type\":\"string\"},{\"default\":null,\"name\":\"Email\",\"type\":[\"null\",\"string\"]}],\"name\":\"GithubUser\",\"namespace\":\"Andreani.WizardApi.Events.Common\",\"type\":\"record\"}},{\"name\":\"Organization\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"long\"},{\"name\":\"Name\",\"type\":\"string\"}],\"name\":\"Organization\",\"namespace\":\"Andreani.WizardApi.Events.Common\",\"type\":\"record\"}}],\"name\":\"Andreani.WizardApi.Events.Record.Repository\",\"type\":\"record\"}"
}

func (r Repository) SchemaName() string {
	return "Andreani.WizardApi.Events.Record.Repository"
}

func (_ Repository) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Repository) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Repository) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Repository) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Repository) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Repository) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Repository) SetString(v string)   { panic("Unsupported operation") }
func (_ Repository) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Repository) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.RepositoryId}

		return w

	case 1:
		w := types.String{Target: &r.Name}

		return w

	case 2:
		r.User = NewGithubUser()

		w := types.Record{Target: &r.User}

		return w

	case 3:
		r.Organization = NewOrganization()

		w := types.Record{Target: &r.Organization}

		return w

	}
	panic("Unknown field index")
}

func (r *Repository) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Repository) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Repository) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Repository) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Repository) HintSize(int)                     { panic("Unsupported operation") }
func (_ Repository) Finalize()                        {}

func (_ Repository) AvroCRC64Fingerprint() []byte {
	return []byte(RepositoryAvroCRC64Fingerprint)
}

func (r Repository) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["RepositoryId"], err = json.Marshal(r.RepositoryId)
	if err != nil {
		return nil, err
	}
	output["Name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	output["User"], err = json.Marshal(r.User)
	if err != nil {
		return nil, err
	}
	output["Organization"], err = json.Marshal(r.Organization)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Repository) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["RepositoryId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RepositoryId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for RepositoryId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["User"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.User); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for User")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Organization"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Organization); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Organization")
	}
	return nil
}
