#!/bin/bash
#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

set -x
set -e

OPERATOR_VERSION=999.0.0

for arg in "$@"; do
   case "$arg" in
      OPERATOR_VERSION=*) OPERATOR_VERSION="${arg#*=}" ;;
   esac
done

cd /home/kogito/

mv sonataflow-db-migrator-runner.jar sonataflow-db-migrator-"$OPERATOR_VERSION"-runner.jar

java -jar sonataflow-db-migrator-"$OPERATOR_VERSION"-runner.jar