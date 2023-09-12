# TCR Access Control (TAC)

Manage the Tencent Cloud TCR access control security policies.

FYI: <https://cloud.tencent.com/document/api/1141/53906>

## Usage

Requirements:
- Go: 1.20+
- OS: Linux/Unix

```bash
# Build
git clone https://github.com/cnrancher/tcr-access-control.git && cd tcr-access-control
go build .

# Show usage
./tcr-access-control -h

# Init config
./tcr-access-control init

# Get security policies
./tcr-access-control status
```

## LICENSE

Copyright 2023 [Rancher Labs, Inc](https://rancher.com).

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
