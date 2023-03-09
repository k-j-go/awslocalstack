package person

type Person struct {
	// Personal details
	name, address, pin string
	// Job details
	workAddress, company, position string
	salary                         int
}

// PersonBuilder struct
type personBuilder struct {
	person *Person
}

// PersonAddressBuilder facet of PersonBuilder
type personAddressBuilder struct {
	personBuilder
}

// PersonJobBuilder facet of PersonBuilder
type personJobBuilder struct {
	personBuilder
}

// Builder NewPersonBuilder constructor for PersonBuilder
func Builder() *personBuilder {
	return &personBuilder{person: &Person{}}
}

// Lives chains to type *PersonBuilder and returns a *PersonAddressBuilder
func (b *personBuilder) Lives() *personAddressBuilder {
	return &personAddressBuilder{*b}
}

// Works chains to type *PersonBuilder and returns a *PersonJobBuilder
func (b *personBuilder) Works() *personJobBuilder {
	return &personJobBuilder{*b}
}

// At adds address to person
func (a *personAddressBuilder) At(address string) *personAddressBuilder {
	a.person.address = address
	return a
}

// WithPostalCode adds postal code to person
func (a *personAddressBuilder) WithPostalCode(pin string) *personAddressBuilder {
	a.person.pin = pin
	return a
}

// As adds position to person
func (j *personJobBuilder) As(position string) *personJobBuilder {
	j.person.position = position
	return j
}

// For adds company to person
func (j *personJobBuilder) For(company string) *personJobBuilder {
	j.person.company = company
	return j
}

// In adds company address to person
func (j *personJobBuilder) In(companyAddress string) *personJobBuilder {
	j.person.workAddress = companyAddress
	return j
}

// WithSalary adds salary to person
func (j *personJobBuilder) WithSalary(salary int) *personJobBuilder {
	j.person.salary = salary
	return j
}

// Build builds a person from PersonBuilder
func (b *personBuilder) Build() *Person {
	return b.person
}
