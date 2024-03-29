// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoDeAlmacenSCE.avsc
 */
package ApiPedidosAlmacenesEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Destinatario struct {
	DatosPersonales DatosPersonales `json:"datosPersonales"`

	OtrosDatos *UnionNullListaDePropiedades `json:"OtrosDatos"`

	Contacto *UnionNullString `json:"Contacto"`
}

const DestinatarioAvroCRC64Fingerprint = "\x8d3\xce\x01\x03\xd92u"

func NewDestinatario() Destinatario {
	r := Destinatario{}
	r.DatosPersonales = NewDatosPersonales()

	return r
}

func DeserializeDestinatario(r io.Reader) (Destinatario, error) {
	t := NewDestinatario()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDestinatarioFromSchema(r io.Reader, schema string) (Destinatario, error) {
	t := NewDestinatario()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDestinatario(r Destinatario, w io.Writer) error {
	var err error
	err = writeDatosPersonales(r.DatosPersonales, w)
	if err != nil {
		return err
	}
	err = writeUnionNullListaDePropiedades(r.OtrosDatos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Contacto, w)
	if err != nil {
		return err
	}
	return err
}

func (r Destinatario) Serialize(w io.Writer) error {
	return writeDestinatario(r, w)
}

func (r Destinatario) Schema() string {
	return "{\"fields\":[{\"name\":\"datosPersonales\",\"type\":{\"fields\":[{\"name\":\"NumeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"name\":\"NombreCompleto\",\"type\":[\"null\",\"string\"]},{\"name\":\"IdInternoDelCliente\",\"type\":[\"null\",\"string\"]},{\"name\":\"EMail\",\"type\":[\"null\",\"string\"]},{\"name\":\"Telefonos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"Tipo\",\"type\":[\"null\",\"string\"]},{\"name\":\"Numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDeTelefonos\",\"type\":\"record\"}]},{\"name\":\"Agrupador\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoDeDocumento\",\"type\":[\"null\",\"string\"]}],\"name\":\"DatosPersonales\",\"type\":\"record\"}},{\"name\":\"OtrosDatos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Meta\",\"type\":\"string\"},{\"name\":\"Contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"name\":\"Contacto\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.WapAltaDePedidoSolicitada.Events.Record.Destinatario\",\"type\":\"record\"}"
}

func (r Destinatario) SchemaName() string {
	return "Andreani.WapAltaDePedidoSolicitada.Events.Record.Destinatario"
}

func (_ Destinatario) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Destinatario) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Destinatario) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Destinatario) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Destinatario) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Destinatario) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Destinatario) SetString(v string)   { panic("Unsupported operation") }
func (_ Destinatario) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Destinatario) Get(i int) types.Field {
	switch i {
	case 0:
		r.DatosPersonales = NewDatosPersonales()

		w := types.Record{Target: &r.DatosPersonales}

		return w

	case 1:
		r.OtrosDatos = NewUnionNullListaDePropiedades()

		return r.OtrosDatos
	case 2:
		r.Contacto = NewUnionNullString()

		return r.Contacto
	}
	panic("Unknown field index")
}

func (r *Destinatario) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Destinatario) NullField(i int) {
	switch i {
	case 1:
		r.OtrosDatos = nil
		return
	case 2:
		r.Contacto = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Destinatario) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Destinatario) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Destinatario) HintSize(int)                     { panic("Unsupported operation") }
func (_ Destinatario) Finalize()                        {}

func (_ Destinatario) AvroCRC64Fingerprint() []byte {
	return []byte(DestinatarioAvroCRC64Fingerprint)
}

func (r Destinatario) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["datosPersonales"], err = json.Marshal(r.DatosPersonales)
	if err != nil {
		return nil, err
	}
	output["OtrosDatos"], err = json.Marshal(r.OtrosDatos)
	if err != nil {
		return nil, err
	}
	output["Contacto"], err = json.Marshal(r.Contacto)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Destinatario) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["datosPersonales"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DatosPersonales); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for datosPersonales")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OtrosDatos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OtrosDatos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OtrosDatos")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Contacto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contacto); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Contacto")
	}
	return nil
}
