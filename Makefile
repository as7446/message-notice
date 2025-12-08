##

BIN_DIR := bin
CMDS := dingtalk mattermost feishu wecom
GO ?= go

.PHONY: build-all $(CMDS) test clean

# 一次性构建全部 CLI，产物放在 bin/ 目录
build-all: $(addprefix build-,$(CMDS))

# 按需构建单个 CLI，例如：make build-feishu
build-%:
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/$* ./cmd/$*
	@echo "生成 $(BIN_DIR)/$* 完成"

# 运行全部测试
test:
	$(GO) test ./...

# 清理产物
clean:
	rm -rf $(BIN_DIR)
	@echo "已清理 bin/ 目录"

