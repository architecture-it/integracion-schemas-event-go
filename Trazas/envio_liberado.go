// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnvioLiberado.avsc
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

type EnvioLiberado struct {
	Traza Traza `json:"traza"`
}

const EnvioLiberadoAvroCRC64Fingerprint = "\x1a4\xd1/\xf8\xbe2\xe2"

func NewEnvioLiberado() EnvioLiberado {
	r := EnvioLiberado{}
	r.Traza = NewTraza()

	return r
}

func DeserializeEnvioLiberado(r io.Reader) (EnvioLiberado, error) {
	t := NewEnvioLiberado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnvioLiberadoFromSchema(r io.Reader, schema string) (EnvioLiberado, error) {
	t := NewEnvioLiberado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnvioLiberado(r EnvioLiberado, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r EnvioLiberado) Serialize(w io.Writer) error {
	return writeEnvioLiberado(r, w)
}

func (r EnvioLiberado) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"LogisticShipmentType\",\"type\":[\"null\",\"string\"]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.EnvioLiberado\",\"type\":\"record\"}"
}

func (r EnvioLiberado) SchemaName() string {
	return "Integracion.Esquemas.Trazas.EnvioLiberado"
}

func (_ EnvioLiberado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EnvioLiberado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EnvioLiberado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EnvioLiberado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EnvioLiberado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EnvioLiberado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EnvioLiberado) SetString(v string)   { panic("Unsupported operation") }
func (_ EnvioLiberado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EnvioLiberado) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *EnvioLiberado) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EnvioLiberado) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EnvioLiberado) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EnvioLiberado) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EnvioLiberado) HintSize(int)                     { panic("Unsupported operation") }
func (_ EnvioLiberado) Finalize()                        {}

func (_ EnvioLiberado) AvroCRC64Fingerprint() []byte {
	return []byte(EnvioLiberadoAvroCRC64Fingerprint)
}

func (r EnvioLiberado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EnvioLiberado) UnmarshalJSON(data []byte) error {
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
