# Method Overriding

เมื่อ go ใช้ composition (go ไม่มี inheritance) หากต้องการ implement signature method ใหม่สามารถที่จะทำได้ พวกเราเรียกว่า Method Overriding
ซึ่งหน้าตาของ method ที่ implement ใหม่จะมี signature เหมือน method เก่า

```go

type A struct{}

func (a *A) Println() {
	fmt.Println("A")
}

type B struct {
	A
}

func (b *B) Println() {
	fmt.Println("B")
}

func main() {
	b := B{}
	b.Println()
}
```

