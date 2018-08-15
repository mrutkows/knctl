/*
Copyright 2018 The Knative Authors

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

package cmd

import (
	"time"

	uitable "github.com/cppforlife/go-cli-ui/ui/table"
	"k8s.io/apimachinery/pkg/util/duration"
)

type ValueAge struct {
	T time.Time
}

var _ uitable.Value = ValueAge{}

func NewValueAge(t time.Time) ValueAge { return ValueAge{T: t} }

func (t ValueAge) String() string {
	if t.T.IsZero() {
		return ""
	}
	return duration.ShortHumanDuration(time.Now().Sub(t.T))
}

func (t ValueAge) Value() uitable.Value { return t }

func (t ValueAge) Compare(other uitable.Value) int {
	otherT := other.(ValueAge).T
	switch {
	case t.T.Equal(otherT):
		return 0
	case t.T.Before(otherT):
		return -1
	default:
		return 1
	}
}
