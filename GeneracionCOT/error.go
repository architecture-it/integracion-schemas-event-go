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

type Error struct {
	TipoError *UnionNullString `json:"TipoError"`

	Codigo int32 `json:"Codigo"`

	Descripcion string `json:"Descripcion"`
}

const ErrorAvroCRC64Fingerprint = "\xea\x81\xfd\x96\xb7U:\xde"

func NewError() Error {
	r := Error{}
	return r
}

func DeserializeError(r io.Reader) (Error, error) {
	t := NewError()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeErrorFromSchema(r io.Reader, schema string) (Error, error) {
	t := NewError()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeError(r Error, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.TipoError, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Codigo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Descripcion, w)
	if err != nil {
		return err
	}
	return err
}

func (r Error) Serialize(w io.Writer) error {
	return writeError(r, w)
}

func (r Error) Schema() string {
	return "{\"fields\":[{\"name\":\"TipoError\",\"type\":[\"null\",\"string\"]},{\"name\":\"Codigo\",\"type\":\"int\"},{\"name\":\"Descripcion\",\"type\":\"string\"}],\"name\":\"Andreani.GeneracionCOT.Events.Record.Error\",\"type\":\"record\"}"
}

func (r Error) SchemaName() string {
	return "Andreani.GeneracionCOT.Events.Record.Error"
}

func (_ Error) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Error) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Error) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Error) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Error) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Error) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Error) SetString(v string)   { panic("Unsupported operation") }
func (_ Error) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Error) Get(i int) types.Field {
	switch i {
	case 0:
		r.TipoError = NewUnionNullString()

		return r.TipoError
	case 1:
		w := types.Int{Target: &r.Codigo}

		return w

	case 2:
		w := types.String{Target: &r.Descripcion}

		return w

	}
	panic("Unknown field index")
}

func (r *Error) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Error) NullField(i int) {
	switch i {
	case 0:
		r.TipoError = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Error) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Error) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Error) HintSize(int)                     { panic("Unsupported operation") }
func (_ Error) Finalize()                        {}

func (_ Error) AvroCRC64Fingerprint() []byte {
	return []byte(ErrorAvroCRC64Fingerprint)
}

func (r Error) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["TipoError"], err = json.Marshal(r.TipoError)
	if err != nil {
		return nil, err
	}
	output["Codigo"], err = json.Marshal(r.Codigo)
	if err != nil {
		return nil, err
	}
	output["Descripcion"], err = json.Marshal(r.Descripcion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Error) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["TipoError"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoError); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoError")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Codigo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Codigo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Codigo")
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
	return nil
}