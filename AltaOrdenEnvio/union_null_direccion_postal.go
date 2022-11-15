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

type UnionNullDireccionPostalTypeEnum int

const (
	UnionNullDireccionPostalTypeEnumDireccionPostal UnionNullDireccionPostalTypeEnum = 1
)

type UnionNullDireccionPostal struct {
	Null            *types.NullVal
	DireccionPostal DireccionPostal
	UnionType       UnionNullDireccionPostalTypeEnum
}

func writeUnionNullDireccionPostal(r *UnionNullDireccionPostal, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullDireccionPostalTypeEnumDireccionPostal:
		return writeDireccionPostal(r.DireccionPostal, w)
	}
	return fmt.Errorf("invalid value for *UnionNullDireccionPostal")
}

func NewUnionNullDireccionPostal() *UnionNullDireccionPostal {
	return &UnionNullDireccionPostal{}
}

func (r *UnionNullDireccionPostal) Serialize(w io.Writer) error {
	return writeUnionNullDireccionPostal(r, w)
}

func DeserializeUnionNullDireccionPostal(r io.Reader) (*UnionNullDireccionPostal, error) {
	t := NewUnionNullDireccionPostal()
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

func DeserializeUnionNullDireccionPostalFromSchema(r io.Reader, schema string) (*UnionNullDireccionPostal, error) {
	t := NewUnionNullDireccionPostal()
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

func (r *UnionNullDireccionPostal) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"casillaDeCorreo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"region\",\"type\":[\"null\",\"string\"]}],\"name\":\"DireccionPostal\",\"type\":\"record\"}]"
}

func (_ *UnionNullDireccionPostal) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullDireccionPostal) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullDireccionPostal) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullDireccionPostal) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullDireccionPostal) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullDireccionPostal) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullDireccionPostal) SetLong(v int64) {

	r.UnionType = (UnionNullDireccionPostalTypeEnum)(v)
}

func (r *UnionNullDireccionPostal) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.DireccionPostal = NewDireccionPostal()
		return &types.Record{Target: (&r.DireccionPostal)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullDireccionPostal) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullDireccionPostal) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullDireccionPostal) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullDireccionPostal) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullDireccionPostal) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullDireccionPostal) Finalize()                        {}

func (r *UnionNullDireccionPostal) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullDireccionPostalTypeEnumDireccionPostal:
		return json.Marshal(map[string]interface{}{"Andreani.AltaOrdenEnvio.Events.Common.DireccionPostal": r.DireccionPostal})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullDireccionPostal")
}

func (r *UnionNullDireccionPostal) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.AltaOrdenEnvio.Events.Common.DireccionPostal"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.DireccionPostal)
	}
	return fmt.Errorf("invalid value for *UnionNullDireccionPostal")
}
