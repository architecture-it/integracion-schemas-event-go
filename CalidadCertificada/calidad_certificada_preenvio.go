// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CalidadCertificadaPreenvio.avsc
 */
package CalidadCertificadaEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type CalidadCertificadaPreenvio struct {
	IpImpresora string `json:"IpImpresora"`

	Detalle DetallePreenvio `json:"detalle"`
}

const CalidadCertificadaPreenvioAvroCRC64Fingerprint = "߷\xcb\xd1Q\x17vH"

func NewCalidadCertificadaPreenvio() CalidadCertificadaPreenvio {
	r := CalidadCertificadaPreenvio{}
	r.Detalle = NewDetallePreenvio()

	return r
}

func DeserializeCalidadCertificadaPreenvio(r io.Reader) (CalidadCertificadaPreenvio, error) {
	t := NewCalidadCertificadaPreenvio()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCalidadCertificadaPreenvioFromSchema(r io.Reader, schema string) (CalidadCertificadaPreenvio, error) {
	t := NewCalidadCertificadaPreenvio()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCalidadCertificadaPreenvio(r CalidadCertificadaPreenvio, w io.Writer) error {
	var err error
	err = vm.WriteString(r.IpImpresora, w)
	if err != nil {
		return err
	}
	err = writeDetallePreenvio(r.Detalle, w)
	if err != nil {
		return err
	}
	return err
}

func (r CalidadCertificadaPreenvio) Serialize(w io.Writer) error {
	return writeCalidadCertificadaPreenvio(r, w)
}

func (r CalidadCertificadaPreenvio) Schema() string {
	return "{\"fields\":[{\"name\":\"IpImpresora\",\"type\":\"string\"},{\"name\":\"detalle\",\"type\":{\"fields\":[{\"name\":\"Contrato\",\"type\":\"string\"},{\"name\":\"IdBulto\",\"type\":\"string\"},{\"name\":\"NumeroRemito\",\"type\":\"string\"},{\"name\":\"Interno\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"},{\"name\":\"FinPedido\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"Retornable\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"NroFactura\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CentroDeCosto\",\"type\":[\"null\",\"string\"]},{\"name\":\"ValorACobrar\",\"type\":\"float\"},{\"name\":\"Origen\",\"type\":{\"fields\":[{\"name\":\"Postal\",\"type\":{\"fields\":[{\"name\":\"CodigoPostal\",\"type\":\"string\"},{\"name\":\"Calle\",\"type\":\"string\"},{\"name\":\"Numero\",\"type\":\"string\"},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"Region\",\"type\":\"string\"},{\"name\":\"Pais\",\"type\":\"string\"},{\"name\":\"ComponentesDeDireccion\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Meta\",\"type\":\"string\"},{\"name\":\"Contenido\",\"type\":\"string\"}],\"name\":\"ComponenteDeDireccion\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Postal\",\"type\":\"record\"}}],\"name\":\"Direccion\",\"type\":\"record\"}},{\"name\":\"Destino\",\"type\":\"Andreani.CalidadCertificada.Events.Record.Direccion\"},{\"name\":\"Remitente\",\"type\":{\"fields\":[{\"name\":\"NombreCompleto\",\"type\":\"string\"},{\"name\":\"Email\",\"type\":\"string\"},{\"name\":\"DocumentoTipo\",\"type\":\"string\"},{\"name\":\"DocumentoNumero\",\"type\":\"string\"},{\"name\":\"Telefonos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Tipo\",\"type\":\"int\"},{\"name\":\"Numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Persona\",\"type\":\"record\"}},{\"name\":\"Destinatario\",\"type\":{\"items\":\"Andreani.CalidadCertificada.Events.Record.Persona\",\"type\":\"array\"}},{\"name\":\"Bultos\",\"type\":{\"items\":{\"fields\":[{\"default\":null,\"name\":\"EstadoBultoId\",\"type\":[\"null\",\"int\"]},{\"name\":\"Kilos\",\"type\":\"float\"},{\"name\":\"LargoCm\",\"type\":\"float\"},{\"name\":\"AltoCm\",\"type\":\"float\"},{\"name\":\"AnchoCm\",\"type\":\"float\"},{\"name\":\"VolumenCm\",\"type\":\"float\"},{\"name\":\"ValorDeclaradoConImpuestos\",\"type\":\"float\"},{\"name\":\"ValorDeclaradoSinImpuestos\",\"type\":\"float\"},{\"default\":null,\"name\":\"MontoValorSeguro\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"MontoGestionCobranza\",\"type\":[\"null\",\"float\"]},{\"name\":\"Referencias\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Meta\",\"type\":\"string\"},{\"name\":\"Contenido\",\"type\":\"string\"}],\"name\":\"Referencia\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Bulto\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"DetallePreenvio\",\"type\":\"record\"}}],\"name\":\"Andreani.CalidadCertificada.Events.Record.CalidadCertificadaPreenvio\",\"type\":\"record\"}"
}

func (r CalidadCertificadaPreenvio) SchemaName() string {
	return "Andreani.CalidadCertificada.Events.Record.CalidadCertificadaPreenvio"
}

func (_ CalidadCertificadaPreenvio) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CalidadCertificadaPreenvio) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CalidadCertificadaPreenvio) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CalidadCertificadaPreenvio) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CalidadCertificadaPreenvio) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CalidadCertificadaPreenvio) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CalidadCertificadaPreenvio) SetString(v string)   { panic("Unsupported operation") }
func (_ CalidadCertificadaPreenvio) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CalidadCertificadaPreenvio) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.IpImpresora}

		return w

	case 1:
		r.Detalle = NewDetallePreenvio()

		w := types.Record{Target: &r.Detalle}

		return w

	}
	panic("Unknown field index")
}

func (r *CalidadCertificadaPreenvio) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *CalidadCertificadaPreenvio) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ CalidadCertificadaPreenvio) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CalidadCertificadaPreenvio) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CalidadCertificadaPreenvio) HintSize(int)                     { panic("Unsupported operation") }
func (_ CalidadCertificadaPreenvio) Finalize()                        {}

func (_ CalidadCertificadaPreenvio) AvroCRC64Fingerprint() []byte {
	return []byte(CalidadCertificadaPreenvioAvroCRC64Fingerprint)
}

func (r CalidadCertificadaPreenvio) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["IpImpresora"], err = json.Marshal(r.IpImpresora)
	if err != nil {
		return nil, err
	}
	output["detalle"], err = json.Marshal(r.Detalle)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CalidadCertificadaPreenvio) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["IpImpresora"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IpImpresora); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IpImpresora")
	}
	val = func() json.RawMessage {
		if v, ok := fields["detalle"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Detalle); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for detalle")
	}
	return nil
}
