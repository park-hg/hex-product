package domain

// 도메인 로직을 위한 interface인 Service와 Repository를 정의합니다.
// interface란, 외부의 adapter가 내부의 domain과 상호작용하기 위한 규격을 제공하는 port입니다.
// business logic과 infrastructure(또는 framework)를 분리하므로써, 더 유연한 설계가 가능합니다.
// 즉, business logic또는 설계 이후의 필요에 따라 infrastructure(framework)를 교체할 수 있게 됩니다.
// 또한 각 구성요소(component)들이 decoupling되어 있으므로, 테스트하기 쉽고 독립적으로 관리할 수 있습니다.

// 외부의 API와 상호작용하기 위한 interface(port)입니다. primary adapter를 위한 port입니다.
type Service interface {
	Find(code string) (*Product, error)
	Store(product *Product) error
	Update(product *Product) error
	FindAll() ([]Product, error)
	Delete(code string) error
}

// 외부의 database와 상호작용하기 위한 interface(port)입니다. secondary adapter를 위한 port입니다.
type Repository interface {
	Find(code string) (*Product, error)
	Store(product *Product) error
	Update(product *Product) error
	FindAll() ([]Product, error)
	Delete(code string) error
}

type service struct {
	productrepo Repository
}

func NewProductService(productrepo Repository) Service {
	return &service{productrepo: productrepo}
}

func (s *service) Find(code string) (*Product, error) {
	return s.productrepo.Find(code)
}

func (s *service) Store(product *Product) error {
	return s.productrepo.Store(product)
}

func (s *service) Update(product *Product) error {
	return s.productrepo.Update(product)
}

func (s *service) FindAll() ([]Product, error) {
	return s.productrepo.FindAll()
}

func (s *service) Delete(code string) error {
	return s.productrepo.Delete(code)
}
