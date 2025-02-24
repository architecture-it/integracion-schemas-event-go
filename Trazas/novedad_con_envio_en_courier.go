// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NovedadConEnvioEnCourier.avsc
 */
package TrazasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type NovedadConEnvioEnCourier struct {
	Traza Traza `json:"traza"`

	QueNovedad string `json:"queNovedad"`

	Courier string `json:"courier"`

	Resultado *UnionNullString `json:"resultado"`

	DatosAdicionales *UnionNullArrayMetadato `json:"datosAdicionales"`
}

const NovedadConEnvioEnCourierAvroCRC64Fingerprint = "\x89Q\x15\x8a\t`\xd8i"

func NewNovedadConEnvioEnCourier() NovedadConEnvioEnCourier {
	r := NovedadConEnvioEnCourier{}
	r.Traza = NewTraza()

	r.Resultado = nil
	r.DatosAdicionales = nil
	return r
}

func DeserializeNovedadConEnvioEnCourier(r io.Reader) (NovedadConEnvioEnCourier, error) {
	t := NewNovedadConEnvioEnCourier()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNovedadConEnvioEnCourierFromSchema(r io.Reader, schema string) (NovedadConEnvioEnCourier, error) {
	t := NewNovedadConEnvioEnCourier()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNovedadConEnvioEnCourier(r NovedadConEnvioEnCourier, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.QueNovedad, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Courier, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Resultado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayMetadato(r.DatosAdicionales, w)
	if err != nil {
		return err
	}
	return err
}

func (r NovedadConEnvioEnCourier) Serialize(w io.Writer) error {
	return writeNovedadConEnvioEnCourier(r, w)
}

func (r NovedadConEnvioEnCourier) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"LogisticShipmentType\",\"type\":[\"null\",\"string\"]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"name\":\"queNovedad\",\"type\":\"string\"},{\"name\":\"courier\",\"type\":\"string\"},{\"default\":null,\"name\":\"resultado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"datosAdicionales\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"Integracion.Esquemas.Trazas.NovedadConEnvioEnCourier\",\"type\":\"record\"}"
}

func (r NovedadConEnvioEnCourier) SchemaName() string {
	return "Integracion.Esquemas.Trazas.NovedadConEnvioEnCourier"
}

func (_ NovedadConEnvioEnCourier) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NovedadConEnvioEnCourier) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NovedadConEnvioEnCourier) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NovedadConEnvioEnCourier) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NovedadConEnvioEnCourier) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NovedadConEnvioEnCourier) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NovedadConEnvioEnCourier) SetString(v string)   { panic("Unsupported operation") }
func (_ NovedadConEnvioEnCourier) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NovedadConEnvioEnCourier) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		w := types.String{Target: &r.QueNovedad}

		return w

	case 2:
		w := types.String{Target: &r.Courier}

		return w

	case 3:
		r.Resultado = NewUnionNullString()

		return r.Resultado
	case 4:
		r.DatosAdicionales = NewUnionNullArrayMetadato()

		return r.DatosAdicionales
	}
	panic("Unknown field index")
}

func (r *NovedadConEnvioEnCourier) SetDefault(i int) {
	switch i {
	case 3:
		r.Resultado = nil
		return
	case 4:
		r.DatosAdicionales = nil
		return
	}
	panic("Unknown field index")
}

func (r *NovedadConEnvioEnCourier) NullField(i int) {
	switch i {
	case 3:
		r.Resultado = nil
		return
	case 4:
		r.DatosAdicionales = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ NovedadConEnvioEnCourier) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NovedadConEnvioEnCourier) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NovedadConEnvioEnCourier) HintSize(int)                     { panic("Unsupported operation") }
func (_ NovedadConEnvioEnCourier) Finalize()                        {}

func (_ NovedadConEnvioEnCourier) AvroCRC64Fingerprint() []byte {
	return []byte(NovedadConEnvioEnCourierAvroCRC64Fingerprint)
}

func (r NovedadConEnvioEnCourier) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["queNovedad"], err = json.Marshal(r.QueNovedad)
	if err != nil {
		return nil, err
	}
	output["courier"], err = json.Marshal(r.Courier)
	if err != nil {
		return nil, err
	}
	output["resultado"], err = json.Marshal(r.Resultado)
	if err != nil {
		return nil, err
	}
	output["datosAdicionales"], err = json.Marshal(r.DatosAdicionales)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NovedadConEnvioEnCourier) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["traza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Traza); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for traza")
	}
	val = func() json.RawMessage {
		if v, ok := fields["queNovedad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.QueNovedad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for queNovedad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["courier"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Courier); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for courier")
	}
	val = func() json.RawMessage {
		if v, ok := fields["resultado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Resultado); err != nil {
			return err
		}
	} else {
		r.Resultado = NewUnionNullString()

		r.Resultado = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["datosAdicionales"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DatosAdicionales); err != nil {
			return err
		}
	} else {
		r.DatosAdicionales = NewUnionNullArrayMetadato()

		r.DatosAdicionales = nil
	}
	return nil
}
