/*
	option模式
*/

package design



type User struct {
	Name string 
	Age int 
	Tags map[string]string
}


// 第一种实现方式
type UserOption func(*User)

func NewUser(opts ...UserOption) *User {
	user := &User{}

	for _, opt := range opts {
		opt(user)	
	}

	return user 
}


func WithName(name string) UserOption {
	return func(u *User) {
		u.Name = name
	}
}

func WithAge(age int) UserOption {
	return func(u *User) {
		u.Age = age	
	}
}


func WithTag(k, v string) UserOption {
	return func (u *User)  {
		if u.Tags == nil {
			u.Tags = make(map[string]string)
		}
		u.Tags[k] = v
	}
}

// 第二种实现方式，Option定义为一个接口
type FilterOption interface {
	Judge(*User) bool
}

type NameEqual struct {
	Name string
}


func (ne * NameEqual) Judge (u *User) bool {
	if ne.Name == u.Name {
		return true
	} else {
		return false
	}
}


type AgeBetween struct {
	FromAge int 
	ToAge int
}

func (ab *AgeBetween) Judge(u *User) bool {
	if u.Age >= ab.FromAge && u.Age <= ab.ToAge {
		return true
	} else {
		return false
	}
}


func UserQuery(filters ...FilterOption) []*User {
	users := []*User{}	// 模拟从数据里检索出了一批User
	filteredUsers := make([]*User, 0, len(users))

L: 
	for _, user := range users {
		for _, opt := range filters {
			if !opt.Judge(user) {	// 有一个条件不成立， 就过滤掉
				continue L
			}
		}
	}


	return filteredUsers
}


/**
********** gorm *************
Open(dialector Dialector, opts ...Option) (db *DB, err error)
type Option interface {
	Apple(*Config) error 
	AfterInitialize(*DB) error
}
**/