// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TransactionEvent.avsc
 */
package AltaOrdenEnvioEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Contrato struct {
	ContratoTms *UnionNullString `json:"contratoTms"`

	ClienteTmsCodigo *UnionNullString `json:"clienteTmsCodigo"`

	Codigo *UnionNullString `json:"codigo"`

	Descripcion *UnionNullString `json:"Descripcion"`

	SolicitanteCodigo *UnionNullString `json:"SolicitanteCodigo"`

	SolicitanteDescripcion *UnionNullString `json:"SolicitanteDescripcion"`

	DestinatarioCodigo *UnionNullString `json:"DestinatarioCodigo"`

	DestinatarioDescripcion *UnionNullString `json:"DestinatarioDescripcion"`

	DestinatarioNombreCorto *UnionNullString `json:"DestinatarioNombreCorto"`

	Segmento *UnionNullString `json:"Segmento"`

	PlazoEntrega *UnionNullString `json:"PlazoEntrega"`

	TipoEntrega *UnionNullString `json:"TipoEntrega"`

	TipoServicio *UnionNullString `json:"TipoServicio"`

	TipoServicioCodigo *UnionNullString `json:"TipoServicioCodigo"`

	InsertTS *UnionNullString `json:"InsertTS"`

	Cuit *UnionNullString `json:"Cuit"`

	Canal *UnionNullString `json:"Canal"`

	SucursalRendicion *UnionNullString `json:"SucursalRendicion"`
}

const ContratoAvroCRC64Fingerprint = "\xceML\xab\x9a\xca\x05\x00"

func NewContrato() Contrato {
	r := Contrato{}
	return r
}

func DeserializeContrato(r io.Reader) (Contrato, error) {
	t := NewContrato()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeContratoFromSchema(r io.Reader, schema string) (Contrato, error) {
	t := NewContrato()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeContrato(r Contrato, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.ContratoTms, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ClienteTmsCodigo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Codigo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Descripcion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.SolicitanteCodigo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.SolicitanteDescripcion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioCodigo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioDescripcion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioNombreCorto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Segmento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.PlazoEntrega, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoEntrega, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoServicio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoServicioCodigo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.InsertTS, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Cuit, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Canal, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.SucursalRendicion, w)
	if err != nil {
		return err
	}
	return err
}

func (r Contrato) Serialize(w io.Writer) error {
	return writeContrato(r, w)
}

func (r Contrato) Schema() string {
	return "{\"fields\":[{\"name\":\"contratoTms\",\"type\":[\"null\",\"string\"]},{\"name\":\"clienteTmsCodigo\",\"type\":[\"null\",\"string\"]},{\"name\":\"codigo\",\"type\":[\"null\",\"string\"]},{\"name\":\"Descripcion\",\"type\":[\"null\",\"string\"]},{\"name\":\"SolicitanteCodigo\",\"type\":[\"null\",\"string\"]},{\"name\":\"SolicitanteDescripcion\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioCodigo\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioDescripcion\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioNombreCorto\",\"type\":[\"null\",\"string\"]},{\"name\":\"Segmento\",\"type\":[\"null\",\"string\"]},{\"name\":\"PlazoEntrega\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoEntrega\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoServicio\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoServicioCodigo\",\"type\":[\"null\",\"string\"]},{\"name\":\"InsertTS\",\"type\":[\"null\",\"string\"]},{\"name\":\"Cuit\",\"type\":[\"null\",\"string\"]},{\"name\":\"Canal\",\"type\":[\"null\",\"string\"]},{\"name\":\"SucursalRendicion\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.AltaOrdenEnvio.Events.Common.Contrato\",\"type\":\"record\"}"
}

func (r Contrato) SchemaName() string {
	return "Andreani.AltaOrdenEnvio.Events.Common.Contrato"
}

func (_ Contrato) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Contrato) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Contrato) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Contrato) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Contrato) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Contrato) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Contrato) SetString(v string)   { panic("Unsupported operation") }
func (_ Contrato) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Contrato) Get(i int) types.Field {
	switch i {
	case 0:
		r.ContratoTms = NewUnionNullString()

		return r.ContratoTms
	case 1:
		r.ClienteTmsCodigo = NewUnionNullString()

		return r.ClienteTmsCodigo
	case 2:
		r.Codigo = NewUnionNullString()

		return r.Codigo
	case 3:
		r.Descripcion = NewUnionNullString()

		return r.Descripcion
	case 4:
		r.SolicitanteCodigo = NewUnionNullString()

		return r.SolicitanteCodigo
	case 5:
		r.SolicitanteDescripcion = NewUnionNullString()

		return r.SolicitanteDescripcion
	case 6:
		r.DestinatarioCodigo = NewUnionNullString()

		return r.DestinatarioCodigo
	case 7:
		r.DestinatarioDescripcion = NewUnionNullString()

		return r.DestinatarioDescripcion
	case 8:
		r.DestinatarioNombreCorto = NewUnionNullString()

		return r.DestinatarioNombreCorto
	case 9:
		r.Segmento = NewUnionNullString()

		return r.Segmento
	case 10:
		r.PlazoEntrega = NewUnionNullString()

		return r.PlazoEntrega
	case 11:
		r.TipoEntrega = NewUnionNullString()

		return r.TipoEntrega
	case 12:
		r.TipoServicio = NewUnionNullString()

		return r.TipoServicio
	case 13:
		r.TipoServicioCodigo = NewUnionNullString()

		return r.TipoServicioCodigo
	case 14:
		r.InsertTS = NewUnionNullString()

		return r.InsertTS
	case 15:
		r.Cuit = NewUnionNullString()

		return r.Cuit
	case 16:
		r.Canal = NewUnionNullString()

		return r.Canal
	case 17:
		r.SucursalRendicion = NewUnionNullString()

		return r.SucursalRendicion
	}
	panic("Unknown field index")
}

func (r *Contrato) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Contrato) NullField(i int) {
	switch i {
	case 0:
		r.ContratoTms = nil
		return
	case 1:
		r.ClienteTmsCodigo = nil
		return
	case 2:
		r.Codigo = nil
		return
	case 3:
		r.Descripcion = nil
		return
	case 4:
		r.SolicitanteCodigo = nil
		return
	case 5:
		r.SolicitanteDescripcion = nil
		return
	case 6:
		r.DestinatarioCodigo = nil
		return
	case 7:
		r.DestinatarioDescripcion = nil
		return
	case 8:
		r.DestinatarioNombreCorto = nil
		return
	case 9:
		r.Segmento = nil
		return
	case 10:
		r.PlazoEntrega = nil
		return
	case 11:
		r.TipoEntrega = nil
		return
	case 12:
		r.TipoServicio = nil
		return
	case 13:
		r.TipoServicioCodigo = nil
		return
	case 14:
		r.InsertTS = nil
		return
	case 15:
		r.Cuit = nil
		return
	case 16:
		r.Canal = nil
		return
	case 17:
		r.SucursalRendicion = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Contrato) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Contrato) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Contrato) HintSize(int)                     { panic("Unsupported operation") }
func (_ Contrato) Finalize()                        {}

func (_ Contrato) AvroCRC64Fingerprint() []byte {
	return []byte(ContratoAvroCRC64Fingerprint)
}

func (r Contrato) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["contratoTms"], err = json.Marshal(r.ContratoTms)
	if err != nil {
		return nil, err
	}
	output["clienteTmsCodigo"], err = json.Marshal(r.ClienteTmsCodigo)
	if err != nil {
		return nil, err
	}
	output["codigo"], err = json.Marshal(r.Codigo)
	if err != nil {
		return nil, err
	}
	output["Descripcion"], err = json.Marshal(r.Descripcion)
	if err != nil {
		return nil, err
	}
	output["SolicitanteCodigo"], err = json.Marshal(r.SolicitanteCodigo)
	if err != nil {
		return nil, err
	}
	output["SolicitanteDescripcion"], err = json.Marshal(r.SolicitanteDescripcion)
	if err != nil {
		return nil, err
	}
	output["DestinatarioCodigo"], err = json.Marshal(r.DestinatarioCodigo)
	if err != nil {
		return nil, err
	}
	output["DestinatarioDescripcion"], err = json.Marshal(r.DestinatarioDescripcion)
	if err != nil {
		return nil, err
	}
	output["DestinatarioNombreCorto"], err = json.Marshal(r.DestinatarioNombreCorto)
	if err != nil {
		return nil, err
	}
	output["Segmento"], err = json.Marshal(r.Segmento)
	if err != nil {
		return nil, err
	}
	output["PlazoEntrega"], err = json.Marshal(r.PlazoEntrega)
	if err != nil {
		return nil, err
	}
	output["TipoEntrega"], err = json.Marshal(r.TipoEntrega)
	if err != nil {
		return nil, err
	}
	output["TipoServicio"], err = json.Marshal(r.TipoServicio)
	if err != nil {
		return nil, err
	}
	output["TipoServicioCodigo"], err = json.Marshal(r.TipoServicioCodigo)
	if err != nil {
		return nil, err
	}
	output["InsertTS"], err = json.Marshal(r.InsertTS)
	if err != nil {
		return nil, err
	}
	output["Cuit"], err = json.Marshal(r.Cuit)
	if err != nil {
		return nil, err
	}
	output["Canal"], err = json.Marshal(r.Canal)
	if err != nil {
		return nil, err
	}
	output["SucursalRendicion"], err = json.Marshal(r.SucursalRendicion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Contrato) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["contratoTms"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ContratoTms); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for contratoTms")
	}
	val = func() json.RawMessage {
		if v, ok := fields["clienteTmsCodigo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ClienteTmsCodigo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for clienteTmsCodigo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Codigo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Descripcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Descripcion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Descripcion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SolicitanteCodigo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SolicitanteCodigo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SolicitanteCodigo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SolicitanteDescripcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SolicitanteDescripcion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SolicitanteDescripcion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioCodigo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioCodigo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioCodigo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioDescripcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioDescripcion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioDescripcion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioNombreCorto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioNombreCorto); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioNombreCorto")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Segmento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Segmento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Segmento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PlazoEntrega"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PlazoEntrega); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PlazoEntrega")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoEntrega"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoEntrega); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoEntrega")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoServicio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoServicio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoServicio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoServicioCodigo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoServicioCodigo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoServicioCodigo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["InsertTS"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.InsertTS); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for InsertTS")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Cuit"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cuit); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Cuit")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Canal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Canal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Canal")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SucursalRendicion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SucursalRendicion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SucursalRendicion")
	}
	return nil
}
