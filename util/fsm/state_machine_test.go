package fsm

import (
	"fmt"
	"testing"
	"time"
)

// 测试默认状态机
func TestDefaultStateMachine(t *testing.T) {
	monster := &Monster{
		Name: "狼人",
	}
	stateMachine := &DefaultStateMachine[*Monster]{Owner: monster}
	monster.StateMachine = stateMachine

	monster.StateMachine.SetInitialState(IdleInstance)

	for true {
		monster.StateMachine.Update()
		time.Sleep(time.Second)
	}

}

// 怪物
type Monster struct {
	Name         string                 //名称
	StateMachine StateMachine[*Monster] //状态机
	IdleTime     int32                  //空闲时间
	PatrolTime   int32                  //巡逻时间
}

var IdleInstance = &Idle{}
var PatrolInstance = &Patrol{}
var AttackInstance = &Attack{}

// 空闲
type Idle struct {
}

func (idle *Idle) Enter(monster *Monster) {
	fmt.Printf("%v 进入空闲\r\n", monster.Name)
}

func (idle *Idle) Update(monster *Monster) {
	fmt.Printf("%v 空闲中...\r\n", monster.Name)
	monster.IdleTime++
	if monster.IdleTime > 3 {
		monster.StateMachine.ChangeState(PatrolInstance)
	}
}

func (idle *Idle) Exit(monster *Monster) {
	fmt.Printf("%v 退出空闲\n", monster.Name)
	monster.IdleTime = 0
}

func (idle *Idle) HandleMessage(monster *Monster, message interface{}) bool {
	return true
}

// 巡逻
type Patrol struct {
	EmptyState[*Monster]
}

func (patrol *Patrol) Enter(monster *Monster) {
	monster.PatrolTime = 0
	fmt.Printf("%v 进入巡逻\n", monster.Name)
}

func (patrol *Patrol) Update(monster *Monster) {
	fmt.Printf("%v 巡逻中...\n", monster.Name)
	monster.PatrolTime++
	if monster.PatrolTime > 5 {
		monster.StateMachine.ChangeState(AttackInstance)
	}
}

// 攻击
type Attack struct {
	EmptyState[*Monster]
}

func (patrol *Attack) Enter(monster *Monster) {
	fmt.Printf("%v 进入攻击，killNpc\n", monster.Name)
	monster.StateMachine.ChangeState(IdleInstance)
}
