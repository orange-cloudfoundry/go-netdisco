package netdisco

import "net/http"

func (c *Client) SearchDevice(query *SearchDeviceQuery) ([]Device, error) {
	devices := make([]Device, 0)
	err := c.Do(http.MethodGet, "/api/v1/search/device?"+query.Serialize().Encode(), nil, &devices)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

