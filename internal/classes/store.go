package classes

// CS stores the classes information
//
// should use mutex here, since it is a shared data
var cs ClassStore

// this line will ensure the ClassStore to implement the methods (compile time check)
var _ Classes = &ClassStore{}

// AddClass adds new class to classes store
func (cs *ClassStore) AddClass(class Class) int {
	cs.nextID++

	class.ID = cs.nextID
	cs.classes = append(cs.classes, class)

	return cs.nextID
}
