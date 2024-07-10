package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func main() {
	people := make(map[string]Man)

	people["Иван Петров"] = Man{Name: "Иван", LastName: "Петров", Age: 30, Gender: "Мужской", Crimes: 2}
	people["Вера Воронова"] = Man{Name: "Вера", LastName: "Воронова", Age: 25, Gender: "Женский", Crimes: 0}
	people["Пётр Иванов"] = Man{Name: "Пёт", LastName: "Иванов", Age: 40, Gender: "Мужской", Crimes: 5}
	people["Наталья Соколова"] = Man{Name: "Наталья", LastName: "Соколова", Age: 35, Gender: "Женский", Crimes: 1}
	people["Сергей Кузнецов"] = Man{Name: "Сергей", LastName: "Кузнецов", Age: 28, Gender: "Мужской", Crimes: 3}
	people["Светлана Семёнова"] = Man{Name: "Светлана", LastName: "Семёнова", Age: 22, Gender: "Женский", Crimes: 0}
	people["Андрей Смирнов"] = Man{Name: "Андрей", LastName: "Смирнов", Age: 50, Gender: "Мужской", Crimes: 4}
	people["Анастасия Сидорова"] = Man{Name: "Анастасия Сидорова", LastName: "Анастасия Сидорова", Age: 27, Gender: "Женский", Crimes: 2}
	people["Дмитрий Васильев"] = Man{Name: "Дмитрий", LastName: "Васильев", Age: 33, Gender: "Мужской", Crimes: 1}
	people["Ирина Попова"] = Man{Name: "Ирина", LastName: "Попова", Age: 29, Gender: "Женский", Crimes: 0}

	suspects := []string{"Иван Петров", "Вера Воронова", "Пётр Иванов", "Наталья Соколова", "Сергей Кузнецов", "Светлана Семёнова", "Андрей Смирнов", "Анастасия Сидорова", "Дмитрий Васильев", "Ирина Попова"}

	fmt.Println("\nПодозреваемый, у которого наибольшее количество совершённых преступлений:")
	fmt.Println("-------------------------------------------------------------------------")
	maxCrimes := 0
	var suspect string

	if len(people) == 0 {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым.")
	}

	fmt.Println("-------------------------------------------------------------------------")

	for _, name := range suspects {
		if person, exists := people[name]; exists {
			if person.Crimes > maxCrimes {
				maxCrimes = person.Crimes
				suspect = name
			}
		}
	}
	fmt.Printf("%s - преступлений: %d\n", suspect, maxCrimes)
	fmt.Println("-------------------------------------------------------------------------")
}
