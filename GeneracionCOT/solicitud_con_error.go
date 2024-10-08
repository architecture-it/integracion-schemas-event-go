// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     SolicitudConError.avsc
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

type SolicitudConError struct {
	Id int32 `json:"Id"`

	FechaHora int64 `json:"FechaHora"`

	Cuit string `json:"Cuit"`

	NumeroComprobante int32 `json:"NumeroComprobante"`

	NombreArchivo string `json:"NombreArchivo"`

	CodigoIntegridad string `json:"CodigoIntegridad"`

	NumeroUnico string `json:"NumeroUnico"`

	Procesado string `json:"Procesado"`

	Errores []Error `json:"Errores"`
}

const SolicitudConErrorAvroCRC64Fingerprint = ")s\xe91\x12\xac_y"

func NewSolicitudConError() SolicitudConError {
	r := SolicitudConError{}
	r.Errores = make([]Error, 0)

	return r
}

func DeserializeSolicitudConError(r io.Reader) (SolicitudConError, error) {
	t := NewSolicitudConError()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSolicitudConErrorFromSchema(r io.Reader, schema string) (SolicitudConError, error) {
	t := NewSolicitudConError()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSolicitudConError(r SolicitudConError, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Id, w)
	if err != nil {
		return err
	}
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
	err = writeArrayError(r.Errores, w)
	if err != nil {
		return err
	}
	return err
}

func (r SolicitudConError) Serialize(w io.Writer) error {
	return writeSolicitudConError(r, w)
}

func (r SolicitudConError) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"FechaHora\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"Cuit\",\"type\":\"string\"},{\"name\":\"NumeroComprobante\",\"type\":\"int\"},{\"name\":\"NombreArchivo\",\"type\":\"string\"},{\"name\":\"CodigoIntegridad\",\"type\":\"string\"},{\"name\":\"NumeroUnico\",\"type\":\"string\"},{\"name\":\"Procesado\",\"type\":\"string\"},{\"name\":\"Errores\",\"type\":{\"items\":{\"fields\":[{\"name\":\"TipoError\",\"type\":[\"null\",\"string\"]},{\"name\":\"Codigo\",\"type\":\"int\"},{\"name\":\"Descripcion\",\"type\":\"string\"}],\"name\":\"Error\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.GeneracionCOT.Events.Record.SolicitudConError\",\"type\":\"record\"}"
}

func (r SolicitudConError) SchemaName() string {
	return "Andreani.GeneracionCOT.Events.Record.SolicitudConError"
}

func (_ SolicitudConError) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ SolicitudConError) SetInt(v int32)       { panic("Unsupported operation") }
func (_ SolicitudConError) SetLong(v int64)      { panic("Unsupported operation") }
func (_ SolicitudConError) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ SolicitudConError) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ SolicitudConError) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ SolicitudConError) SetString(v string)   { panic("Unsupported operation") }
func (_ SolicitudConError) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *SolicitudConError) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Id}

		return w

	case 1:
		w := types.Long{Target: &r.FechaHora}

		return w

	case 2:
		w := types.String{Target: &r.Cuit}

		return w

	case 3:
		w := types.Int{Target: &r.NumeroComprobante}

		return w

	case 4:
		w := types.String{Target: &r.NombreArchivo}

		return w

	case 5:
		w := types.String{Target: &r.CodigoIntegridad}

		return w

	case 6:
		w := types.String{Target: &r.NumeroUnico}

		return w

	case 7:
		w := types.String{Target: &r.Procesado}

		return w

	case 8:
		r.Errores = make([]Error, 0)

		w := ArrayErrorWrapper{Target: &r.Errores}

		return w

	}
	panic("Unknown field index")
}

func (r *SolicitudConError) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *SolicitudConError) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ SolicitudConError) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ SolicitudConError) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ SolicitudConError) HintSize(int)                     { panic("Unsupported operation") }
func (_ SolicitudConError) Finalize()                        {}

func (_ SolicitudConError) AvroCRC64Fingerprint() []byte {
	return []byte(SolicitudConErrorAvroCRC64Fingerprint)
}

func (r SolicitudConError) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["FechaHora"], err = json.Marshal(r.FechaHora)
	if err != nil {
		return nil, err
	}
	output["Cuit"], err = json.Marshal(r.Cuit)
	if err != nil {
		return nil, err
	}
	output["NumeroComprobante"], err = json.Marshal(r.NumeroComprobante)
	if err != nil {
		return nil, err
	}
	output["NombreArchivo"], err = json.Marshal(r.NombreArchivo)
	if err != nil {
		return nil, err
	}
	output["CodigoIntegridad"], err = json.Marshal(r.CodigoIntegridad)
	if err != nil {
		return nil, err
	}
	output["NumeroUnico"], err = json.Marshal(r.NumeroUnico)
	if err != nil {
		return nil, err
	}
	output["Procesado"], err = json.Marshal(r.Procesado)
	if err != nil {
		return nil, err
	}
	output["Errores"], err = json.Marshal(r.Errores)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *SolicitudConError) UnmarshalJSON(data []byte) error {
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
		return fmt.Errorf("no value specified for Id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaHora"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaHora); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaHora")
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
		if v, ok := fields["NumeroComprobante"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroComprobante); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NumeroComprobante")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NombreArchivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NombreArchivo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NombreArchivo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoIntegridad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoIntegridad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodigoIntegridad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroUnico"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroUnico); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NumeroUnico")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Procesado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Procesado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Procesado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Errores"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Errores); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Errores")
	}
	return nil
}
