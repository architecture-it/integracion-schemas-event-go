// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ResponseCotizar.avsc
 */
package CotizarBackendEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ResponseCotizar struct {
	TarifaConIva *UnionNullTarifa `json:"TarifaConIva"`

	TarifaSinIva *UnionNullTarifa `json:"TarifaSinIva"`

	ContratoId *UnionNullString `json:"ContratoId"`

	PesoAforado *UnionNullString `json:"pesoAforado"`

	ModoDeEntrega *UnionNullString `json:"ModoDeEntrega"`

	ModoDeEntregaId *UnionNullString `json:"ModoDeEntregaId"`
}

const ResponseCotizarAvroCRC64Fingerprint = "S|\x87\xcd\xd9\xcfU\xce"

func NewResponseCotizar() ResponseCotizar {
	r := ResponseCotizar{}
	r.TarifaConIva = nil
	r.TarifaSinIva = nil
	r.ContratoId = nil
	r.PesoAforado = nil
	r.ModoDeEntrega = nil
	r.ModoDeEntregaId = nil
	return r
}

func DeserializeResponseCotizar(r io.Reader) (ResponseCotizar, error) {
	t := NewResponseCotizar()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeResponseCotizarFromSchema(r io.Reader, schema string) (ResponseCotizar, error) {
	t := NewResponseCotizar()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeResponseCotizar(r ResponseCotizar, w io.Writer) error {
	var err error
	err = writeUnionNullTarifa(r.TarifaConIva, w)
	if err != nil {
		return err
	}
	err = writeUnionNullTarifa(r.TarifaSinIva, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ContratoId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.PesoAforado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ModoDeEntrega, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ModoDeEntregaId, w)
	if err != nil {
		return err
	}
	return err
}

func (r ResponseCotizar) Serialize(w io.Writer) error {
	return writeResponseCotizar(r, w)
}

func (r ResponseCotizar) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"TarifaConIva\",\"type\":[\"null\",{\"fields\":[{\"name\":\"SeguroDistribucion\",\"type\":\"string\"},{\"name\":\"Distribucion\",\"type\":\"string\"},{\"name\":\"Total\",\"type\":\"string\"}],\"name\":\"Tarifa\",\"namespace\":\"Andreani.CotizarBackend.Events.Cotizar.Common\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"TarifaSinIva\",\"type\":[\"null\",\"Andreani.CotizarBackend.Events.Cotizar.Common.Tarifa\"]},{\"default\":null,\"name\":\"ContratoId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pesoAforado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ModoDeEntrega\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ModoDeEntregaId\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.CotizarBackend.Events.Record.ResponseCotizar\",\"type\":\"record\"}"
}

func (r ResponseCotizar) SchemaName() string {
	return "Andreani.CotizarBackend.Events.Record.ResponseCotizar"
}

func (_ ResponseCotizar) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ResponseCotizar) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ResponseCotizar) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ResponseCotizar) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ResponseCotizar) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ResponseCotizar) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ResponseCotizar) SetString(v string)   { panic("Unsupported operation") }
func (_ ResponseCotizar) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ResponseCotizar) Get(i int) types.Field {
	switch i {
	case 0:
		r.TarifaConIva = NewUnionNullTarifa()

		return r.TarifaConIva
	case 1:
		r.TarifaSinIva = NewUnionNullTarifa()

		return r.TarifaSinIva
	case 2:
		r.ContratoId = NewUnionNullString()

		return r.ContratoId
	case 3:
		r.PesoAforado = NewUnionNullString()

		return r.PesoAforado
	case 4:
		r.ModoDeEntrega = NewUnionNullString()

		return r.ModoDeEntrega
	case 5:
		r.ModoDeEntregaId = NewUnionNullString()

		return r.ModoDeEntregaId
	}
	panic("Unknown field index")
}

func (r *ResponseCotizar) SetDefault(i int) {
	switch i {
	case 0:
		r.TarifaConIva = nil
		return
	case 1:
		r.TarifaSinIva = nil
		return
	case 2:
		r.ContratoId = nil
		return
	case 3:
		r.PesoAforado = nil
		return
	case 4:
		r.ModoDeEntrega = nil
		return
	case 5:
		r.ModoDeEntregaId = nil
		return
	}
	panic("Unknown field index")
}

func (r *ResponseCotizar) NullField(i int) {
	switch i {
	case 0:
		r.TarifaConIva = nil
		return
	case 1:
		r.TarifaSinIva = nil
		return
	case 2:
		r.ContratoId = nil
		return
	case 3:
		r.PesoAforado = nil
		return
	case 4:
		r.ModoDeEntrega = nil
		return
	case 5:
		r.ModoDeEntregaId = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ResponseCotizar) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ResponseCotizar) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ResponseCotizar) HintSize(int)                     { panic("Unsupported operation") }
func (_ ResponseCotizar) Finalize()                        {}

func (_ ResponseCotizar) AvroCRC64Fingerprint() []byte {
	return []byte(ResponseCotizarAvroCRC64Fingerprint)
}

func (r ResponseCotizar) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["TarifaConIva"], err = json.Marshal(r.TarifaConIva)
	if err != nil {
		return nil, err
	}
	output["TarifaSinIva"], err = json.Marshal(r.TarifaSinIva)
	if err != nil {
		return nil, err
	}
	output["ContratoId"], err = json.Marshal(r.ContratoId)
	if err != nil {
		return nil, err
	}
	output["pesoAforado"], err = json.Marshal(r.PesoAforado)
	if err != nil {
		return nil, err
	}
	output["ModoDeEntrega"], err = json.Marshal(r.ModoDeEntrega)
	if err != nil {
		return nil, err
	}
	output["ModoDeEntregaId"], err = json.Marshal(r.ModoDeEntregaId)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ResponseCotizar) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["TarifaConIva"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TarifaConIva); err != nil {
			return err
		}
	} else {
		r.TarifaConIva = NewUnionNullTarifa()

		r.TarifaConIva = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["TarifaSinIva"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TarifaSinIva); err != nil {
			return err
		}
	} else {
		r.TarifaSinIva = NewUnionNullTarifa()

		r.TarifaSinIva = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ContratoId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ContratoId); err != nil {
			return err
		}
	} else {
		r.ContratoId = NewUnionNullString()

		r.ContratoId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["pesoAforado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PesoAforado); err != nil {
			return err
		}
	} else {
		r.PesoAforado = NewUnionNullString()

		r.PesoAforado = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ModoDeEntrega"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ModoDeEntrega); err != nil {
			return err
		}
	} else {
		r.ModoDeEntrega = NewUnionNullString()

		r.ModoDeEntrega = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ModoDeEntregaId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ModoDeEntregaId); err != nil {
			return err
		}
	} else {
		r.ModoDeEntregaId = NewUnionNullString()

		r.ModoDeEntregaId = nil
	}
	return nil
}
