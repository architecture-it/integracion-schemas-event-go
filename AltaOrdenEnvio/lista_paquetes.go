// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     OrdenDeEnvioSolicitada.avsc
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

var _ = fmt.Printf

type ListaPaquetes struct {
	ListaPaquetes *UnionNullArrayDetalleDePaquete `json:"listaPaquetes"`
}

const ListaPaquetesAvroCRC64Fingerprint = "\x9e\xc0\xb2\xd0\x06Y4\xad"

func NewListaPaquetes() ListaPaquetes {
	r := ListaPaquetes{}
	return r
}

func DeserializeListaPaquetes(r io.Reader) (ListaPaquetes, error) {
	t := NewListaPaquetes()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeListaPaquetesFromSchema(r io.Reader, schema string) (ListaPaquetes, error) {
	t := NewListaPaquetes()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeListaPaquetes(r ListaPaquetes, w io.Writer) error {
	var err error
	err = writeUnionNullArrayDetalleDePaquete(r.ListaPaquetes, w)
	if err != nil {
		return err
	}
	return err
}

func (r ListaPaquetes) Serialize(w io.Writer) error {
	return writeListaPaquetes(r, w)
}

func (r ListaPaquetes) Schema() string {
	return "{\"fields\":[{\"name\":\"listaPaquetes\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"pesoEnKg\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"altoEnCm\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"anchoEnCm\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"largoEnCm\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"referenciasDelCliente\",\"type\":[\"null\",{\"type\":\"map\",\"values\":\"string\"}]},{\"default\":null,\"name\":\"volumenEnCm3\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"valorDeclaradoSinImpuesto\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"valorDeclaradoConImpuesto\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"numeroDeBulto\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"valorDeclarado\",\"type\":[\"null\",\"double\"]}],\"name\":\"DetalleDePaquete\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"Andreani.AltaOrdenEnvio.Events.Common.ListaPaquetes\",\"type\":\"record\"}"
}

func (r ListaPaquetes) SchemaName() string {
	return "Andreani.AltaOrdenEnvio.Events.Common.ListaPaquetes"
}

func (_ ListaPaquetes) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ListaPaquetes) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ListaPaquetes) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ListaPaquetes) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ListaPaquetes) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ListaPaquetes) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ListaPaquetes) SetString(v string)   { panic("Unsupported operation") }
func (_ ListaPaquetes) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ListaPaquetes) Get(i int) types.Field {
	switch i {
	case 0:
		r.ListaPaquetes = NewUnionNullArrayDetalleDePaquete()

		return r.ListaPaquetes
	}
	panic("Unknown field index")
}

func (r *ListaPaquetes) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ListaPaquetes) NullField(i int) {
	switch i {
	case 0:
		r.ListaPaquetes = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ListaPaquetes) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ListaPaquetes) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ListaPaquetes) HintSize(int)                     { panic("Unsupported operation") }
func (_ ListaPaquetes) Finalize()                        {}

func (_ ListaPaquetes) AvroCRC64Fingerprint() []byte {
	return []byte(ListaPaquetesAvroCRC64Fingerprint)
}

func (r ListaPaquetes) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["listaPaquetes"], err = json.Marshal(r.ListaPaquetes)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ListaPaquetes) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["listaPaquetes"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ListaPaquetes); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for listaPaquetes")
	}
	return nil
}
