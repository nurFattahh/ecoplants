package repository

import (
	"ecoplant/entity"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return CartRepository{db}
}

func (r *CartRepository) GetProductByID(ID uint) (*entity.Product, error) {
	var product entity.Product

	result := r.db.First(&product, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (r *CartRepository) GetUserCartId(id uint) (*entity.User, error) {
	var user entity.User
	result := r.db.Where("id = ?", id).Take(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *CartRepository) AddProductToCart(id uint, input *entity.CartItem) error {
	var cart entity.CartItem
	err := r.db.Model(&cart).Where("cart_id = ?", id).Create(input).Error
	return err
}

func (r *CartRepository) GetAllProductInCart(ID uint) ([]entity.Cart, error) {
	var cart []entity.Cart
	err := r.db.Model(entity.Cart{}).Where("user_id", ID).Preload("Items.Product").Find(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (r *CartRepository) GetCheckListProduct(ID uint) ([]entity.CartItem, error) {
	var items []entity.CartItem
	err := r.db.Model(entity.CartItem{}).Where("is_check_list =? AND cart_id =?", 1, ID).Find(&items).Error
	return items, err
}

func (r *CartRepository) UpdateTotal(ID uint, total float64) error {
	err := r.db.Model(entity.Cart{}).Where("id =?", ID).Update("total", total).Error
	return err

}

func (r *CartRepository) GetAllCartItem(ID uint) ([]entity.CartItem, error) {
	var items []entity.CartItem
	err := r.db.Model(entity.CartItem{}).Where("cart_id =?", ID).Find(&items).Error
	return items, err
}

func (r *CartRepository) UpdateQuantity(ID uint) error {
	err := r.db.Model(entity.CartItem{}).Where("product_id =?", ID).Update("quantity", gorm.Expr("quantity + ?", 1)).Error
	return err
}

func (r *CartRepository) DeleteItemInCartByID(CartID uint, ProductID uint) error {
	var item entity.CartItem
	err := r.db.Where("product_id =? AND cart_id =? ", ProductID, CartID).Delete(&item).Error
	return err
}
