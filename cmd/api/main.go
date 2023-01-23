package main

func main() {
	server := InitializeServer()
	server.ListenAndServe()
}
