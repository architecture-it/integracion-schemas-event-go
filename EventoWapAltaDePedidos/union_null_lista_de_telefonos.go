// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     AltaDePedidosSolicitada.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullListaDeTelefonosTypeEnum int

const (
	UnionNullListaDeTelefonosTypeEnumListaDeTelefonos UnionNullListaDeTelefonosTypeEnum = 1
)

type UnionNullListaDeTelefonos struct {
	Null             *types.NullVal
	ListaDeTelefonos ListaDeTelefonos
	UnionType        UnionNullListaDeTelefonosTypeEnum
}

func writeUnionNullListaDeTelefonos(r *UnionNullListaDeTelefonos, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullListaDeTelefonosTypeEnumListaDeTelefonos:
		return writeListaDeTelefonos(r.ListaDeTelefonos, w)
	}
	return fmt.Errorf("invalid value for *UnionNullListaDeTelefonos")
}

func NewUnionNullListaDeTelefonos() *UnionNullListaDeTelefonos {
	return &UnionNullListaDeTelefonos{}
}

func (r *UnionNullListaDeTelefonos) Serialize(w io.Writer) error {
	return writeUnionNullListaDeTelefonos(r, w)
}

func DeserializeUnionNullListaDeTelefonos(r io.Reader) (*UnionNullListaDeTelefonos, error) {
	t := NewUnionNullListaDeTelefonos()
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

func DeserializeUnionNullListaDeTelefonosFromSchema(r io.Reader, schema string) (*UnionNullListaDeTelefonos, error) {
	t := NewUnionNullListaDeTelefonos()
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

func (r *UnionNullListaDeTelefonos) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"tipo\",\"type\":[\"null\",\"string\"]},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"namespace\":\"Andreani.Wap.Events.Record\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDeTelefonos\",\"namespace\":\"Andreani.Wap.Events.Record\",\"type\":\"record\"}]"
}

func (_ *UnionNullListaDeTelefonos) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullListaDeTelefonos) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullListaDeTelefonos) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullListaDeTelefonos) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullListaDeTelefonos) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullListaDeTelefonos) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullListaDeTelefonos) SetLong(v int64) {

	r.UnionType = (UnionNullListaDeTelefonosTypeEnum)(v)
}

func (r *UnionNullListaDeTelefonos) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ListaDeTelefonos = NewListaDeTelefonos()
		return &types.Record{Target: (&r.ListaDeTelefonos)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullListaDeTelefonos) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullListaDeTelefonos) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullListaDeTelefonos) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullListaDeTelefonos) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullListaDeTelefonos) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullListaDeTelefonos) Finalize()                        {}

func (r *UnionNullListaDeTelefonos) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullListaDeTelefonosTypeEnumListaDeTelefonos:
		return json.Marshal(map[string]interface{}{"Andreani.Wap.Events.Record.ListaDeTelefonos": r.ListaDeTelefonos})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullListaDeTelefonos")
}

func (r *UnionNullListaDeTelefonos) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.Wap.Events.Record.ListaDeTelefonos"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ListaDeTelefonos)
	}
	return fmt.Errorf("invalid value for *UnionNullListaDeTelefonos")
}
