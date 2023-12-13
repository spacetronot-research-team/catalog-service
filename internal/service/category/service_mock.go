// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=service_mock.go -package=category
//
// Package category is a generated GoMock package.
package category

import (
	context "context"
	reflect "reflect"

	entity "github.com/spacetronot-research-team/catalog-service/internal/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockpersistenceRepository is a mock of persistenceRepository interface.
type MockpersistenceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockpersistenceRepositoryMockRecorder
}

// MockpersistenceRepositoryMockRecorder is the mock recorder for MockpersistenceRepository.
type MockpersistenceRepositoryMockRecorder struct {
	mock *MockpersistenceRepository
}

// NewMockpersistenceRepository creates a new mock instance.
func NewMockpersistenceRepository(ctrl *gomock.Controller) *MockpersistenceRepository {
	mock := &MockpersistenceRepository{ctrl: ctrl}
	mock.recorder = &MockpersistenceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpersistenceRepository) EXPECT() *MockpersistenceRepositoryMockRecorder {
	return m.recorder
}

// DeleteCategory mocks base method.
func (m *MockpersistenceRepository) DeleteCategory(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategory", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategory indicates an expected call of DeleteCategory.
func (mr *MockpersistenceRepositoryMockRecorder) DeleteCategory(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockpersistenceRepository)(nil).DeleteCategory), ctx, id)
}

// FindCategoryByID mocks base method.
func (m *MockpersistenceRepository) FindCategoryByID(ctx context.Context, id int64) (entity.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCategoryByID", ctx, id)
	ret0, _ := ret[0].(entity.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCategoryByID indicates an expected call of FindCategoryByID.
func (mr *MockpersistenceRepositoryMockRecorder) FindCategoryByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCategoryByID", reflect.TypeOf((*MockpersistenceRepository)(nil).FindCategoryByID), ctx, id)
}

// FindCategoryByName mocks base method.
func (m *MockpersistenceRepository) FindCategoryByName(ctx context.Context, name string) (entity.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCategoryByName", ctx, name)
	ret0, _ := ret[0].(entity.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCategoryByName indicates an expected call of FindCategoryByName.
func (mr *MockpersistenceRepositoryMockRecorder) FindCategoryByName(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCategoryByName", reflect.TypeOf((*MockpersistenceRepository)(nil).FindCategoryByName), ctx, name)
}

// InsertCategory mocks base method.
func (m *MockpersistenceRepository) InsertCategory(ctx context.Context, category entity.Category) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertCategory", ctx, category)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertCategory indicates an expected call of InsertCategory.
func (mr *MockpersistenceRepositoryMockRecorder) InsertCategory(ctx, category any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCategory", reflect.TypeOf((*MockpersistenceRepository)(nil).InsertCategory), ctx, category)
}

// UpdateCategory mocks base method.
func (m *MockpersistenceRepository) UpdateCategory(ctx context.Context, category entity.Category) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategory", ctx, category)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCategory indicates an expected call of UpdateCategory.
func (mr *MockpersistenceRepositoryMockRecorder) UpdateCategory(ctx, category any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategory", reflect.TypeOf((*MockpersistenceRepository)(nil).UpdateCategory), ctx, category)
}

// MockcacheRepository is a mock of cacheRepository interface.
type MockcacheRepository struct {
	ctrl     *gomock.Controller
	recorder *MockcacheRepositoryMockRecorder
}

// MockcacheRepositoryMockRecorder is the mock recorder for MockcacheRepository.
type MockcacheRepositoryMockRecorder struct {
	mock *MockcacheRepository
}

// NewMockcacheRepository creates a new mock instance.
func NewMockcacheRepository(ctrl *gomock.Controller) *MockcacheRepository {
	mock := &MockcacheRepository{ctrl: ctrl}
	mock.recorder = &MockcacheRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcacheRepository) EXPECT() *MockcacheRepositoryMockRecorder {
	return m.recorder
}

// FindCategoryByID mocks base method.
func (m *MockcacheRepository) FindCategoryByID(ctx context.Context, id int64) (entity.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCategoryByID", ctx, id)
	ret0, _ := ret[0].(entity.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCategoryByID indicates an expected call of FindCategoryByID.
func (mr *MockcacheRepositoryMockRecorder) FindCategoryByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCategoryByID", reflect.TypeOf((*MockcacheRepository)(nil).FindCategoryByID), ctx, id)
}

// FindCategoryByName mocks base method.
func (m *MockcacheRepository) FindCategoryByName(ctx context.Context, name string) (entity.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCategoryByName", ctx, name)
	ret0, _ := ret[0].(entity.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCategoryByName indicates an expected call of FindCategoryByName.
func (mr *MockcacheRepositoryMockRecorder) FindCategoryByName(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCategoryByName", reflect.TypeOf((*MockcacheRepository)(nil).FindCategoryByName), ctx, name)
}

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateCategory mocks base method.
func (m *MockService) CreateCategory(ctx context.Context, req CreateCategoryRequest) (CreateCategoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", ctx, req)
	ret0, _ := ret[0].(CreateCategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockServiceMockRecorder) CreateCategory(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockService)(nil).CreateCategory), ctx, req)
}
