package design


import "testing"

func Test_UserQuery(t *testing.T) {
	UserQuery()
	UserQuery(&NameEqual{"jonathan"})
	UserQuery(&AgeBetween{18, 28})
	UserQuery(&NameEqual{"jonathan"}, &AgeBetween{18, 28})
}