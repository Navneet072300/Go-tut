package main

import (
	_ "github.com/lib/pq"
)

type user struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}