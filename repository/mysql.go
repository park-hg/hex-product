package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"product_catalog/domain"
)

// db(mysql)과 트랜잭션을 수행하는 코드

type mysqlRepository struct {
	db *gorm.DB
}

func NewMysqlRepository() domain.Repository {
	db, err := gorm.Open(mysql.Open("root:1234@(127.0.0.1:3306)/products"), &gorm.Config{})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}
	repo := &mysqlRepository{db: db}

	return repo
}

func (m *mysqlRepository) Find(code string) (*domain.Product, error) {
	var product domain.Product
	m.db.Where("code = ?", code).First(&product)

	return &product, nil
}

func (m *mysqlRepository) Store(product *domain.Product) error {
	//TODO implement me
	p := domain.Product{
		Code:  product.Code,
		Name:  product.Name,
		Price: product.Price,
	}
	result := m.db.Create(&p)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *mysqlRepository) Update(product *domain.Product) error {
	m.db.Model(&domain.Product{}).Where("code = ?", product.Code).Updates(product)

	return nil
}

func (m *mysqlRepository) FindAll() ([]domain.Product, error) {
	var products []domain.Product
	result := m.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (m *mysqlRepository) Delete(code string) error {
	var product domain.Product
	m.db.Where("code = ?", code).Delete(&product)

	return nil
}
