package db

type YoshiStore struct {
	db map[string]string
	// will soon add mutex
}

func NewStorage() *YoshiStore {
	return &YoshiStore{
		db: make(map[string]string),
	}
}

func (ys *YoshiStore) Set(key, value string) {
	ys.db[key] = value
}

func (ys *YoshiStore) Get(key string) (string, bool) {
	val, ok := ys.db[key]
	if !ok {
		return "", false
	}
	return val, true
}
