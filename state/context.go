package main

import "fmt"

type contextMethod interface {
	AddState(s st)
	ChangeState(s st) error
	GetState() st
	SetStateInfo(info interface{})
	GetStateInfo() interface{}
}
type stateIsNone struct {
	info string
}

func (s *stateIsNone) Error() string {
	return s.info
}

type context struct {
	states    map[st]int
	cur_state st
	//该字段的属性可以引起状态的变化
	state_info interface{}
}

func (c *context) AddState(s st) {
	c.states[s] = 1
}

func (c *context) ChangeState(s st) error {
	if s == nil {
		return &stateIsNone{info: "该状态不存在"}
	} else {
		if c.cur_state == nil {
			fmt.Printf("初始化为%s\n", s.GetName())
		} else {
			fmt.Printf("由%s转为%s\n", c.cur_state.GetName(), s.GetName())
		}
		c.cur_state = s
		return nil
	}
}

func (c *context) GetState() st {
	return c.cur_state
}

func (c *context) SetStateInfo(info interface{}) {
	//TODO 由属性导致状态的变化,应该由具体的类进行重写
	c.state_info = info
	for key, _ := range c.states {
		if key.IsMatch(info) {
			err := c.ChangeState(key)
			if err == nil {
				return
			}
			fmt.Println(err)
		}
	}
}

func (c *context) GetStateInfo() interface{} {

	return c.state_info
}
