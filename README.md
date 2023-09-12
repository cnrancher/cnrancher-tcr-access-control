# TCR Access Control (TAC)

Manage the Tencent Cloud TCR access control security policies.

FYI: <https://cloud.tencent.com/document/api/1141/53906>

## Usage

Requirements:
- Go: 1.20+
- OS: Linux/Unix

```bash
# Build & Install
git clone https://github.com/cnrancher/tcr-access-control.git && cd tcr-access-control
go build . && go install
```

1. Show usage:
    ```console
    $ tcr-access-control -h
    tcr-access-control is a tool for manage the
    Tencent Cloud TCR access security policies.
    ......
    ```

1. Init config (default config path is `$HOME/.tcr_access_control.yaml`):
    ```console
    $ tcr-access-control init
    17:00:00 [INFO] Start init config:
    Default language (zh-CN/en-US) (default: en-US):
    ......
    ```

1. Show existing security policies:
    ```console
    $ tcr-access-control status
    17:00:00 [INFO] External Endpoint Status: Opened
    17:00:00 [INFO] Security Policies:
    INDEX |        CIDR        |  Description
    ------+--------------------+--------------
        0 |      123.12.0.0/24 | Example Description
        1 |          127.0.0.1 | Example
    ------+--------------------+--------------
    ```

1. Add one IP address (CIDR block) to security policy:
    ```console
    $ tcr-access-control allow --ip "8.8.8.8" --description="TEST"
    18:00:00 [INFO] Successfully add "8.8.8.8" to security policy
    ```

1. Remove one IP address from security policy:
    ```console
    $ tcr-access-control remove --ip "8.8.8.8" --index=3
    Security policy index [3] version [14] CIDR [8.8.8.8] will be delete! Confirm [y/N]: y
    DOUBLE CONFIRM! [y/N]: y
    18:00:00 [INFO] Successfully remove "8.8.8.8" from security policy
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
