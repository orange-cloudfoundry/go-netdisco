package netdisco

import (
	"net/url"

	"github.com/cespare/xxhash/v2"
)

const SeparatorByte byte = 255






var separatorByteSlice = []byte{SeparatorByte}

type Device struct {
	UptimeAge         string  `json:"uptime_age"`
	Location          string  `json:"location"`
	SinceLastArpnip   float64 `json:"since_last_arpnip"`
	FirstSeenStamp    string  `json:"first_seen_stamp"`
	OsVer             string  `json:"os_ver"`
	Name              string  `json:"name"`
	LastArpnipStamp   string  `json:"last_arpnip_stamp"`
	Model             string  `json:"model"`
	SinceFirstSeen    float64 `json:"since_first_seen"`
	IP                string  `json:"ip"`
	Serial            string  `json:"serial"`
	SinceLastMacsuck  float64 `json:"since_last_macsuck"`
	DNS               string  `json:"dns"`
	SinceLastDiscover float64 `json:"since_last_discover"`
	LastMacsuckStamp  string  `json:"last_macsuck_stamp"`
	LastDiscoverStamp string  `json:"last_discover_stamp"`

	Contact string `json:"contact"`
	Alias   string `json:"alias"`
	Vendor  string `json:"vendor"`
}

type SearchDeviceQuery struct {
	Q           string `json:"q" yaml:"q"`
	Name        string `json:"name" yaml:"name"`
	Location    string `json:"location" yaml:"location"`
	DNS         string `json:"dns" yaml:"dns"`
	Ip          string `json:"ip" yaml:"ip"`
	Description string `json:"description" yaml:"description"`
	Mac         string `json:"mac" yaml:"mac"`
	Model       string `json:"model" yaml:"model"`
	OS          string `json:"os" yaml:"os"`
	OSVer       string `json:"os_ver" yaml:"os_ver"`
	Vendor      string `json:"vendor" yaml:"vendor"`
	Layers      string `json:"layers" yaml:"layers"`
	Matchall    bool   `json:"matchall" yaml:"matchall"`

	queryId uint64 `json:"-" yaml:"-"`
}

func (q *SearchDeviceQuery) Serialize() url.Values {
	values := make(url.Values)

	if q.Q != "" {
		values["q"] = []string{q.Q}
	}
	if q.Name != "" {
		values["name"] = []string{q.Name}
	}
	if q.Location != "" {
		values["location"] = []string{q.Location}
	}
	if q.DNS != "" {
		values["dns"] = []string{q.DNS}
	}
	if q.Ip != "" {
		values["ip"] = []string{q.Ip}
	}
	if q.Description != "" {
		values["description"] = []string{q.Description}
	}
	if q.Mac != "" {
		values["mac"] = []string{q.Mac}
	}
	if q.Model != "" {
		values["model"] = []string{q.Model}
	}
	if q.OS != "" {
		values["os"] = []string{q.OS}
	}
	if q.OSVer != "" {
		values["os_ver"] = []string{q.OSVer}
	}
	if q.Vendor != "" {
		values["vendor"] = []string{q.Vendor}
	}
	if q.Layers != "" {
		values["layers"] = []string{q.Layers}
	}
	if q.Matchall {
		values["matchall"] = []string{"true"}
	} else {
		values["matchall"] = []string{"false"}
	}
	return values
}

func (q *SearchDeviceQuery) Id() uint64 {
	if q.queryId != 0 {
		return q.queryId
	}
	xxh := xxhash.New()
	for key, val := range q.Serialize() {
		xxh.WriteString("$" + key + "$" + val[0])
		xxh.Write(separatorByteSlice)
	}
	q.queryId = xxh.Sum64()
	return q.queryId
}

type DeviceDetails struct {
	Model        string `json:"model"`
	Fan          string `json:"fan"`
	Vendor       string `json:"vendor"`
	Layers       string `json:"layers"`
	Uptime       int64  `json:"uptime"`
	LastDiscover string `json:"last_discover"`
	Creation     string `json:"creation"`
	OsVer        string `json:"os_ver"`
	Log          string `json:"log"`
	Slots        int    `json:"slots"`
	Description  string `json:"description"`
	IP           string `json:"ip"`
	VtpDomain    string `json:"vtp_domain"`
	ChassisID    string `json:"chassis_id"`
	Ps2Type      string `json:"ps2_type"`
	LastMacsuck  string `json:"last_macsuck"`
	SnmpComm     string `json:"snmp_comm"`
	Ps1Status    string `json:"ps1_status"`
	SnmpEngineid string `json:"snmp_engineid"`
	IsPseudo     int    `json:"is_pseudo"`
	Os           string `json:"os"`
	SnmpVer      int    `json:"snmp_ver"`
	Name         string `json:"name"`
	Ps2Status    string `json:"ps2_status"`
	DNS          string `json:"dns"`
	Location     string `json:"location"`
	Serial       string `json:"serial"`
	Ps1Type      string `json:"ps1_type"`
	SnmpClass    string `json:"snmp_class"`
	Contact      string `json:"contact"`
	LastArpnip   string `json:"last_arpnip"`
	Mac          string `json:"mac"`
}

type DevicePoeStatus struct {
	PoeDisabledPorts   int    `json:"poe_disabled_ports"`
	PoePowerCommitted  string `json:"poe_power_committed"`
	PoeCapablePorts    int    `json:"poe_capable_ports"`
	Name               string `json:"name"`
	Model              string `json:"model"`
	PoePoweredPorts    int    `json:"poe_powered_ports"`
	PoePowerDelivering string `json:"poe_power_delivering"`
	PoeErroredPorts    int    `json:"poe_errored_ports"`
	Location           string `json:"location"`
	DNS                string `json:"dns"`
	IP                 string `json:"ip"`
	Module             int    `json:"module"`
	Power              int    `json:"power"`
	Status             string `json:"status"`
}

type PortUtilization struct {
	IP            string `json:"ip"`
	PortsInUse    int    `json:"ports_in_use"`
	PortsShutdown int    `json:"ports_shutdown"`
	PortCount     int    `json:"port_count"`
	PortsFree     int    `json:"ports_free"`
	DNS           string `json:"dns"`
}

type NodeIPCount struct {
	Name    string `json:"name"`
	Vendor  string `json:"vendor"`
	Mac     string `json:"mac"`
	IPCount int    `json:"ip_count"`
	DNS     string `json:"dns"`
	Port    string `json:"port"`
	Switch  string `json:"switch"`
}

type PortHalfDuplex struct {
	IP     string `json:"ip"`
	Device Device `json:"device"`
	Duplex string `json:"duplex"`
	Port   string `json:"port"`
	Name   string `json:"name"`
}

type PortAdminDown struct {
	IP          string `json:"ip"`
	DNS         string `json:"dns"`
	Port        string `json:"port"`
	Description string `json:"description"`
	Name        string `json:"name"`
	UpAdmin     string `json:"up_admin"`
}

type PortMultiNodes struct {
	MacCount    int    `json:"mac_count"`
	Description string `json:"description"`
	Name        string `json:"name"`
	IP          string `json:"ip"`
	DNS         string `json:"dns"`
	Port        string `json:"port"`
}

type PortErrorDisabled struct {
	DNS    string `json:"dns"`
	Name   string `json:"name"`
	Port   string `json:"port"`
	Reason string `json:"reason"`
	IP     string `json:"ip"`
}

type PortVlanMismatch struct {
	RightDevice string `json:"right_device"`
	RightPort   string `json:"right_port"`
	LeftVlans   string `json:"left_vlans"`
	LeftPort    string `json:"left_port"`
	LeftDevice  string `json:"left_device"`
	RightVlans  string `json:"right_vlans"`
}
