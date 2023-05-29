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

type ListaDeCodigosDeTrazaANMAT struct {
	CodigosDeTrazaANMAT []CodigosDeTrazaANMAT `json:"codigosDeTrazaANMAT"`
}

const ListaDeCodigosDeTrazaANMATAvroCRC64Fingerprint = "2\x99\x82ھ\xf6\x8bY"

func NewListaDeCodigosDeTrazaANMAT() ListaDeCodigosDeTrazaANMAT {
	r := ListaDeCodigosDeTrazaANMAT{}
	r.CodigosDeTrazaANMAT = make([]CodigosDeTrazaANMAT, 0)

	return r
}

func DeserializeListaDeCodigosDeTrazaANMAT(r io.Reader) (ListaDeCodigosDeTrazaANMAT, error) {
	t := NewListaDeCodigosDeTrazaANMAT()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeListaDeCodigosDeTrazaANMATFromSchema(r io.Reader, schema string) (ListaDeCodigosDeTrazaANMAT, error) {
	t := NewListaDeCodigosDeTrazaANMAT()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeListaDeCodigosDeTrazaANMAT(r ListaDeCodigosDeTrazaANMAT, w io.Writer) error {
	var err error
	err = writeArrayCodigosDeTrazaANMAT(r.CodigosDeTrazaANMAT, w)
	if err != nil {
		return err
	}
	return err
}

func (r ListaDeCodigosDeTrazaANMAT) Serialize(w io.Writer) error {
	return writeListaDeCodigosDeTrazaANMAT(r, w)
}

func (r ListaDeCodigosDeTrazaANMAT) Schema() string {
	return "{\"fields\":[{\"name\":\"codigosDeTrazaANMAT\",\"type\":{\"items\":{\"fields\":[{\"name\":\"codigoDeTrazaANMAT\",\"type\":\"int\"}],\"name\":\"CodigosDeTrazaANMAT\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.Wapv2.Events.Record.ListaDeCodigosDeTrazaANMAT\",\"type\":\"record\"}"
}

func (r ListaDeCodigosDeTrazaANMAT) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.ListaDeCodigosDeTrazaANMAT"
}

func (_ ListaDeCodigosDeTrazaANMAT) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ListaDeCodigosDeTrazaANMAT) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ListaDeCodigosDeTrazaANMAT) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ListaDeCodigosDeTrazaANMAT) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ListaDeCodigosDeTrazaANMAT) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ListaDeCodigosDeTrazaANMAT) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ListaDeCodigosDeTrazaANMAT) SetString(v string)   { panic("Unsupported operation") }
func (_ ListaDeCodigosDeTrazaANMAT) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ListaDeCodigosDeTrazaANMAT) Get(i int) types.Field {
	switch i {
	case 0:
		r.CodigosDeTrazaANMAT = make([]CodigosDeTrazaANMAT, 0)

		w := ArrayCodigosDeTrazaANMATWrapper{Target: &r.CodigosDeTrazaANMAT}

		return w

	}
	panic("Unknown field index")
}

func (r *ListaDeCodigosDeTrazaANMAT) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ListaDeCodigosDeTrazaANMAT) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ListaDeCodigosDeTrazaANMAT) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ListaDeCodigosDeTrazaANMAT) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ListaDeCodigosDeTrazaANMAT) HintSize(int)                     { panic("Unsupported operation") }
func (_ ListaDeCodigosDeTrazaANMAT) Finalize()                        {}

func (_ ListaDeCodigosDeTrazaANMAT) AvroCRC64Fingerprint() []byte {
	return []byte(ListaDeCodigosDeTrazaANMATAvroCRC64Fingerprint)
}

func (r ListaDeCodigosDeTrazaANMAT) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["codigosDeTrazaANMAT"], err = json.Marshal(r.CodigosDeTrazaANMAT)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ListaDeCodigosDeTrazaANMAT) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["codigosDeTrazaANMAT"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigosDeTrazaANMAT); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigosDeTrazaANMAT")
	}
	return nil
}