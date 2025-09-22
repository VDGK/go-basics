package main

import (
	"bufio"
	"fmt"
	"os"
)

type bookmarkMap = map[string]string

func main() {
	fmt.Println("___Менеджер Закладок____")
	fmt.Println(`1 - Просмотреть Закладки
2 - Добавить Закладку
3 - Удалить Закладку
4 - Выход
	`)
	bookmarks := make(bookmarkMap)
Menu:
	for {
		input := clientInput()
		switch input {
		case "1":
			listBookmarks(bookmarks)
		case "2":
			addBookmarks(bookmarks)
		case "3":
			removeBookmarks(bookmarks)
		case "4":
			break Menu
		}
	}
}
func clientInput() string {
	fmt.Println("Введите пункт меню:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}
func listBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Пока нет закладок")
	}
	for k, v := range bookmarks {
		fmt.Printf("%s : %s\n", k, v)
	}
}

func addBookmarks(bookmarks bookmarkMap) {
	fmt.Println("Введите ключ: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	k := scanner.Text()
	fmt.Println("Введите значение: ")
	scanner.Scan()
	v := scanner.Text()
	bookmarks[k] = v
}

func removeBookmarks(bookmarks bookmarkMap) {
	fmt.Println("Введите ключи: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	k := scanner.Text()
	delete(bookmarks, k)
}
