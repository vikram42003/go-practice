package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unitsMap := map[string]int{
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }
	return unitsMap    
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	u, exists := units[unit]
    if !exists {
        return false
    }

    bill[item] += u
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	u, exists := units[unit]
    i, exists1 := bill[item]
    if !exists || !exists1 || i - u < 0 {
        return false
    }

    if i - u == 0 {
        delete(bill, item)
    } else {
        bill[item] -= u
    }
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, exists := bill[item]
    if !exists {
        return 0, false
    }
    return qty, true
}
