// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     SobranBultos.avsc
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

type SobranBultos struct {
	Traza Traza `json:"traza"`
}

const SobranBultosAvroCRC64Fingerprint = "lң@\xbf\"\x1aU"

func NewSobranBultos() SobranBultos {
	r := SobranBultos{}
	r.Traza = NewTraza()

	return r
}

func DeserializeSobranBultos(r io.Reader) (SobranBultos, error) {
	t := NewSobranBultos()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSobranBultosFromSchema(r io.Reader, schema string) (SobranBultos, error) {
	t := NewSobranBultos()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSobranBultos(r SobranBultos, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r SobranBultos) Serialize(w io.Writer) error {
	return writeSobranBultos(r, w)
}

func (r SobranBultos) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.SobranBultos\",\"type\":\"record\"}"
}

func (r SobranBultos) SchemaName() string {
	return "Integracion.Esquemas.Trazas.SobranBultos"
}

func (_ SobranBultos) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ SobranBultos) SetInt(v int32)       { panic("Unsupported operation") }
func (_ SobranBultos) SetLong(v int64)      { panic("Unsupported operation") }
func (_ SobranBultos) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ SobranBultos) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ SobranBultos) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ SobranBultos) SetString(v string)   { panic("Unsupported operation") }
func (_ SobranBultos) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *SobranBultos) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *SobranBultos) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *SobranBultos) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ SobranBultos) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ SobranBultos) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ SobranBultos) HintSize(int)                     { panic("Unsupported operation") }
func (_ SobranBultos) Finalize()                        {}

func (_ SobranBultos) AvroCRC64Fingerprint() []byte {
	return []byte(SobranBultosAvroCRC64Fingerprint)
}

func (r SobranBultos) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *SobranBultos) UnmarshalJSON(data []byte) error {
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
