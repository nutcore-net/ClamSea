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
 * Created Date: 2019-12-01 17:21:35 +08:00
 * Author: DTSDAO
 *
 * Last Modified: 2019-12-01 18:26:28 +08:00
 * Modified By: DTSDAO
 * ---------------------------------------------------------------
 */

package shared

import (
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

type config struct {
	Debug    bool
	Addr     string
	LogLoc   string
	Secret   string
	Database databaseConfig
	Modules  modulesConfig
}

type databaseConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

type modulesConfig struct {
	Vote     bool
	Forum    bool
	Festival bool
}

var (
	Version = "V0.1"
	cfg     = &config{
		Debug:  true,
		Addr:   ":9041",
		LogLoc: "logs/",
		Secret: "ClAmsEa",
		Database: databaseConfig{
			Host:     "127.0.0.1",
			Port:     "5432",
			User:     "clam",
			DBName:   "clam",
			Password: "",
		},
		Modules: modulesConfig{
			Vote:     true,
			Forum:    true,
			Festival: true,
		},
	}
	once sync.Once
)

func Config() *config {
	once.Do(func() {
		//TODO Add logger support
		fpath, _ := filepath.Abs("./config.toml")
		//fmt.Printf("parse toml file once. filePath: %s\n", fpath)
		if _, err := toml.DecodeFile(fpath, &cfg); err != nil {
			panic(err)
		}
	})
	return cfg
}
