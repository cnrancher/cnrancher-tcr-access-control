# TCR Access Control (TAC)

Command line tool for managing the Tencent Cloud TCR ([容器镜像服务](https://cloud.tencent.com/document/product/1141/39278))
public access control (访问控制 -> 公网访问白名单) security policies.

FYI: <https://cloud.tencent.com/document/api/1141/53906>

## Usage

Requirements:
- Go: 1.20+
- OS: Linux/Unix

1. Build & Install
    ```console
    $ git clone https://github.com/cnrancher/tcr-access-control.git && cd tcr-access-control
    $ go build . && go install
    ```

1. Show usage:
    ```console
    $ tcr-access-control -h
    tcr-access-control is a tool for manage the Tencent Cloud
    TCR public access (访问控制 -> 公网访问白名单) security policies.
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

    You need to ensure that the External Endpoint (公网访问入口) Status is **Opened** before running this command.

    ```console
    $ tcr-access-control status
    External Endpoint (公网访问入口) Status: Opened

    Security Policies:
    ------+--------------------+--------------
    INDEX |        CIDR        |  Description
    ------+--------------------+--------------
        0 |      123.12.0.0/24 | Example Description
        1 |          127.0.0.1 | Example
    ------+--------------------+--------------
    ```

    Output in JSON format:

    ```console
    $ tcr-access-control status --json
    {
      "status": "Opened",
      "reason": "",
      "policies": [
        {
          "index": 0,
          "cidr": "1.2.3.4",
          "description": "Example"
        },
        {
          "index": 1,
          "cidr": "8.8.8.8/32",
          "description": "TEST"
        }
      ]
    }
    ```

1. Add one IP address (or CIDR block) to security policy:
    ```console
    $ tcr-access-control allow --ip "8.8.8.8" --description="TEST"
    18:00:00 [INFO] Successfully add "8.8.8.8" to security policy
    ```

1. Remove one IP address (or CIDR block) from security policy:
    ```console
    $ tcr-access-control remove --ip "8.8.8.8" --index=3
    Security policy index [3] version [14] CIDR [8.8.8.8] will be delete!
    Confirm [y/N]: y
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
