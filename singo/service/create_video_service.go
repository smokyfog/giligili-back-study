package service

import (
	"singo/model"
	"singo/serializer"
)

// CreateVideoService 视频投稿服务
type CreateVideoService struct {
	Title  string `form:"title" json:"title" binding:"required,min=2,max=300"`
	Info   string `form:"info" json:"info" binding:"required,min=0,max=3000"`
	URL    string `form:"url" json:"url"`
	Avatar string `form:"avatar" json:"avatar"`
}

// Create 创建视频
func (service *CreateVideoService) Create(user *model.User) serializer.Response {
	video := model.Video{
		Title:  service.Title,
		Info:   service.Info,
		URL:    service.URL,
		Avatar: service.Avatar,
		UserID: user.ID,
	}
	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  serializer.CodeParamErr,
			Msg:   "视频保存失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
