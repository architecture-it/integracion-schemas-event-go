// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     BultoContingenciaEnSistemas.avsc
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

type BultoContingenciaEnSistemas struct {
	Traza TrazaDeBulto `json:"traza"`
}

const BultoContingenciaEnSistemasAvroCRC64Fingerprint = "\x8e\xeb\xdcЍm\xd6\xe4"

func NewBultoContingenciaEnSistemas() BultoContingenciaEnSistemas {
	r := BultoContingenciaEnSistemas{}
	r.Traza = NewTrazaDeBulto()

	return r
}

func DeserializeBultoContingenciaEnSistemas(r io.Reader) (BultoContingenciaEnSistemas, error) {
	t := NewBultoContingenciaEnSistemas()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBultoContingenciaEnSistemasFromSchema(r io.Reader, schema string) (BultoContingenciaEnSistemas, error) {
	t := NewBultoContingenciaEnSistemas()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBultoContingenciaEnSistemas(r BultoContingenciaEnSistemas, w io.Writer) error {
	var err error
	err = writeTrazaDeBulto(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r BultoContingenciaEnSistemas) Serialize(w io.Writer) error {
	return writeBultoContingenciaEnSistemas(r, w)
}

func (r BultoContingenciaEnSistemas) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"name\":\"numeroDeBulto\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"tipoDeBulto\",\"type\":[\"null\",{\"name\":\"TipoDeBulto\",\"symbols\":[\"Paquete\",\"Contenedor\"],\"type\":\"enum\"}]}],\"name\":\"TrazaDeBulto\",\"namespace\":\"Integracion.Esquemas.Bulto.Referencias\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Bulto.Trazas.BultoContingenciaEnSistemas\",\"type\":\"record\"}"
}

func (r BultoContingenciaEnSistemas) SchemaName() string {
	return "Integracion.Esquemas.Bulto.Trazas.BultoContingenciaEnSistemas"
}

func (_ BultoContingenciaEnSistemas) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BultoContingenciaEnSistemas) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BultoContingenciaEnSistemas) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BultoContingenciaEnSistemas) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BultoContingenciaEnSistemas) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BultoContingenciaEnSistemas) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BultoContingenciaEnSistemas) SetString(v string)   { panic("Unsupported operation") }
func (_ BultoContingenciaEnSistemas) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BultoContingenciaEnSistemas) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTrazaDeBulto()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *BultoContingenciaEnSistemas) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *BultoContingenciaEnSistemas) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ BultoContingenciaEnSistemas) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ BultoContingenciaEnSistemas) AppendArray() types.Field { panic("Unsupported operation") }
func (_ BultoContingenciaEnSistemas) HintSize(int)             { panic("Unsupported operation") }
func (_ BultoContingenciaEnSistemas) Finalize()                {}

func (_ BultoContingenciaEnSistemas) AvroCRC64Fingerprint() []byte {
	return []byte(BultoContingenciaEnSistemasAvroCRC64Fingerprint)
}

func (r BultoContingenciaEnSistemas) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *BultoContingenciaEnSistemas) UnmarshalJSON(data []byte) error {
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
