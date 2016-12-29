package cite

import mock "github.com/stretchr/testify/mock"
import url "net/url"

// MockResource is an autogenerated mock type for the Resource type
type MockResource struct {
	mock.Mock
}

// Cite provides a mock function with given fields:
func (_m *MockResource) Cite() Citation {
	ret := _m.Called()

	var r0 Citation
	if rf, ok := ret.Get(0).(func() Citation); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Citation)
	}

	return r0
}

// Lines provides a mock function with given fields:
func (_m *MockResource) Lines() LinePredicate {
	ret := _m.Called()

	var r0 LinePredicate
	if rf, ok := ret.Get(0).(func() LinePredicate); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(LinePredicate)
		}
	}

	return r0
}

// URL provides a mock function with given fields:
func (_m *MockResource) URL() *url.URL {
	ret := _m.Called()

	var r0 *url.URL
	if rf, ok := ret.Get(0).(func() *url.URL); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*url.URL)
		}
	}

	return r0
}
