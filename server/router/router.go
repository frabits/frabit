/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package router

import (
	"fmt"
	"time"

	_ "github.com/frabits/frabit/server/service"

	"github.com/gin-gonic/gin"
)

type Router struct {
	gin.Engine
}

func (r *Router) Setup(g gin.RouterGroup) {
	fmt.Println(time.Now().Format(time.RFC3339))
}
