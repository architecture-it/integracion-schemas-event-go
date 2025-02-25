// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnvioDeMuestraMedica.avsc
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

type EnvioDeMuestraMedica struct {
	Traza Traza `json:"traza"`
}

const EnvioDeMuestraMedicaAvroCRC64Fingerprint = "\xf5\xc5u\xa0\xaaŏ`"

func NewEnvioDeMuestraMedica() EnvioDeMuestraMedica {
	r := EnvioDeMuestraMedica{}
	r.Traza = NewTraza()

	return r
}

func DeserializeEnvioDeMuestraMedica(r io.Reader) (EnvioDeMuestraMedica, error) {
	t := NewEnvioDeMuestraMedica()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnvioDeMuestraMedicaFromSchema(r io.Reader, schema string) (EnvioDeMuestraMedica, error) {
	t := NewEnvioDeMuestraMedica()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnvioDeMuestraMedica(r EnvioDeMuestraMedica, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r EnvioDeMuestraMedica) Serialize(w io.Writer) error {
	return writeEnvioDeMuestraMedica(r, w)
}

func (r EnvioDeMuestraMedica) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.EnvioDeMuestraMedica\",\"type\":\"record\"}"
}

func (r EnvioDeMuestraMedica) SchemaName() string {
	return "Integracion.Esquemas.Trazas.EnvioDeMuestraMedica"
}

func (_ EnvioDeMuestraMedica) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EnvioDeMuestraMedica) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EnvioDeMuestraMedica) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EnvioDeMuestraMedica) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EnvioDeMuestraMedica) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EnvioDeMuestraMedica) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EnvioDeMuestraMedica) SetString(v string)   { panic("Unsupported operation") }
func (_ EnvioDeMuestraMedica) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EnvioDeMuestraMedica) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *EnvioDeMuestraMedica) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EnvioDeMuestraMedica) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EnvioDeMuestraMedica) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EnvioDeMuestraMedica) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EnvioDeMuestraMedica) HintSize(int)                     { panic("Unsupported operation") }
func (_ EnvioDeMuestraMedica) Finalize()                        {}

func (_ EnvioDeMuestraMedica) AvroCRC64Fingerprint() []byte {
	return []byte(EnvioDeMuestraMedicaAvroCRC64Fingerprint)
}

func (r EnvioDeMuestraMedica) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EnvioDeMuestraMedica) UnmarshalJSON(data []byte) error {
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
