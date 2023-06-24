package model

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var models *RealModels

func init() {
	models = newRealModels()
}

func GetModels() *RealModels {
	return models
}

type Models interface {
	CreateModel(string) Model
}

type RealModels struct {
	models map[string]Model
	onces  map[string]*sync.Once
}

func newRealModels() *RealModels {
	return &RealModels{
		onces:  make(map[string]*sync.Once),
		models: make(map[string]Model),
	}
}

func (m *RealModels) CreateModel(name string) Model {

	once, ok := m.onces[name]
	if !ok {
		once = &sync.Once{}
		m.onces[name] = once
	}
	once.Do(func() {
		file, err := os.Open(fmt.Sprintf("%s.mat", name))
		if err != nil {
			panic(err)
		}
		defer file.Close()
		matrix := fileToMatrix(file)
		m.models[name] = NewRealModel(matrix)
	})
	return m.models[name]
}

func fileToMatrix(file *os.File) [][]float64 {
	scanner := bufio.NewScanner(file)

	matrix := make([][]float64, 1000)
	for i := 0; i < 1000; i++ {
		matrix[i] = make([]float64, 1000)
	}

	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Split(line, " ")
		for col, elem := range elements {
			num, err := strconv.ParseFloat(elem, 64)
			if err != nil {
				panic(err)
			}
			matrix[row][col] = num
		}
		row++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return matrix
}
