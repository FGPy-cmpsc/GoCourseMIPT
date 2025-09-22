package main

import (
	"fmt"
	"hash/fnv"
)

type Library struct {
	Storage   map[int64]string
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

func (library *Library) AddBook(str string) {
	new_key := library.GetIdFunc(str)
	library.Storage[new_key] = str
}

func (library *Library) GetBook(str string) string {
	book_key := library.GetIdFunc(str)
	book, ok := library.Storage[book_key]
	if !ok {
		return "Error: No book named " + str
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
	books1 := []string{"War and Peace", "Anna Karenina", "Doctor Zhivago", "The Brothers Karamazov"}
	library := &Library{Storage: make(map[int64]string), GetIdFunc: HashStr}
	for _, book_name := range books1 {
		library.AddBook(book_name)
	}
	books2 := []string{"War and Peace", "Crime and Punishment", "Doctor Zhivago", "The Brothers Karamazov"}
	for _, book_name := range books2 {
		fmt.Println(library.GetBook(book_name))
	}

	library.GetIdFunc = AdvancedHash
	books3 := []string{"A Hero Of Our Time", "The Master and Margarita", "Dead Souls"}
	for _, book_name := range books3 {
		library.AddBook(book_name)
	}
	books4 := []string{"Crime and Punishment", "The Master and Margarita"}
	for _, book_name := range books4 {
		fmt.Println(library.GetBook(book_name))
	}
}
