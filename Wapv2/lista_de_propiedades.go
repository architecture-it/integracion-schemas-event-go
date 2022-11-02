// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoDeAlmacenSCE.avsc
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

type ListaDePropiedades struct {
	Metadatos []Metadato `json:"metadatos"`
}

const ListaDePropiedadesAvroCRC64Fingerprint = "\xdfP\xefy(mg\xe8"

func NewListaDePropiedades() ListaDePropiedades {
	r := ListaDePropiedades{}
	r.Metadatos = make([]Metadato, 0)

	return r
}

func DeserializeListaDePropiedades(r io.Reader) (ListaDePropiedades, error) {
	t := NewListaDePropiedades()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeListaDePropiedadesFromSchema(r io.Reader, schema string) (ListaDePropiedades, error) {
	t := NewListaDePropiedades()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeListaDePropiedades(r ListaDePropiedades, w io.Writer) error {
	var err error
	err = writeArrayMetadato(r.Metadatos, w)
	if err != nil {
		return err
	}
	return err
}

func (r ListaDePropiedades) Serialize(w io.Writer) error {
	return writeListaDePropiedades(r, w)
}

func (r ListaDePropiedades) Schema() string {
	return "{\"fields\":[{\"name\":\"metadatos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.Wapv2.Events.Record.ListaDePropiedades\",\"type\":\"record\"}"
}

func (r ListaDePropiedades) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.ListaDePropiedades"
}

func (_ ListaDePropiedades) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ListaDePropiedades) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ListaDePropiedades) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ListaDePropiedades) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ListaDePropiedades) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ListaDePropiedades) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ListaDePropiedades) SetString(v string)   { panic("Unsupported operation") }
func (_ ListaDePropiedades) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ListaDePropiedades) Get(i int) types.Field {
	switch i {
	case 0:
		r.Metadatos = make([]Metadato, 0)

		w := ArrayMetadatoWrapper{Target: &r.Metadatos}

		return w

	}
	panic("Unknown field index")
}

func (r *ListaDePropiedades) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ListaDePropiedades) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ListaDePropiedades) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ListaDePropiedades) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ListaDePropiedades) HintSize(int)                     { panic("Unsupported operation") }
func (_ ListaDePropiedades) Finalize()                        {}

func (_ ListaDePropiedades) AvroCRC64Fingerprint() []byte {
	return []byte(ListaDePropiedadesAvroCRC64Fingerprint)
}

func (r ListaDePropiedades) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["metadatos"], err = json.Marshal(r.Metadatos)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ListaDePropiedades) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["metadatos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Metadatos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for metadatos")
	}
	return nil
}
