package util

func MustOK(err error) {
	if err != nil {
		panic(err)
	}
}
