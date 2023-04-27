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

type ListaDeFiles struct {
	Documentos []File `json:"documentos"`
}

const ListaDeFilesAvroCRC64Fingerprint = "\x04^\xb0\xc6]S\x94/"

func NewListaDeFiles() ListaDeFiles {
	r := ListaDeFiles{}
	r.Documentos = make([]File, 0)

	return r
}

func DeserializeListaDeFiles(r io.Reader) (ListaDeFiles, error) {
	t := NewListaDeFiles()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeListaDeFilesFromSchema(r io.Reader, schema string) (ListaDeFiles, error) {
	t := NewListaDeFiles()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeListaDeFiles(r ListaDeFiles, w io.Writer) error {
	var err error
	err = writeArrayFile(r.Documentos, w)
	if err != nil {
		return err
	}
	return err
}

func (r ListaDeFiles) Serialize(w io.Writer) error {
	return writeListaDeFiles(r, w)
}

func (r ListaDeFiles) Schema() string {
	return "{\"fields\":[{\"name\":\"documentos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"numeroEntidad\",\"type\":\"string\"},{\"name\":\"documento\",\"type\":\"string\"}],\"name\":\"File\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.Wapv2.Events.Record.ListaDeFiles\",\"type\":\"record\"}"
}

func (r ListaDeFiles) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.ListaDeFiles"
}

func (_ ListaDeFiles) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ListaDeFiles) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ListaDeFiles) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ListaDeFiles) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ListaDeFiles) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ListaDeFiles) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ListaDeFiles) SetString(v string)   { panic("Unsupported operation") }
func (_ ListaDeFiles) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ListaDeFiles) Get(i int) types.Field {
	switch i {
	case 0:
		r.Documentos = make([]File, 0)

		w := ArrayFileWrapper{Target: &r.Documentos}

		return w

	}
	panic("Unknown field index")
}

func (r *ListaDeFiles) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ListaDeFiles) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ListaDeFiles) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ListaDeFiles) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ListaDeFiles) HintSize(int)                     { panic("Unsupported operation") }
func (_ ListaDeFiles) Finalize()                        {}

func (_ ListaDeFiles) AvroCRC64Fingerprint() []byte {
	return []byte(ListaDeFilesAvroCRC64Fingerprint)
}

func (r ListaDeFiles) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["documentos"], err = json.Marshal(r.Documentos)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ListaDeFiles) UnmarshalJSON(data []byte) error {
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
