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
	SKU *UnionNullString `json:"SKU"`

	CodigoCliente string `json:"CodigoCliente"`

	Descripcion string `json:"Descripcion"`

	CantidadControlada int32 `json:"CantidadControlada"`

	CantidadPickeada int32 `json:"CantidadPickeada"`

	NroLineaPedido string `json:"NroLineaPedido"`
}

const ArticuloAuditarAvroCRC64Fingerprint = "\xaf\x9e,O\x99\f\t\x7f"

func NewArticuloAuditar() ArticuloAuditar {
	r := ArticuloAuditar{}
	r.SKU = nil
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
	err = writeUnionNullString(r.SKU, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoCliente, w)
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
	return "{\"fields\":[{\"default\":null,\"name\":\"SKU\",\"type\":[\"null\",\"string\"]},{\"name\":\"CodigoCliente\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"name\":\"CantidadControlada\",\"type\":\"int\"},{\"name\":\"CantidadPickeada\",\"type\":\"int\"},{\"name\":\"NroLineaPedido\",\"type\":\"string\"}],\"name\":\"Andreani.Auditoria.Events.Common.ArticuloAuditar\",\"type\":\"record\"}"
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
		r.SKU = NewUnionNullString()

		return r.SKU
	case 1:
		w := types.String{Target: &r.CodigoCliente}

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
		w := types.String{Target: &r.NroLineaPedido}

		return w

	}
	panic("Unknown field index")
}

func (r *ArticuloAuditar) SetDefault(i int) {
	switch i {
	case 0:
		r.SKU = nil
		return
	}
	panic("Unknown field index")
}

func (r *ArticuloAuditar) NullField(i int) {
	switch i {
	case 0:
		r.SKU = nil
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
	output["SKU"], err = json.Marshal(r.SKU)
	if err != nil {
		return nil, err
	}
	output["CodigoCliente"], err = json.Marshal(r.CodigoCliente)
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
		r.SKU = NewUnionNullString()

		r.SKU = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoCliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodigoCliente")
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
