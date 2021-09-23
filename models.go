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
