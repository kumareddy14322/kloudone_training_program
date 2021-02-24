package main

//DBUSER const
const DBUSER string = "postgress"

//DBPASS const
const DBPASS string = "kumar"

//DBNAME const
const DBNAME string = "learn"

//PORT const
const PORT string = ":8080"

func main() {
	a := app{}
	a.Initialize(DBUSER, DBPASS, DBNAME)
	a.Run(PORT)
}