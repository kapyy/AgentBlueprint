package main

import apm_util "golang-client/apm_examples"

func main() {
	funcNodes := []uint64{
		110000001,
	}
	apm_util.WriteGenericMethodAPM(funcNodes, "default_agent.apm")
}
