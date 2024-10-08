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
	Era *UnionNullLong `json:"era"`

	CodigoDeEnvio *UnionNullString `json:"codigoDeEnvio"`

	CalculoEra CalculoEra `json:"calculoEra"`

	Transportista Transportista `json:"transportista"`

	Cuando *UnionNullLong `json:"cuando"`

	Estimacion *UnionNullString `json:"estimacion"`

	EstimacionDescripcion *UnionNullString `json:"estimacionDescripcion"`

	EstimacionI *UnionNullString `json:"estimacionI"`

	EstimacionDescripcionI *UnionNullString `json:"estimacionDescripcionI"`

	EstimacionII *UnionNullString `json:"estimacionII"`

	EstimacionDescripcionII *UnionNullString `json:"estimacionDescripcionII"`

	EstimacionIII *UnionNullString `json:"estimacionIII"`

	EstimacionDescripcionIII *UnionNullString `json:"estimacionDescripcionIII"`
}

const ERAAvroCRC64Fingerprint = "6Y M8\x11\x05T"

func NewERA() ERA {
	r := ERA{}
	r.Era = nil
	r.CodigoDeEnvio = nil
	r.CalculoEra = NewCalculoEra()

	r.Transportista = NewTransportista()

	r.Cuando = nil
	r.Estimacion = nil
	r.EstimacionDescripcion = nil
	r.EstimacionI = nil
	r.EstimacionDescripcionI = nil
	r.EstimacionII = nil
	r.EstimacionDescripcionII = nil
	r.EstimacionIII = nil
	r.EstimacionDescripcionIII = nil
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
	err = writeUnionNullLong(r.Cuando, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Estimacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstimacionDescripcion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstimacionI, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstimacionDescripcionI, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstimacionII, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstimacionDescripcionII, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstimacionIII, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstimacionDescripcionIII, w)
	if err != nil {
		return err
	}
	return err
}

func (r ERA) Serialize(w io.Writer) error {
	return writeERA(r, w)
}

func (r ERA) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"era\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"codigoDeEnvio\",\"type\":[\"null\",\"string\"]},{\"name\":\"calculoEra\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"ordenDeEnvioEnHR\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"numeroHojaDeRuta\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"geocoordenadas\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"recorridoEnSegundos\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"recorridoEnMetros\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"demoraEnDomicilioEnMinutos\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"demoraSalidaSucursalEnMinutos\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"eraAnterior\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"fechaCreacionHojaDeRuta\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"numeroContenedor\",\"type\":[\"null\",\"string\"]}],\"name\":\"CalculoEra\",\"type\":\"record\"}},{\"name\":\"transportista\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"esEventual\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"idGla\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idGli\",\"type\":[\"null\",\"string\"]},{\"name\":\"sucursalDondeTrabaja\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"codigoAlertran\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoIntegra\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]}],\"name\":\"SucursalDondeTrabaja\",\"type\":\"record\"}},{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoDeDocumento\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"cumplimientoSecuenciaHR\",\"type\":[\"null\",\"double\"]}],\"name\":\"Transportista\",\"type\":\"record\"}},{\"default\":null,\"name\":\"cuando\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"estimacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estimacionDescripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estimacionI\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estimacionDescripcionI\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estimacionII\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estimacionDescripcionII\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estimacionIII\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estimacionDescripcionIII\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.DeliveryEstimate.Events.Records.ERA\",\"type\":\"record\"}"
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
		r.Era = NewUnionNullLong()

		return r.Era
	case 1:
		r.CodigoDeEnvio = NewUnionNullString()

		return r.CodigoDeEnvio
	case 2:
		r.CalculoEra = NewCalculoEra()

		w := types.Record{Target: &r.CalculoEra}

		return w

	case 3:
		r.Transportista = NewTransportista()

		w := types.Record{Target: &r.Transportista}

		return w

	case 4:
		r.Cuando = NewUnionNullLong()

		return r.Cuando
	case 5:
		r.Estimacion = NewUnionNullString()

		return r.Estimacion
	case 6:
		r.EstimacionDescripcion = NewUnionNullString()

		return r.EstimacionDescripcion
	case 7:
		r.EstimacionI = NewUnionNullString()

		return r.EstimacionI
	case 8:
		r.EstimacionDescripcionI = NewUnionNullString()

		return r.EstimacionDescripcionI
	case 9:
		r.EstimacionII = NewUnionNullString()

		return r.EstimacionII
	case 10:
		r.EstimacionDescripcionII = NewUnionNullString()

		return r.EstimacionDescripcionII
	case 11:
		r.EstimacionIII = NewUnionNullString()

		return r.EstimacionIII
	case 12:
		r.EstimacionDescripcionIII = NewUnionNullString()

		return r.EstimacionDescripcionIII
	}
	panic("Unknown field index")
}

func (r *ERA) SetDefault(i int) {
	switch i {
	case 0:
		r.Era = nil
		return
	case 1:
		r.CodigoDeEnvio = nil
		return
	case 4:
		r.Cuando = nil
		return
	case 5:
		r.Estimacion = nil
		return
	case 6:
		r.EstimacionDescripcion = nil
		return
	case 7:
		r.EstimacionI = nil
		return
	case 8:
		r.EstimacionDescripcionI = nil
		return
	case 9:
		r.EstimacionII = nil
		return
	case 10:
		r.EstimacionDescripcionII = nil
		return
	case 11:
		r.EstimacionIII = nil
		return
	case 12:
		r.EstimacionDescripcionIII = nil
		return
	}
	panic("Unknown field index")
}

func (r *ERA) NullField(i int) {
	switch i {
	case 0:
		r.Era = nil
		return
	case 1:
		r.CodigoDeEnvio = nil
		return
	case 4:
		r.Cuando = nil
		return
	case 5:
		r.Estimacion = nil
		return
	case 6:
		r.EstimacionDescripcion = nil
		return
	case 7:
		r.EstimacionI = nil
		return
	case 8:
		r.EstimacionDescripcionI = nil
		return
	case 9:
		r.EstimacionII = nil
		return
	case 10:
		r.EstimacionDescripcionII = nil
		return
	case 11:
		r.EstimacionIII = nil
		return
	case 12:
		r.EstimacionDescripcionIII = nil
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
	output["cuando"], err = json.Marshal(r.Cuando)
	if err != nil {
		return nil, err
	}
	output["estimacion"], err = json.Marshal(r.Estimacion)
	if err != nil {
		return nil, err
	}
	output["estimacionDescripcion"], err = json.Marshal(r.EstimacionDescripcion)
	if err != nil {
		return nil, err
	}
	output["estimacionI"], err = json.Marshal(r.EstimacionI)
	if err != nil {
		return nil, err
	}
	output["estimacionDescripcionI"], err = json.Marshal(r.EstimacionDescripcionI)
	if err != nil {
		return nil, err
	}
	output["estimacionII"], err = json.Marshal(r.EstimacionII)
	if err != nil {
		return nil, err
	}
	output["estimacionDescripcionII"], err = json.Marshal(r.EstimacionDescripcionII)
	if err != nil {
		return nil, err
	}
	output["estimacionIII"], err = json.Marshal(r.EstimacionIII)
	if err != nil {
		return nil, err
	}
	output["estimacionDescripcionIII"], err = json.Marshal(r.EstimacionDescripcionIII)
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
	val = func() json.RawMessage {
		if v, ok := fields["estimacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estimacion); err != nil {
			return err
		}
	} else {
		r.Estimacion = NewUnionNullString()

		r.Estimacion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["estimacionDescripcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstimacionDescripcion); err != nil {
			return err
		}
	} else {
		r.EstimacionDescripcion = NewUnionNullString()

		r.EstimacionDescripcion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["estimacionI"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstimacionI); err != nil {
			return err
		}
	} else {
		r.EstimacionI = NewUnionNullString()

		r.EstimacionI = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["estimacionDescripcionI"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstimacionDescripcionI); err != nil {
			return err
		}
	} else {
		r.EstimacionDescripcionI = NewUnionNullString()

		r.EstimacionDescripcionI = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["estimacionII"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstimacionII); err != nil {
			return err
		}
	} else {
		r.EstimacionII = NewUnionNullString()

		r.EstimacionII = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["estimacionDescripcionII"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstimacionDescripcionII); err != nil {
			return err
		}
	} else {
		r.EstimacionDescripcionII = NewUnionNullString()

		r.EstimacionDescripcionII = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["estimacionIII"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstimacionIII); err != nil {
			return err
		}
	} else {
		r.EstimacionIII = NewUnionNullString()

		r.EstimacionIII = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["estimacionDescripcionIII"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstimacionDescripcionIII); err != nil {
			return err
		}
	} else {
		r.EstimacionDescripcionIII = NewUnionNullString()

		r.EstimacionDescripcionIII = nil
	}
	return nil
}
