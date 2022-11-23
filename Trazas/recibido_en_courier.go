// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     RecibidoEnCourier.avsc
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

type RecibidoEnCourier struct {
	Traza Traza `json:"traza"`

	Courier string `json:"courier"`

	DatosAdicionales *UnionNullArrayMetadato `json:"datosAdicionales"`
}

const RecibidoEnCourierAvroCRC64Fingerprint = "\xe3; \xeb\x86o\x80\x1c"

func NewRecibidoEnCourier() RecibidoEnCourier {
	r := RecibidoEnCourier{}
	r.Traza = NewTraza()

	r.DatosAdicionales = nil
	return r
}

func DeserializeRecibidoEnCourier(r io.Reader) (RecibidoEnCourier, error) {
	t := NewRecibidoEnCourier()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRecibidoEnCourierFromSchema(r io.Reader, schema string) (RecibidoEnCourier, error) {
	t := NewRecibidoEnCourier()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRecibidoEnCourier(r RecibidoEnCourier, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Courier, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayMetadato(r.DatosAdicionales, w)
	if err != nil {
		return err
	}
	return err
}

func (r RecibidoEnCourier) Serialize(w io.Writer) error {
	return writeRecibidoEnCourier(r, w)
}

func (r RecibidoEnCourier) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"name\":\"courier\",\"type\":\"string\"},{\"default\":null,\"name\":\"datosAdicionales\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"Integracion.Esquemas.Trazas.RecibidoEnCourier\",\"type\":\"record\"}"
}

func (r RecibidoEnCourier) SchemaName() string {
	return "Integracion.Esquemas.Trazas.RecibidoEnCourier"
}

func (_ RecibidoEnCourier) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ RecibidoEnCourier) SetInt(v int32)       { panic("Unsupported operation") }
func (_ RecibidoEnCourier) SetLong(v int64)      { panic("Unsupported operation") }
func (_ RecibidoEnCourier) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ RecibidoEnCourier) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ RecibidoEnCourier) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ RecibidoEnCourier) SetString(v string)   { panic("Unsupported operation") }
func (_ RecibidoEnCourier) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RecibidoEnCourier) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		w := types.String{Target: &r.Courier}

		return w

	case 2:
		r.DatosAdicionales = NewUnionNullArrayMetadato()

		return r.DatosAdicionales
	}
	panic("Unknown field index")
}

func (r *RecibidoEnCourier) SetDefault(i int) {
	switch i {
	case 2:
		r.DatosAdicionales = nil
		return
	}
	panic("Unknown field index")
}

func (r *RecibidoEnCourier) NullField(i int) {
	switch i {
	case 2:
		r.DatosAdicionales = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ RecibidoEnCourier) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ RecibidoEnCourier) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ RecibidoEnCourier) HintSize(int)                     { panic("Unsupported operation") }
func (_ RecibidoEnCourier) Finalize()                        {}

func (_ RecibidoEnCourier) AvroCRC64Fingerprint() []byte {
	return []byte(RecibidoEnCourierAvroCRC64Fingerprint)
}

func (r RecibidoEnCourier) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["courier"], err = json.Marshal(r.Courier)
	if err != nil {
		return nil, err
	}
	output["datosAdicionales"], err = json.Marshal(r.DatosAdicionales)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *RecibidoEnCourier) UnmarshalJSON(data []byte) error {
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
