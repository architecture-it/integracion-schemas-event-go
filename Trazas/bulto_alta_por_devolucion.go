// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     BultoAltaPorDevolucion.avsc
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

type BultoAltaPorDevolucion struct {
	Traza TrazaDeBulto `json:"traza"`

	Donde DatosSucursal `json:"donde"`

	GeneradaPorNumeroDeEnvio *UnionNullString `json:"generadaPorNumeroDeEnvio"`
}

const BultoAltaPorDevolucionAvroCRC64Fingerprint = "q\\s1\x8f\x8e\x86a"

func NewBultoAltaPorDevolucion() BultoAltaPorDevolucion {
	r := BultoAltaPorDevolucion{}
	r.Traza = NewTrazaDeBulto()

	r.Donde = NewDatosSucursal()

	return r
}

func DeserializeBultoAltaPorDevolucion(r io.Reader) (BultoAltaPorDevolucion, error) {
	t := NewBultoAltaPorDevolucion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBultoAltaPorDevolucionFromSchema(r io.Reader, schema string) (BultoAltaPorDevolucion, error) {
	t := NewBultoAltaPorDevolucion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBultoAltaPorDevolucion(r BultoAltaPorDevolucion, w io.Writer) error {
	var err error
	err = writeTrazaDeBulto(r.Traza, w)
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

func (r BultoAltaPorDevolucion) Serialize(w io.Writer) error {
	return writeBultoAltaPorDevolucion(r, w)
}

func (r BultoAltaPorDevolucion) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"name\":\"numeroDeBulto\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"tipoDeBulto\",\"type\":[\"null\",{\"name\":\"TipoDeBulto\",\"symbols\":[\"Paquete\",\"Contenedor\"],\"type\":\"enum\"}]}],\"name\":\"TrazaDeBulto\",\"namespace\":\"Integracion.Esquemas.Bulto.Referencias\",\"type\":\"record\"}},{\"name\":\"donde\",\"type\":\"Integracion.Esquemas.Referencias.DatosSucursal\"},{\"name\":\"generadaPorNumeroDeEnvio\",\"type\":[\"null\",\"string\"]}],\"name\":\"Integracion.Esquemas.Bulto.Trazas.BultoAltaPorDevolucion\",\"type\":\"record\"}"
}

func (r BultoAltaPorDevolucion) SchemaName() string {
	return "Integracion.Esquemas.Bulto.Trazas.BultoAltaPorDevolucion"
}

func (_ BultoAltaPorDevolucion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BultoAltaPorDevolucion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BultoAltaPorDevolucion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BultoAltaPorDevolucion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BultoAltaPorDevolucion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BultoAltaPorDevolucion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BultoAltaPorDevolucion) SetString(v string)   { panic("Unsupported operation") }
func (_ BultoAltaPorDevolucion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BultoAltaPorDevolucion) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTrazaDeBulto()

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

func (r *BultoAltaPorDevolucion) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *BultoAltaPorDevolucion) NullField(i int) {
	switch i {
	case 2:
		r.GeneradaPorNumeroDeEnvio = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ BultoAltaPorDevolucion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ BultoAltaPorDevolucion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ BultoAltaPorDevolucion) HintSize(int)                     { panic("Unsupported operation") }
func (_ BultoAltaPorDevolucion) Finalize()                        {}

func (_ BultoAltaPorDevolucion) AvroCRC64Fingerprint() []byte {
	return []byte(BultoAltaPorDevolucionAvroCRC64Fingerprint)
}

func (r BultoAltaPorDevolucion) MarshalJSON() ([]byte, error) {
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

func (r *BultoAltaPorDevolucion) UnmarshalJSON(data []byte) error {
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
