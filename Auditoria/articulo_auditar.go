// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ArticulosDeBulto.avsc
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

type ArticuloAuditar struct {
	Ean *UnionNullString `json:"Ean"`

	Sku string `json:"Sku"`

	Descripcion string `json:"Descripcion"`

	CantidadControlada int32 `json:"CantidadControlada"`

	CantidadPickeada int32 `json:"CantidadPickeada"`

	CantidadPedido int32 `json:"CantidadPedido"`

	NroLineaPedido string `json:"NroLineaPedido"`
}

const ArticuloAuditarAvroCRC64Fingerprint = "]v\xa4\x84\xa4\xc6@\xd0"

func NewArticuloAuditar() ArticuloAuditar {
	r := ArticuloAuditar{}
	r.Ean = nil
	return r
}

func DeserializeArticuloAuditar(r io.Reader) (ArticuloAuditar, error) {
	t := NewArticuloAuditar()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeArticuloAuditarFromSchema(r io.Reader, schema string) (ArticuloAuditar, error) {
	t := NewArticuloAuditar()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeArticuloAuditar(r ArticuloAuditar, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Ean, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Sku, w)
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
	err = vm.WriteInt(r.CantidadPickeada, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.CantidadPedido, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NroLineaPedido, w)
	if err != nil {
		return err
	}
	return err
}

func (r ArticuloAuditar) Serialize(w io.Writer) error {
	return writeArticuloAuditar(r, w)
}

func (r ArticuloAuditar) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Ean\",\"type\":[\"null\",\"string\"]},{\"name\":\"Sku\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"name\":\"CantidadControlada\",\"type\":\"int\"},{\"name\":\"CantidadPickeada\",\"type\":\"int\"},{\"name\":\"CantidadPedido\",\"type\":\"int\"},{\"name\":\"NroLineaPedido\",\"type\":\"string\"}],\"name\":\"Andreani.Auditoria.Events.Common.ArticuloAuditar\",\"type\":\"record\"}"
}

func (r ArticuloAuditar) SchemaName() string {
	return "Andreani.Auditoria.Events.Common.ArticuloAuditar"
}

func (_ ArticuloAuditar) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ArticuloAuditar) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ArticuloAuditar) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ArticuloAuditar) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ArticuloAuditar) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ArticuloAuditar) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ArticuloAuditar) SetString(v string)   { panic("Unsupported operation") }
func (_ ArticuloAuditar) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ArticuloAuditar) Get(i int) types.Field {
	switch i {
	case 0:
		r.Ean = NewUnionNullString()

		return r.Ean
	case 1:
		w := types.String{Target: &r.Sku}

		return w

	case 2:
		w := types.String{Target: &r.Descripcion}

		return w

	case 3:
		w := types.Int{Target: &r.CantidadControlada}

		return w

	case 4:
		w := types.Int{Target: &r.CantidadPickeada}

		return w

	case 5:
		w := types.Int{Target: &r.CantidadPedido}

		return w

	case 6:
		w := types.String{Target: &r.NroLineaPedido}

		return w

	}
	panic("Unknown field index")
}

func (r *ArticuloAuditar) SetDefault(i int) {
	switch i {
	case 0:
		r.Ean = nil
		return
	}
	panic("Unknown field index")
}

func (r *ArticuloAuditar) NullField(i int) {
	switch i {
	case 0:
		r.Ean = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ArticuloAuditar) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArticuloAuditar) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ArticuloAuditar) HintSize(int)                     { panic("Unsupported operation") }
func (_ ArticuloAuditar) Finalize()                        {}

func (_ ArticuloAuditar) AvroCRC64Fingerprint() []byte {
	return []byte(ArticuloAuditarAvroCRC64Fingerprint)
}

func (r ArticuloAuditar) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Ean"], err = json.Marshal(r.Ean)
	if err != nil {
		return nil, err
	}
	output["Sku"], err = json.Marshal(r.Sku)
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
	output["CantidadPickeada"], err = json.Marshal(r.CantidadPickeada)
	if err != nil {
		return nil, err
	}
	output["CantidadPedido"], err = json.Marshal(r.CantidadPedido)
	if err != nil {
		return nil, err
	}
	output["NroLineaPedido"], err = json.Marshal(r.NroLineaPedido)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ArticuloAuditar) UnmarshalJSON(data []byte) error {
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
		r.Ean = NewUnionNullString()

		r.Ean = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Sku"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sku); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Sku")
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
		if v, ok := fields["CantidadPickeada"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadPickeada); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CantidadPickeada")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CantidadPedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadPedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CantidadPedido")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NroLineaPedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NroLineaPedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NroLineaPedido")
	}
	return nil
}
