// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     RoturaParcial.avsc
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

type RoturaParcial struct {
	Traza Traza `json:"traza"`
}

const RoturaParcialAvroCRC64Fingerprint = "ѬM\x10\x0f\x8c\xa6z"

func NewRoturaParcial() RoturaParcial {
	r := RoturaParcial{}
	r.Traza = NewTraza()

	return r
}

func DeserializeRoturaParcial(r io.Reader) (RoturaParcial, error) {
	t := NewRoturaParcial()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRoturaParcialFromSchema(r io.Reader, schema string) (RoturaParcial, error) {
	t := NewRoturaParcial()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRoturaParcial(r RoturaParcial, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r RoturaParcial) Serialize(w io.Writer) error {
	return writeRoturaParcial(r, w)
}

func (r RoturaParcial) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.RoturaParcial\",\"type\":\"record\"}"
}

func (r RoturaParcial) SchemaName() string {
	return "Integracion.Esquemas.Trazas.RoturaParcial"
}

func (_ RoturaParcial) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ RoturaParcial) SetInt(v int32)       { panic("Unsupported operation") }
func (_ RoturaParcial) SetLong(v int64)      { panic("Unsupported operation") }
func (_ RoturaParcial) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ RoturaParcial) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ RoturaParcial) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ RoturaParcial) SetString(v string)   { panic("Unsupported operation") }
func (_ RoturaParcial) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RoturaParcial) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *RoturaParcial) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *RoturaParcial) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ RoturaParcial) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ RoturaParcial) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ RoturaParcial) HintSize(int)                     { panic("Unsupported operation") }
func (_ RoturaParcial) Finalize()                        {}

func (_ RoturaParcial) AvroCRC64Fingerprint() []byte {
	return []byte(RoturaParcialAvroCRC64Fingerprint)
}

func (r RoturaParcial) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *RoturaParcial) UnmarshalJSON(data []byte) error {
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
