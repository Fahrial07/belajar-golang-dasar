package helper

var version = "1.0.0"      //tidak bisa di akses dari luar package karena pakai huruf kecil semua
var Application = "Golang" //bisa di akses

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}
