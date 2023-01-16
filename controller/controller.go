package controller

import (
	"liteserver/controller/v1/public"
	"liteserver/controller/v1/system"
)

var ControllerGroup = new(Controller)

type Controller struct {
	System system.SystemGroup
	Public public.PublicGroup
}
