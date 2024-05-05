package lifecycle

import "github.com/fengjx/go-halo/hook"

type lifecycle struct {
	HookName string
	Order    int
}

var (
	// InterfaceAware 接口实现注入
	InterfaceAware = lifecycle{
		HookName: "interface-aware",
		Order:    10,
	}

	// PostProcessor 后置处理
	PostProcessor = lifecycle{
		HookName: "post-processor",
		Order:    100,
	}

	// InitData 数据初始化
	InitData = lifecycle{
		HookName: "init-data",
		Order:    1000,
	}
)

func DoHooks() {
	hook.DoHooks(InterfaceAware.HookName)
	hook.DoHooks(PostProcessor.HookName)
	hook.DoHooks(InitData.HookName)
}

func AddHook(l lifecycle, handler func(), opts ...hook.Option) {
	AddHookWithOrder(l, l.Order, handler, opts...)
}

func AddHookWithOrder(l lifecycle, order int, handler func(), opts ...hook.Option) {
	hook.AddHook(l.HookName, order, handler, opts...)
}
