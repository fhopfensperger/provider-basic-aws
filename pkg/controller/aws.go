/*
Copyright 2019 The Crossplane Authors.

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

package controller

import (
	"time"

	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/logging"

	"github.com/fhopfensperger/provider-basic-aws/pkg/controller/config"
	"github.com/fhopfensperger/provider-basic-aws/pkg/controller/dynamodb/table"
)

// Setup creates all AWS controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter, poll time.Duration) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, time.Duration) error{
		table.SetupTable,
	} {
		if err := setup(mgr, l, rl, poll); err != nil {
			return err
		}
	}

	return config.Setup(mgr, l, rl)
}
