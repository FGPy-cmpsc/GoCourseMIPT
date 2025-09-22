package main

import (
	"fmt"
	"hash/fnv"
)

type Book struct {
	BookName string
	Author   string
	Edition  string
}

func (book Book) GetInfo() string {
	return "book name: " + book.BookName + ", author: " + book.Author + ", edition: " + book.Edition
}

type Library struct {
	Storage   map[int64]Book
	GetIdFunc func(str string) int64
}

func HashStr(str string) int64 {
	hash := int64(0)
	base := int64(26)
	mod := int64(1000000009)
	pow := int64(1)
	for _, char := range str {
		hash += (int64(char) * pow) % mod
		pow = (pow * base) % mod
	}
	return hash
}

func (library *Library) AddBook(name, author, edit string) {
	new_key := library.GetIdFunc(name)
	library.Storage[new_key] = Book{name, author, edit}
}

func (library *Library) GetBook(str string) Book {
	book_key := library.GetIdFunc(str)
	book, ok := library.Storage[book_key]
	if !ok {
		fmt.Println("Error: No book named " + str)
		return book
	}
	delete(library.Storage, book_key)
	return book
}

func AdvancedHash(str string) int64 {
	hash := fnv.New64a()
	hash.Write([]byte(str))
	return int64(hash.Sum64())
}

func main() {
	books1 := [][]string{{"War and Peace", "Tolstoy", "ECSMO"}, {"Anna Karenina", "Tolstoy", "AST"}, {"Doctor Zhivago", "Pasternak", "AST"}, {"The Brothers Karamazov", "Dostoyevsky", "EBook"}}
	library := &Library{Storage: make(map[int64]Book), GetIdFunc: HashStr}
	for _, book_info := range books1 {
		library.AddBook(book_info[0], book_info[1], book_info[2])
	}
	books2 := []string{"War and Peace", "Crime and Punishment", "Doctor Zhivago", "The Brothers Karamazov"}
	for _, book_name := range books2 {
		fmt.Println(library.GetBook(book_name))
	}

	library.GetIdFunc = AdvancedHash
	books3 := [][]string{{"A Hero Of Our Time", "Lermontov", "Azbuka"}, {"The Master and Margarita", "Bulgakov", "AST"}, {"Dead Souls", "Gogol", "Myf"}}
	for _, book_info := range books3 {
		library.AddBook(book_info[0], book_info[1], book_info[2])
	}
	books4 := []string{"Crime and Punishment", "The Master and Margarita"}
	for _, book_name := range books4 {
		fmt.Println(library.GetBook(book_name).GetInfo())
	}
}
