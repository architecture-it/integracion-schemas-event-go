// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhSCEPedidosEmpaquetados.avsc
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

type Detalle struct {
	SKU string `json:"SKU"`

	LineaPedidoWH string `json:"LineaPedidoWH"`

	LineaExterna string `json:"LineaExterna"`

	PaqueteLote string `json:"PaqueteLote"`

	LoteCajitaFabricante string `json:"LoteCajitaFabricante"`

	LoteSecundario string `json:"LoteSecundario"`

	FechaFabricacion *UnionNullLong `json:"FechaFabricacion"`

	FechaVencimiento *UnionNullLong `json:"FechaVencimiento"`

	ProductoTrazable string `json:"ProductoTrazable"`

	Almacen string `json:"Almacen"`

	EstadoLote string `json:"EstadoLote"`

	BloqueoUbicacion string `json:"BloqueoUbicacion"`

	VidaUtilLote string `json:"VidaUtilLote"`

	EntregaAntesDe *UnionNullLong `json:"EntregaAntesDe"`

	ConsumoAntesDe *UnionNullLong `json:"ConsumoAntesDe"`

	EsfuerzoLineaApiPlani *UnionNullString `json:"EsfuerzoLineaApiPlani"`

	VidaUtil *UnionNullString `json:"VidaUtil"`

	Cantidad float32 `json:"Cantidad"`

	CantidadPickeada float32 `json:"CantidadPickeada"`

	CantidadExpedida float32 `json:"CantidadExpedida"`

	EstadoLineaCodigo string `json:"EstadoLineaCodigo"`

	EstadoLineaDescripcion string `json:"EstadoLineaDescripcion"`

	CantidadBultos string `json:"CantidadBultos"`

	Series []Series `json:"Series"`

	Contrato *UnionNullString `json:"Contrato"`

	Nomenclatura *UnionNullString `json:"Nomenclatura"`

	TipoPedido *UnionNullString `json:"TipoPedido"`

	TipoPedidoDescripcion *UnionNullString `json:"TipoPedidoDescripcion"`

	DetalleAcondicionamiento *UnionNullString `json:"DetalleAcondicionamiento"`

	EstadoOTAcondicionamiento *UnionNullString `json:"EstadoOTAcondicionamiento"`
}

const DetalleAvroCRC64Fingerprint = "p\x17R^lt^g"

func NewDetalle() Detalle {
	r := Detalle{}
	r.Series = make([]Series, 0)

	return r
}

func DeserializeDetalle(r io.Reader) (Detalle, error) {
	t := NewDetalle()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDetalleFromSchema(r io.Reader, schema string) (Detalle, error) {
	t := NewDetalle()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDetalle(r Detalle, w io.Writer) error {
	var err error
	err = vm.WriteString(r.SKU, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.LineaPedidoWH, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.LineaExterna, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.PaqueteLote, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.LoteCajitaFabricante, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.LoteSecundario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaFabricacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaVencimiento, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.ProductoTrazable, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Almacen, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EstadoLote, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.BloqueoUbicacion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.VidaUtilLote, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.EntregaAntesDe, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.ConsumoAntesDe, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EsfuerzoLineaApiPlani, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.VidaUtil, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.Cantidad, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.CantidadPickeada, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.CantidadExpedida, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EstadoLineaCodigo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EstadoLineaDescripcion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CantidadBultos, w)
	if err != nil {
		return err
	}
	err = writeArraySeries(r.Series, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Contrato, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Nomenclatura, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoPedido, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoPedidoDescripcion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DetalleAcondicionamiento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstadoOTAcondicionamiento, w)
	if err != nil {
		return err
	}
	return err
}

func (r Detalle) Serialize(w io.Writer) error {
	return writeDetalle(r, w)
}

func (r Detalle) Schema() string {
	return "{\"fields\":[{\"name\":\"SKU\",\"type\":\"string\"},{\"name\":\"LineaPedidoWH\",\"type\":\"string\"},{\"name\":\"LineaExterna\",\"type\":\"string\"},{\"name\":\"PaqueteLote\",\"type\":\"string\"},{\"name\":\"LoteCajitaFabricante\",\"type\":\"string\"},{\"name\":\"LoteSecundario\",\"type\":\"string\"},{\"name\":\"FechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ProductoTrazable\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"EstadoLote\",\"type\":\"string\"},{\"name\":\"BloqueoUbicacion\",\"type\":\"string\"},{\"name\":\"VidaUtilLote\",\"type\":\"string\"},{\"name\":\"EntregaAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ConsumoAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"EsfuerzoLineaApiPlani\",\"type\":[\"null\",\"string\"]},{\"name\":\"VidaUtil\",\"type\":[\"null\",\"string\"]},{\"name\":\"Cantidad\",\"type\":\"float\"},{\"name\":\"CantidadPickeada\",\"type\":\"float\"},{\"name\":\"CantidadExpedida\",\"type\":\"float\"},{\"name\":\"EstadoLineaCodigo\",\"type\":\"string\"},{\"name\":\"EstadoLineaDescripcion\",\"type\":\"string\"},{\"name\":\"CantidadBultos\",\"type\":\"string\"},{\"name\":\"Series\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Serie\",\"type\":\"string\"}],\"name\":\"Series\",\"type\":\"record\"},\"type\":\"array\"}},{\"name\":\"Contrato\",\"type\":[\"null\",\"string\"]},{\"name\":\"Nomenclatura\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoPedido\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoPedidoDescripcion\",\"type\":[\"null\",\"string\"]},{\"name\":\"DetalleAcondicionamiento\",\"type\":[\"null\",\"string\"]},{\"name\":\"EstadoOTAcondicionamiento\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.EventoWhPedidos.Events.SCEPedidosEmpaquetadosCommon.Detalle\",\"type\":\"record\"}"
}

func (r Detalle) SchemaName() string {
	return "Andreani.EventoWhPedidos.Events.SCEPedidosEmpaquetadosCommon.Detalle"
}

func (_ Detalle) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Detalle) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Detalle) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Detalle) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Detalle) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Detalle) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Detalle) SetString(v string)   { panic("Unsupported operation") }
func (_ Detalle) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Detalle) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.SKU}

		return w

	case 1:
		w := types.String{Target: &r.LineaPedidoWH}

		return w

	case 2:
		w := types.String{Target: &r.LineaExterna}

		return w

	case 3:
		w := types.String{Target: &r.PaqueteLote}

		return w

	case 4:
		w := types.String{Target: &r.LoteCajitaFabricante}

		return w

	case 5:
		w := types.String{Target: &r.LoteSecundario}

		return w

	case 6:
		r.FechaFabricacion = NewUnionNullLong()

		return r.FechaFabricacion
	case 7:
		r.FechaVencimiento = NewUnionNullLong()

		return r.FechaVencimiento
	case 8:
		w := types.String{Target: &r.ProductoTrazable}

		return w

	case 9:
		w := types.String{Target: &r.Almacen}

		return w

	case 10:
		w := types.String{Target: &r.EstadoLote}

		return w

	case 11:
		w := types.String{Target: &r.BloqueoUbicacion}

		return w

	case 12:
		w := types.String{Target: &r.VidaUtilLote}

		return w

	case 13:
		r.EntregaAntesDe = NewUnionNullLong()

		return r.EntregaAntesDe
	case 14:
		r.ConsumoAntesDe = NewUnionNullLong()

		return r.ConsumoAntesDe
	case 15:
		r.EsfuerzoLineaApiPlani = NewUnionNullString()

		return r.EsfuerzoLineaApiPlani
	case 16:
		r.VidaUtil = NewUnionNullString()

		return r.VidaUtil
	case 17:
		w := types.Float{Target: &r.Cantidad}

		return w

	case 18:
		w := types.Float{Target: &r.CantidadPickeada}

		return w

	case 19:
		w := types.Float{Target: &r.CantidadExpedida}

		return w

	case 20:
		w := types.String{Target: &r.EstadoLineaCodigo}

		return w

	case 21:
		w := types.String{Target: &r.EstadoLineaDescripcion}

		return w

	case 22:
		w := types.String{Target: &r.CantidadBultos}

		return w

	case 23:
		r.Series = make([]Series, 0)

		w := ArraySeriesWrapper{Target: &r.Series}

		return w

	case 24:
		r.Contrato = NewUnionNullString()

		return r.Contrato
	case 25:
		r.Nomenclatura = NewUnionNullString()

		return r.Nomenclatura
	case 26:
		r.TipoPedido = NewUnionNullString()

		return r.TipoPedido
	case 27:
		r.TipoPedidoDescripcion = NewUnionNullString()

		return r.TipoPedidoDescripcion
	case 28:
		r.DetalleAcondicionamiento = NewUnionNullString()

		return r.DetalleAcondicionamiento
	case 29:
		r.EstadoOTAcondicionamiento = NewUnionNullString()

		return r.EstadoOTAcondicionamiento
	}
	panic("Unknown field index")
}

func (r *Detalle) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Detalle) NullField(i int) {
	switch i {
	case 6:
		r.FechaFabricacion = nil
		return
	case 7:
		r.FechaVencimiento = nil
		return
	case 13:
		r.EntregaAntesDe = nil
		return
	case 14:
		r.ConsumoAntesDe = nil
		return
	case 15:
		r.EsfuerzoLineaApiPlani = nil
		return
	case 16:
		r.VidaUtil = nil
		return
	case 24:
		r.Contrato = nil
		return
	case 25:
		r.Nomenclatura = nil
		return
	case 26:
		r.TipoPedido = nil
		return
	case 27:
		r.TipoPedidoDescripcion = nil
		return
	case 28:
		r.DetalleAcondicionamiento = nil
		return
	case 29:
		r.EstadoOTAcondicionamiento = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Detalle) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Detalle) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Detalle) HintSize(int)                     { panic("Unsupported operation") }
func (_ Detalle) Finalize()                        {}

func (_ Detalle) AvroCRC64Fingerprint() []byte {
	return []byte(DetalleAvroCRC64Fingerprint)
}

func (r Detalle) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["SKU"], err = json.Marshal(r.SKU)
	if err != nil {
		return nil, err
	}
	output["LineaPedidoWH"], err = json.Marshal(r.LineaPedidoWH)
	if err != nil {
		return nil, err
	}
	output["LineaExterna"], err = json.Marshal(r.LineaExterna)
	if err != nil {
		return nil, err
	}
	output["PaqueteLote"], err = json.Marshal(r.PaqueteLote)
	if err != nil {
		return nil, err
	}
	output["LoteCajitaFabricante"], err = json.Marshal(r.LoteCajitaFabricante)
	if err != nil {
		return nil, err
	}
	output["LoteSecundario"], err = json.Marshal(r.LoteSecundario)
	if err != nil {
		return nil, err
	}
	output["FechaFabricacion"], err = json.Marshal(r.FechaFabricacion)
	if err != nil {
		return nil, err
	}
	output["FechaVencimiento"], err = json.Marshal(r.FechaVencimiento)
	if err != nil {
		return nil, err
	}
	output["ProductoTrazable"], err = json.Marshal(r.ProductoTrazable)
	if err != nil {
		return nil, err
	}
	output["Almacen"], err = json.Marshal(r.Almacen)
	if err != nil {
		return nil, err
	}
	output["EstadoLote"], err = json.Marshal(r.EstadoLote)
	if err != nil {
		return nil, err
	}
	output["BloqueoUbicacion"], err = json.Marshal(r.BloqueoUbicacion)
	if err != nil {
		return nil, err
	}
	output["VidaUtilLote"], err = json.Marshal(r.VidaUtilLote)
	if err != nil {
		return nil, err
	}
	output["EntregaAntesDe"], err = json.Marshal(r.EntregaAntesDe)
	if err != nil {
		return nil, err
	}
	output["ConsumoAntesDe"], err = json.Marshal(r.ConsumoAntesDe)
	if err != nil {
		return nil, err
	}
	output["EsfuerzoLineaApiPlani"], err = json.Marshal(r.EsfuerzoLineaApiPlani)
	if err != nil {
		return nil, err
	}
	output["VidaUtil"], err = json.Marshal(r.VidaUtil)
	if err != nil {
		return nil, err
	}
	output["Cantidad"], err = json.Marshal(r.Cantidad)
	if err != nil {
		return nil, err
	}
	output["CantidadPickeada"], err = json.Marshal(r.CantidadPickeada)
	if err != nil {
		return nil, err
	}
	output["CantidadExpedida"], err = json.Marshal(r.CantidadExpedida)
	if err != nil {
		return nil, err
	}
	output["EstadoLineaCodigo"], err = json.Marshal(r.EstadoLineaCodigo)
	if err != nil {
		return nil, err
	}
	output["EstadoLineaDescripcion"], err = json.Marshal(r.EstadoLineaDescripcion)
	if err != nil {
		return nil, err
	}
	output["CantidadBultos"], err = json.Marshal(r.CantidadBultos)
	if err != nil {
		return nil, err
	}
	output["Series"], err = json.Marshal(r.Series)
	if err != nil {
		return nil, err
	}
	output["Contrato"], err = json.Marshal(r.Contrato)
	if err != nil {
		return nil, err
	}
	output["Nomenclatura"], err = json.Marshal(r.Nomenclatura)
	if err != nil {
		return nil, err
	}
	output["TipoPedido"], err = json.Marshal(r.TipoPedido)
	if err != nil {
		return nil, err
	}
	output["TipoPedidoDescripcion"], err = json.Marshal(r.TipoPedidoDescripcion)
	if err != nil {
		return nil, err
	}
	output["DetalleAcondicionamiento"], err = json.Marshal(r.DetalleAcondicionamiento)
	if err != nil {
		return nil, err
	}
	output["EstadoOTAcondicionamiento"], err = json.Marshal(r.EstadoOTAcondicionamiento)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Detalle) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["SKU"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SKU); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SKU")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LineaPedidoWH"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LineaPedidoWH); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LineaPedidoWH")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LineaExterna"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LineaExterna); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LineaExterna")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PaqueteLote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PaqueteLote); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PaqueteLote")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LoteCajitaFabricante"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteCajitaFabricante); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LoteCajitaFabricante")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LoteSecundario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteSecundario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LoteSecundario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaFabricacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaFabricacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaFabricacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaVencimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaVencimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaVencimiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ProductoTrazable"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ProductoTrazable); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ProductoTrazable")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Almacen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Almacen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Almacen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EstadoLote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstadoLote); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EstadoLote")
	}
	val = func() json.RawMessage {
		if v, ok := fields["BloqueoUbicacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BloqueoUbicacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for BloqueoUbicacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["VidaUtilLote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.VidaUtilLote); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for VidaUtilLote")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EntregaAntesDe"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EntregaAntesDe); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EntregaAntesDe")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ConsumoAntesDe"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ConsumoAntesDe); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ConsumoAntesDe")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EsfuerzoLineaApiPlani"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EsfuerzoLineaApiPlani); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EsfuerzoLineaApiPlani")
	}
	val = func() json.RawMessage {
		if v, ok := fields["VidaUtil"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.VidaUtil); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for VidaUtil")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Cantidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cantidad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Cantidad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CantidadPickeada"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadPickeada); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CantidadPickeada")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CantidadExpedida"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadExpedida); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CantidadExpedida")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EstadoLineaCodigo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstadoLineaCodigo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EstadoLineaCodigo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EstadoLineaDescripcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstadoLineaDescripcion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EstadoLineaDescripcion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CantidadBultos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadBultos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CantidadBultos")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Series"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Series); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Series")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Contrato"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contrato); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Contrato")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Nomenclatura"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Nomenclatura); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Nomenclatura")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoPedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoPedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoPedido")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoPedidoDescripcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoPedidoDescripcion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoPedidoDescripcion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DetalleAcondicionamiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DetalleAcondicionamiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DetalleAcondicionamiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EstadoOTAcondicionamiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstadoOTAcondicionamiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EstadoOTAcondicionamiento")
	}
	return nil
}
