// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     BultoExpedicionEnReparto.avsc
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

type BultoExpedicionEnReparto struct {
	Traza TrazaDeBulto `json:"traza"`
}

const BultoExpedicionEnRepartoAvroCRC64Fingerprint = "\xac\xdc\xfa\xa4d\x93\xf9\x17"

func NewBultoExpedicionEnReparto() BultoExpedicionEnReparto {
	r := BultoExpedicionEnReparto{}
	r.Traza = NewTrazaDeBulto()

	return r
}

func DeserializeBultoExpedicionEnReparto(r io.Reader) (BultoExpedicionEnReparto, error) {
	t := NewBultoExpedicionEnReparto()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBultoExpedicionEnRepartoFromSchema(r io.Reader, schema string) (BultoExpedicionEnReparto, error) {
	t := NewBultoExpedicionEnReparto()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBultoExpedicionEnReparto(r BultoExpedicionEnReparto, w io.Writer) error {
	var err error
	err = writeTrazaDeBulto(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r BultoExpedicionEnReparto) Serialize(w io.Writer) error {
	return writeBultoExpedicionEnReparto(r, w)
}

func (r BultoExpedicionEnReparto) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"name\":\"numeroDeBulto\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"tipoDeBulto\",\"type\":[\"null\",{\"name\":\"TipoDeBulto\",\"symbols\":[\"Paquete\",\"Contenedor\"],\"type\":\"enum\"}]}],\"name\":\"TrazaDeBulto\",\"namespace\":\"Integracion.Esquemas.Bulto.Referencias\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Bulto.Trazas.BultoExpedicionEnReparto\",\"type\":\"record\"}"
}

func (r BultoExpedicionEnReparto) SchemaName() string {
	return "Integracion.Esquemas.Bulto.Trazas.BultoExpedicionEnReparto"
}

func (_ BultoExpedicionEnReparto) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BultoExpedicionEnReparto) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BultoExpedicionEnReparto) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BultoExpedicionEnReparto) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BultoExpedicionEnReparto) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BultoExpedicionEnReparto) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BultoExpedicionEnReparto) SetString(v string)   { panic("Unsupported operation") }
func (_ BultoExpedicionEnReparto) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BultoExpedicionEnReparto) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTrazaDeBulto()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *BultoExpedicionEnReparto) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *BultoExpedicionEnReparto) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ BultoExpedicionEnReparto) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ BultoExpedicionEnReparto) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ BultoExpedicionEnReparto) HintSize(int)                     { panic("Unsupported operation") }
func (_ BultoExpedicionEnReparto) Finalize()                        {}

func (_ BultoExpedicionEnReparto) AvroCRC64Fingerprint() []byte {
	return []byte(BultoExpedicionEnRepartoAvroCRC64Fingerprint)
}

func (r BultoExpedicionEnReparto) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *BultoExpedicionEnReparto) UnmarshalJSON(data []byte) error {
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
