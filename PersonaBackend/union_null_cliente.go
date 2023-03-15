// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnviosEvent.avsc
 */
package PersonaBackendEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullClienteTypeEnum int

const (
	UnionNullClienteTypeEnumCliente UnionNullClienteTypeEnum = 1
)

type UnionNullCliente struct {
	Null      *types.NullVal
	Cliente   Cliente
	UnionType UnionNullClienteTypeEnum
}

func writeUnionNullCliente(r *UnionNullCliente, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullClienteTypeEnumCliente:
		return writeCliente(r.Cliente, w)
	}
	return fmt.Errorf("invalid value for *UnionNullCliente")
}

func NewUnionNullCliente() *UnionNullCliente {
	return &UnionNullCliente{}
}

func (r *UnionNullCliente) Serialize(w io.Writer) error {
	return writeUnionNullCliente(r, w)
}

func DeserializeUnionNullCliente(r io.Reader) (*UnionNullCliente, error) {
	t := NewUnionNullCliente()
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

func DeserializeUnionNullClienteFromSchema(r io.Reader, schema string) (*UnionNullCliente, error) {
	t := NewUnionNullCliente()
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

func (r *UnionNullCliente) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"CodigoAndreani\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Logo\",\"type\":\"string\"}],\"name\":\"Cliente\",\"type\":\"record\"}]"
}

func (_ *UnionNullCliente) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullCliente) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullCliente) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullCliente) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullCliente) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullCliente) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullCliente) SetLong(v int64) {

	r.UnionType = (UnionNullClienteTypeEnum)(v)
}

func (r *UnionNullCliente) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Cliente = NewCliente()
		return &types.Record{Target: (&r.Cliente)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullCliente) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullCliente) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullCliente) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullCliente) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullCliente) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullCliente) Finalize()                        {}

func (r *UnionNullCliente) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullClienteTypeEnumCliente:
		return json.Marshal(map[string]interface{}{"Andreani.PersonaBackend.Events.Common.Cliente": r.Cliente})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullCliente")
}

func (r *UnionNullCliente) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.PersonaBackend.Events.Common.Cliente"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Cliente)
	}
	return fmt.Errorf("invalid value for *UnionNullCliente")
}
