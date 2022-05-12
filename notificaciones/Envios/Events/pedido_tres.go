// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     Pedido.avsc
 *     PedidoDos.avsc
 *     PedidoTres.avsc
 */
package Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type PedidoTres struct {
	Id string `json:"id"`

	NumeroDePedido int32 `json:"numeroDePedido"`

	Planta string `json:"planta"`

	Almacen string `json:"almacen"`

	Contrato string `json:"contrato"`
}

const PedidoTresAvroCRC64Fingerprint = "l\xaf\"c\xabg\x10\u007f"

func NewPedidoTres() PedidoTres {
	r := PedidoTres{}
	return r
}

func DeserializePedidoTres(r io.Reader) (PedidoTres, error) {
	t := NewPedidoTres()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePedidoTresFromSchema(r io.Reader, schema string) (PedidoTres, error) {
	t := NewPedidoTres()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePedidoTres(r PedidoTres, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.NumeroDePedido, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Planta, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Almacen, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Contrato, w)
	if err != nil {
		return err
	}
	return err
}

func (r PedidoTres) Serialize(w io.Writer) error {
	return writePedidoTres(r, w)
}

func (r PedidoTres) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"numeroDePedido\",\"type\":\"int\"},{\"name\":\"planta\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"name\":\"contrato\",\"type\":\"string\"}],\"name\":\"Andreani.pipeline.Event.PedidoTres\",\"type\":\"record\"}"
}

func (r PedidoTres) SchemaName() string {
	return "Andreani.pipeline.Event.PedidoTres"
}

func (_ PedidoTres) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ PedidoTres) SetInt(v int32)       { panic("Unsupported operation") }
func (_ PedidoTres) SetLong(v int64)      { panic("Unsupported operation") }
func (_ PedidoTres) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ PedidoTres) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ PedidoTres) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ PedidoTres) SetString(v string)   { panic("Unsupported operation") }
func (_ PedidoTres) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *PedidoTres) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Id}

		return w

	case 1:
		w := types.Int{Target: &r.NumeroDePedido}

		return w

	case 2:
		w := types.String{Target: &r.Planta}

		return w

	case 3:
		w := types.String{Target: &r.Almacen}

		return w

	case 4:
		w := types.String{Target: &r.Contrato}

		return w

	}
	panic("Unknown field index")
}

func (r *PedidoTres) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *PedidoTres) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ PedidoTres) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ PedidoTres) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ PedidoTres) HintSize(int)                     { panic("Unsupported operation") }
func (_ PedidoTres) Finalize()                        {}

func (_ PedidoTres) AvroCRC64Fingerprint() []byte {
	return []byte(PedidoTresAvroCRC64Fingerprint)
}

func (r PedidoTres) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["numeroDePedido"], err = json.Marshal(r.NumeroDePedido)
	if err != nil {
		return nil, err
	}
	output["planta"], err = json.Marshal(r.Planta)
	if err != nil {
		return nil, err
	}
	output["almacen"], err = json.Marshal(r.Almacen)
	if err != nil {
		return nil, err
	}
	output["contrato"], err = json.Marshal(r.Contrato)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *PedidoTres) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroDePedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDePedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroDePedido")
	}
	val = func() json.RawMessage {
		if v, ok := fields["planta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Planta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for planta")
	}
	val = func() json.RawMessage {
		if v, ok := fields["almacen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Almacen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for almacen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["contrato"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contrato); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for contrato")
	}
	return nil
}
