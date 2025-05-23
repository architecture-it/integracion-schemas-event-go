// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnvioCompletoLiberado.avsc
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

type EnvioCompletoLiberado struct {
	Traza Traza `json:"traza"`
}

const EnvioCompletoLiberadoAvroCRC64Fingerprint = "\x8d\xb3\xa2`\xb1\x9ey4"

func NewEnvioCompletoLiberado() EnvioCompletoLiberado {
	r := EnvioCompletoLiberado{}
	r.Traza = NewTraza()

	return r
}

func DeserializeEnvioCompletoLiberado(r io.Reader) (EnvioCompletoLiberado, error) {
	t := NewEnvioCompletoLiberado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnvioCompletoLiberadoFromSchema(r io.Reader, schema string) (EnvioCompletoLiberado, error) {
	t := NewEnvioCompletoLiberado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnvioCompletoLiberado(r EnvioCompletoLiberado, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r EnvioCompletoLiberado) Serialize(w io.Writer) error {
	return writeEnvioCompletoLiberado(r, w)
}

func (r EnvioCompletoLiberado) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.EnvioCompletoLiberado\",\"type\":\"record\"}"
}

func (r EnvioCompletoLiberado) SchemaName() string {
	return "Integracion.Esquemas.Trazas.EnvioCompletoLiberado"
}

func (_ EnvioCompletoLiberado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EnvioCompletoLiberado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EnvioCompletoLiberado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EnvioCompletoLiberado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EnvioCompletoLiberado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EnvioCompletoLiberado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EnvioCompletoLiberado) SetString(v string)   { panic("Unsupported operation") }
func (_ EnvioCompletoLiberado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EnvioCompletoLiberado) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *EnvioCompletoLiberado) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EnvioCompletoLiberado) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EnvioCompletoLiberado) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EnvioCompletoLiberado) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EnvioCompletoLiberado) HintSize(int)                     { panic("Unsupported operation") }
func (_ EnvioCompletoLiberado) Finalize()                        {}

func (_ EnvioCompletoLiberado) AvroCRC64Fingerprint() []byte {
	return []byte(EnvioCompletoLiberadoAvroCRC64Fingerprint)
}

func (r EnvioCompletoLiberado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EnvioCompletoLiberado) UnmarshalJSON(data []byte) error {
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
