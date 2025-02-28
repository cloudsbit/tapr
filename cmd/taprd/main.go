// Copyright 2018 Klaus Birkelund Abildgaard Jensen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/cloudsbit/tapr/config"
	"github.com/cloudsbit/tapr/flags"
	"github.com/cloudsbit/tapr/rpc/ioserver"
	"github.com/cloudsbit/tapr/sim"
	"github.com/cloudsbit/tapr/store"

	// store implementations
	_ "github.com/cloudsbit/tapr/store/fs/service"
	_ "github.com/cloudsbit/tapr/store/tape/service"

	// inventory implementations
	_ "github.com/cloudsbit/tapr/store/tape/inv/postgres"

	// changer implementations
	_ "github.com/cloudsbit/tapr/store/tape/changer/fake"
	_ "github.com/cloudsbit/tapr/store/tape/changer/mtx"

	// format implementations
	_ "github.com/cloudsbit/tapr/format/ltfs"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flags.Parse(flags.Server)

	fmt.Println("taprd: starting")

	if flags.Simulate {
		fmt.Println("taprd: simulation enabled")
		sim.Enable()
	}

	fmt.Printf("taprd: server configuration file: %s\n", flags.ServerConfigFile)

	f, err := os.Open(flags.ServerConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	srvConfig, err := config.InitServerConfig(f)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Fprintf(os.Stderr, "SrvConfig: %v\n", config.PrettyPrint(srvConfig))

	for name, cfg := range srvConfig.Stores {
		// here name is 'default' or 'archive'
		stg, err := store.Create(name, cfg)
		if err != nil {
			log.Fatal(err)
		}

		// io api server
		httpIO := ioserver.New(config.New(), stg)
		http.Handle("/api/v1/"+name+"/io/", httpIO)
	}

	fmt.Println("taprd: server ready")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
