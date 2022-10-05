// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhArticuloAsnConfirmacion.avsc
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

type Cabecera struct {
	SKU string `json:"SKU"`

	Descripcion string `json:"Descripcion"`

	Propietario string `json:"Propietario"`

	TipoOrigen string `json:"TipoOrigen"`

	CodigoOrigenWH string `json:"CodigoOrigenWH"`

	CodigoOrigenExterno string `json:"CodigoOrigenExterno"`

	ContratoServicioIngreso *UnionNullString `json:"ContratoServicioIngreso"`

	NomenclaturaServicioIngreso *UnionNullString `json:"NomenclaturaServicioIngreso"`

	DescripcionServicioIngreso *UnionNullString `json:"DescripcionServicioIngreso"`

	StockAnteriorSKU float32 `json:"StockAnteriorSKU"`

	StockTotalSKU float32 `json:"StockTotalSKU"`

	StockDisponibleSKU float32 `json:"StockDisponibleSKU"`

	StockEnTransitoSKU float32 `json:"StockEnTransitoSKU"`
}

const CabeceraAvroCRC64Fingerprint = "\xa0\x80^\xa9\xb1\xae!S"

func NewCabecera() Cabecera {
	r := Cabecera{}
	return r
}

func DeserializeCabecera(r io.Reader) (Cabecera, error) {
	t := NewCabecera()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCabeceraFromSchema(r io.Reader, schema string) (Cabecera, error) {
	t := NewCabecera()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCabecera(r Cabecera, w io.Writer) error {
	var err error
	err = vm.WriteString(r.SKU, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Descripcion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TipoOrigen, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoOrigenWH, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoOrigenExterno, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ContratoServicioIngreso, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NomenclaturaServicioIngreso, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DescripcionServicioIngreso, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.StockAnteriorSKU, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.StockTotalSKU, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.StockDisponibleSKU, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.StockEnTransitoSKU, w)
	if err != nil {
		return err
	}
	return err
}

func (r Cabecera) Serialize(w io.Writer) error {
	return writeCabecera(r, w)
}

func (r Cabecera) Schema() string {
	return "{\"fields\":[{\"name\":\"SKU\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"TipoOrigen\",\"type\":\"string\"},{\"name\":\"CodigoOrigenWH\",\"type\":\"string\"},{\"name\":\"CodigoOrigenExterno\",\"type\":\"string\"},{\"name\":\"ContratoServicioIngreso\",\"type\":[\"null\",\"string\"]},{\"name\":\"NomenclaturaServicioIngreso\",\"type\":[\"null\",\"string\"]},{\"name\":\"DescripcionServicioIngreso\",\"type\":[\"null\",\"string\"]},{\"name\":\"StockAnteriorSKU\",\"type\":\"float\"},{\"name\":\"StockTotalSKU\",\"type\":\"float\"},{\"name\":\"StockDisponibleSKU\",\"type\":\"float\"},{\"name\":\"StockEnTransitoSKU\",\"type\":\"float\"}],\"name\":\"Andreani.EventoWhArticulos.Events.AsnConfirmadaCommon.Cabecera\",\"type\":\"record\"}"
}

func (r Cabecera) SchemaName() string {
	return "Andreani.EventoWhArticulos.Events.AsnConfirmadaCommon.Cabecera"
}

func (_ Cabecera) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Cabecera) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Cabecera) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Cabecera) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Cabecera) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Cabecera) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Cabecera) SetString(v string)   { panic("Unsupported operation") }
func (_ Cabecera) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Cabecera) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.SKU}

		return w

	case 1:
		w := types.String{Target: &r.Descripcion}

		return w

	case 2:
		w := types.String{Target: &r.Propietario}

		return w

	case 3:
		w := types.String{Target: &r.TipoOrigen}

		return w

	case 4:
		w := types.String{Target: &r.CodigoOrigenWH}

		return w

	case 5:
		w := types.String{Target: &r.CodigoOrigenExterno}

		return w

	case 6:
		r.ContratoServicioIngreso = NewUnionNullString()

		return r.ContratoServicioIngreso
	case 7:
		r.NomenclaturaServicioIngreso = NewUnionNullString()

		return r.NomenclaturaServicioIngreso
	case 8:
		r.DescripcionServicioIngreso = NewUnionNullString()

		return r.DescripcionServicioIngreso
	case 9:
		w := types.Float{Target: &r.StockAnteriorSKU}

		return w

	case 10:
		w := types.Float{Target: &r.StockTotalSKU}

		return w

	case 11:
		w := types.Float{Target: &r.StockDisponibleSKU}

		return w

	case 12:
		w := types.Float{Target: &r.StockEnTransitoSKU}

		return w

	}
	panic("Unknown field index")
}

func (r *Cabecera) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Cabecera) NullField(i int) {
	switch i {
	case 6:
		r.ContratoServicioIngreso = nil
		return
	case 7:
		r.NomenclaturaServicioIngreso = nil
		return
	case 8:
		r.DescripcionServicioIngreso = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Cabecera) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Cabecera) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Cabecera) HintSize(int)                     { panic("Unsupported operation") }
func (_ Cabecera) Finalize()                        {}

func (_ Cabecera) AvroCRC64Fingerprint() []byte {
	return []byte(CabeceraAvroCRC64Fingerprint)
}

func (r Cabecera) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["SKU"], err = json.Marshal(r.SKU)
	if err != nil {
		return nil, err
	}
	output["Descripcion"], err = json.Marshal(r.Descripcion)
	if err != nil {
		return nil, err
	}
	output["Propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["TipoOrigen"], err = json.Marshal(r.TipoOrigen)
	if err != nil {
		return nil, err
	}
	output["CodigoOrigenWH"], err = json.Marshal(r.CodigoOrigenWH)
	if err != nil {
		return nil, err
	}
	output["CodigoOrigenExterno"], err = json.Marshal(r.CodigoOrigenExterno)
	if err != nil {
		return nil, err
	}
	output["ContratoServicioIngreso"], err = json.Marshal(r.ContratoServicioIngreso)
	if err != nil {
		return nil, err
	}
	output["NomenclaturaServicioIngreso"], err = json.Marshal(r.NomenclaturaServicioIngreso)
	if err != nil {
		return nil, err
	}
	output["DescripcionServicioIngreso"], err = json.Marshal(r.DescripcionServicioIngreso)
	if err != nil {
		return nil, err
	}
	output["StockAnteriorSKU"], err = json.Marshal(r.StockAnteriorSKU)
	if err != nil {
		return nil, err
	}
	output["StockTotalSKU"], err = json.Marshal(r.StockTotalSKU)
	if err != nil {
		return nil, err
	}
	output["StockDisponibleSKU"], err = json.Marshal(r.StockDisponibleSKU)
	if err != nil {
		return nil, err
	}
	output["StockEnTransitoSKU"], err = json.Marshal(r.StockEnTransitoSKU)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Cabecera) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["Descripcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Descripcion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Descripcion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Propietario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Propietario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Propietario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoOrigen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoOrigen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoOrigenWH"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoOrigenWH); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodigoOrigenWH")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoOrigenExterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoOrigenExterno); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodigoOrigenExterno")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ContratoServicioIngreso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ContratoServicioIngreso); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ContratoServicioIngreso")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NomenclaturaServicioIngreso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NomenclaturaServicioIngreso); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NomenclaturaServicioIngreso")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DescripcionServicioIngreso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DescripcionServicioIngreso); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DescripcionServicioIngreso")
	}
	val = func() json.RawMessage {
		if v, ok := fields["StockAnteriorSKU"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StockAnteriorSKU); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for StockAnteriorSKU")
	}
	val = func() json.RawMessage {
		if v, ok := fields["StockTotalSKU"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StockTotalSKU); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for StockTotalSKU")
	}
	val = func() json.RawMessage {
		if v, ok := fields["StockDisponibleSKU"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StockDisponibleSKU); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for StockDisponibleSKU")
	}
	val = func() json.RawMessage {
		if v, ok := fields["StockEnTransitoSKU"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StockEnTransitoSKU); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for StockEnTransitoSKU")
	}
	return nil
}