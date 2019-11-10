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
 * Created Date: 2019-11-09 16:58:19 +08:00
 * Author: DTSDAO
 *
 * Last Modified: 2019-11-10 15:48:37 +08:00
 * Modified By: DTSDAO
 * ---------------------------------------------------------------
 */

package models

/*
	UserAllows
	true: permit
	false: deny
*/

type UserAllows struct {
	Msg   bool `json:"msg"`
	Forum bool `json:"forum"`
}
