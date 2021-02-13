package routes

import (
	"hr_manajemen/controllers"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.Use(cors.Default())
	user := r.Group("/user")
	{
		user.GET("/getUser", controllers.GetUser2)
// 		user.GET("getUser/:nik", controllers.GetUserByNIK)
		user.POST("/insert", controllers.InsertUser)
		user.PUT("update/:nik", controllers.UpdateUser)
		user.DELETE("delete/:nik", controllers.DeleteUser)
		user.POST("/login", controllers.LoginUser)
		user.POST("/signup", controllers.SignUpUser)
	}
	pelamar := r.Group("/pelamar")
	{
		// pelamar.GET("/getPelamar", controllers.GetPelamar)
		pelamar.GET("/getPelamar/:nik", controllers.GetPelamarsByNIK)
		pelamar.DELETE("/delete/:nik", controllers.DeletePelamars)
		// pelamar.PUT("update/", controllers.UpdatePelamars)
		pelamar.POST("/insert", controllers.InsertPelamar)
		pelamar.POST("/upload", controllers.UploadFoto)
		pelamar.POST("/login", controllers.LoginPelamar)

	}
	perusahaan := r.Group("/perusahaan")
	{
		perusahaan.GET("/getPerusahaan/:id", controllers.GetPerusahaanByID)
		perusahaan.GET("/getPerusahaan", controllers.GetAllPerusahaans)
		perusahaan.POST("/insert", controllers.InsertPerusahaan)
		perusahaan.DELETE("/delete/:id", controllers.DeletePerusahaan)
		perusahaanhh.PUT("update/:id", controllers.UpdatePerusahaan)

	}
	return r
}
