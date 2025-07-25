// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Ticket.avsc
 */
package EAMEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Ticket struct {
	IdTicket int32 `json:"IdTicket"`

	Solicitante string `json:"Solicitante"`

	Asunto string `json:"Asunto"`

	Descripcion string `json:"Descripcion"`

	Planta string `json:"Planta"`

	TipoSector *UnionNullString `json:"TipoSector"`

	Categoria_Servicio *UnionNullString `json:"Categoria_Servicio"`

	Telefono *UnionNullString `json:"Telefono"`

	Correo_Electronico *UnionNullString `json:"Correo_Electronico"`

	Mesa_Origen *UnionNullString `json:"Mesa_Origen"`
}

const TicketAvroCRC64Fingerprint = "X\xe2\xd2.\xb0\x82\xb7}"

func NewTicket() Ticket {
	r := Ticket{}
	return r
}

func DeserializeTicket(r io.Reader) (Ticket, error) {
	t := NewTicket()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeTicketFromSchema(r io.Reader, schema string) (Ticket, error) {
	t := NewTicket()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeTicket(r Ticket, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.IdTicket, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Solicitante, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Asunto, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Descripcion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Planta, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoSector, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Categoria_Servicio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Telefono, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Correo_Electronico, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Mesa_Origen, w)
	if err != nil {
		return err
	}
	return err
}

func (r Ticket) Serialize(w io.Writer) error {
	return writeTicket(r, w)
}

func (r Ticket) Schema() string {
	return "{\"fields\":[{\"name\":\"IdTicket\",\"type\":\"int\"},{\"name\":\"Solicitante\",\"type\":\"string\"},{\"name\":\"Asunto\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"name\":\"Planta\",\"type\":\"string\"},{\"name\":\"TipoSector\",\"type\":[\"null\",\"string\"]},{\"name\":\"Categoria_Servicio\",\"type\":[\"null\",\"string\"]},{\"name\":\"Telefono\",\"type\":[\"null\",\"string\"]},{\"name\":\"Correo_Electronico\",\"type\":[\"null\",\"string\"]},{\"name\":\"Mesa_Origen\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.EAM.Events.MDA.Ticket\",\"type\":\"record\"}"
}

func (r Ticket) SchemaName() string {
	return "Andreani.EAM.Events.MDA.Ticket"
}

func (_ Ticket) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Ticket) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Ticket) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Ticket) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Ticket) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Ticket) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Ticket) SetString(v string)   { panic("Unsupported operation") }
func (_ Ticket) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Ticket) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.IdTicket}

		return w

	case 1:
		w := types.String{Target: &r.Solicitante}

		return w

	case 2:
		w := types.String{Target: &r.Asunto}

		return w

	case 3:
		w := types.String{Target: &r.Descripcion}

		return w

	case 4:
		w := types.String{Target: &r.Planta}

		return w

	case 5:
		r.TipoSector = NewUnionNullString()

		return r.TipoSector
	case 6:
		r.Categoria_Servicio = NewUnionNullString()

		return r.Categoria_Servicio
	case 7:
		r.Telefono = NewUnionNullString()

		return r.Telefono
	case 8:
		r.Correo_Electronico = NewUnionNullString()

		return r.Correo_Electronico
	case 9:
		r.Mesa_Origen = NewUnionNullString()

		return r.Mesa_Origen
	}
	panic("Unknown field index")
}

func (r *Ticket) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Ticket) NullField(i int) {
	switch i {
	case 5:
		r.TipoSector = nil
		return
	case 6:
		r.Categoria_Servicio = nil
		return
	case 7:
		r.Telefono = nil
		return
	case 8:
		r.Correo_Electronico = nil
		return
	case 9:
		r.Mesa_Origen = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Ticket) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Ticket) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Ticket) HintSize(int)                     { panic("Unsupported operation") }
func (_ Ticket) Finalize()                        {}

func (_ Ticket) AvroCRC64Fingerprint() []byte {
	return []byte(TicketAvroCRC64Fingerprint)
}

func (r Ticket) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["IdTicket"], err = json.Marshal(r.IdTicket)
	if err != nil {
		return nil, err
	}
	output["Solicitante"], err = json.Marshal(r.Solicitante)
	if err != nil {
		return nil, err
	}
	output["Asunto"], err = json.Marshal(r.Asunto)
	if err != nil {
		return nil, err
	}
	output["Descripcion"], err = json.Marshal(r.Descripcion)
	if err != nil {
		return nil, err
	}
	output["Planta"], err = json.Marshal(r.Planta)
	if err != nil {
		return nil, err
	}
	output["TipoSector"], err = json.Marshal(r.TipoSector)
	if err != nil {
		return nil, err
	}
	output["Categoria_Servicio"], err = json.Marshal(r.Categoria_Servicio)
	if err != nil {
		return nil, err
	}
	output["Telefono"], err = json.Marshal(r.Telefono)
	if err != nil {
		return nil, err
	}
	output["Correo_Electronico"], err = json.Marshal(r.Correo_Electronico)
	if err != nil {
		return nil, err
	}
	output["Mesa_Origen"], err = json.Marshal(r.Mesa_Origen)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Ticket) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["IdTicket"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdTicket); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IdTicket")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Solicitante"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Solicitante); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Solicitante")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Asunto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Asunto); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Asunto")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Descripcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Descripcion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Descripcion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Planta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Planta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Planta")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoSector"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoSector); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoSector")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Categoria_Servicio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Categoria_Servicio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Categoria_Servicio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Telefono"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Telefono); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Telefono")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Correo_Electronico"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Correo_Electronico); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Correo_Electronico")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Mesa_Origen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Mesa_Origen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Mesa_Origen")
	}
	return nil
}
