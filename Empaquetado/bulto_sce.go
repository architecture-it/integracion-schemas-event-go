// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoSce.avsc
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

type BultoSce struct {
	ContenedorPreparacion string `json:"ContenedorPreparacion"`

	CodigoEmbalaje string `json:"CodigoEmbalaje"`

	TipoContenedorEmbalaje string `json:"TipoContenedorEmbalaje"`

	Longitud *UnionNullFloat `json:"Longitud"`

	Altura *UnionNullFloat `json:"Altura"`

	Ancho *UnionNullFloat `json:"Ancho"`

	Peso *UnionNullFloat `json:"Peso"`

	ArticulosSce []ArticuloSce `json:"ArticulosSce"`
}

const BultoSceAvroCRC64Fingerprint = "\u009e\x06\x9bq\xf7\xdc\xdd"

func NewBultoSce() BultoSce {
	r := BultoSce{}
	r.Longitud = nil
	r.Altura = nil
	r.Ancho = nil
	r.Peso = nil
	r.ArticulosSce = make([]ArticuloSce, 0)

	return r
}

func DeserializeBultoSce(r io.Reader) (BultoSce, error) {
	t := NewBultoSce()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBultoSceFromSchema(r io.Reader, schema string) (BultoSce, error) {
	t := NewBultoSce()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBultoSce(r BultoSce, w io.Writer) error {
	var err error
	err = vm.WriteString(r.ContenedorPreparacion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoEmbalaje, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TipoContenedorEmbalaje, w)
	if err != nil {
		return err
	}
	err = writeUnionNullFloat(r.Longitud, w)
	if err != nil {
		return err
	}
	err = writeUnionNullFloat(r.Altura, w)
	if err != nil {
		return err
	}
	err = writeUnionNullFloat(r.Ancho, w)
	if err != nil {
		return err
	}
	err = writeUnionNullFloat(r.Peso, w)
	if err != nil {
		return err
	}
	err = writeArrayArticuloSce(r.ArticulosSce, w)
	if err != nil {
		return err
	}
	return err
}

func (r BultoSce) Serialize(w io.Writer) error {
	return writeBultoSce(r, w)
}

func (r BultoSce) Schema() string {
	return "{\"fields\":[{\"name\":\"ContenedorPreparacion\",\"type\":\"string\"},{\"name\":\"CodigoEmbalaje\",\"type\":\"string\"},{\"name\":\"TipoContenedorEmbalaje\",\"type\":\"string\"},{\"default\":null,\"name\":\"Longitud\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Altura\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Ancho\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Peso\",\"type\":[\"null\",\"float\"]},{\"name\":\"ArticulosSce\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Sku\",\"type\":\"string\"},{\"name\":\"Cantidad\",\"type\":\"int\"}],\"name\":\"ArticuloSce\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.Empaquetado.Events.Common.BultoSce\",\"type\":\"record\"}"
}

func (r BultoSce) SchemaName() string {
	return "Andreani.Empaquetado.Events.Common.BultoSce"
}

func (_ BultoSce) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BultoSce) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BultoSce) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BultoSce) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BultoSce) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BultoSce) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BultoSce) SetString(v string)   { panic("Unsupported operation") }
func (_ BultoSce) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BultoSce) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.ContenedorPreparacion}

		return w

	case 1:
		w := types.String{Target: &r.CodigoEmbalaje}

		return w

	case 2:
		w := types.String{Target: &r.TipoContenedorEmbalaje}

		return w

	case 3:
		r.Longitud = NewUnionNullFloat()

		return r.Longitud
	case 4:
		r.Altura = NewUnionNullFloat()

		return r.Altura
	case 5:
		r.Ancho = NewUnionNullFloat()

		return r.Ancho
	case 6:
		r.Peso = NewUnionNullFloat()

		return r.Peso
	case 7:
		r.ArticulosSce = make([]ArticuloSce, 0)

		w := ArrayArticuloSceWrapper{Target: &r.ArticulosSce}

		return w

	}
	panic("Unknown field index")
}

func (r *BultoSce) SetDefault(i int) {
	switch i {
	case 3:
		r.Longitud = nil
		return
	case 4:
		r.Altura = nil
		return
	case 5:
		r.Ancho = nil
		return
	case 6:
		r.Peso = nil
		return
	}
	panic("Unknown field index")
}

func (r *BultoSce) NullField(i int) {
	switch i {
	case 3:
		r.Longitud = nil
		return
	case 4:
		r.Altura = nil
		return
	case 5:
		r.Ancho = nil
		return
	case 6:
		r.Peso = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ BultoSce) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ BultoSce) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ BultoSce) HintSize(int)                     { panic("Unsupported operation") }
func (_ BultoSce) Finalize()                        {}

func (_ BultoSce) AvroCRC64Fingerprint() []byte {
	return []byte(BultoSceAvroCRC64Fingerprint)
}

func (r BultoSce) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["ContenedorPreparacion"], err = json.Marshal(r.ContenedorPreparacion)
	if err != nil {
		return nil, err
	}
	output["CodigoEmbalaje"], err = json.Marshal(r.CodigoEmbalaje)
	if err != nil {
		return nil, err
	}
	output["TipoContenedorEmbalaje"], err = json.Marshal(r.TipoContenedorEmbalaje)
	if err != nil {
		return nil, err
	}
	output["Longitud"], err = json.Marshal(r.Longitud)
	if err != nil {
		return nil, err
	}
	output["Altura"], err = json.Marshal(r.Altura)
	if err != nil {
		return nil, err
	}
	output["Ancho"], err = json.Marshal(r.Ancho)
	if err != nil {
		return nil, err
	}
	output["Peso"], err = json.Marshal(r.Peso)
	if err != nil {
		return nil, err
	}
	output["ArticulosSce"], err = json.Marshal(r.ArticulosSce)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *BultoSce) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["ContenedorPreparacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ContenedorPreparacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ContenedorPreparacion")
	}
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
		if v, ok := fields["TipoContenedorEmbalaje"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoContenedorEmbalaje); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoContenedorEmbalaje")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Longitud"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Longitud); err != nil {
			return err
		}
	} else {
		r.Longitud = NewUnionNullFloat()

		r.Longitud = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Altura"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Altura); err != nil {
			return err
		}
	} else {
		r.Altura = NewUnionNullFloat()

		r.Altura = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Ancho"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ancho); err != nil {
			return err
		}
	} else {
		r.Ancho = NewUnionNullFloat()

		r.Ancho = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Peso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Peso); err != nil {
			return err
		}
	} else {
		r.Peso = NewUnionNullFloat()

		r.Peso = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ArticulosSce"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ArticulosSce); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ArticulosSce")
	}
	return nil
}