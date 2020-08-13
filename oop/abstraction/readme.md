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


## abstraction แตกต่างจาก encapsulation อย่างไร


### Abstraction 
  เป็นการ focus ที่ high-level design โดย object จะไม่สนใจรายละเอียดในสิ่งที่ไม่จำเป็นและซ้อนสิ่งที่มีความซับซ้อนไว้ภายใน ซึ่ง object จะสนใจแค่ว่า มีอะไรบ้างที่สามารถนำไปใช้แก้ไขปัญหาที่กำลัง focus อยู่​ ณ ตอนนี้

### Encapsulation 
  เป็นการ focus ในระดับ implement ว่าจะต้องเก็บรักษาข้อมูลนี้ไว้ในภายในหรือปล่อยให้ภายนอกสามารถเรียกใช้ได้โดยไม่สนใจว่าสิ่งที่กำลังทำอยู่นั้นมีความซับซ้อนมากน้อยเพียงใด