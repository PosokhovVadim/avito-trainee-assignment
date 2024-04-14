package controllers

import "github.com/gin-gonic/gin"

var (
	BadRequest          = gin.H{"description": "Некорректные данные"}
	Unauthorized        = gin.H{"description": "Пользователь не авторизован"}
	AccessDenied        = gin.H{"description": "Пользователь не имеет доступа"}
	BannerNotFound      = gin.H{"description": "Баннер не найден"}
	InternalError       = gin.H{"description": "Внутренняя ошибка"}
	BannerAlreadyExists = gin.H{"description": "Баннер уже существует"}
)
