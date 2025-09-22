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

type SliceStorage struct {
	Storage []Book
}

func (sliceStorage *SliceStorage) GetBook(hash uint64) Book {
	return sliceStorage.Storage[hash%uint64(len(sliceStorage.Storage))]
}

func (sliceStorage *SliceStorage) AddBook(book Book, hash uint64) {
	sliceStorage.Storage[hash%uint64(len(sliceStorage.Storage))] = book
}

func (sliceStorage *SliceStorage) RemoveBook(hash uint64) {
	sliceStorage.Storage[hash%uint64(len(sliceStorage.Storage))] = Book{}
}

type MapStorage struct {
	Storage map[uint64]Book
}

func (mapStorage *MapStorage) GetBook(hash uint64) Book {
	book, ok := mapStorage.Storage[hash]
	if !ok {
		return Book{}
	}
	return book
}

func (mapStorage *MapStorage) AddBook(book Book, hash uint64) {
	mapStorage.Storage[hash] = book
}

func (mapStorage *MapStorage) RemoveBook(hash uint64) {
	delete(mapStorage.Storage, hash)
}

type StorageInterface interface {
	GetBook(uint64) Book
	AddBook(Book, uint64)
	RemoveBook(uint64)
}

type Library struct {
	Storage   StorageInterface
	GetIdFunc func(string) uint64
}

func HashStr(str string) uint64 {
	hash := uint64(0)
	base := uint64(26)
	mod := uint64(1000000009)
	pow := uint64(1)
	for _, char := range str {
		hash += (uint64(char) * pow) % mod
		hash %= mod
		pow = (pow * base) % mod
	}
	return hash
}

func AdvancedHash(str string) uint64 {
	hash := fnv.New64a()
	hash.Write([]byte(str))
	return hash.Sum64()
}

func (library *Library) AddBook(name, author, edit string) {
	book := Book{name, author, edit}
	new_key := library.GetIdFunc(name)
	library.Storage.AddBook(book, new_key)
}

func (library *Library) GetBook(str string) Book {
	book_key := library.GetIdFunc(str)
	book := library.Storage.GetBook(book_key)
	if book == (Book{}) {
		fmt.Println("Error: No book named " + str)
		return book
	}
	library.Storage.RemoveBook(book_key)
	return book
}

func main() {
	sliceStorage := SliceStorage{Storage: make([]Book, 1000000)}
	mapStorage := MapStorage{Storage: make(map[uint64]Book)}

	books1 := [][]string{{"War and Peace", "Tolstoy", "ECSMO"}, {"Anna Karenina", "Tolstoy", "AST"}, {"Doctor Zhivago", "Pasternak", "AST"}, {"The Brothers Karamazov", "Dostoyevsky", "EBook"}}

	library := &Library{Storage: &sliceStorage, GetIdFunc: HashStr}
	for _, book_info := range books1 {
		library.AddBook(book_info[0], book_info[1], book_info[2])
	}
	books2 := []string{"War and Peace", "Crime and Punishment", "Doctor Zhivago", "The Brothers Karamazov"}
	for _, book_name := range books2 {
		fmt.Println(library.GetBook(book_name))
	}

	library.Storage = &mapStorage
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
