// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NovedadGenerada.avsc
 */
package NovedadesCustomEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type NovedadGenerada struct {
	IdContrato int32 `json:"idContrato"`

	IdCliente int32 `json:"idCliente"`

	NombreDeCliente string `json:"nombreDeCliente"`

	NumeroDeCliente *UnionNullString `json:"numeroDeCliente"`

	NumeroDeContratoInterno string `json:"numeroDeContratoInterno"`

	NumeroDeEnvio string `json:"numeroDeEnvio"`

	Evento string `json:"evento"`

	Motivo *UnionNullString `json:"motivo"`

	Submotivo *UnionNullString `json:"Submotivo"`

	Remitente *UnionNullString `json:"remitente"`

	CicloDelEnvio *UnionNullString `json:"cicloDelEnvio"`

	FechaDeIncidencia int64 `json:"fechaDeIncidencia"`

	CodigoCliente1 *UnionNullString `json:"codigoCliente1"`

	CodigoCliente2 *UnionNullString `json:"codigoCliente2"`

	MensajeAEnviar string `json:"mensajeAEnviar"`

	UrlDestino *UnionNullString `json:"urlDestino"`
}

const NovedadGeneradaAvroCRC64Fingerprint = "\xd6DaB\x1e\x86<\xac"

func NewNovedadGenerada() NovedadGenerada {
	r := NovedadGenerada{}
	r.NumeroDeCliente = nil
	r.Motivo = nil
	r.Submotivo = nil
	r.Remitente = nil
	r.CicloDelEnvio = nil
	r.UrlDestino = nil
	return r
}

func DeserializeNovedadGenerada(r io.Reader) (NovedadGenerada, error) {
	t := NewNovedadGenerada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNovedadGeneradaFromSchema(r io.Reader, schema string) (NovedadGenerada, error) {
	t := NewNovedadGenerada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNovedadGenerada(r NovedadGenerada, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.IdContrato, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.IdCliente, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NombreDeCliente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeCliente, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroDeContratoInterno, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroDeEnvio, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Evento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Motivo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Submotivo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Remitente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CicloDelEnvio, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.FechaDeIncidencia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoCliente1, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoCliente2, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.MensajeAEnviar, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.UrlDestino, w)
	if err != nil {
		return err
	}
	return err
}

func (r NovedadGenerada) Serialize(w io.Writer) error {
	return writeNovedadGenerada(r, w)
}

func (r NovedadGenerada) Schema() string {
	return "{\"fields\":[{\"name\":\"idContrato\",\"type\":\"int\"},{\"name\":\"idCliente\",\"type\":\"int\"},{\"name\":\"nombreDeCliente\",\"type\":\"string\"},{\"default\":null,\"name\":\"numeroDeCliente\",\"type\":[\"null\",\"string\"]},{\"name\":\"numeroDeContratoInterno\",\"type\":\"string\"},{\"name\":\"numeroDeEnvio\",\"type\":\"string\"},{\"name\":\"evento\",\"type\":\"string\"},{\"default\":null,\"name\":\"motivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Submotivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remitente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaDeIncidencia\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoCliente1\",\"type\":[\"null\",\"string\"]},{\"name\":\"codigoCliente2\",\"type\":[\"null\",\"string\"]},{\"name\":\"mensajeAEnviar\",\"type\":\"string\"},{\"default\":null,\"name\":\"urlDestino\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.NovedadesCustom.Events.Record.NovedadGenerada\",\"type\":\"record\"}"
}

func (r NovedadGenerada) SchemaName() string {
	return "Andreani.NovedadesCustom.Events.Record.NovedadGenerada"
}

func (_ NovedadGenerada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NovedadGenerada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NovedadGenerada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NovedadGenerada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NovedadGenerada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NovedadGenerada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NovedadGenerada) SetString(v string)   { panic("Unsupported operation") }
func (_ NovedadGenerada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NovedadGenerada) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.IdContrato}

		return w

	case 1:
		w := types.Int{Target: &r.IdCliente}

		return w

	case 2:
		w := types.String{Target: &r.NombreDeCliente}

		return w

	case 3:
		r.NumeroDeCliente = NewUnionNullString()

		return r.NumeroDeCliente
	case 4:
		w := types.String{Target: &r.NumeroDeContratoInterno}

		return w

	case 5:
		w := types.String{Target: &r.NumeroDeEnvio}

		return w

	case 6:
		w := types.String{Target: &r.Evento}

		return w

	case 7:
		r.Motivo = NewUnionNullString()

		return r.Motivo
	case 8:
		r.Submotivo = NewUnionNullString()

		return r.Submotivo
	case 9:
		r.Remitente = NewUnionNullString()

		return r.Remitente
	case 10:
		r.CicloDelEnvio = NewUnionNullString()

		return r.CicloDelEnvio
	case 11:
		w := types.Long{Target: &r.FechaDeIncidencia}

		return w

	case 12:
		r.CodigoCliente1 = NewUnionNullString()

		return r.CodigoCliente1
	case 13:
		r.CodigoCliente2 = NewUnionNullString()

		return r.CodigoCliente2
	case 14:
		w := types.String{Target: &r.MensajeAEnviar}

		return w

	case 15:
		r.UrlDestino = NewUnionNullString()

		return r.UrlDestino
	}
	panic("Unknown field index")
}

func (r *NovedadGenerada) SetDefault(i int) {
	switch i {
	case 3:
		r.NumeroDeCliente = nil
		return
	case 7:
		r.Motivo = nil
		return
	case 8:
		r.Submotivo = nil
		return
	case 9:
		r.Remitente = nil
		return
	case 10:
		r.CicloDelEnvio = nil
		return
	case 15:
		r.UrlDestino = nil
		return
	}
	panic("Unknown field index")
}

func (r *NovedadGenerada) NullField(i int) {
	switch i {
	case 3:
		r.NumeroDeCliente = nil
		return
	case 7:
		r.Motivo = nil
		return
	case 8:
		r.Submotivo = nil
		return
	case 9:
		r.Remitente = nil
		return
	case 10:
		r.CicloDelEnvio = nil
		return
	case 12:
		r.CodigoCliente1 = nil
		return
	case 13:
		r.CodigoCliente2 = nil
		return
	case 15:
		r.UrlDestino = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ NovedadGenerada) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NovedadGenerada) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NovedadGenerada) HintSize(int)                     { panic("Unsupported operation") }
func (_ NovedadGenerada) Finalize()                        {}

func (_ NovedadGenerada) AvroCRC64Fingerprint() []byte {
	return []byte(NovedadGeneradaAvroCRC64Fingerprint)
}

func (r NovedadGenerada) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["idContrato"], err = json.Marshal(r.IdContrato)
	if err != nil {
		return nil, err
	}
	output["idCliente"], err = json.Marshal(r.IdCliente)
	if err != nil {
		return nil, err
	}
	output["nombreDeCliente"], err = json.Marshal(r.NombreDeCliente)
	if err != nil {
		return nil, err
	}
	output["numeroDeCliente"], err = json.Marshal(r.NumeroDeCliente)
	if err != nil {
		return nil, err
	}
	output["numeroDeContratoInterno"], err = json.Marshal(r.NumeroDeContratoInterno)
	if err != nil {
		return nil, err
	}
	output["numeroDeEnvio"], err = json.Marshal(r.NumeroDeEnvio)
	if err != nil {
		return nil, err
	}
	output["evento"], err = json.Marshal(r.Evento)
	if err != nil {
		return nil, err
	}
	output["motivo"], err = json.Marshal(r.Motivo)
	if err != nil {
		return nil, err
	}
	output["Submotivo"], err = json.Marshal(r.Submotivo)
	if err != nil {
		return nil, err
	}
	output["remitente"], err = json.Marshal(r.Remitente)
	if err != nil {
		return nil, err
	}
	output["cicloDelEnvio"], err = json.Marshal(r.CicloDelEnvio)
	if err != nil {
		return nil, err
	}
	output["fechaDeIncidencia"], err = json.Marshal(r.FechaDeIncidencia)
	if err != nil {
		return nil, err
	}
	output["codigoCliente1"], err = json.Marshal(r.CodigoCliente1)
	if err != nil {
		return nil, err
	}
	output["codigoCliente2"], err = json.Marshal(r.CodigoCliente2)
	if err != nil {
		return nil, err
	}
	output["mensajeAEnviar"], err = json.Marshal(r.MensajeAEnviar)
	if err != nil {
		return nil, err
	}
	output["urlDestino"], err = json.Marshal(r.UrlDestino)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NovedadGenerada) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["idContrato"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdContrato); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idContrato")
	}
	val = func() json.RawMessage {
		if v, ok := fields["idCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdCliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idCliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["nombreDeCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NombreDeCliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for nombreDeCliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroDeCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeCliente); err != nil {
			return err
		}
	} else {
		r.NumeroDeCliente = NewUnionNullString()

		r.NumeroDeCliente = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroDeContratoInterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeContratoInterno); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroDeContratoInterno")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroDeEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeEnvio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroDeEnvio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["evento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Evento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for evento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["motivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Motivo); err != nil {
			return err
		}
	} else {
		r.Motivo = NewUnionNullString()

		r.Motivo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Submotivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Submotivo); err != nil {
			return err
		}
	} else {
		r.Submotivo = NewUnionNullString()

		r.Submotivo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["remitente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Remitente); err != nil {
			return err
		}
	} else {
		r.Remitente = NewUnionNullString()

		r.Remitente = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["cicloDelEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CicloDelEnvio); err != nil {
			return err
		}
	} else {
		r.CicloDelEnvio = NewUnionNullString()

		r.CicloDelEnvio = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaDeIncidencia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaDeIncidencia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fechaDeIncidencia")
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoCliente1"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoCliente1); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigoCliente1")
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoCliente2"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoCliente2); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigoCliente2")
	}
	val = func() json.RawMessage {
		if v, ok := fields["mensajeAEnviar"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MensajeAEnviar); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for mensajeAEnviar")
	}
	val = func() json.RawMessage {
		if v, ok := fields["urlDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UrlDestino); err != nil {
			return err
		}
	} else {
		r.UrlDestino = NewUnionNullString()

		r.UrlDestino = nil
	}
	return nil
}
