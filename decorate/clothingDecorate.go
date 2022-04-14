package decorate

import "fmt"

type decorate interface {
	Decorate()
}
type clothesDecorate struct {
	personMethod
	decorated decorate
}

func (c *clothesDecorate) Decorate() {
	//子类重写

}
func (c *clothesDecorate) Decoratefun(de decorate) {
	//子类重写

	c.decorated = de
}
func (c *clothesDecorate) wear() {
	//子类重写
	c.Wear()
	c.Decorate()

}

type CasualPantDecorateor struct {
	clothesDecorate
}

func (c *CasualPantDecorateor) Decorate() {
	//子类重写
	c.wear()
	fmt.Println("一条卡其色休闲裤")
}

type BeltDecorator struct {
	clothesDecorate
}

func (b *BeltDecorator) Decorate() {
	//子类重写
	b.decorated.Decorate()
	fmt.Println("一条银色针扣头的黑色腰带")

}

type LeatherShoeDecorator struct {
	clothesDecorate
}

func (l *LeatherShoeDecorator) Decorate() {
	//子类重写
	l.decorated.Decorate()
	fmt.Println("一双深色休闲皮鞋")

}

type KnittedSweaterDecorator struct {
	clothesDecorate
}

func (k *KnittedSweaterDecorator) Decorate() {
	//子类重写
	k.decorated.Decorate()
	fmt.Println("一件紫红色针织毛衣")

}

type WhiteShirtDecorator struct {
	clothesDecorate
}

func (w *WhiteShirtDecorator) Decorate() {
	//子类重写
	w.decorated.Decorate()
	fmt.Println("一件白色衬衫")
}

type GlassesDecorator struct {
	clothesDecorate
}

func (g *GlassesDecorator) Decorate() {
	//子类重写
	g.decorated.Decorate()
	fmt.Println("一副方形黑框眼睛")

}
