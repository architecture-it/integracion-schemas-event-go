// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     OrdenDeRetiroRechazada.avsc
 */
package SppeApiEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type OrdenDeRetiroRechazada struct {
	Identificacion Identificacion `json:"Identificacion"`

	Cabecera Cabecera `json:"Cabecera"`

	Destino Destino `json:"Destino"`

	Origen Origen `json:"Origen"`

	Contenedores []Contenedor `json:"Contenedores"`
}

const OrdenDeRetiroRechazadaAvroCRC64Fingerprint = "\xf6Wge\x83\xb2Ȍ"

func NewOrdenDeRetiroRechazada() OrdenDeRetiroRechazada {
	r := OrdenDeRetiroRechazada{}
	r.Identificacion = NewIdentificacion()

	r.Cabecera = NewCabecera()

	r.Destino = NewDestino()

	r.Origen = NewOrigen()

	r.Contenedores = make([]Contenedor, 0)

	return r
}

func DeserializeOrdenDeRetiroRechazada(r io.Reader) (OrdenDeRetiroRechazada, error) {
	t := NewOrdenDeRetiroRechazada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeOrdenDeRetiroRechazadaFromSchema(r io.Reader, schema string) (OrdenDeRetiroRechazada, error) {
	t := NewOrdenDeRetiroRechazada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeOrdenDeRetiroRechazada(r OrdenDeRetiroRechazada, w io.Writer) error {
	var err error
	err = writeIdentificacion(r.Identificacion, w)
	if err != nil {
		return err
	}
	err = writeCabecera(r.Cabecera, w)
	if err != nil {
		return err
	}
	err = writeDestino(r.Destino, w)
	if err != nil {
		return err
	}
	err = writeOrigen(r.Origen, w)
	if err != nil {
		return err
	}
	err = writeArrayContenedor(r.Contenedores, w)
	if err != nil {
		return err
	}
	return err
}

func (r OrdenDeRetiroRechazada) Serialize(w io.Writer) error {
	return writeOrdenDeRetiroRechazada(r, w)
}

func (r OrdenDeRetiroRechazada) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.SppeApi.Events.OrdenDeRetiroRechazadaCommon\",\"type\":\"record\"}},{\"name\":\"Cabecera\",\"type\":{\"fields\":[{\"name\":\"Remito\",\"type\":\"string\"},{\"name\":\"OrdenWh\",\"type\":\"string\"},{\"name\":\"OrdenCliente\",\"type\":\"string\"},{\"name\":\"ContratoTMS\",\"type\":\"string\"},{\"name\":\"ValorSeguro\",\"type\":[\"null\",\"string\"]},{\"name\":\"MensajeError\",\"type\":\"string\"}],\"name\":\"Cabecera\",\"namespace\":\"Andreani.SppeApi.Events.OrdenDeRetiroRechazadaCommon\",\"type\":\"record\"}},{\"name\":\"Destino\",\"type\":{\"fields\":[{\"name\":\"DestinatarioCalle\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinarioNumero\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioPiso\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioDepartamento\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioGLNDNI\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioCiudad\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioProvincia\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioCodigoPostal\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioTelefono\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioEmail\",\"type\":[\"null\",\"string\"]}],\"name\":\"Destino\",\"namespace\":\"Andreani.SppeApi.Events.OrdenDeRetiroRechazadaCommon\",\"type\":\"record\"}},{\"name\":\"Origen\",\"type\":{\"fields\":[{\"name\":\"OrigenCiudad\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenCodigoPostal\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenCalle\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenNumero\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenPiso\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenDepartamento\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenEmail\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenTelefono\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenRegion\",\"type\":[\"null\",\"string\"]}],\"name\":\"Origen\",\"namespace\":\"Andreani.SppeApi.Events.OrdenDeRetiroRechazadaCommon\",\"type\":\"record\"}},{\"name\":\"Contenedores\",\"type\":{\"items\":{\"fields\":[{\"name\":\"DropId\",\"type\":\"string\"},{\"name\":\"Altura\",\"type\":\"float\"},{\"name\":\"Profundidad\",\"type\":\"float\"},{\"name\":\"Ancho\",\"type\":\"float\"},{\"name\":\"Peso\",\"type\":\"float\"}],\"name\":\"Contenedor\",\"namespace\":\"Andreani.SppeApi.Events.OrdenDeRetiroRechazadaCommon\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.SppeApi.Events.Record.OrdenDeRetiroRechazada\",\"type\":\"record\"}"
}

func (r OrdenDeRetiroRechazada) SchemaName() string {
	return "Andreani.SppeApi.Events.Record.OrdenDeRetiroRechazada"
}

func (_ OrdenDeRetiroRechazada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ OrdenDeRetiroRechazada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ OrdenDeRetiroRechazada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ OrdenDeRetiroRechazada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ OrdenDeRetiroRechazada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ OrdenDeRetiroRechazada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ OrdenDeRetiroRechazada) SetString(v string)   { panic("Unsupported operation") }
func (_ OrdenDeRetiroRechazada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *OrdenDeRetiroRechazada) Get(i int) types.Field {
	switch i {
	case 0:
		r.Identificacion = NewIdentificacion()

		w := types.Record{Target: &r.Identificacion}

		return w

	case 1:
		r.Cabecera = NewCabecera()

		w := types.Record{Target: &r.Cabecera}

		return w

	case 2:
		r.Destino = NewDestino()

		w := types.Record{Target: &r.Destino}

		return w

	case 3:
		r.Origen = NewOrigen()

		w := types.Record{Target: &r.Origen}

		return w

	case 4:
		r.Contenedores = make([]Contenedor, 0)

		w := ArrayContenedorWrapper{Target: &r.Contenedores}

		return w

	}
	panic("Unknown field index")
}

func (r *OrdenDeRetiroRechazada) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *OrdenDeRetiroRechazada) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ OrdenDeRetiroRechazada) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ OrdenDeRetiroRechazada) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ OrdenDeRetiroRechazada) HintSize(int)                     { panic("Unsupported operation") }
func (_ OrdenDeRetiroRechazada) Finalize()                        {}

func (_ OrdenDeRetiroRechazada) AvroCRC64Fingerprint() []byte {
	return []byte(OrdenDeRetiroRechazadaAvroCRC64Fingerprint)
}

func (r OrdenDeRetiroRechazada) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Identificacion"], err = json.Marshal(r.Identificacion)
	if err != nil {
		return nil, err
	}
	output["Cabecera"], err = json.Marshal(r.Cabecera)
	if err != nil {
		return nil, err
	}
	output["Destino"], err = json.Marshal(r.Destino)
	if err != nil {
		return nil, err
	}
	output["Origen"], err = json.Marshal(r.Origen)
	if err != nil {
		return nil, err
	}
	output["Contenedores"], err = json.Marshal(r.Contenedores)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *OrdenDeRetiroRechazada) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Identificacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Identificacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Identificacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Cabecera"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cabecera); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Cabecera")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Destino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Destino); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Destino")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Origen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Origen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Origen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Contenedores"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contenedores); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Contenedores")
	}
	return nil
}
