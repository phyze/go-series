package animal

type IMammal interface {
	Eat()
	Sound() string
	MakeSound(typ string) string
}

type Mammal struct {
	abstractMammalSound
}

type abstractMammalSound struct {}

func (aa *abstractMammalSound) MakeSound(typ string) string {
	var sound string
	switch  {
	case typ == "pet" :
		sound = "Hok hok hok !!!"
	case typ == "cat" :
		sound = "meow meow meow !!!"
	default:
		sound = ""
	}
	return sound
}

