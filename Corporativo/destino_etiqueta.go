// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EmailEtiquetaEvent.avsc
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

type DestinoEtiqueta struct {
	Sucursal SucursalEtiqueta `json:"Sucursal"`

	PuntoDeTercero PuntoDeTerceroEtiqueta `json:"PuntoDeTercero"`

	DireccionDestino DireccionDestinoEtiqueta `json:"DireccionDestino"`
}

const DestinoEtiquetaAvroCRC64Fingerprint = ".5\xb3\xfc\\u\xbeL"

func NewDestinoEtiqueta() DestinoEtiqueta {
	r := DestinoEtiqueta{}
	r.Sucursal = NewSucursalEtiqueta()

	r.PuntoDeTercero = NewPuntoDeTerceroEtiqueta()

	r.DireccionDestino = NewDireccionDestinoEtiqueta()

	return r
}

func DeserializeDestinoEtiqueta(r io.Reader) (DestinoEtiqueta, error) {
	t := NewDestinoEtiqueta()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDestinoEtiquetaFromSchema(r io.Reader, schema string) (DestinoEtiqueta, error) {
	t := NewDestinoEtiqueta()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDestinoEtiqueta(r DestinoEtiqueta, w io.Writer) error {
	var err error
	err = writeSucursalEtiqueta(r.Sucursal, w)
	if err != nil {
		return err
	}
	err = writePuntoDeTerceroEtiqueta(r.PuntoDeTercero, w)
	if err != nil {
		return err
	}
	err = writeDireccionDestinoEtiqueta(r.DireccionDestino, w)
	if err != nil {
		return err
	}
	return err
}

func (r DestinoEtiqueta) Serialize(w io.Writer) error {
	return writeDestinoEtiqueta(r, w)
}

func (r DestinoEtiqueta) Schema() string {
	return "{\"fields\":[{\"name\":\"Sucursal\",\"type\":{\"fields\":[{\"name\":\"Codigo\",\"type\":\"string\"},{\"name\":\"Numero\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"name\":\"Canal\",\"type\":\"string\"},{\"name\":\"Direccion\",\"type\":{\"fields\":[{\"name\":\"Calle\",\"type\":\"string\"},{\"name\":\"Numero\",\"type\":\"string\"},{\"name\":\"Provincia\",\"type\":\"string\"},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"Region\",\"type\":\"string\"},{\"name\":\"Pais\",\"type\":\"string\"},{\"name\":\"CodigoPostal\",\"type\":\"string\"}],\"name\":\"DireccionSucursalEtiqueta\",\"type\":\"record\"}},{\"name\":\"SeHaceAtencionAlCliente\",\"type\":\"boolean\"},{\"name\":\"Tipo\",\"type\":\"string\"},{\"name\":\"JsonTelefonos\",\"type\":\"string\"},{\"name\":\"Latitud\",\"type\":\"double\"},{\"name\":\"Longitud\",\"type\":\"double\"}],\"name\":\"SucursalEtiqueta\",\"type\":\"record\"}},{\"name\":\"PuntoDeTercero\",\"type\":{\"fields\":[{\"name\":\"Codigo\",\"type\":\"string\"},{\"name\":\"Numero\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"name\":\"Canal\",\"type\":\"string\"},{\"name\":\"Direccion\",\"type\":{\"fields\":[{\"name\":\"Calle\",\"type\":\"string\"},{\"name\":\"Numero\",\"type\":\"string\"},{\"name\":\"Provincia\",\"type\":\"string\"},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"Region\",\"type\":\"string\"},{\"name\":\"Pais\",\"type\":\"string\"},{\"name\":\"CodigoPostal\",\"type\":\"string\"}],\"name\":\"DireccionPuntoDeTerceroEtiqueta\",\"type\":\"record\"}},{\"name\":\"SeHaceAtencionAlCliente\",\"type\":\"boolean\"},{\"name\":\"Tipo\",\"type\":\"string\"},{\"name\":\"JsonTelefonos\",\"type\":\"string\"},{\"name\":\"Latitud\",\"type\":\"double\"},{\"name\":\"Longitud\",\"type\":\"double\"}],\"name\":\"PuntoDeTerceroEtiqueta\",\"type\":\"record\"}},{\"name\":\"DireccionDestino\",\"type\":{\"fields\":[{\"name\":\"Calle\",\"type\":\"string\"},{\"name\":\"Numero\",\"type\":\"string\"},{\"name\":\"Piso\",\"type\":\"string\"},{\"name\":\"Unidad\",\"type\":\"string\"},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"CodigoPostal\",\"type\":\"string\"},{\"name\":\"Provincia\",\"type\":\"string\"},{\"name\":\"DatosAdicionales\",\"type\":\"string\"},{\"name\":\"Latitud\",\"type\":\"string\"},{\"name\":\"Longitud\",\"type\":\"string\"},{\"name\":\"ObservacionesAdicionales\",\"type\":\"string\"}],\"name\":\"DireccionDestinoEtiqueta\",\"type\":\"record\"}}],\"name\":\"Andreani.Corporativo.Events.Record.DestinoEtiqueta\",\"type\":\"record\"}"
}

func (r DestinoEtiqueta) SchemaName() string {
	return "Andreani.Corporativo.Events.Record.DestinoEtiqueta"
}

func (_ DestinoEtiqueta) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DestinoEtiqueta) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DestinoEtiqueta) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DestinoEtiqueta) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DestinoEtiqueta) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DestinoEtiqueta) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DestinoEtiqueta) SetString(v string)   { panic("Unsupported operation") }
func (_ DestinoEtiqueta) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DestinoEtiqueta) Get(i int) types.Field {
	switch i {
	case 0:
		r.Sucursal = NewSucursalEtiqueta()

		w := types.Record{Target: &r.Sucursal}

		return w

	case 1:
		r.PuntoDeTercero = NewPuntoDeTerceroEtiqueta()

		w := types.Record{Target: &r.PuntoDeTercero}

		return w

	case 2:
		r.DireccionDestino = NewDireccionDestinoEtiqueta()

		w := types.Record{Target: &r.DireccionDestino}

		return w

	}
	panic("Unknown field index")
}

func (r *DestinoEtiqueta) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *DestinoEtiqueta) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ DestinoEtiqueta) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DestinoEtiqueta) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DestinoEtiqueta) HintSize(int)                     { panic("Unsupported operation") }
func (_ DestinoEtiqueta) Finalize()                        {}

func (_ DestinoEtiqueta) AvroCRC64Fingerprint() []byte {
	return []byte(DestinoEtiquetaAvroCRC64Fingerprint)
}

func (r DestinoEtiqueta) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Sucursal"], err = json.Marshal(r.Sucursal)
	if err != nil {
		return nil, err
	}
	output["PuntoDeTercero"], err = json.Marshal(r.PuntoDeTercero)
	if err != nil {
		return nil, err
	}
	output["DireccionDestino"], err = json.Marshal(r.DireccionDestino)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DestinoEtiqueta) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Sucursal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sucursal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Sucursal")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PuntoDeTercero"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PuntoDeTercero); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PuntoDeTercero")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DireccionDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DireccionDestino); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DireccionDestino")
	}
	return nil
}
