package main

func PanicFun() {
	defer func() {
		println("Excuted Panic Function")
	}()
	println("Excuting Panic Function")
	panic("Panic Situation Occurred")

}
func HelloWorld() {
	defer func() {
		println("Excuted Hellow World Function")
	}()
	println("Executing Hello Wrold Function")
	PanicFun()

}

func main() {
	println("Panic Example")
	defer func() {
		println("Excuted Main Function")
	}()
	HelloWorld()
}

/* Output:
@AnishriM ➜ /workspaces/guide-to-golang/intermediate-concepts (go-basic-syntax) $ go run panic.go
Panic Example
Executing Hello Wrold Function
Excuting Panic Function
Excuted Panic Function
Excuted Hellow World Function
Excuted Main Function
panic: Panic Situation Occurred

goroutine 1 [running]:
main.PanicFun()
        /workspaces/guide-to-golang/intermediate-concepts/panic.go:8 +0x65
main.HelloWorld()
        /workspaces/guide-to-golang/intermediate-concepts/panic.go:16 +0x56
main.main()
        /workspaces/guide-to-golang/intermediate-concepts/panic.go:25 +0x5b
exit status 2 */
