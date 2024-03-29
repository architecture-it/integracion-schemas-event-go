// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     OrdenDeEnvioCreada.avsc
 */
package TrazasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullTipoPagoDestinoTypeEnum int

const (
	UnionNullTipoPagoDestinoTypeEnumTipoPagoDestino UnionNullTipoPagoDestinoTypeEnum = 1
)

type UnionNullTipoPagoDestino struct {
	Null            *types.NullVal
	TipoPagoDestino TipoPagoDestino
	UnionType       UnionNullTipoPagoDestinoTypeEnum
}

func writeUnionNullTipoPagoDestino(r *UnionNullTipoPagoDestino, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullTipoPagoDestinoTypeEnumTipoPagoDestino:
		return writeTipoPagoDestino(r.TipoPagoDestino, w)
	}
	return fmt.Errorf("invalid value for *UnionNullTipoPagoDestino")
}

func NewUnionNullTipoPagoDestino() *UnionNullTipoPagoDestino {
	return &UnionNullTipoPagoDestino{}
}

func (r *UnionNullTipoPagoDestino) Serialize(w io.Writer) error {
	return writeUnionNullTipoPagoDestino(r, w)
}

func DeserializeUnionNullTipoPagoDestino(r io.Reader) (*UnionNullTipoPagoDestino, error) {
	t := NewUnionNullTipoPagoDestino()
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

func DeserializeUnionNullTipoPagoDestinoFromSchema(r io.Reader, schema string) (*UnionNullTipoPagoDestino, error) {
	t := NewUnionNullTipoPagoDestino()
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

func (r *UnionNullTipoPagoDestino) Schema() string {
	return "[\"null\",{\"name\":\"TipoPagoDestino\",\"symbols\":[\"undefined\",\"P\",\"D\"],\"type\":\"enum\"}]"
}

func (_ *UnionNullTipoPagoDestino) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullTipoPagoDestino) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullTipoPagoDestino) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullTipoPagoDestino) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullTipoPagoDestino) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullTipoPagoDestino) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullTipoPagoDestino) SetLong(v int64) {

	r.UnionType = (UnionNullTipoPagoDestinoTypeEnum)(v)
}

func (r *UnionNullTipoPagoDestino) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		return &TipoPagoDestinoWrapper{Target: (&r.TipoPagoDestino)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullTipoPagoDestino) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullTipoPagoDestino) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullTipoPagoDestino) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullTipoPagoDestino) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullTipoPagoDestino) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullTipoPagoDestino) Finalize()                        {}

func (r *UnionNullTipoPagoDestino) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullTipoPagoDestinoTypeEnumTipoPagoDestino:
		return json.Marshal(map[string]interface{}{"Integracion.Esquemas.Referencias.TipoPagoDestino": r.TipoPagoDestino})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullTipoPagoDestino")
}

func (r *UnionNullTipoPagoDestino) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Integracion.Esquemas.Referencias.TipoPagoDestino"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.TipoPagoDestino)
	}
	return fmt.Errorf("invalid value for *UnionNullTipoPagoDestino")
}
