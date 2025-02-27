// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     BultoSolicitudDeRescate.avsc
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

type BultoSolicitudDeRescate struct {
	Traza TrazaDeBulto `json:"traza"`

	EsInterno bool `json:"esInterno"`

	SucursalDondeSeRescata *UnionNullDatosSucursal `json:"sucursalDondeSeRescata"`
}

const BultoSolicitudDeRescateAvroCRC64Fingerprint = "\xab\xa17/39Ȼ"

func NewBultoSolicitudDeRescate() BultoSolicitudDeRescate {
	r := BultoSolicitudDeRescate{}
	r.Traza = NewTrazaDeBulto()

	r.SucursalDondeSeRescata = nil
	return r
}

func DeserializeBultoSolicitudDeRescate(r io.Reader) (BultoSolicitudDeRescate, error) {
	t := NewBultoSolicitudDeRescate()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBultoSolicitudDeRescateFromSchema(r io.Reader, schema string) (BultoSolicitudDeRescate, error) {
	t := NewBultoSolicitudDeRescate()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBultoSolicitudDeRescate(r BultoSolicitudDeRescate, w io.Writer) error {
	var err error
	err = writeTrazaDeBulto(r.Traza, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.EsInterno, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDatosSucursal(r.SucursalDondeSeRescata, w)
	if err != nil {
		return err
	}
	return err
}

func (r BultoSolicitudDeRescate) Serialize(w io.Writer) error {
	return writeBultoSolicitudDeRescate(r, w)
}

func (r BultoSolicitudDeRescate) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"name\":\"numeroDeBulto\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"tipoDeBulto\",\"type\":[\"null\",{\"name\":\"TipoDeBulto\",\"symbols\":[\"Paquete\",\"Contenedor\"],\"type\":\"enum\"}]}],\"name\":\"TrazaDeBulto\",\"namespace\":\"Integracion.Esquemas.Bulto.Referencias\",\"type\":\"record\"}},{\"name\":\"esInterno\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"sucursalDondeSeRescata\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosSucursal\"]}],\"name\":\"Integracion.Esquemas.Bulto.Trazas.BultoSolicitudDeRescate\",\"type\":\"record\"}"
}

func (r BultoSolicitudDeRescate) SchemaName() string {
	return "Integracion.Esquemas.Bulto.Trazas.BultoSolicitudDeRescate"
}

func (_ BultoSolicitudDeRescate) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BultoSolicitudDeRescate) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BultoSolicitudDeRescate) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BultoSolicitudDeRescate) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BultoSolicitudDeRescate) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BultoSolicitudDeRescate) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BultoSolicitudDeRescate) SetString(v string)   { panic("Unsupported operation") }
func (_ BultoSolicitudDeRescate) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BultoSolicitudDeRescate) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTrazaDeBulto()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		w := types.Boolean{Target: &r.EsInterno}

		return w

	case 2:
		r.SucursalDondeSeRescata = NewUnionNullDatosSucursal()

		return r.SucursalDondeSeRescata
	}
	panic("Unknown field index")
}

func (r *BultoSolicitudDeRescate) SetDefault(i int) {
	switch i {
	case 2:
		r.SucursalDondeSeRescata = nil
		return
	}
	panic("Unknown field index")
}

func (r *BultoSolicitudDeRescate) NullField(i int) {
	switch i {
	case 2:
		r.SucursalDondeSeRescata = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ BultoSolicitudDeRescate) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ BultoSolicitudDeRescate) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ BultoSolicitudDeRescate) HintSize(int)                     { panic("Unsupported operation") }
func (_ BultoSolicitudDeRescate) Finalize()                        {}

func (_ BultoSolicitudDeRescate) AvroCRC64Fingerprint() []byte {
	return []byte(BultoSolicitudDeRescateAvroCRC64Fingerprint)
}

func (r BultoSolicitudDeRescate) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["esInterno"], err = json.Marshal(r.EsInterno)
	if err != nil {
		return nil, err
	}
	output["sucursalDondeSeRescata"], err = json.Marshal(r.SucursalDondeSeRescata)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *BultoSolicitudDeRescate) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["esInterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EsInterno); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for esInterno")
	}
	val = func() json.RawMessage {
		if v, ok := fields["sucursalDondeSeRescata"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SucursalDondeSeRescata); err != nil {
			return err
		}
	} else {
		r.SucursalDondeSeRescata = NewUnionNullDatosSucursal()

		r.SucursalDondeSeRescata = nil
	}
	return nil
}
