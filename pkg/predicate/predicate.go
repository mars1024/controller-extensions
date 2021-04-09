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

package predicate

import (
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	logf "github.com/mars1024/controller-extensions/pkg/log"
)

var log = logf.Log.WithName("predicate")

// IgnoreDeletionPredicate implements a default delete predicate function which always return false.
//
// This predicate will ignore the delete events, if finalizer is being used for delete protection on objects,
// this predicate is recommended because it will avoid the last invalid reconciliation triggered by delete event.
//
// Controller.Watch(
//		&source.Kind{Type: v1.MyCustomKind},
// 		&handler.EnqueueRequestForObject{},
//		predicate.IgnoreDeletionPredicate{},
//		)
type IgnoreDeletionPredicate struct {
	predicate.Funcs
}

func (IgnoreDeletionPredicate) Delete(e event.DeleteEvent) bool {
	return false
}
