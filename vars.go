// Copyright 2018 The SQLer Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.
package main

import (
	"errors"
	"flag"
	"runtime"
)

var (
	flagDBDriver       = flag.String("driver", "mysql", "the sql driver to be used")
	flagDBDSN          = flag.String("dsn", "root:root@tcp(127.0.0.1)/test?multiStatements=true", "the data source name for the selected engine")
	flagAPIFile        = flag.String("config", "./config.example.hcl", "the config file(s) that contains your endpoints configs, it accepts comma seprated list of glob style pattern")
	flagRESTListenAddr = flag.String("rest", ":8025", "the http restful api listen address")
	flagRESPListenAddr = flag.String("resp", ":3678", "the resp (redis protocol) server listen address")
	flagWorkers        = flag.Int("workers", runtime.NumCPU(), "the maximum workers count")
	flagSQLSeparator   = flag.String("sep", `---\\--`, "multi sql query separator")
)

var (
	errNoMacroFound       = errors.New("Resource not found")
	errValidationError    = errors.New("Validation error")
	errAuthorizationError = errors.New("Authorization Error")
)

var (
	errStatusCodeMap = map[error]int{
		errNoMacroFound:       404,
		errValidationError:    422,
		errAuthorizationError: 401,
	}
)

var (
	macrosManager *Manager
)

const (
	sqlerVersion = "v2.2"
	sqlerBrand   = `
                _ 
    ___   __ _ | |  ___  _ __  
   / __| / _| || | / _ \| '__|   
   \__ \| (_| || ||  __/| |     
   |___/ \__, ||_| \___||_|    v2.2 edit by Mou
			|_| 
                       _             _                _     
     _ __ ___    __ _ | |__    ___  | |_   ___   ___ | |__  
    | '_ | _ \  / _| || '_ \  / _ \ | __| / _ \ / __|| '_ \ 
    | | | | | || (_| || |_) || (_) || |_ |  __/| (__ | | | |
    |_| |_| |_| \__,_||_.__/  \___/  \__| \___| \___||_| |_|

turn your SQL queries into safe valid RESTful apis.
	
	`
)
