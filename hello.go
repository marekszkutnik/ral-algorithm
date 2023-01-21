package main

import (
	"fmt"
	"sort"
)

type Replicas_profile struct {
	replicas_num        int
	CPU_reservation     float32
	CPU_aspiration      float32
	RAM_reservation     float32
	RAM_aspiration      float32
	latency_reservation float32
	latency_aspiration  float32
}

func main() {
	// fmt.Println("Hello World!")

	var CPU float32
	var RAM float32
	var latency float32

	var new_replicas_num int

	CPU = 0.5
	RAM = 0.5
	latency = 0.5

	var rp1 Replicas_profile
	var rp2 Replicas_profile
	var rp3 Replicas_profile
	var rp4 Replicas_profile

	rp1.replicas_num = 1
	rp1.CPU_reservation = 0.4
	rp1.CPU_aspiration = 0.3
	rp1.RAM_reservation = 0.4
	rp1.RAM_aspiration = 0.3
	rp1.latency_reservation = 0.4
	rp1.latency_aspiration = 0.3

	rp2.replicas_num = 2
	rp2.CPU_reservation = 0.6
	rp2.CPU_aspiration = 0.4
	rp2.RAM_reservation = 0.6
	rp2.RAM_aspiration = 0.4
	rp2.latency_reservation = 0.6
	rp2.latency_aspiration = 0.4

	rp3.replicas_num = 3
	rp3.CPU_reservation = 0.7
	rp3.CPU_aspiration = 0.3
	rp3.RAM_reservation = 0.7
	rp3.RAM_aspiration = 0.3
	rp3.latency_reservation = 0.7
	rp3.latency_aspiration = 0.5

	rp4.replicas_num = 4
	rp4.CPU_reservation = 0.95
	rp4.CPU_aspiration = 0.6
	rp4.RAM_reservation = 0.95
	rp4.RAM_aspiration = 0.6
	rp4.latency_reservation = 0.95
	rp4.latency_aspiration = 0

	// var prefer_min = []float32{}
	// var prefer_max = []float32{}

	var Replicas_profiles = [4]Replicas_profile{rp1, rp2, rp3, rp4}

	// fmt.Println("Hello World!", CPU, RAM, latency)
	// fmt.Println("MIN", prefer_min)
	// fmt.Println("MAX", prefer_max)
	// fmt.Println("RP", Replicas_profiles)

	var tmp_prefer_min = []float32{}
	var tmp_prefer_max = []float32{}

	var prefer_max = make(map[int]float32)

	// findPreferMaxMap
	for _, val := range Replicas_profiles {
		tmp_CPU := (val.CPU_reservation - CPU) / (val.CPU_reservation - val.CPU_aspiration)
		tmp_RAM := (val.RAM_reservation - RAM) / (val.RAM_reservation - val.RAM_aspiration)
		tmp_latency := (val.latency_reservation - latency) / (val.latency_reservation - val.latency_aspiration)

		/*
			prefer_min = append(prefer_min, tmp_CPU, tmp_RAM, tmp_latency)

			sort.Slice(prefer_min, func(i, j int) bool {
				return prefer_min[i] < prefer_min[j]
			})

			prefer_max = append(prefer_max, prefer_min[0])

			prefer_min = nil

			fmt.Printf("%v\n", prefer_max)
		*/

		tmp_prefer_min := append(tmp_prefer_min, tmp_CPU, tmp_RAM, tmp_latency)

		sort.Slice(tmp_prefer_min, func(i, j int) bool {
			return tmp_prefer_min[i] < tmp_prefer_min[j]
		})

		tmp_prefer_max = append(tmp_prefer_max, tmp_prefer_min[0])

		// fmt.Printf("%v\n", tmp_prefer_max)

		prefer_max[val.replicas_num] = tmp_prefer_max[0]

		tmp_prefer_min = nil
		tmp_prefer_max = nil
	}

	fmt.Printf("Prefer MAX Map: %v\n\n", prefer_max)

	// getReplicasNum
	keys := make([]int, 0, len(prefer_max))
	for key := range prefer_max {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return prefer_max[keys[i]] > prefer_max[keys[j]]
	})

	// fmt.Printf("%v\n\n", keys)

	new_replicas_num = keys[0]

	fmt.Printf("New Replicas num: %v", new_replicas_num)
}
