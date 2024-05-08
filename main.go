package main

func main() {
	var err error
	err = ConfigInit()
	if err != nil {
		panic(err)
	}

}
