// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ComienzoCustodiaEnSucursal.avsc
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

type ComienzoCustodiaEnSucursal struct {
	Traza Traza `json:"traza"`

	CantidadDeDiasDeLaFase int32 `json:"cantidadDeDiasDeLaFase"`

	Motivo *UnionNullString `json:"motivo"`

	Submotivo *UnionNullString `json:"submotivo"`

	SucursalDeCustodia *UnionNullDatosSucursal `json:"sucursalDeCustodia"`
}

const ComienzoCustodiaEnSucursalAvroCRC64Fingerprint = "\xfbt\xb3\xca\a\xe4+\xbd"

func NewComienzoCustodiaEnSucursal() ComienzoCustodiaEnSucursal {
	r := ComienzoCustodiaEnSucursal{}
	r.Traza = NewTraza()

	r.Motivo = nil
	r.Submotivo = nil
	r.SucursalDeCustodia = nil
	return r
}

func DeserializeComienzoCustodiaEnSucursal(r io.Reader) (ComienzoCustodiaEnSucursal, error) {
	t := NewComienzoCustodiaEnSucursal()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeComienzoCustodiaEnSucursalFromSchema(r io.Reader, schema string) (ComienzoCustodiaEnSucursal, error) {
	t := NewComienzoCustodiaEnSucursal()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeComienzoCustodiaEnSucursal(r ComienzoCustodiaEnSucursal, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.CantidadDeDiasDeLaFase, w)
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
	err = writeUnionNullDatosSucursal(r.SucursalDeCustodia, w)
	if err != nil {
		return err
	}
	return err
}

func (r ComienzoCustodiaEnSucursal) Serialize(w io.Writer) error {
	return writeComienzoCustodiaEnSucursal(r, w)
}

func (r ComienzoCustodiaEnSucursal) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"name\":\"cantidadDeDiasDeLaFase\",\"type\":\"int\"},{\"default\":null,\"name\":\"motivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"submotivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalDeCustodia\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosSucursal\"]}],\"name\":\"Integracion.Esquemas.Trazas.ComienzoCustodiaEnSucursal\",\"type\":\"record\"}"
}

func (r ComienzoCustodiaEnSucursal) SchemaName() string {
	return "Integracion.Esquemas.Trazas.ComienzoCustodiaEnSucursal"
}

func (_ ComienzoCustodiaEnSucursal) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ComienzoCustodiaEnSucursal) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ComienzoCustodiaEnSucursal) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ComienzoCustodiaEnSucursal) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ComienzoCustodiaEnSucursal) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ComienzoCustodiaEnSucursal) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ComienzoCustodiaEnSucursal) SetString(v string)   { panic("Unsupported operation") }
func (_ ComienzoCustodiaEnSucursal) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ComienzoCustodiaEnSucursal) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		w := types.Int{Target: &r.CantidadDeDiasDeLaFase}

		return w

	case 2:
		r.Motivo = NewUnionNullString()

		return r.Motivo
	case 3:
		r.Submotivo = NewUnionNullString()

		return r.Submotivo
	case 4:
		r.SucursalDeCustodia = NewUnionNullDatosSucursal()

		return r.SucursalDeCustodia
	}
	panic("Unknown field index")
}

func (r *ComienzoCustodiaEnSucursal) SetDefault(i int) {
	switch i {
	case 2:
		r.Motivo = nil
		return
	case 3:
		r.Submotivo = nil
		return
	case 4:
		r.SucursalDeCustodia = nil
		return
	}
	panic("Unknown field index")
}

func (r *ComienzoCustodiaEnSucursal) NullField(i int) {
	switch i {
	case 2:
		r.Motivo = nil
		return
	case 3:
		r.Submotivo = nil
		return
	case 4:
		r.SucursalDeCustodia = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ComienzoCustodiaEnSucursal) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ComienzoCustodiaEnSucursal) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ComienzoCustodiaEnSucursal) HintSize(int)                     { panic("Unsupported operation") }
func (_ ComienzoCustodiaEnSucursal) Finalize()                        {}

func (_ ComienzoCustodiaEnSucursal) AvroCRC64Fingerprint() []byte {
	return []byte(ComienzoCustodiaEnSucursalAvroCRC64Fingerprint)
}

func (r ComienzoCustodiaEnSucursal) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["cantidadDeDiasDeLaFase"], err = json.Marshal(r.CantidadDeDiasDeLaFase)
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
	output["sucursalDeCustodia"], err = json.Marshal(r.SucursalDeCustodia)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ComienzoCustodiaEnSucursal) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["cantidadDeDiasDeLaFase"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadDeDiasDeLaFase); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cantidadDeDiasDeLaFase")
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
		if v, ok := fields["sucursalDeCustodia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SucursalDeCustodia); err != nil {
			return err
		}
	} else {
		r.SucursalDeCustodia = NewUnionNullDatosSucursal()

		r.SucursalDeCustodia = nil
	}
	return nil
}
