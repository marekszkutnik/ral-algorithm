package main

import (
	"fmt"
	"sort"
)

type ReplicasProfile struct {
	replicasNum    int
	CPUReservation float64
	CPUAspiration  float64
	RAMReservation float64
	RAMAspiration  float64
	LATReservation float64
	LATAspiration  float64
}

const (
	maxCPUPerPod = 2000.0
	maxRAMPerPod = 400.0
	maxLATPerPod = 10.0
)

var (
	curentCPU = 4400.0
	curentRAM = 900.0
	curentLAT = 25.0

	CPUReservationPerPod = 1500.0
	RAMReservationPerPod = 300.0
	LATReservationPerPod = 8.0
	CPUAspirationPerPod  = 200.0
	RAMAspirationPerPod  = 100.0
	LATAspirationPerPod  = 3.0

	rp1  ReplicasProfile
	rp2  ReplicasProfile
	rp3  ReplicasProfile
	rp4  ReplicasProfile
	rp5  ReplicasProfile
	rp6  ReplicasProfile
	rp7  ReplicasProfile
	rp8  ReplicasProfile
	rp9  ReplicasProfile
	rp10 ReplicasProfile

	newReplicasNum int
)

func main() {
	SetProfilesWithNormalizedParams(CPUReservationPerPod, RAMReservationPerPod, LATReservationPerPod, CPUAspirationPerPod, RAMAspirationPerPod, LATAspirationPerPod)
	numberOfReplicas := CalculateDesiredReplicasbyRALAlgorithm(curentCPU, curentRAM, curentLAT)

	fmt.Printf("New Replicas num: %v", float64(numberOfReplicas))
}

func CalculateDesiredReplicasbyRALAlgorithm(curentCPUVal float64, curentRAMVal float64, curentLATVal float64) (numberOfReplicas int) {
	// normalize params
	curentCPUNorm := NormalizeCPU(curentCPUVal)
	curentRAMNorm := NormalizeRAM(curentRAMVal)
	curentLATNorm := NormalizeLAT(curentLATVal)

	ReplicasProfiles := [10]ReplicasProfile{rp1, rp2, rp3, rp4, rp5, rp6, rp7, rp8, rp9, rp10}

	var tmp_prefer_min = []float64{}
	var prefer_max = make(map[int]float64)

	// findPreferMaxMap
	for _, val := range ReplicasProfiles {
		tmp_CPU := (val.CPUReservation - curentCPUNorm) / (val.CPUReservation - val.CPUAspiration)
		tmp_RAM := (val.RAMReservation - curentRAMNorm) / (val.RAMReservation - val.RAMAspiration)
		tmp_latency := (val.LATReservation - curentLATNorm) / (val.LATReservation - val.LATAspiration)

		// check if feasible
		if tmp_CPU >= 0 && tmp_RAM >= 0 && tmp_latency >= 0 {
			tmp_prefer_min = append(tmp_prefer_min, tmp_CPU, tmp_RAM, tmp_latency)
		}

		if len(tmp_prefer_min) > 0 {
			sort.Slice(tmp_prefer_min, func(i, j int) bool {
				return tmp_prefer_min[i] < tmp_prefer_min[j]
			})

			prefer_max[val.replicasNum] = tmp_prefer_min[0]

			tmp_prefer_min = nil
		}
	}

	if len(prefer_max) > 0 {
		fmt.Printf("Prefer MAX Map: %v\n\n", prefer_max)

		// getReplicasNum
		keys := make([]int, 0, len(prefer_max))
		for key := range prefer_max {
			keys = append(keys, key)
		}

		sort.SliceStable(keys, func(i, j int) bool {
			return prefer_max[keys[i]] < prefer_max[keys[j]]
		})

		newReplicasNum = keys[0]
	} else {
		newReplicasNum = len(ReplicasProfiles)
	}

	// fmt.Printf("New Replicas num: %v", new_replicasNum)
	return newReplicasNum
}

func SetProfilesWithNormalizedParams(CPUReservationPerPodVal float64, RAMReservationPerPodVal float64, LATReservationPerPodVal float64, CPUAspirationPerPodVal float64, RAMAspirationPerPodVal float64, LATAspirationPerPodVal float64) {
	CPUReservationPerPodNorm := NormalizeCPU(CPUReservationPerPodVal)
	RAMReservationPerPodNorm := NormalizeRAM(RAMReservationPerPodVal)
	LATReservationPerPodNorm := NormalizeLAT(LATReservationPerPodVal)
	CPUAspirationPerPodNorm := NormalizeCPU(CPUAspirationPerPodVal)
	RAMAspirationPerPodNorm := NormalizeRAM(RAMAspirationPerPodVal)
	LATAspirationPerPodNorm := NormalizeLAT(LATAspirationPerPodVal)

	// set replicas profiles R and A levels
	SetProfilesParams(&rp1, 1, CPUReservationPerPodNorm, RAMReservationPerPodNorm, LATReservationPerPodNorm, CPUAspirationPerPodNorm, RAMAspirationPerPodNorm, LATAspirationPerPodNorm)
	SetProfilesParams(&rp2, 2, CPUReservationPerPodNorm, RAMReservationPerPodNorm, LATReservationPerPodNorm, CPUAspirationPerPodNorm, RAMAspirationPerPodNorm, LATAspirationPerPodNorm)
	SetProfilesParams(&rp3, 3, CPUReservationPerPodNorm, RAMReservationPerPodNorm, LATReservationPerPodNorm, CPUAspirationPerPodNorm, RAMAspirationPerPodNorm, LATAspirationPerPodNorm)
	SetProfilesParams(&rp4, 4, CPUReservationPerPodNorm, RAMReservationPerPodNorm, LATReservationPerPodNorm, CPUAspirationPerPodNorm, RAMAspirationPerPodNorm, LATAspirationPerPodNorm)
	SetProfilesParams(&rp5, 5, CPUReservationPerPodNorm, RAMReservationPerPodNorm, LATReservationPerPodNorm, CPUAspirationPerPodNorm, RAMAspirationPerPodNorm, LATAspirationPerPodNorm)
	SetProfilesParams(&rp6, 6, CPUReservationPerPodNorm, RAMReservationPerPodNorm, LATReservationPerPodNorm, CPUAspirationPerPodNorm, RAMAspirationPerPodNorm, LATAspirationPerPodNorm)
	SetProfilesParams(&rp7, 7, CPUReservationPerPodNorm, RAMReservationPerPodNorm, LATReservationPerPodNorm, CPUAspirationPerPodNorm, RAMAspirationPerPodNorm, LATAspirationPerPodNorm)
	SetProfilesParams(&rp8, 8, CPUReservationPerPodNorm, RAMReservationPerPodNorm, LATReservationPerPodNorm, CPUAspirationPerPodNorm, RAMAspirationPerPodNorm, LATAspirationPerPodNorm)
	SetProfilesParams(&rp9, 9, CPUReservationPerPodNorm, RAMReservationPerPodNorm, LATReservationPerPodNorm, CPUAspirationPerPodNorm, RAMAspirationPerPodNorm, LATAspirationPerPodNorm)
	SetProfilesParams(&rp10, 10, CPUReservationPerPodNorm, RAMReservationPerPodNorm, LATReservationPerPodNorm, CPUAspirationPerPodNorm, RAMAspirationPerPodNorm, LATAspirationPerPodNorm)
}

func SetProfilesParams(rp *ReplicasProfile, rpN int, CPURsv float64, RAMRsv float64, LATRsv float64, CPUAsp float64, RAMAsp float64, LATAsp float64) {
	rp.replicasNum = rpN

	rp.CPUReservation = CPURsv * float64(rpN)
	rp.RAMReservation = RAMRsv * float64(rpN)
	rp.LATReservation = LATRsv * float64(rpN)

	rp.CPUAspiration = CPUAsp * float64(rpN)
	rp.RAMAspiration = RAMAsp * float64(rpN)
	rp.LATAspiration = LATAsp * float64(rpN)
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
