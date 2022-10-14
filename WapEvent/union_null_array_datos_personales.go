// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ListaDeDatosPersonales.avsc
 */
package WapEventEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullArrayDatosPersonalesTypeEnum int

const (
	UnionNullArrayDatosPersonalesTypeEnumArrayDatosPersonales UnionNullArrayDatosPersonalesTypeEnum = 1
)

type UnionNullArrayDatosPersonales struct {
	Null                 *types.NullVal
	ArrayDatosPersonales []DatosPersonales
	UnionType            UnionNullArrayDatosPersonalesTypeEnum
}

func writeUnionNullArrayDatosPersonales(r *UnionNullArrayDatosPersonales, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullArrayDatosPersonalesTypeEnumArrayDatosPersonales:
		return writeArrayDatosPersonales(r.ArrayDatosPersonales, w)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayDatosPersonales")
}

func NewUnionNullArrayDatosPersonales() *UnionNullArrayDatosPersonales {
	return &UnionNullArrayDatosPersonales{}
}

func (r *UnionNullArrayDatosPersonales) Serialize(w io.Writer) error {
	return writeUnionNullArrayDatosPersonales(r, w)
}

func DeserializeUnionNullArrayDatosPersonales(r io.Reader) (*UnionNullArrayDatosPersonales, error) {
	t := NewUnionNullArrayDatosPersonales()
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

func DeserializeUnionNullArrayDatosPersonalesFromSchema(r io.Reader, schema string) (*UnionNullArrayDatosPersonales, error) {
	t := NewUnionNullArrayDatosPersonales()
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

func (r *UnionNullArrayDatosPersonales) Schema() string {
	return "[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idinternocliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"eMail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"tipo\",\"type\":[\"null\",\"string\"]},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDeTelefonos\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"contacto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoDeDocumento\",\"type\":[\"null\",{\"name\":\"TipoDeDocumento\",\"symbols\":[\"undefined\",\"DNI\",\"CUIT\",\"CUIL\"],\"type\":\"enum\"}]}],\"name\":\"DatosPersonales\",\"type\":\"record\"},\"type\":\"array\"}]"
}

func (_ *UnionNullArrayDatosPersonales) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullArrayDatosPersonales) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullArrayDatosPersonales) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullArrayDatosPersonales) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullArrayDatosPersonales) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullArrayDatosPersonales) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullArrayDatosPersonales) SetLong(v int64) {

	r.UnionType = (UnionNullArrayDatosPersonalesTypeEnum)(v)
}

func (r *UnionNullArrayDatosPersonales) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ArrayDatosPersonales = make([]DatosPersonales, 0)
		return &ArrayDatosPersonalesWrapper{Target: (&r.ArrayDatosPersonales)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayDatosPersonales) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullArrayDatosPersonales) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullArrayDatosPersonales) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullArrayDatosPersonales) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayDatosPersonales) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullArrayDatosPersonales) Finalize()                {}

func (r *UnionNullArrayDatosPersonales) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullArrayDatosPersonalesTypeEnumArrayDatosPersonales:
		return json.Marshal(map[string]interface{}{"array": r.ArrayDatosPersonales})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullArrayDatosPersonales")
}

func (r *UnionNullArrayDatosPersonales) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ArrayDatosPersonales)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayDatosPersonales")
}
