package repository

import "github.com/lupguo/go-ddd-sample/domain/entity"

// UploadImgRepo 图片上传相关仓储接口，只要实现了该接口，则可以操作Domain领域实体
type UploadImgRepo interface {
	Save(*entity.UploadImg) (*entity.UploadImg, error)
	Get(uint64) (*entity.UploadImg, error)
	GetAll() ([]entity.UploadImg, error)
	Delete(uint64) error
}
