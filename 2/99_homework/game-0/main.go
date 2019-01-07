package main

import "strings"

func main() {

}

var game *Game

func initGame() {

	game = new(Game)
	game.InitGame()
}

func handleCommand(command string) string {
	parced := strings.Split(command, " ")
	task, exist := game.gamer.tasks[parced[0]]
	if !exist {
		return "неизвестная команда"
	}
	return task.do(parced[1:]...)
}

type Game struct {
	rooms map[string]*room
	doors map[string][]*door
	gamer *gamer
}

func (g *Game) InitGame() {
	g.rooms = make(map[string]*room)
	g.doors = make(map[string][]*door)
	var ownRoomTableThings []*thing
	var ownRoomChairThings []*thing
	var ownRoomFurniture []*furniture

	ownRoomTableThings = append(ownRoomTableThings, NewThing("ключи", false))
	ownRoomTableThings = append(ownRoomTableThings, NewThing("конспекты", false))

	ownRoomChairThings = append(ownRoomChairThings, NewThing("рюкзак", true))

	ownRoomFurniture = append(ownRoomFurniture, NewFurniture("стол", "на столе: ", ownRoomTableThings))
	ownRoomFurniture = append(ownRoomFurniture, NewFurniture("стул", "на стуле - ", ownRoomChairThings))

	g.rooms["кухня"] = NewRoom("кухня", "ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ", "кухня, ничего интересного.", make([]*furniture, 0))
	g.rooms["коридор"] = NewRoom("коридор", "", "ничего интересного.", make([]*furniture, 0))
	g.rooms["комната"] = NewRoom("комната", "", "ты в своей комнате.", ownRoomFurniture)
	g.rooms["улица"] = NewRoom("улица", "", "на улице весна.", ownRoomFurniture)
	g.rooms["дом"] = NewRoom("дом", "", "ничего интересного.", ownRoomFurniture)

	g.doors["кухня"] = append(g.doors["кухня"], &door{"коридор", true})
	g.doors["комната"] = append(g.doors["комната"], &door{"коридор", true})
	g.doors["коридор"] = append(g.doors["коридор"], &door{"кухня", true})
	g.doors["коридор"] = append(g.doors["коридор"], &door{"комната", true})
	g.doors["коридор"] = append(g.doors["коридор"], &door{"улица", false})
	g.doors["улица"] = append(g.doors["улица"], &door{"домой", true})

	g.gamer = NewGamer(g.rooms["кухня"], map[string]*task{})

	g.gamer.tasks["осмотреться"] = &task{do: func(params ...string) string {
		return g.gamer.in_room.greeting + g.getThings() + g.nextWay()
	}}

	g.gamer.tasks["идти"] = &task{do: func(params ...string) string {
		if len(params) == 0 {
			return "не задан путь"
		}
		opened, finded := g.findWay(params[0])
		if !finded {
			return "нет пути в " + params[0]
		}
		if !opened {
			return "дверь закрыта"
		}
		g.gamer.in_room = g.rooms[params[0]]
		return g.gamer.in_room.message + g.nextWay()
	}}

	g.gamer.tasks["одеть"] = &task{do: func(params ...string) string {
		if len(params) == 0 {
			return "не известно что одевать"
		}
		for _, itemFurniture := range g.gamer.in_room.furnitures {
			newThing := make([]*thing, 0, len(itemFurniture.things))
			for _, itemThing := range itemFurniture.things {
				if itemThing.name == params[0] {
					if !itemThing.can_activate_inv {
						return "нельзя одеть " + params[0]
					}
					g.gamer.inventory_active = true
				} else {
					newThing = append(newThing, itemThing)
				}
			}
			if g.gamer.inventory_active {
				itemFurniture.things = newThing
				kitchen := g.rooms["кухня"]
				kitchen.greeting = strings.Replace(kitchen.greeting, "собрать рюкзак и ", "", 1)

				break
			}
		}
		if !g.gamer.inventory_active {
			return "ничего не одели"
		}
		return "вы одели: " + params[0]
	}}

	g.gamer.tasks["взять"] = &task{do: func(params ...string) string {
		var finded bool
		if len(params) == 0 {
			return "не известно что взять"
		}

		if !g.gamer.inventory_active {
			return "некуда класть"
		}

		for _, itemFurniture := range g.gamer.in_room.furnitures {
			newThing := make([]*thing, 0, len(itemFurniture.things))
			for _, itemThing := range itemFurniture.things {
				if itemThing.name == params[0] {
					finded = true
					g.gamer.things = append(g.gamer.things, itemThing)
				} else {
					newThing = append(newThing, itemThing)
				}
			}
			if finded {
				itemFurniture.things = newThing
				break
			}
		}
		if !finded {
			return "нет такого"
		}
		return "предмет добавлен в инвентарь: " + params[0]
	}}

	g.gamer.tasks["применить"] = &task{do: func(params ...string) string {
		var findedThing, findedClosedDoor bool
		if len(params) < 2 {
			return "не задано к чему применять"
		}
		if len(params) < 1 {
			return "не задано что применять"
		}
		//if !g.gamer.inventory_active {
		//	return "нет инвентаря"
		//}

		for _, itemThing := range g.gamer.things {
			if itemThing.name == params[0] {
				findedThing = true
				break
			}
		}
		if !findedThing {
			return "нет предмета в инвентаре - " + params[0]
		}

		for _, itemDoor := range g.doors[g.gamer.in_room.name] {
			if !itemDoor.opened {
				itemDoor.opened = true
				findedClosedDoor = true
			}
		}
		if params[1] != "дверь" {
			return "не к чему применить"
		}

		if !findedClosedDoor {
			return "нет предмета в инвентаре - " + params[0]
		}

		return "дверь открыта"
	}}

}

func (g *Game) nextWay() string {
	res := " можно пройти - "
	for i, itemDoor := range g.doors[g.gamer.in_room.name] {
		if i > 0 {
			res += ", "
		}
		res = res + itemDoor.wayTo
	}
	return res
}

func (g *Game) findWay(nameDoor string) (opened bool, finded bool) {
	for _, itemDoor := range g.doors[g.gamer.in_room.name] {
		if itemDoor.wayTo == nameDoor {
			finded = true
			opened = itemDoor.opened
			break
		}
	}
	return
}

func (g Game) getThings() (res string) {
	for i, itemFurniture := range g.gamer.in_room.furnitures {
		if i > 0 {
			if len(itemFurniture.things) > 0 {
				res += ", "
			}
		}
		if len(itemFurniture.things) > 0 {
			res += itemFurniture.pname
		}
		for j, itemThing := range itemFurniture.things {
			if j > 0 {
				res += ", "
			}
			res += itemThing.name
		}
	}
	if res == "" && g.gamer.in_room.name == "комната" {
		res += "пустая комната"
	}
	res += "."
	return
}

type room struct {
	name       string
	greeting   string // приветствие при команде осмотреться
	message    string //сообщение при переходе в комнату
	furnitures []*furniture
}

func NewRoom(name, greeting, message string, furniture []*furniture) *room {
	return &room{
		name:       name,
		greeting:   greeting,
		message:    message,
		furnitures: furniture,
	}
}

type door struct {
	wayTo  string
	opened bool
}

type thing struct {
	name             string
	can_activate_inv bool
}

func NewThing(name string, can_activate_inv bool) *thing {
	return &thing{
		name:             name,
		can_activate_inv: can_activate_inv,
	}
}

type furniture struct {
	name   string
	pname  string //имя вместе с предлогом  в/на
	things []*thing
}

func NewFurniture(name, pname string, things []*thing) *furniture {
	return &furniture{
		name:   name,
		pname:  pname,
		things: things,
	}
}

type task struct {
	do func(params ...string) string
}

type gamer struct {
	inventory_active bool             //одет ли рюкзак
	things           []*thing         // что в инвентаре игрока
	in_room          *room            // где сейчас игрок
	tasks            map[string]*task // доступные задачи игрока
}

func NewGamer(startRoom *room, tasks map[string]*task) *gamer {
	return &gamer{
		inventory_active: false,
		things:           make([]*thing, 0),
		in_room:          startRoom,
		tasks:            tasks,
	}
}
