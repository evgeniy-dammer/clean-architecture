package http

import (
	"fmt"

	"github.com/evgeniy-dammer/clean-architecture/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("HTTP_PORT", 8080)
	viper.SetDefault("HTTP_HOST", "127.0.0.1")
	viper.SetDefault("IS_PRODUCTION", "false")
}

type Delivery struct {
	options   Options
	ucContact usecase.Contact
	ucGroup   usecase.Group
	router    *gin.Engine
}

type Options struct{}

func New(ucContact usecase.Contact, ucGroup usecase.Group, options Options) *Delivery {
	d := &Delivery{
		ucContact: ucContact,
		ucGroup:   ucGroup,
	}

	d.SetOptions(options)

	d.router = d.initRouter()

	return d
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}

func (d *Delivery) Run() error {
	err := d.router.Run(fmt.Sprintf("%s:%d", viper.GetString("HTTP_HOST"), uint16(viper.GetUint("HTTP_PORT"))))

	return fmt.Errorf("unable to start listenning: %w", err)
}

func checkAuth(c *gin.Context) {
	c.Next()
}
