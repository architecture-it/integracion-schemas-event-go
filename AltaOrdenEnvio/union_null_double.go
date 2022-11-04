// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TransactionEvent.avsc
 */
package AltaOrdenEnvioEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullDoubleTypeEnum int

const (
	UnionNullDoubleTypeEnumDouble UnionNullDoubleTypeEnum = 1
)

type UnionNullDouble struct {
	Null      *types.NullVal
	Double    float64
	UnionType UnionNullDoubleTypeEnum
}

func writeUnionNullDouble(r *UnionNullDouble, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullDoubleTypeEnumDouble:
		return vm.WriteDouble(r.Double, w)
	}
	return fmt.Errorf("invalid value for *UnionNullDouble")
}

func NewUnionNullDouble() *UnionNullDouble {
	return &UnionNullDouble{}
}

func (r *UnionNullDouble) Serialize(w io.Writer) error {
	return writeUnionNullDouble(r, w)
}

func DeserializeUnionNullDouble(r io.Reader) (*UnionNullDouble, error) {
	t := NewUnionNullDouble()
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

func DeserializeUnionNullDoubleFromSchema(r io.Reader, schema string) (*UnionNullDouble, error) {
	t := NewUnionNullDouble()
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

func (r *UnionNullDouble) Schema() string {
	return "[\"null\",\"double\"]"
}

func (_ *UnionNullDouble) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullDouble) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullDouble) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullDouble) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullDouble) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullDouble) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullDouble) SetLong(v int64) {

	r.UnionType = (UnionNullDoubleTypeEnum)(v)
}

func (r *UnionNullDouble) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		return &types.Double{Target: (&r.Double)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullDouble) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullDouble) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullDouble) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullDouble) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullDouble) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullDouble) Finalize()                        {}

func (r *UnionNullDouble) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullDoubleTypeEnumDouble:
		return json.Marshal(map[string]interface{}{"double": r.Double})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullDouble")
}

func (r *UnionNullDouble) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["double"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Double)
	}
	return fmt.Errorf("invalid value for *UnionNullDouble")
}