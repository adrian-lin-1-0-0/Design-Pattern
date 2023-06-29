package main

import (
	"fmt"
	"reflect"
	"sync"

	"5.1.H/model"
)

func main() {
	models := model.GetModels()
	m := models.CreateModel("Reflection")
	m2 := models.CreateModel("Reflection")
	fmt.Println(reflect.ValueOf(m).Pointer() == reflect.ValueOf(m2).Pointer())
	// true

	v := make([]float64, 1000)
	for i := 0; i < 1000; i++ {
		v[i] = 1.0
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			models := model.GetModels()
			m := models.CreateModel("Scaling")
			m.Mul(v)
		}()
	}

	wg.Wait()

}
