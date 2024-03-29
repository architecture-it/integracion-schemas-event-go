// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     MantenimientoDeArticuloSolicitado.avsc
 */
package ApiMantenimientoDeProductoEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Lote struct {
	Codigo string `json:"codigo"`

	LoteDeFabricante string `json:"loteDeFabricante"`

	LoteSecundario string `json:"loteSecundario"`

	FechaDeVencimiento string `json:"fechaDeVencimiento"`

	OtrosDatos []Metadato `json:"otrosDatos"`
}

const LoteAvroCRC64Fingerprint = "\xaf\xbbC\x809\x0e=\x80"

func NewLote() Lote {
	r := Lote{}
	r.OtrosDatos = make([]Metadato, 0)

	return r
}

func DeserializeLote(r io.Reader) (Lote, error) {
	t := NewLote()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeLoteFromSchema(r io.Reader, schema string) (Lote, error) {
	t := NewLote()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeLote(r Lote, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Codigo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.LoteDeFabricante, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.LoteSecundario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaDeVencimiento, w)
	if err != nil {
		return err
	}
	err = writeArrayMetadato(r.OtrosDatos, w)
	if err != nil {
		return err
	}
	return err
}

func (r Lote) Serialize(w io.Writer) error {
	return writeLote(r, w)
}

func (r Lote) Schema() string {
	return "{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"loteDeFabricante\",\"type\":\"string\"},{\"name\":\"loteSecundario\",\"type\":\"string\"},{\"name\":\"fechaDeVencimiento\",\"type\":\"string\"},{\"name\":\"otrosDatos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.ApiMantenimientoDeProducto.Events.Record.Lote\",\"type\":\"record\"}"
}

func (r Lote) SchemaName() string {
	return "Andreani.ApiMantenimientoDeProducto.Events.Record.Lote"
}

func (_ Lote) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Lote) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Lote) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Lote) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Lote) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Lote) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Lote) SetString(v string)   { panic("Unsupported operation") }
func (_ Lote) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Lote) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Codigo}

		return w

	case 1:
		w := types.String{Target: &r.LoteDeFabricante}

		return w

	case 2:
		w := types.String{Target: &r.LoteSecundario}

		return w

	case 3:
		w := types.String{Target: &r.FechaDeVencimiento}

		return w

	case 4:
		r.OtrosDatos = make([]Metadato, 0)

		w := ArrayMetadatoWrapper{Target: &r.OtrosDatos}

		return w

	}
	panic("Unknown field index")
}

func (r *Lote) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Lote) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Lote) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Lote) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Lote) HintSize(int)                     { panic("Unsupported operation") }
func (_ Lote) Finalize()                        {}

func (_ Lote) AvroCRC64Fingerprint() []byte {
	return []byte(LoteAvroCRC64Fingerprint)
}

func (r Lote) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["codigo"], err = json.Marshal(r.Codigo)
	if err != nil {
		return nil, err
	}
	output["loteDeFabricante"], err = json.Marshal(r.LoteDeFabricante)
	if err != nil {
		return nil, err
	}
	output["loteSecundario"], err = json.Marshal(r.LoteSecundario)
	if err != nil {
		return nil, err
	}
	output["fechaDeVencimiento"], err = json.Marshal(r.FechaDeVencimiento)
	if err != nil {
		return nil, err
	}
	output["otrosDatos"], err = json.Marshal(r.OtrosDatos)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Lote) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["codigo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Codigo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["loteDeFabricante"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteDeFabricante); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for loteDeFabricante")
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
		if v, ok := fields["fechaDeVencimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaDeVencimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fechaDeVencimiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["otrosDatos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OtrosDatos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for otrosDatos")
	}
	return nil
}
