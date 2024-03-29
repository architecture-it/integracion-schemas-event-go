// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoValidado.avsc
 */
package CtrlExpEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type PedidoValidado struct {
	NroPedido string `json:"NroPedido"`

	NroPedidoExterno string `json:"NroPedidoExterno"`

	SubPedido string `json:"SubPedido"`

	Bandeja string `json:"Bandeja"`

	Estado string `json:"Estado"`

	IpImpresora string `json:"IpImpresora"`

	Cliente string `json:"Cliente"`

	SchemaDb *UnionNullString `json:"SchemaDb"`
}

const PedidoValidadoAvroCRC64Fingerprint = "\x10\x93\tT\x00\xff\xc0\x8a"

func NewPedidoValidado() PedidoValidado {
	r := PedidoValidado{}
	r.SchemaDb = nil
	return r
}

func DeserializePedidoValidado(r io.Reader) (PedidoValidado, error) {
	t := NewPedidoValidado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePedidoValidadoFromSchema(r io.Reader, schema string) (PedidoValidado, error) {
	t := NewPedidoValidado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePedidoValidado(r PedidoValidado, w io.Writer) error {
	var err error
	err = vm.WriteString(r.NroPedido, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NroPedidoExterno, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.SubPedido, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Bandeja, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Estado, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.IpImpresora, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Cliente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.SchemaDb, w)
	if err != nil {
		return err
	}
	return err
}

func (r PedidoValidado) Serialize(w io.Writer) error {
	return writePedidoValidado(r, w)
}

func (r PedidoValidado) Schema() string {
	return "{\"fields\":[{\"name\":\"NroPedido\",\"type\":\"string\"},{\"name\":\"NroPedidoExterno\",\"type\":\"string\"},{\"name\":\"SubPedido\",\"type\":\"string\"},{\"name\":\"Bandeja\",\"type\":\"string\"},{\"name\":\"Estado\",\"type\":\"string\"},{\"name\":\"IpImpresora\",\"type\":\"string\"},{\"name\":\"Cliente\",\"type\":\"string\"},{\"default\":null,\"name\":\"SchemaDb\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.CtrlExp.Events.Record.PedidoValidado\",\"type\":\"record\"}"
}

func (r PedidoValidado) SchemaName() string {
	return "Andreani.CtrlExp.Events.Record.PedidoValidado"
}

func (_ PedidoValidado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ PedidoValidado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ PedidoValidado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ PedidoValidado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ PedidoValidado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ PedidoValidado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ PedidoValidado) SetString(v string)   { panic("Unsupported operation") }
func (_ PedidoValidado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *PedidoValidado) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.NroPedido}

		return w

	case 1:
		w := types.String{Target: &r.NroPedidoExterno}

		return w

	case 2:
		w := types.String{Target: &r.SubPedido}

		return w

	case 3:
		w := types.String{Target: &r.Bandeja}

		return w

	case 4:
		w := types.String{Target: &r.Estado}

		return w

	case 5:
		w := types.String{Target: &r.IpImpresora}

		return w

	case 6:
		w := types.String{Target: &r.Cliente}

		return w

	case 7:
		r.SchemaDb = NewUnionNullString()

		return r.SchemaDb
	}
	panic("Unknown field index")
}

func (r *PedidoValidado) SetDefault(i int) {
	switch i {
	case 7:
		r.SchemaDb = nil
		return
	}
	panic("Unknown field index")
}

func (r *PedidoValidado) NullField(i int) {
	switch i {
	case 7:
		r.SchemaDb = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ PedidoValidado) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ PedidoValidado) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ PedidoValidado) HintSize(int)                     { panic("Unsupported operation") }
func (_ PedidoValidado) Finalize()                        {}

func (_ PedidoValidado) AvroCRC64Fingerprint() []byte {
	return []byte(PedidoValidadoAvroCRC64Fingerprint)
}

func (r PedidoValidado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["NroPedido"], err = json.Marshal(r.NroPedido)
	if err != nil {
		return nil, err
	}
	output["NroPedidoExterno"], err = json.Marshal(r.NroPedidoExterno)
	if err != nil {
		return nil, err
	}
	output["SubPedido"], err = json.Marshal(r.SubPedido)
	if err != nil {
		return nil, err
	}
	output["Bandeja"], err = json.Marshal(r.Bandeja)
	if err != nil {
		return nil, err
	}
	output["Estado"], err = json.Marshal(r.Estado)
	if err != nil {
		return nil, err
	}
	output["IpImpresora"], err = json.Marshal(r.IpImpresora)
	if err != nil {
		return nil, err
	}
	output["Cliente"], err = json.Marshal(r.Cliente)
	if err != nil {
		return nil, err
	}
	output["SchemaDb"], err = json.Marshal(r.SchemaDb)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *PedidoValidado) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["NroPedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NroPedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NroPedido")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NroPedidoExterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NroPedidoExterno); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NroPedidoExterno")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SubPedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SubPedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SubPedido")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Bandeja"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Bandeja); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Bandeja")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Estado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Estado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["IpImpresora"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IpImpresora); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IpImpresora")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Cliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Cliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SchemaDb"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SchemaDb); err != nil {
			return err
		}
	} else {
		r.SchemaDb = NewUnionNullString()

		r.SchemaDb = nil
	}
	return nil
}
