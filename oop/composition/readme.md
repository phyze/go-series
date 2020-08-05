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

จาก code  ตัวอย่างด้านบนแสดงถึงการทำ composition โดย อ้างถึง Dog และ Cat มาไว้ใน PetComposition และ PetComposition สามารถเข้าถึง type ของ object อื่นได้

```go
func main () {
  dog := Dog{Name: "Ivory"}
	cat := Cat{Name: "Ebony"}
	petCom := PetComposition {
		cat: cat,
		dog: dog,
	}
	fmt.Println("cat name's ", petCom.cat.Name)
	fmt.Println("dog name's ", petCom.dog.Name)
}
```

## ความแตกต่างของ composition กับ inheritance


### **Inheritance**

Inheritance มี relation แบบ is a (เป็น หรือ คือ) ในเวลาเดียวกัน
นั้นหมายความว่า ถ้า child object มีการสืบทอดจาก parent object เจ้า child object จะได้ความสามารถทุกอย่างของ parent object มา ตรงนี้แหละที่ทำให้ child object มีคุณลักษณะ สองอย่างพร้อมกัน

```java

class Paper {}

class Book extends Paper {}

```

class book สามาเป็นได้ทั้งหนังสือและกระดาษในเวลาเดียวกัน และ 2 class ต้องมีความเกี่ยวโยงซึ่งกันและกันด้วย หมายความว่า paper ต้องสามาเป็นหนังสือได้และ book ต้องสามาเป็นกระดาษได้

### **Composition**


