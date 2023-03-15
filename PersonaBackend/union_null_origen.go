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

type UnionNullOrigenTypeEnum int

const (
	UnionNullOrigenTypeEnumOrigen UnionNullOrigenTypeEnum = 1
)

type UnionNullOrigen struct {
	Null      *types.NullVal
	Origen    Origen
	UnionType UnionNullOrigenTypeEnum
}

func writeUnionNullOrigen(r *UnionNullOrigen, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullOrigenTypeEnumOrigen:
		return writeOrigen(r.Origen, w)
	}
	return fmt.Errorf("invalid value for *UnionNullOrigen")
}

func NewUnionNullOrigen() *UnionNullOrigen {
	return &UnionNullOrigen{}
}

func (r *UnionNullOrigen) Serialize(w io.Writer) error {
	return writeUnionNullOrigen(r, w)
}

func DeserializeUnionNullOrigen(r io.Reader) (*UnionNullOrigen, error) {
	t := NewUnionNullOrigen()
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

func DeserializeUnionNullOrigenFromSchema(r io.Reader, schema string) (*UnionNullOrigen, error) {
	t := NewUnionNullOrigen()
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

func (r *UnionNullOrigen) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"Tipo\",\"type\":\"string\"},{\"name\":\"Calle\",\"type\":\"string\"},{\"name\":\"Numero\",\"type\":\"string\"},{\"name\":\"Piso\",\"type\":\"string\"},{\"name\":\"Unidad\",\"type\":\"string\"},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"CodigoPostal\",\"type\":\"string\"},{\"name\":\"Provincia\",\"type\":\"string\"},{\"default\":null,\"name\":\"SucursalId\",\"type\":[\"null\",\"string\"]}],\"name\":\"Origen\",\"type\":\"record\"}]"
}

func (_ *UnionNullOrigen) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullOrigen) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullOrigen) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullOrigen) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullOrigen) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullOrigen) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullOrigen) SetLong(v int64) {

	r.UnionType = (UnionNullOrigenTypeEnum)(v)
}

func (r *UnionNullOrigen) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Origen = NewOrigen()
		return &types.Record{Target: (&r.Origen)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullOrigen) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullOrigen) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullOrigen) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullOrigen) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullOrigen) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullOrigen) Finalize()                        {}

func (r *UnionNullOrigen) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullOrigenTypeEnumOrigen:
		return json.Marshal(map[string]interface{}{"Andreani.PersonaBackend.Events.Common.Origen": r.Origen})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullOrigen")
}

func (r *UnionNullOrigen) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.PersonaBackend.Events.Common.Origen"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Origen)
	}
	return fmt.Errorf("invalid value for *UnionNullOrigen")
}
