# Abstraction

เป็นการแสดงรายละเอียดข้อมูลหรือการทำงานที่จำเป็นสำหรับการใช้งานและซ้อนสิ่งที่ไม่ได้ใช้ไว้ภายในเพื่อลดความซับซ้อนลงและนำไปใช้งานงานโดยไม่ตองสนใจว่าข้างหลังว่ามันมีการทำงานหรือข้อมูลอะไรอยู่บ้าง 

abstraction/animal

```go
    package animal 
    
    type IMammal interface {
      Eat() 
      Sound() string
      MakeSound(mammalType string) string
    }

    type Mammal struct {
      abstractMammalSound
    }

    type abstractMammalSound struct {}
    func (am *abstractMammalSound) MakeSound(mammalType string) string {
      switch {
        case mammalType == "dog" :
          return "hok hok hok !!!"
        case mammalType == "cat" :
          reuturn "meow meow meow !!!"
        default :
          return ""
      }
    }
```

   1. define IMammal เป็น interface ให้มี behavior คือ กิน, ส่งเสียง , และ ทำเสียง
  2. defind Mammal struct ขึ้นมาเพื่อนำไปใช้ต่อเพื่อซ้อนความซับซ้อนไว้ใน Mammal โดยภายใน struct Mammal นั้นมีการใช้ composite ด้วยการนำ  abstractMammalSound ซึ่งเป็น struct เข้ามาไว้ใน Mammal เพื่อซ้อนความซับซ้อนการทำงานไว้ภายใต้ Mammal 
  3. define abstractMammalSound เป็น struct และ implement logic ที่เรียกกับการทำเสียงของสัตว์แต่ละประเภท

abstraction/animal/pet

```go 
  package pet

  type IPet interface {
    animal.IMammal
  }

  func NewDog() IPet {
    return &dog{}
  }

  type dog struct {}
  
  func (d *dog) Sound() string {
    return d.MakeSound("dog")
  }

  func (d *dog) Eat() {}
```

1. define IPet เป็น interface และใช้การ composite ฝัง animal.IMammal  ไว้ภายใน IPet เพื่อเป็นการบอกว่า IPet จะต้องมี behavior เหมือนกับ IMammal นะ เช่น Eat , Sound และ MakeSound
2. ทำการสร้าง Constructor ให้กับ struct dog และ มี return type เป็น IPet ซึ่งเป็น interface คนที่ได้รับ instance ของ dog ไปจะสามารถมองเห็น method ที่มีการประกาศไว้ใน IPet เท่านั้น
3. สร้าง struct dog และ implement method ตาม specification ของ IPet คือ method Sound และ Eat ให้สังเกตที่ method Sound จะมีการเรียกใช้ MakeSound() ภายใต้ method Sound เพื่อเป็นการซ้อนความ complex และ สิ่งที่ไม่จำเป็นที่ end user จะต้องรู้
   
```go 
  package main
  func main() {
    pet := pet.NewDog()
    fmt.Println(dog.Sound())
  }

```

> หากใครอยากทดลองเล่นดูให้ทำการสร้าง Cat struct ดูนะครับโดยดู dog เป็นตัวอย่าง

## abstraction แตกต่างจาก encapsulation อย่างไร


### Abstraction 
  เป็นการ focus ที่ high-level design โดย object จะไม่สนใจรายละเอียดในสิ่งที่ไม่จำเป็นและซ้อนสิ่งที่มีความซับซ้อนไว้ภายใน ซึ่ง object จะสนใจแค่ว่า มีอะไรบ้างที่สามารถนำไปใช้แก้ไขปัญหาที่กำลัง focus อยู่​ ณ ตอนนี้

### Encapsulation 
  เป็นการ focus ในระดับ implement ว่าจะต้องเก็บรักษาข้อมูลนี้ไว้ในภายในหรือปล่อยให้ภายนอกสามารถเรียกใช้ได้โดยไม่สนใจว่าสิ่งที่กำลังทำอยู่นั้นมีความซับซ้อนมากน้อยเพียงใด