// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     FinCustodiaEnSucursal.avsc
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

type FinCustodiaEnSucursal struct {
	Traza Traza `json:"traza"`

	CantidadDeDias int32 `json:"cantidadDeDias"`

	Motivo *UnionNullString `json:"motivo"`

	Submotivo *UnionNullString `json:"submotivo"`

	SucursalDeDistribucion *UnionNullDatosSucursal `json:"sucursalDeDistribucion"`
}

const FinCustodiaEnSucursalAvroCRC64Fingerprint = "\x1b\xfct\x96\xf9#7\x92"

func NewFinCustodiaEnSucursal() FinCustodiaEnSucursal {
	r := FinCustodiaEnSucursal{}
	r.Traza = NewTraza()

	r.Motivo = nil
	r.Submotivo = nil
	r.SucursalDeDistribucion = nil
	return r
}

func DeserializeFinCustodiaEnSucursal(r io.Reader) (FinCustodiaEnSucursal, error) {
	t := NewFinCustodiaEnSucursal()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeFinCustodiaEnSucursalFromSchema(r io.Reader, schema string) (FinCustodiaEnSucursal, error) {
	t := NewFinCustodiaEnSucursal()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeFinCustodiaEnSucursal(r FinCustodiaEnSucursal, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.CantidadDeDias, w)
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
	err = writeUnionNullDatosSucursal(r.SucursalDeDistribucion, w)
	if err != nil {
		return err
	}
	return err
}

func (r FinCustodiaEnSucursal) Serialize(w io.Writer) error {
	return writeFinCustodiaEnSucursal(r, w)
}

func (r FinCustodiaEnSucursal) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"name\":\"cantidadDeDias\",\"type\":\"int\"},{\"default\":null,\"name\":\"motivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"submotivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalDeDistribucion\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosSucursal\"]}],\"name\":\"Integracion.Esquemas.Trazas.FinCustodiaEnSucursal\",\"type\":\"record\"}"
}

func (r FinCustodiaEnSucursal) SchemaName() string {
	return "Integracion.Esquemas.Trazas.FinCustodiaEnSucursal"
}

func (_ FinCustodiaEnSucursal) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ FinCustodiaEnSucursal) SetInt(v int32)       { panic("Unsupported operation") }
func (_ FinCustodiaEnSucursal) SetLong(v int64)      { panic("Unsupported operation") }
func (_ FinCustodiaEnSucursal) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ FinCustodiaEnSucursal) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ FinCustodiaEnSucursal) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ FinCustodiaEnSucursal) SetString(v string)   { panic("Unsupported operation") }
func (_ FinCustodiaEnSucursal) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *FinCustodiaEnSucursal) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		w := types.Int{Target: &r.CantidadDeDias}

		return w

	case 2:
		r.Motivo = NewUnionNullString()

		return r.Motivo
	case 3:
		r.Submotivo = NewUnionNullString()

		return r.Submotivo
	case 4:
		r.SucursalDeDistribucion = NewUnionNullDatosSucursal()

		return r.SucursalDeDistribucion
	}
	panic("Unknown field index")
}

func (r *FinCustodiaEnSucursal) SetDefault(i int) {
	switch i {
	case 2:
		r.Motivo = nil
		return
	case 3:
		r.Submotivo = nil
		return
	case 4:
		r.SucursalDeDistribucion = nil
		return
	}
	panic("Unknown field index")
}

func (r *FinCustodiaEnSucursal) NullField(i int) {
	switch i {
	case 2:
		r.Motivo = nil
		return
	case 3:
		r.Submotivo = nil
		return
	case 4:
		r.SucursalDeDistribucion = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ FinCustodiaEnSucursal) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ FinCustodiaEnSucursal) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ FinCustodiaEnSucursal) HintSize(int)                     { panic("Unsupported operation") }
func (_ FinCustodiaEnSucursal) Finalize()                        {}

func (_ FinCustodiaEnSucursal) AvroCRC64Fingerprint() []byte {
	return []byte(FinCustodiaEnSucursalAvroCRC64Fingerprint)
}

func (r FinCustodiaEnSucursal) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["cantidadDeDias"], err = json.Marshal(r.CantidadDeDias)
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
	output["sucursalDeDistribucion"], err = json.Marshal(r.SucursalDeDistribucion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *FinCustodiaEnSucursal) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["cantidadDeDias"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadDeDias); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cantidadDeDias")
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
		if v, ok := fields["sucursalDeDistribucion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SucursalDeDistribucion); err != nil {
			return err
		}
	} else {
		r.SucursalDeDistribucion = NewUnionNullDatosSucursal()

		r.SucursalDeDistribucion = nil
	}
	return nil
}
