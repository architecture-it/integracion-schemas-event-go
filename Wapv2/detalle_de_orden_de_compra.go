// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     DetalleDeOrdenDeCompra.avsc
 */
package Wapv2Events

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
	NumeroDeLineaDeCliente string `json:"numeroDeLineaDeCliente"`

	OrdenDeCompraDeCliente string `json:"ordenDeCompraDeCliente"`

	NotasDeLinea *UnionNullString `json:"notasDeLinea"`

	NumeroDeLinea string `json:"numeroDeLinea"`

	CantidadPedida float64 `json:"cantidadPedida"`

	Articulo Articulo `json:"articulo"`

	ValorDeclarado *UnionNullString `json:"valorDeclarado"`

	UOM *UnionNullString `json:"UOM"`

	CamposLibres *UnionNullListaDePropiedades `json:"camposLibres"`

	FechaOrdenDeCompra *UnionNullString `json:"fechaOrdenDeCompra"`

	TipoMaterial *UnionNullString `json:"tipoMaterial"`
}

const DetalleDeOrdenDeCompraAvroCRC64Fingerprint = ";\x83x\xc0ƅ\x1f\xf1"

func NewDetalleDeOrdenDeCompra() DetalleDeOrdenDeCompra {
	r := DetalleDeOrdenDeCompra{}
	r.NotasDeLinea = nil
	r.Articulo = NewArticulo()

	r.ValorDeclarado = nil
	r.UOM = nil
	r.CamposLibres = nil
	r.FechaOrdenDeCompra = nil
	r.TipoMaterial = nil
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
	return "{\"fields\":[{\"name\":\"numeroDeLineaDeCliente\",\"type\":\"string\"},{\"name\":\"ordenDeCompraDeCliente\",\"type\":\"string\"},{\"default\":null,\"name\":\"notasDeLinea\",\"type\":[\"null\",\"string\"]},{\"name\":\"numeroDeLinea\",\"type\":\"string\"},{\"name\":\"cantidadPedida\",\"type\":\"double\"},{\"name\":\"articulo\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"cantidad\",\"type\":\"double\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"numeropedido\",\"type\":[\"null\",\"string\"]},{\"name\":\"zonaConsumo\",\"type\":[\"null\",\"string\"]},{\"name\":\"unidadmedida\",\"type\":\"string\"},{\"default\":null,\"name\":\"datosadicionales\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"name\":\"lote\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaDeVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",\"string\"]},{\"name\":\"estado\",\"type\":[\"null\",\"string\"]}],\"name\":\"LoteArticulo\",\"type\":\"record\"}}],\"name\":\"Articulo\",\"type\":\"record\"}},{\"default\":null,\"name\":\"valorDeclarado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"UOM\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"camposLibres\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.ListaDePropiedades\"]},{\"default\":null,\"name\":\"fechaOrdenDeCompra\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoMaterial\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.Wapv2.Events.Record.DetalleDeOrdenDeCompra\",\"type\":\"record\"}"
}

func (r DetalleDeOrdenDeCompra) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.DetalleDeOrdenDeCompra"
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
	output["numeroDeLineaDeCliente"], err = json.Marshal(r.NumeroDeLineaDeCliente)
	if err != nil {
		return nil, err
	}
	output["ordenDeCompraDeCliente"], err = json.Marshal(r.OrdenDeCompraDeCliente)
	if err != nil {
		return nil, err
	}
	output["notasDeLinea"], err = json.Marshal(r.NotasDeLinea)
	if err != nil {
		return nil, err
	}
	output["numeroDeLinea"], err = json.Marshal(r.NumeroDeLinea)
	if err != nil {
		return nil, err
	}
	output["cantidadPedida"], err = json.Marshal(r.CantidadPedida)
	if err != nil {
		return nil, err
	}
	output["articulo"], err = json.Marshal(r.Articulo)
	if err != nil {
		return nil, err
	}
	output["valorDeclarado"], err = json.Marshal(r.ValorDeclarado)
	if err != nil {
		return nil, err
	}
	output["UOM"], err = json.Marshal(r.UOM)
	if err != nil {
		return nil, err
	}
	output["camposLibres"], err = json.Marshal(r.CamposLibres)
	if err != nil {
		return nil, err
	}
	output["fechaOrdenDeCompra"], err = json.Marshal(r.FechaOrdenDeCompra)
	if err != nil {
		return nil, err
	}
	output["tipoMaterial"], err = json.Marshal(r.TipoMaterial)
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
		if v, ok := fields["numeroDeLineaDeCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeLineaDeCliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroDeLineaDeCliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ordenDeCompraDeCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenDeCompraDeCliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ordenDeCompraDeCliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["notasDeLinea"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NotasDeLinea); err != nil {
			return err
		}
	} else {
		r.NotasDeLinea = NewUnionNullString()

		r.NotasDeLinea = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroDeLinea"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeLinea); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroDeLinea")
	}
	val = func() json.RawMessage {
		if v, ok := fields["cantidadPedida"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadPedida); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cantidadPedida")
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
		if v, ok := fields["valorDeclarado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ValorDeclarado); err != nil {
			return err
		}
	} else {
		r.ValorDeclarado = NewUnionNullString()

		r.ValorDeclarado = nil
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
		r.UOM = NewUnionNullString()

		r.UOM = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["camposLibres"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CamposLibres); err != nil {
			return err
		}
	} else {
		r.CamposLibres = NewUnionNullListaDePropiedades()

		r.CamposLibres = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaOrdenDeCompra"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaOrdenDeCompra); err != nil {
			return err
		}
	} else {
		r.FechaOrdenDeCompra = NewUnionNullString()

		r.FechaOrdenDeCompra = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoMaterial"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoMaterial); err != nil {
			return err
		}
	} else {
		r.TipoMaterial = NewUnionNullString()

		r.TipoMaterial = nil
	}
	return nil
}
