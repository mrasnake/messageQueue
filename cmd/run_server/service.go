package main

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"os"
	"strings"
	"sync"
)

type ItemService struct{
	mu sync.Mutex
	storage *Storage
	logs *os.File
}

func NewService(l string) (*ItemService, error) {
	store := NewStorage()
	f, err := os.OpenFile( l, os.O_CREATE | os.O_WRONLY, 0777)
	if err != nil{
		return nil, fmt.Errorf("unable to open logfile: %w", err)
	}
	out := &ItemService{
		storage: store,
		logs: f,
	}
	return out, nil
}

func (i *ItemService) WriteLog(msg string){
	i.mu.Lock()
	defer i.mu.Unlock()

	fmt.Fprintf( i.logs, "Found value: %v\n", msg )
	return
}

func (i *ItemService) ProcessMessage(in []byte){

	s := strings.Fields(string(in))
	switch s[0] {
	case "ADD_ITEM":
		if err := i.AddItem(&AddItemRequest{
			Value: s[1],
		}); err != nil{
			fmt.Printf("failed to process message %v: %v\n", s[0], err.Error())
		}
	case "GET_ITEM":
		val, err := i.GetItem(&GetItemRequest{
			Value: s[1],
		})
		if err != nil {
			fmt.Printf("failed to process message %v: %v\n", s[0], err.Error())
			return
		}
		i.WriteLog(val)
	case "REMOVE_ITEM":
		if err := i.RemoveItem(&RemoveItemRequest{
			Value: s[1],
		}); err != nil{
			fmt.Printf("failed to process message %v: %v\n", s[0], err.Error())
		}
	case "GET_ALL_ITEMS":
		vals, err := i.ListItems()
		if err != nil {
			fmt.Printf("failed to process message %v: %v\n", s[0], err.Error())
			return
		}
		for _, v := range vals {
			i.WriteLog(v)
		}
	default:
		fmt.Println("Invalid request")

	}

	return
}

type AddItemRequest struct {
	Value string
}

func (a AddItemRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Value, validation.Required),
	)
}

type GetItemRequest struct {
	Value string
}

func (g GetItemRequest) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.Value, validation.Required),
	)
}

type RemoveItemRequest struct {
	Value string
}

func (r RemoveItemRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Value, validation.Required),
	)
}

func (i *ItemService) AddItem(req *AddItemRequest) error{
	if err := req.Validate(); err != nil {
		return fmt.Errorf("invalid request: %w", err)
	}

	return i.storage.AddItem(req.Value)
}

func (i *ItemService) GetItem(req *GetItemRequest) (string, error){
	if err := req.Validate(); err != nil {
		return "", fmt.Errorf("invalid request: %w", err)
	}

	return i.storage.GetItem(req.Value)
}

func (i *ItemService) RemoveItem(req *RemoveItemRequest) error{
	if err := req.Validate(); err != nil {
		return fmt.Errorf("invalid request: %w", err)
	}

	return i.storage.RemoveItem(req.Value)
}

func (i *ItemService) ListItems() ([]string, error){
	return i.storage.ListItems()
}