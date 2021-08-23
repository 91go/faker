package faker

import (
	"github.com/91go/faker/faker"
	"reflect"
	"testing"
)

var (
	fk = faker.NewGenerator()
)

func TestAge(t *testing.T) {

	age := fk.Age()
	output := reflect.TypeOf(age).Kind()
	if output != reflect.String {
		t.Errorf("Age generation is failed for")
	}
}

func TestAirport(t *testing.T) {

	iata := fk.AirportIATA()
	output := reflect.TypeOf(iata).Kind()
	if output != reflect.String {
		t.Errorf("IATA generation is failed for")
	}
}
