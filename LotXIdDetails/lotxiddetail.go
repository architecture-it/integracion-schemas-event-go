// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Lotxiddetail.avsc
 */
package LotXIdDetailsEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Lotxiddetail struct {
	Captureby *UnionNullString `json:"captureby"`

	Id *UnionNullString `json:"id"`

	Ioflag *UnionNullString `json:"ioflag"`

	Iqty *UnionNullString `json:"iqty"`

	Itrnkey *UnionNullString `json:"itrnkey"`

	Lot *UnionNullString `json:"lot"`

	Oother1 *UnionNullString `json:"oother1"`

	Oqty *UnionNullString `json:"oqty"`

	Pickdetailkey *UnionNullString `json:"pickdetailkey"`

	Serialnumberlong *UnionNullString `json:"serialnumberlong"`

	Sku *UnionNullString `json:"sku"`

	Sourcekey *UnionNullString `json:"sourcekey"`

	Sourcelinenumber *UnionNullString `json:"sourcelinenumber"`

	Almacen *UnionNullString `json:"almacen"`
}

const LotxiddetailAvroCRC64Fingerprint = "\xcb\xde\xec\xabE\xda\x01U"

func NewLotxiddetail() Lotxiddetail {
	r := Lotxiddetail{}
	r.Captureby = nil
	r.Id = nil
	r.Ioflag = nil
	r.Iqty = nil
	r.Itrnkey = nil
	r.Lot = nil
	r.Oother1 = nil
	r.Oqty = nil
	r.Pickdetailkey = nil
	r.Serialnumberlong = nil
	r.Sku = nil
	r.Sourcekey = nil
	r.Sourcelinenumber = nil
	r.Almacen = nil
	return r
}

func DeserializeLotxiddetail(r io.Reader) (Lotxiddetail, error) {
	t := NewLotxiddetail()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeLotxiddetailFromSchema(r io.Reader, schema string) (Lotxiddetail, error) {
	t := NewLotxiddetail()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeLotxiddetail(r Lotxiddetail, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Captureby, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Id, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Ioflag, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Iqty, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Itrnkey, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Lot, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Oother1, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Oqty, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Pickdetailkey, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Serialnumberlong, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Sku, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Sourcekey, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Sourcelinenumber, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Almacen, w)
	if err != nil {
		return err
	}
	return err
}

func (r Lotxiddetail) Serialize(w io.Writer) error {
	return writeLotxiddetail(r, w)
}

func (r Lotxiddetail) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"captureby\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ioflag\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"iqty\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"itrnkey\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"lot\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"oother1\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"oqty\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pickdetailkey\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"serialnumberlong\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sku\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sourcekey\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sourcelinenumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"almacen\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.LotXIdDetails.Events.Record.Lotxiddetail\",\"type\":\"record\"}"
}

func (r Lotxiddetail) SchemaName() string {
	return "Andreani.LotXIdDetails.Events.Record.Lotxiddetail"
}

func (_ Lotxiddetail) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Lotxiddetail) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Lotxiddetail) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Lotxiddetail) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Lotxiddetail) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Lotxiddetail) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Lotxiddetail) SetString(v string)   { panic("Unsupported operation") }
func (_ Lotxiddetail) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Lotxiddetail) Get(i int) types.Field {
	switch i {
	case 0:
		r.Captureby = NewUnionNullString()

		return r.Captureby
	case 1:
		r.Id = NewUnionNullString()

		return r.Id
	case 2:
		r.Ioflag = NewUnionNullString()

		return r.Ioflag
	case 3:
		r.Iqty = NewUnionNullString()

		return r.Iqty
	case 4:
		r.Itrnkey = NewUnionNullString()

		return r.Itrnkey
	case 5:
		r.Lot = NewUnionNullString()

		return r.Lot
	case 6:
		r.Oother1 = NewUnionNullString()

		return r.Oother1
	case 7:
		r.Oqty = NewUnionNullString()

		return r.Oqty
	case 8:
		r.Pickdetailkey = NewUnionNullString()

		return r.Pickdetailkey
	case 9:
		r.Serialnumberlong = NewUnionNullString()

		return r.Serialnumberlong
	case 10:
		r.Sku = NewUnionNullString()

		return r.Sku
	case 11:
		r.Sourcekey = NewUnionNullString()

		return r.Sourcekey
	case 12:
		r.Sourcelinenumber = NewUnionNullString()

		return r.Sourcelinenumber
	case 13:
		r.Almacen = NewUnionNullString()

		return r.Almacen
	}
	panic("Unknown field index")
}

func (r *Lotxiddetail) SetDefault(i int) {
	switch i {
	case 0:
		r.Captureby = nil
		return
	case 1:
		r.Id = nil
		return
	case 2:
		r.Ioflag = nil
		return
	case 3:
		r.Iqty = nil
		return
	case 4:
		r.Itrnkey = nil
		return
	case 5:
		r.Lot = nil
		return
	case 6:
		r.Oother1 = nil
		return
	case 7:
		r.Oqty = nil
		return
	case 8:
		r.Pickdetailkey = nil
		return
	case 9:
		r.Serialnumberlong = nil
		return
	case 10:
		r.Sku = nil
		return
	case 11:
		r.Sourcekey = nil
		return
	case 12:
		r.Sourcelinenumber = nil
		return
	case 13:
		r.Almacen = nil
		return
	}
	panic("Unknown field index")
}

func (r *Lotxiddetail) NullField(i int) {
	switch i {
	case 0:
		r.Captureby = nil
		return
	case 1:
		r.Id = nil
		return
	case 2:
		r.Ioflag = nil
		return
	case 3:
		r.Iqty = nil
		return
	case 4:
		r.Itrnkey = nil
		return
	case 5:
		r.Lot = nil
		return
	case 6:
		r.Oother1 = nil
		return
	case 7:
		r.Oqty = nil
		return
	case 8:
		r.Pickdetailkey = nil
		return
	case 9:
		r.Serialnumberlong = nil
		return
	case 10:
		r.Sku = nil
		return
	case 11:
		r.Sourcekey = nil
		return
	case 12:
		r.Sourcelinenumber = nil
		return
	case 13:
		r.Almacen = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Lotxiddetail) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Lotxiddetail) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Lotxiddetail) HintSize(int)                     { panic("Unsupported operation") }
func (_ Lotxiddetail) Finalize()                        {}

func (_ Lotxiddetail) AvroCRC64Fingerprint() []byte {
	return []byte(LotxiddetailAvroCRC64Fingerprint)
}

func (r Lotxiddetail) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["captureby"], err = json.Marshal(r.Captureby)
	if err != nil {
		return nil, err
	}
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["ioflag"], err = json.Marshal(r.Ioflag)
	if err != nil {
		return nil, err
	}
	output["iqty"], err = json.Marshal(r.Iqty)
	if err != nil {
		return nil, err
	}
	output["itrnkey"], err = json.Marshal(r.Itrnkey)
	if err != nil {
		return nil, err
	}
	output["lot"], err = json.Marshal(r.Lot)
	if err != nil {
		return nil, err
	}
	output["oother1"], err = json.Marshal(r.Oother1)
	if err != nil {
		return nil, err
	}
	output["oqty"], err = json.Marshal(r.Oqty)
	if err != nil {
		return nil, err
	}
	output["pickdetailkey"], err = json.Marshal(r.Pickdetailkey)
	if err != nil {
		return nil, err
	}
	output["serialnumberlong"], err = json.Marshal(r.Serialnumberlong)
	if err != nil {
		return nil, err
	}
	output["sku"], err = json.Marshal(r.Sku)
	if err != nil {
		return nil, err
	}
	output["sourcekey"], err = json.Marshal(r.Sourcekey)
	if err != nil {
		return nil, err
	}
	output["sourcelinenumber"], err = json.Marshal(r.Sourcelinenumber)
	if err != nil {
		return nil, err
	}
	output["almacen"], err = json.Marshal(r.Almacen)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Lotxiddetail) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["captureby"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Captureby); err != nil {
			return err
		}
	} else {
		r.Captureby = NewUnionNullString()

		r.Captureby = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		r.Id = NewUnionNullString()

		r.Id = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ioflag"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ioflag); err != nil {
			return err
		}
	} else {
		r.Ioflag = NewUnionNullString()

		r.Ioflag = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["iqty"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Iqty); err != nil {
			return err
		}
	} else {
		r.Iqty = NewUnionNullString()

		r.Iqty = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["itrnkey"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Itrnkey); err != nil {
			return err
		}
	} else {
		r.Itrnkey = NewUnionNullString()

		r.Itrnkey = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["lot"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Lot); err != nil {
			return err
		}
	} else {
		r.Lot = NewUnionNullString()

		r.Lot = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["oother1"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Oother1); err != nil {
			return err
		}
	} else {
		r.Oother1 = NewUnionNullString()

		r.Oother1 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["oqty"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Oqty); err != nil {
			return err
		}
	} else {
		r.Oqty = NewUnionNullString()

		r.Oqty = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["pickdetailkey"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Pickdetailkey); err != nil {
			return err
		}
	} else {
		r.Pickdetailkey = NewUnionNullString()

		r.Pickdetailkey = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["serialnumberlong"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Serialnumberlong); err != nil {
			return err
		}
	} else {
		r.Serialnumberlong = NewUnionNullString()

		r.Serialnumberlong = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["sku"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sku); err != nil {
			return err
		}
	} else {
		r.Sku = NewUnionNullString()

		r.Sku = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["sourcekey"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sourcekey); err != nil {
			return err
		}
	} else {
		r.Sourcekey = NewUnionNullString()

		r.Sourcekey = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["sourcelinenumber"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sourcelinenumber); err != nil {
			return err
		}
	} else {
		r.Sourcelinenumber = NewUnionNullString()

		r.Sourcelinenumber = nil
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
		r.Almacen = NewUnionNullString()

		r.Almacen = nil
	}
	return nil
}