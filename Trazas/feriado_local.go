// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     FeriadoLocal.avsc
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

type FeriadoLocal struct {
	Traza Traza `json:"traza"`
}

const FeriadoLocalAvroCRC64Fingerprint = "_\xc5{\xe0E\xc3ˏ"

func NewFeriadoLocal() FeriadoLocal {
	r := FeriadoLocal{}
	r.Traza = NewTraza()

	return r
}

func DeserializeFeriadoLocal(r io.Reader) (FeriadoLocal, error) {
	t := NewFeriadoLocal()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeFeriadoLocalFromSchema(r io.Reader, schema string) (FeriadoLocal, error) {
	t := NewFeriadoLocal()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeFeriadoLocal(r FeriadoLocal, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r FeriadoLocal) Serialize(w io.Writer) error {
	return writeFeriadoLocal(r, w)
}

func (r FeriadoLocal) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.FeriadoLocal\",\"type\":\"record\"}"
}

func (r FeriadoLocal) SchemaName() string {
	return "Integracion.Esquemas.Trazas.FeriadoLocal"
}

func (_ FeriadoLocal) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ FeriadoLocal) SetInt(v int32)       { panic("Unsupported operation") }
func (_ FeriadoLocal) SetLong(v int64)      { panic("Unsupported operation") }
func (_ FeriadoLocal) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ FeriadoLocal) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ FeriadoLocal) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ FeriadoLocal) SetString(v string)   { panic("Unsupported operation") }
func (_ FeriadoLocal) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *FeriadoLocal) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *FeriadoLocal) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *FeriadoLocal) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ FeriadoLocal) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ FeriadoLocal) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ FeriadoLocal) HintSize(int)                     { panic("Unsupported operation") }
func (_ FeriadoLocal) Finalize()                        {}

func (_ FeriadoLocal) AvroCRC64Fingerprint() []byte {
	return []byte(FeriadoLocalAvroCRC64Fingerprint)
}

func (r FeriadoLocal) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *FeriadoLocal) UnmarshalJSON(data []byte) error {
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