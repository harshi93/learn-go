// Maps = Dict

package main

import "fmt"

func main() {

	StudentAge := make(map[string]int)

	StudentAge["Harsh"] = 23
	StudentAge["Avi"] = 20
	StudentAge["Ronald"] = 19
	StudentAge["Ravikishan"] = 34

	//Print Everything
	fmt.Println(StudentAge)

	//Print Ravikishans Age
	fmt.Println(StudentAge["Ravikishan"])

	//get length of maps
	fmt.Println(len(StudentAge))

	//nested maps
	superhero :=
		map[string]map[string]string{
			"Superman": map[string]string{
				"RealName": "Clark Kent",
				"City":     "Metropolis",
			},

			"Batman": map[string]string{
				"RealName": "Bruce Wayne",
				"City":     "Gotham City",
			},
		}

	if temp, hero := superhero["Superman"]; hero {

		fmt.Println(temp["RealName"], temp["City"])

	}

}
