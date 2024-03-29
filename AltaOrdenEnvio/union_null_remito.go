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

type UnionNullRemitoTypeEnum int

const (
	UnionNullRemitoTypeEnumRemito UnionNullRemitoTypeEnum = 1
)

type UnionNullRemito struct {
	Null      *types.NullVal
	Remito    Remito
	UnionType UnionNullRemitoTypeEnum
}

func writeUnionNullRemito(r *UnionNullRemito, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullRemitoTypeEnumRemito:
		return writeRemito(r.Remito, w)
	}
	return fmt.Errorf("invalid value for *UnionNullRemito")
}

func NewUnionNullRemito() *UnionNullRemito {
	return &UnionNullRemito{}
}

func (r *UnionNullRemito) Serialize(w io.Writer) error {
	return writeUnionNullRemito(r, w)
}

func DeserializeUnionNullRemito(r io.Reader) (*UnionNullRemito, error) {
	t := NewUnionNullRemito()
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

func DeserializeUnionNullRemitoFromSchema(r io.Reader, schema string) (*UnionNullRemito, error) {
	t := NewUnionNullRemito()
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

func (r *UnionNullRemito) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"complementarios\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"name\":\"numeroDeRemito\",\"type\":[\"null\",\"string\"]}],\"name\":\"Remito\",\"type\":\"record\"}]"
}

func (_ *UnionNullRemito) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullRemito) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullRemito) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullRemito) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullRemito) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullRemito) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullRemito) SetLong(v int64) {

	r.UnionType = (UnionNullRemitoTypeEnum)(v)
}

func (r *UnionNullRemito) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Remito = NewRemito()
		return &types.Record{Target: (&r.Remito)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullRemito) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullRemito) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullRemito) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullRemito) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullRemito) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullRemito) Finalize()                        {}

func (r *UnionNullRemito) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullRemitoTypeEnumRemito:
		return json.Marshal(map[string]interface{}{"Andreani.AltaOrdenEnvio.Events.Common.Remito": r.Remito})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullRemito")
}

func (r *UnionNullRemito) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.AltaOrdenEnvio.Events.Common.Remito"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Remito)
	}
	return fmt.Errorf("invalid value for *UnionNullRemito")
}
