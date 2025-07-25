// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoDeAlmacenSCE.avsc
 */
package WarehousePedidoEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ImpRemito struct {
	Value string `json:"value"`
}

const ImpRemitoAvroCRC64Fingerprint = "\x14\x1c\xf1\x8f`\n4R"

func NewImpRemito() ImpRemito {
	r := ImpRemito{}
	return r
}

func DeserializeImpRemito(r io.Reader) (ImpRemito, error) {
	t := NewImpRemito()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeImpRemitoFromSchema(r io.Reader, schema string) (ImpRemito, error) {
	t := NewImpRemito()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeImpRemito(r ImpRemito, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Value, w)
	if err != nil {
		return err
	}
	return err
}

func (r ImpRemito) Serialize(w io.Writer) error {
	return writeImpRemito(r, w)
}

func (r ImpRemito) Schema() string {
	return "{\"fields\":[{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"Andreani.WarehousePedido.Events.Record.ImpRemito\",\"type\":\"record\"}"
}

func (r ImpRemito) SchemaName() string {
	return "Andreani.WarehousePedido.Events.Record.ImpRemito"
}

func (_ ImpRemito) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ImpRemito) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ImpRemito) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ImpRemito) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ImpRemito) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ImpRemito) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ImpRemito) SetString(v string)   { panic("Unsupported operation") }
func (_ ImpRemito) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ImpRemito) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Value}

		return w

	}
	panic("Unknown field index")
}

func (r *ImpRemito) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ImpRemito) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ImpRemito) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ImpRemito) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ImpRemito) HintSize(int)                     { panic("Unsupported operation") }
func (_ ImpRemito) Finalize()                        {}

func (_ ImpRemito) AvroCRC64Fingerprint() []byte {
	return []byte(ImpRemitoAvroCRC64Fingerprint)
}

func (r ImpRemito) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["value"], err = json.Marshal(r.Value)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ImpRemito) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["value"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Value); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for value")
	}
	return nil
}
