package data

func FillWithMockUsers(us *UserStore) {
	for id := 0; id < 100; id++ {
		us.Put(id)
	}
}
