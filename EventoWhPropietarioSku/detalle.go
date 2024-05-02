// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhPropietarioSkuCompletado.avsc
 */
package EventoWhPropietarioSkuEvents

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
	NumeroLineaWH string `json:"NumeroLineaWH"`

	PropietarioOrigen string `json:"PropietarioOrigen"`

	PropietarioDestino string `json:"PropietarioDestino"`

	ArticuloOrigen string `json:"ArticuloOrigen"`

	ArticuloDestino string `json:"ArticuloDestino"`

	LoteWhOrigen *UnionNullString `json:"LoteWhOrigen"`

	LoteWhDestino *UnionNullString `json:"LoteWhDestino"`

	LpnOrigen *UnionNullString `json:"LpnOrigen"`

	LpnDestino *UnionNullString `json:"LpnDestino"`

	CantidadOrigen float32 `json:"CantidadOrigen"`

	CantidadDestino float32 `json:"CantidadDestino"`

	PaqueteLote *UnionNullString `json:"PaqueteLote"`

	LoteCajitaFabricante *UnionNullString `json:"LoteCajitaFabricante"`

	LoteSecundario *UnionNullString `json:"LoteSecundario"`

	FechaFabricacion *UnionNullLong `json:"FechaFabricacion"`

	FechaVencimiento *UnionNullLong `json:"FechaVencimiento"`

	ProductoTrazable *UnionNullString `json:"ProductoTrazable"`

	AlmacenConsumo *UnionNullString `json:"AlmacenConsumo"`

	EstadoLote *UnionNullString `json:"EstadoLote"`

	BloqueoUbicacion *UnionNullString `json:"BloqueoUbicacion"`

	VidaUtilLote *UnionNullString `json:"VidaUtilLote"`

	EntregaAntesDe *UnionNullLong `json:"EntregaAntesDe"`

	ConsumoAntesDe *UnionNullLong `json:"ConsumoAntesDe"`

	Estado *UnionNullString `json:"Estado"`

	FechaFinalizacion *UnionNullLong `json:"FechaFinalizacion"`
}

const DetalleAvroCRC64Fingerprint = ">\xb8/\xfe\xf5%\xca."

func NewDetalle() Detalle {
	r := Detalle{}
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
	err = vm.WriteString(r.NumeroLineaWH, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.PropietarioOrigen, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.PropietarioDestino, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.ArticuloOrigen, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.ArticuloDestino, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LoteWhOrigen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LoteWhDestino, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LpnOrigen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LpnDestino, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.CantidadOrigen, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.CantidadDestino, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.PaqueteLote, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LoteCajitaFabricante, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LoteSecundario, w)
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
	err = writeUnionNullString(r.ProductoTrazable, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.AlmacenConsumo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstadoLote, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.BloqueoUbicacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.VidaUtilLote, w)
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
	err = writeUnionNullString(r.Estado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaFinalizacion, w)
	if err != nil {
		return err
	}
	return err
}

func (r Detalle) Serialize(w io.Writer) error {
	return writeDetalle(r, w)
}

func (r Detalle) Schema() string {
	return "{\"fields\":[{\"name\":\"NumeroLineaWH\",\"type\":\"string\"},{\"name\":\"PropietarioOrigen\",\"type\":\"string\"},{\"name\":\"PropietarioDestino\",\"type\":\"string\"},{\"name\":\"ArticuloOrigen\",\"type\":\"string\"},{\"name\":\"ArticuloDestino\",\"type\":\"string\"},{\"name\":\"LoteWhOrigen\",\"type\":[\"null\",\"string\"]},{\"name\":\"LoteWhDestino\",\"type\":[\"null\",\"string\"]},{\"name\":\"LpnOrigen\",\"type\":[\"null\",\"string\"]},{\"name\":\"LpnDestino\",\"type\":[\"null\",\"string\"]},{\"name\":\"CantidadOrigen\",\"type\":\"float\"},{\"name\":\"CantidadDestino\",\"type\":\"float\"},{\"name\":\"PaqueteLote\",\"type\":[\"null\",\"string\"]},{\"name\":\"LoteCajitaFabricante\",\"type\":[\"null\",\"string\"]},{\"name\":\"LoteSecundario\",\"type\":[\"null\",\"string\"]},{\"name\":\"FechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ProductoTrazable\",\"type\":[\"null\",\"string\"]},{\"name\":\"AlmacenConsumo\",\"type\":[\"null\",\"string\"]},{\"name\":\"EstadoLote\",\"type\":[\"null\",\"string\"]},{\"name\":\"BloqueoUbicacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"VidaUtilLote\",\"type\":[\"null\",\"string\"]},{\"name\":\"EntregaAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ConsumoAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"Estado\",\"type\":[\"null\",\"string\"]},{\"name\":\"FechaFinalizacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"Andreani.EventoWhPropietarioSku.Events.CompletadoCommon.Detalle\",\"type\":\"record\"}"
}

func (r Detalle) SchemaName() string {
	return "Andreani.EventoWhPropietarioSku.Events.CompletadoCommon.Detalle"
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
		w := types.String{Target: &r.NumeroLineaWH}

		return w

	case 1:
		w := types.String{Target: &r.PropietarioOrigen}

		return w

	case 2:
		w := types.String{Target: &r.PropietarioDestino}

		return w

	case 3:
		w := types.String{Target: &r.ArticuloOrigen}

		return w

	case 4:
		w := types.String{Target: &r.ArticuloDestino}

		return w

	case 5:
		r.LoteWhOrigen = NewUnionNullString()

		return r.LoteWhOrigen
	case 6:
		r.LoteWhDestino = NewUnionNullString()

		return r.LoteWhDestino
	case 7:
		r.LpnOrigen = NewUnionNullString()

		return r.LpnOrigen
	case 8:
		r.LpnDestino = NewUnionNullString()

		return r.LpnDestino
	case 9:
		w := types.Float{Target: &r.CantidadOrigen}

		return w

	case 10:
		w := types.Float{Target: &r.CantidadDestino}

		return w

	case 11:
		r.PaqueteLote = NewUnionNullString()

		return r.PaqueteLote
	case 12:
		r.LoteCajitaFabricante = NewUnionNullString()

		return r.LoteCajitaFabricante
	case 13:
		r.LoteSecundario = NewUnionNullString()

		return r.LoteSecundario
	case 14:
		r.FechaFabricacion = NewUnionNullLong()

		return r.FechaFabricacion
	case 15:
		r.FechaVencimiento = NewUnionNullLong()

		return r.FechaVencimiento
	case 16:
		r.ProductoTrazable = NewUnionNullString()

		return r.ProductoTrazable
	case 17:
		r.AlmacenConsumo = NewUnionNullString()

		return r.AlmacenConsumo
	case 18:
		r.EstadoLote = NewUnionNullString()

		return r.EstadoLote
	case 19:
		r.BloqueoUbicacion = NewUnionNullString()

		return r.BloqueoUbicacion
	case 20:
		r.VidaUtilLote = NewUnionNullString()

		return r.VidaUtilLote
	case 21:
		r.EntregaAntesDe = NewUnionNullLong()

		return r.EntregaAntesDe
	case 22:
		r.ConsumoAntesDe = NewUnionNullLong()

		return r.ConsumoAntesDe
	case 23:
		r.Estado = NewUnionNullString()

		return r.Estado
	case 24:
		r.FechaFinalizacion = NewUnionNullLong()

		return r.FechaFinalizacion
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
	case 5:
		r.LoteWhOrigen = nil
		return
	case 6:
		r.LoteWhDestino = nil
		return
	case 7:
		r.LpnOrigen = nil
		return
	case 8:
		r.LpnDestino = nil
		return
	case 11:
		r.PaqueteLote = nil
		return
	case 12:
		r.LoteCajitaFabricante = nil
		return
	case 13:
		r.LoteSecundario = nil
		return
	case 14:
		r.FechaFabricacion = nil
		return
	case 15:
		r.FechaVencimiento = nil
		return
	case 16:
		r.ProductoTrazable = nil
		return
	case 17:
		r.AlmacenConsumo = nil
		return
	case 18:
		r.EstadoLote = nil
		return
	case 19:
		r.BloqueoUbicacion = nil
		return
	case 20:
		r.VidaUtilLote = nil
		return
	case 21:
		r.EntregaAntesDe = nil
		return
	case 22:
		r.ConsumoAntesDe = nil
		return
	case 23:
		r.Estado = nil
		return
	case 24:
		r.FechaFinalizacion = nil
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
	output["NumeroLineaWH"], err = json.Marshal(r.NumeroLineaWH)
	if err != nil {
		return nil, err
	}
	output["PropietarioOrigen"], err = json.Marshal(r.PropietarioOrigen)
	if err != nil {
		return nil, err
	}
	output["PropietarioDestino"], err = json.Marshal(r.PropietarioDestino)
	if err != nil {
		return nil, err
	}
	output["ArticuloOrigen"], err = json.Marshal(r.ArticuloOrigen)
	if err != nil {
		return nil, err
	}
	output["ArticuloDestino"], err = json.Marshal(r.ArticuloDestino)
	if err != nil {
		return nil, err
	}
	output["LoteWhOrigen"], err = json.Marshal(r.LoteWhOrigen)
	if err != nil {
		return nil, err
	}
	output["LoteWhDestino"], err = json.Marshal(r.LoteWhDestino)
	if err != nil {
		return nil, err
	}
	output["LpnOrigen"], err = json.Marshal(r.LpnOrigen)
	if err != nil {
		return nil, err
	}
	output["LpnDestino"], err = json.Marshal(r.LpnDestino)
	if err != nil {
		return nil, err
	}
	output["CantidadOrigen"], err = json.Marshal(r.CantidadOrigen)
	if err != nil {
		return nil, err
	}
	output["CantidadDestino"], err = json.Marshal(r.CantidadDestino)
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
	output["AlmacenConsumo"], err = json.Marshal(r.AlmacenConsumo)
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
	output["Estado"], err = json.Marshal(r.Estado)
	if err != nil {
		return nil, err
	}
	output["FechaFinalizacion"], err = json.Marshal(r.FechaFinalizacion)
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
		if v, ok := fields["NumeroLineaWH"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroLineaWH); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NumeroLineaWH")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PropietarioOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PropietarioOrigen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PropietarioOrigen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PropietarioDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PropietarioDestino); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PropietarioDestino")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ArticuloOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ArticuloOrigen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ArticuloOrigen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ArticuloDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ArticuloDestino); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ArticuloDestino")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LoteWhOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteWhOrigen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LoteWhOrigen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LoteWhDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteWhDestino); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LoteWhDestino")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LpnOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LpnOrigen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LpnOrigen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LpnDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LpnDestino); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LpnDestino")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CantidadOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadOrigen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CantidadOrigen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CantidadDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadDestino); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CantidadDestino")
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
		if v, ok := fields["AlmacenConsumo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AlmacenConsumo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for AlmacenConsumo")
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
		if v, ok := fields["FechaFinalizacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaFinalizacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaFinalizacion")
	}
	return nil
}
