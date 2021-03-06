package httpd

import (
	"github.com/kataras/neffos"

	"github.com/jumpserver/koko/pkg/logger"
)

func OnELFinderConnect(c *neffos.NSConn, msg neffos.Message) error {
	logger.Infof("Request %s: web folder ws connect", c.Conn.ID())
	data := EmitSidMsg{Sid: c.Conn.ID()}
	c.Emit("data", neffos.Marshal(data))
	return nil
}

func OnELFinderDisconnect(c *neffos.NSConn, msg neffos.Message) error {
	logger.Infof("Request %s: web folder ws disconnect", c.Conn.ID())
	removeUserVolume(c.Conn.ID())
	return nil
}
