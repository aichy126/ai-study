package main

import "fmt"

//实现一个简单的洗衣机状态机

//状态机接口
type IFSMState interface {
	Enter()                             //进入状态
	Exit()                              //退出状态
	CheckTransition(action string) bool //状态转移检测
	Action() string                     //动作
}

//状态struct
type FSMState struct{}

//待机状态
type StandByState struct {
	action string
	FSMState
}

func NewStandByState() *StandByState {
	return &StandByState{action: "待机"}
}

//进入待机
func (this *StandByState) Enter() {
	fmt.Println("待机: 待机中")
}

//退出待机
func (this *StandByState) Exit() {
	fmt.Println("待机: 待机结束,开始执行动作")
}

//待机动作
func (this *StandByState) Action() string {
	return this.action
}

//清洗状态转移检测
func (this *StandByState) CheckTransition(action string) bool {
	if action == this.action {
		return true
	}
	return false
}

//清洗
type CleaningState struct {
	action string
	FSMState
}

func NewCleaningState() *CleaningState {
	return &CleaningState{action: "清洗"}
}

//开始清洗
func (this *CleaningState) Enter() {
	fmt.Println("清洗: 开始洗衣服")
}

//结束清洗
func (this *CleaningState) Exit() {
	fmt.Println("清洗: 洗衣服结束")
}

//清洗动作
func (this *CleaningState) Action() string {
	return this.action
}

//清洗状态转移检测
func (this *CleaningState) CheckTransition(action string) bool {
	if action == this.action {
		return true
	}
	return false
}

//甩干
type ShakeDryState struct {
	action string
	FSMState
}

func NewShakeDryState() *ShakeDryState {
	return &ShakeDryState{action: "甩干"}
}
func (this *ShakeDryState) Enter() {
	fmt.Println("甩干: 开始甩干")
}
func (this *ShakeDryState) Exit() {
	fmt.Println("甩干: 甩干结束")
}
func (this *ShakeDryState) Action() string {
	return this.action
}

// 状态转移检测
func (this *ShakeDryState) CheckTransition(action string) bool {
	if action == this.action {
		return true
	}
	return false
}

type FSM struct {
	// 持有状态集合
	states map[string]IFSMState
	// 当前状态
	current_state IFSMState
	// 默认状态
	default_state IFSMState
	// 外部输入数据
	input_data string
	// 是否初始化
	inited bool
}

// 初始化FSM
func (this *FSM) Init() {
	this.Reset()
}

// 添加状态到FSM
func (this *FSM) AddState(key string, state IFSMState) {
	if this.states == nil {
		this.states = make(map[string]IFSMState, 2)
	}
	this.states[key] = state
}

// 设置默认的State
// 为什么要设置默认的State?
func (this *FSM) SetDefaultState(state IFSMState) {
	this.default_state = state
}

// 转移状态
// 当状态转移的时候要对当前状态判断,是否执行对当前状态的改变
func (this *FSM) TransitionState() {
	//下一个状态改变为初始状态
	nextState := this.default_state
	//获得要执行的状态
	input_data := this.input_data
	if this.inited {
		//获取所有状态
		for _, v := range this.states {
			//如果输入状态在状态合集里
			if input_data == v.Action() {
				//下一个状态为要执行的状态
				nextState = v
				break
			}
		}
	}
	if ok := nextState.CheckTransition(this.input_data); ok {
		if this.current_state != nil {
			// 退出前一个状态
			this.current_state.Exit()
		}
		this.current_state = nextState
		this.inited = true
		nextState.Enter()
	}
}

// 设置输入数据
func (this *FSM) SetAction(inputData string) {
	this.input_data = inputData
	this.TransitionState()
}

// 重置
func (this *FSM) Reset() {
	this.inited = false
}

func main() {
	standByState := NewStandByState()   //待机
	cleaningState := NewCleaningState() //清洗
	shakeDryState := NewShakeDryState() //甩干
	fsm := new(FSM)
	fsm.AddState("待机", standByState)
	fsm.AddState("清洗", cleaningState)
	fsm.AddState("甩干", shakeDryState)
	fsm.SetDefaultState(standByState)
	fsm.Init()
	fsm.SetAction("清洗")
	fsm.SetAction("清洗")
	fsm.SetAction("甩干")
	fsm.SetAction("清洗")
	fsm.SetAction("甩干")

}
