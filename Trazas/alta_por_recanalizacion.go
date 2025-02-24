// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     AltaPorRecanalizacion.avsc
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

type AltaPorRecanalizacion struct {
	Traza Traza `json:"traza"`

	Donde DatosSucursal `json:"donde"`

	GeneradaPorNumeroDeEnvio *UnionNullString `json:"generadaPorNumeroDeEnvio"`
}

const AltaPorRecanalizacionAvroCRC64Fingerprint = "1\x8f\xfdo\xfe\xd7P~"

func NewAltaPorRecanalizacion() AltaPorRecanalizacion {
	r := AltaPorRecanalizacion{}
	r.Traza = NewTraza()

	r.Donde = NewDatosSucursal()

	return r
}

func DeserializeAltaPorRecanalizacion(r io.Reader) (AltaPorRecanalizacion, error) {
	t := NewAltaPorRecanalizacion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAltaPorRecanalizacionFromSchema(r io.Reader, schema string) (AltaPorRecanalizacion, error) {
	t := NewAltaPorRecanalizacion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAltaPorRecanalizacion(r AltaPorRecanalizacion, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = writeDatosSucursal(r.Donde, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.GeneradaPorNumeroDeEnvio, w)
	if err != nil {
		return err
	}
	return err
}

func (r AltaPorRecanalizacion) Serialize(w io.Writer) error {
	return writeAltaPorRecanalizacion(r, w)
}

func (r AltaPorRecanalizacion) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"LogisticShipmentType\",\"type\":[\"null\",\"string\"]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"name\":\"donde\",\"type\":\"Integracion.Esquemas.Referencias.DatosSucursal\"},{\"name\":\"generadaPorNumeroDeEnvio\",\"type\":[\"null\",\"string\"]}],\"name\":\"Integracion.Esquemas.Trazas.AltaPorRecanalizacion\",\"type\":\"record\"}"
}

func (r AltaPorRecanalizacion) SchemaName() string {
	return "Integracion.Esquemas.Trazas.AltaPorRecanalizacion"
}

func (_ AltaPorRecanalizacion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ AltaPorRecanalizacion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ AltaPorRecanalizacion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ AltaPorRecanalizacion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ AltaPorRecanalizacion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ AltaPorRecanalizacion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ AltaPorRecanalizacion) SetString(v string)   { panic("Unsupported operation") }
func (_ AltaPorRecanalizacion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AltaPorRecanalizacion) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		r.Donde = NewDatosSucursal()

		w := types.Record{Target: &r.Donde}

		return w

	case 2:
		r.GeneradaPorNumeroDeEnvio = NewUnionNullString()

		return r.GeneradaPorNumeroDeEnvio
	}
	panic("Unknown field index")
}

func (r *AltaPorRecanalizacion) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *AltaPorRecanalizacion) NullField(i int) {
	switch i {
	case 2:
		r.GeneradaPorNumeroDeEnvio = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ AltaPorRecanalizacion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ AltaPorRecanalizacion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ AltaPorRecanalizacion) HintSize(int)                     { panic("Unsupported operation") }
func (_ AltaPorRecanalizacion) Finalize()                        {}

func (_ AltaPorRecanalizacion) AvroCRC64Fingerprint() []byte {
	return []byte(AltaPorRecanalizacionAvroCRC64Fingerprint)
}

func (r AltaPorRecanalizacion) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["donde"], err = json.Marshal(r.Donde)
	if err != nil {
		return nil, err
	}
	output["generadaPorNumeroDeEnvio"], err = json.Marshal(r.GeneradaPorNumeroDeEnvio)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *AltaPorRecanalizacion) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["donde"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Donde); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for donde")
	}
	val = func() json.RawMessage {
		if v, ok := fields["generadaPorNumeroDeEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.GeneradaPorNumeroDeEnvio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for generadaPorNumeroDeEnvio")
	}
	return nil
}
