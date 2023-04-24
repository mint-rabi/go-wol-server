package go_wol_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewServer(address string, port int) error {
	r := gin.Default()
	r.GET("/wol/:mac_address", func(c *gin.Context) {
		macAddress := c.Param("mac_address")
		err := Send(macAddress)
		if err != nil {
			c.String(http.StatusBadRequest, "Wake on Lan send failed. cause:%s", err)
			return
		}
		c.String(http.StatusOK, "Wake on Lan send success. MAC Address:%s", macAddress)
		return
	})
	err := r.Run(fmt.Sprintf("%s:%v", address, port))
	if err != nil {
		return err
	}
	return nil
}
