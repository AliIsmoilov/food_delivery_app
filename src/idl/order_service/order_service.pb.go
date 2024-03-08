// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.5
// source: src/idl/order_service/order_service.proto

package order_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_idl_order_service_order_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_src_idl_order_service_order_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_src_idl_order_service_order_service_proto_rawDescGZIP(), []int{0}
}

type PrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PrimaryKey) Reset() {
	*x = PrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_idl_order_service_order_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimaryKey) ProtoMessage() {}

func (x *PrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_src_idl_order_service_order_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimaryKey.ProtoReflect.Descriptor instead.
func (*PrimaryKey) Descriptor() ([]byte, []int) {
	return file_src_idl_order_service_order_service_proto_rawDescGZIP(), []int{1}
}

func (x *PrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code         string  `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Note         string  `protobuf:"bytes,3,opt,name=note,proto3" json:"note,omitempty"`
	TotalPrice   uint32  `protobuf:"varint,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Status       string  `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	RestaurantId string  `protobuf:"bytes,6,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	CustomerId   string  `protobuf:"bytes,7,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CourierId    string  `protobuf:"bytes,8,opt,name=courier_id,json=courierId,proto3" json:"courier_id,omitempty"`
	Adress       string  `protobuf:"bytes,9,opt,name=adress,proto3" json:"adress,omitempty"`
	Lat          float32 `protobuf:"fixed32,10,opt,name=lat,proto3" json:"lat,omitempty"`
	Long         float32 `protobuf:"fixed32,11,opt,name=long,proto3" json:"long,omitempty"`
	Items        []*Item `protobuf:"bytes,12,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_idl_order_service_order_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_src_idl_order_service_order_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_src_idl_order_service_order_service_proto_rawDescGZIP(), []int{2}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Order) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Order) GetTotalPrice() uint32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *Order) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Order) GetCourierId() string {
	if x != nil {
		return x.CourierId
	}
	return ""
}

func (x *Order) GetAdress() string {
	if x != nil {
		return x.Adress
	}
	return ""
}

func (x *Order) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Order) GetLong() float32 {
	if x != nil {
		return x.Long
	}
	return 0
}

func (x *Order) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemId  string `protobuf:"bytes,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	OrderId string `protobuf:"bytes,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Qty     uint32 `protobuf:"varint,4,opt,name=qty,proto3" json:"qty,omitempty"`
	Price   uint32 `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_idl_order_service_order_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_src_idl_order_service_order_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_src_idl_order_service_order_service_proto_rawDescGZIP(), []int{3}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *Item) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Item) GetQty() uint32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

func (x *Item) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type OrderListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *OrderListReq) Reset() {
	*x = OrderListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_idl_order_service_order_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderListReq) ProtoMessage() {}

func (x *OrderListReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_idl_order_service_order_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderListReq.ProtoReflect.Descriptor instead.
func (*OrderListReq) Descriptor() ([]byte, []int) {
	return file_src_idl_order_service_order_service_proto_rawDescGZIP(), []int{4}
}

func (x *OrderListReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type OrderListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  uint32   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Orders []*Order `protobuf:"bytes,2,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *OrderListResp) Reset() {
	*x = OrderListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_idl_order_service_order_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderListResp) ProtoMessage() {}

func (x *OrderListResp) ProtoReflect() protoreflect.Message {
	mi := &file_src_idl_order_service_order_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderListResp.ProtoReflect.Descriptor instead.
func (*OrderListResp) Descriptor() ([]byte, []int) {
	return file_src_idl_order_service_order_service_proto_rawDescGZIP(), []int{5}
}

func (x *OrderListResp) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *OrderListResp) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type Courier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber string  `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Photo       string  `protobuf:"bytes,4,opt,name=photo,proto3" json:"photo,omitempty"`
	CarModel    string  `protobuf:"bytes,5,opt,name=car_model,json=carModel,proto3" json:"car_model,omitempty"`
	CarNumber   string  `protobuf:"bytes,6,opt,name=car_number,json=carNumber,proto3" json:"car_number,omitempty"`
	CarColor    string  `protobuf:"bytes,7,opt,name=car_color,json=carColor,proto3" json:"car_color,omitempty"`
	IsAvailable bool    `protobuf:"varint,8,opt,name=is_available,json=isAvailable,proto3" json:"is_available,omitempty"`
	Lat         float32 `protobuf:"fixed32,9,opt,name=lat,proto3" json:"lat,omitempty"`
	Long        float32 `protobuf:"fixed32,10,opt,name=long,proto3" json:"long,omitempty"`
}

func (x *Courier) Reset() {
	*x = Courier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_idl_order_service_order_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Courier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Courier) ProtoMessage() {}

func (x *Courier) ProtoReflect() protoreflect.Message {
	mi := &file_src_idl_order_service_order_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Courier.ProtoReflect.Descriptor instead.
func (*Courier) Descriptor() ([]byte, []int) {
	return file_src_idl_order_service_order_service_proto_rawDescGZIP(), []int{6}
}

func (x *Courier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Courier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Courier) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Courier) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *Courier) GetCarModel() string {
	if x != nil {
		return x.CarModel
	}
	return ""
}

func (x *Courier) GetCarNumber() string {
	if x != nil {
		return x.CarNumber
	}
	return ""
}

func (x *Courier) GetCarColor() string {
	if x != nil {
		return x.CarColor
	}
	return ""
}

func (x *Courier) GetIsAvailable() bool {
	if x != nil {
		return x.IsAvailable
	}
	return false
}

func (x *Courier) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Courier) GetLong() float32 {
	if x != nil {
		return x.Long
	}
	return 0
}

type CourierListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page   uint32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *CourierListReq) Reset() {
	*x = CourierListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_idl_order_service_order_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourierListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourierListReq) ProtoMessage() {}

func (x *CourierListReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_idl_order_service_order_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourierListReq.ProtoReflect.Descriptor instead.
func (*CourierListReq) Descriptor() ([]byte, []int) {
	return file_src_idl_order_service_order_service_proto_rawDescGZIP(), []int{7}
}

func (x *CourierListReq) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *CourierListReq) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *CourierListReq) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type CourierListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count    int32      `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Couriers []*Courier `protobuf:"bytes,2,rep,name=couriers,proto3" json:"couriers,omitempty"`
}

func (x *CourierListResp) Reset() {
	*x = CourierListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_idl_order_service_order_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourierListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourierListResp) ProtoMessage() {}

func (x *CourierListResp) ProtoReflect() protoreflect.Message {
	mi := &file_src_idl_order_service_order_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourierListResp.ProtoReflect.Descriptor instead.
func (*CourierListResp) Descriptor() ([]byte, []int) {
	return file_src_idl_order_service_order_service_proto_rawDescGZIP(), []int{8}
}

func (x *CourierListResp) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CourierListResp) GetCouriers() []*Courier {
	if x != nil {
		return x.Couriers
	}
	return nil
}

var File_src_idl_order_service_order_service_proto protoreflect.FileDescriptor

var file_src_idl_order_service_order_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x72, 0x63, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x1c, 0x0a, 0x0a, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0xc6, 0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x6f, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x6f, 0x6e, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12,
	0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x72, 0x0a, 0x04, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x71, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x26,
	0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x53, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a,
	0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x88, 0x02, 0x0a, 0x07,
	0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x72, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x69, 0x73, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c,
	0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x22, 0x52, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x5b, 0x0a, 0x0f, 0x43, 0x6f,
	0x75, 0x72, 0x69, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x08, 0x63,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x73, 0x32, 0xe0, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x19, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x09, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x75, 0x72, 0x69, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x1a, 0x19, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_src_idl_order_service_order_service_proto_rawDescOnce sync.Once
	file_src_idl_order_service_order_service_proto_rawDescData = file_src_idl_order_service_order_service_proto_rawDesc
)

func file_src_idl_order_service_order_service_proto_rawDescGZIP() []byte {
	file_src_idl_order_service_order_service_proto_rawDescOnce.Do(func() {
		file_src_idl_order_service_order_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_idl_order_service_order_service_proto_rawDescData)
	})
	return file_src_idl_order_service_order_service_proto_rawDescData
}

var file_src_idl_order_service_order_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_src_idl_order_service_order_service_proto_goTypes = []interface{}{
	(*Empty)(nil),           // 0: order_service.Empty
	(*PrimaryKey)(nil),      // 1: order_service.PrimaryKey
	(*Order)(nil),           // 2: order_service.Order
	(*Item)(nil),            // 3: order_service.Item
	(*OrderListReq)(nil),    // 4: order_service.OrderListReq
	(*OrderListResp)(nil),   // 5: order_service.OrderListResp
	(*Courier)(nil),         // 6: order_service.Courier
	(*CourierListReq)(nil),  // 7: order_service.CourierListReq
	(*CourierListResp)(nil), // 8: order_service.CourierListResp
}
var file_src_idl_order_service_order_service_proto_depIdxs = []int32{
	3, // 0: order_service.Order.items:type_name -> order_service.Item
	2, // 1: order_service.OrderListResp.orders:type_name -> order_service.Order
	6, // 2: order_service.CourierListResp.couriers:type_name -> order_service.Courier
	2, // 3: order_service.OrderService.CreateOrder:input_type -> order_service.Order
	4, // 4: order_service.OrderService.OrderList:input_type -> order_service.OrderListReq
	6, // 5: order_service.OrderService.CreateCourier:input_type -> order_service.Courier
	1, // 6: order_service.OrderService.CreateOrder:output_type -> order_service.PrimaryKey
	5, // 7: order_service.OrderService.OrderList:output_type -> order_service.OrderListResp
	1, // 8: order_service.OrderService.CreateCourier:output_type -> order_service.PrimaryKey
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_src_idl_order_service_order_service_proto_init() }
func file_src_idl_order_service_order_service_proto_init() {
	if File_src_idl_order_service_order_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_idl_order_service_order_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_idl_order_service_order_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimaryKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_idl_order_service_order_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_idl_order_service_order_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_idl_order_service_order_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_idl_order_service_order_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderListResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_idl_order_service_order_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Courier); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_idl_order_service_order_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourierListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_idl_order_service_order_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourierListResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_idl_order_service_order_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_idl_order_service_order_service_proto_goTypes,
		DependencyIndexes: file_src_idl_order_service_order_service_proto_depIdxs,
		MessageInfos:      file_src_idl_order_service_order_service_proto_msgTypes,
	}.Build()
	File_src_idl_order_service_order_service_proto = out.File
	file_src_idl_order_service_order_service_proto_rawDesc = nil
	file_src_idl_order_service_order_service_proto_goTypes = nil
	file_src_idl_order_service_order_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*PrimaryKey, error)
	OrderList(ctx context.Context, in *OrderListReq, opts ...grpc.CallOption) (*OrderListResp, error)
	CreateCourier(ctx context.Context, in *Courier, opts ...grpc.CallOption) (*PrimaryKey, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*PrimaryKey, error) {
	out := new(PrimaryKey)
	err := c.cc.Invoke(ctx, "/order_service.OrderService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) OrderList(ctx context.Context, in *OrderListReq, opts ...grpc.CallOption) (*OrderListResp, error) {
	out := new(OrderListResp)
	err := c.cc.Invoke(ctx, "/order_service.OrderService/OrderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateCourier(ctx context.Context, in *Courier, opts ...grpc.CallOption) (*PrimaryKey, error) {
	out := new(PrimaryKey)
	err := c.cc.Invoke(ctx, "/order_service.OrderService/CreateCourier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	CreateOrder(context.Context, *Order) (*PrimaryKey, error)
	OrderList(context.Context, *OrderListReq) (*OrderListResp, error)
	CreateCourier(context.Context, *Courier) (*PrimaryKey, error)
}

// UnimplementedOrderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (*UnimplementedOrderServiceServer) CreateOrder(context.Context, *Order) (*PrimaryKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (*UnimplementedOrderServiceServer) OrderList(context.Context, *OrderListReq) (*OrderListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderList not implemented")
}
func (*UnimplementedOrderServiceServer) CreateCourier(context.Context, *Courier) (*PrimaryKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourier not implemented")
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_service.OrderService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_OrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).OrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_service.OrderService/OrderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).OrderList(ctx, req.(*OrderListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreateCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Courier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_service.OrderService/CreateCourier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateCourier(ctx, req.(*Courier))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "order_service.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "OrderList",
			Handler:    _OrderService_OrderList_Handler,
		},
		{
			MethodName: "CreateCourier",
			Handler:    _OrderService_CreateCourier_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/idl/order_service/order_service.proto",
}
