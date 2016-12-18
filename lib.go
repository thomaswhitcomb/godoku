package main

import "strings"

type Values map[string]string
type Fact struct {
	cell   string
	domain string
}
type FactList []*Fact

func initializeValues(cells *UnitType) *Values {
	var values = make(Values, 0)
	for c, _ := range *cells {
		values[c] = COL_DOMAIN
	}
	return &values
}

func makeFacts(grid *string, cellIndex *[]string) *FactList {
	var factList = make(FactList, 0, len(Col_ids)*len(Row_ids))
	for i, key := range *cellIndex {
		if (*grid)[i] != '0' {
			fact := Fact{key, string((*grid)[i])}
			factList = append(factList, &fact)
		}
	}
	return &factList
}
func getUnsolvedFact(values *Values) *Fact {
	for _, key := range *CellIndex() {
		value, exists := (*values)[key] // this is cool.  double return checks for existence
		if exists && len(value) != 1 {
			return &Fact{key, value}
		}
	}
	return nil
}
func valuesToGrid(values *Values) *string {
	var grid = ""
	for _, key := range *CellIndex() {
		v := (*values)[key]
		if len(v) == 1 {
			grid = grid + v
		} else {
			grid = grid + "{" + v + "}"
		}
	}
	return &grid
}
func cloneValues(values *Values) *Values {
	var clone = make(Values, 0)
	for k, v := range *values {
		clone[k] = v
	}
	return &clone
}

func assign(values *Values, cell *string, value *string) *Values {
	var emptyValues *Values
	if willItConflict(values, cell, value) {
		var x = make(Values, 0)
		return &x
	}
	(*values)[*cell] = *value
	var domain string
	var unit = (*Peers())[*cell]
	for c, _ := range unit {
		domain = (*values)[c]
		if strings.Contains(domain, *value) {
			domain = strings.Replace(domain, *value, "", 1)
			(*values)[c] = domain
			if len(domain) == 1 {
				if emptyValues = assign(values, &c, &domain); len(*emptyValues) == 0 {
					return emptyValues
				}
			}
		}
	}
	return values
}

func hasConflictInUnit(values *Values, unit *UnitType) bool {
	var domain = ""
	var value string
	for c, _ := range *unit {
		value = (*values)[c]
		if len(value) == 1 {
			if strings.Contains(domain, value) {
				return true
			}
			domain = domain + value
		}
	}
	return false
}

func hasConflictInUnitList(values *Values, unitList *UnitListType) bool {
	for _, unit := range *unitList {
		if hasConflictInUnit(values, &unit) {
			return true
		}
	}
	return false
}

func willItConflict(values *Values, cell *string, value *string) bool {

	var localValues = cloneValues(values)
	(*localValues)[*cell] = *value
	var unitList = (*Units())[*cell]
	return hasConflictInUnitList(localValues, &unitList)

}

func backward(values *Values) *Values {
	var fact = getUnsolvedFact(values)
	if fact == nil {
		return values
	}
	var clone *Values
	var solution *Values
	for _, c := range fact.domain {
		clone = cloneValues(values)
		var cs = string(c)
		if clone = assign(clone, &fact.cell, &cs); len(*clone) != 0 {
			solution = backward(clone)
			if len(*solution) > 0 {
				return solution
			}
		}
	}
	return &Values{}
}

func forward(factList *FactList, values *Values) *Values {
	var emptyValues *Values
	for _, fact := range *factList {
		if emptyValues = assign(values, &fact.cell, &fact.domain); len(*emptyValues) == 0 {
			return emptyValues
		}
	}
	return values
}

func SolveGrid(grid *string) *string {
	values := initializeValues(Cells())
	factList := makeFacts(grid, CellIndex())
	values = forward(factList, values)
	if len(*values) == 0 {
		return new(string)
	}
	var newvalues = backward(values)
	return valuesToGrid(newvalues)
}
