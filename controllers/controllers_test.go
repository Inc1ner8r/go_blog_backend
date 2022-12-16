package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestControllers(t *testing.T) {
	//should connect to db. Get, post and delete data properly
	var db = InitDB("test")

	//var db = InitDB("test")
	assert.Equal(t, 1223, 123, "they should be equal")
	db.Begin()

	//reset db

	//push books

	//get books

}
