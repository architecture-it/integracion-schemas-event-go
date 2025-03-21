// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ExpedicionEnReparto.avsc
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

type ExpedicionEnReparto struct {
	Traza Traza `json:"traza"`
}

const ExpedicionEnRepartoAvroCRC64Fingerprint = "w\xad\xda\xea\x04\xb6j\n"

func NewExpedicionEnReparto() ExpedicionEnReparto {
	r := ExpedicionEnReparto{}
	r.Traza = NewTraza()

	return r
}

func DeserializeExpedicionEnReparto(r io.Reader) (ExpedicionEnReparto, error) {
	t := NewExpedicionEnReparto()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeExpedicionEnRepartoFromSchema(r io.Reader, schema string) (ExpedicionEnReparto, error) {
	t := NewExpedicionEnReparto()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeExpedicionEnReparto(r ExpedicionEnReparto, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r ExpedicionEnReparto) Serialize(w io.Writer) error {
	return writeExpedicionEnReparto(r, w)
}

func (r ExpedicionEnReparto) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.ExpedicionEnReparto\",\"type\":\"record\"}"
}

func (r ExpedicionEnReparto) SchemaName() string {
	return "Integracion.Esquemas.Trazas.ExpedicionEnReparto"
}

func (_ ExpedicionEnReparto) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ExpedicionEnReparto) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ExpedicionEnReparto) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ExpedicionEnReparto) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ExpedicionEnReparto) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ExpedicionEnReparto) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ExpedicionEnReparto) SetString(v string)   { panic("Unsupported operation") }
func (_ ExpedicionEnReparto) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ExpedicionEnReparto) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *ExpedicionEnReparto) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ExpedicionEnReparto) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ExpedicionEnReparto) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ExpedicionEnReparto) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ExpedicionEnReparto) HintSize(int)                     { panic("Unsupported operation") }
func (_ ExpedicionEnReparto) Finalize()                        {}

func (_ ExpedicionEnReparto) AvroCRC64Fingerprint() []byte {
	return []byte(ExpedicionEnRepartoAvroCRC64Fingerprint)
}

func (r ExpedicionEnReparto) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ExpedicionEnReparto) UnmarshalJSON(data []byte) error {
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
