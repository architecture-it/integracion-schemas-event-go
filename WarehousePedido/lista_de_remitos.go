// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoDeAlmacenSCE.avsc
 */
package WarehousePedidoEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ListaDeRemitos struct {
	Documentos []Remito `json:"documentos"`
}

const ListaDeRemitosAvroCRC64Fingerprint = "\xe2\xd2\xcb;u\xce\\\x1b"

func NewListaDeRemitos() ListaDeRemitos {
	r := ListaDeRemitos{}
	r.Documentos = make([]Remito, 0)

	return r
}

func DeserializeListaDeRemitos(r io.Reader) (ListaDeRemitos, error) {
	t := NewListaDeRemitos()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeListaDeRemitosFromSchema(r io.Reader, schema string) (ListaDeRemitos, error) {
	t := NewListaDeRemitos()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeListaDeRemitos(r ListaDeRemitos, w io.Writer) error {
	var err error
	err = writeArrayRemito(r.Documentos, w)
	if err != nil {
		return err
	}
	return err
}

func (r ListaDeRemitos) Serialize(w io.Writer) error {
	return writeListaDeRemitos(r, w)
}

func (r ListaDeRemitos) Schema() string {
	return "{\"fields\":[{\"name\":\"documentos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"numeroEntidad\",\"type\":\"string\"},{\"name\":\"documento\",\"type\":\"string\"}],\"name\":\"Remito\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.WarehousePedido.Events.Record.ListaDeRemitos\",\"type\":\"record\"}"
}

func (r ListaDeRemitos) SchemaName() string {
	return "Andreani.WarehousePedido.Events.Record.ListaDeRemitos"
}

func (_ ListaDeRemitos) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ListaDeRemitos) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ListaDeRemitos) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ListaDeRemitos) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ListaDeRemitos) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ListaDeRemitos) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ListaDeRemitos) SetString(v string)   { panic("Unsupported operation") }
func (_ ListaDeRemitos) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ListaDeRemitos) Get(i int) types.Field {
	switch i {
	case 0:
		r.Documentos = make([]Remito, 0)

		w := ArrayRemitoWrapper{Target: &r.Documentos}

		return w

	}
	panic("Unknown field index")
}

func (r *ListaDeRemitos) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ListaDeRemitos) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ListaDeRemitos) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ListaDeRemitos) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ListaDeRemitos) HintSize(int)                     { panic("Unsupported operation") }
func (_ ListaDeRemitos) Finalize()                        {}

func (_ ListaDeRemitos) AvroCRC64Fingerprint() []byte {
	return []byte(ListaDeRemitosAvroCRC64Fingerprint)
}

func (r ListaDeRemitos) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["documentos"], err = json.Marshal(r.Documentos)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ListaDeRemitos) UnmarshalJSON(data []byte) error {
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
