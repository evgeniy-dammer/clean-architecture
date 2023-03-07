package http

import (
	"fmt"

	"github.com/evgeniy-dammer/clean-architecture/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// @title contact service on clean architecture
// @version 1.0
// @description contact service on clean architecture

// @contact.name API Support
// @contact.email evgeniydammer@gmail.com

// @BasePath /

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetDefault("HTTP_PORT", 8080)
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
	err := d.router.Run(fmt.Sprintf(":%d", uint16(viper.GetUint("HTTP_PORT"))))

	return errors.Wrap(err, "unable to run server")
}

func checkAuth(c *gin.Context) {
	c.Next()
}
