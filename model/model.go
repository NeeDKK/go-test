package model

type AllowGroup struct {
	Id           int64  `json:"id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
	GroupName    string `json:"group_name"`
	Description  string `json:"description"`
	Status       int64  `json:"status"`
	ProtoGroupId string `json:"proto_group_id"`
}
type AllowListIec104 struct {
	Id               int64  `json:"id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
	AllowListGroupId int64  `json:"allow_list_group_id"`
	DeviceName       string `json:"device_name"`
	SrcAndDestIp     string `json:"src_and_dest_ip"`
	SrcAndDestMac    string `json:"src_and_dest_mac"`
	AsduType         string `json:"asdu_type"`
	CausetxType      string `json:"causetx_type"`
}
type AllowListModbus struct {
	Id               int64  `json:"id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
	AllowListGroupId int64  `json:"allow_list_group_id"`
	DeviceName       string `json:"device_name"`
	SrcAndDestIp     string `json:"src_and_dest_ip"`
	SrcAndDestMac    string `json:"src_and_dest_mac"`
	Func             string `json:"func"`
	StartAddr        string `json:"start_addr"`
	EndAddr          string `json:"end_addr"`
}
type AllowListOpcda struct {
	Id               int64  `json:"id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
	AllowListGroupId int64  `json:"allow_list_group_id"`
	DeviceName       string `json:"device_name"`
	SrcAndDestIp     string `json:"src_and_dest_ip"`
	SrcAndDestMac    string `json:"src_and_dest_mac"`
	Uuid             string `json:"uuid"`
	Opnum            string `json:"opnum"`
}
type AllowListS7 struct {
	Id               int64  `json:"id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
	AllowListGroupId int64  `json:"allow_list_group_id"`
	DeviceName       string `json:"device_name"`
	SrcAndDestIp     string `json:"src_and_dest_ip"`
	SrcAndDestMac    string `json:"src_and_dest_mac"`
	PduType          string `json:"pdu_type"`
	DataType         string `json:"data_type"`
	FuncGroupType    string `json:"func_group_type"`
	SubFuncType      string `json:"sub_func_type"`
	FunctionType     string `json:"function_type"`
}
type Assert struct {
	Id                 int64  `json:"id"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	DeletedAt          string `json:"deleted_at"`
	DeviceName         string `json:"device_name"`
	DeviceType         string `json:"device_type"`
	DeviceLevel        int64  `json:"device_level"`
	DeviceIp           string `json:"device_ip"`
	DeviceMac          string `json:"device_mac"`
	DeviceManufacturer string `json:"device_manufacturer"`
	DeviceModel        string `json:"device_model"`
	SerialNumber       string `json:"serial_number"`
	DeviceVersion      string `json:"device_version"`
	DeviceRegion       string `json:"device_region"`
	DeviceNetwork      string `json:"device_network"`
	Protocol           string `json:"protocol"`
	OpenPort           string `json:"open_port"`
	UnsafeProtocol     string `json:"unsafe_protocol"`
}
type AuthorityMenu struct {
	Id          int64  `json:"id"`
	Path        string `json:"path"`
	Icon        string `json:"icon"`
	Name        string `json:"name"`
	Sort        int64  `json:"sort"`
	Title       string `json:"title"`
	Hidden      int64  `json:"hidden"`
	Component   string `json:"component"`
	ParentId    string `json:"parent_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
	KeepAlive   int64  `json:"keep_alive"`
	MenuLevel   int64  `json:"menu_level"`
	DefaultMenu int64  `json:"default_menu"`
	CloseTab    int64  `json:"close_tab"`
	MenuId      int64  `json:"menu_id"`
	AuthorityId string `json:"authority_id"`
}
type BaseLineGroup struct {
	Id          int64  `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
	GroupName   string `json:"group_name"`
	Description string `json:"description"`
	Status      int64  `json:"status"`
}
type BaseLineIpPort struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	Ip              string `json:"ip"`
	Port            string `json:"port"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	About           int64  `json:"about"`
	BaseLineGroupId int64  `json:"base_line_group_id"`
}
type BaseLineLog struct {
	Id          int64  `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
	Ip          string `json:"ip"`
	Mac         string `json:"mac"`
	Protocol    string `json:"protocol"`
	Port        string `json:"port"`
	Description string `json:"description"`
	BaseLineId  int64  `json:"base_line_id"`
}
type BaseLineNetwork struct {
	Id          int64  `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
	NetworkName string `json:"network_name"`
	Description string `json:"description"`
	SrcIp       string `json:"src_ip"`
	DestIp      string `json:"dest_ip"`
	DestPort    string `json:"dest_port"`
	Protocol    string `json:"protocol"`
	MaxTime     int64  `json:"max_time"`
	BaseLineId  int64  `json:"base_line_id"`
}
type BaseLineStudyTemp struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	SrcIp           string `json:"src_ip"`
	DestIp          string `json:"dest_ip"`
	Port            string `json:"port"`
	About           int64  `json:"about"`
	BaseLineGroupId int64  `json:"base_line_group_id"`
	Protocol        string `json:"protocol"`
}
type BaseLineWarning struct {
	Id               int64  `json:"id"`
	SourceIp         string `json:"source_ip"`
	DestinationIp    string `json:"destination_ip"`
	SourceMac        string `json:"source_mac"`
	DestinationMac   string `json:"destination_mac"`
	SourcePort       int64  `json:"source_port"`
	DestinationPort  int64  `json:"destination_port"`
	Protocol         string `json:"protocol"`
	AppLayerProtocol string `json:"app_layer_protocol"`
	PacketLength     int64  `json:"packet_length"`
	Packet           string `json:"packet"`
	SignatureMessage string `json:"signature_message"`
	Timestamp        string `json:"timestamp"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
	WarningNum       int64  `json:"warning_num"`
	HandleStatus     string `json:"handle_status"`
	HandleTime       string `json:"handle_time"`
	HandlePerson     string `json:"handle_person"`
	HandleContent    string `json:"handle_content"`
	BaselineGroupId  string `json:"baseline_group_id"`
}
type BehaviorDetection struct {
	RuleId       string `json:"rule_id"`
	RuleDetail   string `json:"rule_detail"`
	RuleLevel    int64  `json:"rule_level"`
	RuleProtocol string `json:"rule_protocol"`
	RuleType     string `json:"rule_type"`
	RuleStatus   int64  `json:"rule_status"`
	Sid          string `json:"sid"`
}
type BehaviorLog struct {
	Id         int64  `json:"id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
	Ip         string `json:"ip"`
	Mac        string `json:"mac"`
	ThreatType string `json:"threat_type"`
	Protocol   string `json:"protocol"`
	FlowData   string `json:"flow_data"`
	Severity   string `json:"severity"`
	Port       string `json:"port"`
}
type BehaviorWarning struct {
	Id                  int64  `json:"id"`
	SourceIp            string `json:"source_ip"`
	DestinationIp       string `json:"destination_ip"`
	SourceMac           string `json:"source_mac"`
	DestinationMac      string `json:"destination_mac"`
	SourcePort          int64  `json:"source_port"`
	DestinationPort     int64  `json:"destination_port"`
	Protocol            string `json:"protocol"`
	AppLayerProtocol    string `json:"app_layer_protocol"`
	PacketLength        int64  `json:"packet_length"`
	Packet              string `json:"packet"`
	SignatureMessage    string `json:"signature_message"`
	Timestamp           string `json:"timestamp"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
	DeletedAt           string `json:"deleted_at"`
	WarningNum          int64  `json:"warning_num"`
	HandleStatus        string `json:"handle_status"`
	HandleTime          string `json:"handle_time"`
	HandlePerson        string `json:"handle_person"`
	HandleContent       string `json:"handle_content"`
	RuleLevel           int64  `json:"rule_level"`
	BehaviorDescription string `json:"behavior_description"`
	Sid                 string `json:"sid"`
}
type BlacklistWarning struct {
	Id               int64  `json:"id"`
	SourceIp         string `json:"source_ip"`
	DestinationIp    string `json:"destination_ip"`
	SourceMac        string `json:"source_mac"`
	DestinationMac   string `json:"destination_mac"`
	SourcePort       int64  `json:"source_port"`
	DestinationPort  int64  `json:"destination_port"`
	Protocol         string `json:"protocol"`
	AppLayerProtocol string `json:"app_layer_protocol"`
	PacketLength     int64  `json:"packet_length"`
	Packet           string `json:"packet"`
	SignatureMessage string `json:"signature_message"`
	Timestamp        string `json:"timestamp"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
	WarningNum       int64  `json:"warning_num"`
	HandleStatus     string `json:"handle_status"`
	HandleTime       string `json:"handle_time"`
	HandlePerson     string `json:"handle_person"`
	HandleContent    string `json:"handle_content"`
	Sid              string `json:"sid"`
	WarningLevel     string `json:"warning_level"`
	WarningContent   string `json:"warning_content"`
}
type CasbinRule struct {
	PType string `json:"p_type"`
	V0    string `json:"v0"`
	V1    string `json:"v1"`
	V2    string `json:"v2"`
	V3    string `json:"v3"`
	V4    string `json:"v4"`
	V5    string `json:"v5"`
}
type DeviceName struct {
	MacPrefix  string `json:"mac_prefix"`
	DeviceName string `json:"device_name"`
}
type ExaCustomers struct {
	Id                 int64  `json:"id"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	DeletedAt          string `json:"deleted_at"`
	CustomerName       string `json:"customer_name"`
	CustomerPhoneData  string `json:"customer_phone_data"`
	SysUserId          int64  `json:"sys_user_id"`
	SysUserAuthorityId string `json:"sys_user_authority_id"`
}
type ExaFiles struct {
	Id         int64  `json:"id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
	FileName   string `json:"file_name"`
	FileMd5    string `json:"file_md5"`
	FilePath   string `json:"file_path"`
	ChunkTotal int64  `json:"chunk_total"`
	IsFinish   int64  `json:"is_finish"`
}
type ExaFileChunks struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	ExaFileId       int64  `json:"exa_file_id"`
	FileChunkNumber int64  `json:"file_chunk_number"`
	FileChunkPath   string `json:"file_chunk_path"`
}
type ExaFileUploadAndDownloads struct {
	Id        int64  `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Tag       string `json:"tag"`
	Key       string `json:"key"`
}
type FlowDatasDnp3 struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	FlowdataHeadId  int64  `json:"flowdata_head_id"`
	PacketLenth     int64  `json:"packet_lenth"`
	PacketTimestamp string `json:"packet_timestamp"`
	FlowTimestamp   string `json:"flow_timestamp"`
	Direction       int64  `json:"direction"`
	Func            string `json:"func"`
}
type FlowDatasEnipio struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	FlowdataHeadId  int64  `json:"flowdata_head_id"`
	PacketLenth     int64  `json:"packet_lenth"`
	PacketTimestamp string `json:"packet_timestamp"`
	FlowTimestamp   string `json:"flow_timestamp"`
	Direction       int64  `json:"direction"`
	AddressType     string `json:"address_type"`
	DataType        string `json:"data_type"`
}
type FlowDatasEniptcp struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	FlowdataHeadId  int64  `json:"flowdata_head_id"`
	PacketLenth     int64  `json:"packet_lenth"`
	PacketTimestamp string `json:"packet_timestamp"`
	FlowTimestamp   string `json:"flow_timestamp"`
	Direction       int64  `json:"direction"`
	Command         string `json:"command"`
	ServiceName     string `json:"service_name"`
	AddressType     string `json:"address_type"`
	DataType        string `json:"data_type"`
}
type FlowDatasEnipudp struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	FlowdataHeadId  int64  `json:"flowdata_head_id"`
	PacketLenth     int64  `json:"packet_lenth"`
	PacketTimestamp string `json:"packet_timestamp"`
	FlowTimestamp   string `json:"flow_timestamp"`
	Direction       int64  `json:"direction"`
	Command         string `json:"command"`
}
type FlowDatasFtp struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	FlowdataHeadId  int64  `json:"flowdata_head_id"`
	PacketLenth     int64  `json:"packet_lenth"`
	PacketTimestamp string `json:"packet_timestamp"`
	FlowTimestamp   string `json:"flow_timestamp"`
	Direction       int64  `json:"direction"`
	AccountName     string `json:"account_name"`
	Command         string `json:"command"`
	ProtocolDetail  string `json:"protocol_detail"`
}
type FlowDatasGoose struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	FlowdataHeadId  int64  `json:"flowdata_head_id"`
	PacketLenth     int64  `json:"packet_lenth"`
	PacketTimestamp string `json:"packet_timestamp"`
	FlowTimestamp   string `json:"flow_timestamp"`
	Direction       int64  `json:"direction"`
	DatSet          string `json:"dat_set"`
	GoId            string `json:"go_id"`
}
type FlowDatasHttp struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	FlowdataHeadId  int64  `json:"flowdata_head_id"`
	PacketLenth     int64  `json:"packet_lenth"`
	PacketTimestamp string `json:"packet_timestamp"`
	FlowTimestamp   string `json:"flow_timestamp"`
	Direction       int64  `json:"direction"`
	Url             string `json:"url"`
	ProtocolDetail  string `json:"protocol_detail"`
}
type FlowDatasIec104 struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	FlowdataHeadId  int64  `json:"flowdata_head_id"`
	PacketLenth     int64  `json:"packet_lenth"`
	PacketTimestamp string `json:"packet_timestamp"`
	FlowTimestamp   string `json:"flow_timestamp"`
	Direction       int64  `json:"direction"`
	CausetxType     string `json:"causetx_type"`
	AsduType        string `json:"asdu_type"`
}
type FlowDatasMms struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	FlowdataHeadId  int64  `json:"flowdata_head_id"`
	PacketLenth     int64  `json:"packet_lenth"`
	PacketTimestamp string `json:"packet_timestamp"`
	FlowTimestamp   string `json:"flow_timestamp"`
	Direction       int64  `json:"direction"`
	PduType         string `json:"pdu_type"`
	ServiceRequest  string `json:"service_request"`
}
type FlowDatasModbus struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	FlowdataHeadId  int64  `json:"flowdata_head_id"`
	PacketLenth     int64  `json:"packet_lenth"`
	PacketTimestamp string `json:"packet_timestamp"`
	FlowTimestamp   string `json:"flow_timestamp"`
	Direction       int64  `json:"direction"`
	Func            string `json:"func"`
	StartAddr       string `json:"start_addr"`
	EndAddr         string `json:"end_addr"`
}
type FlowDatasOpcda struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	FlowdataHeadId  int64  `json:"flowdata_head_id"`
	PacketLenth     int64  `json:"packet_lenth"`
	PacketTimestamp string `json:"packet_timestamp"`
	FlowTimestamp   string `json:"flow_timestamp"`
	Direction       int64  `json:"direction"`
	OpInt           string `json:"op_int"`
	OpCode          string `json:"op_code"`
}
type FlowDatasOpcua struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	FlowdataHeadId  int64  `json:"flowdata_head_id"`
	PacketLenth     int64  `json:"packet_lenth"`
	PacketTimestamp string `json:"packet_timestamp"`
	FlowTimestamp   string `json:"flow_timestamp"`
	Direction       int64  `json:"direction"`
	ServiceId       string `json:"service_id"`
}
type FlowDatasPnrtdcp struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	FlowdataHeadId  int64  `json:"flowdata_head_id"`
	PacketLenth     int64  `json:"packet_lenth"`
	PacketTimestamp string `json:"packet_timestamp"`
	FlowTimestamp   string `json:"flow_timestamp"`
	Direction       int64  `json:"direction"`
	Frameid         string `json:"frameid"`
	Serviceid       string `json:"serviceid"`
	Servicetype     string `json:"servicetype"`
	Dcpoption       string `json:"dcpoption"`
	Dcpsuboption    string `json:"dcpsuboption"`
}
type FlowDatasPop3 struct {
	Id               int64  `json:"id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
	FlowdataHeadId   int64  `json:"flowdata_head_id"`
	PacketLenth      int64  `json:"packet_lenth"`
	PacketTimestamp  string `json:"packet_timestamp"`
	FlowTimestamp    string `json:"flow_timestamp"`
	Direction        int64  `json:"direction"`
	SouceMailAddress string `json:"souce_mail_address"`
	DestMailAddress  string `json:"dest_mail_address"`
	ProtocolDetail   string `json:"protocol_detail"`
}
type FlowDatasProfinetio struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	FlowdataHeadId  int64  `json:"flowdata_head_id"`
	PacketLenth     int64  `json:"packet_lenth"`
	PacketTimestamp string `json:"packet_timestamp"`
	FlowTimestamp   string `json:"flow_timestamp"`
	Direction       int64  `json:"direction"`
	Func            string `json:"func"`
	OpInt           string `json:"op_int"`
	DataType        string `json:"data_type"`
}
type FlowDatasS7 struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	FlowdataHeadId  int64  `json:"flowdata_head_id"`
	PacketLenth     int64  `json:"packet_lenth"`
	PacketTimestamp string `json:"packet_timestamp"`
	FlowTimestamp   string `json:"flow_timestamp"`
	Direction       int64  `json:"direction"`
	PduType         string `json:"pdu_type"`
	OpType          string `json:"op_type"`
	DataType        string `json:"data_type"`
}
type FlowDatasSmtp struct {
	Id               int64  `json:"id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
	FlowdataHeadId   int64  `json:"flowdata_head_id"`
	PacketLenth      int64  `json:"packet_lenth"`
	PacketTimestamp  string `json:"packet_timestamp"`
	FlowTimestamp    string `json:"flow_timestamp"`
	Direction        int64  `json:"direction"`
	SouceMailAddress string `json:"souce_mail_address"`
	DestMailAddress  string `json:"dest_mail_address"`
	ProtocolDetail   string `json:"protocol_detail"`
}
type FlowDatasSnmp struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	FlowdataHeadId  int64  `json:"flowdata_head_id"`
	PacketLenth     int64  `json:"packet_lenth"`
	PacketTimestamp string `json:"packet_timestamp"`
	FlowTimestamp   string `json:"flow_timestamp"`
	Direction       int64  `json:"direction"`
	PduType         string `json:"pdu_type"`
	Version         string `json:"version"`
	Community       string `json:"community"`
}
type FlowDatasSv struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	FlowdataHeadId  int64  `json:"flowdata_head_id"`
	PacketLenth     int64  `json:"packet_lenth"`
	PacketTimestamp string `json:"packet_timestamp"`
	FlowTimestamp   string `json:"flow_timestamp"`
	Direction       int64  `json:"direction"`
	SvId            string `json:"sv_id"`
	SmpSynch        string `json:"smp_synch"`
}
type FlowDatasTelnet struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	FlowdataHeadId  int64  `json:"flowdata_head_id"`
	PacketLenth     int64  `json:"packet_lenth"`
	PacketTimestamp string `json:"packet_timestamp"`
	FlowTimestamp   string `json:"flow_timestamp"`
	Direction       int64  `json:"direction"`
	AccountName     string `json:"account_name"`
	Command         string `json:"command"`
	ProtocolDetail  string `json:"protocol_detail"`
}
type FlowDataHeads struct {
	Id                 int64  `json:"id"`
	FlowTimestamp      string `json:"flow_timestamp"`
	PacketTimestamp    string `json:"packet_timestamp"`
	SourceMac          string `json:"source_mac"`
	DestinationMac     string `json:"destination_mac"`
	IpVersion          int64  `json:"ip_version"`
	SourcePort         int64  `json:"source_port"`
	SourceIp           string `json:"source_ip"`
	DestinationIp      string `json:"destination_ip"`
	DestinationPort    int64  `json:"destination_port"`
	ProtocolType       int64  `json:"protocol_type"`
	ProtocolTypeName   string `json:"protocol_type_name"`
	PacketLenth        int64  `json:"packet_lenth"`
	PacketContent      string `json:"packet_content"`
	ProtocolSourceName string `json:"protocol_source_name"`
	Directions         int64  `json:"directions"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	DeletedAt          string `json:"deleted_at"`
}
type Icdevicetraffics struct {
	Id          int64  `json:"id"`
	IcDeviceIp  string `json:"ic_device_ip"`
	IcDeviceMac string `json:"ic_device_mac"`
	IcDeviceId  string `json:"ic_device_id"`
	TrafficType int64  `json:"traffic_type"`
	TrafficName string `json:"traffic_name"`
	SendBytes   int64  `json:"send_bytes"`
	RecvBytes   int64  `json:"recv_bytes"`
	SendSpeed   int64  `json:"send_speed"`
	RecvSpeed   int64  `json:"recv_speed"`
	Timestamp   string `json:"timestamp"`
	DeviceName  string `json:"device_name"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
type Icdevicetrafficstats struct {
	Id          int64  `json:"id"`
	IcDeviceIp  string `json:"ic_device_ip"`
	IcDeviceMac string `json:"ic_device_mac"`
	IcDeviceId  string `json:"ic_device_id"`
	TrafficType int64  `json:"traffic_type"`
	TrafficName string `json:"traffic_name"`
	TotalBytes  int64  `json:"total_bytes"`
	Timestamp   string `json:"timestamp"`
	DeviceName  string `json:"device_name"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
type Incidents struct {
	Id               int64  `json:"id"`
	SourceName       string `json:"source_name"`
	DestinationName  string `json:"destination_name"`
	SourceIp         string `json:"source_ip"`
	DestinationIp    string `json:"destination_ip"`
	SourceMac        string `json:"source_mac"`
	DestinationMac   string `json:"destination_mac"`
	IpVersion        int64  `json:"ip_version"`
	SourcePort       int64  `json:"source_port"`
	DestinationPort  int64  `json:"destination_port"`
	Action           int64  `json:"action"`
	Protocol         string `json:"protocol"`
	AppLayerProtocol string `json:"app_layer_protocol"`
	PacketLength     int64  `json:"packet_length"`
	Packet           string `json:"packet"`
	SignatureMessage string `json:"signature_message"`
	SignatureId      int64  `json:"signature_id"`
	MatchedKey       string `json:"matched_key"`
	ProtocolDetail   string `json:"protocol_detail"`
	PayloadLength    int64  `json:"payload_length"`
	Payload          string `json:"payload"`
	SignatureName    string `json:"signature_name"`
	Timestamp        string `json:"timestamp"`
	OccurredDate     string `json:"occurred_date"`
	OccurredTime     string `json:"occurred_time"`
	Dpi              string `json:"dpi"`
	DpiName          string `json:"dpi_name"`
	Level            int64  `json:"level"`
	Status           int64  `json:"status"`
	Deployed         int64  `json:"deployed"`
	AlertType        int64  `json:"alert_type"`
	AlertId          int64  `json:"alert_id"`
	ZoneName         string `json:"zone_name"`
	Memo             string `json:"memo"`
	DpiIp            string `json:"dpi_ip"`
	BoxId            string `json:"box_id"`
	DeviceId         string `json:"device_id"`
	FutureAction     int64  `json:"future_action"`
	TopologyId       string `json:"topology_id"`
	DeviceName       string `json:"device_name"`
	StatisticDetail  string `json:"statistic_detail"`
	RiskLevel        int64  `json:"risk_level"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
	Sid              string `json:"sid"`
}
type IpMacBind struct {
	Id        int64  `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	Remark    string `json:"remark"`
	Ip        string `json:"ip"`
	Mac       string `json:"mac"`
	Status    int64  `json:"status"`
	Hits      int64  `json:"hits"`
}
type IpMacWarning struct {
	Id            int64  `json:"id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at"`
	Ip            string `json:"ip"`
	Mac           string `json:"mac"`
	Description   string `json:"description"`
	HitNum        int64  `json:"hit_num"`
	ProtoType     string `json:"proto_type"`
	HandleStatus  string `json:"handle_status"`
	HandleTime    string `json:"handle_time"`
	HandlePerson  string `json:"handle_person"`
	HandleContent string `json:"handle_content"`
	Timestamp     string `json:"timestamp"`
}
type JwtBlacklists struct {
	Id        int64  `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	Jwt       string `json:"jwt"`
}
type KeyValues struct {
	Id          int64  `json:"id"`
	ParentId    int64  `json:"parent_id"`
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
type LicenseRecord struct {
	Id            int64  `json:"id"`
	Sn            string `json:"sn"`
	Prod          string `json:"prod"`
	Project       string `json:"project"`
	Pid           string `json:"pid"`
	Vid           string `json:"vid"`
	Sku           string `json:"sku"`
	ArgsValue     string `json:"args_value"`
	ApplyDate     string `json:"apply_date"`
	ExpireDate    string `json:"expire_date"`
	EffectiveDay  int64  `json:"effective_day"`
	EffectiveTime string `json:"effective_time"`
	Type          int64  `json:"type"`
	Uuid          string `json:"uuid"`
	Signature     string `json:"signature"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at"`
}
type NetworkWarning struct {
	Id                 int64  `json:"id"`
	SourceIp           string `json:"source_ip"`
	DestinationIp      string `json:"destination_ip"`
	SourceMac          string `json:"source_mac"`
	DestinationMac     string `json:"destination_mac"`
	SourcePort         int64  `json:"source_port"`
	DestinationPort    int64  `json:"destination_port"`
	Protocol           string `json:"protocol"`
	AppLayerProtocol   string `json:"app_layer_protocol"`
	PacketLength       int64  `json:"packet_length"`
	Packet             string `json:"packet"`
	SignatureMessage   string `json:"signature_message"`
	Timestamp          string `json:"timestamp"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	DeletedAt          string `json:"deleted_at"`
	WarningNum         int64  `json:"warning_num"`
	HandleStatus       string `json:"handle_status"`
	HandleTime         string `json:"handle_time"`
	HandlePerson       string `json:"handle_person"`
	HandleContent      string `json:"handle_content"`
	WarningDescription string `json:"warning_description"`
	Sid                string `json:"sid"`
}
type NicStatistic struct {
	Id        int64   `json:"id"`
	Eth       string  `json:"eth"`
	RxBytes   float64 `json:"rx_bytes"`
	TxBytes   float64 `json:"tx_bytes"`
	Timestamp string  `json:"timestamp"`
}
type ProtocolField struct {
	Id        int64  `json:"id"`
	Pid       int64  `json:"pid"`
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	Offset    int64  `json:"offset"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
type ProtoDetailRelation struct {
	Id           int64  `json:"id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
	AllowGroupId int64  `json:"allow_group_id"`
	ProtoGroupId int64  `json:"proto_group_id"`
	ProtoDetail  string `json:"proto_detail"`
}
type ProtoGroup struct {
	Id          int64  `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
	GroupName   string `json:"group_name"`
	Description string `json:"description"`
}
type ProtoIdentification struct {
	Id                int64  `json:"id"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	DeletedAt         string `json:"deleted_at"`
	ProtocolName      string `json:"protocol_name"`
	Reviation         string `json:"reviation"`
	Describe          string `json:"describe"`
	TransportProtocol string `json:"transport_protocol"`
	Port              int64  `json:"port"`
	Status            int64  `json:"status"`
	IsCustom          int64  `json:"is_custom"`
	EigenOffset       int64  `json:"eigen_offset"`
	EigenValue        string `json:"eigen_value"`
	CheckModel        int64  `json:"check_model"`
	ProtoNum          int64  `json:"proto_num"`
}
type ProtoIec104 struct {
	Id               int64  `json:"id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
	ProtoGroupId     int64  `json:"proto_group_id"`
	DeviceName       string `json:"device_name"`
	SrcIp            string `json:"src_ip"`
	DestIp           string `json:"dest_ip"`
	TransmissionMode string `json:"transmission_mode"`
	AsduType         string `json:"asdu_type"`
	CausetxType      string `json:"causetx_type"`
}
type ProtoModbus struct {
	Id               int64  `json:"id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
	ProtoGroupId     int64  `json:"proto_group_id"`
	DeviceName       string `json:"device_name"`
	SrcIp            string `json:"src_ip"`
	DestIp           string `json:"dest_ip"`
	TransmissionMode string `json:"transmission_mode"`
	ObjectName       string `json:"object_name"`
	Func             string `json:"func"`
	StartAddr        string `json:"start_addr"`
	EndAddr          string `json:"end_addr"`
}
type ProtoOpcda struct {
	Id               int64  `json:"id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
	ProtoGroupId     int64  `json:"proto_group_id"`
	DeviceName       string `json:"device_name"`
	SrcIp            string `json:"src_ip"`
	DestIp           string `json:"dest_ip"`
	TransmissionMode string `json:"transmission_mode"`
	Uuid             string `json:"uuid"`
	Opnum            string `json:"opnum"`
}
type ProtoS7 struct {
	Id               int64  `json:"id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
	ProtoGroupId     int64  `json:"proto_group_id"`
	DeviceName       string `json:"device_name"`
	SrcIp            string `json:"src_ip"`
	DestIp           string `json:"dest_ip"`
	TransmissionMode string `json:"transmission_mode"`
	PduType          string `json:"pdu_type"`
	DataType         string `json:"data_type"`
	FuncGroupType    string `json:"func_group_type"`
	SubFuncType      string `json:"sub_func_type"`
	FunctionType     string `json:"function_type"`
}
type RegionManagement struct {
	Id                int64  `json:"id"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	DeletedAt         string `json:"deleted_at"`
	RegionName        string `json:"region_name"`
	RegionDescription string `json:"region_description"`
	ParentId          int64  `json:"parent_id"`
}
type Reports struct {
	Id         int64  `json:"id"`
	Key        string `json:"key"`
	Creator    string `json:"creator"`
	ReportName string `json:"report_name"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	Tag        string `json:"tag"`
	Type       string `json:"type"`
	Comment    string `json:"comment"`
	Status     int64  `json:"status"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
type Rules struct {
	Id          int64  `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
	SourceIp    string `json:"source_ip"`
	DestIp      string `json:"dest_ip"`
	SourceMac   string `json:"source_mac"`
	DestMac     string `json:"dest_mac"`
	Protocol    string `json:"protocol"`
	Command     string `json:"command"`
	AddressType int64  `json:"address_type"`
	DataType    int64  `json:"data_type"`
	Service     string `json:"service"`
	RuleGroup   string `json:"rule_group"`
}
type SessionStatistic struct {
	Time      string `json:"time"`
	ConcCount int64  `json:"conc_count"`
	NewCount  int64  `json:"new_count"`
}
type Signatures struct {
	SignatureId       string `json:"signatureId"`
	VulnerabilityId   string `json:"vulnerabilityId"`
	Priority          int64  `json:"priority"`
	Body              string `json:"body"`
	Rev               int64  `json:"rev"`
	Sid               int64  `json:"sid"`
	Msg               string `json:"msg"`
	Signame           string `json:"signame"`
	Flowbits          string `json:"flowbits"`
	Action            int64  `json:"action"`
	SourceSignatureId string `json:"sourceSignatureId"`
	Status            int64  `json:"status"`
	RiskLevel         int64  `json:"riskLevel"`
	Checked           int64  `json:"checked"`
}
type SysApis struct {
	Id          int64  `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
	Path        string `json:"path"`
	Description string `json:"description"`
	ApiGroup    string `json:"api_group"`
	Method      string `json:"method"`
}
type SysAuthorities struct {
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at"`
	AuthorityId   string `json:"authority_id"`
	AuthorityName string `json:"authority_name"`
	ParentId      string `json:"parent_id"`
	DefaultRouter string `json:"default_router"`
}
type SysAuthorityMenus struct {
	SysAuthorityAuthorityId string `json:"sys_authority_authority_id"`
	SysBaseMenuId           int64  `json:"sys_base_menu_id"`
}
type SysAutoCodeHistories struct {
	Id            int64  `json:"id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at"`
	TableName     string `json:"table_name"`
	RequestMeta   string `json:"request_meta"`
	AutoCodePath  string `json:"auto_code_path"`
	InjectionMeta string `json:"injection_meta"`
	StructName    string `json:"struct_name"`
	StructCnName  string `json:"struct_cn_name"`
	ApiIds        string `json:"api_ids"`
	Flag          int64  `json:"flag"`
}
type SysBaseMenus struct {
	Id          int64  `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
	MenuLevel   int64  `json:"menu_level"`
	ParentId    string `json:"parent_id"`
	Path        string `json:"path"`
	Name        string `json:"name"`
	Hidden      int64  `json:"hidden"`
	Component   string `json:"component"`
	Sort        int64  `json:"sort"`
	KeepAlive   int64  `json:"keep_alive"`
	DefaultMenu int64  `json:"default_menu"`
	Title       string `json:"title"`
	Icon        string `json:"icon"`
	CloseTab    int64  `json:"close_tab"`
}
type SysBaseMenuParameters struct {
	Id            int64  `json:"id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at"`
	SysBaseMenuId int64  `json:"sys_base_menu_id"`
	Type          string `json:"type"`
	Key           string `json:"key"`
	Value         string `json:"value"`
}
type SysDataAuthorityId struct {
	SysAuthorityAuthorityId    string `json:"sys_authority_authority_id"`
	DataAuthorityIdAuthorityId string `json:"data_authority_id_authority_id"`
}
type SysDictionaries struct {
	Id        int64  `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Status    int64  `json:"status"`
	Desc      string `json:"desc"`
}
type SysDictionaryDetails struct {
	Id              int64  `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
	Label           string `json:"label"`
	Value           int64  `json:"value"`
	Status          int64  `json:"status"`
	Sort            int64  `json:"sort"`
	SysDictionaryId int64  `json:"sys_dictionary_id"`
}
type SysOperationRecords struct {
	Id             int64  `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
	Ip             string `json:"ip"`
	Method         string `json:"method"`
	Path           string `json:"path"`
	Status         int64  `json:"status"`
	Latency        int64  `json:"latency"`
	Agent          string `json:"agent"`
	ErrorMessage   string `json:"error_message"`
	Body           string `json:"body"`
	Resp           string `json:"resp"`
	UserId         int64  `json:"user_id"`
	OperationModel string `json:"operation_model"`
	Detail         string `json:"detail"`
}
type SysUsers struct {
	Id            int64  `json:"id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at"`
	Uuid          string `json:"uuid"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	NickName      string `json:"nick_name"`
	SideMode      string `json:"side_mode"`
	HeaderImg     string `json:"header_img"`
	BaseColor     string `json:"base_color"`
	ActiveColor   string `json:"active_color"`
	AuthorityId   string `json:"authority_id"`
	PhoneNum      string `json:"phone_num"`
	Status        string `json:"status"`
	LastLoginTime string `json:"last_login_time"`
}
type SysUserAuthority struct {
	SysAuthorityAuthorityId string `json:"sys_authority_authority_id"`
	SysUserId               int64  `json:"sys_user_id"`
}
type TempRuler struct {
	Id          int64  `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
	TempRuler   string `json:"temp_ruler"`
	AllowListId string `json:"allow_list_id"`
}
type Test struct {
	Test string `json:"test"`
}
type ThreatLog struct {
	Id         int64  `json:"id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
	Ip         string `json:"ip"`
	Mac        string `json:"mac"`
	ThreatType string `json:"threat_type"`
	Protocol   string `json:"protocol"`
	ThreatName string `json:"threat_name"`
	FlowData   string `json:"flow_data"`
	Severity   string `json:"severity"`
	Port       string `json:"port"`
}
type UpgradeRecord struct {
	Id           int64  `json:"id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
	FileName     string `json:"file_name"`
	Version      string `json:"version"`
	Model        string `json:"model"`
	IsBack       string `json:"is_back"`
	LastRecordId string `json:"last_record_id"`
	IsEffective  string `json:"is_effective"`
}
type UserSetting struct {
	Id          int64  `json:"id"`
	SettingCode string `json:"setting_code"`
	SettingName string `json:"setting_name"`
	Description string `json:"description"`
	Args        string `json:"args"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
type Vulnerabilities struct {
	VulnerabilityId string `json:"vulnerabilityId"`
	Severity        int64  `json:"severity"`
	Category        string `json:"category"`
	ThreatName      string `json:"threatName"`
	ThreatNameEng   string `json:"threatNameEng"`
	Description     string `json:"description"`
	DescriptionEng  string `json:"descriptionEng"`
	Requirement     string `json:"requirement"`
	Caused          string `json:"caused"`
	Suggest         string `json:"suggest"`
	Reference       string `json:"reference"`
	Valid           int64  `json:"valid"`
	Cve             string `json:"cve"`
	DefaultAction   int64  `json:"defaultAction"`
	PublishDate     string `json:"publishDate"`
	Sid             int64  `json:"sid"`
	Deleted         int64  `json:"deleted"`
}
type WhiteListLog struct {
	Id          int64  `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
	Ip          string `json:"ip"`
	Mac         string `json:"mac"`
	Protocol    string `json:"protocol"`
	Port        string `json:"port"`
	Description string `json:"description"`
	WhiteListId int64  `json:"white_list_id"`
}
type WhiteListWarning struct {
	Id               int64  `json:"id"`
	SourceName       string `json:"source_name"`
	DestinationName  string `json:"destination_name"`
	SourceIp         string `json:"source_ip"`
	DestinationIp    string `json:"destination_ip"`
	SourceMac        string `json:"source_mac"`
	DestinationMac   string `json:"destination_mac"`
	IpVersion        int64  `json:"ip_version"`
	SourcePort       int64  `json:"source_port"`
	DestinationPort  int64  `json:"destination_port"`
	Action           int64  `json:"action"`
	Protocol         string `json:"protocol"`
	AppLayerProtocol string `json:"app_layer_protocol"`
	PacketLength     int64  `json:"packet_length"`
	Packet           string `json:"packet"`
	SignatureMessage string `json:"signature_message"`
	SignatureId      int64  `json:"signature_id"`
	MatchedKey       string `json:"matched_key"`
	ProtocolDetail   string `json:"protocol_detail"`
	PayloadLength    int64  `json:"payload_length"`
	Payload          string `json:"payload"`
	SignatureName    string `json:"signature_name"`
	Timestamp        string `json:"timestamp"`
	Level            int64  `json:"level"`
	Status           int64  `json:"status"`
	Deployed         int64  `json:"deployed"`
	AlertType        int64  `json:"alert_type"`
	DeviceId         string `json:"device_id"`
	DeviceName       string `json:"device_name"`
	RiskLevel        int64  `json:"risk_level"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
	Sid              string `json:"sid"`
	WarningNum       int64  `json:"warning_num"`
	HandleStatus     string `json:"handle_status"`
	HandleTime       string `json:"handle_time"`
	HandlePerson     string `json:"handle_person"`
	HandleContent    string `json:"handle_content"`
	WhitelistGroupId string `json:"whitelist_group_id"`
}
type WorkingModel struct {
	Id                  int64  `json:"id"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
	DeletedAt           string `json:"deleted_at"`
	WorkingModelStatus  string `json:"working_model_status"`
	StartTime           string `json:"start_time"`
	EndTime             string `json:"end_time"`
	StudyContent        string `json:"study_content"`
	AllowListGroup      string `json:"allow_list_group"`
	BaseLineGroup       string `json:"base_line_group"`
	Protos              string `json:"protos"`
	BaseLines           string `json:"base_lines"`
	RunState            string `json:"run_state"`
	WarningAllowList    string `json:"warning_allow_list"`
	WarningBaseLineList string `json:"warning_base_line_list"`
	WarningState        string `json:"warning_state"`
}
