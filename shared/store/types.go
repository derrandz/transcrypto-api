package store

type (
	Memory interface {
		Get(key string) interface{}
		Set(key string, value interface{}) bool
		Del(key string) bool
	}

	Redis interface {
		Get(key string) interface{}
		Set(key string, value interface{}) bool
		Del(key string) bool
	}

	MongoDb interface {
		InsertOne(arg ...interface{}) bool
		FindOne(arg ...interface{}) bool
	}
)
