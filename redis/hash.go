package redis

// HGet HGet
func (c *Client) HGet(key string, field string) (string, error) {

	keyCommon := c.client.HGet(key, field)
	if keyCommon == nil {
		return "", nil
	}

	return keyCommon.Val(), nil
}

// HSet HSet
func (c *Client) HSet(key, field string, vals interface{}) (bool, error) {

	ret := c.client.HSet(key, field, vals)

	err := ret.Err()
	if err != nil {
		return false, err
	}

	return ret.Val(), nil
}
