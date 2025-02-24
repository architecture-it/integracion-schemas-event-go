// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnvioProcesadoEnSorter.avsc
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

type EnvioProcesadoEnSorter struct {
	Traza Traza `json:"traza"`
}

const EnvioProcesadoEnSorterAvroCRC64Fingerprint = "\xbd\x85V\xa4\xf3\x1eU\xd1"

func NewEnvioProcesadoEnSorter() EnvioProcesadoEnSorter {
	r := EnvioProcesadoEnSorter{}
	r.Traza = NewTraza()

	return r
}

func DeserializeEnvioProcesadoEnSorter(r io.Reader) (EnvioProcesadoEnSorter, error) {
	t := NewEnvioProcesadoEnSorter()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnvioProcesadoEnSorterFromSchema(r io.Reader, schema string) (EnvioProcesadoEnSorter, error) {
	t := NewEnvioProcesadoEnSorter()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnvioProcesadoEnSorter(r EnvioProcesadoEnSorter, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r EnvioProcesadoEnSorter) Serialize(w io.Writer) error {
	return writeEnvioProcesadoEnSorter(r, w)
}

func (r EnvioProcesadoEnSorter) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"LogisticShipmentType\",\"type\":[\"null\",\"string\"]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.EnvioProcesadoEnSorter\",\"type\":\"record\"}"
}

func (r EnvioProcesadoEnSorter) SchemaName() string {
	return "Integracion.Esquemas.Trazas.EnvioProcesadoEnSorter"
}

func (_ EnvioProcesadoEnSorter) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EnvioProcesadoEnSorter) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EnvioProcesadoEnSorter) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EnvioProcesadoEnSorter) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EnvioProcesadoEnSorter) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EnvioProcesadoEnSorter) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EnvioProcesadoEnSorter) SetString(v string)   { panic("Unsupported operation") }
func (_ EnvioProcesadoEnSorter) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EnvioProcesadoEnSorter) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *EnvioProcesadoEnSorter) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EnvioProcesadoEnSorter) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EnvioProcesadoEnSorter) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EnvioProcesadoEnSorter) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EnvioProcesadoEnSorter) HintSize(int)                     { panic("Unsupported operation") }
func (_ EnvioProcesadoEnSorter) Finalize()                        {}

func (_ EnvioProcesadoEnSorter) AvroCRC64Fingerprint() []byte {
	return []byte(EnvioProcesadoEnSorterAvroCRC64Fingerprint)
}

func (r EnvioProcesadoEnSorter) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EnvioProcesadoEnSorter) UnmarshalJSON(data []byte) error {
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
