package redis

// LPush LPush
func (c *Client) LPush(key string, values ...interface{}) (bool, error) {

	ret := c.client.LPush(key, values)

	err := ret.Err()
	if err != nil {
		return false, err
	}

	if ret.Val() > 0 {
		return true, nil
	}

	return false, nil
}
