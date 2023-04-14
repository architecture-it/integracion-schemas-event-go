// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ProcesoIdaVLMSchema.avsc
 */
package PedidosVlmSceEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type PedidoDeVLM struct {
	IdPedido string `json:"idPedido"`

	Descripcion string `json:"descripcion"`

	TipoOperacion string `json:"tipoOperacion"`

	TipoEjecucion string `json:"tipoEjecucion"`

	Prioridad int32 `json:"prioridad"`

	Cliente string `json:"cliente"`

	AlmacenDeOrigen string `json:"almacenDeOrigen"`

	TipoPedido string `json:"tipoPedido"`

	DetallesDePedido []DetallesDePedido `json:"detallesDePedido"`
}

const PedidoDeVLMAvroCRC64Fingerprint = "V(\xb9_\x11\fw\xc3"

func NewPedidoDeVLM() PedidoDeVLM {
	r := PedidoDeVLM{}
	r.DetallesDePedido = make([]DetallesDePedido, 0)

	return r
}

func DeserializePedidoDeVLM(r io.Reader) (PedidoDeVLM, error) {
	t := NewPedidoDeVLM()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePedidoDeVLMFromSchema(r io.Reader, schema string) (PedidoDeVLM, error) {
	t := NewPedidoDeVLM()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePedidoDeVLM(r PedidoDeVLM, w io.Writer) error {
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
	err = vm.WriteInt(r.Prioridad, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Cliente, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.AlmacenDeOrigen, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TipoPedido, w)
	if err != nil {
		return err
	}
	err = writeArrayDetallesDePedido(r.DetallesDePedido, w)
	if err != nil {
		return err
	}
	return err
}

func (r PedidoDeVLM) Serialize(w io.Writer) error {
	return writePedidoDeVLM(r, w)
}

func (r PedidoDeVLM) Schema() string {
	return "{\"fields\":[{\"name\":\"idPedido\",\"type\":\"string\"},{\"name\":\"descripcion\",\"type\":\"string\"},{\"name\":\"tipoOperacion\",\"type\":\"string\"},{\"name\":\"tipoEjecucion\",\"type\":\"string\"},{\"name\":\"prioridad\",\"type\":\"int\"},{\"name\":\"cliente\",\"type\":\"string\"},{\"name\":\"almacenDeOrigen\",\"type\":\"string\"},{\"name\":\"tipoPedido\",\"type\":\"string\"},{\"name\":\"detallesDePedido\",\"type\":{\"items\":{\"fields\":[{\"name\":\"articulo\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"lote\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"fechadeVencimiento\",\"type\":\"string\"}],\"name\":\"Lote\",\"type\":\"record\"}}],\"name\":\"Articulo\",\"type\":\"record\"}},{\"name\":\"unidades\",\"type\":\"int\"},{\"name\":\"categoriaDeUbicacionDeOrigen\",\"type\":\"string\"},{\"name\":\"idDeTarea\",\"type\":\"string\"},{\"name\":\"almacenDeOrigen\",\"type\":\"string\"},{\"name\":\"idPedido\",\"type\":\"string\"}],\"name\":\"DetallesDePedido\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.PedidosVLM.Events.ProcesoIda.PedidoDeVLM\",\"type\":\"record\"}"
}

func (r PedidoDeVLM) SchemaName() string {
	return "Andreani.PedidosVLM.Events.ProcesoIda.PedidoDeVLM"
}

func (_ PedidoDeVLM) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ PedidoDeVLM) SetInt(v int32)       { panic("Unsupported operation") }
func (_ PedidoDeVLM) SetLong(v int64)      { panic("Unsupported operation") }
func (_ PedidoDeVLM) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ PedidoDeVLM) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ PedidoDeVLM) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ PedidoDeVLM) SetString(v string)   { panic("Unsupported operation") }
func (_ PedidoDeVLM) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *PedidoDeVLM) Get(i int) types.Field {
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
		w := types.Int{Target: &r.Prioridad}

		return w

	case 5:
		w := types.String{Target: &r.Cliente}

		return w

	case 6:
		w := types.String{Target: &r.AlmacenDeOrigen}

		return w

	case 7:
		w := types.String{Target: &r.TipoPedido}

		return w

	case 8:
		r.DetallesDePedido = make([]DetallesDePedido, 0)

		w := ArrayDetallesDePedidoWrapper{Target: &r.DetallesDePedido}

		return w

	}
	panic("Unknown field index")
}

func (r *PedidoDeVLM) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *PedidoDeVLM) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ PedidoDeVLM) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ PedidoDeVLM) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ PedidoDeVLM) HintSize(int)                     { panic("Unsupported operation") }
func (_ PedidoDeVLM) Finalize()                        {}

func (_ PedidoDeVLM) AvroCRC64Fingerprint() []byte {
	return []byte(PedidoDeVLMAvroCRC64Fingerprint)
}

func (r PedidoDeVLM) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["idPedido"], err = json.Marshal(r.IdPedido)
	if err != nil {
		return nil, err
	}
	output["descripcion"], err = json.Marshal(r.Descripcion)
	if err != nil {
		return nil, err
	}
	output["tipoOperacion"], err = json.Marshal(r.TipoOperacion)
	if err != nil {
		return nil, err
	}
	output["tipoEjecucion"], err = json.Marshal(r.TipoEjecucion)
	if err != nil {
		return nil, err
	}
	output["prioridad"], err = json.Marshal(r.Prioridad)
	if err != nil {
		return nil, err
	}
	output["cliente"], err = json.Marshal(r.Cliente)
	if err != nil {
		return nil, err
	}
	output["almacenDeOrigen"], err = json.Marshal(r.AlmacenDeOrigen)
	if err != nil {
		return nil, err
	}
	output["tipoPedido"], err = json.Marshal(r.TipoPedido)
	if err != nil {
		return nil, err
	}
	output["detallesDePedido"], err = json.Marshal(r.DetallesDePedido)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *PedidoDeVLM) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["idPedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdPedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idPedido")
	}
	val = func() json.RawMessage {
		if v, ok := fields["descripcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Descripcion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for descripcion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoOperacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoOperacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tipoOperacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoEjecucion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoEjecucion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tipoEjecucion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["prioridad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Prioridad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for prioridad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["cliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["almacenDeOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AlmacenDeOrigen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for almacenDeOrigen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoPedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoPedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tipoPedido")
	}
	val = func() json.RawMessage {
		if v, ok := fields["detallesDePedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DetallesDePedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for detallesDePedido")
	}
	return nil
}
