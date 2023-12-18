// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ETA.avsc
 */
package DeliveryEstimateEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type CalculoEta struct {
	OrdenDeEnvioEnHR *UnionNullInt `json:"OrdenDeEnvioEnHR"`

	NumeroHojaDeRuta *UnionNullString `json:"NumeroHojaDeRuta"`

	Geocoordenadas *UnionNullString `json:"Geocoordenadas"`

	RecorridoEnSegundos *UnionNullDouble `json:"RecorridoEnSegundos"`

	RecorridoEnMetros *UnionNullDouble `json:"RecorridoEnMetros"`

	DemoraEnDomicilioEnMinutos *UnionNullInt `json:"DemoraEnDomicilioEnMinutos"`

	DemoraSalidaSucursalEnMinutos *UnionNullInt `json:"DemoraSalidaSucursalEnMinutos"`

	EtaAnterior *UnionNullLong `json:"EtaAnterior"`

	FechaCreacionHojaDeRuta *UnionNullLong `json:"FechaCreacionHojaDeRuta"`
}

const CalculoEtaAvroCRC64Fingerprint = "Ȥ\xd8\aC\x91n\x9d"

func NewCalculoEta() CalculoEta {
	r := CalculoEta{}
	r.OrdenDeEnvioEnHR = nil
	r.NumeroHojaDeRuta = nil
	r.Geocoordenadas = nil
	r.RecorridoEnSegundos = nil
	r.RecorridoEnMetros = nil
	r.DemoraEnDomicilioEnMinutos = nil
	r.DemoraSalidaSucursalEnMinutos = nil
	r.EtaAnterior = nil
	r.FechaCreacionHojaDeRuta = nil
	return r
}

func DeserializeCalculoEta(r io.Reader) (CalculoEta, error) {
	t := NewCalculoEta()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCalculoEtaFromSchema(r io.Reader, schema string) (CalculoEta, error) {
	t := NewCalculoEta()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCalculoEta(r CalculoEta, w io.Writer) error {
	var err error
	err = writeUnionNullInt(r.OrdenDeEnvioEnHR, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroHojaDeRuta, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Geocoordenadas, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDouble(r.RecorridoEnSegundos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDouble(r.RecorridoEnMetros, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.DemoraEnDomicilioEnMinutos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.DemoraSalidaSucursalEnMinutos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.EtaAnterior, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaCreacionHojaDeRuta, w)
	if err != nil {
		return err
	}
	return err
}

func (r CalculoEta) Serialize(w io.Writer) error {
	return writeCalculoEta(r, w)
}

func (r CalculoEta) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"OrdenDeEnvioEnHR\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"NumeroHojaDeRuta\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Geocoordenadas\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"RecorridoEnSegundos\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"RecorridoEnMetros\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"DemoraEnDomicilioEnMinutos\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"DemoraSalidaSucursalEnMinutos\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"EtaAnterior\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"FechaCreacionHojaDeRuta\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"Andreani.DeliveryEstimate.Events.Records.CalculoEta\",\"type\":\"record\"}"
}

func (r CalculoEta) SchemaName() string {
	return "Andreani.DeliveryEstimate.Events.Records.CalculoEta"
}

func (_ CalculoEta) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CalculoEta) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CalculoEta) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CalculoEta) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CalculoEta) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CalculoEta) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CalculoEta) SetString(v string)   { panic("Unsupported operation") }
func (_ CalculoEta) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CalculoEta) Get(i int) types.Field {
	switch i {
	case 0:
		r.OrdenDeEnvioEnHR = NewUnionNullInt()

		return r.OrdenDeEnvioEnHR
	case 1:
		r.NumeroHojaDeRuta = NewUnionNullString()

		return r.NumeroHojaDeRuta
	case 2:
		r.Geocoordenadas = NewUnionNullString()

		return r.Geocoordenadas
	case 3:
		r.RecorridoEnSegundos = NewUnionNullDouble()

		return r.RecorridoEnSegundos
	case 4:
		r.RecorridoEnMetros = NewUnionNullDouble()

		return r.RecorridoEnMetros
	case 5:
		r.DemoraEnDomicilioEnMinutos = NewUnionNullInt()

		return r.DemoraEnDomicilioEnMinutos
	case 6:
		r.DemoraSalidaSucursalEnMinutos = NewUnionNullInt()

		return r.DemoraSalidaSucursalEnMinutos
	case 7:
		r.EtaAnterior = NewUnionNullLong()

		return r.EtaAnterior
	case 8:
		r.FechaCreacionHojaDeRuta = NewUnionNullLong()

		return r.FechaCreacionHojaDeRuta
	}
	panic("Unknown field index")
}

func (r *CalculoEta) SetDefault(i int) {
	switch i {
	case 0:
		r.OrdenDeEnvioEnHR = nil
		return
	case 1:
		r.NumeroHojaDeRuta = nil
		return
	case 2:
		r.Geocoordenadas = nil
		return
	case 3:
		r.RecorridoEnSegundos = nil
		return
	case 4:
		r.RecorridoEnMetros = nil
		return
	case 5:
		r.DemoraEnDomicilioEnMinutos = nil
		return
	case 6:
		r.DemoraSalidaSucursalEnMinutos = nil
		return
	case 7:
		r.EtaAnterior = nil
		return
	case 8:
		r.FechaCreacionHojaDeRuta = nil
		return
	}
	panic("Unknown field index")
}

func (r *CalculoEta) NullField(i int) {
	switch i {
	case 0:
		r.OrdenDeEnvioEnHR = nil
		return
	case 1:
		r.NumeroHojaDeRuta = nil
		return
	case 2:
		r.Geocoordenadas = nil
		return
	case 3:
		r.RecorridoEnSegundos = nil
		return
	case 4:
		r.RecorridoEnMetros = nil
		return
	case 5:
		r.DemoraEnDomicilioEnMinutos = nil
		return
	case 6:
		r.DemoraSalidaSucursalEnMinutos = nil
		return
	case 7:
		r.EtaAnterior = nil
		return
	case 8:
		r.FechaCreacionHojaDeRuta = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ CalculoEta) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CalculoEta) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CalculoEta) HintSize(int)                     { panic("Unsupported operation") }
func (_ CalculoEta) Finalize()                        {}

func (_ CalculoEta) AvroCRC64Fingerprint() []byte {
	return []byte(CalculoEtaAvroCRC64Fingerprint)
}

func (r CalculoEta) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["OrdenDeEnvioEnHR"], err = json.Marshal(r.OrdenDeEnvioEnHR)
	if err != nil {
		return nil, err
	}
	output["NumeroHojaDeRuta"], err = json.Marshal(r.NumeroHojaDeRuta)
	if err != nil {
		return nil, err
	}
	output["Geocoordenadas"], err = json.Marshal(r.Geocoordenadas)
	if err != nil {
		return nil, err
	}
	output["RecorridoEnSegundos"], err = json.Marshal(r.RecorridoEnSegundos)
	if err != nil {
		return nil, err
	}
	output["RecorridoEnMetros"], err = json.Marshal(r.RecorridoEnMetros)
	if err != nil {
		return nil, err
	}
	output["DemoraEnDomicilioEnMinutos"], err = json.Marshal(r.DemoraEnDomicilioEnMinutos)
	if err != nil {
		return nil, err
	}
	output["DemoraSalidaSucursalEnMinutos"], err = json.Marshal(r.DemoraSalidaSucursalEnMinutos)
	if err != nil {
		return nil, err
	}
	output["EtaAnterior"], err = json.Marshal(r.EtaAnterior)
	if err != nil {
		return nil, err
	}
	output["FechaCreacionHojaDeRuta"], err = json.Marshal(r.FechaCreacionHojaDeRuta)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CalculoEta) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["OrdenDeEnvioEnHR"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenDeEnvioEnHR); err != nil {
			return err
		}
	} else {
		r.OrdenDeEnvioEnHR = NewUnionNullInt()

		r.OrdenDeEnvioEnHR = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroHojaDeRuta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroHojaDeRuta); err != nil {
			return err
		}
	} else {
		r.NumeroHojaDeRuta = NewUnionNullString()

		r.NumeroHojaDeRuta = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Geocoordenadas"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Geocoordenadas); err != nil {
			return err
		}
	} else {
		r.Geocoordenadas = NewUnionNullString()

		r.Geocoordenadas = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["RecorridoEnSegundos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RecorridoEnSegundos); err != nil {
			return err
		}
	} else {
		r.RecorridoEnSegundos = NewUnionNullDouble()

		r.RecorridoEnSegundos = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["RecorridoEnMetros"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RecorridoEnMetros); err != nil {
			return err
		}
	} else {
		r.RecorridoEnMetros = NewUnionNullDouble()

		r.RecorridoEnMetros = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DemoraEnDomicilioEnMinutos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DemoraEnDomicilioEnMinutos); err != nil {
			return err
		}
	} else {
		r.DemoraEnDomicilioEnMinutos = NewUnionNullInt()

		r.DemoraEnDomicilioEnMinutos = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DemoraSalidaSucursalEnMinutos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DemoraSalidaSucursalEnMinutos); err != nil {
			return err
		}
	} else {
		r.DemoraSalidaSucursalEnMinutos = NewUnionNullInt()

		r.DemoraSalidaSucursalEnMinutos = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["EtaAnterior"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EtaAnterior); err != nil {
			return err
		}
	} else {
		r.EtaAnterior = NewUnionNullLong()

		r.EtaAnterior = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaCreacionHojaDeRuta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaCreacionHojaDeRuta); err != nil {
			return err
		}
	} else {
		r.FechaCreacionHojaDeRuta = NewUnionNullLong()

		r.FechaCreacionHojaDeRuta = nil
	}
	return nil
}
