// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     WosTrazaAnmat.avsc
 */
package WosTrazabilidadEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Serie struct {
	Serie string `json:"serie"`

	DescripcionEstado string `json:"descripcionEstado"`
}

const SerieAvroCRC64Fingerprint = "\xbbRəG\xb0\xfd("

func NewSerie() Serie {
	r := Serie{}
	return r
}

func DeserializeSerie(r io.Reader) (Serie, error) {
	t := NewSerie()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSerieFromSchema(r io.Reader, schema string) (Serie, error) {
	t := NewSerie()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSerie(r Serie, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Serie, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.DescripcionEstado, w)
	if err != nil {
		return err
	}
	return err
}

func (r Serie) Serialize(w io.Writer) error {
	return writeSerie(r, w)
}

func (r Serie) Schema() string {
	return "{\"fields\":[{\"name\":\"serie\",\"type\":\"string\"},{\"name\":\"descripcionEstado\",\"type\":\"string\"}],\"name\":\"Andreani.WosTrazabilidad.Events.AnmatCommon.Serie\",\"type\":\"record\"}"
}

func (r Serie) SchemaName() string {
	return "Andreani.WosTrazabilidad.Events.AnmatCommon.Serie"
}

func (_ Serie) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Serie) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Serie) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Serie) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Serie) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Serie) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Serie) SetString(v string)   { panic("Unsupported operation") }
func (_ Serie) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Serie) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Serie}

		return w

	case 1:
		w := types.String{Target: &r.DescripcionEstado}

		return w

	}
	panic("Unknown field index")
}

func (r *Serie) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Serie) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Serie) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Serie) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Serie) HintSize(int)                     { panic("Unsupported operation") }
func (_ Serie) Finalize()                        {}

func (_ Serie) AvroCRC64Fingerprint() []byte {
	return []byte(SerieAvroCRC64Fingerprint)
}

func (r Serie) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["serie"], err = json.Marshal(r.Serie)
	if err != nil {
		return nil, err
	}
	output["descripcionEstado"], err = json.Marshal(r.DescripcionEstado)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Serie) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["serie"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Serie); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for serie")
	}
	val = func() json.RawMessage {
		if v, ok := fields["descripcionEstado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DescripcionEstado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for descripcionEstado")
	}
	return nil
}
