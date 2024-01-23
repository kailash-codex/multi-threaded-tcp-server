package main

func main() {
	listener, err := net.Listen("tcp", ":2212")
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("waiting for a client to connect")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
	}
}