package main

import (
	"fmt"
	"math/rand"

	ga "5.2.H"
)

/*
產品
- 產品 1：價格 100 元，重量 2 公斤，類別 A。
- 產品 2：價格 200 元，重量 3 公斤，類別 A。
- 產品 3：價格 150 元，重量 5 公斤，類別 B。
- 產品 4：價格 300 元，重量 4 公斤，類別 B。
- 產品 5：價格 180 元，重量 6 公斤，類別 C。
- 產品 6：價格 250 元，重量 7 公斤，類別 C。

客戶的喜好度:
- 類別 A：80% -> 0.8
- 類別 B：60% -> 0.2
- 類別 C：20% -> 0.1

例如，假設客戶預算為 700 元，購物袋最大承載重量為 15 公斤，根據客戶的喜好度，基因演算法可能推薦以下購買清單：


Gene -> 0~7
Chromosome -> [Gene, Gene, Gene, Gene, Gene, Gene]
*/

type Organism struct {
	chromosome ga.Chromosome
	fitness    float64
}

func (o Organism) GetChromosome() ga.Chromosome {
	return o.chromosome
}

type Item struct {
	Price      int
	Weight     int
	Preference float64
}

func (o Organism) Fitness() float64 {
	if o.fitness != 0 {
		return o.fitness
	}
	amount := 0
	weight := 0
	var fitness float64
	chromosome := o.chromosome

	items := []Item{
		{100, 2, 0.8},
		{200, 3, 0.8},
		{150, 5, 0.2},
		{300, 4, 0.2},
		{180, 6, 0.1},
		{250, 7, 0.1},
	}

	for i, gene := range chromosome {
		tmpPrice := items[i].Price * int(gene)
		tmpWeight := items[i].Weight * int(gene)
		amount += tmpPrice
		weight += tmpWeight

		if tmpPrice < 700 && tmpWeight < 15 {
			fitness += items[i].Preference * float64(gene) * 5
		}

	}

	if amount < 700 {
		fitness += 60
	}

	if weight < 15 {
		fitness += 60
	}

	fitness++
	o.fitness = fitness

	return o.fitness
}

type OrganismFactory struct {
}

func (o *OrganismFactory) Create() ga.Organism {
	chromosome := make(ga.Chromosome, 6)
	for i := 0; i < 6; i++ {
		chromosome[i] = ga.Gene(rand.Intn(8))
	}
	return Organism{chromosome, 0}
}

func (o *OrganismFactory) CreateWithChromosome(chromosome ga.Chromosome) ga.Organism {
	return Organism{chromosome, 0}
}

func (o *OrganismFactory) CreateGene() ga.Gene {
	return ga.Gene(rand.Intn(8))
}

func main() {

	gaOpts := (&ga.GeneticAlgorithmOptions{}).
		SetMaxGeneration(200).
		SetMutationRate(0.01).
		SetPopulationSize(2000).
		SetOrganismFactory(&OrganismFactory{}).
		SetFitnessThreshold(100)

	geneticAlgorithm := ga.NewGeneticAlgorithm(gaOpts)
	res := geneticAlgorithm.Run()
	fmt.Println(res.GetChromosome())
	// [3 1 1 0 0 0]
	/*
		產品 1： *3
		產品 2： *1
		產品 3： *1
		產品 4： *0
		產品 5： *0
		產品 6： *0
	*/
	fmt.Println(res.Fitness())
	// 138
}
