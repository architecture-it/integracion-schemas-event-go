// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Traza.avsc
 */
package NovedadesCustomEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Traza struct {
	Id *UnionNullString `json:"Id"`

	Evento *UnionNullString `json:"Evento"`

	NumeroDeEnvio *UnionNullString `json:"NumeroDeEnvio"`

	NumeroDeContratoInterno *UnionNullString `json:"NumeroDeContratoInterno"`

	Cuando int64 `json:"cuando"`

	CicloDelEnvio *UnionNullString `json:"CicloDelEnvio"`

	Motivo *UnionNullString `json:"Motivo"`

	SubMotivo *UnionNullString `json:"SubMotivo"`

	DatosAdicionales *UnionNullArrayMetadata `json:"DatosAdicionales"`
}

const TrazaAvroCRC64Fingerprint = "Tg\x14\xc5\x03t1\v"

func NewTraza() Traza {
	r := Traza{}
	r.Id = nil
	r.Evento = nil
	r.NumeroDeEnvio = nil
	r.NumeroDeContratoInterno = nil
	r.CicloDelEnvio = nil
	r.Motivo = nil
	r.SubMotivo = nil
	r.DatosAdicionales = nil
	return r
}

func DeserializeTraza(r io.Reader) (Traza, error) {
	t := NewTraza()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeTrazaFromSchema(r io.Reader, schema string) (Traza, error) {
	t := NewTraza()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeTraza(r Traza, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Id, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Evento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeEnvio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeContratoInterno, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Cuando, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CicloDelEnvio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Motivo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.SubMotivo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayMetadata(r.DatosAdicionales, w)
	if err != nil {
		return err
	}
	return err
}

func (r Traza) Serialize(w io.Writer) error {
	return writeTraza(r, w)
}

func (r Traza) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Evento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeContratoInterno\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"CicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Motivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SubMotivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DatosAdicionales\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"Key\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Value\",\"type\":[\"null\",\"string\"]}],\"name\":\"Metadata\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"Andreani.NovedadesCustom.Events.Common.Traza\",\"type\":\"record\"}"
}

func (r Traza) SchemaName() string {
	return "Andreani.NovedadesCustom.Events.Common.Traza"
}

func (_ Traza) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Traza) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Traza) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Traza) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Traza) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Traza) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Traza) SetString(v string)   { panic("Unsupported operation") }
func (_ Traza) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Traza) Get(i int) types.Field {
	switch i {
	case 0:
		r.Id = NewUnionNullString()

		return r.Id
	case 1:
		r.Evento = NewUnionNullString()

		return r.Evento
	case 2:
		r.NumeroDeEnvio = NewUnionNullString()

		return r.NumeroDeEnvio
	case 3:
		r.NumeroDeContratoInterno = NewUnionNullString()

		return r.NumeroDeContratoInterno
	case 4:
		w := types.Long{Target: &r.Cuando}

		return w

	case 5:
		r.CicloDelEnvio = NewUnionNullString()

		return r.CicloDelEnvio
	case 6:
		r.Motivo = NewUnionNullString()

		return r.Motivo
	case 7:
		r.SubMotivo = NewUnionNullString()

		return r.SubMotivo
	case 8:
		r.DatosAdicionales = NewUnionNullArrayMetadata()

		return r.DatosAdicionales
	}
	panic("Unknown field index")
}

func (r *Traza) SetDefault(i int) {
	switch i {
	case 0:
		r.Id = nil
		return
	case 1:
		r.Evento = nil
		return
	case 2:
		r.NumeroDeEnvio = nil
		return
	case 3:
		r.NumeroDeContratoInterno = nil
		return
	case 5:
		r.CicloDelEnvio = nil
		return
	case 6:
		r.Motivo = nil
		return
	case 7:
		r.SubMotivo = nil
		return
	case 8:
		r.DatosAdicionales = nil
		return
	}
	panic("Unknown field index")
}

func (r *Traza) NullField(i int) {
	switch i {
	case 0:
		r.Id = nil
		return
	case 1:
		r.Evento = nil
		return
	case 2:
		r.NumeroDeEnvio = nil
		return
	case 3:
		r.NumeroDeContratoInterno = nil
		return
	case 5:
		r.CicloDelEnvio = nil
		return
	case 6:
		r.Motivo = nil
		return
	case 7:
		r.SubMotivo = nil
		return
	case 8:
		r.DatosAdicionales = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Traza) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Traza) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Traza) HintSize(int)                     { panic("Unsupported operation") }
func (_ Traza) Finalize()                        {}

func (_ Traza) AvroCRC64Fingerprint() []byte {
	return []byte(TrazaAvroCRC64Fingerprint)
}

func (r Traza) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["Evento"], err = json.Marshal(r.Evento)
	if err != nil {
		return nil, err
	}
	output["NumeroDeEnvio"], err = json.Marshal(r.NumeroDeEnvio)
	if err != nil {
		return nil, err
	}
	output["NumeroDeContratoInterno"], err = json.Marshal(r.NumeroDeContratoInterno)
	if err != nil {
		return nil, err
	}
	output["cuando"], err = json.Marshal(r.Cuando)
	if err != nil {
		return nil, err
	}
	output["CicloDelEnvio"], err = json.Marshal(r.CicloDelEnvio)
	if err != nil {
		return nil, err
	}
	output["Motivo"], err = json.Marshal(r.Motivo)
	if err != nil {
		return nil, err
	}
	output["SubMotivo"], err = json.Marshal(r.SubMotivo)
	if err != nil {
		return nil, err
	}
	output["DatosAdicionales"], err = json.Marshal(r.DatosAdicionales)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Traza) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		r.Id = NewUnionNullString()

		r.Id = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Evento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Evento); err != nil {
			return err
		}
	} else {
		r.Evento = NewUnionNullString()

		r.Evento = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeEnvio); err != nil {
			return err
		}
	} else {
		r.NumeroDeEnvio = NewUnionNullString()

		r.NumeroDeEnvio = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeContratoInterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeContratoInterno); err != nil {
			return err
		}
	} else {
		r.NumeroDeContratoInterno = NewUnionNullString()

		r.NumeroDeContratoInterno = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["cuando"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cuando); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cuando")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CicloDelEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CicloDelEnvio); err != nil {
			return err
		}
	} else {
		r.CicloDelEnvio = NewUnionNullString()

		r.CicloDelEnvio = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Motivo"]; ok {
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
		if v, ok := fields["SubMotivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SubMotivo); err != nil {
			return err
		}
	} else {
		r.SubMotivo = NewUnionNullString()

		r.SubMotivo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DatosAdicionales"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DatosAdicionales); err != nil {
			return err
		}
	} else {
		r.DatosAdicionales = NewUnionNullArrayMetadata()

		r.DatosAdicionales = nil
	}
	return nil
}
