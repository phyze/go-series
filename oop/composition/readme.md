# Composition

ในภาษา go นั้นไม่มีการ inheritance เลยแนะนำกลไกลที่คล้ายกับ inheritance 


Composition  คือ object หนึ่ง ๆ ที่ต้องการ refer objects มาไว้ที่ ๆ เดียวเพื่อการใช้งานบางสิ่งบางอย่าง

```go
type PetComposition struct {
	cat Cat
	dog Dog
}

type Dog struct {
	Name string
}

type Cat struct {
  Name string
}
```

จาก code  ตัวอย่างด้านบนแสดงถึงการทำ composition โดย อ้างถึง Dog และ Cat มาไว้ใน PetComposition และ PetComposition สามารถเข้าถึง properties และ method ของ object อื่นได้