package main

import "sort"

var CellIndex *[]string
var Cells *UnitType
var UnitList *UnitListType
var Units *UnitListMapType
var Peers *UnitMapType

func init() {
	Cells = Cross(Row_ids, Col_ids)

	CellIndex = createCellIndex(Cells)

	var rowPartition *PartitionType = BuildPartition(ROWS_IN_QUADRANT, Row_ids)
	var colPartition *PartitionType = BuildPartition(COLS_IN_QUADRANT, Col_ids)
	UnitList = BuildUnitList(rowPartition, colPartition)

	Units = buildUnits()

	Peers = buildPeers()

}

func Cross(s1 string, s2 string) *UnitType {
	var unit = make(UnitType, 0)
	for _, c1 := range s1 {
		for _, c2 := range s2 {
			unit[string(c1)+string(c2)] = true
		}
	}
	return &unit
}

func BuildPartition(size int, s1 string) *PartitionType {
	var partition = make(PartitionType, 0, len(s1)/size)
	for i := 0; i < len(s1); i = i + size {
		partition = append(partition, s1[i:i+size])
	}
	return &partition
}

func BuildUnitList(rowPartition *PartitionType, colPartition *PartitionType) *UnitListType {
	var ul = make(UnitListType, 0, len(Col_ids)+len(Row_ids)+(len(*rowPartition)*len(*rowPartition)))
	var unit *UnitType
	for _, c := range Col_ids {
		unit = Cross(Row_ids, string(c))
		ul = append(ul, *unit)
	}
	for _, c := range Row_ids {
		unit = Cross(string(c), Col_ids)
		ul = append(ul, *unit)
	}
	for _, s1 := range *rowPartition {
		for _, s2 := range *colPartition {
			unit = Cross(s1, s2)
			ul = append(ul, *unit)
		}
	}
	return &ul
}
func createCellIndex(cells *UnitType) *[]string {
	index := make([]string, 0, len(*cells))
	for key, _ := range *cells {
		index = append(index, key)
	}
	sort.Strings(index)
	return &index
}

func buildUnits() *UnitListMapType {
	var ul UnitListType
	var units = make(UnitListMapType, 0)
	for c, _ := range *Cells {
		ul = make(UnitListType, 0, 3)
		units[c] = ul
		for _, a_ul := range *UnitList {
			if a_ul[c] {
				units[c] = append(units[c], a_ul)
			}
		}
	}
	return &units
}
func buildPeers() *UnitMapType {
	var ul UnitListType
	var peers = make(UnitMapType, 0)
	for c, _ := range *Cells {
		ul = (*Units)[c]
		var u = make(UnitType, 0)
		for _, unit := range ul {
			for cell, _ := range unit {
				u[cell] = true
			}
		}
		delete(u, c)
		peers[c] = u
	}
	return &peers
}
