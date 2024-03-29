// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     MantenimientoDeLoteSolicitado.avsc
 */
package ApiMantenimientoDeLoteEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type MantenimientoDeLote struct {
	Propietario string `json:"propietario"`

	Articulo string `json:"articulo"`

	Paquete *UnionNullString `json:"paquete"`

	LoteCaja *UnionNullString `json:"loteCaja"`

	LoteSecundario *UnionNullString `json:"loteSecundario"`

	LoteSap *UnionNullString `json:"loteSap"`

	FechaFabricacion *UnionNullString `json:"fechaFabricacion"`

	FechaVencimiento *UnionNullString `json:"fechaVencimiento"`

	Trazable *UnionNullString `json:"trazable"`

	Estado *UnionNullString `json:"estado"`

	Procedencia *UnionNullString `json:"procedencia"`

	CampoLibre1 *UnionNullString `json:"campoLibre1"`

	CampoLibre2 *UnionNullString `json:"campoLibre2"`

	CampoLibre3 *UnionNullString `json:"campoLibre3"`

	CampoLibre4 *UnionNullString `json:"campoLibre4"`

	CampoLibre5 *UnionNullString `json:"campoLibre5"`

	LoteExternoCliente *UnionNullString `json:"loteExternoCliente"`

	DeliverByDate *UnionNullString `json:"deliverByDate"`

	BestByDate *UnionNullString `json:"bestByDate"`

	FechaCreacion string `json:"fechaCreacion"`

	UsuarioCreacion string `json:"usuarioCreacion"`

	FechaEdicion string `json:"fechaEdicion"`

	UsuarioEdicion string `json:"usuarioEdicion"`
}

const MantenimientoDeLoteAvroCRC64Fingerprint = "\xf4\x15œ8\xef\x9d\x1c"

func NewMantenimientoDeLote() MantenimientoDeLote {
	r := MantenimientoDeLote{}
	return r
}

func DeserializeMantenimientoDeLote(r io.Reader) (MantenimientoDeLote, error) {
	t := NewMantenimientoDeLote()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeMantenimientoDeLoteFromSchema(r io.Reader, schema string) (MantenimientoDeLote, error) {
	t := NewMantenimientoDeLote()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeMantenimientoDeLote(r MantenimientoDeLote, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Articulo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Paquete, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LoteCaja, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LoteSecundario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LoteSap, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaFabricacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaVencimiento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Trazable, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Estado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Procedencia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CampoLibre1, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CampoLibre2, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CampoLibre3, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CampoLibre4, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CampoLibre5, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LoteExternoCliente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DeliverByDate, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.BestByDate, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaCreacion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.UsuarioCreacion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaEdicion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.UsuarioEdicion, w)
	if err != nil {
		return err
	}
	return err
}

func (r MantenimientoDeLote) Serialize(w io.Writer) error {
	return writeMantenimientoDeLote(r, w)
}

func (r MantenimientoDeLote) Schema() string {
	return "{\"fields\":[{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"articulo\",\"type\":\"string\"},{\"name\":\"paquete\",\"type\":[\"null\",\"string\"]},{\"name\":\"loteCaja\",\"type\":[\"null\",\"string\"]},{\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"name\":\"loteSap\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaFabricacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaVencimiento\",\"type\":[\"null\",\"string\"]},{\"name\":\"trazable\",\"type\":[\"null\",\"string\"]},{\"name\":\"estado\",\"type\":[\"null\",\"string\"]},{\"name\":\"procedencia\",\"type\":[\"null\",\"string\"]},{\"name\":\"campoLibre1\",\"type\":[\"null\",\"string\"]},{\"name\":\"campoLibre2\",\"type\":[\"null\",\"string\"]},{\"name\":\"campoLibre3\",\"type\":[\"null\",\"string\"]},{\"name\":\"campoLibre4\",\"type\":[\"null\",\"string\"]},{\"name\":\"campoLibre5\",\"type\":[\"null\",\"string\"]},{\"name\":\"loteExternoCliente\",\"type\":[\"null\",\"string\"]},{\"name\":\"deliverByDate\",\"type\":[\"null\",\"string\"]},{\"name\":\"bestByDate\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaCreacion\",\"type\":\"string\"},{\"name\":\"usuarioCreacion\",\"type\":\"string\"},{\"name\":\"fechaEdicion\",\"type\":\"string\"},{\"name\":\"usuarioEdicion\",\"type\":\"string\"}],\"name\":\"Andreani.ApiMantenimientoDeLote.Events.Record.MantenimientoDeLote\",\"type\":\"record\"}"
}

func (r MantenimientoDeLote) SchemaName() string {
	return "Andreani.ApiMantenimientoDeLote.Events.Record.MantenimientoDeLote"
}

func (_ MantenimientoDeLote) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ MantenimientoDeLote) SetInt(v int32)       { panic("Unsupported operation") }
func (_ MantenimientoDeLote) SetLong(v int64)      { panic("Unsupported operation") }
func (_ MantenimientoDeLote) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ MantenimientoDeLote) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ MantenimientoDeLote) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ MantenimientoDeLote) SetString(v string)   { panic("Unsupported operation") }
func (_ MantenimientoDeLote) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MantenimientoDeLote) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Propietario}

		return w

	case 1:
		w := types.String{Target: &r.Articulo}

		return w

	case 2:
		r.Paquete = NewUnionNullString()

		return r.Paquete
	case 3:
		r.LoteCaja = NewUnionNullString()

		return r.LoteCaja
	case 4:
		r.LoteSecundario = NewUnionNullString()

		return r.LoteSecundario
	case 5:
		r.LoteSap = NewUnionNullString()

		return r.LoteSap
	case 6:
		r.FechaFabricacion = NewUnionNullString()

		return r.FechaFabricacion
	case 7:
		r.FechaVencimiento = NewUnionNullString()

		return r.FechaVencimiento
	case 8:
		r.Trazable = NewUnionNullString()

		return r.Trazable
	case 9:
		r.Estado = NewUnionNullString()

		return r.Estado
	case 10:
		r.Procedencia = NewUnionNullString()

		return r.Procedencia
	case 11:
		r.CampoLibre1 = NewUnionNullString()

		return r.CampoLibre1
	case 12:
		r.CampoLibre2 = NewUnionNullString()

		return r.CampoLibre2
	case 13:
		r.CampoLibre3 = NewUnionNullString()

		return r.CampoLibre3
	case 14:
		r.CampoLibre4 = NewUnionNullString()

		return r.CampoLibre4
	case 15:
		r.CampoLibre5 = NewUnionNullString()

		return r.CampoLibre5
	case 16:
		r.LoteExternoCliente = NewUnionNullString()

		return r.LoteExternoCliente
	case 17:
		r.DeliverByDate = NewUnionNullString()

		return r.DeliverByDate
	case 18:
		r.BestByDate = NewUnionNullString()

		return r.BestByDate
	case 19:
		w := types.String{Target: &r.FechaCreacion}

		return w

	case 20:
		w := types.String{Target: &r.UsuarioCreacion}

		return w

	case 21:
		w := types.String{Target: &r.FechaEdicion}

		return w

	case 22:
		w := types.String{Target: &r.UsuarioEdicion}

		return w

	}
	panic("Unknown field index")
}

func (r *MantenimientoDeLote) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *MantenimientoDeLote) NullField(i int) {
	switch i {
	case 2:
		r.Paquete = nil
		return
	case 3:
		r.LoteCaja = nil
		return
	case 4:
		r.LoteSecundario = nil
		return
	case 5:
		r.LoteSap = nil
		return
	case 6:
		r.FechaFabricacion = nil
		return
	case 7:
		r.FechaVencimiento = nil
		return
	case 8:
		r.Trazable = nil
		return
	case 9:
		r.Estado = nil
		return
	case 10:
		r.Procedencia = nil
		return
	case 11:
		r.CampoLibre1 = nil
		return
	case 12:
		r.CampoLibre2 = nil
		return
	case 13:
		r.CampoLibre3 = nil
		return
	case 14:
		r.CampoLibre4 = nil
		return
	case 15:
		r.CampoLibre5 = nil
		return
	case 16:
		r.LoteExternoCliente = nil
		return
	case 17:
		r.DeliverByDate = nil
		return
	case 18:
		r.BestByDate = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ MantenimientoDeLote) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ MantenimientoDeLote) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ MantenimientoDeLote) HintSize(int)                     { panic("Unsupported operation") }
func (_ MantenimientoDeLote) Finalize()                        {}

func (_ MantenimientoDeLote) AvroCRC64Fingerprint() []byte {
	return []byte(MantenimientoDeLoteAvroCRC64Fingerprint)
}

func (r MantenimientoDeLote) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["articulo"], err = json.Marshal(r.Articulo)
	if err != nil {
		return nil, err
	}
	output["paquete"], err = json.Marshal(r.Paquete)
	if err != nil {
		return nil, err
	}
	output["loteCaja"], err = json.Marshal(r.LoteCaja)
	if err != nil {
		return nil, err
	}
	output["loteSecundario"], err = json.Marshal(r.LoteSecundario)
	if err != nil {
		return nil, err
	}
	output["loteSap"], err = json.Marshal(r.LoteSap)
	if err != nil {
		return nil, err
	}
	output["fechaFabricacion"], err = json.Marshal(r.FechaFabricacion)
	if err != nil {
		return nil, err
	}
	output["fechaVencimiento"], err = json.Marshal(r.FechaVencimiento)
	if err != nil {
		return nil, err
	}
	output["trazable"], err = json.Marshal(r.Trazable)
	if err != nil {
		return nil, err
	}
	output["estado"], err = json.Marshal(r.Estado)
	if err != nil {
		return nil, err
	}
	output["procedencia"], err = json.Marshal(r.Procedencia)
	if err != nil {
		return nil, err
	}
	output["campoLibre1"], err = json.Marshal(r.CampoLibre1)
	if err != nil {
		return nil, err
	}
	output["campoLibre2"], err = json.Marshal(r.CampoLibre2)
	if err != nil {
		return nil, err
	}
	output["campoLibre3"], err = json.Marshal(r.CampoLibre3)
	if err != nil {
		return nil, err
	}
	output["campoLibre4"], err = json.Marshal(r.CampoLibre4)
	if err != nil {
		return nil, err
	}
	output["campoLibre5"], err = json.Marshal(r.CampoLibre5)
	if err != nil {
		return nil, err
	}
	output["loteExternoCliente"], err = json.Marshal(r.LoteExternoCliente)
	if err != nil {
		return nil, err
	}
	output["deliverByDate"], err = json.Marshal(r.DeliverByDate)
	if err != nil {
		return nil, err
	}
	output["bestByDate"], err = json.Marshal(r.BestByDate)
	if err != nil {
		return nil, err
	}
	output["fechaCreacion"], err = json.Marshal(r.FechaCreacion)
	if err != nil {
		return nil, err
	}
	output["usuarioCreacion"], err = json.Marshal(r.UsuarioCreacion)
	if err != nil {
		return nil, err
	}
	output["fechaEdicion"], err = json.Marshal(r.FechaEdicion)
	if err != nil {
		return nil, err
	}
	output["usuarioEdicion"], err = json.Marshal(r.UsuarioEdicion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *MantenimientoDeLote) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["propietario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Propietario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for propietario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["articulo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Articulo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for articulo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["paquete"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Paquete); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for paquete")
	}
	val = func() json.RawMessage {
		if v, ok := fields["loteCaja"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteCaja); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for loteCaja")
	}
	val = func() json.RawMessage {
		if v, ok := fields["loteSecundario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteSecundario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for loteSecundario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["loteSap"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteSap); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for loteSap")
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaFabricacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaFabricacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fechaFabricacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaVencimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaVencimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fechaVencimiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["trazable"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Trazable); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for trazable")
	}
	val = func() json.RawMessage {
		if v, ok := fields["estado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for estado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["procedencia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Procedencia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for procedencia")
	}
	val = func() json.RawMessage {
		if v, ok := fields["campoLibre1"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CampoLibre1); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for campoLibre1")
	}
	val = func() json.RawMessage {
		if v, ok := fields["campoLibre2"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CampoLibre2); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for campoLibre2")
	}
	val = func() json.RawMessage {
		if v, ok := fields["campoLibre3"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CampoLibre3); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for campoLibre3")
	}
	val = func() json.RawMessage {
		if v, ok := fields["campoLibre4"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CampoLibre4); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for campoLibre4")
	}
	val = func() json.RawMessage {
		if v, ok := fields["campoLibre5"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CampoLibre5); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for campoLibre5")
	}
	val = func() json.RawMessage {
		if v, ok := fields["loteExternoCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteExternoCliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for loteExternoCliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["deliverByDate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DeliverByDate); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for deliverByDate")
	}
	val = func() json.RawMessage {
		if v, ok := fields["bestByDate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BestByDate); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for bestByDate")
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaCreacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaCreacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fechaCreacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["usuarioCreacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UsuarioCreacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for usuarioCreacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaEdicion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaEdicion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fechaEdicion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["usuarioEdicion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UsuarioEdicion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for usuarioEdicion")
	}
	return nil
}
