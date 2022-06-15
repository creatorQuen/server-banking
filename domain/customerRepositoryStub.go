package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustommerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Anohin", "New Delhpi", "110011", "2000-01-01", "1"},
		{"1002", "Rob", "New Delhpi", "110011", "2000-01-01", "1"},
	}

	return CustomerRepositoryStub{customers}
}
