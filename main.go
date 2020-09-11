package main

import (
	"fmt"
	"runtime"

	"github.com/ChizhovVadim/CounterGo/engine"
	"github.com/ChizhovVadim/CounterGo/eval"
	"github.com/ChizhovVadim/CounterGo/uci"
)

/*
Counter Copyright (C) 2017-2020 Vadim Chizhov
This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
You should have received a copy of the GNU General Public License along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

const (
	name   = "Counter"
	author = "Vadim Chizhov"
)

var (
	versionName = "sciabarracom"
	buildDate   = "2020/10/10"
	gitRevision = "master"
)

// Main is the entry point of OpenWhisk
func Main(args map[string]interface{}) map[string]interface{} {
	fmt.Println(name,
		"VersionName", versionName,
		"BuildDate", buildDate,
		"GitRevision", gitRevision,
		"RuntimeVersion", runtime.Version())

	var engine = engine.NewEngine(func() engine.Evaluator {
		return eval.NewEvaluationService()
	})

	var protocol = &uci.Protocol{
		Name:    name,
		Author:  author,
		Version: versionName,
		Engine:  engine,
		Options: []uci.Option{
			&uci.IntOption{Name: "Hash", Min: 4, Max: 1 << 16, Value: &engine.Hash},
			&uci.IntOption{Name: "Threads", Min: 1, Max: runtime.NumCPU(), Value: &engine.Threads},
			&uci.BoolOption{Name: "ExperimentSettings", Value: &engine.ExperimentSettings},
		},
	}

	res := make(map[string]interface{})

	time := "1000"
	ok := true
	fen := ""

	if t, ok := args["time"].(string); ok {
		time = t
	}

	if fen, ok = args["fen"].(string); !ok {
		res["error"] = "fen is a required argument"
		return res
	}

	move, err := uci.Play(protocol, fen, time)
	if err != nil {
		res["error"] = err.Error()
		return res
	}

	res["move"] = move
	return map[string]interface{}{
		"body": res,
	}
}
