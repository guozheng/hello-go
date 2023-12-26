package main

import (
	"fmt"
	"hello_proto/model"
	"log"

	"google.golang.org/protobuf/proto"
)

//go:generate protoc --go_out=. person.proto
func main() {
	person := &model.Person{
		Name:  "John Doe",
		Id:    1234,
		Email: "jdoe@example.com",
	}

	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Println(data)

	newPerson := &model.Person{}
	err = proto.Unmarshal(data, newPerson)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	if person.GetName() != newPerson.GetName() {
		log.Fatalf("data mismatch %q != %q", person.GetName(), newPerson.GetName())
	}

	fmt.Println(newPerson)
}
