package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

func task1() {
	var x int = 42
	var f float32 = 25.5
	var b bool = true
	var s string = "hello"
	var e []int

	px := &x
	pf := &f
	pb := &b
	ps := &s
	pe := &e

	fmt.Println("Значение x:", x)
	fmt.Println("Значение f:", f)
	fmt.Println("Значение b:", b)
	fmt.Println("Значение s:", s)
	fmt.Println("Значение e:", e)

	fmt.Printf("Тип px: %T\n", px)
	fmt.Printf("Тип pf: %T\n", pf)
	fmt.Printf("Тип pb: %T\n", pb)
	fmt.Printf("Тип ps: %T\n", ps)
	fmt.Printf("Тип pe: %T\n", pe)
}

func task2() {
	strNum := "23"
	intNum, err := strconv.Atoi(strNum)
	if err != nil {
		fmt.Println("Ошибка при преобразовании:", err)
		return
	}
	fmt.Println("Целое число:", intNum)

	foo := "23abc"
	num := ""
	for _, r := range foo {
		if unicode.IsDigit(r) {
			num += string(r)
		} else {
			break
		}
	}
	n, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Ошибка при преобразовании:", err)
		return
	}
	fmt.Println("Целое число:", n)

	// Преобразуйте переменную age в Boolean
	age := "123abc"
	bul1 := age != ""
	fmt.Println("age в Boolean:", bul1)

	// Преобразуйте переменную flag в Boolean
	flag := 1
	bul2 := flag != 0
	fmt.Println("flag в Boolean:", bul2)

	// Преобразуйте значение в Boolean
	str_one := "Privet"
	str_two := ""
	b1 := str_one != ""
	b2 := str_two != ""
	fmt.Println("str_one в Boolean:", b1)
	fmt.Println("str_two в Boolean:", b2)

	// Преобразуйте false в строку
	bl := false
	s_bl := strconv.FormatBool(bl)
	fmt.Println("false в строку:", s_bl)
}

func task3() {
	// Нужно обменять значения переменных местами.
	var age1 float64 = 36.6
	var temperature float64 = 25
	var bufer float64

	bufer = age1
	age1 = temperature
	temperature = bufer

	fmt.Println("обменяли значения переменных местами", age1, temperature)
}

func task4() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	m := x
	if y > m {
		m = y
	}
	if z > m {
		m = z
	}
	fmt.Println("Наибольшее число", m)
}

func task5() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(float64(a) / float64(b))
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	fmt.Println(math.Pow(float64(a), float64(b)))
}

func task6() {
	var a float64
	fmt.Scan(&a)
	fmt.Println(a * a)
}

func task7() {
	var A, B, C float64
	fmt.Scan(&A, &B, &C)

	ac := math.Abs(A - C)
	bc := math.Abs(B - C)
	fmt.Println(ac, bc, ac+bc)
}

func task8() {
	var n int
	fmt.Scan(&n)

	a := n / 1000
	b := (n / 100) % 10
	c := (n / 10) % 10
	d := n % 10
	fmt.Println(a == d && b == c)
}

// task9
func task9() {
	var A, B float64
	fmt.Scan(&A, &B)
	fmt.Println(-B / A)
}

// task10
func task10() {
	var A int
	fmt.Scan(&A)
	fmt.Println(A%2 == 0)
	fmt.Println(A%2 != 0)
}

// task11
func task11() {
	var m int
	fmt.Scan(&m)

	if m == 12 || m == 1 || m == 2 {
		fmt.Println("зима")
	} else if m >= 3 && m <= 5 {
		fmt.Println("весна")
	} else if m >= 6 && m <= 8 {
		fmt.Println("лето")
	} else if m >= 9 && m <= 11 {
		fmt.Println("осень")
	}
}

func task12() {
	var unit int
	var M, k float64
	fmt.Scan(&unit, &M)

	switch unit {
	case 1:
		k = 1
	case 2:
		k = 1e-6
	case 3:
		k = 1e-3
	case 4:
		k = 1000
	case 5:
		k = 100
	}
	fmt.Println(M * k)
}

func task13() {
	var year int
	fmt.Scan(&year)

	colors := []string{"зеленой", "красной", "желтой", "белой", "черной"}
	animals := []string{"крысы", "коровы", "тигра", "зайца", "дракона", "змеи", "лошади", "овцы", "обезьяны", "курицы", "собаки", "свиньи"}

	off := year - 1984
	off = ((off % 60) + 60) % 60

	fmt.Println(colors[(off/12)%5], animals[off%12])
}

func task14() {
	var N int
	fmt.Scan(&N)

	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}

	last := map[int]int{}
	bestD := map[int]int{}
	bestI := map[int][2]int{}
	order := []int{}
	seen := map[int]bool{}

	for i, v := range a {
		if !seen[v] {
			seen[v] = true
			order = append(order, v)
		}
		if j, ok := last[v]; ok {
			d := i - j
			if bd, ok2 := bestD[v]; !ok2 || d < bd {
				bestD[v] = d
				bestI[v] = [2]int{j, i}
			}
		}
		last[v] = i
	}

	for _, v := range order {
		if p, ok := bestI[v]; ok {
			fmt.Println(v, p[0], p[1])
		}
	}
}

func task15() {
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Scan(&a[i])
		a[i]++
	}
	for i := 0; i < 10; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(a[i])
	}
	fmt.Println()
}

// task16
type Base struct {
	Name, Class       string
	Level, Exp        int
	HP, MaxHP, Attack int
}

func (b *Base) GetName() string { return b.Name }
func (b *Base) GetHP() int      { return b.HP }
func (b *Base) IsAlive() bool   { return b.HP > 0 }

func (b *Base) Heal(x int) {
	b.HP += x
	if b.HP > b.MaxHP {
		b.HP = b.MaxHP
	}
}
func (b *Base) Damage(x int) {
	b.HP -= x
	if b.HP < 0 {
		b.HP = 0
	}
}
func (b *Base) Strike() int { return b.Attack }

func (b *Base) SuperStrike() int { return b.Attack * 2 }

type Warrior struct {
	Base
	Armor int
}
type Mage struct {
	Base
	Mana int
}
type Archer struct {
	Base
	Agility int
}

type Character interface {
	GetName() string
	GetHP() int
	IsAlive() bool
	Damage(int)
	Heal(int)
	Strike() int
	SuperStrike() int
}

func (w *Warrior) Damage(x int) { w.Base.Damage(x - x*w.Armor/100) }

func (m *Mage) SuperStrike() int {
	if m.Mana >= 10 {
		m.Mana -= 10
		return m.Attack*2 + 5
	}
	return m.Attack * 2
}

func (a *Archer) Strike() int { return a.Attack + a.Agility/5 }

func fight(a, b Character) {
	turn := 0
	for a.IsAlive() && b.IsAlive() {
		turn++
		dmg := a.Strike()
		if turn%3 == 0 {
			dmg = a.SuperStrike()
		}
		b.Damage(dmg)
		fmt.Println(a.GetName(), "удар", dmg, "=>", b.GetName(), "HP", b.GetHP())
		a, b = b, a
	}
	if a.IsAlive() {
		fmt.Println("Победил", a.GetName())
	} else {
		fmt.Println("Победил", b.GetName())
	}
}

func task16() {
	p1 := &Warrior{Base: Base{Name: "Thor", Class: "Warrior", Level: 1, HP: 60, MaxHP: 60, Attack: 10}, Armor: 20}
	p2 := &Mage{Base: Base{Name: "Merlin", Class: "Mage", Level: 1, HP: 45, MaxHP: 45, Attack: 12}, Mana: 30}
	fight(p1, p2)
}

func task17() {
	allUsers := []string{"id3", "id5", "id9", "id8", "id2", "id1", "id4", "id6", "id7", "id10"}
	offline := []string{"id3", "id9", "id7", "id2", "id4", "id6"}

	off := map[string]bool{}
	for _, u := range offline {
		off[u] = true
	}
	for _, u := range allUsers {
		if !off[u] {
			fmt.Println(u)
		}
	}
}

func task18() {
	books := []string{"id3", "id5", "id9", "id8", "id2", "id1"}
	mags := []string{"id8", "id2", "id1", "id4", "id6", "id7", "id10"}

	set := map[string]bool{}
	for _, u := range books {
		set[u] = true
	}
	for _, u := range mags {
		if set[u] {
			fmt.Println(u)
		}
	}
}

// task19
// todo: Реализовать метод Next() для структуры Node !!!
// Реализовать метод подсчёта времени и расстояния  calcDistance( x Node, y Node) от узла X до узла Y для структуры Route.
type Node struct {
	name string
	time int
	dist float32
	next *Node
}

func (n *Node) Next() *Node {
	if n == nil {
		return nil
	}
	return n.next
}

type Route struct {
	head   *Node
	length int
}

func (r *Route) insertAtHead(name string, time int, dist float32) {
	r.head = &Node{name, time, dist, r.head}
	r.length++
}

func absI(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func absF(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}

func calcDistance(x, y *Node) (int, float32) {
	return absI(y.time - x.time), absF(y.dist - x.dist)
}

func (r *Route) find(name string) *Node {
	for n := r.head; n != nil; n = n.next {
		if n.name == name {
			return n
		}
	}
	return nil
}

func (r *Route) calcDistance(xName, yName string) (int, float32, bool) {
	x, y := r.find(xName), r.find(yName)
	if x == nil || y == nil {
		return 0, 0, false
	}
	t, d := calcDistance(x, y)
	return t, d, true
}

func task19() {

	list := Route{nil, 0}
	list.insertAtHead("C", 15, 18)
	list.insertAtHead("B", 5, 6)
	list.insertAtHead("A", 0, 0)

	t, d, ok := list.calcDistance("A", "C")
	if ok {
		fmt.Println(t, d)
	}
}
