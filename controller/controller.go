package controller

import (
	"github.com/246859/lite-server-go/controller/v1c/public"
	"github.com/246859/lite-server-go/controller/v1c/system"
)

var ControllerGroup = new(Controller)

type Controller struct {
	System system.SystemGroup
	Public public.PublicGroup
}
