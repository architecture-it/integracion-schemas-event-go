// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Person.avsc
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

type UnionNullTeamTypeEnum int

const (
	UnionNullTeamTypeEnumTeam UnionNullTeamTypeEnum = 1
)

type UnionNullTeam struct {
	Null      *types.NullVal
	Team      Team
	UnionType UnionNullTeamTypeEnum
}

func writeUnionNullTeam(r *UnionNullTeam, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullTeamTypeEnumTeam:
		return writeTeam(r.Team, w)
	}
	return fmt.Errorf("invalid value for *UnionNullTeam")
}

func NewUnionNullTeam() *UnionNullTeam {
	return &UnionNullTeam{}
}

func (r *UnionNullTeam) Serialize(w io.Writer) error {
	return writeUnionNullTeam(r, w)
}

func DeserializeUnionNullTeam(r io.Reader) (*UnionNullTeam, error) {
	t := NewUnionNullTeam()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionNullTeamFromSchema(r io.Reader, schema string) (*UnionNullTeam, error) {
	t := NewUnionNullTeam()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r *UnionNullTeam) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"tl\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"boss\",\"type\":[\"null\",\"string\"]},{\"name\":\"members\",\"type\":{\"items\":\"string\",\"type\":\"array\"}}],\"name\":\"Team\",\"type\":\"record\"}]"
}

func (_ *UnionNullTeam) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullTeam) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullTeam) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullTeam) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullTeam) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullTeam) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullTeam) SetLong(v int64) {

	r.UnionType = (UnionNullTeamTypeEnum)(v)
}

func (r *UnionNullTeam) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Team = NewTeam()
		return &types.Record{Target: (&r.Team)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullTeam) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullTeam) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullTeam) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullTeam) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullTeam) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullTeam) Finalize()                        {}

func (r *UnionNullTeam) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullTeamTypeEnumTeam:
		return json.Marshal(map[string]interface{}{"Andreani.Test.Events.Record.Common.Team": r.Team})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullTeam")
}

func (r *UnionNullTeam) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.Test.Events.Record.Common.Team"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Team)
	}
	return fmt.Errorf("invalid value for *UnionNullTeam")
}
