package schema

import (
	"geeorm/dialect"
	"go/ast"
	"reflect"
)

/**
Dialect 实现了一些特定的 SQL 语句的转换，接下来我们将要实现 ORM 框架中最为核心的转换——对象(object)和表(table)的转换。给定一个任意的对象，转换为关系型数据库中的表结构。

在数据库中创建一张表需要哪些要素呢？

表名(table name) —— 结构体名(struct name)
字段名和字段类型 —— 成员变量和类型。
额外的约束条件(例如非空、主键等) —— 成员变量的Tag（Go 语言通过 Tag 实现，Java、Python 等语言通过注解实现）
*/

type Field struct {
	Name string //字段名
	Type string //类型
	Tag  string //约束条件
}
type Schema struct {
	Model      interface{}       //被映射的原始对象
	Name       string            //表名
	Fields     []*Field          //字段
	FieldNames []string          //所有的字段名
	fieldMap   map[string]*Field //记录字段名和Field映射关系，方便之后直接使用，不用遍历Fields切片
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

func (schem *Schema) RecordValues(dest interface{}) []interface{} {

	//Indirect返回v指向的值，如果v是个nil指针，Indirect返回零值，如果v不是指针，Indirect返回v本身。这里也理解了为啥要起Indirect这个名字，间接获取到了v指向的值。
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range schem.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}

// 将任意的对象解析为 Schema 实例。
func Parse(dest interface{}, d dialect.Dialect) *Schema {

	/**
	todo
		TypeOf() 和 ValueOf() 是 reflect 包最为基本也是最重要的 2 个方法，分别用来返回入参的类型和值。因为设计的入参是一个对象的指针，因此需要 reflect.Indirect() 获取指针指向的实例。
		modelType.Name() 获取到结构体的名称作为表名。
		NumField() 获取实例的字段的个数，然后通过下标获取到特定字段 p := modelType.Field(i)。
		p.Name 即字段名，p.Type 即字段类型，通过 (Dialect).DataTypeOf() 转换为数据库的字段类型，p.Tag 即额外的约束条件。
	*/

	/**
	比如传入的dest是结构体，reflect.ValueOf(dest)，拿到是dest的反射Value；
	reflect.Indirect(reflect.ValueOf(dest)).Type() ，拿到是具体的dest的哪个包的的那个结构体类型，，比如是  model.User
	*/
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model: dest,
		//modelType.Name() 拿到传入的dest的类型名称，比如传入的是结构体的话，则是拿到是比如：User，这个字符串的名称
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}

	//
	for i := 0; i < modelType.NumField(); i++ {

		//todo modelType.Field(i) 拿到的是某个结构体中的某个字段信息StructField
		//type StructField struct {
		// Name is the field name.
		//Name string

		// PkgPath is the package path that qualifies a lower case (unexported)
		// field name. It is empty for upper case (exported) field names.
		// See https://golang.org/ref/spec#Uniqueness_of_identifiers
		//PkgPath string

		//Type      Type      // field type
		//Tag       StructTag // field tag string
		//Offset    uintptr   // offset within struct, in bytes
		//Index     []int     // index sequence for Type.FieldByIndex
		//Anonymous bool      // is an embedded field
		//}
		p := modelType.Field(i)

		//Anonymous 是否是匿名的嵌入式字段  && 字段是否可以被导出
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}
