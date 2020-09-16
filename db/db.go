package db

// NoSQL database interface
type NoSQL interface {
	UpdateDocument(collection, key string, data interface{}) int64
}
