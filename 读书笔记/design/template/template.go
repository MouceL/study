package template

type controller interface {
	start()
	stop()
	restart()
}

type manager struct {
	name string
	ctrl controller
}

// manager 拥有一个 接口类型的变量


