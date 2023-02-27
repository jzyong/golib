package fsm

// State 状态
type State[E any] interface {
	Enter(entity E)                                   //状态进入
	Update(entity E)                                  //正常定时监测更新
	Exit(entity E)                                    //状态退出
	HandleMessage(entity E, message interface{}) bool //向当前的状态或全局状态发送消息
}

// EmptyState 空状态，默认实现所有方法
type EmptyState[E any] struct {
}

func (state *EmptyState[E]) Enter(entity E) {
}

func (state *EmptyState[E]) Update(entity E) {
}

func (state *EmptyState[E]) Exit(entity E) {
}

func (state *EmptyState[E]) HandleMessage(entity E, message interface{}) bool {
	return false
}
