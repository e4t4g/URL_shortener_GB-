package usecase

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/e4t4g/URL_shortener_GB-/internal/app/usecase.UseCase -o ./use_case_mock_test.go -n UseCaseMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// UseCaseMock implements UseCase
type UseCaseMock struct {
	t minimock.Tester

	funcCreate          func(ctx context.Context, fullURL *URLData) (up1 *URLData, err error)
	inspectFuncCreate   func(ctx context.Context, fullURL *URLData)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mUseCaseMockCreate

	funcGetStat          func(ctx context.Context, id int) (up1 *URLData, err error)
	inspectFuncGetStat   func(ctx context.Context, id int)
	afterGetStatCounter  uint64
	beforeGetStatCounter uint64
	GetStatMock          mUseCaseMockGetStat

	funcRedirect          func(ctx context.Context, token string) (up1 *URLData, err error)
	inspectFuncRedirect   func(ctx context.Context, token string)
	afterRedirectCounter  uint64
	beforeRedirectCounter uint64
	RedirectMock          mUseCaseMockRedirect
}

// NewUseCaseMock returns a mock for UseCase
func NewUseCaseMock(t minimock.Tester) *UseCaseMock {
	m := &UseCaseMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mUseCaseMockCreate{mock: m}
	m.CreateMock.callArgs = []*UseCaseMockCreateParams{}

	m.GetStatMock = mUseCaseMockGetStat{mock: m}
	m.GetStatMock.callArgs = []*UseCaseMockGetStatParams{}

	m.RedirectMock = mUseCaseMockRedirect{mock: m}
	m.RedirectMock.callArgs = []*UseCaseMockRedirectParams{}

	return m
}

type mUseCaseMockCreate struct {
	mock               *UseCaseMock
	defaultExpectation *UseCaseMockCreateExpectation
	expectations       []*UseCaseMockCreateExpectation

	callArgs []*UseCaseMockCreateParams
	mutex    sync.RWMutex
}

// UseCaseMockCreateExpectation specifies expectation struct of the UseCase.Create
type UseCaseMockCreateExpectation struct {
	mock    *UseCaseMock
	params  *UseCaseMockCreateParams
	results *UseCaseMockCreateResults
	Counter uint64
}

// UseCaseMockCreateParams contains parameters of the UseCase.Create
type UseCaseMockCreateParams struct {
	ctx     context.Context
	fullURL *URLData
}

// UseCaseMockCreateResults contains results of the UseCase.Create
type UseCaseMockCreateResults struct {
	up1 *URLData
	err error
}

// Expect sets up expected params for UseCase.Create
func (mmCreate *mUseCaseMockCreate) Expect(ctx context.Context, fullURL *URLData) *mUseCaseMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UseCaseMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UseCaseMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &UseCaseMockCreateParams{ctx, fullURL}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the UseCase.Create
func (mmCreate *mUseCaseMockCreate) Inspect(f func(ctx context.Context, fullURL *URLData)) *mUseCaseMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for UseCaseMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by UseCase.Create
func (mmCreate *mUseCaseMockCreate) Return(up1 *URLData, err error) *UseCaseMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UseCaseMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UseCaseMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &UseCaseMockCreateResults{up1, err}
	return mmCreate.mock
}

//Set uses given function f to mock the UseCase.Create method
func (mmCreate *mUseCaseMockCreate) Set(f func(ctx context.Context, fullURL *URLData) (up1 *URLData, err error)) *UseCaseMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the UseCase.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the UseCase.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the UseCase.Create which will trigger the stat defined by the following
// Then helper
func (mmCreate *mUseCaseMockCreate) When(ctx context.Context, fullURL *URLData) *UseCaseMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UseCaseMock.Create mock is already set by Set")
	}

	expectation := &UseCaseMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &UseCaseMockCreateParams{ctx, fullURL},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up UseCase.Create return parameters for the expectation previously defined by the When method
func (e *UseCaseMockCreateExpectation) Then(up1 *URLData, err error) *UseCaseMock {
	e.results = &UseCaseMockCreateResults{up1, err}
	return e.mock
}

// Create implements UseCase
func (mmCreate *UseCaseMock) Create(ctx context.Context, fullURL *URLData) (up1 *URLData, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, fullURL)
	}

	mm_params := &UseCaseMockCreateParams{ctx, fullURL}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := UseCaseMockCreateParams{ctx, fullURL}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("UseCaseMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the UseCaseMock.Create")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, fullURL)
	}
	mmCreate.t.Fatalf("Unexpected call to UseCaseMock.Create. %v %v", ctx, fullURL)
	return
}

// CreateAfterCounter returns a count of finished UseCaseMock.Create invocations
func (mmCreate *UseCaseMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of UseCaseMock.Create invocations
func (mmCreate *UseCaseMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to UseCaseMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mUseCaseMockCreate) Calls() []*UseCaseMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*UseCaseMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *UseCaseMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *UseCaseMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UseCaseMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UseCaseMock.Create")
		} else {
			m.t.Errorf("Expected call to UseCaseMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to UseCaseMock.Create")
	}
}

type mUseCaseMockGetStat struct {
	mock               *UseCaseMock
	defaultExpectation *UseCaseMockGetStatExpectation
	expectations       []*UseCaseMockGetStatExpectation

	callArgs []*UseCaseMockGetStatParams
	mutex    sync.RWMutex
}

// UseCaseMockGetStatExpectation specifies expectation struct of the UseCase.GetStat
type UseCaseMockGetStatExpectation struct {
	mock    *UseCaseMock
	params  *UseCaseMockGetStatParams
	results *UseCaseMockGetStatResults
	Counter uint64
}

// UseCaseMockGetStatParams contains parameters of the UseCase.GetStat
type UseCaseMockGetStatParams struct {
	ctx context.Context
	id  int
}

// UseCaseMockGetStatResults contains results of the UseCase.GetStat
type UseCaseMockGetStatResults struct {
	up1 *URLData
	err error
}

// Expect sets up expected params for UseCase.GetStat
func (mmGetStat *mUseCaseMockGetStat) Expect(ctx context.Context, id int) *mUseCaseMockGetStat {
	if mmGetStat.mock.funcGetStat != nil {
		mmGetStat.mock.t.Fatalf("UseCaseMock.GetStat mock is already set by Set")
	}

	if mmGetStat.defaultExpectation == nil {
		mmGetStat.defaultExpectation = &UseCaseMockGetStatExpectation{}
	}

	mmGetStat.defaultExpectation.params = &UseCaseMockGetStatParams{ctx, id}
	for _, e := range mmGetStat.expectations {
		if minimock.Equal(e.params, mmGetStat.defaultExpectation.params) {
			mmGetStat.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetStat.defaultExpectation.params)
		}
	}

	return mmGetStat
}

// Inspect accepts an inspector function that has same arguments as the UseCase.GetStat
func (mmGetStat *mUseCaseMockGetStat) Inspect(f func(ctx context.Context, id int)) *mUseCaseMockGetStat {
	if mmGetStat.mock.inspectFuncGetStat != nil {
		mmGetStat.mock.t.Fatalf("Inspect function is already set for UseCaseMock.GetStat")
	}

	mmGetStat.mock.inspectFuncGetStat = f

	return mmGetStat
}

// Return sets up results that will be returned by UseCase.GetStat
func (mmGetStat *mUseCaseMockGetStat) Return(up1 *URLData, err error) *UseCaseMock {
	if mmGetStat.mock.funcGetStat != nil {
		mmGetStat.mock.t.Fatalf("UseCaseMock.GetStat mock is already set by Set")
	}

	if mmGetStat.defaultExpectation == nil {
		mmGetStat.defaultExpectation = &UseCaseMockGetStatExpectation{mock: mmGetStat.mock}
	}
	mmGetStat.defaultExpectation.results = &UseCaseMockGetStatResults{up1, err}
	return mmGetStat.mock
}

//Set uses given function f to mock the UseCase.GetStat method
func (mmGetStat *mUseCaseMockGetStat) Set(f func(ctx context.Context, id int) (up1 *URLData, err error)) *UseCaseMock {
	if mmGetStat.defaultExpectation != nil {
		mmGetStat.mock.t.Fatalf("Default expectation is already set for the UseCase.GetStat method")
	}

	if len(mmGetStat.expectations) > 0 {
		mmGetStat.mock.t.Fatalf("Some expectations are already set for the UseCase.GetStat method")
	}

	mmGetStat.mock.funcGetStat = f
	return mmGetStat.mock
}

// When sets expectation for the UseCase.GetStat which will trigger the stat defined by the following
// Then helper
func (mmGetStat *mUseCaseMockGetStat) When(ctx context.Context, id int) *UseCaseMockGetStatExpectation {
	if mmGetStat.mock.funcGetStat != nil {
		mmGetStat.mock.t.Fatalf("UseCaseMock.GetStat mock is already set by Set")
	}

	expectation := &UseCaseMockGetStatExpectation{
		mock:   mmGetStat.mock,
		params: &UseCaseMockGetStatParams{ctx, id},
	}
	mmGetStat.expectations = append(mmGetStat.expectations, expectation)
	return expectation
}

// Then sets up UseCase.GetStat return parameters for the expectation previously defined by the When method
func (e *UseCaseMockGetStatExpectation) Then(up1 *URLData, err error) *UseCaseMock {
	e.results = &UseCaseMockGetStatResults{up1, err}
	return e.mock
}

// GetStat implements UseCase
func (mmGetStat *UseCaseMock) GetStat(ctx context.Context, id int) (up1 *URLData, err error) {
	mm_atomic.AddUint64(&mmGetStat.beforeGetStatCounter, 1)
	defer mm_atomic.AddUint64(&mmGetStat.afterGetStatCounter, 1)

	if mmGetStat.inspectFuncGetStat != nil {
		mmGetStat.inspectFuncGetStat(ctx, id)
	}

	mm_params := &UseCaseMockGetStatParams{ctx, id}

	// Record call args
	mmGetStat.GetStatMock.mutex.Lock()
	mmGetStat.GetStatMock.callArgs = append(mmGetStat.GetStatMock.callArgs, mm_params)
	mmGetStat.GetStatMock.mutex.Unlock()

	for _, e := range mmGetStat.GetStatMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmGetStat.GetStatMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetStat.GetStatMock.defaultExpectation.Counter, 1)
		mm_want := mmGetStat.GetStatMock.defaultExpectation.params
		mm_got := UseCaseMockGetStatParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetStat.t.Errorf("UseCaseMock.GetStat got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetStat.GetStatMock.defaultExpectation.results
		if mm_results == nil {
			mmGetStat.t.Fatal("No results are set for the UseCaseMock.GetStat")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmGetStat.funcGetStat != nil {
		return mmGetStat.funcGetStat(ctx, id)
	}
	mmGetStat.t.Fatalf("Unexpected call to UseCaseMock.GetStat. %v %v", ctx, id)
	return
}

// GetStatAfterCounter returns a count of finished UseCaseMock.GetStat invocations
func (mmGetStat *UseCaseMock) GetStatAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetStat.afterGetStatCounter)
}

// GetStatBeforeCounter returns a count of UseCaseMock.GetStat invocations
func (mmGetStat *UseCaseMock) GetStatBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetStat.beforeGetStatCounter)
}

// Calls returns a list of arguments used in each call to UseCaseMock.GetStat.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetStat *mUseCaseMockGetStat) Calls() []*UseCaseMockGetStatParams {
	mmGetStat.mutex.RLock()

	argCopy := make([]*UseCaseMockGetStatParams, len(mmGetStat.callArgs))
	copy(argCopy, mmGetStat.callArgs)

	mmGetStat.mutex.RUnlock()

	return argCopy
}

// MinimockGetStatDone returns true if the count of the GetStat invocations corresponds
// the number of defined expectations
func (m *UseCaseMock) MinimockGetStatDone() bool {
	for _, e := range m.GetStatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetStatMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetStatCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetStat != nil && mm_atomic.LoadUint64(&m.afterGetStatCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetStatInspect logs each unmet expectation
func (m *UseCaseMock) MinimockGetStatInspect() {
	for _, e := range m.GetStatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UseCaseMock.GetStat with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetStatMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetStatCounter) < 1 {
		if m.GetStatMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UseCaseMock.GetStat")
		} else {
			m.t.Errorf("Expected call to UseCaseMock.GetStat with params: %#v", *m.GetStatMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetStat != nil && mm_atomic.LoadUint64(&m.afterGetStatCounter) < 1 {
		m.t.Error("Expected call to UseCaseMock.GetStat")
	}
}

type mUseCaseMockRedirect struct {
	mock               *UseCaseMock
	defaultExpectation *UseCaseMockRedirectExpectation
	expectations       []*UseCaseMockRedirectExpectation

	callArgs []*UseCaseMockRedirectParams
	mutex    sync.RWMutex
}

// UseCaseMockRedirectExpectation specifies expectation struct of the UseCase.Redirect
type UseCaseMockRedirectExpectation struct {
	mock    *UseCaseMock
	params  *UseCaseMockRedirectParams
	results *UseCaseMockRedirectResults
	Counter uint64
}

// UseCaseMockRedirectParams contains parameters of the UseCase.Redirect
type UseCaseMockRedirectParams struct {
	ctx   context.Context
	token string
}

// UseCaseMockRedirectResults contains results of the UseCase.Redirect
type UseCaseMockRedirectResults struct {
	up1 *URLData
	err error
}

// Expect sets up expected params for UseCase.Redirect
func (mmRedirect *mUseCaseMockRedirect) Expect(ctx context.Context, token string) *mUseCaseMockRedirect {
	if mmRedirect.mock.funcRedirect != nil {
		mmRedirect.mock.t.Fatalf("UseCaseMock.Redirect mock is already set by Set")
	}

	if mmRedirect.defaultExpectation == nil {
		mmRedirect.defaultExpectation = &UseCaseMockRedirectExpectation{}
	}

	mmRedirect.defaultExpectation.params = &UseCaseMockRedirectParams{ctx, token}
	for _, e := range mmRedirect.expectations {
		if minimock.Equal(e.params, mmRedirect.defaultExpectation.params) {
			mmRedirect.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmRedirect.defaultExpectation.params)
		}
	}

	return mmRedirect
}

// Inspect accepts an inspector function that has same arguments as the UseCase.Redirect
func (mmRedirect *mUseCaseMockRedirect) Inspect(f func(ctx context.Context, token string)) *mUseCaseMockRedirect {
	if mmRedirect.mock.inspectFuncRedirect != nil {
		mmRedirect.mock.t.Fatalf("Inspect function is already set for UseCaseMock.Redirect")
	}

	mmRedirect.mock.inspectFuncRedirect = f

	return mmRedirect
}

// Return sets up results that will be returned by UseCase.Redirect
func (mmRedirect *mUseCaseMockRedirect) Return(up1 *URLData, err error) *UseCaseMock {
	if mmRedirect.mock.funcRedirect != nil {
		mmRedirect.mock.t.Fatalf("UseCaseMock.Redirect mock is already set by Set")
	}

	if mmRedirect.defaultExpectation == nil {
		mmRedirect.defaultExpectation = &UseCaseMockRedirectExpectation{mock: mmRedirect.mock}
	}
	mmRedirect.defaultExpectation.results = &UseCaseMockRedirectResults{up1, err}
	return mmRedirect.mock
}

//Set uses given function f to mock the UseCase.Redirect method
func (mmRedirect *mUseCaseMockRedirect) Set(f func(ctx context.Context, token string) (up1 *URLData, err error)) *UseCaseMock {
	if mmRedirect.defaultExpectation != nil {
		mmRedirect.mock.t.Fatalf("Default expectation is already set for the UseCase.Redirect method")
	}

	if len(mmRedirect.expectations) > 0 {
		mmRedirect.mock.t.Fatalf("Some expectations are already set for the UseCase.Redirect method")
	}

	mmRedirect.mock.funcRedirect = f
	return mmRedirect.mock
}

// When sets expectation for the UseCase.Redirect which will trigger the stat defined by the following
// Then helper
func (mmRedirect *mUseCaseMockRedirect) When(ctx context.Context, token string) *UseCaseMockRedirectExpectation {
	if mmRedirect.mock.funcRedirect != nil {
		mmRedirect.mock.t.Fatalf("UseCaseMock.Redirect mock is already set by Set")
	}

	expectation := &UseCaseMockRedirectExpectation{
		mock:   mmRedirect.mock,
		params: &UseCaseMockRedirectParams{ctx, token},
	}
	mmRedirect.expectations = append(mmRedirect.expectations, expectation)
	return expectation
}

// Then sets up UseCase.Redirect return parameters for the expectation previously defined by the When method
func (e *UseCaseMockRedirectExpectation) Then(up1 *URLData, err error) *UseCaseMock {
	e.results = &UseCaseMockRedirectResults{up1, err}
	return e.mock
}

// Redirect implements UseCase
func (mmRedirect *UseCaseMock) Redirect(ctx context.Context, token string) (up1 *URLData, err error) {
	mm_atomic.AddUint64(&mmRedirect.beforeRedirectCounter, 1)
	defer mm_atomic.AddUint64(&mmRedirect.afterRedirectCounter, 1)

	if mmRedirect.inspectFuncRedirect != nil {
		mmRedirect.inspectFuncRedirect(ctx, token)
	}

	mm_params := &UseCaseMockRedirectParams{ctx, token}

	// Record call args
	mmRedirect.RedirectMock.mutex.Lock()
	mmRedirect.RedirectMock.callArgs = append(mmRedirect.RedirectMock.callArgs, mm_params)
	mmRedirect.RedirectMock.mutex.Unlock()

	for _, e := range mmRedirect.RedirectMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmRedirect.RedirectMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRedirect.RedirectMock.defaultExpectation.Counter, 1)
		mm_want := mmRedirect.RedirectMock.defaultExpectation.params
		mm_got := UseCaseMockRedirectParams{ctx, token}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmRedirect.t.Errorf("UseCaseMock.Redirect got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmRedirect.RedirectMock.defaultExpectation.results
		if mm_results == nil {
			mmRedirect.t.Fatal("No results are set for the UseCaseMock.Redirect")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmRedirect.funcRedirect != nil {
		return mmRedirect.funcRedirect(ctx, token)
	}
	mmRedirect.t.Fatalf("Unexpected call to UseCaseMock.Redirect. %v %v", ctx, token)
	return
}

// RedirectAfterCounter returns a count of finished UseCaseMock.Redirect invocations
func (mmRedirect *UseCaseMock) RedirectAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRedirect.afterRedirectCounter)
}

// RedirectBeforeCounter returns a count of UseCaseMock.Redirect invocations
func (mmRedirect *UseCaseMock) RedirectBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRedirect.beforeRedirectCounter)
}

// Calls returns a list of arguments used in each call to UseCaseMock.Redirect.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmRedirect *mUseCaseMockRedirect) Calls() []*UseCaseMockRedirectParams {
	mmRedirect.mutex.RLock()

	argCopy := make([]*UseCaseMockRedirectParams, len(mmRedirect.callArgs))
	copy(argCopy, mmRedirect.callArgs)

	mmRedirect.mutex.RUnlock()

	return argCopy
}

// MinimockRedirectDone returns true if the count of the Redirect invocations corresponds
// the number of defined expectations
func (m *UseCaseMock) MinimockRedirectDone() bool {
	for _, e := range m.RedirectMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RedirectMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRedirectCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRedirect != nil && mm_atomic.LoadUint64(&m.afterRedirectCounter) < 1 {
		return false
	}
	return true
}

// MinimockRedirectInspect logs each unmet expectation
func (m *UseCaseMock) MinimockRedirectInspect() {
	for _, e := range m.RedirectMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UseCaseMock.Redirect with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RedirectMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRedirectCounter) < 1 {
		if m.RedirectMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UseCaseMock.Redirect")
		} else {
			m.t.Errorf("Expected call to UseCaseMock.Redirect with params: %#v", *m.RedirectMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRedirect != nil && mm_atomic.LoadUint64(&m.afterRedirectCounter) < 1 {
		m.t.Error("Expected call to UseCaseMock.Redirect")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UseCaseMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateInspect()

		m.MinimockGetStatInspect()

		m.MinimockRedirectInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UseCaseMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *UseCaseMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockGetStatDone() &&
		m.MinimockRedirectDone()
}