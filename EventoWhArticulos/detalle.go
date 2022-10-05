// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhArticuloExpedicion.avsc
 */
package EventoWhArticulosEvents

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
	PaqueteLote string `json:"PaqueteLote"`

	LoteCajitaFabricante string `json:"LoteCajitaFabricante"`

	LoteSecundario string `json:"LoteSecundario"`

	FechaFabricacion *UnionNullLong `json:"FechaFabricacion"`

	FechaVencimiento *UnionNullLong `json:"FechaVencimiento"`

	ProductoTrazable string `json:"ProductoTrazable"`

	AlmacenConsumo string `json:"AlmacenConsumo"`

	EstadoLote string `json:"EstadoLote"`

	BloqueoUbicacion string `json:"BloqueoUbicacion"`

	VidaUtilLote string `json:"VidaUtilLote"`

	Lpn *UnionNullString `json:"Lpn"`

	EntregaAntesDe *UnionNullLong `json:"EntregaAntesDe"`

	ConsumoAntesDe *UnionNullLong `json:"ConsumoAntesDe"`

	StockDisponible float32 `json:"StockDisponible"`

	StockEnTransito float32 `json:"StockEnTransito"`

	StockAnterior float32 `json:"StockAnterior"`

	StockTotal float32 `json:"StockTotal"`
}

const DetalleAvroCRC64Fingerprint = "\u007fF\x16\xabD\xbf\x16i"

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
	err = vm.WriteString(r.AlmacenConsumo, w)
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
	err = writeUnionNullString(r.Lpn, w)
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
	err = vm.WriteFloat(r.StockDisponible, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.StockEnTransito, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.StockAnterior, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.StockTotal, w)
	if err != nil {
		return err
	}
	return err
}

func (r Detalle) Serialize(w io.Writer) error {
	return writeDetalle(r, w)
}

func (r Detalle) Schema() string {
	return "{\"fields\":[{\"name\":\"PaqueteLote\",\"type\":\"string\"},{\"name\":\"LoteCajitaFabricante\",\"type\":\"string\"},{\"name\":\"LoteSecundario\",\"type\":\"string\"},{\"name\":\"FechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ProductoTrazable\",\"type\":\"string\"},{\"name\":\"AlmacenConsumo\",\"type\":\"string\"},{\"name\":\"EstadoLote\",\"type\":\"string\"},{\"name\":\"BloqueoUbicacion\",\"type\":\"string\"},{\"name\":\"VidaUtilLote\",\"type\":\"string\"},{\"name\":\"Lpn\",\"type\":[\"null\",\"string\"]},{\"name\":\"EntregaAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ConsumoAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"StockDisponible\",\"type\":\"float\"},{\"name\":\"StockEnTransito\",\"type\":\"float\"},{\"name\":\"StockAnterior\",\"type\":\"float\"},{\"name\":\"StockTotal\",\"type\":\"float\"}],\"name\":\"Andreani.EventoWhArticulos.Events.ArticuloExpedicionCommon.Detalle\",\"type\":\"record\"}"
}

func (r Detalle) SchemaName() string {
	return "Andreani.EventoWhArticulos.Events.ArticuloExpedicionCommon.Detalle"
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
		w := types.String{Target: &r.PaqueteLote}

		return w

	case 1:
		w := types.String{Target: &r.LoteCajitaFabricante}

		return w

	case 2:
		w := types.String{Target: &r.LoteSecundario}

		return w

	case 3:
		r.FechaFabricacion = NewUnionNullLong()

		return r.FechaFabricacion
	case 4:
		r.FechaVencimiento = NewUnionNullLong()

		return r.FechaVencimiento
	case 5:
		w := types.String{Target: &r.ProductoTrazable}

		return w

	case 6:
		w := types.String{Target: &r.AlmacenConsumo}

		return w

	case 7:
		w := types.String{Target: &r.EstadoLote}

		return w

	case 8:
		w := types.String{Target: &r.BloqueoUbicacion}

		return w

	case 9:
		w := types.String{Target: &r.VidaUtilLote}

		return w

	case 10:
		r.Lpn = NewUnionNullString()

		return r.Lpn
	case 11:
		r.EntregaAntesDe = NewUnionNullLong()

		return r.EntregaAntesDe
	case 12:
		r.ConsumoAntesDe = NewUnionNullLong()

		return r.ConsumoAntesDe
	case 13:
		w := types.Float{Target: &r.StockDisponible}

		return w

	case 14:
		w := types.Float{Target: &r.StockEnTransito}

		return w

	case 15:
		w := types.Float{Target: &r.StockAnterior}

		return w

	case 16:
		w := types.Float{Target: &r.StockTotal}

		return w

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
	case 3:
		r.FechaFabricacion = nil
		return
	case 4:
		r.FechaVencimiento = nil
		return
	case 10:
		r.Lpn = nil
		return
	case 11:
		r.EntregaAntesDe = nil
		return
	case 12:
		r.ConsumoAntesDe = nil
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
	output["Lpn"], err = json.Marshal(r.Lpn)
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
	output["StockDisponible"], err = json.Marshal(r.StockDisponible)
	if err != nil {
		return nil, err
	}
	output["StockEnTransito"], err = json.Marshal(r.StockEnTransito)
	if err != nil {
		return nil, err
	}
	output["StockAnterior"], err = json.Marshal(r.StockAnterior)
	if err != nil {
		return nil, err
	}
	output["StockTotal"], err = json.Marshal(r.StockTotal)
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
		if v, ok := fields["Lpn"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Lpn); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Lpn")
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
		if v, ok := fields["StockDisponible"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StockDisponible); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for StockDisponible")
	}
	val = func() json.RawMessage {
		if v, ok := fields["StockEnTransito"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StockEnTransito); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for StockEnTransito")
	}
	val = func() json.RawMessage {
		if v, ok := fields["StockAnterior"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StockAnterior); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for StockAnterior")
	}
	val = func() json.RawMessage {
		if v, ok := fields["StockTotal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StockTotal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for StockTotal")
	}
	return nil
}
