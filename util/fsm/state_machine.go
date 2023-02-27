package fsm

// StateMachine 处理实体状态移动或者处理接收到的消息
type StateMachine[E any] interface {
	Update()                                //定时监测更新状态
	ChangeState(state State[E])             //执行状态转移
	RevertToPreviousState() bool            //返回到之前状态 true 成功
	SetInitialState(state State[E])         //设置初始状态
	IsInState(state State[E]) bool          //当前状态是否在给定的状态
	HandleMessage(message interface{}) bool //处理消息
}

// DefaultStateMachine 默认状态机管理器
// 另外可实现堆栈状态机，可依次返回之前状态 ，如菜单列表
type DefaultStateMachine[E any] struct {
	Owner         E        //状态机拥有者
	CurrentState  State[E] //当前状态
	PreviousState State[E] //之前状态
	GlobalState   State[E] //全局状态，每个State Update都会调用
}

func (sm *DefaultStateMachine[E]) SetInitialState(state State[E]) {
	sm.PreviousState = nil
	sm.CurrentState = state
}

// Update 更新当前或全局状态
func (sm *DefaultStateMachine[E]) Update() {
	if sm.GlobalState != nil {
		sm.GlobalState.Update(sm.Owner)
	}
	if sm.CurrentState != nil {
		sm.CurrentState.Update(sm.Owner)
	}
}

func (sm *DefaultStateMachine[E]) ChangeState(newState State[E]) {
	sm.PreviousState = sm.CurrentState
	if sm.CurrentState != nil {
		sm.CurrentState.Exit(sm.Owner)
	}
	sm.CurrentState = newState
	if sm.CurrentState != nil {
		sm.CurrentState.Enter(sm.Owner)
	}
}

func (sm *DefaultStateMachine[E]) RevertToPreviousState() bool {
	if sm.PreviousState == nil {
		return false
	}
	sm.ChangeState(sm.PreviousState)
	return true
}

func (sm *DefaultStateMachine[E]) IsInState(state State[E]) bool {
	return sm.CurrentState == state
}

func (sm *DefaultStateMachine[E]) HandleMessage(message interface{}) bool {
	if sm.CurrentState != nil && sm.CurrentState.HandleMessage(sm.Owner, message) {
		return true
	}
	if sm.GlobalState != nil && sm.GlobalState.HandleMessage(sm.Owner, message) {
		return true
	}
	return false
}
