package dilawar

type Type string

const (
	TypeCredit Type = "credit"
	TypeDebit  Type = "debit"
)

type Storable interface {
	Create(*Transaction) error
	Read(int) *Transaction
	List() []Transaction
}
