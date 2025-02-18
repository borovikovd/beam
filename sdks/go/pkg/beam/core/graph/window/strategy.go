// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package window contains window representation, windowing strategies and utilities.
package window

import "github.com/apache/beam/sdks/v2/go/pkg/beam/core/graph/window/trigger"

type AccumulationMode string

const (
	Unspecified  AccumulationMode = "AccumulationMode_UNSPECIFIED"
	Discarding   AccumulationMode = "AccumulationMode_DISCARDING"
	Accumulating AccumulationMode = "AccumulationMode_ACCUMULATING"
	Retracting   AccumulationMode = "AccumulationMode_RETRACTING"
)

// WindowingStrategy defines the types of windowing used in a pipeline and contains
// the data to support executing a windowing strategy.
type WindowingStrategy struct {
	Fn               *Fn
	Trigger          trigger.Trigger
	AccumulationMode AccumulationMode
	AllowedLateness  int // in milliseconds
}

func (ws *WindowingStrategy) Equals(o *WindowingStrategy) bool {
	return ws.Fn.Equals(o.Fn)
}

func (ws *WindowingStrategy) String() string {
	return ws.Fn.String()
}

// DefaultWindowingStrategy returns the default windowing strategy.
func DefaultWindowingStrategy() *WindowingStrategy {
	return &WindowingStrategy{Fn: NewGlobalWindows(), Trigger: trigger.Trigger{Kind: trigger.DefaultTrigger}, AccumulationMode: Discarding, AllowedLateness: 0}
}
