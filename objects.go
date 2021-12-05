package netdisco

import "net/http"

func (c *Client) ObjectDeviceByIP(ip string) (DeviceDetails, error) {
	var device DeviceDetails
	err := c.Do(http.MethodGet, "/api/v1/object/device/"+ip, nil, &device)
	if err != nil {
		return DeviceDetails{}, err
	}
	return device, nil
}
