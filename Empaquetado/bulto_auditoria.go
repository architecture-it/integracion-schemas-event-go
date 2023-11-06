// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoAuditar.avsc
 */
package EmpaquetadoEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type BultoAuditoria struct {
	CodigoEmbalaje string `json:"CodigoEmbalaje"`

	Articulos []ArticuloAuditoria `json:"Articulos"`
}

const BultoAuditoriaAvroCRC64Fingerprint = "[%\x17n\x06jƘ"

func NewBultoAuditoria() BultoAuditoria {
	r := BultoAuditoria{}
	r.Articulos = make([]ArticuloAuditoria, 0)

	return r
}

func DeserializeBultoAuditoria(r io.Reader) (BultoAuditoria, error) {
	t := NewBultoAuditoria()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBultoAuditoriaFromSchema(r io.Reader, schema string) (BultoAuditoria, error) {
	t := NewBultoAuditoria()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBultoAuditoria(r BultoAuditoria, w io.Writer) error {
	var err error
	err = vm.WriteString(r.CodigoEmbalaje, w)
	if err != nil {
		return err
	}
	err = writeArrayArticuloAuditoria(r.Articulos, w)
	if err != nil {
		return err
	}
	return err
}

func (r BultoAuditoria) Serialize(w io.Writer) error {
	return writeBultoAuditoria(r, w)
}

func (r BultoAuditoria) Schema() string {
	return "{\"fields\":[{\"name\":\"CodigoEmbalaje\",\"type\":\"string\"},{\"name\":\"Articulos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Sku\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"default\":null,\"name\":\"Ean\",\"type\":[\"null\",\"string\"]},{\"name\":\"NroLineaPedido\",\"type\":\"string\"},{\"name\":\"CantidadPedido\",\"type\":\"int\"},{\"name\":\"CantidadPickeada\",\"type\":\"int\"},{\"name\":\"CantidadEmpacada\",\"type\":\"int\"},{\"default\":null,\"name\":\"Diferencia\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"Error\",\"type\":[\"null\",\"string\"]}],\"name\":\"ArticuloAuditoria\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.Empaquetado.Events.Common.BultoAuditoria\",\"type\":\"record\"}"
}

func (r BultoAuditoria) SchemaName() string {
	return "Andreani.Empaquetado.Events.Common.BultoAuditoria"
}

func (_ BultoAuditoria) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BultoAuditoria) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BultoAuditoria) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BultoAuditoria) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BultoAuditoria) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BultoAuditoria) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BultoAuditoria) SetString(v string)   { panic("Unsupported operation") }
func (_ BultoAuditoria) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BultoAuditoria) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.CodigoEmbalaje}

		return w

	case 1:
		r.Articulos = make([]ArticuloAuditoria, 0)

		w := ArrayArticuloAuditoriaWrapper{Target: &r.Articulos}

		return w

	}
	panic("Unknown field index")
}

func (r *BultoAuditoria) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *BultoAuditoria) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ BultoAuditoria) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ BultoAuditoria) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ BultoAuditoria) HintSize(int)                     { panic("Unsupported operation") }
func (_ BultoAuditoria) Finalize()                        {}

func (_ BultoAuditoria) AvroCRC64Fingerprint() []byte {
	return []byte(BultoAuditoriaAvroCRC64Fingerprint)
}

func (r BultoAuditoria) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["CodigoEmbalaje"], err = json.Marshal(r.CodigoEmbalaje)
	if err != nil {
		return nil, err
	}
	output["Articulos"], err = json.Marshal(r.Articulos)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *BultoAuditoria) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["CodigoEmbalaje"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoEmbalaje); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodigoEmbalaje")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Articulos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Articulos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Articulos")
	}
	return nil
}
