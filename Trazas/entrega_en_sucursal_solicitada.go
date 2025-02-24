// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EntregaEnSucursalSolicitada.avsc
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

type EntregaEnSucursalSolicitada struct {
	Traza Traza `json:"traza"`

	Cantidad int32 `json:"cantidad"`
}

const EntregaEnSucursalSolicitadaAvroCRC64Fingerprint = "ŸM\r&R\xf41"

func NewEntregaEnSucursalSolicitada() EntregaEnSucursalSolicitada {
	r := EntregaEnSucursalSolicitada{}
	r.Traza = NewTraza()

	return r
}

func DeserializeEntregaEnSucursalSolicitada(r io.Reader) (EntregaEnSucursalSolicitada, error) {
	t := NewEntregaEnSucursalSolicitada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEntregaEnSucursalSolicitadaFromSchema(r io.Reader, schema string) (EntregaEnSucursalSolicitada, error) {
	t := NewEntregaEnSucursalSolicitada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEntregaEnSucursalSolicitada(r EntregaEnSucursalSolicitada, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Cantidad, w)
	if err != nil {
		return err
	}
	return err
}

func (r EntregaEnSucursalSolicitada) Serialize(w io.Writer) error {
	return writeEntregaEnSucursalSolicitada(r, w)
}

func (r EntregaEnSucursalSolicitada) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"LogisticShipmentType\",\"type\":[\"null\",\"string\"]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"name\":\"cantidad\",\"type\":\"int\"}],\"name\":\"Integracion.Esquemas.Trazas.EntregaEnSucursalSolicitada\",\"type\":\"record\"}"
}

func (r EntregaEnSucursalSolicitada) SchemaName() string {
	return "Integracion.Esquemas.Trazas.EntregaEnSucursalSolicitada"
}

func (_ EntregaEnSucursalSolicitada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EntregaEnSucursalSolicitada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EntregaEnSucursalSolicitada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EntregaEnSucursalSolicitada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EntregaEnSucursalSolicitada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EntregaEnSucursalSolicitada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EntregaEnSucursalSolicitada) SetString(v string)   { panic("Unsupported operation") }
func (_ EntregaEnSucursalSolicitada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EntregaEnSucursalSolicitada) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		w := types.Int{Target: &r.Cantidad}

		return w

	}
	panic("Unknown field index")
}

func (r *EntregaEnSucursalSolicitada) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EntregaEnSucursalSolicitada) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EntregaEnSucursalSolicitada) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ EntregaEnSucursalSolicitada) AppendArray() types.Field { panic("Unsupported operation") }
func (_ EntregaEnSucursalSolicitada) HintSize(int)             { panic("Unsupported operation") }
func (_ EntregaEnSucursalSolicitada) Finalize()                {}

func (_ EntregaEnSucursalSolicitada) AvroCRC64Fingerprint() []byte {
	return []byte(EntregaEnSucursalSolicitadaAvroCRC64Fingerprint)
}

func (r EntregaEnSucursalSolicitada) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["cantidad"], err = json.Marshal(r.Cantidad)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EntregaEnSucursalSolicitada) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["cantidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cantidad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cantidad")
	}
	return nil
}
