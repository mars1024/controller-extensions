/*
 Copyright 2021 Controller Extensions Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package log

import (
	"sync"

	"github.com/go-logr/logr"
	"k8s.io/klog/v2/klogr"
)

// Log is the base logger used by controller-extensions.
var Log logr.Logger

var onlySetOnce = sync.Once{}

func init() {
	// using klog as default logger
	Log = klogr.New()
}

// SetLogger sets a logger implementation globally, and this function will only work on first call.
func SetLogger(l logr.Logger) {
	onlySetOnce.Do(func() {
		Log = l
	})
}
