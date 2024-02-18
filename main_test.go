package main

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД

	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		require.NoError(t, err)
	}
	defer db.Close()

	clientID := 1

	// напиши тест здесь
	cl, err := selectClient(db, clientID)

	require.NoError(t, err)

	assert.Equal(t, clientID, cl.ID)
	assert.NotEmpty(t, cl.FIO)
	assert.NotEmpty(t, cl.Login)
	assert.NotEmpty(t, cl.Birthday)
	assert.NotEmpty(t, cl.Email)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД

	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	clientID := -1

	// напиши тест здесь
	cl, err := selectClient(db, clientID)

	require.Equal(t, sql.ErrNoRows, err)

	assert.Empty(t, cl.ID)
	assert.Empty(t, cl.FIO)
	assert.Empty(t, cl.Login)
	assert.Empty(t, cl.Birthday)
	assert.Empty(t, cl.Email)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД

	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	cl.ID, err = insertClient(db, cl)

	require.NoError(t, err)
	require.NotEmpty(t, cl.ID)

	stored, err := selectClient(db, cl.ID)
	require.NoError(t, err)

	assert.Equal(t, cl, stored)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД

	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	id, err := insertClient(db, cl)

	require.NoError(t, err)
	require.NotEmpty(t, id)

	_, err = selectClient(db, id)
	require.NoError(t, err)

	err = deleteClient(db, id)
	require.NoError(t, err)

	_, err = selectClient(db, id)
	require.Equal(t, sql.ErrNoRows, err)
}
