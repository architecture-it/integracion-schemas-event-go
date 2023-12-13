// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Preasignacion.avsc
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

type Preasignacion struct {
	PaqueteLote string `json:"PaqueteLote"`

	LoteCajitaFabricante string `json:"LoteCajitaFabricante"`

	LoteSecundario string `json:"LoteSecundario"`

	FechaFabricacion *UnionNullLong `json:"FechaFabricacion"`

	FechaVencimiento *UnionNullLong `json:"FechaVencimiento"`

	ProductoTrazable string `json:"ProductoTrazable"`

	AlmacenConsumo string `json:"AlmacenConsumo"`

	EstadoLote string `json:"EstadoLote"`

	BloqueoUbicacion string `json:"BloqueoUbicacion"`

	VitaUtilLote string `json:"VitaUtilLote"`

	LoteSce string `json:"LoteSce"`

	CantidadPreasignada float32 `json:"CantidadPreasignada"`

	FechaPreasignacion *UnionNullLong `json:"FechaPreasignacion"`
}

const PreasignacionAvroCRC64Fingerprint = "\x85\xd3\xe6|\xf0ͳ\xfa"

func NewPreasignacion() Preasignacion {
	r := Preasignacion{}
	return r
}

func DeserializePreasignacion(r io.Reader) (Preasignacion, error) {
	t := NewPreasignacion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePreasignacionFromSchema(r io.Reader, schema string) (Preasignacion, error) {
	t := NewPreasignacion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePreasignacion(r Preasignacion, w io.Writer) error {
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
	err = vm.WriteString(r.VitaUtilLote, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.LoteSce, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.CantidadPreasignada, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaPreasignacion, w)
	if err != nil {
		return err
	}
	return err
}

func (r Preasignacion) Serialize(w io.Writer) error {
	return writePreasignacion(r, w)
}

func (r Preasignacion) Schema() string {
	return "{\"fields\":[{\"name\":\"PaqueteLote\",\"type\":\"string\"},{\"name\":\"LoteCajitaFabricante\",\"type\":\"string\"},{\"name\":\"LoteSecundario\",\"type\":\"string\"},{\"name\":\"FechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ProductoTrazable\",\"type\":\"string\"},{\"name\":\"AlmacenConsumo\",\"type\":\"string\"},{\"name\":\"EstadoLote\",\"type\":\"string\"},{\"name\":\"BloqueoUbicacion\",\"type\":\"string\"},{\"name\":\"VitaUtilLote\",\"type\":\"string\"},{\"name\":\"LoteSce\",\"type\":\"string\"},{\"name\":\"CantidadPreasignada\",\"type\":\"float\"},{\"name\":\"FechaPreasignacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"Andreani.EventoWhPedidos.Events.PreasignadoCommon.Preasignacion\",\"type\":\"record\"}"
}

func (r Preasignacion) SchemaName() string {
	return "Andreani.EventoWhPedidos.Events.PreasignadoCommon.Preasignacion"
}

func (_ Preasignacion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Preasignacion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Preasignacion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Preasignacion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Preasignacion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Preasignacion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Preasignacion) SetString(v string)   { panic("Unsupported operation") }
func (_ Preasignacion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Preasignacion) Get(i int) types.Field {
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
		w := types.String{Target: &r.VitaUtilLote}

		return w

	case 10:
		w := types.String{Target: &r.LoteSce}

		return w

	case 11:
		w := types.Float{Target: &r.CantidadPreasignada}

		return w

	case 12:
		r.FechaPreasignacion = NewUnionNullLong()

		return r.FechaPreasignacion
	}
	panic("Unknown field index")
}

func (r *Preasignacion) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Preasignacion) NullField(i int) {
	switch i {
	case 3:
		r.FechaFabricacion = nil
		return
	case 4:
		r.FechaVencimiento = nil
		return
	case 12:
		r.FechaPreasignacion = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Preasignacion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Preasignacion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Preasignacion) HintSize(int)                     { panic("Unsupported operation") }
func (_ Preasignacion) Finalize()                        {}

func (_ Preasignacion) AvroCRC64Fingerprint() []byte {
	return []byte(PreasignacionAvroCRC64Fingerprint)
}

func (r Preasignacion) MarshalJSON() ([]byte, error) {
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
	output["VitaUtilLote"], err = json.Marshal(r.VitaUtilLote)
	if err != nil {
		return nil, err
	}
	output["LoteSce"], err = json.Marshal(r.LoteSce)
	if err != nil {
		return nil, err
	}
	output["CantidadPreasignada"], err = json.Marshal(r.CantidadPreasignada)
	if err != nil {
		return nil, err
	}
	output["FechaPreasignacion"], err = json.Marshal(r.FechaPreasignacion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Preasignacion) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["VitaUtilLote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.VitaUtilLote); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for VitaUtilLote")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LoteSce"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteSce); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LoteSce")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CantidadPreasignada"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadPreasignada); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CantidadPreasignada")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaPreasignacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaPreasignacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaPreasignacion")
	}
	return nil
}
