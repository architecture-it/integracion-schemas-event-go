// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     UOHdrCreada.avsc
 */
package UOPublisherHdrEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullArrayRetiroDeValorTypeEnum int

const (
	UnionNullArrayRetiroDeValorTypeEnumArrayRetiroDeValor UnionNullArrayRetiroDeValorTypeEnum = 1
)

type UnionNullArrayRetiroDeValor struct {
	Null               *types.NullVal
	ArrayRetiroDeValor []RetiroDeValor
	UnionType          UnionNullArrayRetiroDeValorTypeEnum
}

func writeUnionNullArrayRetiroDeValor(r *UnionNullArrayRetiroDeValor, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullArrayRetiroDeValorTypeEnumArrayRetiroDeValor:
		return writeArrayRetiroDeValor(r.ArrayRetiroDeValor, w)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayRetiroDeValor")
}

func NewUnionNullArrayRetiroDeValor() *UnionNullArrayRetiroDeValor {
	return &UnionNullArrayRetiroDeValor{}
}

func (r *UnionNullArrayRetiroDeValor) Serialize(w io.Writer) error {
	return writeUnionNullArrayRetiroDeValor(r, w)
}

func DeserializeUnionNullArrayRetiroDeValor(r io.Reader) (*UnionNullArrayRetiroDeValor, error) {
	t := NewUnionNullArrayRetiroDeValor()
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

func DeserializeUnionNullArrayRetiroDeValorFromSchema(r io.Reader, schema string) (*UnionNullArrayRetiroDeValor, error) {
	t := NewUnionNullArrayRetiroDeValor()
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

func (r *UnionNullArrayRetiroDeValor) Schema() string {
	return "[\"null\",{\"items\":{\"fields\":[{\"name\":\"Monto\",\"type\":\"double\"},{\"name\":\"PagoExacto\",\"type\":\"boolean\"}],\"name\":\"RetiroDeValor\",\"type\":\"record\"},\"type\":\"array\"}]"
}

func (_ *UnionNullArrayRetiroDeValor) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullArrayRetiroDeValor) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullArrayRetiroDeValor) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullArrayRetiroDeValor) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullArrayRetiroDeValor) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullArrayRetiroDeValor) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullArrayRetiroDeValor) SetLong(v int64) {

	r.UnionType = (UnionNullArrayRetiroDeValorTypeEnum)(v)
}

func (r *UnionNullArrayRetiroDeValor) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ArrayRetiroDeValor = make([]RetiroDeValor, 0)
		return &ArrayRetiroDeValorWrapper{Target: (&r.ArrayRetiroDeValor)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayRetiroDeValor) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullArrayRetiroDeValor) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullArrayRetiroDeValor) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullArrayRetiroDeValor) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayRetiroDeValor) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullArrayRetiroDeValor) Finalize()                {}

func (r *UnionNullArrayRetiroDeValor) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullArrayRetiroDeValorTypeEnumArrayRetiroDeValor:
		return json.Marshal(map[string]interface{}{"array": r.ArrayRetiroDeValor})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullArrayRetiroDeValor")
}

func (r *UnionNullArrayRetiroDeValor) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ArrayRetiroDeValor)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayRetiroDeValor")
}
