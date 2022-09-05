package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	// UpdateTransaction(transaction models.Transaction) (models.Transaction, error)
	// DeleteTransaction(transaction models.Transaction) (models.Transaction, error)
	// GetUserTransaction(ID int) ([]models.User, error)
	UpdateTransactions(status string, ID string) error
	GetOneTransaction(ID string) (models.Transaction, error) // Declare GetOneTransaction repository method ...
	// InputIdTransToCart(product models.Cart) (models.Cart, error)
	FindTransactionId(UserID int) ([]models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Carts").Preload("Carts.Product").Find(&transaction, ID).Error
	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}

func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Carts").Preload("Carts.Product").Preload("User").First(&transaction, "id = ?", ID).Error

	return transaction, err
}

func (r *repository) UpdateTransactions(status string, ID string) error {
	var transaction models.Transaction
	r.db.Preload("Product").First(&transaction, ID)

	// If is different & Status is "success" decrement product quantity
	if status != transaction.Status && status == "success" {
		var product models.Product
		r.db.First(&product, transaction.ID)
	}

	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return err
}

func (r *repository) FindTransactionId(UserID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("User").Preload("Carts").Preload("Carts.Product").Find(&transaction, "user_id = ?", UserID).Error

	return transaction, err
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Preload("Carts").Preload("Carts.Product").Find(&transactions).Error
	return transactions, err
}
