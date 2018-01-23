package main

import "testing"
import "fmt"

func TestBuildUnitList(t *testing.T) {
	var rowPartition *PartitionType = BuildPartition(ROWS_IN_QUADRANT, Row_ids)
	var colPartition *PartitionType = BuildPartition(COLS_IN_QUADRANT, Col_ids)
	if i := len(*BuildUnitList(rowPartition, colPartition)); i != 27 {
		t.Error(fmt.Sprintf("TestBuildUnitList - expected 27 but got %d", i))
	}
}

func TestPartition(t *testing.T) {
	var partition *PartitionType
	partition = BuildPartition(3, "abcdef")
	if (*partition)[0] != "abc" || (*partition)[1] != "def" {
		t.Error("TestPartition - didnt get abc or def")
	}

	if len(*partition) != 2 {
		t.Error(fmt.Sprintf("TestPartition - expected 2 but got %d", len(*partition)))
	}

	var rowPartition *PartitionType = BuildPartition(ROWS_IN_QUADRANT, Row_ids)
	var colPartition *PartitionType = BuildPartition(COLS_IN_QUADRANT, Col_ids)

	if (*rowPartition)[0] != "ABC" || (*rowPartition)[1] != "DEF" || (*rowPartition)[2] != "GHI" {
		t.Error("TestPartition - Bad rowPartitions.")
	}
	if (*colPartition)[0] != "123" || (*colPartition)[1] != "456" || (*colPartition)[2] != "789" {
		t.Error("TestPartition - Bad colPartitions.")
	}
}

func TestUnits(t *testing.T) {
	var ul UnitListType
	for cell, _ := range *Cells {
		ul = (*Units)[cell]
		if len(ul) != 3 {
			t.Error(fmt.Sprintf("TestUnits - Bad unit count."))
		}
	}
	c2_ul := (*Units)["C2"]
	c2_ul_0 := c2_ul[0]
	c2_ul_1 := c2_ul[1]
	c2_ul_2 := c2_ul[2]
	if !c2_ul_0["A2"] || !c2_ul_0["B2"] || !c2_ul_0["C2"] || !c2_ul_0["D2"] || !c2_ul_0["E2"] || !c2_ul_0["F2"] || !c2_ul_0["G2"] || !c2_ul_0["H2"] || !c2_ul_0["I2"] {
		t.Error("TestUnits - Bad A2 unit", c2_ul_0)
	}
	if !c2_ul_1["C1"] || !c2_ul_1["C2"] || !c2_ul_1["C3"] || !c2_ul_1["C4"] || !c2_ul_1["C5"] || !c2_ul_1["C6"] || !c2_ul_1["C7"] || !c2_ul_1["C8"] || !c2_ul_1["C9"] {
		t.Error("TestUnits - Bad C1 unit", c2_ul_1)
	}
	if !c2_ul_2["A1"] || !c2_ul_2["A2"] || !c2_ul_2["A3"] || !c2_ul_2["B1"] || !c2_ul_2["B2"] || !c2_ul_2["B3"] || !c2_ul_2["C1"] || !c2_ul_2["C2"] || !c2_ul_2["C3"] {
		t.Error("TestUnits - Bad A1 unit", c2_ul_2)
	}

}

func TestPeers(t *testing.T) {
	var unit UnitType
	for cell, _ := range *Cells {
		unit = (*Peers)[cell]
		if len(unit) != 20 {
			t.Error("TestPeers - bad length ", cell, len(unit))
		}
	}
	unit = (*Peers)["C2"]
	if !unit["A2"] || !unit["B2"] || !unit["D2"] || !unit["E2"] || !unit["F2"] || !unit["G2"] || !unit["H2"] || !unit["I2"] || !unit["C1"] || !unit["C3"] || !unit["C4"] || !unit["C5"] || !unit["C6"] || !unit["C7"] || !unit["C8"] || !unit["C9"] || !unit["A1"] || !unit["A3"] || !unit["B1"] || !unit["B3"] {
		t.Error("TestPeers - Bad C2")
	}
}

func TestGetUnsolvedValues(t *testing.T) {
	var values = Values{"A1": "1"}
	var fact = getUnsolvedFact(&values)
	if fact != nil {
		t.Error("TestGetUnsolvedValues - Test failure ", fact)
	}
	values = Values{"A1": "12"}
	fact = getUnsolvedFact(&values)
	if fact == nil {
		t.Error("TestGetUnsolvedValues - Test failure ")
	}
	if fact.cell != "A1" {
		t.Error("TestGetUnsolvedValues - Bad cell ")
	}
	if fact.domain != "12" {
		t.Error("TestGetUnsolvedValues - Bad domain ")
	}
}

func TestHasConflictInUnit(t *testing.T) {

	var values = make(Values)
	var unit = make(UnitType)

	values["A1"] = "5"
	values["A2"] = "5"
	unit["A1"] = true
	unit["A2"] = true
	if !hasConflictInUnit(&values, &unit) {
		t.Error("TestHasConflictInUnit - test 1 failed ")
	}
	values["A1"] = "5"
	values["A2"] = "6"
	unit["A1"] = true
	unit["A2"] = true
	if hasConflictInUnit(&values, &unit) {
		t.Error("TestHasConflictInUnit - test 2 failed ")
	}

}

func TestHasConflictInUnitList(t *testing.T) {
	var values = make(Values, 0)
	var unitList = make(UnitListType, 2)

	values["A1"] = "5"
	values["A2"] = "6"
	values["A3"] = "6"
	values["A4"] = "7"
	values["A5"] = "5"

	var unit1 = UnitType{"A1": true, "A2": true}
	var unit2 = UnitType{"A2": true, "A3": true}
	unitList[0] = unit1
	unitList[1] = unit2
	if !hasConflictInUnitList(&values, &unitList) {
		t.Error("TestHasConflictInUnitList - test 1 failed ")
	}

	unit1 = UnitType{"A1": true, "A2": true}
	unit2 = UnitType{"A3": true, "A4": true}
	var unit3 = UnitType{"A4": true, "A5": true, "A1": true}

	unitList = make(UnitListType, 3)
	unitList[0] = unit1
	unitList[1] = unit2
	unitList[2] = unit3

	if !hasConflictInUnitList(&values, &unitList) {
		t.Error("TestHasConflictInUnitList - test 2 failed ")
	}

	unitList = make(UnitListType, 3)
	unit3 = UnitType{"A4": true, "A5": true}
	unitList[0] = unit1
	unitList[1] = unit2
	unitList[2] = unit3

	if hasConflictInUnitList(&values, &unitList) {
		t.Error("TestHasConflictInUnitList - test 3 failed ")
	}
}

func TestWillItConflict(t *testing.T) {

	var values = make(Values, 0)
	values["A1"] = "5"
	values["A2"] = "6"
	values["A3"] = "3"

	var cell string = "A1"
	var value string = "6"

	if !willItConflict(&values, &cell, &value) {
		t.Error("TestWillItConflict - test 1 failed ")
	}

	values = make(Values, 0)
	values["A1"] = "5"
	values["A2"] = "4"
	values["A3"] = "3"

	if willItConflict(&values, &cell, &value) {
		t.Error("TestWillItConflict - test 2 failed ")
	}

	values = make(Values, 0)
	values["A1"] = "5"
	values["A2"] = "4"
	values["A3"] = "6"

	if !willItConflict(&values, &cell, &value) {
		t.Error("TestWillItConflict - test 3 failed ")
	}
}

func TestMakeFacts(t *testing.T) {
}

func TestSolveGrid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"1", "000000000000000000000000000000000000000000000000000000000000000000000000000000000", "123456789456789123789123456214365897365897214897214365531642978642978531978531642"},
		{"2", "920705003600000080004600000006080120040301070071060500000004800010000002400902061", "928745613657213489134698257396587124845321976271469538762134895519876342483952761"},
		{"3", "003010002850649700070002000016080070204701609030020810000500020009273086300060900", "493817562852649731671352498916485273284731659735926814168594327549273186327168945"},
		{"4", "020007000609000008000950200035000070407000809080000120001034000700000602000100030", "128467395659312748374958261235891476417623859986745123561234987743589612892176534"},
		{"5", "600000084003060000001000502100074000720906035000320008305000200000050900240000007", "652719384483265791971438562138574629724986135569321478395847216817652943246193857"},
		{"6", "800000000003600000070090200050007000000045700000100030001000068008500010090000400", "812753649943682175675491283154237896369845721287169534521974368438526917796318452"},
		{"7", "800000000003600000070090200050007000000045700000100030001000068008500010090000403", ""},
		{"8", "800000000003600000070090200050007000000045700000100030001000068008500010090000404", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := SolveGrid(&tt.input)
			fmt.Printf("answer: %s\n", *ans)
			if *ans != tt.expected {
				t.Errorf("Started with %s and expected %s", tt.input, *ans)
			}
		})
	}
}
