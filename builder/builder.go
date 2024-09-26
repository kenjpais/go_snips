package builder

import "fmt"

// Person struct and builder
type Person struct {
	Address, PostCode, City string
	CompanyName, Position   string
	AnnualIncome            int
}

type PersonBuilder struct {
	person *Person
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (pab *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	pab.person.City = city
	return pab
}

func (pab *PersonAddressBuilder) At(Address string) *PersonAddressBuilder {
	pab.person.Address = Address
	return pab
}

func (pab *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	pab.person.PostCode = postcode
	return pab
}

func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}

func (pjb *PersonJobBuilder) At(company string) *PersonJobBuilder {
	pjb.person.CompanyName = company
	return pjb
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

// Address struct and builder
type Address struct {
	FullAddress, Area string
	Pincode, StreetNo int
}

type addressMod func(*Address)

type AddressBuilder struct {
	actions []addressMod
}

func (ab *AddressBuilder) AddFullAddress(FullAddress string) *AddressBuilder {
	ab.actions = append(ab.actions, func(a *Address) {
		a.FullAddress = FullAddress
	})
	return ab
}

func (ab *AddressBuilder) AddArea(Area string) *AddressBuilder {
	ab.actions = append(ab.actions, func(a *Address) {
		a.Area = Area
	})
	return ab
}

func (ab *AddressBuilder) AddPincode(Pincode int) *AddressBuilder {
	ab.actions = append(ab.actions, func(a *Address) {
		a.Pincode = Pincode
	})
	return ab
}

func (ab *AddressBuilder) AddStreetNo(streetNo int) *AddressBuilder {
	ab.actions = append(ab.actions, func(a *Address) {
		a.StreetNo = streetNo
	})
	return ab
}

func (ab *AddressBuilder) Build() *Address {
	ad := Address{}
	for _, a := range ab.actions {
		a(&ad)
	}
	return &ad
}

// TestBuilder to test the builder pattern
func TestBuilder() {
	// Testing PersonBuilder
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("123 London Road").
		In("London").
		WithPostcode("12122").
		Works().
		At("Fabrikan").
		AsA("Programmer").
		Earning(123123)
	person := pb.Build()
	fmt.Println("Person:", *person)

	// Testing AddressBuilder
	ab := AddressBuilder{}
	address := ab.
		AddStreetNo(12).
		AddFullAddress("456 Elm Street").
		AddArea("Downtown").
		AddPincode(78910).
		Build()
	fmt.Println("Address:", *address)
}
