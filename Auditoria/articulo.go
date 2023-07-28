// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     BultoValidar.avsc
 */
package AuditoriaEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Articulo struct {
	Ean string `json:"Ean"`

	SKU string `json:"SKU"`

	Descripcion string `json:"Descripcion"`

	CantidadControlada int32 `json:"CantidadControlada"`

	ValorUnitario *UnionNullString `json:"ValorUnitario"`

	Picker *UnionNullString `json:"Picker"`

	Packer *UnionNullString `json:"Packer"`
}

const ArticuloAvroCRC64Fingerprint = "\"ϯ(\xf0\xaa$C"

func NewArticulo() Articulo {
	r := Articulo{}
	r.ValorUnitario = nil
	r.Picker = nil
	r.Packer = nil
	return r
}

func DeserializeArticulo(r io.Reader) (Articulo, error) {
	t := NewArticulo()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeArticuloFromSchema(r io.Reader, schema string) (Articulo, error) {
	t := NewArticulo()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeArticulo(r Articulo, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Ean, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.SKU, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Descripcion, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.CantidadControlada, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ValorUnitario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Picker, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Packer, w)
	if err != nil {
		return err
	}
	return err
}

func (r Articulo) Serialize(w io.Writer) error {
	return writeArticulo(r, w)
}

func (r Articulo) Schema() string {
	return "{\"fields\":[{\"name\":\"Ean\",\"type\":\"string\"},{\"name\":\"SKU\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"name\":\"CantidadControlada\",\"type\":\"int\"},{\"default\":null,\"name\":\"ValorUnitario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Picker\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Packer\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.Auditoria.Events.Common.Articulo\",\"type\":\"record\"}"
}

func (r Articulo) SchemaName() string {
	return "Andreani.Auditoria.Events.Common.Articulo"
}

func (_ Articulo) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Articulo) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Articulo) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Articulo) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Articulo) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Articulo) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Articulo) SetString(v string)   { panic("Unsupported operation") }
func (_ Articulo) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Articulo) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Ean}

		return w

	case 1:
		w := types.String{Target: &r.SKU}

		return w

	case 2:
		w := types.String{Target: &r.Descripcion}

		return w

	case 3:
		w := types.Int{Target: &r.CantidadControlada}

		return w

	case 4:
		r.ValorUnitario = NewUnionNullString()

		return r.ValorUnitario
	case 5:
		r.Picker = NewUnionNullString()

		return r.Picker
	case 6:
		r.Packer = NewUnionNullString()

		return r.Packer
	}
	panic("Unknown field index")
}

func (r *Articulo) SetDefault(i int) {
	switch i {
	case 4:
		r.ValorUnitario = nil
		return
	case 5:
		r.Picker = nil
		return
	case 6:
		r.Packer = nil
		return
	}
	panic("Unknown field index")
}

func (r *Articulo) NullField(i int) {
	switch i {
	case 4:
		r.ValorUnitario = nil
		return
	case 5:
		r.Picker = nil
		return
	case 6:
		r.Packer = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Articulo) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Articulo) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Articulo) HintSize(int)                     { panic("Unsupported operation") }
func (_ Articulo) Finalize()                        {}

func (_ Articulo) AvroCRC64Fingerprint() []byte {
	return []byte(ArticuloAvroCRC64Fingerprint)
}

func (r Articulo) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Ean"], err = json.Marshal(r.Ean)
	if err != nil {
		return nil, err
	}
	output["SKU"], err = json.Marshal(r.SKU)
	if err != nil {
		return nil, err
	}
	output["Descripcion"], err = json.Marshal(r.Descripcion)
	if err != nil {
		return nil, err
	}
	output["CantidadControlada"], err = json.Marshal(r.CantidadControlada)
	if err != nil {
		return nil, err
	}
	output["ValorUnitario"], err = json.Marshal(r.ValorUnitario)
	if err != nil {
		return nil, err
	}
	output["Picker"], err = json.Marshal(r.Picker)
	if err != nil {
		return nil, err
	}
	output["Packer"], err = json.Marshal(r.Packer)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Articulo) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Ean"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ean); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Ean")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SKU"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SKU); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SKU")
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
		if v, ok := fields["CantidadControlada"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadControlada); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CantidadControlada")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ValorUnitario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ValorUnitario); err != nil {
			return err
		}
	} else {
		r.ValorUnitario = NewUnionNullString()

		r.ValorUnitario = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Picker"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Picker); err != nil {
			return err
		}
	} else {
		r.Picker = NewUnionNullString()

		r.Picker = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Packer"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Packer); err != nil {
			return err
		}
	} else {
		r.Packer = NewUnionNullString()

		r.Packer = nil
	}
	return nil
}
