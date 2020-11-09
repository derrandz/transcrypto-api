package memory

type (
	memory struct {
		db map[string]interface{}
	}
)

func NewMemoryStore() *memory {
	return &memory{
		db: make(map[string]interface{}),
	}
}

func (p *memory) Get(key string) interface{} {
	return p.db[key]
}

func (p *memory) Set(key string, value interface{}) bool {
	p.db[key] = value
	return true
}

func (p *memory) Del(key string) bool {
	delete(p.db, key)
	return true
}
