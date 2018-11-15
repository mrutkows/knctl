#!/bin/bash

# Copyright 2018 The Knative Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e -x -u

GOOS=darwin GOARCH=amd64 go build -o knctl-darwin-amd64 ./cmd/...
GOOS=linux GOARCH=amd64 go build -o knctl-linux-amd64 ./cmd/...
GOOS=windows GOARCH=amd64 go build -o knctl-windows-amd64 ./cmd/...

shasum -a 256 ./knctl-*-amd64
