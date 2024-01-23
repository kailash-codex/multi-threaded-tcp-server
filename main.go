package main

func main() {
	listener, err := net.Listen("tcp", ":2212")
	if err != nil {
		log.Fatal(err)
	}
}