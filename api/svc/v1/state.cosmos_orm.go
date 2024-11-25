// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package svcv1

import (
	context "context"
	ormlist "cosmossdk.io/orm/model/ormlist"
	ormtable "cosmossdk.io/orm/model/ormtable"
	ormerrors "cosmossdk.io/orm/types/ormerrors"
)

type MetadataTable interface {
	Insert(ctx context.Context, metadata *Metadata) error
	InsertReturningId(ctx context.Context, metadata *Metadata) (uint64, error)
	LastInsertedSequence(ctx context.Context) (uint64, error)
	Update(ctx context.Context, metadata *Metadata) error
	Save(ctx context.Context, metadata *Metadata) error
	Delete(ctx context.Context, metadata *Metadata) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id uint64) (*Metadata, error)
	HasByOrigin(ctx context.Context, origin string) (found bool, err error)
	// GetByOrigin returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByOrigin(ctx context.Context, origin string) (*Metadata, error)
	List(ctx context.Context, prefixKey MetadataIndexKey, opts ...ormlist.Option) (MetadataIterator, error)
	ListRange(ctx context.Context, from, to MetadataIndexKey, opts ...ormlist.Option) (MetadataIterator, error)
	DeleteBy(ctx context.Context, prefixKey MetadataIndexKey) error
	DeleteRange(ctx context.Context, from, to MetadataIndexKey) error

	doNotImplement()
}

type MetadataIterator struct {
	ormtable.Iterator
}

func (i MetadataIterator) Value() (*Metadata, error) {
	var metadata Metadata
	err := i.UnmarshalMessage(&metadata)
	return &metadata, err
}

type MetadataIndexKey interface {
	id() uint32
	values() []interface{}
	metadataIndexKey()
}

// primary key starting index..
type MetadataPrimaryKey = MetadataIdIndexKey

type MetadataIdIndexKey struct {
	vs []interface{}
}

func (x MetadataIdIndexKey) id() uint32            { return 0 }
func (x MetadataIdIndexKey) values() []interface{} { return x.vs }
func (x MetadataIdIndexKey) metadataIndexKey()     {}

func (this MetadataIdIndexKey) WithId(id uint64) MetadataIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type MetadataOriginIndexKey struct {
	vs []interface{}
}

func (x MetadataOriginIndexKey) id() uint32            { return 1 }
func (x MetadataOriginIndexKey) values() []interface{} { return x.vs }
func (x MetadataOriginIndexKey) metadataIndexKey()     {}

func (this MetadataOriginIndexKey) WithOrigin(origin string) MetadataOriginIndexKey {
	this.vs = []interface{}{origin}
	return this
}

type metadataTable struct {
	table ormtable.AutoIncrementTable
}

func (this metadataTable) Insert(ctx context.Context, metadata *Metadata) error {
	return this.table.Insert(ctx, metadata)
}

func (this metadataTable) Update(ctx context.Context, metadata *Metadata) error {
	return this.table.Update(ctx, metadata)
}

func (this metadataTable) Save(ctx context.Context, metadata *Metadata) error {
	return this.table.Save(ctx, metadata)
}

func (this metadataTable) Delete(ctx context.Context, metadata *Metadata) error {
	return this.table.Delete(ctx, metadata)
}

func (this metadataTable) InsertReturningId(ctx context.Context, metadata *Metadata) (uint64, error) {
	return this.table.InsertReturningPKey(ctx, metadata)
}

func (this metadataTable) LastInsertedSequence(ctx context.Context) (uint64, error) {
	return this.table.LastInsertedSequence(ctx)
}

func (this metadataTable) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this metadataTable) Get(ctx context.Context, id uint64) (*Metadata, error) {
	var metadata Metadata
	found, err := this.table.PrimaryKey().Get(ctx, &metadata, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &metadata, nil
}

func (this metadataTable) HasByOrigin(ctx context.Context, origin string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		origin,
	)
}

func (this metadataTable) GetByOrigin(ctx context.Context, origin string) (*Metadata, error) {
	var metadata Metadata
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &metadata,
		origin,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &metadata, nil
}

func (this metadataTable) List(ctx context.Context, prefixKey MetadataIndexKey, opts ...ormlist.Option) (MetadataIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return MetadataIterator{it}, err
}

func (this metadataTable) ListRange(ctx context.Context, from, to MetadataIndexKey, opts ...ormlist.Option) (MetadataIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return MetadataIterator{it}, err
}

func (this metadataTable) DeleteBy(ctx context.Context, prefixKey MetadataIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this metadataTable) DeleteRange(ctx context.Context, from, to MetadataIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this metadataTable) doNotImplement() {}

var _ MetadataTable = metadataTable{}

func NewMetadataTable(db ormtable.Schema) (MetadataTable, error) {
	table := db.GetTable(&Metadata{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Metadata{}).ProtoReflect().Descriptor().FullName()))
	}
	return metadataTable{table.(ormtable.AutoIncrementTable)}, nil
}

type ProfileTable interface {
	Insert(ctx context.Context, profile *Profile) error
	Update(ctx context.Context, profile *Profile) error
	Save(ctx context.Context, profile *Profile) error
	Delete(ctx context.Context, profile *Profile) error
	Has(ctx context.Context, id string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id string) (*Profile, error)
	HasBySubjectOrigin(ctx context.Context, subject string, origin string) (found bool, err error)
	// GetBySubjectOrigin returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetBySubjectOrigin(ctx context.Context, subject string, origin string) (*Profile, error)
	List(ctx context.Context, prefixKey ProfileIndexKey, opts ...ormlist.Option) (ProfileIterator, error)
	ListRange(ctx context.Context, from, to ProfileIndexKey, opts ...ormlist.Option) (ProfileIterator, error)
	DeleteBy(ctx context.Context, prefixKey ProfileIndexKey) error
	DeleteRange(ctx context.Context, from, to ProfileIndexKey) error

	doNotImplement()
}

type ProfileIterator struct {
	ormtable.Iterator
}

func (i ProfileIterator) Value() (*Profile, error) {
	var profile Profile
	err := i.UnmarshalMessage(&profile)
	return &profile, err
}

type ProfileIndexKey interface {
	id() uint32
	values() []interface{}
	profileIndexKey()
}

// primary key starting index..
type ProfilePrimaryKey = ProfileIdIndexKey

type ProfileIdIndexKey struct {
	vs []interface{}
}

func (x ProfileIdIndexKey) id() uint32            { return 0 }
func (x ProfileIdIndexKey) values() []interface{} { return x.vs }
func (x ProfileIdIndexKey) profileIndexKey()      {}

func (this ProfileIdIndexKey) WithId(id string) ProfileIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type ProfileSubjectOriginIndexKey struct {
	vs []interface{}
}

func (x ProfileSubjectOriginIndexKey) id() uint32            { return 1 }
func (x ProfileSubjectOriginIndexKey) values() []interface{} { return x.vs }
func (x ProfileSubjectOriginIndexKey) profileIndexKey()      {}

func (this ProfileSubjectOriginIndexKey) WithSubject(subject string) ProfileSubjectOriginIndexKey {
	this.vs = []interface{}{subject}
	return this
}

func (this ProfileSubjectOriginIndexKey) WithSubjectOrigin(subject string, origin string) ProfileSubjectOriginIndexKey {
	this.vs = []interface{}{subject, origin}
	return this
}

type profileTable struct {
	table ormtable.Table
}

func (this profileTable) Insert(ctx context.Context, profile *Profile) error {
	return this.table.Insert(ctx, profile)
}

func (this profileTable) Update(ctx context.Context, profile *Profile) error {
	return this.table.Update(ctx, profile)
}

func (this profileTable) Save(ctx context.Context, profile *Profile) error {
	return this.table.Save(ctx, profile)
}

func (this profileTable) Delete(ctx context.Context, profile *Profile) error {
	return this.table.Delete(ctx, profile)
}

func (this profileTable) Has(ctx context.Context, id string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this profileTable) Get(ctx context.Context, id string) (*Profile, error) {
	var profile Profile
	found, err := this.table.PrimaryKey().Get(ctx, &profile, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &profile, nil
}

func (this profileTable) HasBySubjectOrigin(ctx context.Context, subject string, origin string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		subject,
		origin,
	)
}

func (this profileTable) GetBySubjectOrigin(ctx context.Context, subject string, origin string) (*Profile, error) {
	var profile Profile
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &profile,
		subject,
		origin,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &profile, nil
}

func (this profileTable) List(ctx context.Context, prefixKey ProfileIndexKey, opts ...ormlist.Option) (ProfileIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return ProfileIterator{it}, err
}

func (this profileTable) ListRange(ctx context.Context, from, to ProfileIndexKey, opts ...ormlist.Option) (ProfileIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return ProfileIterator{it}, err
}

func (this profileTable) DeleteBy(ctx context.Context, prefixKey ProfileIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this profileTable) DeleteRange(ctx context.Context, from, to ProfileIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this profileTable) doNotImplement() {}

var _ ProfileTable = profileTable{}

func NewProfileTable(db ormtable.Schema) (ProfileTable, error) {
	table := db.GetTable(&Profile{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Profile{}).ProtoReflect().Descriptor().FullName()))
	}
	return profileTable{table}, nil
}

type StateStore interface {
	MetadataTable() MetadataTable
	ProfileTable() ProfileTable

	doNotImplement()
}

type stateStore struct {
	metadata MetadataTable
	profile  ProfileTable
}

func (x stateStore) MetadataTable() MetadataTable {
	return x.metadata
}

func (x stateStore) ProfileTable() ProfileTable {
	return x.profile
}

func (stateStore) doNotImplement() {}

var _ StateStore = stateStore{}

func NewStateStore(db ormtable.Schema) (StateStore, error) {
	metadataTable, err := NewMetadataTable(db)
	if err != nil {
		return nil, err
	}

	profileTable, err := NewProfileTable(db)
	if err != nil {
		return nil, err
	}

	return stateStore{
		metadataTable,
		profileTable,
	}, nil
}
