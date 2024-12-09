// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     RemitenteEtiqueta.avsc
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

type DireccionRemitenteEtiqueta struct {
	Calle string `json:"Calle"`

	Numero string `json:"Numero"`

	Piso string `json:"Piso"`

	Unidad string `json:"Unidad"`

	Localidad string `json:"Localidad"`

	CodigoPostal string `json:"CodigoPostal"`

	Provincia string `json:"Provincia"`

	DatosAdicionales string `json:"DatosAdicionales"`

	Latitud string `json:"Latitud"`

	Longitud string `json:"Longitud"`
}

const DireccionRemitenteEtiquetaAvroCRC64Fingerprint = "\x8d\x83\x914\x97\x02\xaf\x86"

func NewDireccionRemitenteEtiqueta() DireccionRemitenteEtiqueta {
	r := DireccionRemitenteEtiqueta{}
	return r
}

func DeserializeDireccionRemitenteEtiqueta(r io.Reader) (DireccionRemitenteEtiqueta, error) {
	t := NewDireccionRemitenteEtiqueta()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDireccionRemitenteEtiquetaFromSchema(r io.Reader, schema string) (DireccionRemitenteEtiqueta, error) {
	t := NewDireccionRemitenteEtiqueta()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDireccionRemitenteEtiqueta(r DireccionRemitenteEtiqueta, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Calle, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Numero, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Piso, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Unidad, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Localidad, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoPostal, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Provincia, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.DatosAdicionales, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Latitud, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Longitud, w)
	if err != nil {
		return err
	}
	return err
}

func (r DireccionRemitenteEtiqueta) Serialize(w io.Writer) error {
	return writeDireccionRemitenteEtiqueta(r, w)
}

func (r DireccionRemitenteEtiqueta) Schema() string {
	return "{\"fields\":[{\"name\":\"Calle\",\"type\":\"string\"},{\"name\":\"Numero\",\"type\":\"string\"},{\"name\":\"Piso\",\"type\":\"string\"},{\"name\":\"Unidad\",\"type\":\"string\"},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"CodigoPostal\",\"type\":\"string\"},{\"name\":\"Provincia\",\"type\":\"string\"},{\"name\":\"DatosAdicionales\",\"type\":\"string\"},{\"name\":\"Latitud\",\"type\":\"string\"},{\"name\":\"Longitud\",\"type\":\"string\"}],\"name\":\"Andreani.Corporativo.Events.Record.DireccionRemitenteEtiqueta\",\"type\":\"record\"}"
}

func (r DireccionRemitenteEtiqueta) SchemaName() string {
	return "Andreani.Corporativo.Events.Record.DireccionRemitenteEtiqueta"
}

func (_ DireccionRemitenteEtiqueta) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DireccionRemitenteEtiqueta) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DireccionRemitenteEtiqueta) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DireccionRemitenteEtiqueta) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DireccionRemitenteEtiqueta) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DireccionRemitenteEtiqueta) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DireccionRemitenteEtiqueta) SetString(v string)   { panic("Unsupported operation") }
func (_ DireccionRemitenteEtiqueta) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DireccionRemitenteEtiqueta) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Calle}

		return w

	case 1:
		w := types.String{Target: &r.Numero}

		return w

	case 2:
		w := types.String{Target: &r.Piso}

		return w

	case 3:
		w := types.String{Target: &r.Unidad}

		return w

	case 4:
		w := types.String{Target: &r.Localidad}

		return w

	case 5:
		w := types.String{Target: &r.CodigoPostal}

		return w

	case 6:
		w := types.String{Target: &r.Provincia}

		return w

	case 7:
		w := types.String{Target: &r.DatosAdicionales}

		return w

	case 8:
		w := types.String{Target: &r.Latitud}

		return w

	case 9:
		w := types.String{Target: &r.Longitud}

		return w

	}
	panic("Unknown field index")
}

func (r *DireccionRemitenteEtiqueta) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *DireccionRemitenteEtiqueta) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ DireccionRemitenteEtiqueta) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DireccionRemitenteEtiqueta) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DireccionRemitenteEtiqueta) HintSize(int)                     { panic("Unsupported operation") }
func (_ DireccionRemitenteEtiqueta) Finalize()                        {}

func (_ DireccionRemitenteEtiqueta) AvroCRC64Fingerprint() []byte {
	return []byte(DireccionRemitenteEtiquetaAvroCRC64Fingerprint)
}

func (r DireccionRemitenteEtiqueta) MarshalJSON() ([]byte, error) {
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
	output["Piso"], err = json.Marshal(r.Piso)
	if err != nil {
		return nil, err
	}
	output["Unidad"], err = json.Marshal(r.Unidad)
	if err != nil {
		return nil, err
	}
	output["Localidad"], err = json.Marshal(r.Localidad)
	if err != nil {
		return nil, err
	}
	output["CodigoPostal"], err = json.Marshal(r.CodigoPostal)
	if err != nil {
		return nil, err
	}
	output["Provincia"], err = json.Marshal(r.Provincia)
	if err != nil {
		return nil, err
	}
	output["DatosAdicionales"], err = json.Marshal(r.DatosAdicionales)
	if err != nil {
		return nil, err
	}
	output["Latitud"], err = json.Marshal(r.Latitud)
	if err != nil {
		return nil, err
	}
	output["Longitud"], err = json.Marshal(r.Longitud)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DireccionRemitenteEtiqueta) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["Piso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Piso); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Piso")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Unidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Unidad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Unidad")
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
		if v, ok := fields["DatosAdicionales"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DatosAdicionales); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DatosAdicionales")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Latitud"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Latitud); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Latitud")
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
		return fmt.Errorf("no value specified for Longitud")
	}
	return nil
}
