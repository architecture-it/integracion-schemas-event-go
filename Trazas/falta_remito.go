// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     FaltaRemito.avsc
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

type FaltaRemito struct {
	Traza Traza `json:"traza"`
}

const FaltaRemitoAvroCRC64Fingerprint = "lu\xd8\x15u\x811\xe7"

func NewFaltaRemito() FaltaRemito {
	r := FaltaRemito{}
	r.Traza = NewTraza()

	return r
}

func DeserializeFaltaRemito(r io.Reader) (FaltaRemito, error) {
	t := NewFaltaRemito()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeFaltaRemitoFromSchema(r io.Reader, schema string) (FaltaRemito, error) {
	t := NewFaltaRemito()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeFaltaRemito(r FaltaRemito, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r FaltaRemito) Serialize(w io.Writer) error {
	return writeFaltaRemito(r, w)
}

func (r FaltaRemito) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.FaltaRemito\",\"type\":\"record\"}"
}

func (r FaltaRemito) SchemaName() string {
	return "Integracion.Esquemas.Trazas.FaltaRemito"
}

func (_ FaltaRemito) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ FaltaRemito) SetInt(v int32)       { panic("Unsupported operation") }
func (_ FaltaRemito) SetLong(v int64)      { panic("Unsupported operation") }
func (_ FaltaRemito) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ FaltaRemito) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ FaltaRemito) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ FaltaRemito) SetString(v string)   { panic("Unsupported operation") }
func (_ FaltaRemito) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *FaltaRemito) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *FaltaRemito) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *FaltaRemito) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ FaltaRemito) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ FaltaRemito) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ FaltaRemito) HintSize(int)                     { panic("Unsupported operation") }
func (_ FaltaRemito) Finalize()                        {}

func (_ FaltaRemito) AvroCRC64Fingerprint() []byte {
	return []byte(FaltaRemitoAvroCRC64Fingerprint)
}

func (r FaltaRemito) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *FaltaRemito) UnmarshalJSON(data []byte) error {
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
