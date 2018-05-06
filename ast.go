package graphql_parser

import (
	"github.com/vektah/graphql-parser/lexer"
)

type Source struct {
	Body           string
	Name           string
	LocationOffset *Location
}

// Location contains a range of UTF-8 character offsets and token references that
// identify the region of the source from which the AST derived.
type Location struct {
	// The character offset at which this Node begins.
	Start int

	// The character offset at which this Node ends.
	End int

	// The Token at which this Node begins.
	StartToken Token

	// The Token at which this Node ends.
	EndToken Token

	// The Source document the AST represents.
	Source Source
}

// Token represents a range of characters represented by a lexical token
// within a Source.
type Token struct {
	// The kind of Token.
	Kind lexer.Type

	// The character offset at which this Node begins.
	Start int

	// The character offset at which this Node ends.
	End int

	// The 1-indexed line number on which this Token appears.
	Line int

	// The 1-indexed column number at which this Token begins.
	Column int

	// For non-punctuation tokens, represents the interpreted value of the token.
	Value string

	// Tokens exist as nodes in a double-linked-list amongst all tokens
	// including ignored tokens. <SOF> is always the first node and <EOF>
	// the last.
	Prev *Token
	Next *Token
}

// Name

type Name struct {
	Loc   Location
	Value string
}

// Document

type ExecutableDocument struct {
	Loc        Location
	Operations []OperationDefinition
	Fragments  []FragmentDefinition
}

type SchemaDocument struct {
	Loc                 Location
	SchemaDefinitions   []SchemaDefinition
	TypeDefinitions     []TypeDefinition
	DirectiveDefinition []DirectiveDefinition
	SchemaExtensions    []SchemaExtension
	TypeExtensions      []TypeExtension
}

type OperationDefinition struct {
	Loc                 Location
	Operation           string
	Name                Name
	VariableDefinitions []VariableDefinition
	Directives          []Directive
	SelectionSet        SelectionSet
}

type VariableDefinition struct {
	Loc          Location
	Variable     Variable
	Type         Type
	DefaultValue Value
}

type Variable struct {
	Loc  Location
	Name Name
}

type SelectionSet struct {
	Loc        Location
	Selections []Selection
}

type Selection interface {
	isSelection()
}

func (Field) isSelection()          {}
func (FragmentSpread) isSelection() {}
func (InlineFragment) isSelection() {}

type Field struct {
	Loc          Location
	Alias        Name
	Name         Name
	Arguments    []Argument
	Directives   []Directive
	SelectionSet SelectionSet
}

type Argument struct {
	Loc   Location
	Name  Name
	Value Value
}

// Fragments

type FragmentSpread struct {
	Loc        Location
	Name       Name
	Directives []Directive
}

type InlineFragment struct {
	Loc           Location
	TypeCondition NamedType
	Directives    []Directive
	SelectionSet  []SelectionSet
}

type FragmentDefinition struct {
	Loc  Location
	Name Name
	// Note: fragment variable definitions are experimental and may be changed
	// or removed in the future.
	VariableDefinition []VariableDefinition
	TypeCondition      NamedType
	Directives         []Directive
	SelectionSet       []SelectionSet
}

// Values

type Value interface {
	isValue()
}

func (Variable) isValue()     {}
func (IntValue) isValue()     {}
func (FloatValue) isValue()   {}
func (StringValue) isValue()  {}
func (BooleanValue) isValue() {}
func (NullValue) isValue()    {}
func (EnumValue) isValue()    {}
func (ListValue) isValue()    {}
func (ObjectValue) isValue()  {}

type IntValue struct {
	Loc   Location
	Value string
}

type FloatValue struct {
	Loc   Location
	Value string
}

type StringValue struct {
	Loc   Location
	Value string
	Block bool
}

type BooleanValue struct {
	Loc   Location
	Value bool
}

type NullValue struct {
	Loc Location
}

type EnumValue struct {
	Loc   Location
	Value string
}

type ListValue struct {
	Loc    Location
	Values []Value
}

type ObjectValue struct {
	Loc    Location
	Fields []ObjectField
}

type ObjectField struct {
	Loc   Location
	Name  Name
	Value Value
}

// Directives

type Directive struct {
	Loc       Location
	Name      Name
	Arguments []Argument
}

// Type Reference

type Type interface {
	isType()
}

func (NamedType) isType()   {}
func (ListType) isType()    {}
func (NonNullType) isType() {}

type NamedType struct {
	Loc  Location
	Name Name
}

type ListType struct {
	Loc  Location
	Type Type
}

type NonNullType struct {
	Loc  Location
	Type Type
}

// Type System Definition

type TypeDefinition interface {
	isTypeDefinition()
}

func (SchemaDefinition) isTypeDefinition()          {}
func (DirectiveDefinition) isTypeDefinition()       {}
func (ScalarTypeDefinition) isTypeDefinition()      {}
func (ObjectTypeDefinition) isTypeDefinition()      {}
func (InterfaceTypeDefinition) isTypeDefinition()   {}
func (UnionTypeDefinition) isTypeDefinition()       {}
func (EnumTypeDefinition) isTypeDefinition()        {}
func (InputObjectTypeDefinition) isTypeDefinition() {}

type SchemaDefinition struct {
	Loc            Location
	Directives     []Directive
	OperationTypes []OperationTypeDefinition
}

type OperationTypeDefinition struct {
	Loc       Location
	Operation string
	Type      NamedType
}

// Type Definition

type ScalarTypeDefinition struct {
	Loc         Location
	Description StringValue
	Name        Name
	Directives  []Directive
}

type ObjectTypeDefinition struct {
	Loc         Location
	Description StringValue
	Name        Name
	Interfaces  []NamedType
	Directives  []Directive
	Fields      []FieldDefinition
}

type FieldDefinition struct {
	Loc         Location
	Description StringValue
	Name        Name
	Arguments   []InputValueDefinition
	Type        Type
	Directives  []Directive
}

type InputValueDefinition struct {
	Loc          Location
	Description  StringValue
	Name         Name
	Type         Type
	DefaultValue Value
	Directives   []Directive
}

type InterfaceTypeDefinition struct {
	Loc         Location
	Description StringValue
	Name        Name
	Directives  []Directive
	Fields      []FieldDefinition
}

type UnionTypeDefinition struct {
	Loc         Location
	Description StringValue
	Name        Name
	Directives  []Directive
	Types       []NamedType
}

type EnumTypeDefinition struct {
	Loc         Location
	Description StringValue
	Name        Name
	Directives  []Directive
	Values      []EnumValueDefinition
}

type EnumValueDefinition struct {
	Loc         Location
	Description StringValue
	Name        Name
	Directives  []Directive
}

type InputObjectTypeDefinition struct {
	Loc         Location
	Description StringValue
	Name        Name
	Directives  []Directive
	Fields      []InputValueDefinition
}

// Directive Definitions

type DirectiveDefinition struct {
	Loc         Location
	Description StringValue
	Name        Name
	Arguments   InputValueDefinition
	Locations   Name
}

// Type System Extensions

type SchemaExtension struct {
	Loc            Location
	Directives     []Directive
	OperationTypes []OperationTypeDefinition
}

type TypeExtension interface {
	isTypeExtension()
}

func (ScalarTypeExtension) isTypeExtension()      {}
func (ObjectTypeExtension) isTypeExtension()      {}
func (InterfaceTypeExtension) isTypeExtension()   {}
func (UnionTypeExtension) isTypeExtension()       {}
func (EnumTypeExtension) isTypeExtension()        {}
func (InputObjectTypeExtension) isTypeExtension() {}

// Type Extensions

type ScalarTypeExtension struct {
	Loc        Location
	Name       Name
	Directives []Directive
}

type ObjectTypeExtension struct {
	Loc        Location
	Name       Name
	Interfaces NamedType
	Directives []Directive
	Fields     FieldDefinition
}

type InterfaceTypeExtension struct {
	Loc        Location
	Name       Name
	Directives []Directive
	Fields     []FieldDefinition
}

type UnionTypeExtension struct {
	Loc        Location
	Name       Name
	Directives []Directive
	Types      NamedType
}

type EnumTypeExtension struct {
	Loc        Location
	Name       Name
	Directives []Directive
	Values     EnumValueDefinition
}

type InputObjectTypeExtension struct {
	Loc        Location
	Name       Name
	Directives []Directive
	Fields     InputValueDefinition
}
