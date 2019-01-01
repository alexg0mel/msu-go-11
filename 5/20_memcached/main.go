package main

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"time"
)

func getRecord(mkey string) *memcache.Item {
	println("get", mkey)
	// получаем одиночную запись
	item, err := memcacheClient.Get(mkey)
	// если записи нету, то для этого есть специальная ошибка, её надо обрабатывать отдеьно, это почти штатная ситуация, а не что-то страшное
	if err == memcache.ErrCacheMiss {
		fmt.Println("Record not found in memcache")
		return nil
	} else if err != nil {
		PanicOnErr(err)
	}
	return item
}

var (
	memcacheClient *memcache.Client
)

func main() {
	MemcachedAddresses := []string{"127.0.0.1:11211"}
	memcacheClient = memcache.New(MemcachedAddresses...)

	mkey := "record_21"

	item := getRecord(mkey)
	fmt.Printf("first get %+v\n", item)

	ttl := 5
	// Set устанавливает значение, не зависимо от того, была там запись или нет
	err := memcacheClient.Set(&memcache.Item{
		Key:   mkey,
		Value: []byte("1"),
		// Expiration указывает через сколько секунд запись пропадёт из кеша
		Expiration: int32(ttl),
	})
	PanicOnErr(err)

	time.Sleep(time.Microsecond)
	// time.Sleep(7 * time.Second)

	item = getRecord(mkey)

	// second get &{Key:record_21 Value:[49] Flags:0 Expiration:0 casid:10977}
	// Value:[49] возвращается слайс байт!
	fmt.Printf("second get %+v\n", item)

	// Add добавляет запись если её ещё не было
	err = memcacheClient.Add(&memcache.Item{
		Key:        mkey,
		Value:      []byte("2"),
		Expiration: int32(ttl),
	})
	// если запись не была добавлена, то вернётся соотвтетсувющая ошибка
	if err == memcache.ErrNotStored {
		fmt.Println("Record not stored")
	} else if err != nil {
		PanicOnErr(err)
	}
	item = getRecord(mkey)
	fmt.Printf("third get %+v\n", item)

	// увеличиваем счетчик внутри на 2
	afterIncrement, err := memcacheClient.Increment(mkey, uint64(2))
	PanicOnErr(err)
	fmt.Println("afterIncrement by 2 ", mkey, "is", afterIncrement)

	// уменьашем счетчик внутри на 1
	afterDecrement, err := memcacheClient.Decrement(mkey, uint64(1))
	PanicOnErr(err)
	fmt.Println("afterDecrement by 1 ", mkey, "is", afterDecrement)

	// для несуществующей записи инкремент невозможен
	afterIncrement, err = memcacheClient.Increment(mkey+"_not_exist", uint64(1))
	fmt.Println("afterIncrement not existing record ", afterIncrement)
	if err == memcache.ErrCacheMiss {
		fmt.Println("Record not exist")
	} else if err == memcache.ErrNotStored {
		fmt.Println("Record not stored")
	} else if err != nil {
		PanicOnErr(err)
	} else {
		fmt.Println("afterDecrement by 1 ", mkey, "is", afterDecrement)
	}

	// получаем несколько записей за раз
	mkeys := []string{mkey, "record_22"}
	println("get multiple", mkeys)
	multipleItems, err := memcacheClient.GetMulti(mkeys)
	PanicOnErr(err)
	fmt.Println("multipleItems", multipleItems)

}

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
