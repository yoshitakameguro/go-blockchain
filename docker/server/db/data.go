package db

func CreateInitialUserData() {
	for i := 0; i < 3; i++ {
		FakeUserWithWallet()
	}
}
