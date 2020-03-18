package main

import "demochain/core"

func main() {
	bc := core.NewBlockchain()
	bc.SendData("send 1 BTC to Jack")
	bc.SendData("send 1 EOS to Jack")
	bc.Print()
}
