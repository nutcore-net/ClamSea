/*
 * Copyright 2019 Nutcore
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * ---------------------------------------------------------------
 * Project: ClamCloud
 *
 * Created Date: 2019-11-09 17:39:41 +08:00
 * Author: DTSDAO
 *
 * Last Modified: 2019-12-01 18:30:04 +08:00
 * Modified By: DTSDAO
 * ---------------------------------------------------------------
 */

package forum

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.RouterGroup) {
	//TODO Add real routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": 0,
			"data":   "hello",
		})
	})
}
