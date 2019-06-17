package main

import(

	"fmt"
	
)

type Creature struct{

	Name string
	Real bool

}

type FlyingCreature struct{

	Creature
	WingSpan int
}

type Foo struct{
}

type Unicorn struct{
	Creature
}

type Fooer interface{
	Foo1()
	Foo2()
}

type Pterodactyl struct{
	FlyingCreature

}

type Door struct{
	Thickness int
	Color string
}

func (f Foo) Foo1(){
	fmt.Println("Foo1() here")
}

func (f Foo) Foo2(){
	fmt.Println("Foo2() here")
}

func (c Creature) Dump(){

	fmt.Printf("Name: %s, Real %t\n",  c.Name, c.Real)
}

func Dump(c *Creature){
	fmt.Printf("Name: %s, Real %t\n",  c.Name, c.Real)
}

func teste(){
	fmt.Println("teste")
}

func NewPterodactyl (wingSpan int) *Pterodactyl{
	pet := &Pterodactyl{
		FlyingCreature{
			Creature{
				"Pterodactyl",
				true,
			},
			wingSpan,
		},	
	}
	return pet

}

func main(){

	creature := Creature{
		Name:"Monster Diogo",
		Real: true,
	}
	unicorn := Unicorn{
		Creature{"Uni Joao", true, },
	}

	myFlying := FlyingCreature{
		Creature{"brazilian flying",false,},
		500,
	}
	/*
	pterodactilo := Pterodactyl{
		FlyingCreature{
			Creature{
				"Pterodactilo",
				true,
			},
			5,
		},
	}
	*/
	pet2 := NewPterodactyl(8)

	door := &Door{3, "red"}

	
	
}