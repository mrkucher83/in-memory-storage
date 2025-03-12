package engine

type HashTable struct {
	data map[string]string // in-memory
}

func NewHashTable() *HashTable {
	return &HashTable{
		data: make(map[string]string),
	}
}

func (s *HashTable) Set(key, val string) {
	s.data[key] = val
}

func (s *HashTable) Get(key string) (string, bool) {
	val, ok := s.data[key]
	return val, ok
}

func (s *HashTable) Del(key string) {
	delete(s.data, key)
}
