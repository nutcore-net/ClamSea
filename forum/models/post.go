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
 * Created Date: 2019-11-09 17:04:06 +08:00
 * Author: DTSDAO
 *
 * Last Modified: 2019-11-10 15:49:07 +08:00
 * Modified By: DTSDAO
 * ---------------------------------------------------------------
 */

package models

type Post struct {
	ID uint `json:"id"`

	// Reply tree
	To    uint `json:"to"`
	Root  uint `json:"root"`
	Depth uint `json:"depth"`

	//Statistics
	Votes int  `json:"votes"`
	Views uint `json:"views"`

	// Info
	Title   string `json:"title"`
	Content string `json:"content"`

	Tags   []string `json:"tags"`
	Author uint     `json:"Author"`

	// Time
	CreatedAt uint `json:"created_at"`
	UpdatedAt uint `json:"updated_at"`
}
