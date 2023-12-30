package test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/k0kubun/pp"
	"golang.org/x/sync/singleflight"
)

func TestSingleFlight(t *testing.T) {
	var group singleflight.Group

	// Fungsi untuk mendapatkan data cuaca (contoh fungsi yang memerlukan waktu untuk dieksekusi)
	getWeather := func(city string) (interface{}, error) {
		pp.Printf("Fetching weather data for %s\n", city)
		// Misalkan proses pengambilan data cuaca membutuhkan waktu
		time.Sleep(time.Microsecond * 1)
		return fmt.Sprintf("Weather data for %s", city), nil
	}

	// Fungsi untuk mendapatkan data cuaca dengan penggunaan singleflight
	getWeatherWithSingleFlight := func(city string) (interface{}, error) {
		val, err, _ := group.Do(city, func() (interface{}, error) {
			return getWeather(city)
		})
		return val, err
	}

	var wg sync.WaitGroup
	// Pengujian dengan beberapa pemanggilan ke fungsi getWeatherWithSingleFlight
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, city string) {
			defer wg.Done()
			data, err := getWeatherWithSingleFlight(city)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}
			fmt.Printf("Data cuaca: %s\n", data)
		}(&wg, "Semarang")
	}

	wg.Wait()

	// Tunggu beberapa saat agar semua goroutine selesai
	// time.Sleep(5 * time.Second)
}
