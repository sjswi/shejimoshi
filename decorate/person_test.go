package decorate

import "testing"

func Test(t *testing.T) {
	tony := engineer{
		person: person{
			name: "Tony",
		},
		skill: "客户端",
	}
	c := &CasualPantDecorateor{clothesDecorate{
		personMethod: &tony,
		decorated:    nil,
	}}

	w := new(WhiteShirtDecorator)
	b := new(BeltDecorator)
	g := new(GlassesDecorator)
	k := new(KnittedSweaterDecorator)
	l := new(LeatherShoeDecorator)

	w.Decoratefun(c)
	b.Decoratefun(w)
	g.Decoratefun(b)
	k.Decoratefun(g)
	l.Decoratefun(k)
	//TODO 一个问题，go的继承机制导致只能将接口传入，因此输出只能是反方向输出，因为engineer无法传到最外层
	l.Decorate()
}
