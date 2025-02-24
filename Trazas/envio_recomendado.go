// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnvioRecomendado.avsc
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

type EnvioRecomendado struct {
	Traza Traza `json:"traza"`
}

const EnvioRecomendadoAvroCRC64Fingerprint = "\xbfM\xbe\x1a\x0f\x02\\\x9e"

func NewEnvioRecomendado() EnvioRecomendado {
	r := EnvioRecomendado{}
	r.Traza = NewTraza()

	return r
}

func DeserializeEnvioRecomendado(r io.Reader) (EnvioRecomendado, error) {
	t := NewEnvioRecomendado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnvioRecomendadoFromSchema(r io.Reader, schema string) (EnvioRecomendado, error) {
	t := NewEnvioRecomendado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnvioRecomendado(r EnvioRecomendado, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r EnvioRecomendado) Serialize(w io.Writer) error {
	return writeEnvioRecomendado(r, w)
}

func (r EnvioRecomendado) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"LogisticShipmentType\",\"type\":[\"null\",\"string\"]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.EnvioRecomendado\",\"type\":\"record\"}"
}

func (r EnvioRecomendado) SchemaName() string {
	return "Integracion.Esquemas.Trazas.EnvioRecomendado"
}

func (_ EnvioRecomendado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EnvioRecomendado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EnvioRecomendado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EnvioRecomendado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EnvioRecomendado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EnvioRecomendado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EnvioRecomendado) SetString(v string)   { panic("Unsupported operation") }
func (_ EnvioRecomendado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EnvioRecomendado) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *EnvioRecomendado) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EnvioRecomendado) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EnvioRecomendado) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EnvioRecomendado) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EnvioRecomendado) HintSize(int)                     { panic("Unsupported operation") }
func (_ EnvioRecomendado) Finalize()                        {}

func (_ EnvioRecomendado) AvroCRC64Fingerprint() []byte {
	return []byte(EnvioRecomendadoAvroCRC64Fingerprint)
}

func (r EnvioRecomendado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EnvioRecomendado) UnmarshalJSON(data []byte) error {
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
