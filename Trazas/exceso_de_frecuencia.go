// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ExcesoDeFrecuencia.avsc
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

type ExcesoDeFrecuencia struct {
	Traza Traza `json:"traza"`
}

const ExcesoDeFrecuenciaAvroCRC64Fingerprint = "\x9c[\xa8\x99\xa9\ue85e"

func NewExcesoDeFrecuencia() ExcesoDeFrecuencia {
	r := ExcesoDeFrecuencia{}
	r.Traza = NewTraza()

	return r
}

func DeserializeExcesoDeFrecuencia(r io.Reader) (ExcesoDeFrecuencia, error) {
	t := NewExcesoDeFrecuencia()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeExcesoDeFrecuenciaFromSchema(r io.Reader, schema string) (ExcesoDeFrecuencia, error) {
	t := NewExcesoDeFrecuencia()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeExcesoDeFrecuencia(r ExcesoDeFrecuencia, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r ExcesoDeFrecuencia) Serialize(w io.Writer) error {
	return writeExcesoDeFrecuencia(r, w)
}

func (r ExcesoDeFrecuencia) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.ExcesoDeFrecuencia\",\"type\":\"record\"}"
}

func (r ExcesoDeFrecuencia) SchemaName() string {
	return "Integracion.Esquemas.Trazas.ExcesoDeFrecuencia"
}

func (_ ExcesoDeFrecuencia) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ExcesoDeFrecuencia) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ExcesoDeFrecuencia) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ExcesoDeFrecuencia) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ExcesoDeFrecuencia) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ExcesoDeFrecuencia) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ExcesoDeFrecuencia) SetString(v string)   { panic("Unsupported operation") }
func (_ ExcesoDeFrecuencia) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ExcesoDeFrecuencia) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *ExcesoDeFrecuencia) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ExcesoDeFrecuencia) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ExcesoDeFrecuencia) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ExcesoDeFrecuencia) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ExcesoDeFrecuencia) HintSize(int)                     { panic("Unsupported operation") }
func (_ ExcesoDeFrecuencia) Finalize()                        {}

func (_ ExcesoDeFrecuencia) AvroCRC64Fingerprint() []byte {
	return []byte(ExcesoDeFrecuenciaAvroCRC64Fingerprint)
}

func (r ExcesoDeFrecuencia) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ExcesoDeFrecuencia) UnmarshalJSON(data []byte) error {
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