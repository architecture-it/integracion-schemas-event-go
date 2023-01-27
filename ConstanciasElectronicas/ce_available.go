// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CeAvailable.avsc
 */
package ConstanciasElectronicasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type CeAvailable struct {
	Traza Traza `json:"traza"`

	Contrato Contrato `json:"contrato"`

	Template Template `json:"template"`

	Constancia Constancia `json:"constancia"`
}

const CeAvailableAvroCRC64Fingerprint = "\x8cyC\x17!\x85\xf3\xc8"

func NewCeAvailable() CeAvailable {
	r := CeAvailable{}
	r.Traza = NewTraza()

	r.Contrato = NewContrato()

	r.Template = NewTemplate()

	r.Constancia = NewConstancia()

	return r
}

func DeserializeCeAvailable(r io.Reader) (CeAvailable, error) {
	t := NewCeAvailable()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCeAvailableFromSchema(r io.Reader, schema string) (CeAvailable, error) {
	t := NewCeAvailable()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCeAvailable(r CeAvailable, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = writeContrato(r.Contrato, w)
	if err != nil {
		return err
	}
	err = writeTemplate(r.Template, w)
	if err != nil {
		return err
	}
	err = writeConstancia(r.Constancia, w)
	if err != nil {
		return err
	}
	return err
}

func (r CeAvailable) Serialize(w io.Writer) error {
	return writeCeAvailable(r, w)
}

func (r CeAvailable) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}}],\"name\":\"Traza\",\"namespace\":\"Andreani.ConstanciasElectronicas.Common\",\"type\":\"record\"}},{\"name\":\"contrato\",\"type\":{\"fields\":[{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"name\":\"codigoDeClienteInterno\",\"type\":\"string\"},{\"name\":\"status\",\"type\":{\"name\":\"EntityStatus\",\"symbols\":[\"Inactive\",\"Active\",\"Expired\",\"Deleted\"],\"type\":\"enum\"}}],\"name\":\"Contrato\",\"namespace\":\"Andreani.ConstanciasElectronicas.Common\",\"type\":\"record\"}},{\"name\":\"template\",\"type\":{\"fields\":[{\"name\":\"type\",\"type\":{\"name\":\"TemplateType\",\"symbols\":[\"Undefined\",\"Delivery\",\"NoDelivered\"],\"type\":\"enum\"}},{\"name\":\"modules\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"Andreani.ConstanciasElectronicas.Common.EntityStatus\"}],\"name\":\"Template\",\"namespace\":\"Andreani.ConstanciasElectronicas.Common\",\"type\":\"record\"}},{\"name\":\"constancia\",\"type\":{\"fields\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"url\",\"type\":\"string\"},{\"name\":\"guid\",\"type\":\"string\"},{\"name\":\"attemps\",\"type\":\"int\"}],\"name\":\"Constancia\",\"namespace\":\"Andreani.ConstanciasElectronicas.Common\",\"type\":\"record\"}}],\"name\":\"Andreani.ConstanciasElectronicas.CeAvailable\",\"type\":\"record\"}"
}

func (r CeAvailable) SchemaName() string {
	return "Andreani.ConstanciasElectronicas.CeAvailable"
}

func (_ CeAvailable) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CeAvailable) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CeAvailable) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CeAvailable) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CeAvailable) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CeAvailable) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CeAvailable) SetString(v string)   { panic("Unsupported operation") }
func (_ CeAvailable) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CeAvailable) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		r.Contrato = NewContrato()

		w := types.Record{Target: &r.Contrato}

		return w

	case 2:
		r.Template = NewTemplate()

		w := types.Record{Target: &r.Template}

		return w

	case 3:
		r.Constancia = NewConstancia()

		w := types.Record{Target: &r.Constancia}

		return w

	}
	panic("Unknown field index")
}

func (r *CeAvailable) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *CeAvailable) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ CeAvailable) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CeAvailable) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CeAvailable) HintSize(int)                     { panic("Unsupported operation") }
func (_ CeAvailable) Finalize()                        {}

func (_ CeAvailable) AvroCRC64Fingerprint() []byte {
	return []byte(CeAvailableAvroCRC64Fingerprint)
}

func (r CeAvailable) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["contrato"], err = json.Marshal(r.Contrato)
	if err != nil {
		return nil, err
	}
	output["template"], err = json.Marshal(r.Template)
	if err != nil {
		return nil, err
	}
	output["constancia"], err = json.Marshal(r.Constancia)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CeAvailable) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["contrato"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contrato); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for contrato")
	}
	val = func() json.RawMessage {
		if v, ok := fields["template"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Template); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for template")
	}
	val = func() json.RawMessage {
		if v, ok := fields["constancia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Constancia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for constancia")
	}
	return nil
}
