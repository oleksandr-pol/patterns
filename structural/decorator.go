package structural

import "log"

func LogDecorate(fn func(float64) float64) func(float64) float64 {
	return func(n float64) float64 {
		log.Println("Starting the execution with the integer", n)

		result := fn(n)

		log.Println("Execution is completed with the result", result)

		return result
	}
}
