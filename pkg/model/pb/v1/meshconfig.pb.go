// Code generated by protoc-gen-go. DO NOT EDIT.
// source: meshconfig.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MeshService struct {
	Id *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 所属网格ID
	MeshId *wrappers.StringValue `protobuf:"bytes,2,opt,name=mesh_id,proto3" json:"mesh_id,omitempty"`
	// 所属网格名字
	MeshName *wrappers.StringValue `protobuf:"bytes,3,opt,name=mesh_name,proto3" json:"mesh_name,omitempty"`
	// 所属网格token
	MeshToken *wrappers.StringValue `protobuf:"bytes,4,opt,name=mesh_token,proto3" json:"mesh_token,omitempty"`
	// 在北极星上面的服务的ID
	ServiceId *wrappers.StringValue `protobuf:"bytes,5,opt,name=service_id,proto3" json:"service_id,omitempty"`
	// 在北极星上面的服务名
	Service *wrappers.StringValue `protobuf:"bytes,6,opt,name=service,proto3" json:"service,omitempty"`
	// 在北极星上面的命名空间
	Namespace *wrappers.StringValue `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// 服务的部门
	Department *wrappers.StringValue `protobuf:"bytes,8,opt,name=department,proto3" json:"department,omitempty"`
	// 服务的业务
	Business *wrappers.StringValue `protobuf:"bytes,9,opt,name=business,proto3" json:"business,omitempty"`
	// 服务负责人
	Owners *wrappers.StringValue `protobuf:"bytes,10,opt,name=owners,proto3" json:"owners,omitempty"`
	// 在服务网格里面的命名空间
	MeshNamespace *wrappers.StringValue `protobuf:"bytes,11,opt,name=mesh_namespace,proto3" json:"mesh_namespace,omitempty"`
	// 在服务网格里面的服务名
	MeshService *wrappers.StringValue `protobuf:"bytes,12,opt,name=mesh_service,proto3" json:"mesh_service,omitempty"`
	// 是网格内部还是外部服务
	Location *wrappers.StringValue `protobuf:"bytes,13,opt,name=location,proto3" json:"location,omitempty"`
	// 可以被网格的哪些命名空间看到
	ExportTo *wrappers.StringValue `protobuf:"bytes,14,opt,name=export_to,proto3" json:"export_to,omitempty"`
	// 网格服务的revision
	Revision *wrappers.StringValue `protobuf:"bytes,15,opt,name=revision,proto3" json:"revision,omitempty"`
	// 服务订阅时间
	Ctime *wrappers.StringValue `protobuf:"bytes,16,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// 服务订阅信息修改时间
	Mtime                *wrappers.StringValue `protobuf:"bytes,17,opt,name=mtime,proto3" json:"mtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MeshService) Reset()         { *m = MeshService{} }
func (m *MeshService) String() string { return proto.CompactTextString(m) }
func (*MeshService) ProtoMessage()    {}
func (*MeshService) Descriptor() ([]byte, []int) {
	return fileDescriptor_meshconfig_6952e779a0f6f92b, []int{0}
}
func (m *MeshService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshService.Unmarshal(m, b)
}
func (m *MeshService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshService.Marshal(b, m, deterministic)
}
func (dst *MeshService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshService.Merge(dst, src)
}
func (m *MeshService) XXX_Size() int {
	return xxx_messageInfo_MeshService.Size(m)
}
func (m *MeshService) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshService.DiscardUnknown(m)
}

var xxx_messageInfo_MeshService proto.InternalMessageInfo

func (m *MeshService) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MeshService) GetMeshId() *wrappers.StringValue {
	if m != nil {
		return m.MeshId
	}
	return nil
}

func (m *MeshService) GetMeshName() *wrappers.StringValue {
	if m != nil {
		return m.MeshName
	}
	return nil
}

func (m *MeshService) GetMeshToken() *wrappers.StringValue {
	if m != nil {
		return m.MeshToken
	}
	return nil
}

func (m *MeshService) GetServiceId() *wrappers.StringValue {
	if m != nil {
		return m.ServiceId
	}
	return nil
}

func (m *MeshService) GetService() *wrappers.StringValue {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *MeshService) GetNamespace() *wrappers.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *MeshService) GetDepartment() *wrappers.StringValue {
	if m != nil {
		return m.Department
	}
	return nil
}

func (m *MeshService) GetBusiness() *wrappers.StringValue {
	if m != nil {
		return m.Business
	}
	return nil
}

func (m *MeshService) GetOwners() *wrappers.StringValue {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *MeshService) GetMeshNamespace() *wrappers.StringValue {
	if m != nil {
		return m.MeshNamespace
	}
	return nil
}

func (m *MeshService) GetMeshService() *wrappers.StringValue {
	if m != nil {
		return m.MeshService
	}
	return nil
}

func (m *MeshService) GetLocation() *wrappers.StringValue {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *MeshService) GetExportTo() *wrappers.StringValue {
	if m != nil {
		return m.ExportTo
	}
	return nil
}

func (m *MeshService) GetRevision() *wrappers.StringValue {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *MeshService) GetCtime() *wrappers.StringValue {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *MeshService) GetMtime() *wrappers.StringValue {
	if m != nil {
		return m.Mtime
	}
	return nil
}

type Mesh struct {
	Id *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 网格名字
	Name *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 网格所属业务
	Business *wrappers.StringValue `protobuf:"bytes,3,opt,name=business,proto3" json:"business,omitempty"`
	// 网格所属部门
	Department *wrappers.StringValue `protobuf:"bytes,4,opt,name=department,proto3" json:"department,omitempty"`
	// 网格版本号
	Revision *wrappers.StringValue `protobuf:"bytes,5,opt,name=revision,proto3" json:"revision,omitempty"`
	// 网格token
	Token *wrappers.StringValue `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
	// 网格属主
	Owners *wrappers.StringValue `protobuf:"bytes,7,opt,name=owners,proto3" json:"owners,omitempty"`
	// 是否为托管网格
	Managed *wrappers.BoolValue `protobuf:"bytes,8,opt,name=managed,proto3" json:"managed,omitempty"`
	// istio的版本
	IstioVersion *wrappers.StringValue `protobuf:"bytes,9,opt,name=istio_version,proto3" json:"istio_version,omitempty"`
	// 该网格的数据面集群
	DataCluster *wrappers.StringValue `protobuf:"bytes,10,opt,name=data_cluster,proto3" json:"data_cluster,omitempty"`
	// 网格订阅的服务
	Services []*MeshService `protobuf:"bytes,11,rep,name=services,proto3" json:"services,omitempty"`
	// 网格描述
	Comment *wrappers.StringValue `protobuf:"bytes,12,opt,name=comment,proto3" json:"comment,omitempty"`
	// 网格创建时间
	Ctime *wrappers.StringValue `protobuf:"bytes,13,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// 网格修改时间
	Mtime                *wrappers.StringValue `protobuf:"bytes,14,opt,name=mtime,proto3" json:"mtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Mesh) Reset()         { *m = Mesh{} }
func (m *Mesh) String() string { return proto.CompactTextString(m) }
func (*Mesh) ProtoMessage()    {}
func (*Mesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_meshconfig_6952e779a0f6f92b, []int{1}
}
func (m *Mesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mesh.Unmarshal(m, b)
}
func (m *Mesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mesh.Marshal(b, m, deterministic)
}
func (dst *Mesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mesh.Merge(dst, src)
}
func (m *Mesh) XXX_Size() int {
	return xxx_messageInfo_Mesh.Size(m)
}
func (m *Mesh) XXX_DiscardUnknown() {
	xxx_messageInfo_Mesh.DiscardUnknown(m)
}

var xxx_messageInfo_Mesh proto.InternalMessageInfo

func (m *Mesh) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Mesh) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Mesh) GetBusiness() *wrappers.StringValue {
	if m != nil {
		return m.Business
	}
	return nil
}

func (m *Mesh) GetDepartment() *wrappers.StringValue {
	if m != nil {
		return m.Department
	}
	return nil
}

func (m *Mesh) GetRevision() *wrappers.StringValue {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *Mesh) GetToken() *wrappers.StringValue {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *Mesh) GetOwners() *wrappers.StringValue {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *Mesh) GetManaged() *wrappers.BoolValue {
	if m != nil {
		return m.Managed
	}
	return nil
}

func (m *Mesh) GetIstioVersion() *wrappers.StringValue {
	if m != nil {
		return m.IstioVersion
	}
	return nil
}

func (m *Mesh) GetDataCluster() *wrappers.StringValue {
	if m != nil {
		return m.DataCluster
	}
	return nil
}

func (m *Mesh) GetServices() []*MeshService {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *Mesh) GetComment() *wrappers.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *Mesh) GetCtime() *wrappers.StringValue {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *Mesh) GetMtime() *wrappers.StringValue {
	if m != nil {
		return m.Mtime
	}
	return nil
}

type MeshResource struct {
	Id *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 所属网格ID
	MeshId *wrappers.StringValue `protobuf:"bytes,2,opt,name=mesh_id,proto3" json:"mesh_id,omitempty"`
	// 所属网格名字
	MeshName *wrappers.StringValue `protobuf:"bytes,3,opt,name=mesh_name,proto3" json:"mesh_name,omitempty"`
	// 所属网格命名空间
	MeshNamespace *wrappers.StringValue `protobuf:"bytes,4,opt,name=mesh_namespace,proto3" json:"mesh_namespace,omitempty"`
	// 网格规则的类型
	TypeUrl *wrappers.StringValue `protobuf:"bytes,5,opt,name=type_url,proto3" json:"type_url,omitempty"`
	// 网格规则名字
	Name *wrappers.StringValue `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// 网格规则修订号
	Revision *wrappers.StringValue `protobuf:"bytes,7,opt,name=revision,proto3" json:"revision,omitempty"`
	// 网格规则的内容
	Body *wrappers.StringValue `protobuf:"bytes,8,opt,name=body,proto3" json:"body,omitempty"`
	// 所属网格的token
	MeshToken *wrappers.StringValue `protobuf:"bytes,9,opt,name=mesh_token,proto3" json:"mesh_token,omitempty"`
	// 规则创建时间
	Ctime *wrappers.StringValue `protobuf:"bytes,10,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// 规则修改时间
	Mtime                *wrappers.StringValue `protobuf:"bytes,11,opt,name=mtime,proto3" json:"mtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MeshResource) Reset()         { *m = MeshResource{} }
func (m *MeshResource) String() string { return proto.CompactTextString(m) }
func (*MeshResource) ProtoMessage()    {}
func (*MeshResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_meshconfig_6952e779a0f6f92b, []int{2}
}
func (m *MeshResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshResource.Unmarshal(m, b)
}
func (m *MeshResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshResource.Marshal(b, m, deterministic)
}
func (dst *MeshResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshResource.Merge(dst, src)
}
func (m *MeshResource) XXX_Size() int {
	return xxx_messageInfo_MeshResource.Size(m)
}
func (m *MeshResource) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshResource.DiscardUnknown(m)
}

var xxx_messageInfo_MeshResource proto.InternalMessageInfo

func (m *MeshResource) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MeshResource) GetMeshId() *wrappers.StringValue {
	if m != nil {
		return m.MeshId
	}
	return nil
}

func (m *MeshResource) GetMeshName() *wrappers.StringValue {
	if m != nil {
		return m.MeshName
	}
	return nil
}

func (m *MeshResource) GetMeshNamespace() *wrappers.StringValue {
	if m != nil {
		return m.MeshNamespace
	}
	return nil
}

func (m *MeshResource) GetTypeUrl() *wrappers.StringValue {
	if m != nil {
		return m.TypeUrl
	}
	return nil
}

func (m *MeshResource) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MeshResource) GetRevision() *wrappers.StringValue {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *MeshResource) GetBody() *wrappers.StringValue {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *MeshResource) GetMeshToken() *wrappers.StringValue {
	if m != nil {
		return m.MeshToken
	}
	return nil
}

func (m *MeshResource) GetCtime() *wrappers.StringValue {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *MeshResource) GetMtime() *wrappers.StringValue {
	if m != nil {
		return m.Mtime
	}
	return nil
}

type MeshConfig struct {
	// 所属网格ID
	MeshId *wrappers.StringValue `protobuf:"bytes,1,opt,name=mesh_id,proto3" json:"mesh_id,omitempty"`
	// 所属网格名字
	MeshName *wrappers.StringValue `protobuf:"bytes,2,opt,name=mesh_name,proto3" json:"mesh_name,omitempty"`
	// 请求的配置类型
	TypeUrl *wrappers.StringValue `protobuf:"bytes,3,opt,name=type_url,proto3" json:"type_url,omitempty"`
	// 具体的各个网格规则
	Resources []*MeshResource `protobuf:"bytes,4,rep,name=resources,proto3" json:"resources,omitempty"`
	// 总体的修订版本号
	Revision             *wrappers.StringValue `protobuf:"bytes,5,opt,name=revision,proto3" json:"revision,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MeshConfig) Reset()         { *m = MeshConfig{} }
func (m *MeshConfig) String() string { return proto.CompactTextString(m) }
func (*MeshConfig) ProtoMessage()    {}
func (*MeshConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_meshconfig_6952e779a0f6f92b, []int{3}
}
func (m *MeshConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshConfig.Unmarshal(m, b)
}
func (m *MeshConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshConfig.Marshal(b, m, deterministic)
}
func (dst *MeshConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshConfig.Merge(dst, src)
}
func (m *MeshConfig) XXX_Size() int {
	return xxx_messageInfo_MeshConfig.Size(m)
}
func (m *MeshConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MeshConfig proto.InternalMessageInfo

func (m *MeshConfig) GetMeshId() *wrappers.StringValue {
	if m != nil {
		return m.MeshId
	}
	return nil
}

func (m *MeshConfig) GetMeshName() *wrappers.StringValue {
	if m != nil {
		return m.MeshName
	}
	return nil
}

func (m *MeshConfig) GetTypeUrl() *wrappers.StringValue {
	if m != nil {
		return m.TypeUrl
	}
	return nil
}

func (m *MeshConfig) GetResources() []*MeshResource {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *MeshConfig) GetRevision() *wrappers.StringValue {
	if m != nil {
		return m.Revision
	}
	return nil
}

func init() {
	proto.RegisterType((*MeshService)(nil), "v1.MeshService")
	proto.RegisterType((*Mesh)(nil), "v1.Mesh")
	proto.RegisterType((*MeshResource)(nil), "v1.MeshResource")
	proto.RegisterType((*MeshConfig)(nil), "v1.MeshConfig")
}

func init() { proto.RegisterFile("meshconfig.proto", fileDescriptor_meshconfig_6952e779a0f6f92b) }

var fileDescriptor_meshconfig_6952e779a0f6f92b = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x96, 0xbf, 0x6e, 0xdb, 0x30,
	0x10, 0xc6, 0x11, 0x59, 0xb1, 0xec, 0x53, 0xfe, 0xb8, 0x9c, 0x88, 0xa0, 0x28, 0x82, 0x4c, 0x01,
	0x5a, 0x28, 0x89, 0x1b, 0x14, 0x41, 0xd1, 0xa1, 0x48, 0xbb, 0x76, 0x71, 0x80, 0xae, 0x86, 0x2c,
	0x5d, 0x1c, 0xa2, 0x92, 0x28, 0x90, 0x94, 0xd3, 0x3c, 0x43, 0x5f, 0xa3, 0xaf, 0xd7, 0xad, 0x0f,
	0x50, 0x90, 0xfa, 0x63, 0x3b, 0x59, 0x4e, 0xde, 0xba, 0x59, 0xf2, 0xef, 0x33, 0x79, 0xbe, 0xef,
	0x3b, 0x12, 0x26, 0x39, 0xea, 0x87, 0x44, 0x16, 0xf7, 0x62, 0x19, 0x95, 0x4a, 0x1a, 0xc9, 0xbc,
	0xd5, 0xd5, 0xc9, 0x9b, 0xa5, 0x94, 0xcb, 0x0c, 0x2f, 0xdc, 0x9b, 0x45, 0x75, 0x7f, 0xf1, 0xa8,
	0xe2, 0xb2, 0x44, 0xa5, 0x6b, 0xe6, 0xec, 0x4f, 0x00, 0xe1, 0x37, 0xd4, 0x0f, 0x77, 0xa8, 0x56,
	0x22, 0x41, 0xf6, 0x0e, 0x3c, 0x91, 0xf2, 0xbd, 0xd3, 0xbd, 0xf3, 0x70, 0xfa, 0x3a, 0xaa, 0xc5,
	0x51, 0x2b, 0x8e, 0xee, 0x8c, 0x12, 0xc5, 0xf2, 0x7b, 0x9c, 0x55, 0x38, 0xf3, 0x44, 0xca, 0x3e,
	0x40, 0x60, 0x57, 0x9d, 0x8b, 0x94, 0x7b, 0x04, 0x49, 0x0b, 0xb3, 0x8f, 0x30, 0x76, 0x1f, 0x8b,
	0x38, 0x47, 0x3e, 0x20, 0x28, 0xd7, 0x38, 0xfb, 0x04, 0xe0, 0x1e, 0x8c, 0xfc, 0x81, 0x05, 0xf7,
	0x09, 0xe2, 0x0d, 0xde, 0xaa, 0x75, 0x5d, 0xaa, 0xdd, 0xf4, 0x3e, 0x45, 0xbd, 0xe6, 0x6d, 0xbd,
	0xcd, 0x13, 0x1f, 0x52, 0xea, 0x6d, 0x60, 0x5b, 0xaf, 0xdd, 0xbb, 0x2e, 0xe3, 0x04, 0x79, 0x40,
	0xa9, 0xb7, 0xc3, 0xed, 0x8e, 0x53, 0x2c, 0x63, 0x65, 0x72, 0x2c, 0x0c, 0x1f, 0x51, 0x76, 0xbc,
	0xe6, 0xd9, 0x0d, 0x8c, 0x16, 0x95, 0x16, 0x05, 0x6a, 0xcd, 0xc7, 0x04, 0x6d, 0x47, 0xb3, 0x6b,
	0x18, 0xca, 0xc7, 0x02, 0x95, 0xe6, 0x40, 0xd0, 0x35, 0x2c, 0xfb, 0x0a, 0x47, 0x5d, 0xab, 0xea,
	0x72, 0x43, 0x82, 0xfa, 0x99, 0x86, 0x7d, 0x86, 0x03, 0xf7, 0xa6, 0xfd, 0xb3, 0x0f, 0x08, 0xbf,
	0xb1, 0xa5, 0xb0, 0x75, 0x67, 0x32, 0x89, 0x8d, 0x90, 0x05, 0x3f, 0xa4, 0xd4, 0xdd, 0xd2, 0xb6,
	0x57, 0xf8, 0xb3, 0x94, 0xca, 0xcc, 0x8d, 0xe4, 0x47, 0x94, 0x5e, 0x75, 0xb8, 0x5d, 0x55, 0xe1,
	0x4a, 0x68, 0xbb, 0xea, 0x31, 0x65, 0xd5, 0x96, 0x66, 0x53, 0xd8, 0x4f, 0x8c, 0xc8, 0x91, 0x4f,
	0x08, 0xb2, 0x1a, 0xb5, 0x9a, 0xdc, 0x69, 0x5e, 0x51, 0x34, 0x0e, 0x3d, 0xfb, 0x35, 0x04, 0xdf,
	0xe6, 0xbd, 0x67, 0xd0, 0x2f, 0xc1, 0x77, 0x59, 0xa5, 0xa4, 0xdc, 0x91, 0x5b, 0xc6, 0x1b, 0xf4,
	0x32, 0xde, 0xb6, 0xe1, 0xfd, 0xfe, 0x86, 0xef, 0x5a, 0xb0, 0xdf, 0xb7, 0x05, 0xf5, 0x4c, 0xa1,
	0x44, 0xbb, 0x46, 0x37, 0x42, 0x12, 0xf4, 0x08, 0xc9, 0x35, 0x04, 0x79, 0x5c, 0xc4, 0x4b, 0x4c,
	0x9b, 0x3c, 0x9f, 0xbc, 0x90, 0xdd, 0x4a, 0x99, 0xb5, 0x43, 0xb3, 0x46, 0xd9, 0x2d, 0x1c, 0x0a,
	0x6d, 0x84, 0x9c, 0xaf, 0x50, 0xb9, 0xf2, 0x28, 0x79, 0xde, 0x96, 0xd8, 0x60, 0xa5, 0xb1, 0x89,
	0xe7, 0x49, 0x56, 0x69, 0x83, 0x8a, 0x14, 0xed, 0x2d, 0x05, 0x7b, 0x0b, 0xa3, 0x26, 0x63, 0x9a,
	0x87, 0xa7, 0x83, 0xf3, 0x70, 0x7a, 0x1c, 0xad, 0xae, 0xa2, 0x8d, 0x33, 0x64, 0xd6, 0x01, 0x76,
	0x5e, 0x26, 0x32, 0x77, 0x7d, 0xa4, 0x44, 0xb8, 0x85, 0xd7, 0x69, 0x38, 0xdc, 0x21, 0x0d, 0x47,
	0xf4, 0x34, 0xfc, 0xf5, 0xe1, 0xc0, 0xee, 0x7c, 0x86, 0x5a, 0x56, 0xea, 0xbf, 0x38, 0xfe, 0x5e,
	0x0e, 0x58, 0x7f, 0x87, 0x01, 0x7b, 0x03, 0x23, 0xf3, 0x54, 0xe2, 0xbc, 0x52, 0x19, 0x2d, 0x25,
	0x2d, 0xdd, 0x4d, 0x82, 0x61, 0x9f, 0x49, 0xd0, 0x25, 0x32, 0xe8, 0x95, 0xc8, 0x4b, 0xf0, 0x17,
	0x32, 0x7d, 0x22, 0x1d, 0x7a, 0x8e, 0x7c, 0x76, 0x39, 0x18, 0xf7, 0xbc, 0x1c, 0x74, 0xb6, 0x83,
	0x1d, 0x6c, 0x17, 0xd2, 0x6d, 0xf7, 0xdb, 0x03, 0xb0, 0xb6, 0xfb, 0xe2, 0x6e, 0x6b, 0x9b, 0x36,
	0xda, 0xdb, 0xd9, 0x46, 0x5e, 0x3f, 0x1b, 0x6d, 0x1a, 0x60, 0xd0, 0xcb, 0x00, 0x11, 0x8c, 0x55,
	0x13, 0x17, 0xcd, 0x7d, 0x37, 0x01, 0x26, 0xed, 0x04, 0x68, 0x73, 0x34, 0x5b, 0x23, 0xbb, 0x0f,
	0xe4, 0xc5, 0xd0, 0x7d, 0xff, 0xfe, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x0e, 0x3d, 0xfd,
	0xda, 0x0a, 0x00, 0x00,
}
