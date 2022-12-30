// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TransmitLogExt.avsc
 */
package SppeDocumentosEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type TransmitLogExt struct {
	Propietario string `json:"Propietario"`

	OrdenCliente string `json:"OrdenCliente"`

	Instancia string `json:"Instancia"`

	Almacen string `json:"Almacen"`

	Remito *UnionNullString `json:"Remito"`

	Contrato *UnionNullString `json:"Contrato"`
}

const TransmitLogExtAvroCRC64Fingerprint = "\x04#l]%\\N\xa4"

func NewTransmitLogExt() TransmitLogExt {
	r := TransmitLogExt{}
	r.Remito = nil
	r.Contrato = nil
	return r
}

func DeserializeTransmitLogExt(r io.Reader) (TransmitLogExt, error) {
	t := NewTransmitLogExt()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeTransmitLogExtFromSchema(r io.Reader, schema string) (TransmitLogExt, error) {
	t := NewTransmitLogExt()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeTransmitLogExt(r TransmitLogExt, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.OrdenCliente, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Instancia, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Almacen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Remito, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Contrato, w)
	if err != nil {
		return err
	}
	return err
}

func (r TransmitLogExt) Serialize(w io.Writer) error {
	return writeTransmitLogExt(r, w)
}

func (r TransmitLogExt) Schema() string {
	return "{\"fields\":[{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"OrdenCliente\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"default\":null,\"name\":\"Remito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Contrato\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.SppeDocumentos.Events.Record.TransmitLogExt\",\"type\":\"record\"}"
}

func (r TransmitLogExt) SchemaName() string {
	return "Andreani.SppeDocumentos.Events.Record.TransmitLogExt"
}

func (_ TransmitLogExt) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ TransmitLogExt) SetInt(v int32)       { panic("Unsupported operation") }
func (_ TransmitLogExt) SetLong(v int64)      { panic("Unsupported operation") }
func (_ TransmitLogExt) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ TransmitLogExt) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ TransmitLogExt) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ TransmitLogExt) SetString(v string)   { panic("Unsupported operation") }
func (_ TransmitLogExt) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *TransmitLogExt) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Propietario}

		return w

	case 1:
		w := types.String{Target: &r.OrdenCliente}

		return w

	case 2:
		w := types.String{Target: &r.Instancia}

		return w

	case 3:
		w := types.String{Target: &r.Almacen}

		return w

	case 4:
		r.Remito = NewUnionNullString()

		return r.Remito
	case 5:
		r.Contrato = NewUnionNullString()

		return r.Contrato
	}
	panic("Unknown field index")
}

func (r *TransmitLogExt) SetDefault(i int) {
	switch i {
	case 4:
		r.Remito = nil
		return
	case 5:
		r.Contrato = nil
		return
	}
	panic("Unknown field index")
}

func (r *TransmitLogExt) NullField(i int) {
	switch i {
	case 4:
		r.Remito = nil
		return
	case 5:
		r.Contrato = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ TransmitLogExt) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ TransmitLogExt) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ TransmitLogExt) HintSize(int)                     { panic("Unsupported operation") }
func (_ TransmitLogExt) Finalize()                        {}

func (_ TransmitLogExt) AvroCRC64Fingerprint() []byte {
	return []byte(TransmitLogExtAvroCRC64Fingerprint)
}

func (r TransmitLogExt) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["OrdenCliente"], err = json.Marshal(r.OrdenCliente)
	if err != nil {
		return nil, err
	}
	output["Instancia"], err = json.Marshal(r.Instancia)
	if err != nil {
		return nil, err
	}
	output["Almacen"], err = json.Marshal(r.Almacen)
	if err != nil {
		return nil, err
	}
	output["Remito"], err = json.Marshal(r.Remito)
	if err != nil {
		return nil, err
	}
	output["Contrato"], err = json.Marshal(r.Contrato)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *TransmitLogExt) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
		if v, ok := fields["Instancia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Instancia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Instancia")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Almacen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Almacen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Almacen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Remito"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Remito); err != nil {
			return err
		}
	} else {
		r.Remito = NewUnionNullString()

		r.Remito = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Contrato"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contrato); err != nil {
			return err
		}
	} else {
		r.Contrato = NewUnionNullString()

		r.Contrato = nil
	}
	return nil
}
