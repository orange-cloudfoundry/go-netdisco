package netdisco

import (
	"fmt"
	"net/http"
)

func (c *Client) ReportsDeviceAddrNoDns() ([]Device, error) {
	devices := make([]Device, 0)
	err := c.Do(http.MethodGet, "/api/v1/report/device/deviceaddrnodns", nil, &devices)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (c *Client) ReportsDeviceByLocation() ([]Device, error) {
	devices := make([]Device, 0)
	err := c.Do(http.MethodGet, "/api/v1/report/device/devicebylocation", nil, &devices)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (c *Client) ReportsDeviceDnsMismatch() ([]Device, error) {
	devices := make([]Device, 0)
	err := c.Do(http.MethodGet, "/api/v1/report/device/devicednsmismatch", nil, &devices)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (c *Client) ReportsDevicePoeStatus() ([]DevicePoeStatus, error) {
	devices := make([]DevicePoeStatus, 0)
	err := c.Do(http.MethodGet, "/api/v1/report/device/devicepoestatus", nil, &devices)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

type MarkAsFreeIfDownForRequest struct {
	AgeNum  int    `json:"age_num"`
	AgeUnit string `json:"age_unit"`
}

func (c *Client) ReportsDevicePortUtilization(req *MarkAsFreeIfDownForRequest) ([]PortUtilization, error) {
	if req == nil {
		req = &MarkAsFreeIfDownForRequest{
			AgeNum:  3,
			AgeUnit: "months",
		}
	}
	devices := make([]PortUtilization, 0)
	err := c.Do(http.MethodGet, fmt.Sprintf(
		"/api/v1/report/device/portutilization?age_num=%d&age_unit=%s", req.AgeNum, req.AgeUnit,
	), nil, &devices)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (c *Client) ReportsNodeMultiIps() ([]NodeIPCount, error) {
	nodes := make([]NodeIPCount, 0)
	err := c.Do(http.MethodGet, "/api/v1/report/node/nodemultiips", nil, &nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (c *Client) ReportsPortHalfDuplex() ([]PortHalfDuplex, error) {
	ports := make([]PortHalfDuplex, 0)
	err := c.Do(http.MethodGet, "/api/v1/report/port/halfduplex", nil, &ports)
	if err != nil {
		return nil, err
	}
	return ports, nil
}

func (c *Client) ReportsPortAdminDown() ([]PortAdminDown, error) {
	ports := make([]PortAdminDown, 0)
	err := c.Do(http.MethodGet, "/api/v1/report/port/portadmindown", nil, &ports)
	if err != nil {
		return nil, err
	}
	return ports, nil
}

func (c *Client) ReportsPortMultiNodes(filterByVlan int) ([]PortMultiNodes, error) {
	vlanParam := ""
	if filterByVlan > 0 {
		vlanParam = fmt.Sprintf("?vlan=%d", filterByVlan)
	}
	ports := make([]PortMultiNodes, 0)
	err := c.Do(http.MethodGet, "/api/v1/report/port/portmultinodes"+vlanParam, nil, &ports)
	if err != nil {
		return nil, err
	}
	return ports, nil
}

func (c *Client) ReportsPortErrorDisabled() ([]PortErrorDisabled, error) {
	ports := make([]PortErrorDisabled, 0)
	err := c.Do(http.MethodGet, "/api/v1/report/port/portserrordisabled", nil, &ports)
	if err != nil {
		return nil, err
	}
	return ports, nil
}

func (c *Client) ReportsPortVlanMismatch() ([]PortVlanMismatch, error) {
	ports := make([]PortVlanMismatch, 0)
	err := c.Do(http.MethodGet, "/api/v1/report/port/portvlanmismatch", nil, &ports)
	if err != nil {
		return nil, err
	}
	return ports, nil
}
