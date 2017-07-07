// Code generated by mockery v1.0.1 DO NOT EDIT.

package nlmock

import mock "github.com/stretchr/testify/mock"
import net "net"
import nl "github.com/quilt/quilt/minion/nl"

// I is an autogenerated mock type for the I type
type I struct {
	mock.Mock
}

// AddVeth provides a mock function with given fields: name, peer, mtu
func (_m *I) AddVeth(name string, peer string, mtu int) error {
	ret := _m.Called(name, peer, mtu)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int) error); ok {
		r0 = rf(name, peer, mtu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddrAdd provides a mock function with given fields: link, ip
func (_m *I) AddrAdd(link nl.Link, ip net.IPNet) error {
	ret := _m.Called(link, ip)

	var r0 error
	if rf, ok := ret.Get(0).(func(nl.Link, net.IPNet) error); ok {
		r0 = rf(link, ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkByIndex provides a mock function with given fields: index
func (_m *I) LinkByIndex(index int) (nl.Link, error) {
	ret := _m.Called(index)

	var r0 nl.Link
	if rf, ok := ret.Get(0).(func(int) nl.Link); ok {
		r0 = rf(index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(nl.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LinkByName provides a mock function with given fields: name
func (_m *I) LinkByName(name string) (nl.Link, error) {
	ret := _m.Called(name)

	var r0 nl.Link
	if rf, ok := ret.Get(0).(func(string) nl.Link); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(nl.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LinkDel provides a mock function with given fields: link
func (_m *I) LinkDel(link nl.Link) error {
	ret := _m.Called(link)

	var r0 error
	if rf, ok := ret.Get(0).(func(nl.Link) error); ok {
		r0 = rf(link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetUp provides a mock function with given fields: link
func (_m *I) LinkSetUp(link nl.Link) error {
	ret := _m.Called(link)

	var r0 error
	if rf, ok := ret.Get(0).(func(nl.Link) error); ok {
		r0 = rf(link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RouteList provides a mock function with given fields: family
func (_m *I) RouteList(family int) ([]nl.Route, error) {
	ret := _m.Called(family)

	var r0 []nl.Route
	if rf, ok := ret.Get(0).(func(int) []nl.Route); ok {
		r0 = rf(family)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]nl.Route)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(family)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
