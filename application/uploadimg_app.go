package application

import (
	"github.com/lupguo/go-ddd-sample/domain/entity"
	"github.com/lupguo/go-ddd-sample/domain/repository"
	"os"
)

// UploadImgAppIer 应用接口
type UploadImgAppIer interface {
	Save(*entity.UploadImg) (*entity.UploadImg, error)
	Get(uint64) (*entity.UploadImg, error)
	GetAll() ([]entity.UploadImg, error)
	Delete(uint64) error
}

// UploadImgApp 上传应用结构体，也是依赖领域仓储接口
type UploadImgApp struct {
	db repository.UploadImgRepo
}

// NewUploadImgApp 初始化上传图片应用
func NewUploadImgApp(db repository.UploadImgRepo) *UploadImgApp {
	return &UploadImgApp{db: db}
}

func (app *UploadImgApp) Save(img *entity.UploadImg) (*entity.UploadImg, error) {
	img, err := app.db.Save(img)
	if err != nil {
		return nil, err
	}
	img.Url = rawUrl(img.Path)
	return img, nil
}

func (app *UploadImgApp) Get(id uint64) (*entity.UploadImg, error) {
	img, err := app.db.Get(id)
	if err != nil {
		return nil, err
	}
	img.Url = rawUrl(img.Path)
	return img, nil
}

func (app *UploadImgApp) GetAll() ([]entity.UploadImg, error) {
	imgs, err := app.db.GetAll()
	if err != nil {
		return nil, err
	}
	for i, img := range imgs {
		imgs[i].Url = rawUrl(img.Path)
	}
	return imgs, nil
}

func (app *UploadImgApp) Delete(id uint64) error {
	return app.db.Delete(id)
}

func rawUrl(path string) string {
	return os.Getenv("IMAGE_DOMAIN") + os.Getenv("LISTEN_PORT") + path
}