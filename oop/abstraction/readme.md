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
  เป็นการ focus ที่ high-level design โดย object จะไม่สนใจรายละเอียดหรือการทำงานที่มีการเจาะจงในบางอย่าง ซึ่ง object จะสนใจแค่ว่า มีการดำเนินการอะไรบ้าง มีข้อมูลอะไรบ้างที่ต้องอยู่ภายใต้ object