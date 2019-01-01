/*Package world документация ко всему пакету*/
package world

// импорт распространяется на ФАЙЛ, а не на ПАКЕТ
import "fmt"

//StartingRoom is name of room which user enter at game start
const StartingRoom = "lobby"

const startingLevel = 1

// GetStartingLevel документация к методу
func GetStartingLevel() int {
	return startingLevel
}

func checkStartingLevel(level int) bool {
	return level == startingLevel
}

//PrintStartRoom выводит стартовую позицию
func PrintStartRoom() {
	fmt.Println("Start level is ", StartingRoom)
}
