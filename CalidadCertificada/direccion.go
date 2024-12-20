// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Direccion.avsc
 */
package CalidadCertificadaEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Direccion struct {
	Postal Postal `json:"Postal"`
}

const DireccionAvroCRC64Fingerprint = "\x1f\xae\a\r~\xb7\xba\x19"

func NewDireccion() Direccion {
	r := Direccion{}
	r.Postal = NewPostal()

	return r
}

func DeserializeDireccion(r io.Reader) (Direccion, error) {
	t := NewDireccion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDireccionFromSchema(r io.Reader, schema string) (Direccion, error) {
	t := NewDireccion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDireccion(r Direccion, w io.Writer) error {
	var err error
	err = writePostal(r.Postal, w)
	if err != nil {
		return err
	}
	return err
}

func (r Direccion) Serialize(w io.Writer) error {
	return writeDireccion(r, w)
}

func (r Direccion) Schema() string {
	return "{\"fields\":[{\"name\":\"Postal\",\"type\":{\"fields\":[{\"name\":\"CodigoPostal\",\"type\":\"string\"},{\"name\":\"Calle\",\"type\":\"string\"},{\"name\":\"Numero\",\"type\":\"string\"},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"Region\",\"type\":\"string\"},{\"name\":\"Pais\",\"type\":\"string\"},{\"name\":\"ComponentesDeDireccion\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Meta\",\"type\":\"string\"},{\"name\":\"Contenido\",\"type\":\"string\"}],\"name\":\"ComponenteDeDireccion\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Postal\",\"type\":\"record\"}}],\"name\":\"Andreani.CalidadCertificada.Events.Record.Direccion\",\"type\":\"record\"}"
}

func (r Direccion) SchemaName() string {
	return "Andreani.CalidadCertificada.Events.Record.Direccion"
}

func (_ Direccion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Direccion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Direccion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Direccion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Direccion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Direccion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Direccion) SetString(v string)   { panic("Unsupported operation") }
func (_ Direccion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Direccion) Get(i int) types.Field {
	switch i {
	case 0:
		r.Postal = NewPostal()

		w := types.Record{Target: &r.Postal}

		return w

	}
	panic("Unknown field index")
}

func (r *Direccion) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Direccion) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Direccion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Direccion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Direccion) HintSize(int)                     { panic("Unsupported operation") }
func (_ Direccion) Finalize()                        {}

func (_ Direccion) AvroCRC64Fingerprint() []byte {
	return []byte(DireccionAvroCRC64Fingerprint)
}

func (r Direccion) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Postal"], err = json.Marshal(r.Postal)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Direccion) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Postal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Postal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Postal")
	}
	return nil
}
