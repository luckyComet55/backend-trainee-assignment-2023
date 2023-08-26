package user

func InitMockData() *UserMockDatabase {
	db := NewUserMockDatabase()
	users := []User{{1}, {5}, {7}, {12}, {2}, {700}, {212}, {151}, {3}}
	for _, u := range users {
		if err := db.CreateObject(u); err != nil {
			panic("something went wrong, check mock values!")
		}
	}
	return db
}
