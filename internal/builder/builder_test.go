package builder

import (
	"fmt"
	"github.com/k-j-go/awslocalstack/internal/builder/person"
	"testing"
)

func TestBuilder(t *testing.T) {
	pb := person.Builder()
	pb.Lives().
		At("").
		WithPostalCode("560102").
		Works().
		As("Software Engineer").
		For("IBM").
		In("Bangalore").
		WithSalary(150000)

	person := pb.Build()

	fmt.Println(person)
}
