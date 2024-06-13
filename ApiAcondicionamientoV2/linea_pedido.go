// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoAcondi.avsc
 */
package ApiAcondicionamientoV2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type LineaPedido struct {
	AlmacenCliente *UnionNullString `json:"almacenCliente"`

	Codigo *UnionNullString `json:"codigo"`

	Cantidad *UnionNullString `json:"cantidad"`

	LoteDeFabricante *UnionNullString `json:"loteDeFabricante"`

	LoteSecundario *UnionNullString `json:"loteSecundario"`

	FechaDeVencimiento *UnionNullString `json:"fechaDeVencimiento"`

	EstadoLote *UnionNullString `json:"estadoLote"`
}

const LineaPedidoAvroCRC64Fingerprint = "#\xc8\xc4y'\xcb\tM"

func NewLineaPedido() LineaPedido {
	r := LineaPedido{}
	r.AlmacenCliente = nil
	r.Codigo = nil
	r.Cantidad = nil
	r.LoteDeFabricante = nil
	r.LoteSecundario = nil
	r.FechaDeVencimiento = nil
	r.EstadoLote = nil
	return r
}

func DeserializeLineaPedido(r io.Reader) (LineaPedido, error) {
	t := NewLineaPedido()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeLineaPedidoFromSchema(r io.Reader, schema string) (LineaPedido, error) {
	t := NewLineaPedido()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeLineaPedido(r LineaPedido, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.AlmacenCliente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Codigo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Cantidad, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LoteDeFabricante, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LoteSecundario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaDeVencimiento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstadoLote, w)
	if err != nil {
		return err
	}
	return err
}

func (r LineaPedido) Serialize(w io.Writer) error {
	return writeLineaPedido(r, w)
}

func (r LineaPedido) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"almacenCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cantidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaDeVencimiento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoLote\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.ApiAcondicionamientoV2.Events.Common.LineaPedido\",\"type\":\"record\"}"
}

func (r LineaPedido) SchemaName() string {
	return "Andreani.ApiAcondicionamientoV2.Events.Common.LineaPedido"
}

func (_ LineaPedido) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ LineaPedido) SetInt(v int32)       { panic("Unsupported operation") }
func (_ LineaPedido) SetLong(v int64)      { panic("Unsupported operation") }
func (_ LineaPedido) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ LineaPedido) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ LineaPedido) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ LineaPedido) SetString(v string)   { panic("Unsupported operation") }
func (_ LineaPedido) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *LineaPedido) Get(i int) types.Field {
	switch i {
	case 0:
		r.AlmacenCliente = NewUnionNullString()

		return r.AlmacenCliente
	case 1:
		r.Codigo = NewUnionNullString()

		return r.Codigo
	case 2:
		r.Cantidad = NewUnionNullString()

		return r.Cantidad
	case 3:
		r.LoteDeFabricante = NewUnionNullString()

		return r.LoteDeFabricante
	case 4:
		r.LoteSecundario = NewUnionNullString()

		return r.LoteSecundario
	case 5:
		r.FechaDeVencimiento = NewUnionNullString()

		return r.FechaDeVencimiento
	case 6:
		r.EstadoLote = NewUnionNullString()

		return r.EstadoLote
	}
	panic("Unknown field index")
}

func (r *LineaPedido) SetDefault(i int) {
	switch i {
	case 0:
		r.AlmacenCliente = nil
		return
	case 1:
		r.Codigo = nil
		return
	case 2:
		r.Cantidad = nil
		return
	case 3:
		r.LoteDeFabricante = nil
		return
	case 4:
		r.LoteSecundario = nil
		return
	case 5:
		r.FechaDeVencimiento = nil
		return
	case 6:
		r.EstadoLote = nil
		return
	}
	panic("Unknown field index")
}

func (r *LineaPedido) NullField(i int) {
	switch i {
	case 0:
		r.AlmacenCliente = nil
		return
	case 1:
		r.Codigo = nil
		return
	case 2:
		r.Cantidad = nil
		return
	case 3:
		r.LoteDeFabricante = nil
		return
	case 4:
		r.LoteSecundario = nil
		return
	case 5:
		r.FechaDeVencimiento = nil
		return
	case 6:
		r.EstadoLote = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ LineaPedido) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ LineaPedido) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ LineaPedido) HintSize(int)                     { panic("Unsupported operation") }
func (_ LineaPedido) Finalize()                        {}

func (_ LineaPedido) AvroCRC64Fingerprint() []byte {
	return []byte(LineaPedidoAvroCRC64Fingerprint)
}

func (r LineaPedido) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["almacenCliente"], err = json.Marshal(r.AlmacenCliente)
	if err != nil {
		return nil, err
	}
	output["codigo"], err = json.Marshal(r.Codigo)
	if err != nil {
		return nil, err
	}
	output["cantidad"], err = json.Marshal(r.Cantidad)
	if err != nil {
		return nil, err
	}
	output["loteDeFabricante"], err = json.Marshal(r.LoteDeFabricante)
	if err != nil {
		return nil, err
	}
	output["loteSecundario"], err = json.Marshal(r.LoteSecundario)
	if err != nil {
		return nil, err
	}
	output["fechaDeVencimiento"], err = json.Marshal(r.FechaDeVencimiento)
	if err != nil {
		return nil, err
	}
	output["estadoLote"], err = json.Marshal(r.EstadoLote)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *LineaPedido) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["almacenCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AlmacenCliente); err != nil {
			return err
		}
	} else {
		r.AlmacenCliente = NewUnionNullString()

		r.AlmacenCliente = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Codigo); err != nil {
			return err
		}
	} else {
		r.Codigo = NewUnionNullString()

		r.Codigo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["cantidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cantidad); err != nil {
			return err
		}
	} else {
		r.Cantidad = NewUnionNullString()

		r.Cantidad = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["loteDeFabricante"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteDeFabricante); err != nil {
			return err
		}
	} else {
		r.LoteDeFabricante = NewUnionNullString()

		r.LoteDeFabricante = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["loteSecundario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteSecundario); err != nil {
			return err
		}
	} else {
		r.LoteSecundario = NewUnionNullString()

		r.LoteSecundario = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaDeVencimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaDeVencimiento); err != nil {
			return err
		}
	} else {
		r.FechaDeVencimiento = NewUnionNullString()

		r.FechaDeVencimiento = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["estadoLote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstadoLote); err != nil {
			return err
		}
	} else {
		r.EstadoLote = NewUnionNullString()

		r.EstadoLote = nil
	}
	return nil
}
