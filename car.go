package golangtest

type Speeder interface {
	MaxSpeed() int
}

func NewCar(speeder Speeder) *Car {
	return &Car{
		Speeder: speeder,
	}
}
//--------------------component 3
type Car struct {
	Speeder Speeder
}

func (c Car) Speed() int {
	defaultSpeed := 80
	if c.Speeder.MaxSpeed() <10{ //butuh test double
		return 20
	}
	if defaultSpeed > c.Speeder.MaxSpeed() {
		return c.Speeder.MaxSpeed()
	}
	return defaultSpeed
}
//---------------------component 2
type DefaultEngine struct{}

func (e DefaultEngine) MaxSpeed() int {
	return 50
}
//---------------------component 1
type TurboEngine struct{}

func (e TurboEngine) MaxSpeed() int {
	return 500
}
