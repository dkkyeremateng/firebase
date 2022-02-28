package store

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

var (
	// ErrNotFound is used when a specific Product is requested but does not exist.
	ErrNotFound = errors.New("document not found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidID = errors.New("ID is not in its proper form")

	// ErrDocsNotFound occurs when a documents could not be found on firbase
	ErrDocsNotFound = errors.New("error getting documents snapshots")

	// ErrForbidden occurs when a user tries to do something that is forbidden to them according to our access control policies.
	ErrForbidden = errors.New("attempted action is not allowed")
)

// New returns a new instance of firestore client
func New(ctx context.Context, app *firebase.App) (*firestore.Client, error) {

	fc, err := app.Firestore(ctx)
	if err != nil {
		return nil, errors.New("error getting a new firestore client instance")
	}

	return fc, nil
}

// FindOneByField returns a document by field
func FindOneByField(ctx context.Context, client firestore.Client, c, f, op string, v interface{}) (*firestore.DocumentSnapshot, error) {

	ds, err := client.Collection(c).Where(f, op, v).Limit(1).Documents(ctx).GetAll()
	if err != nil {
		return nil, ErrDocsNotFound
	}

	if len(ds) <= 0 {
		return nil, ErrNotFound
	}

	return ds[0], nil
}

// FindOneByTwoFields returns a document by field
func FindOneByTwoFields(ctx context.Context, client firestore.Client, c, ff, fop string, fv interface{}, sf, sop string, sv interface{}) (*firestore.DocumentSnapshot, error) {

	ds, err := client.Collection(c).Where(ff, fop, fv).Where(sf, sop, sv).Limit(1).Documents(ctx).GetAll()
	if err != nil {
		return nil, ErrDocsNotFound
	}

	if len(ds) <= 0 {
		return nil, ErrNotFound
	}

	return ds[0], nil
}

// Delete removes a document by id
func Delete(ctx context.Context, client firestore.Client, ref *firestore.DocumentRef) error {

	if _, err := ref.Delete(ctx); err != nil {
		return err
	}

	return nil
}

// FindAllByField returns a document by field
func FindAllByField(ctx context.Context, client firestore.Client, c, f, op string, v interface{}) ([]*firestore.DocumentSnapshot, error) {

	ds, err := client.Collection(c).Where(f, op, v).Documents(ctx).GetAll()
	if err != nil {
		return nil, ErrDocsNotFound
	}

	if len(ds) <= 0 {
		return nil, ErrNotFound
	}

	return ds, nil
}

// FindAllByFieldAndOrder returns all documents by field in order
func FindAllByFieldAndOrder(ctx context.Context, client firestore.Client, c, f, op string, v interface{}, p string, dir firestore.Direction) ([]*firestore.DocumentSnapshot, error) {

	ds, err := client.Collection(c).Where(f, op, v).OrderBy(p, dir).Documents(ctx).GetAll()
	if err != nil {
		fmt.Println(err)
		return nil, ErrDocsNotFound
	}

	if len(ds) <= 0 {
		return nil, ErrNotFound
	}

	return ds, nil
}

// FindAllByTwoFields returns a document by field
func FindAllByTwoFields(ctx context.Context, client firestore.Client, c string, ff string, fop string, fv interface{}, sf string, sop string, sv interface{}) ([]*firestore.DocumentSnapshot, error) {

	ds, err := client.Collection(c).Where(ff, fop, fv).Where(sf, sop, sv).Documents(ctx).GetAll()
	if err != nil {
		return nil, ErrDocsNotFound
	}

	if len(ds) <= 0 {
		return nil, ErrNotFound
	}

	return ds, nil
}

// FindFromArray returns a documents by field
func FindFromArray(ctx context.Context, client firestore.Client, c, f, v string) ([]*firestore.DocumentSnapshot, error) {

	ds, err := client.Collection(c).Where(f, "array-contains", v).Documents(ctx).GetAll()
	if err != nil {
		return nil, ErrDocsNotFound
	}

	if len(ds) <= 0 {
		return nil, ErrNotFound
	}

	return ds, nil
}

// GetAll returns all documents in a colloctions
func GetAll(ctx context.Context, client firestore.Client, c string) ([]*firestore.DocumentSnapshot, error) {

	ds, err := client.Collection(c).Documents(ctx).GetAll()
	if err != nil {
		return nil, ErrDocsNotFound
	}

	if len(ds) <= 0 {
		return nil, ErrNotFound
	}

	return ds, nil
}

// GetAllByOrder returns all documents in a colloctions
func GetAllByOrder(ctx context.Context, client firestore.Client, c, p string, dir firestore.Direction) ([]*firestore.DocumentSnapshot, error) {

	ds, err := client.Collection(c).OrderBy(p, dir).Documents(ctx).GetAll()
	if err != nil {
		return nil, ErrDocsNotFound
	}

	if len(ds) <= 0 {
		return nil, ErrNotFound
	}

	return ds, nil
}

// Add adds a new document to a collection
func Add(ctx context.Context, client firestore.Client, c string, d interface{}) (*firestore.DocumentSnapshot, error) {

	dRef, _, err := client.Collection(c).Add(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("adding document: %w", err)
	}

	ds, err := dRef.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting document snapshot from firebase: %w", err)
	}

	return ds, nil
}

// Update updates a document ref and return updated ref
func Update(ctx context.Context, client firestore.Client, df *firestore.DocumentRef, data interface{}) error {

	if _, err := df.Set(ctx, data); err != nil {
		return fmt.Errorf("updating document: %w", err)
	}

	return nil
}
