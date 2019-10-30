package teststore

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/model"
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/store/sqlstore"
	"reflect"
	"testing"
)

func TestJobRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer func() {
		mock.ExpectClose()
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	store := sqlstore.New(db)
	rows := sqlmock.
		NewRows([]string{"id"})

	var elemID int64 = 1
	expect := []*model.Job{
		{ ID: elemID },
	}

	for _, item := range expect {
		rows = rows.AddRow(item.ID)
	}

	u := testUser(t)
	u.ID = 1
	m := testManager(t, u)
	m.ID = 1
	j := testJob(t, m)

	// TODO: uncomment when validation will be implemented
	/*if err := f.Validate(); err != nil {
		t.Fatal()
	}*/

	//INSERT INTO jobs (managerId, title, description, files, specialityId, experienceLevelId, paymentAmount, " +
	//			"country, city, jobTypeId

	//ok query
	mock.
		ExpectQuery(`INSERT INTO jobs`).
		WithArgs(j.HireManagerId, j.Title, j.Description, j.Files, j.SpecialityId, j.ExperienceLevelId,
			j.PaymentAmount, j.Country, j.City, j.JobTypeId).
		WillReturnRows(rows)

	err = store.Job().Create(j, m)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if j.ID != 1 {
		t.Errorf("bad id: want %v, have %v", u.ID, 1)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// query error
	mock.
		ExpectQuery(`INSERT INTO jobs`).
		WithArgs(j.HireManagerId, j.Title, j.Description, j.Files, j.SpecialityId, j.ExperienceLevelId,
			j.PaymentAmount, j.Country, j.City, j.JobTypeId).
		WillReturnError(fmt.Errorf("bad query"))

	err = store.Job().Create(j, m)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestJobRepository_Find(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer func() {
		mock.ExpectClose()
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	var elemID int64 = 1
	// good query
	rows := sqlmock.
		NewRows([]string{"id", "managerId", "title", "description", "files", "specialityId", "experienceLevelId",
			"paymentAmount", "country", "city", "jobTypeId" })

	u := testUser(t)
	u.ID = 1
	m := testManager(t, u)
	m.ID = 1
	expect := []*model.Job{
		testJob(t, m),
	}

	for _, item := range expect {
		rows = rows.AddRow(item.ID, item.HireManagerId, item.Title, item.Description, item.Files, item.SpecialityId,
			item.ExperienceLevelId, item.PaymentAmount, item.Country, item.City, item.JobTypeId)
	}

	mock.
		ExpectQuery("SELECT id, managerId, title, description, files, specialityId, experienceLevelId, paymentAmount, " +
		"country, city, jobTypeId FROM jobs WHERE").
		WithArgs(elemID).
		WillReturnRows(rows)

	store := sqlstore.New(db)

	item, err := store.Job().Find(elemID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	if !reflect.DeepEqual(item, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], item)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT id, managerId, title, description, files, specialityId, experienceLevelId, paymentAmount, " +
			"country, city, jobTypeId FROM jobs WHERE").
		WithArgs(elemID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = store.Job().Find(elemID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// row scan error
	expect = []*model.Job{
		testJob(t, m),
	}

	mock.
		ExpectQuery("SELECT id, managerId, title, description, files, specialityId, experienceLevelId, paymentAmount, " +
			"country, city, jobTypeId FROM jobs WHERE").
		WithArgs(elemID).
		WillReturnRows(rows)

	_, err = store.Job().Find(elemID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestJobRepository_Edit(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}

	defer func() {
		mock.ExpectClose()
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	store := sqlstore.New(db)

	rows := sqlmock.
		NewRows([]string{"id"})

	var elemID int64 = 1
	expect := []*model.Job{
		{ ID: elemID },
	}

	for _, item := range expect {
		rows = rows.AddRow(item.ID)
	}

	u := testUser(t)
	u.ID = 1
	m := testManager(t, u)
	m.ID = 1
	j := testJob(t, m)

	// TODO: uncomment when validation will be implemented
	/*if err := j.Validate(); err != nil {
		t.Fatal()
	}*/

	//ok query
	j.City = "sim city"
	j.Country = "nnn"
	j.Description = "no description"

	mock.
		ExpectQuery(`UPDATE jobs SET`).
		WithArgs(j.Title, j.Description, j.Files, j.SpecialityId, j.ExperienceLevelId, j.PaymentAmount, j.Country,
			j.City, j.JobTypeId, j.ID).
		WillReturnRows(rows)

	err = store.Job().Edit(j)
	if err != nil {
		t.Fatal(err)
	}

	if j.ID != 1 {
		t.Errorf("bad id: want %v, have %v", j.ID, 1)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}