package controllers

import "github.com/gin-gonic/gin"

var (
	BadRequest     = gin.H{"description": "Некорректные данные"}           // 400
	Unauthorized   = gin.H{"description": "Пользователь не авторизован"}   // 401
	AccessDenied   = gin.H{"description": "Пользователь не имеет доступа"} // 403
	BannerNotFound = gin.H{"description": "Баннер не найден"}              // 404
	InternalError  = gin.H{"description": "Внутренняя ошибка"}             // 500

)
