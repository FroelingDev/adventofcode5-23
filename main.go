package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	seeds := seedsNumbers()
	fmt.Println(seeds)
	fmt.Println()
	fmt.Println()


	seedToSoil := MapAll("seedToSoil.txt")
	fmt.Println(seedToSoil)
	seedToSoilCor := getCorresponding(seedToSoil)
	fmt.Printf("Corresponding: %v\n", seedToSoilCor)
	seedToSoilDif := getDifference(seedToSoil)
	fmt.Printf("Difference: %v\n", seedToSoilDif)
	fmt.Println()
	fmt.Println()


	soilToFertilizer := MapAll("soilToFertilizer.txt")
	fmt.Println(soilToFertilizer)
	soilToFertilizerCor := getCorresponding(soilToFertilizer)
	fmt.Printf("Corresponding: %v\n", soilToFertilizerCor)
	soilToFertilizerDif := getDifference(soilToFertilizer)
	fmt.Printf("Difference: %v\n", soilToFertilizerDif)
	fmt.Println()
	fmt.Println()


	fertilizerToWater := MapAll("fertilizerToWater.txt")
	fmt.Println(fertilizerToWater)
	fertilizerToWaterCor := getCorresponding(fertilizerToWater)
	fmt.Printf("Corresponding: %v\n", fertilizerToWaterCor)
	fertilizerToWaterDif := getDifference(fertilizerToWater)
	fmt.Printf("Difference: %v\n", fertilizerToWaterDif)
	fmt.Println()
	fmt.Println()


	waterToLight := MapAll("waterToLight.txt")
	fmt.Println(waterToLight)
	waterToLightCor := getCorresponding(waterToLight)
	fmt.Printf("Corresponding: %v\n", waterToLightCor)
	waterToLightDif := getDifference(waterToLight)
	fmt.Printf("Difference: %v\n", waterToLightDif)
	fmt.Println()
	fmt.Println()


	lightToTemperature := MapAll("lightToTemperature.txt")
	fmt.Println(lightToTemperature)
	lightToTemperatureCor := getCorresponding(lightToTemperature)
	fmt.Printf("Corresponding: %v\n", lightToTemperatureCor)
	lightToTemperatureDif := getDifference(lightToTemperature)
	fmt.Printf("Difference: %v\n", lightToTemperatureDif)
	fmt.Println()
	fmt.Println()


	temperatureToHumidity := MapAll("temperatureToHumidity.txt")
	fmt.Println(temperatureToHumidity)
	temperatureToHumidityCor := getCorresponding(temperatureToHumidity)
	fmt.Printf("Corresponding: %v\n", temperatureToHumidityCor)
	temperatureToHumidityDif := getDifference(temperatureToHumidity)
	fmt.Printf("Difference: %v\n", temperatureToHumidityDif)
	fmt.Println()
	fmt.Println()


	humidityToLocation := MapAll("humidityToLocation.txt")
	fmt.Println(humidityToLocation)
	humidityToLocationCor := getCorresponding(humidityToLocation)
	fmt.Printf("Corresponding: %v\n", humidityToLocationCor)
	humidityToLocationDif := getDifference(humidityToLocation)
	fmt.Printf("Difference: %v\n", humidityToLocationDif)
	fmt.Println()
	fmt.Println()


	for i :=0; i < len(seeds); i++ {
		seeds[i] = applyDifference(seedToSoil, seedToSoilDif, seedToSoilCor, seeds[i])
		seeds[i] = applyDifference(soilToFertilizer, soilToFertilizerDif, soilToFertilizerCor, seeds[i])
		seeds[i] = applyDifference(fertilizerToWater, fertilizerToWaterDif, fertilizerToWaterCor, seeds[i])
		seeds[i] = applyDifference(waterToLight, waterToLightDif, waterToLightCor, seeds[i])
		seeds[i] = applyDifference(lightToTemperature, lightToTemperatureDif, lightToTemperatureCor, seeds[i])
		seeds[i] = applyDifference(temperatureToHumidity, temperatureToHumidityDif, temperatureToHumidityCor, seeds[i])
		seeds[i] = applyDifference(humidityToLocation, humidityToLocationDif, humidityToLocationCor, seeds[i])
	}

	smallest := seeds[0]
	for _, seed := range seeds {
		if seed < smallest {
			smallest = seed
		}
	}
	fmt.Println("Smallest seed:", smallest)


}

func applyDifference(mat [][]int, diff []int, cor [][]int, seed int) int {
	for i := range mat {
		if seed >= mat[i][1] && seed <= cor[i][1] {
			seed += diff[i]
			break
		}
	}
	return seed
}

func getCorresponding(mat [][]int) [][]int {
	corresponding := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		corresponding[i] = make([]int, 2)
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < mat[i][2]; j++ {
			corresponding[i][0] = mat[i][0] + j
			corresponding[i][1] = mat[i][1] + j
		}
	}

	return corresponding
}

func getDifference(mat [][]int) []int {
	res := make([]int, len(mat))
	for i:=0; i < len(mat); i++ {
		res[i] = mat[i][0] - mat[i][1]
	}
	return res
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func seedsNumbers() []int {
	seeds, err := ioutil.ReadFile("seeds.txt")
	check(err)

	numStr := strings.Fields(string(seeds))
	numbers := make([]int, len(numStr))

	for i, s := range numStr {
		num, err := strconv.Atoi(s)
		check(err)
		numbers[i] = num
	}

	result := []int{}

	for i := 0; i < len(numbers); i += 2 {
		if i+1 < len(numbers) {
			start := numbers[i]
			count := numbers[i+1]
			for j := 0; j < count; j++ {
				result = append(result, start+j)
			}
		}
	}

	return result
}


func MapAll(filename string) [][]int {
	seeds, err := ioutil.ReadFile(filename)
	check(err)

	// Split by newlines to get each row
	rows := strings.Split(strings.TrimSpace(string(seeds)), "\n")
	seedMap := make([][]int, len(rows))

	for i, row := range rows {
		// Split each row by spaces to get the numbers
		numStr := strings.Split(row, " ")
		if len(numStr) != 3 {
			panic("Each row should contain exactly 3 values")
		}
		
		seedMap[i] = make([]int, 3)
		for j, s := range numStr {
			num, err := strconv.Atoi(s)
			check(err)
			seedMap[i][j] = num
		}
	}
	
	return seedMap
}


