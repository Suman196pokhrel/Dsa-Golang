package advancedgo

import "fmt"

// STRUCT COMPOSITION - a simple way of achieving code reuse by building new types using existing types. sub concepts like Promotion & embedding are utilized fo this.
// PROMOTION - promotion is a feature in go that allow fields and methods of an embedded struct to be accessed as if they were part of the outer struct.
// EMBEDDING - Embedding is the process of including one struct type inside another struct type. The fields and methods of the embedded struct become part of the outer struct.

// USECASE
// Suppose you have different types of employees: regular employees,
// managers, and executives. They all share some common attributes like Name and Age, but each type has its own unique attributes and behaviors.

// Person represents a person with name and age
type Person struct {
	Name string
	Age  int
}

// Employee represents a generic employee, embedding Person
type Employee struct {
	Person // Embedded struct
	ID     int
	Role   string
}

// Manager represents a manager, embedding Employee
type Manager struct {
	Employee   // Embedded struct
	Department string
}

// Executive represents an executive, embedding Employee
type Executive struct {
	Employee     // Embedded struct
	StockOptions int
}

// Method specific to Employee
func (e Employee) PrintDetails() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %d\n", e.Name, e.Age, e.ID)
}

func TestPromotionFeature() {
	// Create instances of different types of employees
	employee := Employee{
		Person: Person{Name: "John", Age: 30},
		ID:     1001,
		Role:   "Software Engineer",
	}

	manager := Manager{
		Employee:   Employee{Person: Person{Name: "Alice", Age: 40}, ID: 2001, Role: "Manager"},
		Department: "Engineering",
	}

	executive := Executive{
		Employee:     Employee{Person: Person{Name: "Bob", Age: 50}, ID: 3001, Role: "CEO"},
		StockOptions: 10000,
	}

	// Access fields from embedded structs
	fmt.Println("Employee:", employee.Name, "Role:", employee.Role)
	fmt.Println("Manager:", manager.Name, "Role:", manager.Role, "Department:", manager.Department)
	fmt.Println("Executive:", executive.Name, "Role:", executive.Role, "Stock Options:", executive.StockOptions)

	// Update fields of the embedded struct
	employee.Age = 31
	fmt.Println("Updated Age of Employee:", employee.Age)
}

// BUT - EMBEDDING IS NOT INHERITANCE
// strict type checking allows only similar types of variable assignment i.e, varible of type manager cannot be assigned to anyother type
// there is no dynamic dispatch for concrete types in Go. The methods on the embedded field have no idea they are embedded.
