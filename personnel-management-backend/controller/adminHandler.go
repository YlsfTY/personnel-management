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

// func GetUsers(ctx *gin.Context) {
// 	pageStr := ctx.Query("userPage")
// 	page, err := strconv.Atoi(pageStr)
// 	if err != nil {
// 		ulits.ResponseWithError(ctx, http.StatusBadRequest, "Failed to parse userPage")
// 		return
// 	}
// 	userList, err := dao.GetUserList(10, page)
// 	if err != nil {
// 		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "Failed to get user list")
// 		return
// 	}

// 	if len(userList) == 0 {
// 		ulits.ResponseWithData(ctx, http.StatusOK, ulits.ResponseData{
// 			Code: 200,
// 			Data: gin.H{
// 				"userList": gin.H{},
// 			},
// 			Msg: "Get user list successfully",
// 		})
// 		return
// 	}

// 	uNameArr := make([]string, len(userList))

// 	for i, u := range userList {
// 		uNameArr[i] = u.Name
// 	}

// 	ulits.ResponseWithData(ctx, http.StatusOK, ulits.ResponseData{
// 		Code: 200,
// 		Data: gin.H{
// 			"userList": uNameArr,
// 		},
// 		Msg: "Get user list successfully",
// 	})
// }

func IsAdmin(ctx *gin.Context) {
	ulits.ResponseWithData(ctx, http.StatusOK, ulits.ResponseData{
		Code: 200,
		Data: gin.H{
			"isAdmin": true,
		},
		Msg: "Get user list successfully",
	})
}

// func GetPersonnelList(ctx *gin.Context) {
// 	pageStr := ctx.Query("perPage")
// 	page, err := strconv.Atoi(pageStr)
// 	if err != nil {
// 		ulits.ResponseWithError(ctx, http.StatusBadRequest, "Failed to parse perPage")
// 		return
// 	}
// 	userList, err := dao.GetPerList(10, page)
// 	if err != nil {
// 		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "Failed to get user list")
// 		return
// 	}
// 	ulits.ResponseWithData(ctx, http.StatusOK, ulits.ResponseData{
// 		Code: 200,
// 		Data: gin.H{
// 			"userList": userList,
// 		},
// 		Msg: "Get user list successfully",
// 	})
// }

func AdminGetListData(ctx *gin.Context) {
	resList, err := dao.GetDataList()
	if err != nil {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "Failed to get data list")
		return
	}
	ulits.ResponseWithData(ctx, http.StatusOK, ulits.ResponseData{
		Code: 200,
		Data: gin.H{
			"dataList": resList,
		},
		Msg: "Get data list successfully",
	})

}
func AdminGetPersonnelData(ctx *gin.Context) {
	u := &dao.User{
		Name: ctx.Query("userName"),
	}
	if exist, err := u.SearchUserName(); err != nil {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if !exist {
		ulits.ResponseWithError(ctx, http.StatusBadRequest, "The user does not exist.")
		return
	}
	p := &dao.Personnel{}
	if err := p.GetPersonnel(u); err != nil {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "Failed to get personnel")
		return
	}
	ulits.ResponseWithData(ctx, http.StatusOK, ulits.ResponseData{
		Code: 200,
		Data: gin.H{
			"personnel": p,
		},
		Msg: "Get user data successfully",
	})
}

func AdminUploadImage(ctx *gin.Context) {
	u := &dao.User{
		Name: ctx.Query("userName"),
	}
	if exist, err := u.SearchUserName(); err != nil {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if !exist {
		ulits.ResponseWithError(ctx, http.StatusBadRequest, "The user does not exist.")
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

func AdminCreatePer(ctx *gin.Context) {
	u := &dao.User{
		Name: ctx.Query("userName"),
	}
	if exist, err := u.SearchUserName(); err != nil {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if !exist {
		ulits.ResponseWithError(ctx, http.StatusBadRequest, "The user does not exist.")
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
		if err := p.CreatePersonnel(u); err != nil {
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

func AdminGetImage(ctx *gin.Context) {
	u := &dao.User{
		Name: ctx.Query("userName"),
	}
	if exist, err := u.SearchUserName(); err != nil {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if !exist {
		ulits.ResponseWithError(ctx, http.StatusBadRequest, "The user does not exist.")
		return
	}

	cfg := config.Conf.Server
	baseURL := fmt.Sprintf("http://%s:%s", cfg.Host, cfg.Port)

	imagePath, _ := getImagePath(u.Name)

	if imagePath == "" {
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

func AdminDeletePersonnel(ctx *gin.Context) {
	u := &dao.User{
		Name: ctx.Query("userName"),
	}
	if exist, err := u.SearchUserName(); err != nil {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if !exist {
		ulits.ResponseWithError(ctx, http.StatusBadRequest, "The user does not exist.")
		return
	}
	if err := dao.DeletePersonnel(u); err != nil {
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, "Internal Server Error.")
		return
	}
	ulits.ResponseWithData(ctx, http.StatusOK, ulits.ResponseData{
		Code: 200,
		Data: nil,
		Msg:  "Delete user successfully",
	})
}
