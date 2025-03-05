package storage

type Storage struct {
	data map[string]string // in-memory
}

func NewStorage() *Storage {
	return &Storage{
		data: make(map[string]string),
	}
}

func (s *Storage) Set(key, val string) {
	s.data[key] = val
}

func (s *Storage) Get(key string) (string, bool) {
	val, ok := s.data[key]
	return val, ok
}

func (s *Storage) Del(key string) {
	delete(s.data, key)
}
