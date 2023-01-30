package main

import (
	"fmt"
	"sort"
)

type Replicas_profile struct {
	replicas_num        int
	CPU_reservation     float64
	CPU_aspiration      float64
	RAM_reservation     float64
	RAM_aspiration      float64
	latency_reservation float64
	latency_aspiration  float64
}

const (
	maxCPUPerPod = 2000.0
	maxRAMPerPod = 400.0
	maxLATPerPod = 10.0
)

var (
	curentCPU = 9100.0
	curentRAM = 1700.0
	curentLAT = 43.0

	CPUReservationPerPod = 1500.0
	RAMReservationPerPod = 300.0
	LATReservationPerPod = 8.0
	CPUAspirationPerPod  = 200.0
	RAMAspirationPerPod  = 100.0
	LATAspirationPerPod  = 3.0

	rp1 Replicas_profile
	rp2 Replicas_profile
	rp3 Replicas_profile
	rp4 Replicas_profile
	rp5 Replicas_profile
	rp6 Replicas_profile
	rp7 Replicas_profile
	rp8 Replicas_profile

	new_replicas_num int
)

func main() {
	// fmt.Println("Hello World!")

	// normalize params
	curentCPUNorm := NormalizeCPU(curentCPU)
	curentRAMNorm := NormalizeRAM(curentRAM)
	curentLATNorm := NormalizeLAT(curentLAT)

	CPUReservationPerPodNorm := NormalizeCPU(CPUReservationPerPod)
	RAMReservationPerPodNorm := NormalizeRAM(RAMReservationPerPod)
	LATReservationPerPodNorm := NormalizeLAT(LATReservationPerPod)
	CPUAspirationPerPodNorm := NormalizeCPU(CPUAspirationPerPod)
	RAMAspirationPerPodNorm := NormalizeRAM(RAMAspirationPerPod)
	LATAspirationPerPodNorm := NormalizeLAT(LATAspirationPerPod)

	// set replicas profiles R and A levels
	SetProfilesParams(&rp1, 1, CPUReservationPerPodNorm, RAMReservationPerPodNorm, LATReservationPerPodNorm, CPUAspirationPerPodNorm, RAMAspirationPerPodNorm, LATAspirationPerPodNorm)
	SetProfilesParams(&rp2, 2, CPUReservationPerPodNorm, RAMReservationPerPodNorm, LATReservationPerPodNorm, CPUAspirationPerPodNorm, RAMAspirationPerPodNorm, LATAspirationPerPodNorm)
	SetProfilesParams(&rp3, 3, CPUReservationPerPodNorm, RAMReservationPerPodNorm, LATReservationPerPodNorm, CPUAspirationPerPodNorm, RAMAspirationPerPodNorm, LATAspirationPerPodNorm)
	SetProfilesParams(&rp4, 4, CPUReservationPerPodNorm, RAMReservationPerPodNorm, LATReservationPerPodNorm, CPUAspirationPerPodNorm, RAMAspirationPerPodNorm, LATAspirationPerPodNorm)
	SetProfilesParams(&rp5, 5, CPUReservationPerPodNorm, RAMReservationPerPodNorm, LATReservationPerPodNorm, CPUAspirationPerPodNorm, RAMAspirationPerPodNorm, LATAspirationPerPodNorm)
	SetProfilesParams(&rp6, 6, CPUReservationPerPodNorm, RAMReservationPerPodNorm, LATReservationPerPodNorm, CPUAspirationPerPodNorm, RAMAspirationPerPodNorm, LATAspirationPerPodNorm)
	SetProfilesParams(&rp7, 7, CPUReservationPerPodNorm, RAMReservationPerPodNorm, LATReservationPerPodNorm, CPUAspirationPerPodNorm, RAMAspirationPerPodNorm, LATAspirationPerPodNorm)
	SetProfilesParams(&rp8, 8, CPUReservationPerPodNorm, RAMReservationPerPodNorm, LATReservationPerPodNorm, CPUAspirationPerPodNorm, RAMAspirationPerPodNorm, LATAspirationPerPodNorm)

	Replicas_profiles := [8]Replicas_profile{rp1, rp2, rp3, rp4, rp5, rp6, rp7, rp8}

	// for _, val := range Replicas_profiles {
	// 	fmt.Println(val.replicas_num)
	// 	fmt.Println(val.CPU_reservation)
	// 	fmt.Println(val.RAM_reservation)
	// 	fmt.Println(val.latency_reservation)
	// 	fmt.Println(val.CPU_aspiration)
	// 	fmt.Println(val.RAM_aspiration)
	// 	fmt.Println(val.latency_aspiration)
	// 	fmt.Println("--------------")
	// }

	// fmt.Println("")

	var tmp_prefer_min = []float64{}
	var prefer_max = make(map[int]float64)

	// findPreferMaxMap
	for _, val := range Replicas_profiles {
		tmp_CPU := (val.CPU_reservation - curentCPUNorm) / (val.CPU_reservation - val.CPU_aspiration)
		tmp_RAM := (val.RAM_reservation - curentRAMNorm) / (val.RAM_reservation - val.RAM_aspiration)
		tmp_latency := (val.latency_reservation - curentLATNorm) / (val.latency_reservation - val.latency_aspiration)

		// fmt.Println("CPU ", tmp_CPU, " RAM ", tmp_RAM, " LAT ", tmp_latency)

		// check if feasible
		if tmp_CPU >= 0 && tmp_RAM >= 0 && tmp_latency >= 0 {
			tmp_prefer_min = append(tmp_prefer_min, tmp_CPU, tmp_RAM, tmp_latency)
		}

		if len(tmp_prefer_min) > 0 {
			sort.Slice(tmp_prefer_min, func(i, j int) bool {
				return tmp_prefer_min[i] < tmp_prefer_min[j]
			})

			prefer_max[val.replicas_num] = tmp_prefer_min[0]

			tmp_prefer_min = nil
		}
	}

	fmt.Printf("Prefer MAX Map: %v\n\n", prefer_max)

	// getReplicasNum
	keys := make([]int, 0, len(prefer_max))
	for key := range prefer_max {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return prefer_max[keys[i]] < prefer_max[keys[j]]
	})

	// fmt.Printf("%v\n\n", keys)

	new_replicas_num = keys[0]

	fmt.Printf("New Replicas num: %v", new_replicas_num)
}

func SetProfilesParams(rp *Replicas_profile, rpN int, CPURsv float64, RAMRsv float64, LATRsv float64, CPUAsp float64, RAMAsp float64, LATAsp float64) {
	rp.replicas_num = rpN

	rp.CPU_reservation = CPURsv * float64(rpN)
	rp.RAM_reservation = RAMRsv * float64(rpN)
	rp.latency_reservation = LATRsv * float64(rpN)

	rp.CPU_aspiration = CPUAsp * float64(rpN)
	rp.RAM_aspiration = RAMAsp * float64(rpN)
	rp.latency_aspiration = LATAsp * float64(rpN)
}

func NormalizeCPU(CPU float64) float64 {
	normalizedCPU := CPU / maxCPUPerPod

	return normalizedCPU
}

func NormalizeRAM(RAM float64) float64 {
	normalizedRAM := RAM / maxRAMPerPod

	return normalizedRAM
}

func NormalizeLAT(LAT float64) float64 {
	normalizedLAT := LAT / maxLATPerPod

	return normalizedLAT
}
