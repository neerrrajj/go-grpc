package main

func main() {
	hTTPServer := NewHttpServer(":8000")
	go hTTPServer.Run()
	grpcServer := NewGRPCServer(":9000")
	grpcServer.Run()
}
