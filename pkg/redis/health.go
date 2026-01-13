package redis

func (c *RedisClient) Close() error {
	return c.Client.Close()
}
