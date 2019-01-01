package main

import (
	"database/sql"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"

	"encoding/json"
	"reflect"
)

var (
	c  redis.Conn
	db *sql.DB
)

type CacheItem struct {
	Data interface{} // можем класть любые данные
	Tags map[string]int
}

type Articles []Article

type Article struct {
	Name string
	From string
}

func getCacheRecord(mkey string) string {
	println("get", mkey)
	// получает запись, https://redis.io/commands/get
	item, err := redis.String(c.Do("GET", mkey))
	// если записи нету, то для этого есть специальная ошибка, её надо обрабатывать отдеьно, это почти штатная ситуация, а не что-то страшное
	if err == redis.ErrNil {
		fmt.Println("Record not found in redis (return value is nil)")
		return ""
	} else if err != nil {
		PanicOnErr(err)
	}
	return item
}

// PrintByID print student by id
func GetFioByID(id int64) (string, error) {
	println("call SQL")
	var fio string
	row := db.QueryRow("SELECT fio FROM students WHERE id = ?", id)
	err := row.Scan(&fio)
	return fio, err
}

func getCachedFio(mkey string) (string, error) {
	println("redis get", mkey)
	// получает запись, https://redis.io/commands/get
	data, err := c.Do("GET", mkey)
	item, err := redis.String(data, err)
	// если записи нету, то для этого есть специальная ошибка, её надо обрабатывать отдеьно, это почти штатная ситуация, а не что-то страшное
	if err == redis.ErrNil {
		fmt.Println("Record not found in redis (return value is nil)")
		return "", redis.ErrNil
	} else if err != nil {
		return "", err
	}
	return item, nil
}

func main() {
	var err error
	// соединение
	c, err = redis.DialURL("redis://user:@localhost:6379/0")
	// c, err = redis.DialURL(os.Getenv("REDIS_URL"))
	PanicOnErr(err)
	defer c.Close()

	db, err = sql.Open("mysql", "root@tcp(localhost:3306)/msu-go-11?charset=utf8&interpolateParams=true")
	PanicOnErr(err)

	println("--- taggedCache")
	taggedCache()
	println("--- lockCacheRebuild")
	lockCacheRebuild()
}

func taggedCache() {
	top := Articles{
		Article{"Джава и Докер - это должен знать каждый", "Хабр"},
		Article{"Как взрываются базовые станиции", "Гиктаймс"},
	}

	item := CacheItem{
		Data: top,
		// Tags это метки валидности записи в кеше
		// данная запись если эти метки в кеше имеют такое же значение
		Tags: map[string]int{
			"Habr": 1,
			"GT":   1,
		},
	}

	jsonData, _ := json.Marshal(item)
	println("json to store: ", string(jsonData))

	mkey := "top_news_mobile"
	// добавляет запись, https://redis.io/commands/set
	result, err := redis.String(c.Do("SET", mkey, jsonData))
	if result != "OK" {
		panic("result not ok: " + result)
	}
	c.Do("SET", "Habr", 1)
	c.Do("SET", "GT", 1)

	// если раскомментировать эту строчку, то наш кеш перестанет быть валидным
	// c.Do("INCR", "Habr")

	topCache, err := getCachedFio(mkey)
	fmt.Println("top Cache", topCache)

	cItems := CacheItem{}
	_ = json.Unmarshal([]byte(topCache), &cItems)
	fmt.Printf("top Cache unpacked %+v\n", cItems)

	keys := make([]interface{}, 0)
	toCompare := make([]int, 0)
	for key, val := range cItems.Tags {
		keys = append(keys, key)
		toCompare = append(toCompare, val)
	}

	// https: //redis.io/commands/mget
	reply, err := redis.Ints(c.Do("MGET", keys...))
	PanicOnErr(err)

	fmt.Println("compare cached values", toCompare, "with current values", reply)

	println("cache record is valid:", reflect.DeepEqual(toCompare, reply))

}

func lockCacheRebuild() {
	userID := 11
	mkey := "top_user_" + strconv.Itoa(userID)

	var fio string
	var err error
	// если кеш есть, то мы выходим сразу, если нет - у нас 4 попытки получить значение
	for i := 0; i < 4; i++ {
		fio, err = getCachedFio(mkey)
		// записи по кешу не нашлось
		if err == redis.ErrNil {
			// пытаемся сказать "я строю этот кеш, другие - ждите"
			lockStatus, _ := redis.String(c.Do("SET", mkey+"_lock", fio, "EX", 3, "NX"))
			if lockStatus != "OK" {
				// кто-то другой держит лок, подождём и попробуем получить запись еще раз
				println("sleep", i)
				time.Sleep(time.Millisecond * 10)
			} else {
				// успешло залочились, можем строить кеш
				break
			}
		} else if err != nil {
			PanicOnErr(err)
		} else {
			// запись нашлась
			break
		}
	}
	// если записи нету, то надо её построить и положить туда
	if err == redis.ErrNil {
		println("Create cache data")

		// основная работа по доставанию данных для кеша
		// потенциально тяжелая операция
		fio, err = GetFioByID(1)
		PanicOnErr(err)

		// кладём запись в кеш
		ttl := 50
		// добавляет запись, https://redis.io/commands/set
		result, err := redis.String(c.Do("SET", mkey, fio, "EX", ttl))
		PanicOnErr(err)
		if result != "OK" {
			panic("result not ok: " + result)
		}

		// удаляем лок на построение
		n, err := redis.Int(c.Do("DEL", mkey+"_lock"))
		PanicOnErr(err)
		println("lock deleted:", n)
	}

	println(fio)
}

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
