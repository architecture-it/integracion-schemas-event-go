// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CapacidadOperativaExcedida.avsc
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

type CapacidadOperativaExcedida struct {
	Traza Traza `json:"traza"`
}

const CapacidadOperativaExcedidaAvroCRC64Fingerprint = "\fδpP<\x9b\x15"

func NewCapacidadOperativaExcedida() CapacidadOperativaExcedida {
	r := CapacidadOperativaExcedida{}
	r.Traza = NewTraza()

	return r
}

func DeserializeCapacidadOperativaExcedida(r io.Reader) (CapacidadOperativaExcedida, error) {
	t := NewCapacidadOperativaExcedida()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCapacidadOperativaExcedidaFromSchema(r io.Reader, schema string) (CapacidadOperativaExcedida, error) {
	t := NewCapacidadOperativaExcedida()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCapacidadOperativaExcedida(r CapacidadOperativaExcedida, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r CapacidadOperativaExcedida) Serialize(w io.Writer) error {
	return writeCapacidadOperativaExcedida(r, w)
}

func (r CapacidadOperativaExcedida) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.CapacidadOperativaExcedida\",\"type\":\"record\"}"
}

func (r CapacidadOperativaExcedida) SchemaName() string {
	return "Integracion.Esquemas.Trazas.CapacidadOperativaExcedida"
}

func (_ CapacidadOperativaExcedida) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CapacidadOperativaExcedida) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CapacidadOperativaExcedida) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CapacidadOperativaExcedida) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CapacidadOperativaExcedida) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CapacidadOperativaExcedida) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CapacidadOperativaExcedida) SetString(v string)   { panic("Unsupported operation") }
func (_ CapacidadOperativaExcedida) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CapacidadOperativaExcedida) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *CapacidadOperativaExcedida) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *CapacidadOperativaExcedida) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ CapacidadOperativaExcedida) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CapacidadOperativaExcedida) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CapacidadOperativaExcedida) HintSize(int)                     { panic("Unsupported operation") }
func (_ CapacidadOperativaExcedida) Finalize()                        {}

func (_ CapacidadOperativaExcedida) AvroCRC64Fingerprint() []byte {
	return []byte(CapacidadOperativaExcedidaAvroCRC64Fingerprint)
}

func (r CapacidadOperativaExcedida) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CapacidadOperativaExcedida) UnmarshalJSON(data []byte) error {
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
