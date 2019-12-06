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
 * Created Date: 2019-12-06 21:21:12 +08:00
 * Author: DTSDAO
 *
 * Last Modified: 2019-12-06 21:48:36 +08:00
 * Modified By: DTSDAO
 * ---------------------------------------------------------------
 */

package global

import (
	"fmt"

	// Config
	"github.com/BurntSushi/toml"
	"path/filepath"

	// Database
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	Version = "V0.1"

	DB *gorm.DB

	Config = &config{
		Debug:  true,
		Addr:   ":9041",
		LogLoc: "logs/",
		Secret: "ClAmsEa",
		Database: &databaseConfig{
			Host:     "127.0.0.1",
			Port:     "5432",
			User:     "clam",
			DBName:   "clam",
			Password: "",
		},
		Modules: &modulesConfig{
			Vote:     true,
			Forum:    true,
			Festival: true,
		},
	}
)

func Setup() error {
	//
	// Load Config
	//
	//TODO Add logger support and report issue
	cfgpath, _ := filepath.Abs("./config.toml")
	//fmt.Printf("parse toml file once. filePath: %s\n", fpath)
	if _, err := toml.DecodeFile(cfgpath, &Config); err != nil {
		return err
	}

	//
	// Load DB
	//
	addr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		Config.Database.Host,
		Config.Database.Port,
		Config.Database.User,
		Config.Database.DBName,
		Config.Database.Password,
	)
	DB, err := gorm.Open("postgres", addr)
	if err != nil {
		//TODO log and exit
		DB.Close()
		return err
	}

	return nil
}
