
func dailyTemperatures(temperatures []int) []int {

    result := [][]int{}
    resultTemperatures := make([]int, len(temperatures))

    for idx, value := range temperatures {

        for len(result) > 0 &&
            value > result[len(result)-1][1] {

            previousIndex := result[len(result)-1][0]

            // Number of days until a warmer temperature
            resultTemperatures[previousIndex] = idx - previousIndex

            // Pop
            result = result[:len(result)-1]
        }

        // Push [index, temperature]
        result = append(result, []int{idx, value})
    }

    return resultTemperatures
}
