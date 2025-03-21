// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EntregaPactadaEnReparto.avsc
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

type EntregaPactadaEnReparto struct {
	Traza Traza `json:"traza"`
}

const EntregaPactadaEnRepartoAvroCRC64Fingerprint = "\x89#\x7f\xd0\xf5q\xc1S"

func NewEntregaPactadaEnReparto() EntregaPactadaEnReparto {
	r := EntregaPactadaEnReparto{}
	r.Traza = NewTraza()

	return r
}

func DeserializeEntregaPactadaEnReparto(r io.Reader) (EntregaPactadaEnReparto, error) {
	t := NewEntregaPactadaEnReparto()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEntregaPactadaEnRepartoFromSchema(r io.Reader, schema string) (EntregaPactadaEnReparto, error) {
	t := NewEntregaPactadaEnReparto()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEntregaPactadaEnReparto(r EntregaPactadaEnReparto, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r EntregaPactadaEnReparto) Serialize(w io.Writer) error {
	return writeEntregaPactadaEnReparto(r, w)
}

func (r EntregaPactadaEnReparto) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.EntregaPactadaEnReparto\",\"type\":\"record\"}"
}

func (r EntregaPactadaEnReparto) SchemaName() string {
	return "Integracion.Esquemas.Trazas.EntregaPactadaEnReparto"
}

func (_ EntregaPactadaEnReparto) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EntregaPactadaEnReparto) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EntregaPactadaEnReparto) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EntregaPactadaEnReparto) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EntregaPactadaEnReparto) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EntregaPactadaEnReparto) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EntregaPactadaEnReparto) SetString(v string)   { panic("Unsupported operation") }
func (_ EntregaPactadaEnReparto) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EntregaPactadaEnReparto) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *EntregaPactadaEnReparto) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EntregaPactadaEnReparto) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EntregaPactadaEnReparto) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EntregaPactadaEnReparto) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EntregaPactadaEnReparto) HintSize(int)                     { panic("Unsupported operation") }
func (_ EntregaPactadaEnReparto) Finalize()                        {}

func (_ EntregaPactadaEnReparto) AvroCRC64Fingerprint() []byte {
	return []byte(EntregaPactadaEnRepartoAvroCRC64Fingerprint)
}

func (r EntregaPactadaEnReparto) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EntregaPactadaEnReparto) UnmarshalJSON(data []byte) error {
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
