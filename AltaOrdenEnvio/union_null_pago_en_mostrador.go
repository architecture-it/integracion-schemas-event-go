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

type UnionNullPagoEnMostradorTypeEnum int

const (
	UnionNullPagoEnMostradorTypeEnumPagoEnMostrador UnionNullPagoEnMostradorTypeEnum = 1
)

type UnionNullPagoEnMostrador struct {
	Null            *types.NullVal
	PagoEnMostrador PagoEnMostrador
	UnionType       UnionNullPagoEnMostradorTypeEnum
}

func writeUnionNullPagoEnMostrador(r *UnionNullPagoEnMostrador, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullPagoEnMostradorTypeEnumPagoEnMostrador:
		return writePagoEnMostrador(r.PagoEnMostrador, w)
	}
	return fmt.Errorf("invalid value for *UnionNullPagoEnMostrador")
}

func NewUnionNullPagoEnMostrador() *UnionNullPagoEnMostrador {
	return &UnionNullPagoEnMostrador{}
}

func (r *UnionNullPagoEnMostrador) Serialize(w io.Writer) error {
	return writeUnionNullPagoEnMostrador(r, w)
}

func DeserializeUnionNullPagoEnMostrador(r io.Reader) (*UnionNullPagoEnMostrador, error) {
	t := NewUnionNullPagoEnMostrador()
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

func DeserializeUnionNullPagoEnMostradorFromSchema(r io.Reader, schema string) (*UnionNullPagoEnMostrador, error) {
	t := NewUnionNullPagoEnMostrador()
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

func (r *UnionNullPagoEnMostrador) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"MontoACobrar\",\"type\":\"float\"},{\"default\":null,\"name\":\"documentoTipo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"documentoNumero\",\"type\":[\"null\",\"string\"]}],\"name\":\"PagoEnMostrador\",\"type\":\"record\"}]"
}

func (_ *UnionNullPagoEnMostrador) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullPagoEnMostrador) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullPagoEnMostrador) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullPagoEnMostrador) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullPagoEnMostrador) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullPagoEnMostrador) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullPagoEnMostrador) SetLong(v int64) {

	r.UnionType = (UnionNullPagoEnMostradorTypeEnum)(v)
}

func (r *UnionNullPagoEnMostrador) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.PagoEnMostrador = NewPagoEnMostrador()
		return &types.Record{Target: (&r.PagoEnMostrador)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullPagoEnMostrador) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullPagoEnMostrador) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullPagoEnMostrador) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullPagoEnMostrador) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullPagoEnMostrador) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullPagoEnMostrador) Finalize()                        {}

func (r *UnionNullPagoEnMostrador) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullPagoEnMostradorTypeEnumPagoEnMostrador:
		return json.Marshal(map[string]interface{}{"Andreani.AltaOrdenEnvio.Events.Common.PagoEnMostrador": r.PagoEnMostrador})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullPagoEnMostrador")
}

func (r *UnionNullPagoEnMostrador) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.AltaOrdenEnvio.Events.Common.PagoEnMostrador"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.PagoEnMostrador)
	}
	return fmt.Errorf("invalid value for *UnionNullPagoEnMostrador")
}
