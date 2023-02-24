package fsm

// StateMachine 处理实体状态移动或者处理接收到的消息
type StateMachine interface {
	Update()                           //定时监测更新状态
	ChangeState(state State)           //执行状态转移
	RevertToPreviousState() bool       //返回到之前状态 true 成功
	SetInitialState(state State)       //设置初始状态
	SetGlobalState(state State)        //设置全局状态
	GetCurrentState() State            //获得当前状态
	GetGlobalState() State             //获得全局状态
	GetPreviousState() State           //获得之前状态
	IsInState(state State)             //当前状态是否在给定的状态
	HandleMessage(message interface{}) //处理消息
}

// DefaultStateMachine 默认状态机管理器
// 另外可实现堆栈状态机，可依次返回之前状态 ，如菜单列表
type DefaultStateMachine struct {
	Owner         interface{} //状态机拥有者
	CurrentState  State       //当前状态
	PreviousState State       //之前状态
	GlobalState   State       //全局状态，每个State Update都会调用
}

//TODO 实现方法
