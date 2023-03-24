package service

type RedisService interface {
	Put(string, string) error
	Remove(string) error
	Get(string) (*string, error)
}
