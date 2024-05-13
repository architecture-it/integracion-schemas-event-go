// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhGLNCambio.avsc
 */
package EventosWhGLNEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Cabecera struct {
	OrdenWH string `json:"OrdenWH"`

	OrdenCliente string `json:"OrdenCliente"`

	Remito string `json:"Remito"`

	Propietario string `json:"Propietario"`

	GLNOriginal string `json:"GLNOriginal"`

	GLNNuevo string `json:"GLNNuevo"`
}

const CabeceraAvroCRC64Fingerprint = "\a\x92ގ5\x16\x85\b"

func NewCabecera() Cabecera {
	r := Cabecera{}
	return r
}

func DeserializeCabecera(r io.Reader) (Cabecera, error) {
	t := NewCabecera()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCabeceraFromSchema(r io.Reader, schema string) (Cabecera, error) {
	t := NewCabecera()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCabecera(r Cabecera, w io.Writer) error {
	var err error
	err = vm.WriteString(r.OrdenWH, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.OrdenCliente, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Remito, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.GLNOriginal, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.GLNNuevo, w)
	if err != nil {
		return err
	}
	return err
}

func (r Cabecera) Serialize(w io.Writer) error {
	return writeCabecera(r, w)
}

func (r Cabecera) Schema() string {
	return "{\"fields\":[{\"name\":\"OrdenWH\",\"type\":\"string\"},{\"name\":\"OrdenCliente\",\"type\":\"string\"},{\"name\":\"Remito\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"GLNOriginal\",\"type\":\"string\"},{\"name\":\"GLNNuevo\",\"type\":\"string\"}],\"name\":\"Andreani.EventosWhGLN.Events.CambioCommon.Cabecera\",\"type\":\"record\"}"
}

func (r Cabecera) SchemaName() string {
	return "Andreani.EventosWhGLN.Events.CambioCommon.Cabecera"
}

func (_ Cabecera) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Cabecera) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Cabecera) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Cabecera) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Cabecera) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Cabecera) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Cabecera) SetString(v string)   { panic("Unsupported operation") }
func (_ Cabecera) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Cabecera) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.OrdenWH}

		return w

	case 1:
		w := types.String{Target: &r.OrdenCliente}

		return w

	case 2:
		w := types.String{Target: &r.Remito}

		return w

	case 3:
		w := types.String{Target: &r.Propietario}

		return w

	case 4:
		w := types.String{Target: &r.GLNOriginal}

		return w

	case 5:
		w := types.String{Target: &r.GLNNuevo}

		return w

	}
	panic("Unknown field index")
}

func (r *Cabecera) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Cabecera) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Cabecera) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Cabecera) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Cabecera) HintSize(int)                     { panic("Unsupported operation") }
func (_ Cabecera) Finalize()                        {}

func (_ Cabecera) AvroCRC64Fingerprint() []byte {
	return []byte(CabeceraAvroCRC64Fingerprint)
}

func (r Cabecera) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["OrdenWH"], err = json.Marshal(r.OrdenWH)
	if err != nil {
		return nil, err
	}
	output["OrdenCliente"], err = json.Marshal(r.OrdenCliente)
	if err != nil {
		return nil, err
	}
	output["Remito"], err = json.Marshal(r.Remito)
	if err != nil {
		return nil, err
	}
	output["Propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["GLNOriginal"], err = json.Marshal(r.GLNOriginal)
	if err != nil {
		return nil, err
	}
	output["GLNNuevo"], err = json.Marshal(r.GLNNuevo)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Cabecera) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["OrdenWH"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenWH); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrdenWH")
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
		return fmt.Errorf("no value specified for Remito")
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
		if v, ok := fields["GLNOriginal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.GLNOriginal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for GLNOriginal")
	}
	val = func() json.RawMessage {
		if v, ok := fields["GLNNuevo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.GLNNuevo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for GLNNuevo")
	}
	return nil
}