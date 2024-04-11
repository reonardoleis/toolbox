package http

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	views_handlers "github.com/reonardoleis/views/internal/adapters/transport/http/handlers/views"
	views_domain "github.com/reonardoleis/views/internal/core/domain/views"
)

type Server struct {
	r *gin.Engine
}

func NewServer(viewsUsecase views_domain.ViewUsecase) *Server {
	r := gin.Default()

	viewsHandler := views_handlers.New(viewsUsecase)

	r.POST("/views", viewsHandler.AddView)

	return &Server{r: r}
}

func (s Server) Run() error {
	port := fmt.Sprintf(":%s", os.Getenv("TOOLBOX_PORT"))
	if port == ":" {
		port = ":3000"
	}

	return s.r.Run(port)
}