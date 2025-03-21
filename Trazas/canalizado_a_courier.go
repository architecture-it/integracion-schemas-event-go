// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CanalizadoACourier.avsc
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

type CanalizadoACourier struct {
	Traza Traza `json:"traza"`

	DatosAdicionales *UnionNullArrayMetadato `json:"datosAdicionales"`

	Courier *UnionNullString `json:"courier"`
}

const CanalizadoACourierAvroCRC64Fingerprint = "\"\xd8\xd2%?\xb5\x8am"

func NewCanalizadoACourier() CanalizadoACourier {
	r := CanalizadoACourier{}
	r.Traza = NewTraza()

	r.DatosAdicionales = nil
	r.Courier = nil
	return r
}

func DeserializeCanalizadoACourier(r io.Reader) (CanalizadoACourier, error) {
	t := NewCanalizadoACourier()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCanalizadoACourierFromSchema(r io.Reader, schema string) (CanalizadoACourier, error) {
	t := NewCanalizadoACourier()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCanalizadoACourier(r CanalizadoACourier, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayMetadato(r.DatosAdicionales, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Courier, w)
	if err != nil {
		return err
	}
	return err
}

func (r CanalizadoACourier) Serialize(w io.Writer) error {
	return writeCanalizadoACourier(r, w)
}

func (r CanalizadoACourier) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"default\":null,\"name\":\"datosAdicionales\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"courier\",\"type\":[\"null\",\"string\"]}],\"name\":\"Integracion.Esquemas.Trazas.CanalizadoACourier\",\"type\":\"record\"}"
}

func (r CanalizadoACourier) SchemaName() string {
	return "Integracion.Esquemas.Trazas.CanalizadoACourier"
}

func (_ CanalizadoACourier) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CanalizadoACourier) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CanalizadoACourier) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CanalizadoACourier) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CanalizadoACourier) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CanalizadoACourier) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CanalizadoACourier) SetString(v string)   { panic("Unsupported operation") }
func (_ CanalizadoACourier) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CanalizadoACourier) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		r.DatosAdicionales = NewUnionNullArrayMetadato()

		return r.DatosAdicionales
	case 2:
		r.Courier = NewUnionNullString()

		return r.Courier
	}
	panic("Unknown field index")
}

func (r *CanalizadoACourier) SetDefault(i int) {
	switch i {
	case 1:
		r.DatosAdicionales = nil
		return
	case 2:
		r.Courier = nil
		return
	}
	panic("Unknown field index")
}

func (r *CanalizadoACourier) NullField(i int) {
	switch i {
	case 1:
		r.DatosAdicionales = nil
		return
	case 2:
		r.Courier = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ CanalizadoACourier) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CanalizadoACourier) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CanalizadoACourier) HintSize(int)                     { panic("Unsupported operation") }
func (_ CanalizadoACourier) Finalize()                        {}

func (_ CanalizadoACourier) AvroCRC64Fingerprint() []byte {
	return []byte(CanalizadoACourierAvroCRC64Fingerprint)
}

func (r CanalizadoACourier) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["datosAdicionales"], err = json.Marshal(r.DatosAdicionales)
	if err != nil {
		return nil, err
	}
	output["courier"], err = json.Marshal(r.Courier)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CanalizadoACourier) UnmarshalJSON(data []byte) error {
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
		r.Courier = NewUnionNullString()

		r.Courier = nil
	}
	return nil
}
