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

	c.Decorate()
	c.Wear()
}

type CasualPantDecorateor struct {
	clothesDecorate
}

func (c *CasualPantDecorateor) Decorate() {
	//子类重写
	fmt.Println("一条卡其色休闲裤")
	c.wear()
}

type BeltDecorator struct {
	clothesDecorate
}

func (b *BeltDecorator) Decorate() {
	//子类重写
	fmt.Println("一条银色针扣头的黑色腰带")
	b.decorated.Decorate()
}

type LeatherShoeDecorator struct {
	clothesDecorate
}

func (l *LeatherShoeDecorator) Decorate() {
	//子类重写
	fmt.Println("一双深色休闲皮鞋")
	l.decorated.Decorate()
}

type KnittedSweaterDecorator struct {
	clothesDecorate
}

func (k *KnittedSweaterDecorator) Decorate() {
	//子类重写
	fmt.Println("一件紫红色针织毛衣")
	k.decorated.Decorate()
}

type WhiteShirtDecorator struct {
	clothesDecorate
}

func (w *WhiteShirtDecorator) Decorate() {
	//子类重写
	fmt.Println("一件白色衬衫")
	w.decorated.Decorate()
}

type GlassesDecorator struct {
	clothesDecorate
}

func (g *GlassesDecorator) Decorate() {
	//子类重写
	fmt.Println("一副方形黑框眼睛")
	g.decorated.Decorate()
}
