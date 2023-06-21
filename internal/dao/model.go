package dao

type Model interface {
	GetID() uint64
	TableName() string
	Validate() error
}
