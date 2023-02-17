// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ListaDeTiposDeAcciones.avsc
 */
package Wapv2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ListaDeTiposDeAcciones struct {
	TiposDeAcciones []TiposDeAcciones `json:"tiposDeAcciones"`
}

const ListaDeTiposDeAccionesAvroCRC64Fingerprint = "\xbb\x9f\x19\x1b\xd2\xc2\x12\xdd"

func NewListaDeTiposDeAcciones() ListaDeTiposDeAcciones {
	r := ListaDeTiposDeAcciones{}
	r.TiposDeAcciones = make([]TiposDeAcciones, 0)

	return r
}

func DeserializeListaDeTiposDeAcciones(r io.Reader) (ListaDeTiposDeAcciones, error) {
	t := NewListaDeTiposDeAcciones()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeListaDeTiposDeAccionesFromSchema(r io.Reader, schema string) (ListaDeTiposDeAcciones, error) {
	t := NewListaDeTiposDeAcciones()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeListaDeTiposDeAcciones(r ListaDeTiposDeAcciones, w io.Writer) error {
	var err error
	err = writeArrayTiposDeAcciones(r.TiposDeAcciones, w)
	if err != nil {
		return err
	}
	return err
}

func (r ListaDeTiposDeAcciones) Serialize(w io.Writer) error {
	return writeListaDeTiposDeAcciones(r, w)
}

func (r ListaDeTiposDeAcciones) Schema() string {
	return "{\"fields\":[{\"name\":\"tiposDeAcciones\",\"type\":{\"items\":{\"fields\":[{\"name\":\"tipoDeAccion\",\"type\":\"int\"}],\"name\":\"TiposDeAcciones\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.Wapv2.Events.Record.ListaDeTiposDeAcciones\",\"type\":\"record\"}"
}

func (r ListaDeTiposDeAcciones) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.ListaDeTiposDeAcciones"
}

func (_ ListaDeTiposDeAcciones) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ListaDeTiposDeAcciones) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ListaDeTiposDeAcciones) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ListaDeTiposDeAcciones) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ListaDeTiposDeAcciones) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ListaDeTiposDeAcciones) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ListaDeTiposDeAcciones) SetString(v string)   { panic("Unsupported operation") }
func (_ ListaDeTiposDeAcciones) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ListaDeTiposDeAcciones) Get(i int) types.Field {
	switch i {
	case 0:
		r.TiposDeAcciones = make([]TiposDeAcciones, 0)

		w := ArrayTiposDeAccionesWrapper{Target: &r.TiposDeAcciones}

		return w

	}
	panic("Unknown field index")
}

func (r *ListaDeTiposDeAcciones) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ListaDeTiposDeAcciones) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ListaDeTiposDeAcciones) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ListaDeTiposDeAcciones) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ListaDeTiposDeAcciones) HintSize(int)                     { panic("Unsupported operation") }
func (_ ListaDeTiposDeAcciones) Finalize()                        {}

func (_ ListaDeTiposDeAcciones) AvroCRC64Fingerprint() []byte {
	return []byte(ListaDeTiposDeAccionesAvroCRC64Fingerprint)
}

func (r ListaDeTiposDeAcciones) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["tiposDeAcciones"], err = json.Marshal(r.TiposDeAcciones)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ListaDeTiposDeAcciones) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["tiposDeAcciones"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TiposDeAcciones); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tiposDeAcciones")
	}
	return nil
}
