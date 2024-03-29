// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     OrdenDeEnvioRechazada.avsc
 */
package SppeApiEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Bulto struct {
	DropId string `json:"DropId"`

	Altura float32 `json:"Altura"`

	Profundidad float32 `json:"Profundidad"`

	Ancho float32 `json:"Ancho"`

	Peso float32 `json:"Peso"`
}

const BultoAvroCRC64Fingerprint = "\xb9w;\xe8\x17\xfc\x00\xb3"

func NewBulto() Bulto {
	r := Bulto{}
	return r
}

func DeserializeBulto(r io.Reader) (Bulto, error) {
	t := NewBulto()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBultoFromSchema(r io.Reader, schema string) (Bulto, error) {
	t := NewBulto()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBulto(r Bulto, w io.Writer) error {
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
	return err
}

func (r Bulto) Serialize(w io.Writer) error {
	return writeBulto(r, w)
}

func (r Bulto) Schema() string {
	return "{\"fields\":[{\"name\":\"DropId\",\"type\":\"string\"},{\"name\":\"Altura\",\"type\":\"float\"},{\"name\":\"Profundidad\",\"type\":\"float\"},{\"name\":\"Ancho\",\"type\":\"float\"},{\"name\":\"Peso\",\"type\":\"float\"}],\"name\":\"Andreani.SppeApi.Events.OrdenDeEnvioRechazadaCommon.Bulto\",\"type\":\"record\"}"
}

func (r Bulto) SchemaName() string {
	return "Andreani.SppeApi.Events.OrdenDeEnvioRechazadaCommon.Bulto"
}

func (_ Bulto) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Bulto) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Bulto) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Bulto) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Bulto) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Bulto) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Bulto) SetString(v string)   { panic("Unsupported operation") }
func (_ Bulto) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Bulto) Get(i int) types.Field {
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

	}
	panic("Unknown field index")
}

func (r *Bulto) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Bulto) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Bulto) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Bulto) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Bulto) HintSize(int)                     { panic("Unsupported operation") }
func (_ Bulto) Finalize()                        {}

func (_ Bulto) AvroCRC64Fingerprint() []byte {
	return []byte(BultoAvroCRC64Fingerprint)
}

func (r Bulto) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(output)
}

func (r *Bulto) UnmarshalJSON(data []byte) error {
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
	return nil
}
