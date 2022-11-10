// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     MantenimientoDeLoteRequest.avsc
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

type MantenimientoDeLoteRequest struct {
	Destinatario *UnionNullString `json:"destinatario"`

	NumeroDeOrden string `json:"numeroDeOrden"`

	Vencimiento string `json:"vencimiento"`

	Contrato string `json:"contrato"`

	Almacen *UnionNullString `json:"almacen"`

	AlmacenSap *UnionNullString `json:"almacenSap"`

	Planta string `json:"planta"`

	UriConsulta *UnionNullString `json:"uriConsulta"`

	MantenimientoDeLote MantenimientoDeLoteSolicitado `json:"mantenimientoDeLote"`
}

const MantenimientoDeLoteRequestAvroCRC64Fingerprint = "zT\x8cN鬎\xdc"

func NewMantenimientoDeLoteRequest() MantenimientoDeLoteRequest {
	r := MantenimientoDeLoteRequest{}
	r.MantenimientoDeLote = NewMantenimientoDeLoteSolicitado()

	return r
}

func DeserializeMantenimientoDeLoteRequest(r io.Reader) (MantenimientoDeLoteRequest, error) {
	t := NewMantenimientoDeLoteRequest()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeMantenimientoDeLoteRequestFromSchema(r io.Reader, schema string) (MantenimientoDeLoteRequest, error) {
	t := NewMantenimientoDeLoteRequest()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeMantenimientoDeLoteRequest(r MantenimientoDeLoteRequest, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Destinatario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroDeOrden, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Vencimiento, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Contrato, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Almacen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.AlmacenSap, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Planta, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.UriConsulta, w)
	if err != nil {
		return err
	}
	err = writeMantenimientoDeLoteSolicitado(r.MantenimientoDeLote, w)
	if err != nil {
		return err
	}
	return err
}

func (r MantenimientoDeLoteRequest) Serialize(w io.Writer) error {
	return writeMantenimientoDeLoteRequest(r, w)
}

func (r MantenimientoDeLoteRequest) Schema() string {
	return "{\"fields\":[{\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"name\":\"numeroDeOrden\",\"type\":\"string\"},{\"name\":\"vencimiento\",\"type\":\"string\"},{\"name\":\"contrato\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":[\"null\",\"string\"]},{\"name\":\"almacenSap\",\"type\":[\"null\",\"string\"]},{\"name\":\"planta\",\"type\":\"string\"},{\"name\":\"uriConsulta\",\"type\":[\"null\",\"string\"]},{\"name\":\"mantenimientoDeLote\",\"type\":{\"fields\":[{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"articulo\",\"type\":\"string\"},{\"name\":\"paquete\",\"type\":[\"null\",\"string\"]},{\"name\":\"loteCaja\",\"type\":[\"null\",\"string\"]},{\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"name\":\"loteSap\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaFabricacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaVencimiento\",\"type\":[\"null\",\"string\"]},{\"name\":\"trazable\",\"type\":[\"null\",\"string\"]},{\"name\":\"estado\",\"type\":[\"null\",\"string\"]},{\"name\":\"procedencia\",\"type\":[\"null\",\"string\"]},{\"name\":\"campoLibre1\",\"type\":[\"null\",\"string\"]},{\"name\":\"campoLibre2\",\"type\":[\"null\",\"string\"]},{\"name\":\"campoLibre3\",\"type\":[\"null\",\"string\"]},{\"name\":\"campoLibre4\",\"type\":[\"null\",\"string\"]},{\"name\":\"campoLibre5\",\"type\":[\"null\",\"string\"]},{\"name\":\"loteExternoCliente\",\"type\":[\"null\",\"string\"]},{\"name\":\"deliverByDate\",\"type\":[\"null\",\"string\"]},{\"name\":\"bestByDate\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaCreacion\",\"type\":\"string\"},{\"name\":\"usuarioCreacion\",\"type\":\"string\"},{\"name\":\"fechaEdicion\",\"type\":\"string\"},{\"name\":\"usuarioEdicion\",\"type\":\"string\"}],\"name\":\"MantenimientoDeLoteSolicitado\",\"type\":\"record\"}}],\"name\":\"Andreani.ApiMantenimientoDeLote.Events.Record.MantenimientoDeLoteRequest\",\"type\":\"record\"}"
}

func (r MantenimientoDeLoteRequest) SchemaName() string {
	return "Andreani.ApiMantenimientoDeLote.Events.Record.MantenimientoDeLoteRequest"
}

func (_ MantenimientoDeLoteRequest) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRequest) SetInt(v int32)       { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRequest) SetLong(v int64)      { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRequest) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRequest) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRequest) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRequest) SetString(v string)   { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRequest) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MantenimientoDeLoteRequest) Get(i int) types.Field {
	switch i {
	case 0:
		r.Destinatario = NewUnionNullString()

		return r.Destinatario
	case 1:
		w := types.String{Target: &r.NumeroDeOrden}

		return w

	case 2:
		w := types.String{Target: &r.Vencimiento}

		return w

	case 3:
		w := types.String{Target: &r.Contrato}

		return w

	case 4:
		r.Almacen = NewUnionNullString()

		return r.Almacen
	case 5:
		r.AlmacenSap = NewUnionNullString()

		return r.AlmacenSap
	case 6:
		w := types.String{Target: &r.Planta}

		return w

	case 7:
		r.UriConsulta = NewUnionNullString()

		return r.UriConsulta
	case 8:
		r.MantenimientoDeLote = NewMantenimientoDeLoteSolicitado()

		w := types.Record{Target: &r.MantenimientoDeLote}

		return w

	}
	panic("Unknown field index")
}

func (r *MantenimientoDeLoteRequest) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *MantenimientoDeLoteRequest) NullField(i int) {
	switch i {
	case 0:
		r.Destinatario = nil
		return
	case 4:
		r.Almacen = nil
		return
	case 5:
		r.AlmacenSap = nil
		return
	case 7:
		r.UriConsulta = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ MantenimientoDeLoteRequest) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRequest) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRequest) HintSize(int)                     { panic("Unsupported operation") }
func (_ MantenimientoDeLoteRequest) Finalize()                        {}

func (_ MantenimientoDeLoteRequest) AvroCRC64Fingerprint() []byte {
	return []byte(MantenimientoDeLoteRequestAvroCRC64Fingerprint)
}

func (r MantenimientoDeLoteRequest) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["destinatario"], err = json.Marshal(r.Destinatario)
	if err != nil {
		return nil, err
	}
	output["numeroDeOrden"], err = json.Marshal(r.NumeroDeOrden)
	if err != nil {
		return nil, err
	}
	output["vencimiento"], err = json.Marshal(r.Vencimiento)
	if err != nil {
		return nil, err
	}
	output["contrato"], err = json.Marshal(r.Contrato)
	if err != nil {
		return nil, err
	}
	output["almacen"], err = json.Marshal(r.Almacen)
	if err != nil {
		return nil, err
	}
	output["almacenSap"], err = json.Marshal(r.AlmacenSap)
	if err != nil {
		return nil, err
	}
	output["planta"], err = json.Marshal(r.Planta)
	if err != nil {
		return nil, err
	}
	output["uriConsulta"], err = json.Marshal(r.UriConsulta)
	if err != nil {
		return nil, err
	}
	output["mantenimientoDeLote"], err = json.Marshal(r.MantenimientoDeLote)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *MantenimientoDeLoteRequest) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["destinatario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Destinatario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for destinatario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroDeOrden"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeOrden); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroDeOrden")
	}
	val = func() json.RawMessage {
		if v, ok := fields["vencimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Vencimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for vencimiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["contrato"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contrato); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for contrato")
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
		if v, ok := fields["almacenSap"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AlmacenSap); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for almacenSap")
	}
	val = func() json.RawMessage {
		if v, ok := fields["planta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Planta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for planta")
	}
	val = func() json.RawMessage {
		if v, ok := fields["uriConsulta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UriConsulta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for uriConsulta")
	}
	val = func() json.RawMessage {
		if v, ok := fields["mantenimientoDeLote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MantenimientoDeLote); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for mantenimientoDeLote")
	}
	return nil
}
