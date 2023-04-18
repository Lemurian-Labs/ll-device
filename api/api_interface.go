package api

type IDeviceConfiguration interface {
	// resource configuration
	InitDevice() int
}

type IDeviceContext interface {
	Create() int
	Delete() int
}
