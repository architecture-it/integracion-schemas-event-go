// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     BultoRoturaParcial.avsc
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

type BultoRoturaParcial struct {
	Traza TrazaDeBulto `json:"traza"`
}

const BultoRoturaParcialAvroCRC64Fingerprint = ",j\x81W\xf8\xa2\xbf\xc7"

func NewBultoRoturaParcial() BultoRoturaParcial {
	r := BultoRoturaParcial{}
	r.Traza = NewTrazaDeBulto()

	return r
}

func DeserializeBultoRoturaParcial(r io.Reader) (BultoRoturaParcial, error) {
	t := NewBultoRoturaParcial()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBultoRoturaParcialFromSchema(r io.Reader, schema string) (BultoRoturaParcial, error) {
	t := NewBultoRoturaParcial()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBultoRoturaParcial(r BultoRoturaParcial, w io.Writer) error {
	var err error
	err = writeTrazaDeBulto(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r BultoRoturaParcial) Serialize(w io.Writer) error {
	return writeBultoRoturaParcial(r, w)
}

func (r BultoRoturaParcial) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"name\":\"numeroDeBulto\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"tipoDeBulto\",\"type\":[\"null\",{\"name\":\"TipoDeBulto\",\"symbols\":[\"Paquete\",\"Contenedor\"],\"type\":\"enum\"}]}],\"name\":\"TrazaDeBulto\",\"namespace\":\"Integracion.Esquemas.Bulto.Referencias\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Bulto.Trazas.BultoRoturaParcial\",\"type\":\"record\"}"
}

func (r BultoRoturaParcial) SchemaName() string {
	return "Integracion.Esquemas.Bulto.Trazas.BultoRoturaParcial"
}

func (_ BultoRoturaParcial) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BultoRoturaParcial) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BultoRoturaParcial) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BultoRoturaParcial) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BultoRoturaParcial) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BultoRoturaParcial) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BultoRoturaParcial) SetString(v string)   { panic("Unsupported operation") }
func (_ BultoRoturaParcial) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BultoRoturaParcial) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTrazaDeBulto()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *BultoRoturaParcial) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *BultoRoturaParcial) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ BultoRoturaParcial) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ BultoRoturaParcial) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ BultoRoturaParcial) HintSize(int)                     { panic("Unsupported operation") }
func (_ BultoRoturaParcial) Finalize()                        {}

func (_ BultoRoturaParcial) AvroCRC64Fingerprint() []byte {
	return []byte(BultoRoturaParcialAvroCRC64Fingerprint)
}

func (r BultoRoturaParcial) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *BultoRoturaParcial) UnmarshalJSON(data []byte) error {
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
