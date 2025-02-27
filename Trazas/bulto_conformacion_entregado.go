// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     BultoConformacionEntregado.avsc
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

type BultoConformacionEntregado struct {
	Traza TrazaDeBulto `json:"traza"`
}

const BultoConformacionEntregadoAvroCRC64Fingerprint = "\xfd\xc1\xaa6\xcc\xe7\xb2\xdc"

func NewBultoConformacionEntregado() BultoConformacionEntregado {
	r := BultoConformacionEntregado{}
	r.Traza = NewTrazaDeBulto()

	return r
}

func DeserializeBultoConformacionEntregado(r io.Reader) (BultoConformacionEntregado, error) {
	t := NewBultoConformacionEntregado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBultoConformacionEntregadoFromSchema(r io.Reader, schema string) (BultoConformacionEntregado, error) {
	t := NewBultoConformacionEntregado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBultoConformacionEntregado(r BultoConformacionEntregado, w io.Writer) error {
	var err error
	err = writeTrazaDeBulto(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r BultoConformacionEntregado) Serialize(w io.Writer) error {
	return writeBultoConformacionEntregado(r, w)
}

func (r BultoConformacionEntregado) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"name\":\"numeroDeBulto\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"tipoDeBulto\",\"type\":[\"null\",{\"name\":\"TipoDeBulto\",\"symbols\":[\"Paquete\",\"Contenedor\"],\"type\":\"enum\"}]}],\"name\":\"TrazaDeBulto\",\"namespace\":\"Integracion.Esquemas.Bulto.Referencias\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Bulto.Trazas.BultoConformacionEntregado\",\"type\":\"record\"}"
}

func (r BultoConformacionEntregado) SchemaName() string {
	return "Integracion.Esquemas.Bulto.Trazas.BultoConformacionEntregado"
}

func (_ BultoConformacionEntregado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BultoConformacionEntregado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BultoConformacionEntregado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BultoConformacionEntregado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BultoConformacionEntregado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BultoConformacionEntregado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BultoConformacionEntregado) SetString(v string)   { panic("Unsupported operation") }
func (_ BultoConformacionEntregado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BultoConformacionEntregado) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTrazaDeBulto()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *BultoConformacionEntregado) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *BultoConformacionEntregado) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ BultoConformacionEntregado) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ BultoConformacionEntregado) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ BultoConformacionEntregado) HintSize(int)                     { panic("Unsupported operation") }
func (_ BultoConformacionEntregado) Finalize()                        {}

func (_ BultoConformacionEntregado) AvroCRC64Fingerprint() []byte {
	return []byte(BultoConformacionEntregadoAvroCRC64Fingerprint)
}

func (r BultoConformacionEntregado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *BultoConformacionEntregado) UnmarshalJSON(data []byte) error {
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
