// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Envio.avsc
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

type Envio struct {
	CodigoDeEnvio string `json:"codigoDeEnvio"`

	CantidadDeBultos int32 `json:"cantidadDeBultos"`

	CantidadDeBultosLeidos int32 `json:"cantidadDeBultosLeidos"`

	Totalizador string `json:"totalizador"`

	Bultos *UnionNullArrayBulto `json:"bultos"`
}

const EnvioAvroCRC64Fingerprint = "\x969\"\x13\x00e\xe5a"

func NewEnvio() Envio {
	r := Envio{}
	r.Bultos = nil
	return r
}

func DeserializeEnvio(r io.Reader) (Envio, error) {
	t := NewEnvio()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnvioFromSchema(r io.Reader, schema string) (Envio, error) {
	t := NewEnvio()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnvio(r Envio, w io.Writer) error {
	var err error
	err = vm.WriteString(r.CodigoDeEnvio, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.CantidadDeBultos, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.CantidadDeBultosLeidos, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Totalizador, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayBulto(r.Bultos, w)
	if err != nil {
		return err
	}
	return err
}

func (r Envio) Serialize(w io.Writer) error {
	return writeEnvio(r, w)
}

func (r Envio) Schema() string {
	return "{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"name\":\"cantidadDeBultos\",\"type\":\"int\"},{\"name\":\"cantidadDeBultosLeidos\",\"type\":\"int\"},{\"name\":\"totalizador\",\"type\":\"string\"},{\"default\":null,\"name\":\"bultos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"numeroDeBulto\",\"type\":\"string\"},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}}],\"name\":\"Bulto\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"Integracion.Esquemas.Contenedor.Referencias.Envio\",\"type\":\"record\"}"
}

func (r Envio) SchemaName() string {
	return "Integracion.Esquemas.Contenedor.Referencias.Envio"
}

func (_ Envio) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Envio) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Envio) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Envio) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Envio) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Envio) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Envio) SetString(v string)   { panic("Unsupported operation") }
func (_ Envio) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Envio) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.CodigoDeEnvio}

		return w

	case 1:
		w := types.Int{Target: &r.CantidadDeBultos}

		return w

	case 2:
		w := types.Int{Target: &r.CantidadDeBultosLeidos}

		return w

	case 3:
		w := types.String{Target: &r.Totalizador}

		return w

	case 4:
		r.Bultos = NewUnionNullArrayBulto()

		return r.Bultos
	}
	panic("Unknown field index")
}

func (r *Envio) SetDefault(i int) {
	switch i {
	case 4:
		r.Bultos = nil
		return
	}
	panic("Unknown field index")
}

func (r *Envio) NullField(i int) {
	switch i {
	case 4:
		r.Bultos = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Envio) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Envio) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Envio) HintSize(int)                     { panic("Unsupported operation") }
func (_ Envio) Finalize()                        {}

func (_ Envio) AvroCRC64Fingerprint() []byte {
	return []byte(EnvioAvroCRC64Fingerprint)
}

func (r Envio) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["codigoDeEnvio"], err = json.Marshal(r.CodigoDeEnvio)
	if err != nil {
		return nil, err
	}
	output["cantidadDeBultos"], err = json.Marshal(r.CantidadDeBultos)
	if err != nil {
		return nil, err
	}
	output["cantidadDeBultosLeidos"], err = json.Marshal(r.CantidadDeBultosLeidos)
	if err != nil {
		return nil, err
	}
	output["totalizador"], err = json.Marshal(r.Totalizador)
	if err != nil {
		return nil, err
	}
	output["bultos"], err = json.Marshal(r.Bultos)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Envio) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["codigoDeEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeEnvio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigoDeEnvio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["cantidadDeBultos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadDeBultos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cantidadDeBultos")
	}
	val = func() json.RawMessage {
		if v, ok := fields["cantidadDeBultosLeidos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadDeBultosLeidos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cantidadDeBultosLeidos")
	}
	val = func() json.RawMessage {
		if v, ok := fields["totalizador"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Totalizador); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for totalizador")
	}
	val = func() json.RawMessage {
		if v, ok := fields["bultos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Bultos); err != nil {
			return err
		}
	} else {
		r.Bultos = NewUnionNullArrayBulto()

		r.Bultos = nil
	}
	return nil
}
