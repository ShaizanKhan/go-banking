package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "John Doe", City: "New York", Zipcode: "10001", DateofBirth: "1985-05-20", Status: "Active"},
		{Id: "1002", Name: "Jane Smith", City: "Los Angeles", Zipcode: "90001", DateofBirth: "1990-08-15", Status: "Inactive"},
		{Id: "1003", Name: "Emily Johnson", City: "Chicago", Zipcode: "60601", DateofBirth: "1979-12-10", Status: "Active"},
		{Id: "1004", Name: "Michael Brown", City: "Houston", Zipcode: "77001", DateofBirth: "1982-03-30", Status: "Active"},
		{Id: "1005", Name: "Emma Davis", City: "Phoenix", Zipcode: "85001", DateofBirth: "1995-11-25", Status: "Inactive"},
	}

	return CustomerRepositoryStub{customers: customers}
}
