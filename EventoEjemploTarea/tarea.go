// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Tarea.avsc
 */
package EventoEjemploTareaEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Tarea struct {
	IdPedido string `json:"IdPedido"`

	Descripcion string `json:"Descripcion"`

	TipoOperacion string `json:"TipoOperacion"`

	TipoEjecucion string `json:"TipoEjecucion"`

	Prioridad string `json:"Prioridad"`

	Cliente string `json:"Cliente"`

	Almacen string `json:"Almacen"`

	TipoPedido string `json:"TipoPedido"`

	CodigoArticulo string `json:"CodigoArticulo"`

	CodigoLoteArticulo string `json:"CodigoLoteArticulo"`

	FechaVencimientoLote *UnionNullLong `json:"FechaVencimientoLote"`

	Unidades float32 `json:"Unidades"`

	CategoriaDeUbicacion string `json:"CategoriaDeUbicacion"`

	IdDeTarea string `json:"IdDeTarea"`

	NombreWMSOrigen string `json:"NombreWMSOrigen"`

	NombreWMSDestino string `json:"NombreWMSDestino"`
}

const TareaAvroCRC64Fingerprint = "Gi\x05\x91\x9a\xad\x16\xaf"

func NewTarea() Tarea {
	r := Tarea{}
	return r
}

func DeserializeTarea(r io.Reader) (Tarea, error) {
	t := NewTarea()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeTareaFromSchema(r io.Reader, schema string) (Tarea, error) {
	t := NewTarea()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeTarea(r Tarea, w io.Writer) error {
	var err error
	err = vm.WriteString(r.IdPedido, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Descripcion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TipoOperacion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TipoEjecucion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Prioridad, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Cliente, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Almacen, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TipoPedido, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoArticulo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoLoteArticulo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaVencimientoLote, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.Unidades, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CategoriaDeUbicacion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.IdDeTarea, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NombreWMSOrigen, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NombreWMSDestino, w)
	if err != nil {
		return err
	}
	return err
}

func (r Tarea) Serialize(w io.Writer) error {
	return writeTarea(r, w)
}

func (r Tarea) Schema() string {
	return "{\"fields\":[{\"name\":\"IdPedido\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"name\":\"TipoOperacion\",\"type\":\"string\"},{\"name\":\"TipoEjecucion\",\"type\":\"string\"},{\"name\":\"Prioridad\",\"type\":\"string\"},{\"name\":\"Cliente\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"TipoPedido\",\"type\":\"string\"},{\"name\":\"CodigoArticulo\",\"type\":\"string\"},{\"name\":\"CodigoLoteArticulo\",\"type\":\"string\"},{\"name\":\"FechaVencimientoLote\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"Unidades\",\"type\":\"float\"},{\"name\":\"CategoriaDeUbicacion\",\"type\":\"string\"},{\"name\":\"IdDeTarea\",\"type\":\"string\"},{\"name\":\"NombreWMSOrigen\",\"type\":\"string\"},{\"name\":\"NombreWMSDestino\",\"type\":\"string\"}],\"name\":\"Andreani.EventoEjemploTarea.Events.Common.Tarea\",\"type\":\"record\"}"
}

func (r Tarea) SchemaName() string {
	return "Andreani.EventoEjemploTarea.Events.Common.Tarea"
}

func (_ Tarea) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Tarea) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Tarea) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Tarea) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Tarea) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Tarea) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Tarea) SetString(v string)   { panic("Unsupported operation") }
func (_ Tarea) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Tarea) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.IdPedido}

		return w

	case 1:
		w := types.String{Target: &r.Descripcion}

		return w

	case 2:
		w := types.String{Target: &r.TipoOperacion}

		return w

	case 3:
		w := types.String{Target: &r.TipoEjecucion}

		return w

	case 4:
		w := types.String{Target: &r.Prioridad}

		return w

	case 5:
		w := types.String{Target: &r.Cliente}

		return w

	case 6:
		w := types.String{Target: &r.Almacen}

		return w

	case 7:
		w := types.String{Target: &r.TipoPedido}

		return w

	case 8:
		w := types.String{Target: &r.CodigoArticulo}

		return w

	case 9:
		w := types.String{Target: &r.CodigoLoteArticulo}

		return w

	case 10:
		r.FechaVencimientoLote = NewUnionNullLong()

		return r.FechaVencimientoLote
	case 11:
		w := types.Float{Target: &r.Unidades}

		return w

	case 12:
		w := types.String{Target: &r.CategoriaDeUbicacion}

		return w

	case 13:
		w := types.String{Target: &r.IdDeTarea}

		return w

	case 14:
		w := types.String{Target: &r.NombreWMSOrigen}

		return w

	case 15:
		w := types.String{Target: &r.NombreWMSDestino}

		return w

	}
	panic("Unknown field index")
}

func (r *Tarea) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Tarea) NullField(i int) {
	switch i {
	case 10:
		r.FechaVencimientoLote = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Tarea) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Tarea) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Tarea) HintSize(int)                     { panic("Unsupported operation") }
func (_ Tarea) Finalize()                        {}

func (_ Tarea) AvroCRC64Fingerprint() []byte {
	return []byte(TareaAvroCRC64Fingerprint)
}

func (r Tarea) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["IdPedido"], err = json.Marshal(r.IdPedido)
	if err != nil {
		return nil, err
	}
	output["Descripcion"], err = json.Marshal(r.Descripcion)
	if err != nil {
		return nil, err
	}
	output["TipoOperacion"], err = json.Marshal(r.TipoOperacion)
	if err != nil {
		return nil, err
	}
	output["TipoEjecucion"], err = json.Marshal(r.TipoEjecucion)
	if err != nil {
		return nil, err
	}
	output["Prioridad"], err = json.Marshal(r.Prioridad)
	if err != nil {
		return nil, err
	}
	output["Cliente"], err = json.Marshal(r.Cliente)
	if err != nil {
		return nil, err
	}
	output["Almacen"], err = json.Marshal(r.Almacen)
	if err != nil {
		return nil, err
	}
	output["TipoPedido"], err = json.Marshal(r.TipoPedido)
	if err != nil {
		return nil, err
	}
	output["CodigoArticulo"], err = json.Marshal(r.CodigoArticulo)
	if err != nil {
		return nil, err
	}
	output["CodigoLoteArticulo"], err = json.Marshal(r.CodigoLoteArticulo)
	if err != nil {
		return nil, err
	}
	output["FechaVencimientoLote"], err = json.Marshal(r.FechaVencimientoLote)
	if err != nil {
		return nil, err
	}
	output["Unidades"], err = json.Marshal(r.Unidades)
	if err != nil {
		return nil, err
	}
	output["CategoriaDeUbicacion"], err = json.Marshal(r.CategoriaDeUbicacion)
	if err != nil {
		return nil, err
	}
	output["IdDeTarea"], err = json.Marshal(r.IdDeTarea)
	if err != nil {
		return nil, err
	}
	output["NombreWMSOrigen"], err = json.Marshal(r.NombreWMSOrigen)
	if err != nil {
		return nil, err
	}
	output["NombreWMSDestino"], err = json.Marshal(r.NombreWMSDestino)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Tarea) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["IdPedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdPedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IdPedido")
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
		if v, ok := fields["TipoOperacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoOperacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoOperacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoEjecucion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoEjecucion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoEjecucion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Prioridad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Prioridad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Prioridad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Cliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Cliente")
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
		if v, ok := fields["TipoPedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoPedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoPedido")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoArticulo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoArticulo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodigoArticulo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoLoteArticulo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoLoteArticulo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodigoLoteArticulo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaVencimientoLote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaVencimientoLote); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaVencimientoLote")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Unidades"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Unidades); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Unidades")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CategoriaDeUbicacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CategoriaDeUbicacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CategoriaDeUbicacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["IdDeTarea"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdDeTarea); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IdDeTarea")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NombreWMSOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NombreWMSOrigen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NombreWMSOrigen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NombreWMSDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NombreWMSDestino); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NombreWMSDestino")
	}
	return nil
}
