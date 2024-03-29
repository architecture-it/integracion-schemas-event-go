// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhPedidosEmpaquetado.avsc
 */
package EventoWhPedidosEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Bultos struct {
	DropId string `json:"DropId"`

	Altura float32 `json:"Altura"`

	Profundidad float32 `json:"Profundidad"`

	Ancho float32 `json:"Ancho"`

	Peso float32 `json:"Peso"`

	Contenedor *UnionNullContenedor `json:"Contenedor"`
}

const BultosAvroCRC64Fingerprint = "\xa3K\xe7y\xea\t\xc74"

func NewBultos() Bultos {
	r := Bultos{}
	r.Contenedor = nil
	return r
}

func DeserializeBultos(r io.Reader) (Bultos, error) {
	t := NewBultos()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBultosFromSchema(r io.Reader, schema string) (Bultos, error) {
	t := NewBultos()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBultos(r Bultos, w io.Writer) error {
	var err error
	err = vm.WriteString(r.DropId, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.Altura, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.Profundidad, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.Ancho, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.Peso, w)
	if err != nil {
		return err
	}
	err = writeUnionNullContenedor(r.Contenedor, w)
	if err != nil {
		return err
	}
	return err
}

func (r Bultos) Serialize(w io.Writer) error {
	return writeBultos(r, w)
}

func (r Bultos) Schema() string {
	return "{\"fields\":[{\"name\":\"DropId\",\"type\":\"string\"},{\"name\":\"Altura\",\"type\":\"float\"},{\"name\":\"Profundidad\",\"type\":\"float\"},{\"name\":\"Ancho\",\"type\":\"float\"},{\"name\":\"Peso\",\"type\":\"float\"},{\"default\":null,\"name\":\"Contenedor\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Tara\",\"type\":\"float\"},{\"name\":\"CartonType\",\"type\":\"string\"},{\"name\":\"ContratoRetornable\",\"type\":[\"null\",\"string\"]}],\"name\":\"Contenedor\",\"type\":\"record\"}]}],\"name\":\"Andreani.EventoWhPedidos.Events.EmpaquetadoCommon.Bultos\",\"type\":\"record\"}"
}

func (r Bultos) SchemaName() string {
	return "Andreani.EventoWhPedidos.Events.EmpaquetadoCommon.Bultos"
}

func (_ Bultos) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Bultos) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Bultos) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Bultos) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Bultos) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Bultos) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Bultos) SetString(v string)   { panic("Unsupported operation") }
func (_ Bultos) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Bultos) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.DropId}

		return w

	case 1:
		w := types.Float{Target: &r.Altura}

		return w

	case 2:
		w := types.Float{Target: &r.Profundidad}

		return w

	case 3:
		w := types.Float{Target: &r.Ancho}

		return w

	case 4:
		w := types.Float{Target: &r.Peso}

		return w

	case 5:
		r.Contenedor = NewUnionNullContenedor()

		return r.Contenedor
	}
	panic("Unknown field index")
}

func (r *Bultos) SetDefault(i int) {
	switch i {
	case 5:
		r.Contenedor = nil
		return
	}
	panic("Unknown field index")
}

func (r *Bultos) NullField(i int) {
	switch i {
	case 5:
		r.Contenedor = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Bultos) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Bultos) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Bultos) HintSize(int)                     { panic("Unsupported operation") }
func (_ Bultos) Finalize()                        {}

func (_ Bultos) AvroCRC64Fingerprint() []byte {
	return []byte(BultosAvroCRC64Fingerprint)
}

func (r Bultos) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["DropId"], err = json.Marshal(r.DropId)
	if err != nil {
		return nil, err
	}
	output["Altura"], err = json.Marshal(r.Altura)
	if err != nil {
		return nil, err
	}
	output["Profundidad"], err = json.Marshal(r.Profundidad)
	if err != nil {
		return nil, err
	}
	output["Ancho"], err = json.Marshal(r.Ancho)
	if err != nil {
		return nil, err
	}
	output["Peso"], err = json.Marshal(r.Peso)
	if err != nil {
		return nil, err
	}
	output["Contenedor"], err = json.Marshal(r.Contenedor)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Bultos) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["DropId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DropId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DropId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Altura"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Altura); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Altura")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Profundidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Profundidad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Profundidad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Ancho"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ancho); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Ancho")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Peso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Peso); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Peso")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Contenedor"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contenedor); err != nil {
			return err
		}
	} else {
		r.Contenedor = NewUnionNullContenedor()

		r.Contenedor = nil
	}
	return nil
}
