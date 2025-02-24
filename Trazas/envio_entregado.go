// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnvioEntregado.avsc
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

type EnvioEntregado struct {
	Traza Traza `json:"traza"`

	TipoDeEntrega TipoDeEntrega `json:"tipoDeEntrega"`
}

const EnvioEntregadoAvroCRC64Fingerprint = "l\xd8'\xe2c\x88rf"

func NewEnvioEntregado() EnvioEntregado {
	r := EnvioEntregado{}
	r.Traza = NewTraza()

	return r
}

func DeserializeEnvioEntregado(r io.Reader) (EnvioEntregado, error) {
	t := NewEnvioEntregado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnvioEntregadoFromSchema(r io.Reader, schema string) (EnvioEntregado, error) {
	t := NewEnvioEntregado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnvioEntregado(r EnvioEntregado, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = writeTipoDeEntrega(r.TipoDeEntrega, w)
	if err != nil {
		return err
	}
	return err
}

func (r EnvioEntregado) Serialize(w io.Writer) error {
	return writeEnvioEntregado(r, w)
}

func (r EnvioEntregado) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"LogisticShipmentType\",\"type\":[\"null\",\"string\"]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"name\":\"tipoDeEntrega\",\"type\":{\"name\":\"TipoDeEntrega\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"symbols\":[\"undefined\",\"distribucion\",\"deliveryWindow\"],\"type\":\"enum\"}}],\"name\":\"Integracion.Esquemas.Trazas.EnvioEntregado\",\"type\":\"record\"}"
}

func (r EnvioEntregado) SchemaName() string {
	return "Integracion.Esquemas.Trazas.EnvioEntregado"
}

func (_ EnvioEntregado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EnvioEntregado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EnvioEntregado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EnvioEntregado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EnvioEntregado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EnvioEntregado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EnvioEntregado) SetString(v string)   { panic("Unsupported operation") }
func (_ EnvioEntregado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EnvioEntregado) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		w := TipoDeEntregaWrapper{Target: &r.TipoDeEntrega}

		return w

	}
	panic("Unknown field index")
}

func (r *EnvioEntregado) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EnvioEntregado) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EnvioEntregado) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EnvioEntregado) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EnvioEntregado) HintSize(int)                     { panic("Unsupported operation") }
func (_ EnvioEntregado) Finalize()                        {}

func (_ EnvioEntregado) AvroCRC64Fingerprint() []byte {
	return []byte(EnvioEntregadoAvroCRC64Fingerprint)
}

func (r EnvioEntregado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["tipoDeEntrega"], err = json.Marshal(r.TipoDeEntrega)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EnvioEntregado) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["tipoDeEntrega"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeEntrega); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tipoDeEntrega")
	}
	return nil
}
