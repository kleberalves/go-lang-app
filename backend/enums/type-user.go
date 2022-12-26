package enums

type TypeUser int

const (
	Salesman TypeUser = iota + 1
	Customer
)

func (t TypeUser) String() string {
	return [...]string{"salesman", "customer"}[t-1]
}
func (t TypeUser) EnumIndex() int {
	return int(t)
}
