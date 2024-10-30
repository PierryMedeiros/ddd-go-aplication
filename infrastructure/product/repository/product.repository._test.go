package repository_test

import (
    "testing"
    "desafio-ddd-go/domain/product/entity"
    "desafio-ddd-go/infrastructure/product/repository"
    "desafio-ddd-go/infrastructure/product/repository/orm"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/stretchr/testify/assert"
)

func setupTestDB() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    if err := db.AutoMigrate(&orm.ProductModel{}); err != nil {
        return nil, err
    }
    return db, nil
}

func TestProductRepository_Create(t *testing.T) {
    db, err := setupTestDB()
    assert.NoError(t, err)

    repo := repository.NewProductRepository(db)
    product, err := entity.NewProduct("1", "Test Product", 100.0)
    assert.NoError(t, err)

    err = repo.Create(product)
    assert.NoError(t, err)

    var productModel orm.ProductModel
    err = db.First(&productModel, "id = ?", product.GetID()).Error
    assert.NoError(t, err)
    assert.Equal(t, product.GetName(), productModel.Name)
    assert.Equal(t, product.GetPrice(), productModel.Price)
}

func TestProductRepository_Update(t *testing.T) {
    db, err := setupTestDB()
    assert.NoError(t, err)

    repo := repository.NewProductRepository(db)
    product, err := entity.NewProduct("1", "Test Product", 100.0)
    assert.NoError(t, err)
    _ = repo.Create(product)

    product.ChangeName("Updated Product")
    product.ChangePrice(200.0)
    err = repo.Update(product)
    assert.NoError(t, err)

    var productModel orm.ProductModel
    err = db.First(&productModel, "id = ?", product.GetID()).Error
    assert.NoError(t, err)
    assert.Equal(t, "Updated Product", productModel.Name)
    assert.Equal(t, 200.0, productModel.Price)
}

func TestProductRepository_Find(t *testing.T) {
    db, err := setupTestDB()
    assert.NoError(t, err)

    repo := repository.NewProductRepository(db)
    product, err := entity.NewProduct("1", "Test Product", 100.0)
    assert.NoError(t, err)
    _ = repo.Create(product)

    foundProduct, err := repo.Find(product.GetID())
    assert.NoError(t, err)
    assert.Equal(t, product.GetID(), foundProduct.GetID())
    assert.Equal(t, product.GetName(), foundProduct.GetName())
    assert.Equal(t, product.GetPrice(), foundProduct.GetPrice())
}

func TestProductRepository_FindAll(t *testing.T) {
    db, err := setupTestDB()
    assert.NoError(t, err)

    repo := repository.NewProductRepository(db)
    product1, err := entity.NewProduct("1", "Product 1", 100.0)
    assert.NoError(t, err)
    product2, err := entity.NewProduct("2", "Product 2", 200.0)
    assert.NoError(t, err)
    _ = repo.Create(product1)
    _ = repo.Create(product2)

    products, err := repo.FindAll()
    assert.NoError(t, err)
    assert.Len(t, products, 2)
}
