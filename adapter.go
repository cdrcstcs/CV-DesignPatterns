package main

import "fmt"

// Volt struct represents a voltage
type Volt struct {
	volts int
}

// Socket struct represents a socket
type Socket struct{}

// getVolt method returns a Volt instance with a default value of 120 volts
func (s *Socket) getVolt() Volt {
	return Volt{volts: 120}
}

// SocketAdapter interface defines methods for obtaining voltages of 120V, 12V, and 3V
type SocketAdapter interface {
	get120Volt() Volt
	get12Volt() Volt
	get3Volt() Volt
}

// SocketClassAdapterImpl struct is a class adapter implementing the SocketAdapter interface
type SocketClassAdapterImpl struct {
	Socket Socket
}

// get120Volt method returns 120 volts
func (a *SocketClassAdapterImpl) get120Volt() Volt {
	return a.Socket.getVolt()
}

// get12Volt method returns 12 volts
func (a *SocketClassAdapterImpl) get12Volt() Volt {
	v := a.Socket.getVolt()
	return a.convertVolt(v, 10)
}

// get3Volt method returns 3 volts
func (a *SocketClassAdapterImpl) get3Volt() Volt {
	v := a.Socket.getVolt()
	return a.convertVolt(v, 40)
}

// convertVolt method converts the voltage
func (a *SocketClassAdapterImpl) convertVolt(v Volt, i int) Volt {
	return Volt{volts: v.volts / i}
}

// SocketObjectAdapterImpl struct is an object adapter implementing the SocketAdapter interface
type SocketObjectAdapterImpl struct {
	sock Socket
}

// get120Volt method returns 120 volts
func (a *SocketObjectAdapterImpl) get120Volt() Volt {
	return a.sock.getVolt()
}

// get12Volt method returns 12 volts
func (a *SocketObjectAdapterImpl) get12Volt() Volt {
	v := a.sock.getVolt()
	return a.convertVolt(v, 10)
}

// get3Volt method returns 3 volts
func (a *SocketObjectAdapterImpl) get3Volt() Volt {
	v := a.sock.getVolt()
	return a.convertVolt(v, 40)
}

// convertVolt method converts the voltage
func (a *SocketObjectAdapterImpl) convertVolt(v Volt, i int) Volt {
	return Volt{volts: v.volts / i}
}

// main function to test the adapters
func main() {
	testClassAdapter()
	testObjectAdapter()
}

// testClassAdapter function tests the class adapter
func testClassAdapter() {
	sockAdapter := &SocketClassAdapterImpl{}
	v3 := getVolt(sockAdapter, 3)
	v12 := getVolt(sockAdapter, 12)
	v120 := getVolt(sockAdapter, 120)

	fmt.Printf("v3 volts using Class Adapter=%d\n", v3.volts)
	fmt.Printf("v12 volts using Class Adapter=%d\n", v12.volts)
	fmt.Printf("v120 volts using Class Adapter=%d\n", v120.volts)
}

// testObjectAdapter function tests the object adapter
func testObjectAdapter() {
	sockAdapter := &SocketObjectAdapterImpl{}
	v3 := getVolt(sockAdapter, 3)
	v12 := getVolt(sockAdapter, 12)
	v120 := getVolt(sockAdapter, 120)

	fmt.Printf("v3 volts using Object Adapter=%d\n", v3.volts)
	fmt.Printf("v12 volts using Object Adapter=%d\n", v12.volts)
	fmt.Printf("v120 volts using Object Adapter=%d\n", v120.volts)
}

// getVolt function returns the desired voltage based on the case
func getVolt(sockAdapter SocketAdapter, i int) Volt {
	switch i {
	case 3:
		return sockAdapter.get3Volt()
	case 12:
		return sockAdapter.get12Volt()
	case 120:
		return sockAdapter.get120Volt()
	default:
		return sockAdapter.get120Volt()
	}
}
