// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     BultoEliminarMarcaDeDestruccion.avsc
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

type BultoEliminarMarcaDeDestruccion struct {
	Traza TrazaDeBulto `json:"traza"`
}

const BultoEliminarMarcaDeDestruccionAvroCRC64Fingerprint = "{@\xdc;\xe8M\xa6\xf1"

func NewBultoEliminarMarcaDeDestruccion() BultoEliminarMarcaDeDestruccion {
	r := BultoEliminarMarcaDeDestruccion{}
	r.Traza = NewTrazaDeBulto()

	return r
}

func DeserializeBultoEliminarMarcaDeDestruccion(r io.Reader) (BultoEliminarMarcaDeDestruccion, error) {
	t := NewBultoEliminarMarcaDeDestruccion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBultoEliminarMarcaDeDestruccionFromSchema(r io.Reader, schema string) (BultoEliminarMarcaDeDestruccion, error) {
	t := NewBultoEliminarMarcaDeDestruccion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBultoEliminarMarcaDeDestruccion(r BultoEliminarMarcaDeDestruccion, w io.Writer) error {
	var err error
	err = writeTrazaDeBulto(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r BultoEliminarMarcaDeDestruccion) Serialize(w io.Writer) error {
	return writeBultoEliminarMarcaDeDestruccion(r, w)
}

func (r BultoEliminarMarcaDeDestruccion) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"name\":\"numeroDeBulto\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"tipoDeBulto\",\"type\":[\"null\",{\"name\":\"TipoDeBulto\",\"symbols\":[\"Paquete\",\"Contenedor\"],\"type\":\"enum\"}]}],\"name\":\"TrazaDeBulto\",\"namespace\":\"Integracion.Esquemas.Bulto.Referencias\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Bulto.Trazas.BultoEliminarMarcaDeDestruccion\",\"type\":\"record\"}"
}

func (r BultoEliminarMarcaDeDestruccion) SchemaName() string {
	return "Integracion.Esquemas.Bulto.Trazas.BultoEliminarMarcaDeDestruccion"
}

func (_ BultoEliminarMarcaDeDestruccion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BultoEliminarMarcaDeDestruccion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BultoEliminarMarcaDeDestruccion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BultoEliminarMarcaDeDestruccion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BultoEliminarMarcaDeDestruccion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BultoEliminarMarcaDeDestruccion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BultoEliminarMarcaDeDestruccion) SetString(v string)   { panic("Unsupported operation") }
func (_ BultoEliminarMarcaDeDestruccion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BultoEliminarMarcaDeDestruccion) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTrazaDeBulto()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *BultoEliminarMarcaDeDestruccion) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *BultoEliminarMarcaDeDestruccion) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ BultoEliminarMarcaDeDestruccion) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ BultoEliminarMarcaDeDestruccion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ BultoEliminarMarcaDeDestruccion) HintSize(int)             { panic("Unsupported operation") }
func (_ BultoEliminarMarcaDeDestruccion) Finalize()                {}

func (_ BultoEliminarMarcaDeDestruccion) AvroCRC64Fingerprint() []byte {
	return []byte(BultoEliminarMarcaDeDestruccionAvroCRC64Fingerprint)
}

func (r BultoEliminarMarcaDeDestruccion) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *BultoEliminarMarcaDeDestruccion) UnmarshalJSON(data []byte) error {
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
