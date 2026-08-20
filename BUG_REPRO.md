# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

同一站点两次巡检并发完成时，“最近巡检时间”偶尔会倒退到较早的一次，race 检测同时报告共享状态竞争。请只做诊断，暂时不要修改代码；找出未同步读写的文件和符号，解释旧结果覆盖新结果的时序，并给出确定性交错或 race 输出。

## 含 Bug 版本

- 仓库：11DingKing/chargeguard-bug-26
- 仓库地址：https://github.com/11DingKing/chargeguard-bug-26.git
- parent SHA：0bd9e503b633297cd46a6e28cecca4ff5e659ab0

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/chargeguard-bug-26.git bug-repro
cd bug-repro
git checkout --detach 0bd9e503b633297cd46a6e28cecca4ff5e659ab0
go test ./internal/httpapi -run TestTaskBehavior -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
    task_behavior_test.go:22: latest=2026-03-01 08:00:00 +0000 UTC
FAIL
FAIL	chargeguard/internal/httpapi	0.063s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
    task_behavior_test.go:22: latest=2026-03-01 08:00:00 +0000 UTC
FAIL
FAIL	chargeguard/internal/httpapi	0.003s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

根因结论必须准确写明出问题的 Go 文件、具体符号和完整失效机制，并由实际复现、源码调查和验证证据支撑；调查结束时目标仓库代码、测试和配置零改动。
