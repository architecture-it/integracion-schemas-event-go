// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     SelfTransportShipmentEvent.avsc
 */
package MEunoApiEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type MeliOrderInfo struct {
	SellerId int64 `json:"sellerId"`

	ShipmentId int64 `json:"shipmentId"`

	MLAId string `json:"MLAId"`

	PublicationName string `json:"publicationName"`

	OrderID int64 `json:"orderID"`

	EstimatedDeliveryTime string `json:"estimatedDeliveryTime"`

	Created string `json:"created"`

	PackId *UnionNullLong `json:"packId"`

	ListCost *UnionNullString `json:"listCost"`
}

const MeliOrderInfoAvroCRC64Fingerprint = "\xba9\x1d\x1c\x1cx~\xdd"

func NewMeliOrderInfo() MeliOrderInfo {
	r := MeliOrderInfo{}
	return r
}

func DeserializeMeliOrderInfo(r io.Reader) (MeliOrderInfo, error) {
	t := NewMeliOrderInfo()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeMeliOrderInfoFromSchema(r io.Reader, schema string) (MeliOrderInfo, error) {
	t := NewMeliOrderInfo()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeMeliOrderInfo(r MeliOrderInfo, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.SellerId, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.ShipmentId, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.MLAId, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.PublicationName, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.OrderID, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EstimatedDeliveryTime, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Created, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.PackId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ListCost, w)
	if err != nil {
		return err
	}
	return err
}

func (r MeliOrderInfo) Serialize(w io.Writer) error {
	return writeMeliOrderInfo(r, w)
}

func (r MeliOrderInfo) Schema() string {
	return "{\"fields\":[{\"name\":\"sellerId\",\"type\":\"long\"},{\"name\":\"shipmentId\",\"type\":\"long\"},{\"name\":\"MLAId\",\"type\":\"string\"},{\"name\":\"publicationName\",\"type\":\"string\"},{\"name\":\"orderID\",\"type\":\"long\"},{\"name\":\"estimatedDeliveryTime\",\"type\":\"string\"},{\"name\":\"created\",\"type\":\"string\"},{\"name\":\"packId\",\"type\":[\"null\",\"long\"]},{\"name\":\"listCost\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.MEunoApi.Events.Record.Structs.MeliOrderInfo\",\"type\":\"record\"}"
}

func (r MeliOrderInfo) SchemaName() string {
	return "Andreani.MEunoApi.Events.Record.Structs.MeliOrderInfo"
}

func (_ MeliOrderInfo) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ MeliOrderInfo) SetInt(v int32)       { panic("Unsupported operation") }
func (_ MeliOrderInfo) SetLong(v int64)      { panic("Unsupported operation") }
func (_ MeliOrderInfo) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ MeliOrderInfo) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ MeliOrderInfo) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ MeliOrderInfo) SetString(v string)   { panic("Unsupported operation") }
func (_ MeliOrderInfo) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MeliOrderInfo) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.SellerId}

		return w

	case 1:
		w := types.Long{Target: &r.ShipmentId}

		return w

	case 2:
		w := types.String{Target: &r.MLAId}

		return w

	case 3:
		w := types.String{Target: &r.PublicationName}

		return w

	case 4:
		w := types.Long{Target: &r.OrderID}

		return w

	case 5:
		w := types.String{Target: &r.EstimatedDeliveryTime}

		return w

	case 6:
		w := types.String{Target: &r.Created}

		return w

	case 7:
		r.PackId = NewUnionNullLong()

		return r.PackId
	case 8:
		r.ListCost = NewUnionNullString()

		return r.ListCost
	}
	panic("Unknown field index")
}

func (r *MeliOrderInfo) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *MeliOrderInfo) NullField(i int) {
	switch i {
	case 7:
		r.PackId = nil
		return
	case 8:
		r.ListCost = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ MeliOrderInfo) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ MeliOrderInfo) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ MeliOrderInfo) HintSize(int)                     { panic("Unsupported operation") }
func (_ MeliOrderInfo) Finalize()                        {}

func (_ MeliOrderInfo) AvroCRC64Fingerprint() []byte {
	return []byte(MeliOrderInfoAvroCRC64Fingerprint)
}

func (r MeliOrderInfo) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["sellerId"], err = json.Marshal(r.SellerId)
	if err != nil {
		return nil, err
	}
	output["shipmentId"], err = json.Marshal(r.ShipmentId)
	if err != nil {
		return nil, err
	}
	output["MLAId"], err = json.Marshal(r.MLAId)
	if err != nil {
		return nil, err
	}
	output["publicationName"], err = json.Marshal(r.PublicationName)
	if err != nil {
		return nil, err
	}
	output["orderID"], err = json.Marshal(r.OrderID)
	if err != nil {
		return nil, err
	}
	output["estimatedDeliveryTime"], err = json.Marshal(r.EstimatedDeliveryTime)
	if err != nil {
		return nil, err
	}
	output["created"], err = json.Marshal(r.Created)
	if err != nil {
		return nil, err
	}
	output["packId"], err = json.Marshal(r.PackId)
	if err != nil {
		return nil, err
	}
	output["listCost"], err = json.Marshal(r.ListCost)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *MeliOrderInfo) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["sellerId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SellerId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for sellerId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["shipmentId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ShipmentId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for shipmentId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["MLAId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MLAId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for MLAId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["publicationName"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PublicationName); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for publicationName")
	}
	val = func() json.RawMessage {
		if v, ok := fields["orderID"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrderID); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for orderID")
	}
	val = func() json.RawMessage {
		if v, ok := fields["estimatedDeliveryTime"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstimatedDeliveryTime); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for estimatedDeliveryTime")
	}
	val = func() json.RawMessage {
		if v, ok := fields["created"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Created); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for created")
	}
	val = func() json.RawMessage {
		if v, ok := fields["packId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PackId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for packId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["listCost"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ListCost); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for listCost")
	}
	return nil
}
