#### protoc-gen-enum
用于生成中文枚举的 plugin

##### Useage
install
```shell script
go install github.com/yrzs/protoc-gen-enum

```
proto
```protobuf
syntax = "proto3";

package enum;

import "patch/go.proto";
import "gogo/enum.proto";

option go_package = "./;enum";

// 状态
enum Status {
  option (go.enum) = {stringer_name: 'OriginString'}; //重新命名String Func
  Status_Unspecified = 0 [(gogo.enumvalue_cn) = 'Status_Unspecified'];
  Enable = 1 [(gogo.enumvalue_cn) = '启用'];
  Disable = 2 [(gogo.enumvalue_cn) = "禁用"];
  Delete = 3 [(gogo.enumvalue_cn) = "已删除"];
}

```
makefile

```makefile
.PHONY: enums
# generate enums
enums:
        protoc --proto_path=. \
           --proto_path=./third_party \
           --go-patch_out=plugin=go,paths=source_relative:. \
           --enum_out=plugin=protoc-gen-enum,paths=source_relative:. \
           $(ENUM_PROTO_FILES)
```