# Function 

function ทั่วไปคงรู้กันอยู่แล้วแต่ผมจะพาทัว anonymous function และ higher order function ซึ่งใน go ทำได้

## 1. Anonymous function 

    hello := func() {
      fmt.Println("hello")
    }

    hello() 

    // hello

เป็นการประกาศ function จากนั้น assign ให้กับ variable ที่ชื่อว่า hello เป็นการ define ไว้ก่อนแต่ยังไม่ทำงานจนกว่าจะมีการเรียกใช้ หากต้องการนำไปใช้งานให้ใส่ () ต่อหลังชื่อ variable นั้น


## 2. pass arg to Anonymous function with return 

    result := func(a,b int) int {
      return a+b
    }(2,3)
    
    //or 

    plus := func(a,b int) int {
      return a+b
    }

    fmt.Println(result)
    fmt.Println(plus(3,2))

    // 5

ส่วนนี้จะเหมือน ตัวอย่างที่ 1 เลยแต่ต่างกันที่มี return เพิ่มเข้ามา

## 3. higher order function 

Higher order function หรือ  HOF เป็น function ที่มีการทำงานร่วมกับ function อื่นการที่จะเป็น HOF ได้นั้นจำเป็นค้องส่ง function เข้ามาหรือ return เป็น function 

    func  pick(alphabets []string) func(pick string) []string { ... } 

    alphabets := []string{"a","a","b","r","t","o","f"}
    picked := pick(alphabets)("a")
    fmt.Println(picked)

    // ["a","a"]

