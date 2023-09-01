// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Pedido.avsc
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

type ListaDeDocumentos struct {
	Documentos []Documentos `json:"documentos"`
}

const ListaDeDocumentosAvroCRC64Fingerprint = "\xff\xb4/\xb3<\xbc\x02\n"

func NewListaDeDocumentos() ListaDeDocumentos {
	r := ListaDeDocumentos{}
	r.Documentos = make([]Documentos, 0)

	return r
}

func DeserializeListaDeDocumentos(r io.Reader) (ListaDeDocumentos, error) {
	t := NewListaDeDocumentos()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeListaDeDocumentosFromSchema(r io.Reader, schema string) (ListaDeDocumentos, error) {
	t := NewListaDeDocumentos()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeListaDeDocumentos(r ListaDeDocumentos, w io.Writer) error {
	var err error
	err = writeArrayDocumentos(r.Documentos, w)
	if err != nil {
		return err
	}
	return err
}

func (r ListaDeDocumentos) Serialize(w io.Writer) error {
	return writeListaDeDocumentos(r, w)
}

func (r ListaDeDocumentos) Schema() string {
	return "{\"fields\":[{\"name\":\"documentos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"documentoBase64\",\"type\":\"string\"}],\"name\":\"Documentos\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.Wapv2.Events.Record.ListaDeDocumentos\",\"type\":\"record\"}"
}

func (r ListaDeDocumentos) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.ListaDeDocumentos"
}

func (_ ListaDeDocumentos) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ListaDeDocumentos) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ListaDeDocumentos) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ListaDeDocumentos) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ListaDeDocumentos) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ListaDeDocumentos) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ListaDeDocumentos) SetString(v string)   { panic("Unsupported operation") }
func (_ ListaDeDocumentos) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ListaDeDocumentos) Get(i int) types.Field {
	switch i {
	case 0:
		r.Documentos = make([]Documentos, 0)

		w := ArrayDocumentosWrapper{Target: &r.Documentos}

		return w

	}
	panic("Unknown field index")
}

func (r *ListaDeDocumentos) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ListaDeDocumentos) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ListaDeDocumentos) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ListaDeDocumentos) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ListaDeDocumentos) HintSize(int)                     { panic("Unsupported operation") }
func (_ ListaDeDocumentos) Finalize()                        {}

func (_ ListaDeDocumentos) AvroCRC64Fingerprint() []byte {
	return []byte(ListaDeDocumentosAvroCRC64Fingerprint)
}

func (r ListaDeDocumentos) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["documentos"], err = json.Marshal(r.Documentos)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ListaDeDocumentos) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["documentos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Documentos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for documentos")
	}
	return nil
}
