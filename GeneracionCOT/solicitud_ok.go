// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     SolicitudOk.avsc
 */
package GeneracionCOTEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type SolicitudOk struct {
	FechaHora int64 `json:"fechaHora"`

	Cuit string `json:"cuit"`

	NumeroComprobante int32 `json:"numeroComprobante"`

	NombreArchivo string `json:"nombreArchivo"`

	CodigoIntegridad string `json:"codigoIntegridad"`

	NumeroUnico string `json:"numeroUnico"`

	Procesado string `json:"procesado"`

	Cot string `json:"cot"`
}

const SolicitudOkAvroCRC64Fingerprint = "\xf2N}\xa6\x8d\x80\xac\b"

func NewSolicitudOk() SolicitudOk {
	r := SolicitudOk{}
	return r
}

func DeserializeSolicitudOk(r io.Reader) (SolicitudOk, error) {
	t := NewSolicitudOk()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSolicitudOkFromSchema(r io.Reader, schema string) (SolicitudOk, error) {
	t := NewSolicitudOk()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSolicitudOk(r SolicitudOk, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.FechaHora, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Cuit, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.NumeroComprobante, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NombreArchivo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoIntegridad, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroUnico, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Procesado, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Cot, w)
	if err != nil {
		return err
	}
	return err
}

func (r SolicitudOk) Serialize(w io.Writer) error {
	return writeSolicitudOk(r, w)
}

func (r SolicitudOk) Schema() string {
	return "{\"fields\":[{\"name\":\"fechaHora\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"cuit\",\"type\":\"string\"},{\"name\":\"numeroComprobante\",\"type\":\"int\"},{\"name\":\"nombreArchivo\",\"type\":\"string\"},{\"name\":\"codigoIntegridad\",\"type\":\"string\"},{\"name\":\"numeroUnico\",\"type\":\"string\"},{\"name\":\"procesado\",\"type\":\"string\"},{\"name\":\"cot\",\"type\":\"string\"}],\"name\":\"Andreani.GeneracionCOT.Events.Record.SolicitudOk\",\"type\":\"record\"}"
}

func (r SolicitudOk) SchemaName() string {
	return "Andreani.GeneracionCOT.Events.Record.SolicitudOk"
}

func (_ SolicitudOk) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ SolicitudOk) SetInt(v int32)       { panic("Unsupported operation") }
func (_ SolicitudOk) SetLong(v int64)      { panic("Unsupported operation") }
func (_ SolicitudOk) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ SolicitudOk) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ SolicitudOk) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ SolicitudOk) SetString(v string)   { panic("Unsupported operation") }
func (_ SolicitudOk) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *SolicitudOk) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.FechaHora}

		return w

	case 1:
		w := types.String{Target: &r.Cuit}

		return w

	case 2:
		w := types.Int{Target: &r.NumeroComprobante}

		return w

	case 3:
		w := types.String{Target: &r.NombreArchivo}

		return w

	case 4:
		w := types.String{Target: &r.CodigoIntegridad}

		return w

	case 5:
		w := types.String{Target: &r.NumeroUnico}

		return w

	case 6:
		w := types.String{Target: &r.Procesado}

		return w

	case 7:
		w := types.String{Target: &r.Cot}

		return w

	}
	panic("Unknown field index")
}

func (r *SolicitudOk) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *SolicitudOk) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ SolicitudOk) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ SolicitudOk) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ SolicitudOk) HintSize(int)                     { panic("Unsupported operation") }
func (_ SolicitudOk) Finalize()                        {}

func (_ SolicitudOk) AvroCRC64Fingerprint() []byte {
	return []byte(SolicitudOkAvroCRC64Fingerprint)
}

func (r SolicitudOk) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["fechaHora"], err = json.Marshal(r.FechaHora)
	if err != nil {
		return nil, err
	}
	output["cuit"], err = json.Marshal(r.Cuit)
	if err != nil {
		return nil, err
	}
	output["numeroComprobante"], err = json.Marshal(r.NumeroComprobante)
	if err != nil {
		return nil, err
	}
	output["nombreArchivo"], err = json.Marshal(r.NombreArchivo)
	if err != nil {
		return nil, err
	}
	output["codigoIntegridad"], err = json.Marshal(r.CodigoIntegridad)
	if err != nil {
		return nil, err
	}
	output["numeroUnico"], err = json.Marshal(r.NumeroUnico)
	if err != nil {
		return nil, err
	}
	output["procesado"], err = json.Marshal(r.Procesado)
	if err != nil {
		return nil, err
	}
	output["cot"], err = json.Marshal(r.Cot)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *SolicitudOk) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["fechaHora"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaHora); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fechaHora")
	}
	val = func() json.RawMessage {
		if v, ok := fields["cuit"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cuit); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cuit")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroComprobante"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroComprobante); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroComprobante")
	}
	val = func() json.RawMessage {
		if v, ok := fields["nombreArchivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NombreArchivo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for nombreArchivo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoIntegridad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoIntegridad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigoIntegridad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroUnico"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroUnico); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroUnico")
	}
	val = func() json.RawMessage {
		if v, ok := fields["procesado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Procesado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for procesado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["cot"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cot); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cot")
	}
	return nil
}
