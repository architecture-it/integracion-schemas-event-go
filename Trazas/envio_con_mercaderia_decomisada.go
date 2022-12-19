// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnvioConMercaderiaDecomisada.avsc
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

type EnvioConMercaderiaDecomisada struct {
	Traza Traza `json:"traza"`
}

const EnvioConMercaderiaDecomisadaAvroCRC64Fingerprint = "\xf1\x88\x91^\x02\xa0\xf7\xbf"

func NewEnvioConMercaderiaDecomisada() EnvioConMercaderiaDecomisada {
	r := EnvioConMercaderiaDecomisada{}
	r.Traza = NewTraza()

	return r
}

func DeserializeEnvioConMercaderiaDecomisada(r io.Reader) (EnvioConMercaderiaDecomisada, error) {
	t := NewEnvioConMercaderiaDecomisada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnvioConMercaderiaDecomisadaFromSchema(r io.Reader, schema string) (EnvioConMercaderiaDecomisada, error) {
	t := NewEnvioConMercaderiaDecomisada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnvioConMercaderiaDecomisada(r EnvioConMercaderiaDecomisada, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r EnvioConMercaderiaDecomisada) Serialize(w io.Writer) error {
	return writeEnvioConMercaderiaDecomisada(r, w)
}

func (r EnvioConMercaderiaDecomisada) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.EnvioConMercaderiaDecomisada\",\"type\":\"record\"}"
}

func (r EnvioConMercaderiaDecomisada) SchemaName() string {
	return "Integracion.Esquemas.Trazas.EnvioConMercaderiaDecomisada"
}

func (_ EnvioConMercaderiaDecomisada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EnvioConMercaderiaDecomisada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EnvioConMercaderiaDecomisada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EnvioConMercaderiaDecomisada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EnvioConMercaderiaDecomisada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EnvioConMercaderiaDecomisada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EnvioConMercaderiaDecomisada) SetString(v string)   { panic("Unsupported operation") }
func (_ EnvioConMercaderiaDecomisada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EnvioConMercaderiaDecomisada) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *EnvioConMercaderiaDecomisada) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EnvioConMercaderiaDecomisada) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EnvioConMercaderiaDecomisada) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ EnvioConMercaderiaDecomisada) AppendArray() types.Field { panic("Unsupported operation") }
func (_ EnvioConMercaderiaDecomisada) HintSize(int)             { panic("Unsupported operation") }
func (_ EnvioConMercaderiaDecomisada) Finalize()                {}

func (_ EnvioConMercaderiaDecomisada) AvroCRC64Fingerprint() []byte {
	return []byte(EnvioConMercaderiaDecomisadaAvroCRC64Fingerprint)
}

func (r EnvioConMercaderiaDecomisada) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EnvioConMercaderiaDecomisada) UnmarshalJSON(data []byte) error {
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
