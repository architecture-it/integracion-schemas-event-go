// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ERA.avsc
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

type ERA struct {
	EraRangoInicio *UnionNullLong `json:"eraRangoInicio"`

	EraRangoFin *UnionNullLong `json:"eraRangoFin"`

	Era *UnionNullLong `json:"era"`

	CodigoDeEnvio *UnionNullString `json:"codigoDeEnvio"`

	CalculoEra CalculoEra `json:"calculoEra"`

	Transportista Transportista `json:"transportista"`

	Timestamp *UnionNullLong `json:"timestamp"`

	Cuando *UnionNullLong `json:"cuando"`
}

const ERAAvroCRC64Fingerprint = "5\b\xa7\x80\\\x98\x06+"

func NewERA() ERA {
	r := ERA{}
	r.EraRangoInicio = nil
	r.EraRangoFin = nil
	r.Era = nil
	r.CodigoDeEnvio = nil
	r.CalculoEra = NewCalculoEra()

	r.Transportista = NewTransportista()

	r.Timestamp = nil
	r.Cuando = nil
	return r
}

func DeserializeERA(r io.Reader) (ERA, error) {
	t := NewERA()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeERAFromSchema(r io.Reader, schema string) (ERA, error) {
	t := NewERA()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeERA(r ERA, w io.Writer) error {
	var err error
	err = writeUnionNullLong(r.EraRangoInicio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.EraRangoFin, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.Era, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoDeEnvio, w)
	if err != nil {
		return err
	}
	err = writeCalculoEra(r.CalculoEra, w)
	if err != nil {
		return err
	}
	err = writeTransportista(r.Transportista, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.Timestamp, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.Cuando, w)
	if err != nil {
		return err
	}
	return err
}

func (r ERA) Serialize(w io.Writer) error {
	return writeERA(r, w)
}

func (r ERA) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"eraRangoInicio\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"eraRangoFin\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"era\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"codigoDeEnvio\",\"type\":[\"null\",\"string\"]},{\"name\":\"calculoEra\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"ordenDeEnvioEnHR\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"numeroHojaDeRuta\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"geocoordenadas\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"recorridoEnSegundos\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"recorridoEnMetros\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"demoraEnDomicilioEnMinutos\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"demoraSalidaSucursalEnMinutos\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"eraAnterior\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"fechaCreacionHojaDeRuta\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"CalculoEra\",\"type\":\"record\"}},{\"name\":\"transportista\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"esEventual\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"idGla\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idGli\",\"type\":[\"null\",\"string\"]},{\"name\":\"SucursalDondeTrabaja\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"codigoAlertran\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoIntegra\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]}],\"name\":\"SucursalDondeTrabaja\",\"type\":\"record\"}},{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoDeDocumento\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"cumplimientoSecuenciaHR\",\"type\":[\"null\",\"double\"]}],\"name\":\"Transportista\",\"type\":\"record\"}},{\"default\":null,\"name\":\"timestamp\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"cuando\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"Andreani.DeliveryEstimate.Events.Records.ERA\",\"type\":\"record\"}"
}

func (r ERA) SchemaName() string {
	return "Andreani.DeliveryEstimate.Events.Records.ERA"
}

func (_ ERA) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ERA) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ERA) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ERA) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ERA) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ERA) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ERA) SetString(v string)   { panic("Unsupported operation") }
func (_ ERA) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ERA) Get(i int) types.Field {
	switch i {
	case 0:
		r.EraRangoInicio = NewUnionNullLong()

		return r.EraRangoInicio
	case 1:
		r.EraRangoFin = NewUnionNullLong()

		return r.EraRangoFin
	case 2:
		r.Era = NewUnionNullLong()

		return r.Era
	case 3:
		r.CodigoDeEnvio = NewUnionNullString()

		return r.CodigoDeEnvio
	case 4:
		r.CalculoEra = NewCalculoEra()

		w := types.Record{Target: &r.CalculoEra}

		return w

	case 5:
		r.Transportista = NewTransportista()

		w := types.Record{Target: &r.Transportista}

		return w

	case 6:
		r.Timestamp = NewUnionNullLong()

		return r.Timestamp
	case 7:
		r.Cuando = NewUnionNullLong()

		return r.Cuando
	}
	panic("Unknown field index")
}

func (r *ERA) SetDefault(i int) {
	switch i {
	case 0:
		r.EraRangoInicio = nil
		return
	case 1:
		r.EraRangoFin = nil
		return
	case 2:
		r.Era = nil
		return
	case 3:
		r.CodigoDeEnvio = nil
		return
	case 6:
		r.Timestamp = nil
		return
	case 7:
		r.Cuando = nil
		return
	}
	panic("Unknown field index")
}

func (r *ERA) NullField(i int) {
	switch i {
	case 0:
		r.EraRangoInicio = nil
		return
	case 1:
		r.EraRangoFin = nil
		return
	case 2:
		r.Era = nil
		return
	case 3:
		r.CodigoDeEnvio = nil
		return
	case 6:
		r.Timestamp = nil
		return
	case 7:
		r.Cuando = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ERA) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ERA) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ERA) HintSize(int)                     { panic("Unsupported operation") }
func (_ ERA) Finalize()                        {}

func (_ ERA) AvroCRC64Fingerprint() []byte {
	return []byte(ERAAvroCRC64Fingerprint)
}

func (r ERA) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["eraRangoInicio"], err = json.Marshal(r.EraRangoInicio)
	if err != nil {
		return nil, err
	}
	output["eraRangoFin"], err = json.Marshal(r.EraRangoFin)
	if err != nil {
		return nil, err
	}
	output["era"], err = json.Marshal(r.Era)
	if err != nil {
		return nil, err
	}
	output["codigoDeEnvio"], err = json.Marshal(r.CodigoDeEnvio)
	if err != nil {
		return nil, err
	}
	output["calculoEra"], err = json.Marshal(r.CalculoEra)
	if err != nil {
		return nil, err
	}
	output["transportista"], err = json.Marshal(r.Transportista)
	if err != nil {
		return nil, err
	}
	output["timestamp"], err = json.Marshal(r.Timestamp)
	if err != nil {
		return nil, err
	}
	output["cuando"], err = json.Marshal(r.Cuando)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ERA) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["eraRangoInicio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EraRangoInicio); err != nil {
			return err
		}
	} else {
		r.EraRangoInicio = NewUnionNullLong()

		r.EraRangoInicio = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["eraRangoFin"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EraRangoFin); err != nil {
			return err
		}
	} else {
		r.EraRangoFin = NewUnionNullLong()

		r.EraRangoFin = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["era"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Era); err != nil {
			return err
		}
	} else {
		r.Era = NewUnionNullLong()

		r.Era = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoDeEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeEnvio); err != nil {
			return err
		}
	} else {
		r.CodigoDeEnvio = NewUnionNullString()

		r.CodigoDeEnvio = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["calculoEra"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CalculoEra); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for calculoEra")
	}
	val = func() json.RawMessage {
		if v, ok := fields["transportista"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Transportista); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for transportista")
	}
	val = func() json.RawMessage {
		if v, ok := fields["timestamp"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Timestamp); err != nil {
			return err
		}
	} else {
		r.Timestamp = NewUnionNullLong()

		r.Timestamp = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["cuando"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cuando); err != nil {
			return err
		}
	} else {
		r.Cuando = NewUnionNullLong()

		r.Cuando = nil
	}
	return nil
}
