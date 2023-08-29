// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CambioEstadoTarea.avsc
 */
package WosPickingEstadoTareaEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type CambioEstadoTarea struct {
	TareaId int32 `json:"TareaId"`

	PedidoId int32 `json:"PedidoId"`

	TareaIdExterno string `json:"TareaIdExterno"`

	PedidoIdExterno string `json:"PedidoIdExterno"`

	PlantaOperacionId int32 `json:"PlantaOperacionId"`

	EstadoId string `json:"EstadoId"`

	Fecha int64 `json:"Fecha"`

	MotivoId *UnionNullString `json:"MotivoId"`

	UsuarioId *UnionNullInt `json:"UsuarioId"`

	ContenedorId *UnionNullString `json:"ContenedorId"`

	CantidadPickeados *UnionNullInt `json:"CantidadPickeados"`

	TiempoReal *UnionNullInt `json:"TiempoReal"`
}

const CambioEstadoTareaAvroCRC64Fingerprint = "\xfa=\xdd -\xf0\\\x8d"

func NewCambioEstadoTarea() CambioEstadoTarea {
	r := CambioEstadoTarea{}
	r.MotivoId = nil
	r.UsuarioId = nil
	r.ContenedorId = nil
	r.CantidadPickeados = nil
	r.TiempoReal = nil
	return r
}

func DeserializeCambioEstadoTarea(r io.Reader) (CambioEstadoTarea, error) {
	t := NewCambioEstadoTarea()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCambioEstadoTareaFromSchema(r io.Reader, schema string) (CambioEstadoTarea, error) {
	t := NewCambioEstadoTarea()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCambioEstadoTarea(r CambioEstadoTarea, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.TareaId, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.PedidoId, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TareaIdExterno, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.PedidoIdExterno, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.PlantaOperacionId, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EstadoId, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Fecha, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.MotivoId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.UsuarioId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ContenedorId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.CantidadPickeados, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.TiempoReal, w)
	if err != nil {
		return err
	}
	return err
}

func (r CambioEstadoTarea) Serialize(w io.Writer) error {
	return writeCambioEstadoTarea(r, w)
}

func (r CambioEstadoTarea) Schema() string {
	return "{\"fields\":[{\"name\":\"TareaId\",\"type\":\"int\"},{\"name\":\"PedidoId\",\"type\":\"int\"},{\"name\":\"TareaIdExterno\",\"type\":\"string\"},{\"name\":\"PedidoIdExterno\",\"type\":\"string\"},{\"name\":\"PlantaOperacionId\",\"type\":\"int\"},{\"name\":\"EstadoId\",\"type\":\"string\"},{\"name\":\"Fecha\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"MotivoId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"UsuarioId\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"ContenedorId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CantidadPickeados\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"TiempoReal\",\"type\":[\"null\",\"int\"]}],\"name\":\"Andreani.WosPickingEstadoTarea.Events.Record.CambioEstadoTarea\",\"type\":\"record\"}"
}

func (r CambioEstadoTarea) SchemaName() string {
	return "Andreani.WosPickingEstadoTarea.Events.Record.CambioEstadoTarea"
}

func (_ CambioEstadoTarea) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CambioEstadoTarea) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CambioEstadoTarea) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CambioEstadoTarea) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CambioEstadoTarea) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CambioEstadoTarea) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CambioEstadoTarea) SetString(v string)   { panic("Unsupported operation") }
func (_ CambioEstadoTarea) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CambioEstadoTarea) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.TareaId}

		return w

	case 1:
		w := types.Int{Target: &r.PedidoId}

		return w

	case 2:
		w := types.String{Target: &r.TareaIdExterno}

		return w

	case 3:
		w := types.String{Target: &r.PedidoIdExterno}

		return w

	case 4:
		w := types.Int{Target: &r.PlantaOperacionId}

		return w

	case 5:
		w := types.String{Target: &r.EstadoId}

		return w

	case 6:
		w := types.Long{Target: &r.Fecha}

		return w

	case 7:
		r.MotivoId = NewUnionNullString()

		return r.MotivoId
	case 8:
		r.UsuarioId = NewUnionNullInt()

		return r.UsuarioId
	case 9:
		r.ContenedorId = NewUnionNullString()

		return r.ContenedorId
	case 10:
		r.CantidadPickeados = NewUnionNullInt()

		return r.CantidadPickeados
	case 11:
		r.TiempoReal = NewUnionNullInt()

		return r.TiempoReal
	}
	panic("Unknown field index")
}

func (r *CambioEstadoTarea) SetDefault(i int) {
	switch i {
	case 7:
		r.MotivoId = nil
		return
	case 8:
		r.UsuarioId = nil
		return
	case 9:
		r.ContenedorId = nil
		return
	case 10:
		r.CantidadPickeados = nil
		return
	case 11:
		r.TiempoReal = nil
		return
	}
	panic("Unknown field index")
}

func (r *CambioEstadoTarea) NullField(i int) {
	switch i {
	case 7:
		r.MotivoId = nil
		return
	case 8:
		r.UsuarioId = nil
		return
	case 9:
		r.ContenedorId = nil
		return
	case 10:
		r.CantidadPickeados = nil
		return
	case 11:
		r.TiempoReal = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ CambioEstadoTarea) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CambioEstadoTarea) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CambioEstadoTarea) HintSize(int)                     { panic("Unsupported operation") }
func (_ CambioEstadoTarea) Finalize()                        {}

func (_ CambioEstadoTarea) AvroCRC64Fingerprint() []byte {
	return []byte(CambioEstadoTareaAvroCRC64Fingerprint)
}

func (r CambioEstadoTarea) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["TareaId"], err = json.Marshal(r.TareaId)
	if err != nil {
		return nil, err
	}
	output["PedidoId"], err = json.Marshal(r.PedidoId)
	if err != nil {
		return nil, err
	}
	output["TareaIdExterno"], err = json.Marshal(r.TareaIdExterno)
	if err != nil {
		return nil, err
	}
	output["PedidoIdExterno"], err = json.Marshal(r.PedidoIdExterno)
	if err != nil {
		return nil, err
	}
	output["PlantaOperacionId"], err = json.Marshal(r.PlantaOperacionId)
	if err != nil {
		return nil, err
	}
	output["EstadoId"], err = json.Marshal(r.EstadoId)
	if err != nil {
		return nil, err
	}
	output["Fecha"], err = json.Marshal(r.Fecha)
	if err != nil {
		return nil, err
	}
	output["MotivoId"], err = json.Marshal(r.MotivoId)
	if err != nil {
		return nil, err
	}
	output["UsuarioId"], err = json.Marshal(r.UsuarioId)
	if err != nil {
		return nil, err
	}
	output["ContenedorId"], err = json.Marshal(r.ContenedorId)
	if err != nil {
		return nil, err
	}
	output["CantidadPickeados"], err = json.Marshal(r.CantidadPickeados)
	if err != nil {
		return nil, err
	}
	output["TiempoReal"], err = json.Marshal(r.TiempoReal)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CambioEstadoTarea) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["TareaId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TareaId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TareaId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PedidoId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PedidoId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PedidoId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TareaIdExterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TareaIdExterno); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TareaIdExterno")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PedidoIdExterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PedidoIdExterno); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PedidoIdExterno")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PlantaOperacionId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PlantaOperacionId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PlantaOperacionId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EstadoId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstadoId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EstadoId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Fecha"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Fecha); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Fecha")
	}
	val = func() json.RawMessage {
		if v, ok := fields["MotivoId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MotivoId); err != nil {
			return err
		}
	} else {
		r.MotivoId = NewUnionNullString()

		r.MotivoId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["UsuarioId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UsuarioId); err != nil {
			return err
		}
	} else {
		r.UsuarioId = NewUnionNullInt()

		r.UsuarioId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ContenedorId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ContenedorId); err != nil {
			return err
		}
	} else {
		r.ContenedorId = NewUnionNullString()

		r.ContenedorId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["CantidadPickeados"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadPickeados); err != nil {
			return err
		}
	} else {
		r.CantidadPickeados = NewUnionNullInt()

		r.CantidadPickeados = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["TiempoReal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TiempoReal); err != nil {
			return err
		}
	} else {
		r.TiempoReal = NewUnionNullInt()

		r.TiempoReal = nil
	}
	return nil
}