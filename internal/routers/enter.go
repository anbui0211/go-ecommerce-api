package routers

import (
	"ecommerce/internal/routers/manage"
	"ecommerce/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
