// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnvioEtiqueta.avsc
 */
package CorporativoEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type DestinatarioEtiqueta struct {
	Nombre string `json:"Nombre"`

	Apellido string `json:"Apellido"`

	Telefono string `json:"Telefono"`

	DNI string `json:"DNI"`

	Email string `json:"Email"`

	DestinatarioEmail string `json:"DestinatarioEmail"`
}

const DestinatarioEtiquetaAvroCRC64Fingerprint = "%Q\x8f\xeby\xbc\x87\x94"

func NewDestinatarioEtiqueta() DestinatarioEtiqueta {
	r := DestinatarioEtiqueta{}
	return r
}

func DeserializeDestinatarioEtiqueta(r io.Reader) (DestinatarioEtiqueta, error) {
	t := NewDestinatarioEtiqueta()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDestinatarioEtiquetaFromSchema(r io.Reader, schema string) (DestinatarioEtiqueta, error) {
	t := NewDestinatarioEtiqueta()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDestinatarioEtiqueta(r DestinatarioEtiqueta, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Nombre, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Apellido, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Telefono, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.DNI, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Email, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.DestinatarioEmail, w)
	if err != nil {
		return err
	}
	return err
}

func (r DestinatarioEtiqueta) Serialize(w io.Writer) error {
	return writeDestinatarioEtiqueta(r, w)
}

func (r DestinatarioEtiqueta) Schema() string {
	return "{\"fields\":[{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Apellido\",\"type\":\"string\"},{\"name\":\"Telefono\",\"type\":\"string\"},{\"name\":\"DNI\",\"type\":\"string\"},{\"name\":\"Email\",\"type\":\"string\"},{\"name\":\"DestinatarioEmail\",\"type\":\"string\"}],\"name\":\"Andreani.Corporativo.Events.Record.DestinatarioEtiqueta\",\"type\":\"record\"}"
}

func (r DestinatarioEtiqueta) SchemaName() string {
	return "Andreani.Corporativo.Events.Record.DestinatarioEtiqueta"
}

func (_ DestinatarioEtiqueta) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DestinatarioEtiqueta) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DestinatarioEtiqueta) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DestinatarioEtiqueta) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DestinatarioEtiqueta) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DestinatarioEtiqueta) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DestinatarioEtiqueta) SetString(v string)   { panic("Unsupported operation") }
func (_ DestinatarioEtiqueta) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DestinatarioEtiqueta) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Nombre}

		return w

	case 1:
		w := types.String{Target: &r.Apellido}

		return w

	case 2:
		w := types.String{Target: &r.Telefono}

		return w

	case 3:
		w := types.String{Target: &r.DNI}

		return w

	case 4:
		w := types.String{Target: &r.Email}

		return w

	case 5:
		w := types.String{Target: &r.DestinatarioEmail}

		return w

	}
	panic("Unknown field index")
}

func (r *DestinatarioEtiqueta) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *DestinatarioEtiqueta) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ DestinatarioEtiqueta) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DestinatarioEtiqueta) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DestinatarioEtiqueta) HintSize(int)                     { panic("Unsupported operation") }
func (_ DestinatarioEtiqueta) Finalize()                        {}

func (_ DestinatarioEtiqueta) AvroCRC64Fingerprint() []byte {
	return []byte(DestinatarioEtiquetaAvroCRC64Fingerprint)
}

func (r DestinatarioEtiqueta) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Nombre"], err = json.Marshal(r.Nombre)
	if err != nil {
		return nil, err
	}
	output["Apellido"], err = json.Marshal(r.Apellido)
	if err != nil {
		return nil, err
	}
	output["Telefono"], err = json.Marshal(r.Telefono)
	if err != nil {
		return nil, err
	}
	output["DNI"], err = json.Marshal(r.DNI)
	if err != nil {
		return nil, err
	}
	output["Email"], err = json.Marshal(r.Email)
	if err != nil {
		return nil, err
	}
	output["DestinatarioEmail"], err = json.Marshal(r.DestinatarioEmail)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DestinatarioEtiqueta) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Nombre"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Nombre); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Nombre")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Apellido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Apellido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Apellido")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Telefono"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Telefono); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Telefono")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DNI"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DNI); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DNI")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Email"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Email); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Email")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioEmail"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioEmail); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioEmail")
	}
	return nil
}
