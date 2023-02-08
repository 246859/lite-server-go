package service

import (
	"liteserver/service/public"
	"liteserver/service/system"
)

var AppService = new(ServiceGroup)

type ServiceGroup struct {
	system.SystemService
	public.PublicService
}
