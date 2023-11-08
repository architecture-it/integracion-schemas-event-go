// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     SugerenciaContenedores.avsc
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

type ArticuloContenedorPreparacion struct {
	Sku string `json:"Sku"`

	Descripcion string `json:"Descripcion"`

	Ean *UnionNullString `json:"Ean"`

	NroLineaPedido string `json:"NroLineaPedido"`

	CantidadPedido int32 `json:"CantidadPedido"`

	CantidadPickeada int32 `json:"CantidadPickeada"`

	Lote *UnionNullString `json:"Lote"`

	Serie *UnionNullString `json:"Serie"`

	Zona *UnionNullString `json:"Zona"`

	CodigoZona *UnionNullString `json:"CodigoZona"`

	DescripcionZona *UnionNullString `json:"DescripcionZona"`

	Longitud *UnionNullFloat `json:"Longitud"`

	Altura *UnionNullFloat `json:"Altura"`

	Ancho *UnionNullFloat `json:"Ancho"`

	Peso *UnionNullFloat `json:"Peso"`

	InstruccionesEmbalaje *UnionNullFloat `json:"InstruccionesEmbalaje"`
}

const ArticuloContenedorPreparacionAvroCRC64Fingerprint = "\xa0\xcf\x0e\xd6m\x18\x85\xa6"

func NewArticuloContenedorPreparacion() ArticuloContenedorPreparacion {
	r := ArticuloContenedorPreparacion{}
	r.Ean = nil
	r.Lote = nil
	r.Serie = nil
	r.Zona = nil
	r.CodigoZona = nil
	r.DescripcionZona = nil
	r.Longitud = nil
	r.Altura = nil
	r.Ancho = nil
	r.Peso = nil
	r.InstruccionesEmbalaje = nil
	return r
}

func DeserializeArticuloContenedorPreparacion(r io.Reader) (ArticuloContenedorPreparacion, error) {
	t := NewArticuloContenedorPreparacion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeArticuloContenedorPreparacionFromSchema(r io.Reader, schema string) (ArticuloContenedorPreparacion, error) {
	t := NewArticuloContenedorPreparacion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeArticuloContenedorPreparacion(r ArticuloContenedorPreparacion, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Sku, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Descripcion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Ean, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NroLineaPedido, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.CantidadPedido, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.CantidadPickeada, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Lote, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Serie, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Zona, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoZona, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DescripcionZona, w)
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
	err = writeUnionNullFloat(r.InstruccionesEmbalaje, w)
	if err != nil {
		return err
	}
	return err
}

func (r ArticuloContenedorPreparacion) Serialize(w io.Writer) error {
	return writeArticuloContenedorPreparacion(r, w)
}

func (r ArticuloContenedorPreparacion) Schema() string {
	return "{\"fields\":[{\"name\":\"Sku\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"default\":null,\"name\":\"Ean\",\"type\":[\"null\",\"string\"]},{\"name\":\"NroLineaPedido\",\"type\":\"string\"},{\"name\":\"CantidadPedido\",\"type\":\"int\"},{\"name\":\"CantidadPickeada\",\"type\":\"int\"},{\"default\":null,\"name\":\"Lote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Serie\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Zona\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoZona\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DescripcionZona\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Longitud\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Altura\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Ancho\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Peso\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"InstruccionesEmbalaje\",\"type\":[\"null\",\"float\"]}],\"name\":\"Andreani.Empaquetado.Events.Common.ArticuloContenedorPreparacion\",\"type\":\"record\"}"
}

func (r ArticuloContenedorPreparacion) SchemaName() string {
	return "Andreani.Empaquetado.Events.Common.ArticuloContenedorPreparacion"
}

func (_ ArticuloContenedorPreparacion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ArticuloContenedorPreparacion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ArticuloContenedorPreparacion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ArticuloContenedorPreparacion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ArticuloContenedorPreparacion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ArticuloContenedorPreparacion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ArticuloContenedorPreparacion) SetString(v string)   { panic("Unsupported operation") }
func (_ ArticuloContenedorPreparacion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ArticuloContenedorPreparacion) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Sku}

		return w

	case 1:
		w := types.String{Target: &r.Descripcion}

		return w

	case 2:
		r.Ean = NewUnionNullString()

		return r.Ean
	case 3:
		w := types.String{Target: &r.NroLineaPedido}

		return w

	case 4:
		w := types.Int{Target: &r.CantidadPedido}

		return w

	case 5:
		w := types.Int{Target: &r.CantidadPickeada}

		return w

	case 6:
		r.Lote = NewUnionNullString()

		return r.Lote
	case 7:
		r.Serie = NewUnionNullString()

		return r.Serie
	case 8:
		r.Zona = NewUnionNullString()

		return r.Zona
	case 9:
		r.CodigoZona = NewUnionNullString()

		return r.CodigoZona
	case 10:
		r.DescripcionZona = NewUnionNullString()

		return r.DescripcionZona
	case 11:
		r.Longitud = NewUnionNullFloat()

		return r.Longitud
	case 12:
		r.Altura = NewUnionNullFloat()

		return r.Altura
	case 13:
		r.Ancho = NewUnionNullFloat()

		return r.Ancho
	case 14:
		r.Peso = NewUnionNullFloat()

		return r.Peso
	case 15:
		r.InstruccionesEmbalaje = NewUnionNullFloat()

		return r.InstruccionesEmbalaje
	}
	panic("Unknown field index")
}

func (r *ArticuloContenedorPreparacion) SetDefault(i int) {
	switch i {
	case 2:
		r.Ean = nil
		return
	case 6:
		r.Lote = nil
		return
	case 7:
		r.Serie = nil
		return
	case 8:
		r.Zona = nil
		return
	case 9:
		r.CodigoZona = nil
		return
	case 10:
		r.DescripcionZona = nil
		return
	case 11:
		r.Longitud = nil
		return
	case 12:
		r.Altura = nil
		return
	case 13:
		r.Ancho = nil
		return
	case 14:
		r.Peso = nil
		return
	case 15:
		r.InstruccionesEmbalaje = nil
		return
	}
	panic("Unknown field index")
}

func (r *ArticuloContenedorPreparacion) NullField(i int) {
	switch i {
	case 2:
		r.Ean = nil
		return
	case 6:
		r.Lote = nil
		return
	case 7:
		r.Serie = nil
		return
	case 8:
		r.Zona = nil
		return
	case 9:
		r.CodigoZona = nil
		return
	case 10:
		r.DescripcionZona = nil
		return
	case 11:
		r.Longitud = nil
		return
	case 12:
		r.Altura = nil
		return
	case 13:
		r.Ancho = nil
		return
	case 14:
		r.Peso = nil
		return
	case 15:
		r.InstruccionesEmbalaje = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ArticuloContenedorPreparacion) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArticuloContenedorPreparacion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ ArticuloContenedorPreparacion) HintSize(int)             { panic("Unsupported operation") }
func (_ ArticuloContenedorPreparacion) Finalize()                {}

func (_ ArticuloContenedorPreparacion) AvroCRC64Fingerprint() []byte {
	return []byte(ArticuloContenedorPreparacionAvroCRC64Fingerprint)
}

func (r ArticuloContenedorPreparacion) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Sku"], err = json.Marshal(r.Sku)
	if err != nil {
		return nil, err
	}
	output["Descripcion"], err = json.Marshal(r.Descripcion)
	if err != nil {
		return nil, err
	}
	output["Ean"], err = json.Marshal(r.Ean)
	if err != nil {
		return nil, err
	}
	output["NroLineaPedido"], err = json.Marshal(r.NroLineaPedido)
	if err != nil {
		return nil, err
	}
	output["CantidadPedido"], err = json.Marshal(r.CantidadPedido)
	if err != nil {
		return nil, err
	}
	output["CantidadPickeada"], err = json.Marshal(r.CantidadPickeada)
	if err != nil {
		return nil, err
	}
	output["Lote"], err = json.Marshal(r.Lote)
	if err != nil {
		return nil, err
	}
	output["Serie"], err = json.Marshal(r.Serie)
	if err != nil {
		return nil, err
	}
	output["Zona"], err = json.Marshal(r.Zona)
	if err != nil {
		return nil, err
	}
	output["CodigoZona"], err = json.Marshal(r.CodigoZona)
	if err != nil {
		return nil, err
	}
	output["DescripcionZona"], err = json.Marshal(r.DescripcionZona)
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
	output["InstruccionesEmbalaje"], err = json.Marshal(r.InstruccionesEmbalaje)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ArticuloContenedorPreparacion) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
		if v, ok := fields["Lote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Lote); err != nil {
			return err
		}
	} else {
		r.Lote = NewUnionNullString()

		r.Lote = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Serie"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Serie); err != nil {
			return err
		}
	} else {
		r.Serie = NewUnionNullString()

		r.Serie = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Zona"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Zona); err != nil {
			return err
		}
	} else {
		r.Zona = NewUnionNullString()

		r.Zona = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoZona"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoZona); err != nil {
			return err
		}
	} else {
		r.CodigoZona = NewUnionNullString()

		r.CodigoZona = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DescripcionZona"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DescripcionZona); err != nil {
			return err
		}
	} else {
		r.DescripcionZona = NewUnionNullString()

		r.DescripcionZona = nil
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
		if v, ok := fields["InstruccionesEmbalaje"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.InstruccionesEmbalaje); err != nil {
			return err
		}
	} else {
		r.InstruccionesEmbalaje = NewUnionNullFloat()

		r.InstruccionesEmbalaje = nil
	}
	return nil
}
