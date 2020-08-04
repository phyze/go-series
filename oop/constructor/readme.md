# Constructor

ภาษา go ไม่มี constructor เหมือนในภาษา java หรือ C# แต่ go นั้นสามารถสร้าง function แทน construct ได้

    func New() *Box {
      return &Box{}
    }

    type Box {
      Width int
      Hight int
    }


