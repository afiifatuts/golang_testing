package golangtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// test double with fake
type FakeEngine struct{}

func (e FakeEngine) MaxSpeed() int {
	return 5
}

func TestDefaultEngine_MaxSpeed(t *testing.T) {
	test := []struct {
		name     string
		expexted int
	}{
		{
			name:     "must have 50",
			expexted: 50,
		},
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			d := DefaultEngine{}
			ex := d.MaxSpeed()
			assert.Equal(t, ex, tc.expexted)
		})
	}
}
func TestTurboEngine_MaxSpeed(t *testing.T) {
	test := []struct {
		name     string
		expexted int
	}{
		{
			name:     "must have 500",
			expexted: 500,
		},
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			d := TurboEngine{}
			ex := d.MaxSpeed()
			assert.Equal(t, ex, tc.expexted)
		})
	}
}

func TestCar_Speed(t *testing.T) {
	type fields struct {
		Speeder Speeder
	}
	testCases := []struct {
		name     string
		field    fields
		expexted int
	}{
		{
			name:     "must be 50 when use default engine",
			field:    fields{Speeder: &DefaultEngine{}}, //-> pass pointer biar dia tau yg di tunjuk default engine
			expexted: 50,
		},
		{
			name:     "must be 80 when use turbo engine",
			field:    fields{Speeder: &TurboEngine{}}, //-> pass pointer biar dia tau yg di tunjuk default engine
			expexted: 80,
		},
		{
			name:     "must be 10 when use turbo engine",
			field:    fields{Speeder: &FakeEngine{}}, //-> pass pointer biar dia tau yg di tunjuk default engine
			expexted: 20,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d := Car{
				Speeder: tc.field.Speeder,
			}
			e := d.Speed()
			assert.Equal(t, e, tc.expexted)
		})
	}
}

// Testing dengan Mock
type MockEngine struct {
	mock.Mock
}

func (m MockEngine) MaxSpeed() int {
	args := m.Called()       //-> ketika mock dipanggil
	return args.Get(0).(int) //-> ambil index ke 0 berupa integer
}

func TestCar_Speed_WithMock(t *testing.T) {
	mock := new(MockEngine)
	car := Car{
		Speeder: mock,
	}

	mock.On("MaxSpeed").Return(9)

	speed := car.Speed()
	assert.Equal(t, 20, speed)
}
