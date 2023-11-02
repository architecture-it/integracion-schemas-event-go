// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CamaraPacking.avsc
 */
package appnameEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type CamaraPacking struct {
	Mesa *UnionNullString `json:"Mesa"`

	OrdenWh *UnionNullString `json:"OrdenWh"`

	OrdenCliente *UnionNullString `json:"OrdenCliente"`

	Propietario *UnionNullString `json:"Propietario"`

	Estado *UnionNullString `json:"Estado"`
}

const CamaraPackingAvroCRC64Fingerprint = "4\x99\xc0OsA\xf60"

func NewCamaraPacking() CamaraPacking {
	r := CamaraPacking{}
	return r
}

func DeserializeCamaraPacking(r io.Reader) (CamaraPacking, error) {
	t := NewCamaraPacking()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCamaraPackingFromSchema(r io.Reader, schema string) (CamaraPacking, error) {
	t := NewCamaraPacking()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCamaraPacking(r CamaraPacking, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Mesa, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.OrdenWh, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.OrdenCliente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Estado, w)
	if err != nil {
		return err
	}
	return err
}

func (r CamaraPacking) Serialize(w io.Writer) error {
	return writeCamaraPacking(r, w)
}

func (r CamaraPacking) Schema() string {
	return "{\"fields\":[{\"name\":\"Mesa\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrdenWh\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrdenCliente\",\"type\":[\"null\",\"string\"]},{\"name\":\"Propietario\",\"type\":[\"null\",\"string\"]},{\"name\":\"Estado\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.Camara.Events.Record.CamaraPacking\",\"type\":\"record\"}"
}

func (r CamaraPacking) SchemaName() string {
	return "Andreani.Camara.Events.Record.CamaraPacking"
}

func (_ CamaraPacking) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CamaraPacking) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CamaraPacking) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CamaraPacking) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CamaraPacking) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CamaraPacking) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CamaraPacking) SetString(v string)   { panic("Unsupported operation") }
func (_ CamaraPacking) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CamaraPacking) Get(i int) types.Field {
	switch i {
	case 0:
		r.Mesa = NewUnionNullString()

		return r.Mesa
	case 1:
		r.OrdenWh = NewUnionNullString()

		return r.OrdenWh
	case 2:
		r.OrdenCliente = NewUnionNullString()

		return r.OrdenCliente
	case 3:
		r.Propietario = NewUnionNullString()

		return r.Propietario
	case 4:
		r.Estado = NewUnionNullString()

		return r.Estado
	}
	panic("Unknown field index")
}

func (r *CamaraPacking) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *CamaraPacking) NullField(i int) {
	switch i {
	case 0:
		r.Mesa = nil
		return
	case 1:
		r.OrdenWh = nil
		return
	case 2:
		r.OrdenCliente = nil
		return
	case 3:
		r.Propietario = nil
		return
	case 4:
		r.Estado = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ CamaraPacking) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CamaraPacking) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CamaraPacking) HintSize(int)                     { panic("Unsupported operation") }
func (_ CamaraPacking) Finalize()                        {}

func (_ CamaraPacking) AvroCRC64Fingerprint() []byte {
	return []byte(CamaraPackingAvroCRC64Fingerprint)
}

func (r CamaraPacking) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Mesa"], err = json.Marshal(r.Mesa)
	if err != nil {
		return nil, err
	}
	output["OrdenWh"], err = json.Marshal(r.OrdenWh)
	if err != nil {
		return nil, err
	}
	output["OrdenCliente"], err = json.Marshal(r.OrdenCliente)
	if err != nil {
		return nil, err
	}
	output["Propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["Estado"], err = json.Marshal(r.Estado)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CamaraPacking) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Mesa"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Mesa); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Mesa")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OrdenWh"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenWh); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrdenWh")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OrdenCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenCliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrdenCliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Propietario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Propietario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Propietario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Estado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Estado")
	}
	return nil
}
