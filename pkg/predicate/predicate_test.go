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

package predicate_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/event"

	"github.com/mars1024/controller-extensions/pkg/predicate"
)

var _ = Describe("Predicate", func() {
	Describe("IgnoreDeletionPredicate", func() {
		pred := &predicate.IgnoreDeletionPredicate{}

		It("should return true on create event", func() {
			Expect(pred.Create(event.CreateEvent{})).Should(BeTrue())
		})

		It("should return true on update event", func() {
			Expect(pred.Update(event.UpdateEvent{})).Should(BeTrue())
		})

		It("should return false on delete event", func() {
			Expect(pred.Delete(event.DeleteEvent{})).Should(BeFalse())
		})

		It(" should return true on generic event", func() {
			Expect(pred.Generic(event.GenericEvent{})).Should(BeTrue())
		})
	})
})
