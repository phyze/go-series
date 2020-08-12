# Abstraction

เป็นการแสดงรายละเอียดข้อมูลหรือการทำงานที่จำเป็นสำหรับการใช้งานและซ้อนสิ่งที่ไม่ได้ใช้ไว้ภายในเพื่อลดความซับซ้อนลงและนำไปใช้งานงานโดยไม่ตองสนใจว่าข้างหลังมันมีการทำงานหรือข้อมูลอะไรอยู่บ้าง 


    type Animal interface {
      Eat() 
      Sound() string
      MakeSound()
    }


    type abstractAnimalSound struct {}

    func (aa *abstractAnimalSound) MakeSound(type string) {

    }

    type Dog struct {
      abstractAnimalSound
    }

    func (d *Dog) Eat() { ... }

    func (d *Dog) Sound() string 

