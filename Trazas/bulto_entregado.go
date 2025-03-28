// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     BultoEntregado.avsc
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

type BultoEntregado struct {
	Traza TrazaDeBulto `json:"traza"`

	TipoDeEntrega TipoDeEntrega `json:"tipoDeEntrega"`
}

const BultoEntregadoAvroCRC64Fingerprint = "\x10\xb2\xab\x15\xf7[\xe6u"

func NewBultoEntregado() BultoEntregado {
	r := BultoEntregado{}
	r.Traza = NewTrazaDeBulto()

	return r
}

func DeserializeBultoEntregado(r io.Reader) (BultoEntregado, error) {
	t := NewBultoEntregado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBultoEntregadoFromSchema(r io.Reader, schema string) (BultoEntregado, error) {
	t := NewBultoEntregado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBultoEntregado(r BultoEntregado, w io.Writer) error {
	var err error
	err = writeTrazaDeBulto(r.Traza, w)
	if err != nil {
		return err
	}
	err = writeTipoDeEntrega(r.TipoDeEntrega, w)
	if err != nil {
		return err
	}
	return err
}

func (r BultoEntregado) Serialize(w io.Writer) error {
	return writeBultoEntregado(r, w)
}

func (r BultoEntregado) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"name\":\"numeroDeBulto\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"tipoDeBulto\",\"type\":[\"null\",{\"name\":\"TipoDeBulto\",\"symbols\":[\"Paquete\",\"Contenedor\"],\"type\":\"enum\"}]}],\"name\":\"TrazaDeBulto\",\"namespace\":\"Integracion.Esquemas.Bulto.Referencias\",\"type\":\"record\"}},{\"name\":\"tipoDeEntrega\",\"type\":{\"name\":\"TipoDeEntrega\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"symbols\":[\"undefined\",\"distribucion\",\"deliveryWindow\"],\"type\":\"enum\"}}],\"name\":\"Integracion.Esquemas.Bulto.Trazas.BultoEntregado\",\"type\":\"record\"}"
}

func (r BultoEntregado) SchemaName() string {
	return "Integracion.Esquemas.Bulto.Trazas.BultoEntregado"
}

func (_ BultoEntregado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BultoEntregado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BultoEntregado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BultoEntregado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BultoEntregado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BultoEntregado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BultoEntregado) SetString(v string)   { panic("Unsupported operation") }
func (_ BultoEntregado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BultoEntregado) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTrazaDeBulto()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		w := TipoDeEntregaWrapper{Target: &r.TipoDeEntrega}

		return w

	}
	panic("Unknown field index")
}

func (r *BultoEntregado) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *BultoEntregado) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ BultoEntregado) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ BultoEntregado) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ BultoEntregado) HintSize(int)                     { panic("Unsupported operation") }
func (_ BultoEntregado) Finalize()                        {}

func (_ BultoEntregado) AvroCRC64Fingerprint() []byte {
	return []byte(BultoEntregadoAvroCRC64Fingerprint)
}

func (r BultoEntregado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["tipoDeEntrega"], err = json.Marshal(r.TipoDeEntrega)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *BultoEntregado) UnmarshalJSON(data []byte) error {
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
	val = func() json.RawMessage {
		if v, ok := fields["tipoDeEntrega"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeEntrega); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tipoDeEntrega")
	}
	return nil
}
