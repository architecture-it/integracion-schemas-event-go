// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     BultoNoEntregado.avsc
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

type BultoNoEntregado struct {
	Traza TrazaDeBulto `json:"traza"`

	Motivo *UnionNullString `json:"motivo"`

	Submotivo *UnionNullString `json:"submotivo"`

	ReemplazadoPorNumeroDeEnvio *UnionNullString `json:"reemplazadoPorNumeroDeEnvio"`
}

const BultoNoEntregadoAvroCRC64Fingerprint = "\x90\xcb\xd7&\xb6\x92߂"

func NewBultoNoEntregado() BultoNoEntregado {
	r := BultoNoEntregado{}
	r.Traza = NewTrazaDeBulto()

	r.Motivo = nil
	r.Submotivo = nil
	r.ReemplazadoPorNumeroDeEnvio = nil
	return r
}

func DeserializeBultoNoEntregado(r io.Reader) (BultoNoEntregado, error) {
	t := NewBultoNoEntregado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBultoNoEntregadoFromSchema(r io.Reader, schema string) (BultoNoEntregado, error) {
	t := NewBultoNoEntregado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBultoNoEntregado(r BultoNoEntregado, w io.Writer) error {
	var err error
	err = writeTrazaDeBulto(r.Traza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Motivo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Submotivo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ReemplazadoPorNumeroDeEnvio, w)
	if err != nil {
		return err
	}
	return err
}

func (r BultoNoEntregado) Serialize(w io.Writer) error {
	return writeBultoNoEntregado(r, w)
}

func (r BultoNoEntregado) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"name\":\"numeroDeBulto\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"tipoDeBulto\",\"type\":[\"null\",{\"name\":\"TipoDeBulto\",\"symbols\":[\"Paquete\",\"Contenedor\"],\"type\":\"enum\"}]}],\"name\":\"TrazaDeBulto\",\"namespace\":\"Integracion.Esquemas.Bulto.Referencias\",\"type\":\"record\"}},{\"default\":null,\"name\":\"motivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"submotivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"reemplazadoPorNumeroDeEnvio\",\"type\":[\"null\",\"string\"]}],\"name\":\"Integracion.Esquemas.Bulto.Trazas.BultoNoEntregado\",\"type\":\"record\"}"
}

func (r BultoNoEntregado) SchemaName() string {
	return "Integracion.Esquemas.Bulto.Trazas.BultoNoEntregado"
}

func (_ BultoNoEntregado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BultoNoEntregado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BultoNoEntregado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BultoNoEntregado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BultoNoEntregado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BultoNoEntregado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BultoNoEntregado) SetString(v string)   { panic("Unsupported operation") }
func (_ BultoNoEntregado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BultoNoEntregado) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTrazaDeBulto()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		r.Motivo = NewUnionNullString()

		return r.Motivo
	case 2:
		r.Submotivo = NewUnionNullString()

		return r.Submotivo
	case 3:
		r.ReemplazadoPorNumeroDeEnvio = NewUnionNullString()

		return r.ReemplazadoPorNumeroDeEnvio
	}
	panic("Unknown field index")
}

func (r *BultoNoEntregado) SetDefault(i int) {
	switch i {
	case 1:
		r.Motivo = nil
		return
	case 2:
		r.Submotivo = nil
		return
	case 3:
		r.ReemplazadoPorNumeroDeEnvio = nil
		return
	}
	panic("Unknown field index")
}

func (r *BultoNoEntregado) NullField(i int) {
	switch i {
	case 1:
		r.Motivo = nil
		return
	case 2:
		r.Submotivo = nil
		return
	case 3:
		r.ReemplazadoPorNumeroDeEnvio = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ BultoNoEntregado) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ BultoNoEntregado) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ BultoNoEntregado) HintSize(int)                     { panic("Unsupported operation") }
func (_ BultoNoEntregado) Finalize()                        {}

func (_ BultoNoEntregado) AvroCRC64Fingerprint() []byte {
	return []byte(BultoNoEntregadoAvroCRC64Fingerprint)
}

func (r BultoNoEntregado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["motivo"], err = json.Marshal(r.Motivo)
	if err != nil {
		return nil, err
	}
	output["submotivo"], err = json.Marshal(r.Submotivo)
	if err != nil {
		return nil, err
	}
	output["reemplazadoPorNumeroDeEnvio"], err = json.Marshal(r.ReemplazadoPorNumeroDeEnvio)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *BultoNoEntregado) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["motivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Motivo); err != nil {
			return err
		}
	} else {
		r.Motivo = NewUnionNullString()

		r.Motivo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["submotivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Submotivo); err != nil {
			return err
		}
	} else {
		r.Submotivo = NewUnionNullString()

		r.Submotivo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["reemplazadoPorNumeroDeEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ReemplazadoPorNumeroDeEnvio); err != nil {
			return err
		}
	} else {
		r.ReemplazadoPorNumeroDeEnvio = NewUnionNullString()

		r.ReemplazadoPorNumeroDeEnvio = nil
	}
	return nil
}
