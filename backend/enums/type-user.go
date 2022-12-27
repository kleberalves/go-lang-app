package enums

type TypeUser int

const (
	Salesman TypeUser = iota + 1
	Customer
	Sysadmin
)

func (t TypeUser) String() string {
	a := [...]string{"salesman", "customer", "sysadmin"}

	if int(t) > len(a) || int(t) < 1 {
		return ""
	} else {
		return a[t-1]
	}
}
func (t TypeUser) EnumIndex() int {
	return int(t)
}
