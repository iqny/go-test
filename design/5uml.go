package main

import (
	"fmt"
	"strconv"
)

// Water 水
type Water struct {
	weight int64
}
type Animal struct {
	name string
}

// Grow 成长
func (animal *Animal) Grow(water *Water) {
	weight := strconv.FormatInt(water.weight, 10)
	fmt.Printf("长大需要喝%s吨水\n", weight)
}

type Wing struct{}

// Bird 鸟类
type Bird struct {
	Animal
	featherCol string
	wing       *Wing
}

// LayEggs 下蛋
func (bird *Bird) LayEggs() {
	fmt.Println(bird.featherCol + "鸟下蛋了哈")
}

type Fly interface {
	fly()
}
type Goose struct {
	Bird
}

//fly 飞
func (goose *Goose) fly() {
	fmt.Println("大雁可以飞")
}

// Geese 鸭伪继鸟
type Geese struct {
	geese []*Goose
}
func (g *Geese) flyV() {
	fmt.Printf("雁群V形飞行，共有%d只大雁\n", len(g.geese))
}

type Duck struct {
	Bird
}
type Speak interface {
	Speak()
}
type DonaldDuck struct {
	Duck
}

// Speak 唐老鸭实现说话接口
func (donaldDuck *DonaldDuck) Speak() {
	fmt.Println("唐老鸭说人话了")
}

// Penguin 企鹅伪继承鸟，并关联天气
type Penguin struct {
	Bird
	climate *Climate
}

func (p *Penguin) LayEggs() {
	fmt.Printf("企鹅在%s的天气下下蛋了\n", p.climate.content)
}

type Climate struct {
	content string
}

func showUML() {
	water := &Water{weight: 10}
	//动物
	fmt.Println("-----------动物篇-依赖关系")
	animal := &Animal{}
	animal.Grow(water)
	//鸟
	fmt.Println("-----------鸟篇-伪继承/组合关系")
	bird := &Bird{
		featherCol: "五彩斑斓的",
		wing:       &Wing{},
	}
	bird.Grow(water)
	bird.LayEggs()
	//大雁
	fmt.Println("-----------大雁篇-实现关系")
	goose := &Goose{}
	goose.featherCol = "黑黑的"
	goose.Grow(water)
	goose.LayEggs()
	goose.fly()
	//鸭
	fmt.Println("-----------鸭篇-伪继承关系")
	duck := &Duck{}
	duck.featherCol = "黄色的"
	duck.Grow(water)
	duck.LayEggs()
	//企鹅
	fmt.Println("-----------企鹅篇-伪继承/关联关系")
	penguin := &Penguin{}
	penguin.featherCol = "白色的"
	penguin.Grow(water)
	climate := &Climate{content: "寒冷的"}
	penguin.climate = climate
	penguin.LayEggs()
	//雁群
	fmt.Println("-----------雁群篇-聚合关系")
	g := &Geese{}
	g.geese = append(g.geese, goose)
	g.flyV()
	//唐老鸭
	fmt.Println("-----------唐老鸭篇-实现关系")
	donaldDuck := &DonaldDuck{}
	donaldDuck.featherCol = "红色的"
	donaldDuck.Grow(water)
	donaldDuck.LayEggs()
	donaldDuck.Speak()
}
func main() {
	showUML()
}
