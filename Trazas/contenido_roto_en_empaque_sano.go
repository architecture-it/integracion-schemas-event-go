// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ContenidoRotoEnEmpaqueSano.avsc
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

type ContenidoRotoEnEmpaqueSano struct {
	Traza Traza `json:"traza"`
}

const ContenidoRotoEnEmpaqueSanoAvroCRC64Fingerprint = "a)\x8e\x83&\x10\xa6Q"

func NewContenidoRotoEnEmpaqueSano() ContenidoRotoEnEmpaqueSano {
	r := ContenidoRotoEnEmpaqueSano{}
	r.Traza = NewTraza()

	return r
}

func DeserializeContenidoRotoEnEmpaqueSano(r io.Reader) (ContenidoRotoEnEmpaqueSano, error) {
	t := NewContenidoRotoEnEmpaqueSano()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeContenidoRotoEnEmpaqueSanoFromSchema(r io.Reader, schema string) (ContenidoRotoEnEmpaqueSano, error) {
	t := NewContenidoRotoEnEmpaqueSano()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeContenidoRotoEnEmpaqueSano(r ContenidoRotoEnEmpaqueSano, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r ContenidoRotoEnEmpaqueSano) Serialize(w io.Writer) error {
	return writeContenidoRotoEnEmpaqueSano(r, w)
}

func (r ContenidoRotoEnEmpaqueSano) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.ContenidoRotoEnEmpaqueSano\",\"type\":\"record\"}"
}

func (r ContenidoRotoEnEmpaqueSano) SchemaName() string {
	return "Integracion.Esquemas.Trazas.ContenidoRotoEnEmpaqueSano"
}

func (_ ContenidoRotoEnEmpaqueSano) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ContenidoRotoEnEmpaqueSano) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ContenidoRotoEnEmpaqueSano) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ContenidoRotoEnEmpaqueSano) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ContenidoRotoEnEmpaqueSano) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ContenidoRotoEnEmpaqueSano) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ContenidoRotoEnEmpaqueSano) SetString(v string)   { panic("Unsupported operation") }
func (_ ContenidoRotoEnEmpaqueSano) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ContenidoRotoEnEmpaqueSano) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *ContenidoRotoEnEmpaqueSano) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ContenidoRotoEnEmpaqueSano) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ContenidoRotoEnEmpaqueSano) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ContenidoRotoEnEmpaqueSano) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ContenidoRotoEnEmpaqueSano) HintSize(int)                     { panic("Unsupported operation") }
func (_ ContenidoRotoEnEmpaqueSano) Finalize()                        {}

func (_ ContenidoRotoEnEmpaqueSano) AvroCRC64Fingerprint() []byte {
	return []byte(ContenidoRotoEnEmpaqueSanoAvroCRC64Fingerprint)
}

func (r ContenidoRotoEnEmpaqueSano) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ContenidoRotoEnEmpaqueSano) UnmarshalJSON(data []byte) error {
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
