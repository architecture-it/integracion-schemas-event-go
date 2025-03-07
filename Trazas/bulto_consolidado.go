// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     BultoConsolidado.avsc
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

type BultoConsolidado struct {
	Traza TrazaDeBulto `json:"traza"`

	CodigoDeContenedor string `json:"codigoDeContenedor"`
}

const BultoConsolidadoAvroCRC64Fingerprint = "\xe1&\xcb(\xa9\x92I\a"

func NewBultoConsolidado() BultoConsolidado {
	r := BultoConsolidado{}
	r.Traza = NewTrazaDeBulto()

	return r
}

func DeserializeBultoConsolidado(r io.Reader) (BultoConsolidado, error) {
	t := NewBultoConsolidado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBultoConsolidadoFromSchema(r io.Reader, schema string) (BultoConsolidado, error) {
	t := NewBultoConsolidado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBultoConsolidado(r BultoConsolidado, w io.Writer) error {
	var err error
	err = writeTrazaDeBulto(r.Traza, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoDeContenedor, w)
	if err != nil {
		return err
	}
	return err
}

func (r BultoConsolidado) Serialize(w io.Writer) error {
	return writeBultoConsolidado(r, w)
}

func (r BultoConsolidado) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"name\":\"numeroDeBulto\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"tipoDeBulto\",\"type\":[\"null\",{\"name\":\"TipoDeBulto\",\"symbols\":[\"Paquete\",\"Contenedor\"],\"type\":\"enum\"}]}],\"name\":\"TrazaDeBulto\",\"namespace\":\"Integracion.Esquemas.Bulto.Referencias\",\"type\":\"record\"}},{\"name\":\"codigoDeContenedor\",\"type\":\"string\"}],\"name\":\"Integracion.Esquemas.Bulto.Trazas.BultoConsolidado\",\"type\":\"record\"}"
}

func (r BultoConsolidado) SchemaName() string {
	return "Integracion.Esquemas.Bulto.Trazas.BultoConsolidado"
}

func (_ BultoConsolidado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BultoConsolidado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BultoConsolidado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BultoConsolidado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BultoConsolidado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BultoConsolidado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BultoConsolidado) SetString(v string)   { panic("Unsupported operation") }
func (_ BultoConsolidado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BultoConsolidado) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTrazaDeBulto()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		w := types.String{Target: &r.CodigoDeContenedor}

		return w

	}
	panic("Unknown field index")
}

func (r *BultoConsolidado) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *BultoConsolidado) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ BultoConsolidado) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ BultoConsolidado) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ BultoConsolidado) HintSize(int)                     { panic("Unsupported operation") }
func (_ BultoConsolidado) Finalize()                        {}

func (_ BultoConsolidado) AvroCRC64Fingerprint() []byte {
	return []byte(BultoConsolidadoAvroCRC64Fingerprint)
}

func (r BultoConsolidado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["codigoDeContenedor"], err = json.Marshal(r.CodigoDeContenedor)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *BultoConsolidado) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["codigoDeContenedor"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeContenedor); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigoDeContenedor")
	}
	return nil
}
