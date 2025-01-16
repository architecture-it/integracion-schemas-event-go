// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Digitalizacion.avsc
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

type Digitalizacion struct {
	Traza Traza `json:"traza"`
}

const DigitalizacionAvroCRC64Fingerprint = "\x05\xf3\x9di-\x19֠"

func NewDigitalizacion() Digitalizacion {
	r := Digitalizacion{}
	r.Traza = NewTraza()

	return r
}

func DeserializeDigitalizacion(r io.Reader) (Digitalizacion, error) {
	t := NewDigitalizacion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDigitalizacionFromSchema(r io.Reader, schema string) (Digitalizacion, error) {
	t := NewDigitalizacion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDigitalizacion(r Digitalizacion, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r Digitalizacion) Serialize(w io.Writer) error {
	return writeDigitalizacion(r, w)
}

func (r Digitalizacion) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.Digitalizacion\",\"type\":\"record\"}"
}

func (r Digitalizacion) SchemaName() string {
	return "Integracion.Esquemas.Trazas.Digitalizacion"
}

func (_ Digitalizacion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Digitalizacion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Digitalizacion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Digitalizacion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Digitalizacion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Digitalizacion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Digitalizacion) SetString(v string)   { panic("Unsupported operation") }
func (_ Digitalizacion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Digitalizacion) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *Digitalizacion) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Digitalizacion) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Digitalizacion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Digitalizacion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Digitalizacion) HintSize(int)                     { panic("Unsupported operation") }
func (_ Digitalizacion) Finalize()                        {}

func (_ Digitalizacion) AvroCRC64Fingerprint() []byte {
	return []byte(DigitalizacionAvroCRC64Fingerprint)
}

func (r Digitalizacion) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Digitalizacion) UnmarshalJSON(data []byte) error {
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
