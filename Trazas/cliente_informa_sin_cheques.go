// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ClienteInformaSinCheques.avsc
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

type ClienteInformaSinCheques struct {
	Traza Traza `json:"traza"`
}

const ClienteInformaSinChequesAvroCRC64Fingerprint = "\x97.g\xa9a\xa7\x9e,"

func NewClienteInformaSinCheques() ClienteInformaSinCheques {
	r := ClienteInformaSinCheques{}
	r.Traza = NewTraza()

	return r
}

func DeserializeClienteInformaSinCheques(r io.Reader) (ClienteInformaSinCheques, error) {
	t := NewClienteInformaSinCheques()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeClienteInformaSinChequesFromSchema(r io.Reader, schema string) (ClienteInformaSinCheques, error) {
	t := NewClienteInformaSinCheques()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeClienteInformaSinCheques(r ClienteInformaSinCheques, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r ClienteInformaSinCheques) Serialize(w io.Writer) error {
	return writeClienteInformaSinCheques(r, w)
}

func (r ClienteInformaSinCheques) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.ClienteInformaSinCheques\",\"type\":\"record\"}"
}

func (r ClienteInformaSinCheques) SchemaName() string {
	return "Integracion.Esquemas.Trazas.ClienteInformaSinCheques"
}

func (_ ClienteInformaSinCheques) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ClienteInformaSinCheques) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ClienteInformaSinCheques) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ClienteInformaSinCheques) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ClienteInformaSinCheques) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ClienteInformaSinCheques) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ClienteInformaSinCheques) SetString(v string)   { panic("Unsupported operation") }
func (_ ClienteInformaSinCheques) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ClienteInformaSinCheques) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *ClienteInformaSinCheques) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ClienteInformaSinCheques) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ClienteInformaSinCheques) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ClienteInformaSinCheques) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ClienteInformaSinCheques) HintSize(int)                     { panic("Unsupported operation") }
func (_ ClienteInformaSinCheques) Finalize()                        {}

func (_ ClienteInformaSinCheques) AvroCRC64Fingerprint() []byte {
	return []byte(ClienteInformaSinChequesAvroCRC64Fingerprint)
}

func (r ClienteInformaSinCheques) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ClienteInformaSinCheques) UnmarshalJSON(data []byte) error {
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