package entity

import (
	"fmt"
	"testing"
)

// customID demonstrates a user-defined ID type that implements String().
type customID struct {
	id float64
}

func (c customID) String() string {
	return fmt.Sprintf("%v", c.id)
}

func TestNew_WithStringID_and_value_using_helpers(t *testing.T) {
	// Use NewWithStringID so only V is specified.
	e := NewWithStringID[string](WithStringID("s100"), WithValueForStringID("v"))
	if got := e.StrID(); got != "s100" {
		t.Fatalf("expected StrID to be %q, got %q", "s100", got)
	}
	if got := e.Value(); got != "v" {
		t.Fatalf("expected Value to be %q, got %q", "v", got)
	}
}

func TestNew_WithIntID_using_helpers(t *testing.T) {
	e := NewWithIntID[string](WithIntID(100), WithValueForIntID("v"))
	if got := e.StrID(); got != "100" {
		t.Fatalf("expected StrID to be %q, got %q", "100", got)
	}
}

func TestNew_WithCustomIDType(t *testing.T) {
	// Use a custom ID type that implements String()
	e := New[customID, string](WithID[customID, string](customID{id: 100.2}), WithValue[customID, string]("x"))
	if got := e.StrID(); got != "100.2" {
		t.Fatalf("expected StrID to be %q, got %q", "100.2", got)
	}
	if got := e.Value(); got != "x" {
		t.Fatalf("expected Value to be %q, got %q", "x", got)
	}
}

func TestNew_Empty_for_int_id_and_value(t *testing.T) {
	// Using K=int, V=int; zero-values used
	e := New[int, int]()
	var zero int
	if e.Value() != zero {
		t.Fatalf("expected zero value %v, got %v", zero, e.Value())
	}
	// ID is zero value for K (here int), StrID falls back to fmt and should be "0"
	if got := e.StrID(); got != "0" {
		t.Fatalf("expected StrID to be %q, got %q", "0", got)
	}
}

func TestNew_Empty_for_string_id_and_value(t *testing.T) {
	// Using K=string, V=int; zero-values used
	e := New[string, int]()
	var zero int
	if e.Value() != zero {
		t.Fatalf("expected zero value %v, got %v", zero, e.Value())
	}
	// ID is zero value for string (""), StrID should be empty string
	if got := e.StrID(); got != "" {
		t.Fatalf("expected StrID to be empty string, got %q", got)
	}
}

// --- One-shot constructor tests ---

func TestNewStringIDEntity_OneShot(t *testing.T) {
	e := NewStringIDEntity[string]("one-shot", "payload")
	if got := e.StrID(); got != "one-shot" {
		t.Fatalf("expected StrID to be %q, got %q", "one-shot", got)
	}
	if got := e.Value(); got != "payload" {
		t.Fatalf("expected Value to be %q, got %q", "payload", got)
	}
}

func TestNewIntIDEntity_OneShot(t *testing.T) {
	e := NewIntIDEntity[string](7, "seven")
	if got := e.StrID(); got != "7" {
		t.Fatalf("expected StrID to be %q, got %q", "7", got)
	}
	if got := e.Value(); got != "seven" {
		t.Fatalf("expected Value to be %q, got %q", "seven", got)
	}
}

func TestNewEntityWithID_GenericOneShot(t *testing.T) {
	// customID is a struct type implementing String(); used as K
	k := customID{id: 3.14}
	e := NewEntityWithID[customID, string](k, "pi")
	if got := e.StrID(); got != "3.14" {
		t.Fatalf("expected StrID to be %q, got %q", "3.14", got)
	}
	if got := e.Value(); got != "pi" {
		t.Fatalf("expected Value to be %q, got %q", "pi", got)
	}
}
