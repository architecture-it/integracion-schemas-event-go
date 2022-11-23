// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     RendicionConRemito.avsc
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

type RendicionConRemito struct {
	Traza Traza `json:"traza"`
}

const RendicionConRemitoAvroCRC64Fingerprint = "\x1ei\xe5\xa6'W\x97\xe8"

func NewRendicionConRemito() RendicionConRemito {
	r := RendicionConRemito{}
	r.Traza = NewTraza()

	return r
}

func DeserializeRendicionConRemito(r io.Reader) (RendicionConRemito, error) {
	t := NewRendicionConRemito()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRendicionConRemitoFromSchema(r io.Reader, schema string) (RendicionConRemito, error) {
	t := NewRendicionConRemito()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRendicionConRemito(r RendicionConRemito, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r RendicionConRemito) Serialize(w io.Writer) error {
	return writeRendicionConRemito(r, w)
}

func (r RendicionConRemito) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.RendicionConRemito\",\"type\":\"record\"}"
}

func (r RendicionConRemito) SchemaName() string {
	return "Integracion.Esquemas.Trazas.RendicionConRemito"
}

func (_ RendicionConRemito) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ RendicionConRemito) SetInt(v int32)       { panic("Unsupported operation") }
func (_ RendicionConRemito) SetLong(v int64)      { panic("Unsupported operation") }
func (_ RendicionConRemito) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ RendicionConRemito) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ RendicionConRemito) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ RendicionConRemito) SetString(v string)   { panic("Unsupported operation") }
func (_ RendicionConRemito) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RendicionConRemito) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *RendicionConRemito) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *RendicionConRemito) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ RendicionConRemito) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ RendicionConRemito) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ RendicionConRemito) HintSize(int)                     { panic("Unsupported operation") }
func (_ RendicionConRemito) Finalize()                        {}

func (_ RendicionConRemito) AvroCRC64Fingerprint() []byte {
	return []byte(RendicionConRemitoAvroCRC64Fingerprint)
}

func (r RendicionConRemito) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *RendicionConRemito) UnmarshalJSON(data []byte) error {
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
