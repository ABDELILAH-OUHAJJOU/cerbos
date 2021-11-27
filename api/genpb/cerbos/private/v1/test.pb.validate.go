// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: cerbos/private/v1/test.proto

package privatev1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on EngineTestCase with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EngineTestCase) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Description

	for idx, item := range m.GetInputs() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EngineTestCaseValidationError{
					field:  fmt.Sprintf("Inputs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetWantOutputs() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EngineTestCaseValidationError{
					field:  fmt.Sprintf("WantOutputs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for WantError

	return nil
}

// EngineTestCaseValidationError is the validation error returned by
// EngineTestCase.Validate if the designated constraints aren't met.
type EngineTestCaseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EngineTestCaseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EngineTestCaseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EngineTestCaseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EngineTestCaseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EngineTestCaseValidationError) ErrorName() string { return "EngineTestCaseValidationError" }

// Error satisfies the builtin error interface
func (e EngineTestCaseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEngineTestCase.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EngineTestCaseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EngineTestCaseValidationError{}

// Validate checks the field values on ServerTestCase with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ServerTestCase) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for WantError

	if v, ok := interface{}(m.GetWantStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerTestCaseValidationError{
				field:  "WantStatus",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.CallKind.(type) {

	case *ServerTestCase_CheckResourceSet:

		if v, ok := interface{}(m.GetCheckResourceSet()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServerTestCaseValidationError{
					field:  "CheckResourceSet",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ServerTestCase_CheckResourceBatch:

		if v, ok := interface{}(m.GetCheckResourceBatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServerTestCaseValidationError{
					field:  "CheckResourceBatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ServerTestCase_PlaygroundValidate:

		if v, ok := interface{}(m.GetPlaygroundValidate()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServerTestCaseValidationError{
					field:  "PlaygroundValidate",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ServerTestCase_PlaygroundEvaluate:

		if v, ok := interface{}(m.GetPlaygroundEvaluate()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServerTestCaseValidationError{
					field:  "PlaygroundEvaluate",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ServerTestCase_AdminAddOrUpdatePolicy:

		if v, ok := interface{}(m.GetAdminAddOrUpdatePolicy()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServerTestCaseValidationError{
					field:  "AdminAddOrUpdatePolicy",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ServerTestCase_PlaygroundProxy:

		if v, ok := interface{}(m.GetPlaygroundProxy()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServerTestCaseValidationError{
					field:  "PlaygroundProxy",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ServerTestCaseValidationError is the validation error returned by
// ServerTestCase.Validate if the designated constraints aren't met.
type ServerTestCaseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerTestCaseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerTestCaseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerTestCaseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerTestCaseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerTestCaseValidationError) ErrorName() string { return "ServerTestCaseValidationError" }

// Error satisfies the builtin error interface
func (e ServerTestCaseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerTestCase.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerTestCaseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerTestCaseValidationError{}

// Validate checks the field values on IndexBuilderTestCase with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *IndexBuilderTestCase) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Files

	// no validation rules for WantErrJson

	// no validation rules for WantErr

	return nil
}

// IndexBuilderTestCaseValidationError is the validation error returned by
// IndexBuilderTestCase.Validate if the designated constraints aren't met.
type IndexBuilderTestCaseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndexBuilderTestCaseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndexBuilderTestCaseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndexBuilderTestCaseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndexBuilderTestCaseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndexBuilderTestCaseValidationError) ErrorName() string {
	return "IndexBuilderTestCaseValidationError"
}

// Error satisfies the builtin error interface
func (e IndexBuilderTestCaseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndexBuilderTestCase.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndexBuilderTestCaseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndexBuilderTestCaseValidationError{}

// Validate checks the field values on CompileTestCase with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CompileTestCase) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for MainDef

	for key, val := range m.GetInputDefs() {
		_ = val

		// no validation rules for InputDefs[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CompileTestCaseValidationError{
					field:  fmt.Sprintf("InputDefs[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CompileTestCaseValidationError is the validation error returned by
// CompileTestCase.Validate if the designated constraints aren't met.
type CompileTestCaseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompileTestCaseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompileTestCaseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompileTestCaseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompileTestCaseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompileTestCaseValidationError) ErrorName() string { return "CompileTestCaseValidationError" }

// Error satisfies the builtin error interface
func (e CompileTestCaseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompileTestCase.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompileTestCaseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompileTestCaseValidationError{}

// Validate checks the field values on CodeGenTestCase with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CodeGenTestCase) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetInputPolicy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CodeGenTestCaseValidationError{
				field:  "InputPolicy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for WantError

	// no validation rules for WantRego

	// no validation rules for WantNumConditions

	return nil
}

// CodeGenTestCaseValidationError is the validation error returned by
// CodeGenTestCase.Validate if the designated constraints aren't met.
type CodeGenTestCaseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CodeGenTestCaseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CodeGenTestCaseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CodeGenTestCaseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CodeGenTestCaseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CodeGenTestCaseValidationError) ErrorName() string { return "CodeGenTestCaseValidationError" }

// Error satisfies the builtin error interface
func (e CodeGenTestCaseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCodeGenTestCase.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CodeGenTestCaseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CodeGenTestCaseValidationError{}

// Validate checks the field values on CelTestCase with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CelTestCase) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCondition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CelTestCaseValidationError{
				field:  "Condition",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetInput()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CelTestCaseValidationError{
				field:  "Input",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Want

	// no validation rules for WantError

	return nil
}

// CelTestCaseValidationError is the validation error returned by
// CelTestCase.Validate if the designated constraints aren't met.
type CelTestCaseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CelTestCaseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CelTestCaseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CelTestCaseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CelTestCaseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CelTestCaseValidationError) ErrorName() string { return "CelTestCaseValidationError" }

// Error satisfies the builtin error interface
func (e CelTestCaseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCelTestCase.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CelTestCaseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CelTestCaseValidationError{}

// Validate checks the field values on SchemaTestCase with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SchemaTestCase) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Description

	if v, ok := interface{}(m.GetSchema()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SchemaTestCaseValidationError{
				field:  "Schema",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetInput()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SchemaTestCaseValidationError{
				field:  "Input",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for InvalidSchema

	// no validation rules for WantError

	for idx, item := range m.GetWantValidationErrors() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SchemaTestCaseValidationError{
					field:  fmt.Sprintf("WantValidationErrors[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// SchemaTestCaseValidationError is the validation error returned by
// SchemaTestCase.Validate if the designated constraints aren't met.
type SchemaTestCaseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SchemaTestCaseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SchemaTestCaseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SchemaTestCaseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SchemaTestCaseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SchemaTestCaseValidationError) ErrorName() string { return "SchemaTestCaseValidationError" }

// Error satisfies the builtin error interface
func (e SchemaTestCaseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSchemaTestCase.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SchemaTestCaseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SchemaTestCaseValidationError{}

// Validate checks the field values on ValidationErrContainer with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ValidationErrContainer) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetErrors() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ValidationErrContainerValidationError{
					field:  fmt.Sprintf("Errors[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ValidationErrContainerValidationError is the validation error returned by
// ValidationErrContainer.Validate if the designated constraints aren't met.
type ValidationErrContainerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValidationErrContainerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValidationErrContainerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValidationErrContainerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValidationErrContainerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValidationErrContainerValidationError) ErrorName() string {
	return "ValidationErrContainerValidationError"
}

// Error satisfies the builtin error interface
func (e ValidationErrContainerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValidationErrContainer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValidationErrContainerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValidationErrContainerValidationError{}

// Validate checks the field values on AttrWrapper with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AttrWrapper) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetAttr() {
		_ = val

		// no validation rules for Attr[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AttrWrapperValidationError{
					field:  fmt.Sprintf("Attr[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// AttrWrapperValidationError is the validation error returned by
// AttrWrapper.Validate if the designated constraints aren't met.
type AttrWrapperValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttrWrapperValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttrWrapperValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttrWrapperValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttrWrapperValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttrWrapperValidationError) ErrorName() string { return "AttrWrapperValidationError" }

// Error satisfies the builtin error interface
func (e AttrWrapperValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttrWrapper.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttrWrapperValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttrWrapperValidationError{}

// Validate checks the field values on ServerTestCase_CheckResourceSetCall with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ServerTestCase_CheckResourceSetCall) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetInput()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerTestCase_CheckResourceSetCallValidationError{
				field:  "Input",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetWantResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerTestCase_CheckResourceSetCallValidationError{
				field:  "WantResponse",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ServerTestCase_CheckResourceSetCallValidationError is the validation error
// returned by ServerTestCase_CheckResourceSetCall.Validate if the designated
// constraints aren't met.
type ServerTestCase_CheckResourceSetCallValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerTestCase_CheckResourceSetCallValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerTestCase_CheckResourceSetCallValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerTestCase_CheckResourceSetCallValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerTestCase_CheckResourceSetCallValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerTestCase_CheckResourceSetCallValidationError) ErrorName() string {
	return "ServerTestCase_CheckResourceSetCallValidationError"
}

// Error satisfies the builtin error interface
func (e ServerTestCase_CheckResourceSetCallValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerTestCase_CheckResourceSetCall.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerTestCase_CheckResourceSetCallValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerTestCase_CheckResourceSetCallValidationError{}

// Validate checks the field values on ServerTestCase_CheckResourceBatchCall
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *ServerTestCase_CheckResourceBatchCall) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetInput()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerTestCase_CheckResourceBatchCallValidationError{
				field:  "Input",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetWantResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerTestCase_CheckResourceBatchCallValidationError{
				field:  "WantResponse",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ServerTestCase_CheckResourceBatchCallValidationError is the validation error
// returned by ServerTestCase_CheckResourceBatchCall.Validate if the
// designated constraints aren't met.
type ServerTestCase_CheckResourceBatchCallValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerTestCase_CheckResourceBatchCallValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerTestCase_CheckResourceBatchCallValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerTestCase_CheckResourceBatchCallValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerTestCase_CheckResourceBatchCallValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerTestCase_CheckResourceBatchCallValidationError) ErrorName() string {
	return "ServerTestCase_CheckResourceBatchCallValidationError"
}

// Error satisfies the builtin error interface
func (e ServerTestCase_CheckResourceBatchCallValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerTestCase_CheckResourceBatchCall.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerTestCase_CheckResourceBatchCallValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerTestCase_CheckResourceBatchCallValidationError{}

// Validate checks the field values on ServerTestCase_PlaygroundValidateCall
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *ServerTestCase_PlaygroundValidateCall) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetInput()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerTestCase_PlaygroundValidateCallValidationError{
				field:  "Input",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetWantResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerTestCase_PlaygroundValidateCallValidationError{
				field:  "WantResponse",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ServerTestCase_PlaygroundValidateCallValidationError is the validation error
// returned by ServerTestCase_PlaygroundValidateCall.Validate if the
// designated constraints aren't met.
type ServerTestCase_PlaygroundValidateCallValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerTestCase_PlaygroundValidateCallValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerTestCase_PlaygroundValidateCallValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerTestCase_PlaygroundValidateCallValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerTestCase_PlaygroundValidateCallValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerTestCase_PlaygroundValidateCallValidationError) ErrorName() string {
	return "ServerTestCase_PlaygroundValidateCallValidationError"
}

// Error satisfies the builtin error interface
func (e ServerTestCase_PlaygroundValidateCallValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerTestCase_PlaygroundValidateCall.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerTestCase_PlaygroundValidateCallValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerTestCase_PlaygroundValidateCallValidationError{}

// Validate checks the field values on ServerTestCase_PlaygroundEvaluateCall
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *ServerTestCase_PlaygroundEvaluateCall) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetInput()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerTestCase_PlaygroundEvaluateCallValidationError{
				field:  "Input",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetWantResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerTestCase_PlaygroundEvaluateCallValidationError{
				field:  "WantResponse",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ServerTestCase_PlaygroundEvaluateCallValidationError is the validation error
// returned by ServerTestCase_PlaygroundEvaluateCall.Validate if the
// designated constraints aren't met.
type ServerTestCase_PlaygroundEvaluateCallValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerTestCase_PlaygroundEvaluateCallValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerTestCase_PlaygroundEvaluateCallValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerTestCase_PlaygroundEvaluateCallValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerTestCase_PlaygroundEvaluateCallValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerTestCase_PlaygroundEvaluateCallValidationError) ErrorName() string {
	return "ServerTestCase_PlaygroundEvaluateCallValidationError"
}

// Error satisfies the builtin error interface
func (e ServerTestCase_PlaygroundEvaluateCallValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerTestCase_PlaygroundEvaluateCall.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerTestCase_PlaygroundEvaluateCallValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerTestCase_PlaygroundEvaluateCallValidationError{}

// Validate checks the field values on ServerTestCase_PlaygroundProxyCall with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ServerTestCase_PlaygroundProxyCall) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetInput()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerTestCase_PlaygroundProxyCallValidationError{
				field:  "Input",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetWantResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerTestCase_PlaygroundProxyCallValidationError{
				field:  "WantResponse",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ServerTestCase_PlaygroundProxyCallValidationError is the validation error
// returned by ServerTestCase_PlaygroundProxyCall.Validate if the designated
// constraints aren't met.
type ServerTestCase_PlaygroundProxyCallValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerTestCase_PlaygroundProxyCallValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerTestCase_PlaygroundProxyCallValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerTestCase_PlaygroundProxyCallValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerTestCase_PlaygroundProxyCallValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerTestCase_PlaygroundProxyCallValidationError) ErrorName() string {
	return "ServerTestCase_PlaygroundProxyCallValidationError"
}

// Error satisfies the builtin error interface
func (e ServerTestCase_PlaygroundProxyCallValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerTestCase_PlaygroundProxyCall.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerTestCase_PlaygroundProxyCallValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerTestCase_PlaygroundProxyCallValidationError{}

// Validate checks the field values on
// ServerTestCase_AdminAddOrUpdatePolicyCall with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ServerTestCase_AdminAddOrUpdatePolicyCall) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetInput()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerTestCase_AdminAddOrUpdatePolicyCallValidationError{
				field:  "Input",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetWantResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerTestCase_AdminAddOrUpdatePolicyCallValidationError{
				field:  "WantResponse",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ServerTestCase_AdminAddOrUpdatePolicyCallValidationError is the validation
// error returned by ServerTestCase_AdminAddOrUpdatePolicyCall.Validate if the
// designated constraints aren't met.
type ServerTestCase_AdminAddOrUpdatePolicyCallValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerTestCase_AdminAddOrUpdatePolicyCallValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerTestCase_AdminAddOrUpdatePolicyCallValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerTestCase_AdminAddOrUpdatePolicyCallValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerTestCase_AdminAddOrUpdatePolicyCallValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerTestCase_AdminAddOrUpdatePolicyCallValidationError) ErrorName() string {
	return "ServerTestCase_AdminAddOrUpdatePolicyCallValidationError"
}

// Error satisfies the builtin error interface
func (e ServerTestCase_AdminAddOrUpdatePolicyCallValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerTestCase_AdminAddOrUpdatePolicyCall.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerTestCase_AdminAddOrUpdatePolicyCallValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerTestCase_AdminAddOrUpdatePolicyCallValidationError{}

// Validate checks the field values on ServerTestCase_Status with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ServerTestCase_Status) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for HttpStatusCode

	// no validation rules for GrpcStatusCode

	return nil
}

// ServerTestCase_StatusValidationError is the validation error returned by
// ServerTestCase_Status.Validate if the designated constraints aren't met.
type ServerTestCase_StatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerTestCase_StatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerTestCase_StatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerTestCase_StatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerTestCase_StatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerTestCase_StatusValidationError) ErrorName() string {
	return "ServerTestCase_StatusValidationError"
}

// Error satisfies the builtin error interface
func (e ServerTestCase_StatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerTestCase_Status.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerTestCase_StatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerTestCase_StatusValidationError{}
