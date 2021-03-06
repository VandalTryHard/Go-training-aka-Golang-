// Команды инициализации для «if»

package main

import (
	"io/ioutil"
	"log"
)

func saveString(fileName string, str string) error {
	err := ioutil.WriteFile(fileName, []byte(str), 0600)
	return err
}

func main() {
	if err := saveString("english.txt", "Hello"); err != nil { //Область видимости первой переменной «err».
		log.Fatal(err)
	}

	if err := saveString("hindi.txt", "Namaste"); err != nil { // Область видимости второй переменной «err».
		log.Fatal(err)
	}
}
