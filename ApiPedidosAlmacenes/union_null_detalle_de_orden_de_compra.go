// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoDeAlmacenSCE.avsc
 */
package ApiPedidosAlmacenesEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullDetalleDeOrdenDeCompraTypeEnum int

const (
	UnionNullDetalleDeOrdenDeCompraTypeEnumDetalleDeOrdenDeCompra UnionNullDetalleDeOrdenDeCompraTypeEnum = 1
)

type UnionNullDetalleDeOrdenDeCompra struct {
	Null                   *types.NullVal
	DetalleDeOrdenDeCompra DetalleDeOrdenDeCompra
	UnionType              UnionNullDetalleDeOrdenDeCompraTypeEnum
}

func writeUnionNullDetalleDeOrdenDeCompra(r *UnionNullDetalleDeOrdenDeCompra, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullDetalleDeOrdenDeCompraTypeEnumDetalleDeOrdenDeCompra:
		return writeDetalleDeOrdenDeCompra(r.DetalleDeOrdenDeCompra, w)
	}
	return fmt.Errorf("invalid value for *UnionNullDetalleDeOrdenDeCompra")
}

func NewUnionNullDetalleDeOrdenDeCompra() *UnionNullDetalleDeOrdenDeCompra {
	return &UnionNullDetalleDeOrdenDeCompra{}
}

func (r *UnionNullDetalleDeOrdenDeCompra) Serialize(w io.Writer) error {
	return writeUnionNullDetalleDeOrdenDeCompra(r, w)
}

func DeserializeUnionNullDetalleDeOrdenDeCompra(r io.Reader) (*UnionNullDetalleDeOrdenDeCompra, error) {
	t := NewUnionNullDetalleDeOrdenDeCompra()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionNullDetalleDeOrdenDeCompraFromSchema(r io.Reader, schema string) (*UnionNullDetalleDeOrdenDeCompra, error) {
	t := NewUnionNullDetalleDeOrdenDeCompra()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r *UnionNullDetalleDeOrdenDeCompra) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"NumeroDeLineaDeCliente\",\"type\":\"string\"},{\"name\":\"OrdenDeCompraDeCliente\",\"type\":\"string\"},{\"name\":\"NotasDeLinea\",\"type\":[\"null\",\"string\"]},{\"name\":\"NumeroDeLinea\",\"type\":\"string\"},{\"name\":\"CantidadPedida\",\"type\":\"double\"},{\"name\":\"articulo\",\"type\":{\"fields\":[{\"name\":\"Codigo\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Lote\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Codigo\",\"type\":\"string\"},{\"name\":\"LoteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"name\":\"LoteSecundario\",\"type\":[\"null\",\"string\"]},{\"name\":\"FechaDeVencimiento\",\"type\":\"int\"},{\"name\":\"OtrosDatos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Meta\",\"type\":\"string\"},{\"name\":\"Contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]}],\"name\":\"LoteArticulo\",\"type\":\"record\"}]},{\"name\":\"OtrosDatos\",\"type\":[\"null\",\"Andreani.WapAltaDePedidoSolicitada.Events.Record.ListaDePropiedades\"]}],\"name\":\"Articulo\",\"type\":\"record\"}},{\"name\":\"ValorDeclarado\",\"type\":[\"null\",\"string\"]},{\"name\":\"UOM\",\"type\":[\"null\",\"string\"]},{\"name\":\"CamposLibres\",\"type\":[\"null\",\"Andreani.WapAltaDePedidoSolicitada.Events.Record.ListaDePropiedades\"]},{\"name\":\"FechaOrdenDeCompra\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoMaterial\",\"type\":[\"null\",\"string\"]}],\"name\":\"DetalleDeOrdenDeCompra\",\"type\":\"record\"}]"
}

func (_ *UnionNullDetalleDeOrdenDeCompra) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullDetalleDeOrdenDeCompra) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullDetalleDeOrdenDeCompra) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullDetalleDeOrdenDeCompra) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullDetalleDeOrdenDeCompra) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullDetalleDeOrdenDeCompra) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullDetalleDeOrdenDeCompra) SetLong(v int64) {

	r.UnionType = (UnionNullDetalleDeOrdenDeCompraTypeEnum)(v)
}

func (r *UnionNullDetalleDeOrdenDeCompra) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.DetalleDeOrdenDeCompra = NewDetalleDeOrdenDeCompra()
		return &types.Record{Target: (&r.DetalleDeOrdenDeCompra)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullDetalleDeOrdenDeCompra) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullDetalleDeOrdenDeCompra) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullDetalleDeOrdenDeCompra) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullDetalleDeOrdenDeCompra) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullDetalleDeOrdenDeCompra) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullDetalleDeOrdenDeCompra) Finalize()                {}

func (r *UnionNullDetalleDeOrdenDeCompra) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullDetalleDeOrdenDeCompraTypeEnumDetalleDeOrdenDeCompra:
		return json.Marshal(map[string]interface{}{"Andreani.WapAltaDePedidoSolicitada.Events.Record.DetalleDeOrdenDeCompra": r.DetalleDeOrdenDeCompra})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullDetalleDeOrdenDeCompra")
}

func (r *UnionNullDetalleDeOrdenDeCompra) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.WapAltaDePedidoSolicitada.Events.Record.DetalleDeOrdenDeCompra"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.DetalleDeOrdenDeCompra)
	}
	return fmt.Errorf("invalid value for *UnionNullDetalleDeOrdenDeCompra")
}
