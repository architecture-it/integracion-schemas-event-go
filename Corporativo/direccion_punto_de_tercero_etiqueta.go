// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PuntoDeTerceroEtiqueta.avsc
 */
package CorporativoEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type DireccionPuntoDeTerceroEtiqueta struct {
	Calle string `json:"Calle"`

	Numero string `json:"Numero"`

	Provincia string `json:"Provincia"`

	Localidad string `json:"Localidad"`

	Region string `json:"Region"`

	Pais string `json:"Pais"`

	CodigoPostal string `json:"CodigoPostal"`
}

const DireccionPuntoDeTerceroEtiquetaAvroCRC64Fingerprint = "{\x19\xe7w\xe2aZ\xfc"

func NewDireccionPuntoDeTerceroEtiqueta() DireccionPuntoDeTerceroEtiqueta {
	r := DireccionPuntoDeTerceroEtiqueta{}
	return r
}

func DeserializeDireccionPuntoDeTerceroEtiqueta(r io.Reader) (DireccionPuntoDeTerceroEtiqueta, error) {
	t := NewDireccionPuntoDeTerceroEtiqueta()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDireccionPuntoDeTerceroEtiquetaFromSchema(r io.Reader, schema string) (DireccionPuntoDeTerceroEtiqueta, error) {
	t := NewDireccionPuntoDeTerceroEtiqueta()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDireccionPuntoDeTerceroEtiqueta(r DireccionPuntoDeTerceroEtiqueta, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Calle, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Numero, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Provincia, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Localidad, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Region, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Pais, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoPostal, w)
	if err != nil {
		return err
	}
	return err
}

func (r DireccionPuntoDeTerceroEtiqueta) Serialize(w io.Writer) error {
	return writeDireccionPuntoDeTerceroEtiqueta(r, w)
}

func (r DireccionPuntoDeTerceroEtiqueta) Schema() string {
	return "{\"fields\":[{\"name\":\"Calle\",\"type\":\"string\"},{\"name\":\"Numero\",\"type\":\"string\"},{\"name\":\"Provincia\",\"type\":\"string\"},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"Region\",\"type\":\"string\"},{\"name\":\"Pais\",\"type\":\"string\"},{\"name\":\"CodigoPostal\",\"type\":\"string\"}],\"name\":\"Andreani.Corporativo.Events.Record.DireccionPuntoDeTerceroEtiqueta\",\"type\":\"record\"}"
}

func (r DireccionPuntoDeTerceroEtiqueta) SchemaName() string {
	return "Andreani.Corporativo.Events.Record.DireccionPuntoDeTerceroEtiqueta"
}

func (_ DireccionPuntoDeTerceroEtiqueta) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DireccionPuntoDeTerceroEtiqueta) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DireccionPuntoDeTerceroEtiqueta) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DireccionPuntoDeTerceroEtiqueta) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DireccionPuntoDeTerceroEtiqueta) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DireccionPuntoDeTerceroEtiqueta) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DireccionPuntoDeTerceroEtiqueta) SetString(v string)   { panic("Unsupported operation") }
func (_ DireccionPuntoDeTerceroEtiqueta) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DireccionPuntoDeTerceroEtiqueta) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Calle}

		return w

	case 1:
		w := types.String{Target: &r.Numero}

		return w

	case 2:
		w := types.String{Target: &r.Provincia}

		return w

	case 3:
		w := types.String{Target: &r.Localidad}

		return w

	case 4:
		w := types.String{Target: &r.Region}

		return w

	case 5:
		w := types.String{Target: &r.Pais}

		return w

	case 6:
		w := types.String{Target: &r.CodigoPostal}

		return w

	}
	panic("Unknown field index")
}

func (r *DireccionPuntoDeTerceroEtiqueta) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *DireccionPuntoDeTerceroEtiqueta) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ DireccionPuntoDeTerceroEtiqueta) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ DireccionPuntoDeTerceroEtiqueta) AppendArray() types.Field { panic("Unsupported operation") }
func (_ DireccionPuntoDeTerceroEtiqueta) HintSize(int)             { panic("Unsupported operation") }
func (_ DireccionPuntoDeTerceroEtiqueta) Finalize()                {}

func (_ DireccionPuntoDeTerceroEtiqueta) AvroCRC64Fingerprint() []byte {
	return []byte(DireccionPuntoDeTerceroEtiquetaAvroCRC64Fingerprint)
}

func (r DireccionPuntoDeTerceroEtiqueta) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Calle"], err = json.Marshal(r.Calle)
	if err != nil {
		return nil, err
	}
	output["Numero"], err = json.Marshal(r.Numero)
	if err != nil {
		return nil, err
	}
	output["Provincia"], err = json.Marshal(r.Provincia)
	if err != nil {
		return nil, err
	}
	output["Localidad"], err = json.Marshal(r.Localidad)
	if err != nil {
		return nil, err
	}
	output["Region"], err = json.Marshal(r.Region)
	if err != nil {
		return nil, err
	}
	output["Pais"], err = json.Marshal(r.Pais)
	if err != nil {
		return nil, err
	}
	output["CodigoPostal"], err = json.Marshal(r.CodigoPostal)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DireccionPuntoDeTerceroEtiqueta) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Calle"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Calle); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Calle")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Numero"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Numero); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Numero")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Provincia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Provincia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Provincia")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Localidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Localidad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Localidad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Region"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Region); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Region")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Pais"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Pais); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Pais")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoPostal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoPostal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodigoPostal")
	}
	return nil
}
