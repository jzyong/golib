package fsm

// State 状态
type State interface {
	Enter(entity interface{})                  //状态进入
	Update(entity interface{})                 //正常定时监测更新
	Exit(entity interface{})                   //状态退出
	HandleMessage(entity, message interface{}) //向当前的状态或全局状态发送消息
}
