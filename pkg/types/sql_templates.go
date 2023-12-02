package types

import (
	"fmt"
	"reflect"
)

type SQLType string

const (
	SQLTypeInt    SQLType = "int"
	SQLTypeSerial SQLType = "serial"

	SQLTypeBool SQLType = "bool"

	SQLTypeTimestamp   SQLType = "timestamp"
	SQLTypeTimestamptz SQLType = "timestamptz"

	SQLTypeVarchar SQLType = "varchar"
	SQLTypeText    SQLType = "text"
)

type SQLColType interface {
	int | string | bool | float64
}

type SqlFKConfig struct {
	field    string
	refTable string
	refField string
}

type SqlColConfig[T SQLColType] struct {
	useNotNull    bool
	useUnique     bool
	usePrimaryKey bool
	useDefault    bool
	useTypeParam  bool
	useCheck      bool
	useForeignKey bool

	valueDefault   T
	valueTypeParam T

	valueCheck string

	valueForeignKey []*SqlFKConfig
}

type SqlCol[T SQLColType] struct {
	name     string
	basetype T
	sqltype  SQLType

	config *SqlColConfig[T]
}

type SqlTableConfig struct {
	hasPKCol   bool
	pkColIndex int
	pkColName  string
}

type SqlTable[T SQLColType] struct {
	name    string
	config  *SqlTableConfig
	columns []*SqlCol[T]
}

func (s *SqlTable[T]) SyncIdentSeq() string {
	return fmt.Sprintf(`
		ALTER TABLE %s
			ALTER %s SET NOT NULL,
			ALTER %s ADD GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1);
		
		SELECT setval(pg_get_serial_sequence('%s', '%s'),
			(SELECT MAX(%s) FROM %s));
	`, s.name, s.config.pkColName, s.config.pkColName, s.name, s.config.pkColName, s.config.pkColName, s.name)
}

/*****************************/
/* TEMPLATE SETUP & DEFAULTS */
/*****************************/

/* CONFIG DEFAULTS */

func defaultSqlTableConfig(t ...bool) *SqlTableConfig {
	config := &SqlTableConfig{
		hasPKCol:   false,
		pkColIndex: -1,
		pkColName:  "",
	}

	if len(t) > 0 && t[0] {
		config.hasPKCol = true
		config.pkColIndex = 0
		config.pkColName = "id"
	}

	return config
}

func defaultSqlFKConfig(kf, rt, rf string) *SqlFKConfig {
	return &SqlFKConfig{
		field:    kf,
		refTable: rt,
		refField: rf,
	}
}

func defaultSqlColConfig[T SQLColType](defaultValue ...T) *SqlColConfig[T] {
	var baseTypeDefault T

	baseTypeReflect := reflect.TypeOf(baseTypeDefault)
	baseTypeReflectVal := reflect.ValueOf(&baseTypeDefault)

	switch baseTypeReflect.Kind() {
	case reflect.Int:
		baseTypeReflectVal.Elem().SetInt(1)
	case reflect.Bool:
		baseTypeReflectVal.Elem().SetBool(false)
	case reflect.String:
		baseTypeReflectVal.Elem().SetString("")
	case reflect.Float64:
		baseTypeReflectVal.Elem().SetFloat(0.0)
	}

	config := &SqlColConfig[T]{
		useNotNull:      false,
		useUnique:       false,
		usePrimaryKey:   false,
		useDefault:      false,
		useTypeParam:    false,
		useCheck:        false,
		useForeignKey:   false,
		valueDefault:    baseTypeDefault,
		valueTypeParam:  baseTypeDefault,
		valueCheck:      "",
		valueForeignKey: nil,
	}

	if len(defaultValue) > 0 && reflect.TypeOf(defaultValue[0]) == baseTypeReflect {
		config.valueDefault = defaultValue[0]
	}

	return config
}

/* INTERNAL FNS */

func (s *SqlTable[T]) update() {
	if s.columns != nil && len(s.columns) > 0 {
		for i, col := range s.columns {
			if col.config.usePrimaryKey {
				s.config.hasPKCol = true
				s.config.pkColIndex = i
				s.config.pkColName = col.name
			}
		}
	}
}

/* CONFIG FNS */
type sqlColConfigFunc[T SQLColType] func(*SqlColConfig[T]) *SqlColConfig[T]

func UseNotNull[T SQLColType]() sqlColConfigFunc[T] {
	return func(c *SqlColConfig[T]) *SqlColConfig[T] {
		c.useNotNull = true
		return c
	}
}

func UseUnique[T SQLColType]() sqlColConfigFunc[T] {
	return func(c *SqlColConfig[T]) *SqlColConfig[T] {
		c.useUnique = true
		return c
	}
}

func UsePrimaryKey[T SQLColType]() sqlColConfigFunc[T] {
	return func(c *SqlColConfig[T]) *SqlColConfig[T] {
		c.usePrimaryKey = true
		return c
	}
}

func UseDefault[T SQLColType](v T) sqlColConfigFunc[T] {
	return func(c *SqlColConfig[T]) *SqlColConfig[T] {
		c.useDefault = true
		c.valueDefault = v
		return c
	}
}

func UseTypeParam[T SQLColType](v T) sqlColConfigFunc[T] {
	return func(c *SqlColConfig[T]) *SqlColConfig[T] {
		c.useTypeParam = true
		c.valueTypeParam = v
		return c
	}
}

func UseCheck[T SQLColType](s string) sqlColConfigFunc[T] {
	return func(c *SqlColConfig[T]) *SqlColConfig[T] {
		c.useCheck = true
		c.valueCheck = s
		return c
	}
}

func UseForeignKey[T SQLColType](fk ...*SqlFKConfig) sqlColConfigFunc[T] {
	return func(c *SqlColConfig[T]) *SqlColConfig[T] {
		c.useForeignKey = true
		c.valueForeignKey = fk
		return c
	}
}

/******/

func (t *SqlTable[T]) NewCol() {
	t.update()
}

/* ENTRY POINTS */

func NewSqlTable[T SQLColType](name string) *SqlTable[T] {
	config := defaultSqlTableConfig()
	table := &SqlTable[T]{
		name:    name,
		config:  config,
		columns: make([]*SqlCol[T], 0),
	}
	return table
}

func NewCol[T SQLColType](name string, basetype T, sqltype SQLType, cfg ...sqlColConfigFunc[T]) *SqlCol[T] {
	config := defaultSqlColConfig[T]()
	for _, fn := range cfg {
		config = fn(config)
	}
	col := &SqlCol[T]{
		name:     name,
		basetype: basetype,
		sqltype:  sqltype,
		config:   config,
	}

	return col
}

func NewFK(name, refTable, refField string) *SqlFKConfig {
	return defaultSqlFKConfig(name, refTable, refField)
}
