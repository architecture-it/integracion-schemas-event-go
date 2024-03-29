// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ProcesoVueltaVLMSchema.avsc
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

type NovedadDePedidoDeVLM struct {
	IdTarea string `json:"idTarea"`

	Almacen string `json:"almacen"`

	NumeroContenedorInterno string `json:"numeroContenedorInterno"`

	IdentificadorDeCaja string `json:"identificadorDeCaja"`

	Unidades int32 `json:"unidades"`

	EsTareaNueva bool `json:"esTareaNueva"`

	Series *UnionNullArrayString `json:"series"`
}

const NovedadDePedidoDeVLMAvroCRC64Fingerprint = "I\xdfx(\x0e\xfeė"

func NewNovedadDePedidoDeVLM() NovedadDePedidoDeVLM {
	r := NovedadDePedidoDeVLM{}
	r.Series = nil
	return r
}

func DeserializeNovedadDePedidoDeVLM(r io.Reader) (NovedadDePedidoDeVLM, error) {
	t := NewNovedadDePedidoDeVLM()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNovedadDePedidoDeVLMFromSchema(r io.Reader, schema string) (NovedadDePedidoDeVLM, error) {
	t := NewNovedadDePedidoDeVLM()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNovedadDePedidoDeVLM(r NovedadDePedidoDeVLM, w io.Writer) error {
	var err error
	err = vm.WriteString(r.IdTarea, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Almacen, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroContenedorInterno, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.IdentificadorDeCaja, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Unidades, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.EsTareaNueva, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayString(r.Series, w)
	if err != nil {
		return err
	}
	return err
}

func (r NovedadDePedidoDeVLM) Serialize(w io.Writer) error {
	return writeNovedadDePedidoDeVLM(r, w)
}

func (r NovedadDePedidoDeVLM) Schema() string {
	return "{\"fields\":[{\"name\":\"idTarea\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"name\":\"numeroContenedorInterno\",\"type\":\"string\"},{\"name\":\"identificadorDeCaja\",\"type\":\"string\"},{\"name\":\"unidades\",\"type\":\"int\"},{\"name\":\"esTareaNueva\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"series\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]}],\"name\":\"Andreani.PedidosVlmSce.Events.ProcesoVuelta.NovedadDePedidoDeVLM\",\"type\":\"record\"}"
}

func (r NovedadDePedidoDeVLM) SchemaName() string {
	return "Andreani.PedidosVlmSce.Events.ProcesoVuelta.NovedadDePedidoDeVLM"
}

func (_ NovedadDePedidoDeVLM) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NovedadDePedidoDeVLM) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NovedadDePedidoDeVLM) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NovedadDePedidoDeVLM) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NovedadDePedidoDeVLM) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NovedadDePedidoDeVLM) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NovedadDePedidoDeVLM) SetString(v string)   { panic("Unsupported operation") }
func (_ NovedadDePedidoDeVLM) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NovedadDePedidoDeVLM) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.IdTarea}

		return w

	case 1:
		w := types.String{Target: &r.Almacen}

		return w

	case 2:
		w := types.String{Target: &r.NumeroContenedorInterno}

		return w

	case 3:
		w := types.String{Target: &r.IdentificadorDeCaja}

		return w

	case 4:
		w := types.Int{Target: &r.Unidades}

		return w

	case 5:
		w := types.Boolean{Target: &r.EsTareaNueva}

		return w

	case 6:
		r.Series = NewUnionNullArrayString()

		return r.Series
	}
	panic("Unknown field index")
}

func (r *NovedadDePedidoDeVLM) SetDefault(i int) {
	switch i {
	case 6:
		r.Series = nil
		return
	}
	panic("Unknown field index")
}

func (r *NovedadDePedidoDeVLM) NullField(i int) {
	switch i {
	case 6:
		r.Series = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ NovedadDePedidoDeVLM) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NovedadDePedidoDeVLM) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NovedadDePedidoDeVLM) HintSize(int)                     { panic("Unsupported operation") }
func (_ NovedadDePedidoDeVLM) Finalize()                        {}

func (_ NovedadDePedidoDeVLM) AvroCRC64Fingerprint() []byte {
	return []byte(NovedadDePedidoDeVLMAvroCRC64Fingerprint)
}

func (r NovedadDePedidoDeVLM) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["idTarea"], err = json.Marshal(r.IdTarea)
	if err != nil {
		return nil, err
	}
	output["almacen"], err = json.Marshal(r.Almacen)
	if err != nil {
		return nil, err
	}
	output["numeroContenedorInterno"], err = json.Marshal(r.NumeroContenedorInterno)
	if err != nil {
		return nil, err
	}
	output["identificadorDeCaja"], err = json.Marshal(r.IdentificadorDeCaja)
	if err != nil {
		return nil, err
	}
	output["unidades"], err = json.Marshal(r.Unidades)
	if err != nil {
		return nil, err
	}
	output["esTareaNueva"], err = json.Marshal(r.EsTareaNueva)
	if err != nil {
		return nil, err
	}
	output["series"], err = json.Marshal(r.Series)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NovedadDePedidoDeVLM) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["idTarea"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdTarea); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idTarea")
	}
	val = func() json.RawMessage {
		if v, ok := fields["almacen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Almacen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for almacen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroContenedorInterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroContenedorInterno); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroContenedorInterno")
	}
	val = func() json.RawMessage {
		if v, ok := fields["identificadorDeCaja"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdentificadorDeCaja); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for identificadorDeCaja")
	}
	val = func() json.RawMessage {
		if v, ok := fields["unidades"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Unidades); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for unidades")
	}
	val = func() json.RawMessage {
		if v, ok := fields["esTareaNueva"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EsTareaNueva); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for esTareaNueva")
	}
	val = func() json.RawMessage {
		if v, ok := fields["series"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Series); err != nil {
			return err
		}
	} else {
		r.Series = NewUnionNullArrayString()

		r.Series = nil
	}
	return nil
}
