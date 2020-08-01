# Custom type

เป็นการสร้าง type ใหม่ขึ้นมาเพื่อนิยามให้ชื่อ type มีความชัดเจนต่อการนำไปใช้งาน 


การใช้ custom type ร่วมกับการแปลง Celsius ไปเป็น Fahrenheit 

    type Celsius float
    type Fahrenheit float
    func ConvertCelsToFah(c Celsius) Fahrenheit {...}


ใน go นั้นไม่มี type date ที่ build-in มาให้ มีแต่ type time เท่านั้น หากจำเป็นต้องใช้ type date เราสามารถทำ custom type ได้

    type Date time.Time
    func (d Date) String() string {...}

    //using
    var date Date
    fmt.Println(date.String())