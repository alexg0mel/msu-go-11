package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// сколько в среднем спим при эмуляции работы
const AvgSleep = 50

func TrackTimingToContext(ctx context.Context, metricName string, start time.Time) {
	// получаем сколько прошлоо времени с момента как таймер стартовал
	elapsed := time.Since(start)
	// получаем тайминги из контекста
	// поскольку там пустой интерфейс, то нам надо преобразовать к нужному типу
	timings, ok := ctx.Value(timingsKey).(*ctxTimings)
	if !ok {
		return
	}
	// лочимся на случай конкурентной записи в мапку
	timings.Lock()
	defer timings.Unlock()
	// если меткри ещё нет - мы её создадим, если есть - допишем в существующую
	if metric, metricExist := timings.Data[metricName]; !metricExist {
		timings.Data[metricName] = &Timing{
			Count:    1,
			Duration: elapsed,
		}
	} else {
		metric.Count++
		metric.Duration += elapsed
	}
}

func checkSession(ctx context.Context) {
	defer TrackTimingToContext(ctx, "checkSession", time.Now())
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(AvgSleep)))
}

func getPosts(ctx context.Context) {
	defer TrackTimingToContext(ctx, "getPosts", time.Now())
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(AvgSleep)))
}

func loadProfile(ctx context.Context) {
	defer TrackTimingToContext(ctx, "loadProfile", time.Now())
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(AvgSleep)))
}

type Timing struct {
	Count    int
	Duration time.Duration
}

type ctxTimings struct {
	sync.Mutex
	Data map[string]*Timing
}

// линтер ругается если используем базовые типы в Value контекста
// типа так безопаснее разграничивать
type key int

const timingsKey key = 1

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// получаем контект из запроса
		ctx := req.Context()

		// получаем новый контект с хранилищем таймингов
		ctx = context.WithValue(ctx, timingsKey, &ctxTimings{
			Data: make(map[string]*Timing),
		})

		// в конце каждого запроса пишем аксес лог сколько заняло времени
		defer func() {
			// получаем тайминги из контекста
			// поскольку там пустой интерфейс, то нам надо преобразовать к нужному типу
			timings, ok := ctx.Value(timingsKey).(*ctxTimings)
			if !ok {
				return
			}
			accessLog := req.URL.String()
			var total time.Duration
			for timing, value := range timings.Data {
				total += value.Duration
				accessLog += fmt.Sprintf(", %s(%d): %s", timing, value.Count, value.Duration)
			}
			accessLog += fmt.Sprintf(", total: %s", total)

			fmt.Println(accessLog)
			fmt.Fprintln(w, accessLog)
		}()

		// эмулируем какую-то работу
		checkSession(ctx)
		loadProfile(ctx)
		getPosts(ctx)
		getPosts(ctx)
		getPosts(ctx)

		fmt.Fprintln(w, "Request done")
	}))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
