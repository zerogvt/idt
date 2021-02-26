package data

import "fmt"

func FillWithMockUsers(us *UserStore) {
	for id := 0; id < 100; id++ {
		us.Update(id, fmt.Sprintf("%06d", id))
	}
}
