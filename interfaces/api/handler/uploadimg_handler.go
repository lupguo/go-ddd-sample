package handler

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"github.com/lupguo/go-ddd-sample/application"
	"github.com/lupguo/go-ddd-sample/domain/entity"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

// UploadImgHandle 上传处理
func UploadImgHandle(c echo.Context) error {
	callback := c.QueryParam("callback")
	var content struct {
		Response  string    `json:"response"`
		Timestamp time.Time `json:"timestamp"`
		Random    int       `json:"random"`
	}
	content.Response = "Sent via JSONP"
	content.Timestamp = time.Now().UTC()
	content.Random = rand.Intn(1000)
	return c.JSONP(http.StatusOK, callback, &content)
}

// UploadImgHandler 图片上传接口层处理
type UploadImgHandler struct {
	uploadImgApp application.UploadImgAppIer
}

// NewUploadImgHandler 初始化一个图片上传接口
func NewUploadImgHandler(app application.UploadImgAppIer) *UploadImgHandler {
	return &UploadImgHandler{uploadImgApp: app}
}

func (h *UploadImgHandler) Save(c echo.Context) error {
	forms, err := c.MultipartForm()
	if err != nil {
		return err
	}
	var imgs []*entity.UploadImg
	for _, file := range forms.File["upload"] {
		fo, err := file.Open()
		if err != nil {
			continue
		}
		// file storage path
		_, err = os.Stat(os.Getenv("IMAGE_STORAGE"))
		if err != nil {
			if os.IsNotExist(err) {
				if err := os.MkdirAll(os.Getenv("IMAGE_STORAGE"), 0755); err != nil {
					return err
				}
			} else {
				return err
			}
		}
		// file save
		ext := path.Ext(file.Filename)
		tempFile, err := ioutil.TempFile(os.Getenv("IMAGE_STORAGE"), "img_*"+ext)
		if err != nil {
			return err
		}
		_, err = io.Copy(tempFile, fo)
		if err != nil {
			return err
		}
		// upload
		uploadImg := entity.UploadImg{
			Name:      file.Filename,
			Path:      tempFile.Name(),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}
		img, err := h.uploadImgApp.Save(&uploadImg)
		if err != nil {
			return err
		}
		imgs = append(imgs, img)
	}
	return c.JSON(http.StatusOK, imgs)
}

func (h *UploadImgHandler) Get(c echo.Context) error {
	strID := c.Param("id")
	if strID == "" {
		return errors.New("the input image ID is empty")
	}
	id, err := strconv.ParseUint(strID, 10, 0)
	if err != nil {
		return err
	}
	img, err := h.uploadImgApp.Get(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, img)
}

func (h *UploadImgHandler) GetAll(c echo.Context) error {
	imgs, err := h.uploadImgApp.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, imgs)
}

func (h *UploadImgHandler) Delete(c echo.Context) error {
	strID := c.Param("id")
	if strID == "" {
		return errors.New("the deleted image ID is empty")
	}
	id, err := strconv.ParseUint(strID, 10, 0)
	if err != nil {
		return err
	}
	err = h.uploadImgApp.Delete(id)
	if err != nil {
		return err
	}
	msg := fmt.Sprintf(`{"msg": "delete Imgage ID:%s success"`, strID)
	return c.JSON(http.StatusOK, msg)
}
