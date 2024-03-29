package controller

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"personnel-management-backend/config"
	"personnel-management-backend/dao"
	"personnel-management-backend/ulits"

	"github.com/gin-gonic/gin"
)

func CreatePer(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "User not found in request context")
		return
	}
	u, ok := user.(dao.User)
	if !ok {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "User type error")
		return
	}

	p := &dao.Personnel{}
	fmt.Println(ctx)
	if err := ctx.ShouldBindJSON(p); err != nil {
		ulits.ResponseWithError(ctx, http.StatusBadRequest, "Failed to parse JSON")
		return
	}

	if err := p.Validate(); err != nil {
		ulits.ResponseWithError(ctx, http.StatusBadRequest, "Failed to validate JSON")
		return
	}
	if ok, err := u.IsUserWithPer(); err != nil {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "Server exception")
		return
	} else if !ok {
		if err := p.CreatePersonnel(&u); err != nil {
			ulits.ResponseWithError(ctx, http.StatusInternalServerError, "Failed to add personnel")
			return
		}
		ulits.ResponseWithData(ctx, http.StatusCreated, ulits.ResponseData{
			Code: 201,
			Data: nil,
			Msg:  "Personnel created successfully",
		})
		return
	}
	if err := p.UpdatePersonnel(); err != nil {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "Failed to update personnel")
		return
	}
	ulits.ResponseWithData(ctx, http.StatusOK, ulits.ResponseData{
		Code: 200,
		Data: nil,
		Msg:  "Personnel update successfully",
	})
}

func GetPer(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "User not found in request context")
		return
	}
	u, ok := user.(dao.User)
	if !ok {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "User type error")
		return
	}
	p := &dao.Personnel{}
	if err := p.GetPersonnel(&u); err != nil {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "Failed to get personnel")
		return
	}
	ulits.ResponseWithData(ctx, http.StatusOK, ulits.ResponseData{
		Code: 200,
		Data: gin.H{
			"personnel": p,
		},
		Msg: "Get personnel successfully",
	})

}

func UploadImage(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "User not found in request context")
		return
	}
	u, ok := user.(dao.User)
	if !ok {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "User type error")
		return
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		ulits.ResponseWithError(ctx, http.StatusBadRequest, "Fail to upload")
		return
	}
	ext := filepath.Ext(file.Filename)
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
		ulits.ResponseWithError(ctx, http.StatusBadRequest, "Fail to upload")
		return
	}

	oldImagePath, err := getImagePath(u.Name)
	if err != nil {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "Fail to upload")
	}

	if oldImagePath != "" {
		if err := os.Remove(oldImagePath); err != nil {
			ulits.ResponseWithError(ctx, http.StatusInternalServerError, "Fail to upload")
			return
		}
	}

	if err := ctx.SaveUploadedFile(file, "./static/image/"+url.PathEscape(u.Name)+ext); err != nil {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "Fail to upload")
		return
	}
	ulits.ResponseWithData(ctx, http.StatusOK, ulits.ResponseData{
		Code: 200,
		Data: nil,
		Msg:  "Image uploaded successfully",
	})
}

func GetImage(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "User not found in request context")
		return
	}
	u, ok := user.(dao.User)
	if !ok {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "User type error")
		return
	}

	cfg := config.Conf.Server
	baseURL := fmt.Sprintf("http://%s:%s", cfg.Host, cfg.Port)

	imagePath, _ := getImagePath(u.Name)

	if imagePath == "" {
		// ulits.ResponseWithError(ctx, http.StatusNotFound, "Image not found")
		ulits.ResponseWithData(ctx, http.StatusOK, ulits.ResponseData{
			Code: 200,
			Data: gin.H{
				"image": fmt.Sprintf("%s%s", baseURL, "/static/image/default.png"),
			},
			Msg: "Get image successfully",
		})
		return
	}

	imageURL := fmt.Sprintf("%s%s", baseURL, imagePath[1:])

	ulits.ResponseWithData(ctx, http.StatusOK, ulits.ResponseData{
		Code: 200,
		Data: gin.H{
			"image": imageURL,
		},
		Msg: "Get image successfully",
	})

}

func getImagePath(uName string) (string, error) {
	filePath := fmt.Sprintf("./static/image/%s", url.PathEscape(uName))
	imagePath := ""

	exts := []string{".jpg", ".jpeg", ".png"}

	for _, ext := range exts {
		if _, err := os.Stat(filePath + ext); os.IsNotExist(err) {
			continue
		} else if err != nil {
			return "", err
		}
		imagePath = filePath + ext
		break
	}
	return imagePath, nil
}

