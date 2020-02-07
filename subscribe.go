package watchman

import "github.com/yookoala/realpath"

// Subscribe is the return object of the Subscribe call
type Subscribe struct {
	Clock     string `bser:"clock"`
	Subscribe string `bser:"subscribe"`
}

// Subscribe subscribes to changes against a specified root and requests that they be sent to the client via its connection. The updates will continue to be sent while the connection is open. If the connection is closed, the subscription is implicitly removed
// https://facebook.github.io/watchman/docs/cmd/subscribe.html
func (c *Client) Subscribe(path, name string) (*Subscribe, error) {
	path, err := realpath.Realpath(path)
	if err != nil {
		return nil, err
	}

	var data struct {
		Subscribe
		base
	}

	if err := c.Send(&data, "subscribe", path); err != nil {
		return nil, err
	}

	if data.Error != "" {
		return nil, data.Error
	}

	return &data.Subscribe, nil
}
