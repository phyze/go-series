# Composition

ในภาษา go นั้นไม่มีการ inheritance เลยแนะนำกลไกลที่คล้ายกับ inheritance 


**Composition**  คือ object หนึ่ง ๆ ที่ต้องการ refer objects มาไว้ที่ ๆ เดียวเพื่อนำมาใช้งาน

```go
type PetComposition struct {
	Cat Cat
	Dog Dog
}

type Dog struct {
	Name string
}

type Cat struct {
  Name string
}
```

จาก code  ตัวอย่างด้านบนแสดงถึงการทำ composition โดย อ้างถึง Dog และ Cat มาไว้ใน PetComposition และ PetComposition สามารถเข้าถึง type ของ object อื่นได้

```go
func main () {
  dog := Dog{Name: "Ivory"}
	cat := Cat{Name: "Ebony"}
	petCom := PetComposition {
		Cat: cat,
		Dog: dog,
	}
	fmt.Println("cat name's ", petCom.Cat.Name)
	fmt.Println("dog name's ", petCom.Dog.Name)
}
```

## ความแตกต่างของ composition กับ inheritance


### **Inheritance**

**Inheritance** มี relation แบบ is  (เป็น หรือ คือ) ในเวลาเดียวกัน
นั้นหมายความว่า ถ้า child object มีการสืบทอดจาก parent object เจ้า child object จะได้ความสามารถทุกอย่างของ parent object มา ตรงนี้แหละที่ทำให้ child object มีคุณลักษณะ สองอย่างพร้อมกัน

```java

class Paper {}

class Book extends Paper {}

```

class book สามาเป็นได้ทั้งหนังสือและกระดาษในเวลาเดียวกัน และ 2 class ต้องมีความหมายที่เชื่อมโยงซึ่งกันและกัน หมายความว่า paper ต้องสามากลายเป็นหนังสือได้และ book ต้องสามาเป็นกระดาษได้

### **Composition**

**composition** มี relation แบบ has (มี หรือ ประกอบด้วย) นั้นหมายความว่าการที่จะเรียกใช้ object อื่นไม่จำเป็นต้องมีความเกี่ยวข้องหรืออาจจะมีความเกี่ยวข้องกันก็ได้ซึ่งต่างจาก inheritance ที่ต้องมีความหมายที่เชื่อมโยงกัน นี้จึงเป็นข้อดีของ composition ที่มีความยืดยุ่นกว่า inheritance และยังช่วยลด coupling ของ objects อีกด้วย

```go

type Car struct {
  Engine Engine
  Wheels  []Wheel
}

type Engine struct {}
type Wheel struct {}
```

