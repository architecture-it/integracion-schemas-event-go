// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     BultoSiniestro.avsc
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

type BultoSiniestro struct {
	Traza TrazaDeBulto `json:"traza"`

	NumeroDeEnvioDelCliente string `json:"numeroDeEnvioDelCliente"`

	Motivo *UnionNullString `json:"motivo"`

	Submotivo *UnionNullString `json:"submotivo"`
}

const BultoSiniestroAvroCRC64Fingerprint = "\xb1\xbb\x05\x05\xe0\xd1*\xcf"

func NewBultoSiniestro() BultoSiniestro {
	r := BultoSiniestro{}
	r.Traza = NewTrazaDeBulto()

	r.Motivo = nil
	r.Submotivo = nil
	return r
}

func DeserializeBultoSiniestro(r io.Reader) (BultoSiniestro, error) {
	t := NewBultoSiniestro()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBultoSiniestroFromSchema(r io.Reader, schema string) (BultoSiniestro, error) {
	t := NewBultoSiniestro()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBultoSiniestro(r BultoSiniestro, w io.Writer) error {
	var err error
	err = writeTrazaDeBulto(r.Traza, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroDeEnvioDelCliente, w)
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
	return err
}

func (r BultoSiniestro) Serialize(w io.Writer) error {
	return writeBultoSiniestro(r, w)
}

func (r BultoSiniestro) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"name\":\"numeroDeBulto\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"tipoDeBulto\",\"type\":[\"null\",{\"name\":\"TipoDeBulto\",\"symbols\":[\"Paquete\",\"Contenedor\"],\"type\":\"enum\"}]}],\"name\":\"TrazaDeBulto\",\"namespace\":\"Integracion.Esquemas.Bulto.Referencias\",\"type\":\"record\"}},{\"name\":\"numeroDeEnvioDelCliente\",\"type\":\"string\"},{\"default\":null,\"name\":\"motivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"submotivo\",\"type\":[\"null\",\"string\"]}],\"name\":\"Integracion.Esquemas.Bulto.Trazas.BultoSiniestro\",\"type\":\"record\"}"
}

func (r BultoSiniestro) SchemaName() string {
	return "Integracion.Esquemas.Bulto.Trazas.BultoSiniestro"
}

func (_ BultoSiniestro) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BultoSiniestro) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BultoSiniestro) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BultoSiniestro) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BultoSiniestro) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BultoSiniestro) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BultoSiniestro) SetString(v string)   { panic("Unsupported operation") }
func (_ BultoSiniestro) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BultoSiniestro) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTrazaDeBulto()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		w := types.String{Target: &r.NumeroDeEnvioDelCliente}

		return w

	case 2:
		r.Motivo = NewUnionNullString()

		return r.Motivo
	case 3:
		r.Submotivo = NewUnionNullString()

		return r.Submotivo
	}
	panic("Unknown field index")
}

func (r *BultoSiniestro) SetDefault(i int) {
	switch i {
	case 2:
		r.Motivo = nil
		return
	case 3:
		r.Submotivo = nil
		return
	}
	panic("Unknown field index")
}

func (r *BultoSiniestro) NullField(i int) {
	switch i {
	case 2:
		r.Motivo = nil
		return
	case 3:
		r.Submotivo = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ BultoSiniestro) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ BultoSiniestro) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ BultoSiniestro) HintSize(int)                     { panic("Unsupported operation") }
func (_ BultoSiniestro) Finalize()                        {}

func (_ BultoSiniestro) AvroCRC64Fingerprint() []byte {
	return []byte(BultoSiniestroAvroCRC64Fingerprint)
}

func (r BultoSiniestro) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["numeroDeEnvioDelCliente"], err = json.Marshal(r.NumeroDeEnvioDelCliente)
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
	return json.Marshal(output)
}

func (r *BultoSiniestro) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["numeroDeEnvioDelCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeEnvioDelCliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroDeEnvioDelCliente")
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
	return nil
}
