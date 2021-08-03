// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: contract.proto

package contract

import (
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

// BackoffPolicyType is the type for backoff policies
type BackoffPolicy int32

const (
	// Exponential backoff policy
	BackoffPolicy_Exponential BackoffPolicy = 0
	// Linear backoff policy
	BackoffPolicy_Linear BackoffPolicy = 1
)

// Enum value maps for BackoffPolicy.
var (
	BackoffPolicy_name = map[int32]string{
		0: "Exponential",
		1: "Linear",
	}
	BackoffPolicy_value = map[string]int32{
		"Exponential": 0,
		"Linear":      1,
	}
)

func (x BackoffPolicy) Enum() *BackoffPolicy {
	p := new(BackoffPolicy)
	*p = x
	return p
}

func (x BackoffPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BackoffPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_contract_proto_enumTypes[0].Descriptor()
}

func (BackoffPolicy) Type() protoreflect.EnumType {
	return &file_contract_proto_enumTypes[0]
}

func (x BackoffPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BackoffPolicy.Descriptor instead.
func (BackoffPolicy) EnumDescriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{0}
}

// Check dev.knative.eventing.kafka.broker.dispatcher.consumer.DeliveryOrder for more details
type DeliveryOrder int32

const (
	DeliveryOrder_UNORDERED DeliveryOrder = 0
	DeliveryOrder_ORDERED   DeliveryOrder = 1
)

// Enum value maps for DeliveryOrder.
var (
	DeliveryOrder_name = map[int32]string{
		0: "UNORDERED",
		1: "ORDERED",
	}
	DeliveryOrder_value = map[string]int32{
		"UNORDERED": 0,
		"ORDERED":   1,
	}
)

func (x DeliveryOrder) Enum() *DeliveryOrder {
	p := new(DeliveryOrder)
	*p = x
	return p
}

func (x DeliveryOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeliveryOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_contract_proto_enumTypes[1].Descriptor()
}

func (DeliveryOrder) Type() protoreflect.EnumType {
	return &file_contract_proto_enumTypes[1]
}

func (x DeliveryOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeliveryOrder.Descriptor instead.
func (DeliveryOrder) EnumDescriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{1}
}

// CloudEvent content mode
type ContentMode int32

const (
	ContentMode_BINARY     ContentMode = 0
	ContentMode_STRUCTURED ContentMode = 1
)

// Enum value maps for ContentMode.
var (
	ContentMode_name = map[int32]string{
		0: "BINARY",
		1: "STRUCTURED",
	}
	ContentMode_value = map[string]int32{
		"BINARY":     0,
		"STRUCTURED": 1,
	}
)

func (x ContentMode) Enum() *ContentMode {
	p := new(ContentMode)
	*p = x
	return p
}

func (x ContentMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContentMode) Descriptor() protoreflect.EnumDescriptor {
	return file_contract_proto_enumTypes[2].Descriptor()
}

func (ContentMode) Type() protoreflect.EnumType {
	return &file_contract_proto_enumTypes[2]
}

func (x ContentMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContentMode.Descriptor instead.
func (ContentMode) EnumDescriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{2}
}

// We don't use the google.protobuf.Empty type because
// configuring the include directory is a mess for the contributors and for the build scripts.
// Hence, more than dealing with contributors that can't get their dev environment
// working with the project, we prefer to have this additional single line of code.
// Protobuf include nightmare? No thanks!
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[0]
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
	return file_contract_proto_rawDescGZIP(), []int{0}
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// attributes filters events by exact match on event context attributes.
	// Each key in the map is compared with the equivalent key in the event
	// context. An event passes the filter if all values are equal to the
	// specified values.
	//
	// Nested context attributes are not supported as keys. Only string values are supported.
	Attributes map[string]string `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{1}
}

func (x *Filter) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type EgressConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Dead letter is where the event is sent when something goes wrong
	DeadLetter string `protobuf:"bytes,1,opt,name=deadLetter,proto3" json:"deadLetter,omitempty"`
	// retry is the minimum number of retries the sender should attempt when
	// sending an event before moving it to the dead letter sink.
	//
	// Setting retry to 0 means don't retry.
	Retry uint32 `protobuf:"varint,2,opt,name=retry,proto3" json:"retry,omitempty"`
	// backoffPolicy is the retry backoff policy (linear, exponential).
	BackoffPolicy BackoffPolicy `protobuf:"varint,3,opt,name=backoffPolicy,proto3,enum=BackoffPolicy" json:"backoffPolicy,omitempty"`
	// backoffDelay is the delay before retrying in milliseconds.
	BackoffDelay uint64 `protobuf:"varint,4,opt,name=backoffDelay,proto3" json:"backoffDelay,omitempty"`
	// timeout is the single request timeout (not the overall retry timeout)
	Timeout uint64 `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *EgressConfig) Reset() {
	*x = EgressConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EgressConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EgressConfig) ProtoMessage() {}

func (x *EgressConfig) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EgressConfig.ProtoReflect.Descriptor instead.
func (*EgressConfig) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{2}
}

func (x *EgressConfig) GetDeadLetter() string {
	if x != nil {
		return x.DeadLetter
	}
	return ""
}

func (x *EgressConfig) GetRetry() uint32 {
	if x != nil {
		return x.Retry
	}
	return 0
}

func (x *EgressConfig) GetBackoffPolicy() BackoffPolicy {
	if x != nil {
		return x.BackoffPolicy
	}
	return BackoffPolicy_Exponential
}

func (x *EgressConfig) GetBackoffDelay() uint64 {
	if x != nil {
		return x.BackoffDelay
	}
	return 0
}

func (x *EgressConfig) GetTimeout() uint64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type Egress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// consumer group name
	ConsumerGroup string `protobuf:"bytes,1,opt,name=consumerGroup,proto3" json:"consumerGroup,omitempty"`
	// destination is the sink where events are sent.
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	// Types that are assignable to ReplyStrategy:
	//	*Egress_ReplyUrl
	//	*Egress_ReplyToOriginalTopic
	ReplyStrategy isEgress_ReplyStrategy `protobuf_oneof:"replyStrategy"`
	Filter        *Filter                `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
	// Id of the egress
	// It's the same as the Kubernetes resource uid
	Uid string `protobuf:"bytes,6,opt,name=uid,proto3" json:"uid,omitempty"`
	// Egress configuration.
	// It overrides Resource's EgressConfig.
	EgressConfig *EgressConfig `protobuf:"bytes,7,opt,name=egressConfig,proto3" json:"egressConfig,omitempty"`
	// Delivery guarantee to use
	// Empty defaults to unordered
	DeliveryOrder DeliveryOrder `protobuf:"varint,8,opt,name=deliveryOrder,proto3,enum=DeliveryOrder" json:"deliveryOrder,omitempty"`
}

func (x *Egress) Reset() {
	*x = Egress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Egress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Egress) ProtoMessage() {}

func (x *Egress) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Egress.ProtoReflect.Descriptor instead.
func (*Egress) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{3}
}

func (x *Egress) GetConsumerGroup() string {
	if x != nil {
		return x.ConsumerGroup
	}
	return ""
}

func (x *Egress) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (m *Egress) GetReplyStrategy() isEgress_ReplyStrategy {
	if m != nil {
		return m.ReplyStrategy
	}
	return nil
}

func (x *Egress) GetReplyUrl() string {
	if x, ok := x.GetReplyStrategy().(*Egress_ReplyUrl); ok {
		return x.ReplyUrl
	}
	return ""
}

func (x *Egress) GetReplyToOriginalTopic() *Empty {
	if x, ok := x.GetReplyStrategy().(*Egress_ReplyToOriginalTopic); ok {
		return x.ReplyToOriginalTopic
	}
	return nil
}

func (x *Egress) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *Egress) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Egress) GetEgressConfig() *EgressConfig {
	if x != nil {
		return x.EgressConfig
	}
	return nil
}

func (x *Egress) GetDeliveryOrder() DeliveryOrder {
	if x != nil {
		return x.DeliveryOrder
	}
	return DeliveryOrder_UNORDERED
}

type isEgress_ReplyStrategy interface {
	isEgress_ReplyStrategy()
}

type Egress_ReplyUrl struct {
	// Send the response to an url
	ReplyUrl string `protobuf:"bytes,3,opt,name=replyUrl,proto3,oneof"`
}

type Egress_ReplyToOriginalTopic struct {
	// Send the response to a Kafka topic
	ReplyToOriginalTopic *Empty `protobuf:"bytes,4,opt,name=replyToOriginalTopic,proto3,oneof"`
}

func (*Egress_ReplyUrl) isEgress_ReplyStrategy() {}

func (*Egress_ReplyToOriginalTopic) isEgress_ReplyStrategy() {}

type Ingress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional content mode to use when pushing messages to Kafka
	ContentMode ContentMode `protobuf:"varint,1,opt,name=contentMode,proto3,enum=ContentMode" json:"contentMode,omitempty"`
	// Ingress can both listen on a specific HTTP path
	// or listen to the / path but match the Host header
	//
	// Types that are assignable to IngressType:
	//	*Ingress_Path
	//	*Ingress_Host
	IngressType isIngress_IngressType `protobuf_oneof:"ingressType"`
}

func (x *Ingress) Reset() {
	*x = Ingress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ingress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ingress) ProtoMessage() {}

func (x *Ingress) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ingress.ProtoReflect.Descriptor instead.
func (*Ingress) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{4}
}

func (x *Ingress) GetContentMode() ContentMode {
	if x != nil {
		return x.ContentMode
	}
	return ContentMode_BINARY
}

func (m *Ingress) GetIngressType() isIngress_IngressType {
	if m != nil {
		return m.IngressType
	}
	return nil
}

func (x *Ingress) GetPath() string {
	if x, ok := x.GetIngressType().(*Ingress_Path); ok {
		return x.Path
	}
	return ""
}

func (x *Ingress) GetHost() string {
	if x, ok := x.GetIngressType().(*Ingress_Host); ok {
		return x.Host
	}
	return ""
}

type isIngress_IngressType interface {
	isIngress_IngressType()
}

type Ingress_Path struct {
	// path to listen for incoming events.
	Path string `protobuf:"bytes,2,opt,name=path,proto3,oneof"`
}

type Ingress_Host struct {
	// host header to match
	Host string `protobuf:"bytes,3,opt,name=host,proto3,oneof"`
}

func (*Ingress_Path) isIngress_IngressType() {}

func (*Ingress_Host) isIngress_IngressType() {}

// Kubernetes resource reference.
type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Object id.
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Object namespace.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Object name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Object version.
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{5}
}

func (x *Reference) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Reference) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Reference) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Reference) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of the resource
	// It's the same as the Kubernetes resource uid
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// Topics name
	// Note: If there is an ingress configured, then this field must have exactly 1 element otherwise,
	//  if the resource does just dispatch from Kafka, then this topic list can contain multiple elements
	Topics []string `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
	// A comma separated list of host/port pairs to use for establishing the initial connection to the Kafka cluster.
	// Note: we're using a comma separated list simply because that's how java kafka client likes it.
	BootstrapServers string `protobuf:"bytes,3,opt,name=bootstrapServers,proto3" json:"bootstrapServers,omitempty"`
	// Optional ingress for this topic
	Ingress *Ingress `protobuf:"bytes,4,opt,name=ingress,proto3" json:"ingress,omitempty"`
	// Optional configuration of egress valid for the whole resource
	EgressConfig *EgressConfig `protobuf:"bytes,5,opt,name=egressConfig,proto3" json:"egressConfig,omitempty"`
	// Optional egresses for this topic
	Egresses []*Egress `protobuf:"bytes,6,rep,name=egresses,proto3" json:"egresses,omitempty"`
	// Types that are assignable to Auth:
	//	*Resource_AbsentAuth
	//	*Resource_AuthSecret
	Auth isResource_Auth `protobuf_oneof:"Auth"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{6}
}

func (x *Resource) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Resource) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *Resource) GetBootstrapServers() string {
	if x != nil {
		return x.BootstrapServers
	}
	return ""
}

func (x *Resource) GetIngress() *Ingress {
	if x != nil {
		return x.Ingress
	}
	return nil
}

func (x *Resource) GetEgressConfig() *EgressConfig {
	if x != nil {
		return x.EgressConfig
	}
	return nil
}

func (x *Resource) GetEgresses() []*Egress {
	if x != nil {
		return x.Egresses
	}
	return nil
}

func (m *Resource) GetAuth() isResource_Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (x *Resource) GetAbsentAuth() *Empty {
	if x, ok := x.GetAuth().(*Resource_AbsentAuth); ok {
		return x.AbsentAuth
	}
	return nil
}

func (x *Resource) GetAuthSecret() *Reference {
	if x, ok := x.GetAuth().(*Resource_AuthSecret); ok {
		return x.AuthSecret
	}
	return nil
}

type isResource_Auth interface {
	isResource_Auth()
}

type Resource_AbsentAuth struct {
	// No auth configured.
	AbsentAuth *Empty `protobuf:"bytes,7,opt,name=absentAuth,proto3,oneof"`
}

type Resource_AuthSecret struct {
	// Secret reference.
	//
	// Secret format:
	//
	//   protocol: (PLAINTEXT | SASL_PLAINTEXT | SSL | SASL_SSL)
	//   sasl.mechanism: (SCRAM-SHA-256 | SCRAM-SHA-512)
	//   ca.crt: <CA PEM certificate>
	//   user.crt: <User PEM certificate>
	//   user.key: <User PEM key>
	//   user: <SASL username>
	//   password: <SASL password>
	//
	// Validation:
	//   - protocol=PLAINTEXT
	//   - protocol=SSL
	//     - required:
	//       - ca.crt
	//       - user.crt
	//       - user.key
	//   - protocol=SASL_PLAINTEXT
	//     - required:
	//       - sasl.mechanism
	//       - user
	//       - password
	//   - protocol=SASL_SSL
	//     - required:
	//       - sasl.mechanism
	//       - ca.crt
	//       - user.crt
	//       - user.key
	//       - user
	//       - password
	AuthSecret *Reference `protobuf:"bytes,8,opt,name=authSecret,proto3,oneof"`
}

func (*Resource_AbsentAuth) isResource_Auth() {}

func (*Resource_AuthSecret) isResource_Auth() {}

type Contract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Count each contract update.
	// Make sure each data plane pod has the same contract generation number.
	Generation uint64      `protobuf:"varint,1,opt,name=generation,proto3" json:"generation,omitempty"`
	Resources  []*Resource `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *Contract) Reset() {
	*x = Contract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contract) ProtoMessage() {}

func (x *Contract) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contract.ProtoReflect.Descriptor instead.
func (*Contract) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{7}
}

func (x *Contract) GetGeneration() uint64 {
	if x != nil {
		return x.Generation
	}
	return 0
}

func (x *Contract) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

var File_contract_proto protoreflect.FileDescriptor

var file_contract_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x80, 0x01, 0x0a, 0x06, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x3d, 0x0a,
	0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb8, 0x01, 0x0a,
	0x0c, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e, 0x0a,
	0x0a, 0x64, 0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x64, 0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x72, 0x65, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x65,
	0x74, 0x72, 0x79, 0x12, 0x34, 0x0a, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x42, 0x61, 0x63,
	0x6b, 0x6f, 0x66, 0x66, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0d, 0x62, 0x61, 0x63, 0x6b,
	0x6f, 0x66, 0x66, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x61, 0x63,
	0x6b, 0x6f, 0x66, 0x66, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0c, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0xd9, 0x02, 0x0a, 0x06, 0x45, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x08, 0x72, 0x65,
	0x70, 0x6c, 0x79, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08,
	0x72, 0x65, 0x70, 0x6c, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x3c, 0x0a, 0x14, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x54, 0x6f, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00,
	0x52, 0x14, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1f, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x0c, 0x65, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c,
	0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x34, 0x0a, 0x0d,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x0f, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x22, 0x74, 0x0a, 0x07, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e,
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x14,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x22, 0x6b, 0x0a, 0x09, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xbc, 0x02, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x2a, 0x0a,
	0x10, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72,
	0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x07, 0x69, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x31, 0x0a,
	0x0c, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x0c, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x23, 0x0a, 0x08, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x08, 0x65, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0a, 0x61, 0x62, 0x73, 0x65, 0x6e, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x62, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x2c, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x48,
	0x00, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x42, 0x06, 0x0a,
	0x04, 0x41, 0x75, 0x74, 0x68, 0x22, 0x53, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x27, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2a, 0x2c, 0x0a, 0x0d, 0x42, 0x61,
	0x63, 0x6b, 0x6f, 0x66, 0x66, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0f, 0x0a, 0x0b, 0x45,
	0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x10, 0x01, 0x2a, 0x2b, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x4f,
	0x52, 0x44, 0x45, 0x52, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x45, 0x44, 0x10, 0x01, 0x2a, 0x29, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x52, 0x55, 0x43, 0x54, 0x55, 0x52, 0x45, 0x44, 0x10, 0x01,
	0x42, 0x5b, 0x0a, 0x2a, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42, 0x11,
	0x44, 0x61, 0x74, 0x61, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x5a, 0x1a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_proto_rawDescOnce sync.Once
	file_contract_proto_rawDescData = file_contract_proto_rawDesc
)

func file_contract_proto_rawDescGZIP() []byte {
	file_contract_proto_rawDescOnce.Do(func() {
		file_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_proto_rawDescData)
	})
	return file_contract_proto_rawDescData
}

var file_contract_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_contract_proto_goTypes = []interface{}{
	(BackoffPolicy)(0),   // 0: BackoffPolicy
	(DeliveryOrder)(0),   // 1: DeliveryOrder
	(ContentMode)(0),     // 2: ContentMode
	(*Empty)(nil),        // 3: Empty
	(*Filter)(nil),       // 4: Filter
	(*EgressConfig)(nil), // 5: EgressConfig
	(*Egress)(nil),       // 6: Egress
	(*Ingress)(nil),      // 7: Ingress
	(*Reference)(nil),    // 8: Reference
	(*Resource)(nil),     // 9: Resource
	(*Contract)(nil),     // 10: Contract
	nil,                  // 11: Filter.AttributesEntry
}
var file_contract_proto_depIdxs = []int32{
	11, // 0: Filter.attributes:type_name -> Filter.AttributesEntry
	0,  // 1: EgressConfig.backoffPolicy:type_name -> BackoffPolicy
	3,  // 2: Egress.replyToOriginalTopic:type_name -> Empty
	4,  // 3: Egress.filter:type_name -> Filter
	5,  // 4: Egress.egressConfig:type_name -> EgressConfig
	1,  // 5: Egress.deliveryOrder:type_name -> DeliveryOrder
	2,  // 6: Ingress.contentMode:type_name -> ContentMode
	7,  // 7: Resource.ingress:type_name -> Ingress
	5,  // 8: Resource.egressConfig:type_name -> EgressConfig
	6,  // 9: Resource.egresses:type_name -> Egress
	3,  // 10: Resource.absentAuth:type_name -> Empty
	8,  // 11: Resource.authSecret:type_name -> Reference
	9,  // 12: Contract.resources:type_name -> Resource
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_contract_proto_init() }
func file_contract_proto_init() {
	if File_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EgressConfig); i {
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
		file_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Egress); i {
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
		file_contract_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ingress); i {
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
		file_contract_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reference); i {
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
		file_contract_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_contract_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contract); i {
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
	file_contract_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Egress_ReplyUrl)(nil),
		(*Egress_ReplyToOriginalTopic)(nil),
	}
	file_contract_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Ingress_Path)(nil),
		(*Ingress_Host)(nil),
	}
	file_contract_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*Resource_AbsentAuth)(nil),
		(*Resource_AuthSecret)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contract_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contract_proto_goTypes,
		DependencyIndexes: file_contract_proto_depIdxs,
		EnumInfos:         file_contract_proto_enumTypes,
		MessageInfos:      file_contract_proto_msgTypes,
	}.Build()
	File_contract_proto = out.File
	file_contract_proto_rawDesc = nil
	file_contract_proto_goTypes = nil
	file_contract_proto_depIdxs = nil
}
