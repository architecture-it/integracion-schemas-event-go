// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ContenedorRecepcionado.avsc
 */
package TrazasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ContenedorRecepcionado struct {
	Traza TrazaContendor `json:"traza"`
}

const ContenedorRecepcionadoAvroCRC64Fingerprint = "ʵ\xa5]\xba+x\xb2"

func NewContenedorRecepcionado() ContenedorRecepcionado {
	r := ContenedorRecepcionado{}
	r.Traza = NewTrazaContendor()

	return r
}

func DeserializeContenedorRecepcionado(r io.Reader) (ContenedorRecepcionado, error) {
	t := NewContenedorRecepcionado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeContenedorRecepcionadoFromSchema(r io.Reader, schema string) (ContenedorRecepcionado, error) {
	t := NewContenedorRecepcionado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeContenedorRecepcionado(r ContenedorRecepcionado, w io.Writer) error {
	var err error
	err = writeTrazaContendor(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r ContenedorRecepcionado) Serialize(w io.Writer) error {
	return writeContenedorRecepcionado(r, w)
}

func (r ContenedorRecepcionado) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"numero\",\"type\":\"string\"},{\"name\":\"tipo\",\"type\":\"string\"},{\"name\":\"ciclo\",\"type\":\"string\"},{\"name\":\"estado\",\"type\":\"string\"},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"TrazaContendor\",\"namespace\":\"Integracion.Esquemas.Contenedor.Referencias\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Contenedor.Trazas.ContenedorRecepcionado\",\"type\":\"record\"}"
}

func (r ContenedorRecepcionado) SchemaName() string {
	return "Integracion.Esquemas.Contenedor.Trazas.ContenedorRecepcionado"
}

func (_ ContenedorRecepcionado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ContenedorRecepcionado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ContenedorRecepcionado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ContenedorRecepcionado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ContenedorRecepcionado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ContenedorRecepcionado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ContenedorRecepcionado) SetString(v string)   { panic("Unsupported operation") }
func (_ ContenedorRecepcionado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ContenedorRecepcionado) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTrazaContendor()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *ContenedorRecepcionado) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ContenedorRecepcionado) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ContenedorRecepcionado) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ContenedorRecepcionado) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ContenedorRecepcionado) HintSize(int)                     { panic("Unsupported operation") }
func (_ ContenedorRecepcionado) Finalize()                        {}

func (_ ContenedorRecepcionado) AvroCRC64Fingerprint() []byte {
	return []byte(ContenedorRecepcionadoAvroCRC64Fingerprint)
}

func (r ContenedorRecepcionado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ContenedorRecepcionado) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["traza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Traza); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for traza")
	}
	return nil
}
