package decorate

import "fmt"

type person struct {
	name string
}

func (p *person) GetName() string {
	//TODO implement me
	return p.name
}

func (p *person) Decorate() {
	//TODO implement me
	panic("implement me")
}

func (p *person) Wear() {
	fmt.Println("穿衣服")
}

type personMethod interface {
	Wear()
	GetName() string
}

type engineer struct {
	person
	skill string
}

func (e *engineer) Skill() string {
	return e.skill
}

func (e *engineer) Wear() {
	fmt.Printf("我是%s工程师 我叫%s 着装:\n", e.Skill(), e.name)
	e.person.Wear()
}

type teacher struct {
	person
	title string
}

func (t *teacher) Title() string {
	return t.title
}
func (t *teacher) Wear() {
	fmt.Printf("我是%s %s\n", t.name, t.title)
	t.person.Wear()
}
