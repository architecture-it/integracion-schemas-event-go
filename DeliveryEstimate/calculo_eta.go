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
	OrdenDeEnvioEnHR int32 `json:"OrdenDeEnvioEnHR"`

	NumeroHojaDeRuta string `json:"NumeroHojaDeRuta"`

	Geocoordenadas string `json:"Geocoordenadas"`

	RecorridoEnSegundos float64 `json:"RecorridoEnSegundos"`

	RecorridoEnMetros float64 `json:"RecorridoEnMetros"`

	DemoraEnDomicilioEnMinutos int32 `json:"DemoraEnDomicilioEnMinutos"`

	DemoraSalidaSucursalEnMinutos int32 `json:"DemoraSalidaSucursalEnMinutos"`

	EtaAnterior *UnionNullLong `json:"EtaAnterior"`

	FechaCreacionHojaDeRuta int64 `json:"FechaCreacionHojaDeRuta"`
}

const CalculoEtaAvroCRC64Fingerprint = "зX\xe8\\c\xd7\x03"

func NewCalculoEta() CalculoEta {
	r := CalculoEta{}
	r.EtaAnterior = nil
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
	err = vm.WriteInt(r.OrdenDeEnvioEnHR, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroHojaDeRuta, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Geocoordenadas, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.RecorridoEnSegundos, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.RecorridoEnMetros, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.DemoraEnDomicilioEnMinutos, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.DemoraSalidaSucursalEnMinutos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.EtaAnterior, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.FechaCreacionHojaDeRuta, w)
	if err != nil {
		return err
	}
	return err
}

func (r CalculoEta) Serialize(w io.Writer) error {
	return writeCalculoEta(r, w)
}

func (r CalculoEta) Schema() string {
	return "{\"fields\":[{\"name\":\"OrdenDeEnvioEnHR\",\"type\":\"int\"},{\"name\":\"NumeroHojaDeRuta\",\"type\":\"string\"},{\"name\":\"Geocoordenadas\",\"type\":\"string\"},{\"name\":\"RecorridoEnSegundos\",\"type\":\"double\"},{\"name\":\"RecorridoEnMetros\",\"type\":\"double\"},{\"name\":\"DemoraEnDomicilioEnMinutos\",\"type\":\"int\"},{\"name\":\"DemoraSalidaSucursalEnMinutos\",\"type\":\"int\"},{\"default\":null,\"name\":\"EtaAnterior\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaCreacionHojaDeRuta\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}}],\"name\":\"Andreani.DeliveryEstimate.Events.Records.CalculoEta\",\"type\":\"record\"}"
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
		w := types.Int{Target: &r.OrdenDeEnvioEnHR}

		return w

	case 1:
		w := types.String{Target: &r.NumeroHojaDeRuta}

		return w

	case 2:
		w := types.String{Target: &r.Geocoordenadas}

		return w

	case 3:
		w := types.Double{Target: &r.RecorridoEnSegundos}

		return w

	case 4:
		w := types.Double{Target: &r.RecorridoEnMetros}

		return w

	case 5:
		w := types.Int{Target: &r.DemoraEnDomicilioEnMinutos}

		return w

	case 6:
		w := types.Int{Target: &r.DemoraSalidaSucursalEnMinutos}

		return w

	case 7:
		r.EtaAnterior = NewUnionNullLong()

		return r.EtaAnterior
	case 8:
		w := types.Long{Target: &r.FechaCreacionHojaDeRuta}

		return w

	}
	panic("Unknown field index")
}

func (r *CalculoEta) SetDefault(i int) {
	switch i {
	case 7:
		r.EtaAnterior = nil
		return
	}
	panic("Unknown field index")
}

func (r *CalculoEta) NullField(i int) {
	switch i {
	case 7:
		r.EtaAnterior = nil
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
		return fmt.Errorf("no value specified for OrdenDeEnvioEnHR")
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
		return fmt.Errorf("no value specified for NumeroHojaDeRuta")
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
		return fmt.Errorf("no value specified for Geocoordenadas")
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
		return fmt.Errorf("no value specified for RecorridoEnSegundos")
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
		return fmt.Errorf("no value specified for RecorridoEnMetros")
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
		return fmt.Errorf("no value specified for DemoraEnDomicilioEnMinutos")
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
		return fmt.Errorf("no value specified for DemoraSalidaSucursalEnMinutos")
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
		return fmt.Errorf("no value specified for FechaCreacionHojaDeRuta")
	}
	return nil
}
