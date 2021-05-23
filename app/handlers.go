package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Employee struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

type EmployeeAll struct {
	EmpAll []Employee `json:"AllEmp"` // JSON name should be the resource name given in json to be parsed
}

var emp EmployeeAll

func readJSON() {
	// Open our jsonFile
	jsonFile, err := os.Open("D:\\GoTest\\TestApI\\Employee.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")

	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	fbytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Successsfully read json file")
	fmt.Println(fbytes)

	err = json.Unmarshal(fbytes, &emp)

	if err != nil {
		fmt.Println("Unmarshalling failed")
	}

	fmt.Println("Unmarshalled")
	fmt.Println(emp)
	for i := 0; i < len(emp.EmpAll); i++ {
		fmt.Println(emp.EmpAll[i].Name)
		fmt.Println(emp.EmpAll[i].Email)
	}

}

func getEmployees(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") // This line should before anything is send as response.
	fmt.Println(w)
	fmt.Println(*r)
	fmt.Fprint(w, "Hello\n")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(emp.EmpAll)
}
