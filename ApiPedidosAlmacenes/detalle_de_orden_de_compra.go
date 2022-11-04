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

var _ = fmt.Printf

type DetalleDeOrdenDeCompra struct {
	NumeroDeLineaDeCliente string `json:"NumeroDeLineaDeCliente"`

	OrdenDeCompraDeCliente string `json:"OrdenDeCompraDeCliente"`

	NotasDeLinea *UnionNullString `json:"NotasDeLinea"`

	NumeroDeLinea string `json:"NumeroDeLinea"`

	CantidadPedida float64 `json:"CantidadPedida"`

	Articulo Articulo `json:"articulo"`

	ValorDeclarado *UnionNullString `json:"ValorDeclarado"`

	UOM *UnionNullString `json:"UOM"`

	CamposLibres *UnionNullListaDePropiedades `json:"CamposLibres"`

	FechaOrdenDeCompra *UnionNullString `json:"FechaOrdenDeCompra"`

	TipoMaterial *UnionNullString `json:"TipoMaterial"`
}

const DetalleDeOrdenDeCompraAvroCRC64Fingerprint = "%\xeb \x9e\xc2&9\xe5"

func NewDetalleDeOrdenDeCompra() DetalleDeOrdenDeCompra {
	r := DetalleDeOrdenDeCompra{}
	r.Articulo = NewArticulo()

	return r
}

func DeserializeDetalleDeOrdenDeCompra(r io.Reader) (DetalleDeOrdenDeCompra, error) {
	t := NewDetalleDeOrdenDeCompra()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDetalleDeOrdenDeCompraFromSchema(r io.Reader, schema string) (DetalleDeOrdenDeCompra, error) {
	t := NewDetalleDeOrdenDeCompra()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDetalleDeOrdenDeCompra(r DetalleDeOrdenDeCompra, w io.Writer) error {
	var err error
	err = vm.WriteString(r.NumeroDeLineaDeCliente, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.OrdenDeCompraDeCliente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NotasDeLinea, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroDeLinea, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.CantidadPedida, w)
	if err != nil {
		return err
	}
	err = writeArticulo(r.Articulo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ValorDeclarado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.UOM, w)
	if err != nil {
		return err
	}
	err = writeUnionNullListaDePropiedades(r.CamposLibres, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaOrdenDeCompra, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoMaterial, w)
	if err != nil {
		return err
	}
	return err
}

func (r DetalleDeOrdenDeCompra) Serialize(w io.Writer) error {
	return writeDetalleDeOrdenDeCompra(r, w)
}

func (r DetalleDeOrdenDeCompra) Schema() string {
	return "{\"fields\":[{\"name\":\"NumeroDeLineaDeCliente\",\"type\":\"string\"},{\"name\":\"OrdenDeCompraDeCliente\",\"type\":\"string\"},{\"name\":\"NotasDeLinea\",\"type\":[\"null\",\"string\"]},{\"name\":\"NumeroDeLinea\",\"type\":\"string\"},{\"name\":\"CantidadPedida\",\"type\":\"double\"},{\"name\":\"articulo\",\"type\":{\"fields\":[{\"name\":\"Codigo\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Lote\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Codigo\",\"type\":\"string\"},{\"name\":\"LoteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"name\":\"LoteSecundario\",\"type\":[\"null\",\"string\"]},{\"name\":\"FechaDeVencimiento\",\"type\":\"int\"},{\"name\":\"OtrosDatos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Meta\",\"type\":\"string\"},{\"name\":\"Contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]}],\"name\":\"LoteArticulo\",\"type\":\"record\"}]},{\"name\":\"OtrosDatos\",\"type\":[\"null\",\"Andreani.WapAltaDePedidoSolicitada.Events.Record.ListaDePropiedades\"]}],\"name\":\"Articulo\",\"type\":\"record\"}},{\"name\":\"ValorDeclarado\",\"type\":[\"null\",\"string\"]},{\"name\":\"UOM\",\"type\":[\"null\",\"string\"]},{\"name\":\"CamposLibres\",\"type\":[\"null\",\"Andreani.WapAltaDePedidoSolicitada.Events.Record.ListaDePropiedades\"]},{\"name\":\"FechaOrdenDeCompra\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoMaterial\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.WapAltaDePedidoSolicitada.Events.Record.DetalleDeOrdenDeCompra\",\"type\":\"record\"}"
}

func (r DetalleDeOrdenDeCompra) SchemaName() string {
	return "Andreani.WapAltaDePedidoSolicitada.Events.Record.DetalleDeOrdenDeCompra"
}

func (_ DetalleDeOrdenDeCompra) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeCompra) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeCompra) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeCompra) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeCompra) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeCompra) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeCompra) SetString(v string)   { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeCompra) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DetalleDeOrdenDeCompra) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.NumeroDeLineaDeCliente}

		return w

	case 1:
		w := types.String{Target: &r.OrdenDeCompraDeCliente}

		return w

	case 2:
		r.NotasDeLinea = NewUnionNullString()

		return r.NotasDeLinea
	case 3:
		w := types.String{Target: &r.NumeroDeLinea}

		return w

	case 4:
		w := types.Double{Target: &r.CantidadPedida}

		return w

	case 5:
		r.Articulo = NewArticulo()

		w := types.Record{Target: &r.Articulo}

		return w

	case 6:
		r.ValorDeclarado = NewUnionNullString()

		return r.ValorDeclarado
	case 7:
		r.UOM = NewUnionNullString()

		return r.UOM
	case 8:
		r.CamposLibres = NewUnionNullListaDePropiedades()

		return r.CamposLibres
	case 9:
		r.FechaOrdenDeCompra = NewUnionNullString()

		return r.FechaOrdenDeCompra
	case 10:
		r.TipoMaterial = NewUnionNullString()

		return r.TipoMaterial
	}
	panic("Unknown field index")
}

func (r *DetalleDeOrdenDeCompra) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *DetalleDeOrdenDeCompra) NullField(i int) {
	switch i {
	case 2:
		r.NotasDeLinea = nil
		return
	case 6:
		r.ValorDeclarado = nil
		return
	case 7:
		r.UOM = nil
		return
	case 8:
		r.CamposLibres = nil
		return
	case 9:
		r.FechaOrdenDeCompra = nil
		return
	case 10:
		r.TipoMaterial = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ DetalleDeOrdenDeCompra) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeCompra) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeCompra) HintSize(int)                     { panic("Unsupported operation") }
func (_ DetalleDeOrdenDeCompra) Finalize()                        {}

func (_ DetalleDeOrdenDeCompra) AvroCRC64Fingerprint() []byte {
	return []byte(DetalleDeOrdenDeCompraAvroCRC64Fingerprint)
}

func (r DetalleDeOrdenDeCompra) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["NumeroDeLineaDeCliente"], err = json.Marshal(r.NumeroDeLineaDeCliente)
	if err != nil {
		return nil, err
	}
	output["OrdenDeCompraDeCliente"], err = json.Marshal(r.OrdenDeCompraDeCliente)
	if err != nil {
		return nil, err
	}
	output["NotasDeLinea"], err = json.Marshal(r.NotasDeLinea)
	if err != nil {
		return nil, err
	}
	output["NumeroDeLinea"], err = json.Marshal(r.NumeroDeLinea)
	if err != nil {
		return nil, err
	}
	output["CantidadPedida"], err = json.Marshal(r.CantidadPedida)
	if err != nil {
		return nil, err
	}
	output["articulo"], err = json.Marshal(r.Articulo)
	if err != nil {
		return nil, err
	}
	output["ValorDeclarado"], err = json.Marshal(r.ValorDeclarado)
	if err != nil {
		return nil, err
	}
	output["UOM"], err = json.Marshal(r.UOM)
	if err != nil {
		return nil, err
	}
	output["CamposLibres"], err = json.Marshal(r.CamposLibres)
	if err != nil {
		return nil, err
	}
	output["FechaOrdenDeCompra"], err = json.Marshal(r.FechaOrdenDeCompra)
	if err != nil {
		return nil, err
	}
	output["TipoMaterial"], err = json.Marshal(r.TipoMaterial)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DetalleDeOrdenDeCompra) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeLineaDeCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeLineaDeCliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NumeroDeLineaDeCliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OrdenDeCompraDeCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenDeCompraDeCliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrdenDeCompraDeCliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NotasDeLinea"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NotasDeLinea); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NotasDeLinea")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeLinea"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeLinea); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NumeroDeLinea")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CantidadPedida"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadPedida); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CantidadPedida")
	}
	val = func() json.RawMessage {
		if v, ok := fields["articulo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Articulo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for articulo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ValorDeclarado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ValorDeclarado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ValorDeclarado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["UOM"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UOM); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for UOM")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CamposLibres"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CamposLibres); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CamposLibres")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaOrdenDeCompra"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaOrdenDeCompra); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaOrdenDeCompra")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoMaterial"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoMaterial); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoMaterial")
	}
	return nil
}