// persistence 通过依赖注入方式，实现领域对持久化存储的控制反转（IOC）
package persistence

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/lupguo/go-ddd-sample/domain/entity"
)

// UploadImgPersis 上传图片的持久化结构体
type UploadImgPersis struct {
	db *gorm.DB
}

// NewUploadImgPersis 创建上传图片DB存储实例
func NewUploadImgPersis(db *gorm.DB) *UploadImgPersis {
	return &UploadImgPersis{db}
}

// Save 保存一张上传图片
func (p *UploadImgPersis) Save(img *entity.UploadImg) (*entity.UploadImg, error) {
	err := p.db.Create(img).Error
	if err != nil {
		return nil, err
	}
	return img, nil
}

// Get 获取一张上传图片
func (p *UploadImgPersis) Get(id uint64) (*entity.UploadImg, error) {
	var img entity.UploadImg
	err := p.db.Where("id = ?", id).Take(&img).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("upload image not found")
	}
	if err != nil {
		return nil, err
	}
	return &img, nil
}

// GetAll 获取一组上传图片
func (p *UploadImgPersis) GetAll() ([]entity.UploadImg, error) {
	var imgs []entity.UploadImg
	err := p.db.Limit(50).Order("created_at desc").Find(&imgs).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("upload images not found")
	}
	if err != nil {
		return nil, err
	}
	return imgs, nil
}

// Delete 删除一张图片
func (p *UploadImgPersis) Delete(id uint64) error {
	var img entity.UploadImg
	err := p.db.Where("id = ?", id).Delete(&img).Error
	if err != nil {
		return err
	}
	return nil
}