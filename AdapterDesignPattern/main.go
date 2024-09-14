package main

func main() {
	macComputer := &Mac{}
	macComputer.InsertIntoLighteningPort()

	windowsComputer := &Windows{}
	windowAdapter := &WindowsAdaptor{WindowMachine: windowsComputer}
	windowAdapter.InsertIntoLighteningPort()
}
