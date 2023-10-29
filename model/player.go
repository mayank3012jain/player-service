package model

type Player struct {
    ID   string    `csv:"id" json:"id"`
    Name string `csv:"name" json:"name"`
    Age  int    `csv:"age" json:"age"`
}