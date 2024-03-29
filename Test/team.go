// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Team.avsc
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

type Team struct {
	Tl *UnionNullString `json:"tl"`

	Boss string `json:"boss"`
}

const TeamAvroCRC64Fingerprint = "^\xd1\xfc\v2G,L"

func NewTeam() Team {
	r := Team{}
	r.Tl = nil
	return r
}

func DeserializeTeam(r io.Reader) (Team, error) {
	t := NewTeam()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeTeamFromSchema(r io.Reader, schema string) (Team, error) {
	t := NewTeam()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeTeam(r Team, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Tl, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Boss, w)
	if err != nil {
		return err
	}
	return err
}

func (r Team) Serialize(w io.Writer) error {
	return writeTeam(r, w)
}

func (r Team) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"tl\",\"type\":[\"null\",\"string\"]},{\"name\":\"boss\",\"type\":\"string\"}],\"name\":\"Andreani.Test.Events.Record.Common.Team\",\"type\":\"record\"}"
}

func (r Team) SchemaName() string {
	return "Andreani.Test.Events.Record.Common.Team"
}

func (_ Team) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Team) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Team) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Team) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Team) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Team) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Team) SetString(v string)   { panic("Unsupported operation") }
func (_ Team) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Team) Get(i int) types.Field {
	switch i {
	case 0:
		r.Tl = NewUnionNullString()

		return r.Tl
	case 1:
		w := types.String{Target: &r.Boss}

		return w

	}
	panic("Unknown field index")
}

func (r *Team) SetDefault(i int) {
	switch i {
	case 0:
		r.Tl = nil
		return
	}
	panic("Unknown field index")
}

func (r *Team) NullField(i int) {
	switch i {
	case 0:
		r.Tl = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Team) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Team) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Team) HintSize(int)                     { panic("Unsupported operation") }
func (_ Team) Finalize()                        {}

func (_ Team) AvroCRC64Fingerprint() []byte {
	return []byte(TeamAvroCRC64Fingerprint)
}

func (r Team) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["tl"], err = json.Marshal(r.Tl)
	if err != nil {
		return nil, err
	}
	output["boss"], err = json.Marshal(r.Boss)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Team) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["tl"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Tl); err != nil {
			return err
		}
	} else {
		r.Tl = NewUnionNullString()

		r.Tl = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["boss"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Boss); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for boss")
	}
	return nil
}
