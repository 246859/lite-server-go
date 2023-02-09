package controller

import (
	"liteserver/controller/v1c/public"
	"liteserver/controller/v1c/system"
)

var ControllerGroup = new(Controller)

type Controller struct {
	System system.SystemGroup
	Public public.PublicGroup
}
