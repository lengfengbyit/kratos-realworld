// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: article/v1/article.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on SlugRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SlugRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SlugRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SlugRequestMultiError, or
// nil if none found.
func (m *SlugRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SlugRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetSlug()) < 1 {
		err := SlugRequestValidationError{
			field:  "Slug",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SlugRequestMultiError(errors)
	}

	return nil
}

// SlugRequestMultiError is an error wrapping multiple validation errors
// returned by SlugRequest.ValidateAll() if the designated constraints aren't met.
type SlugRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SlugRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SlugRequestMultiError) AllErrors() []error { return m }

// SlugRequestValidationError is the validation error returned by
// SlugRequest.Validate if the designated constraints aren't met.
type SlugRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SlugRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SlugRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SlugRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SlugRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SlugRequestValidationError) ErrorName() string { return "SlugRequestValidationError" }

// Error satisfies the builtin error interface
func (e SlugRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSlugRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SlugRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SlugRequestValidationError{}

// Validate checks the field values on CreateArticleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateArticleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateArticleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateArticleRequestMultiError, or nil if none found.
func (m *CreateArticleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateArticleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetArticle() == nil {
		err := CreateArticleRequestValidationError{
			field:  "Article",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetArticle()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateArticleRequestValidationError{
					field:  "Article",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateArticleRequestValidationError{
					field:  "Article",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetArticle()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateArticleRequestValidationError{
				field:  "Article",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateArticleRequestMultiError(errors)
	}

	return nil
}

// CreateArticleRequestMultiError is an error wrapping multiple validation
// errors returned by CreateArticleRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateArticleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateArticleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateArticleRequestMultiError) AllErrors() []error { return m }

// CreateArticleRequestValidationError is the validation error returned by
// CreateArticleRequest.Validate if the designated constraints aren't met.
type CreateArticleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateArticleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateArticleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateArticleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateArticleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateArticleRequestValidationError) ErrorName() string {
	return "CreateArticleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateArticleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateArticleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateArticleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateArticleRequestValidationError{}

// Validate checks the field values on SaveArticleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SaveArticleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SaveArticleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SaveArticleRequestMultiError, or nil if none found.
func (m *SaveArticleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SaveArticleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetArticle() == nil {
		err := SaveArticleRequestValidationError{
			field:  "Article",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetArticle()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SaveArticleRequestValidationError{
					field:  "Article",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SaveArticleRequestValidationError{
					field:  "Article",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetArticle()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SaveArticleRequestValidationError{
				field:  "Article",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetSlug()) < 1 {
		err := SaveArticleRequestValidationError{
			field:  "Slug",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SaveArticleRequestMultiError(errors)
	}

	return nil
}

// SaveArticleRequestMultiError is an error wrapping multiple validation errors
// returned by SaveArticleRequest.ValidateAll() if the designated constraints
// aren't met.
type SaveArticleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SaveArticleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SaveArticleRequestMultiError) AllErrors() []error { return m }

// SaveArticleRequestValidationError is the validation error returned by
// SaveArticleRequest.Validate if the designated constraints aren't met.
type SaveArticleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveArticleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveArticleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveArticleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveArticleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveArticleRequestValidationError) ErrorName() string {
	return "SaveArticleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SaveArticleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveArticleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveArticleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveArticleRequestValidationError{}

// Validate checks the field values on ListArticleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListArticleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListArticleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListArticleRequestMultiError, or nil if none found.
func (m *ListArticleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListArticleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Tag

	// no validation rules for Author

	// no validation rules for Favorited

	if m.GetLimit() != 0 {

		if val := m.GetLimit(); val <= 0 || val >= 100 {
			err := ListArticleRequestValidationError{
				field:  "Limit",
				reason: "value must be inside range (0, 100)",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetOffset() != 0 {

		if m.GetOffset() < 0 {
			err := ListArticleRequestValidationError{
				field:  "Offset",
				reason: "value must be greater than or equal to 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return ListArticleRequestMultiError(errors)
	}

	return nil
}

// ListArticleRequestMultiError is an error wrapping multiple validation errors
// returned by ListArticleRequest.ValidateAll() if the designated constraints
// aren't met.
type ListArticleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListArticleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListArticleRequestMultiError) AllErrors() []error { return m }

// ListArticleRequestValidationError is the validation error returned by
// ListArticleRequest.Validate if the designated constraints aren't met.
type ListArticleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListArticleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListArticleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListArticleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListArticleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListArticleRequestValidationError) ErrorName() string {
	return "ListArticleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListArticleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListArticleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListArticleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListArticleRequestValidationError{}

// Validate checks the field values on ListArticleReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListArticleReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListArticleReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListArticleReplyMultiError, or nil if none found.
func (m *ListArticleReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListArticleReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetArticles() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListArticleReplyValidationError{
						field:  fmt.Sprintf("Articles[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListArticleReplyValidationError{
						field:  fmt.Sprintf("Articles[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListArticleReplyValidationError{
					field:  fmt.Sprintf("Articles[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for ArticlesCount

	if len(errors) > 0 {
		return ListArticleReplyMultiError(errors)
	}

	return nil
}

// ListArticleReplyMultiError is an error wrapping multiple validation errors
// returned by ListArticleReply.ValidateAll() if the designated constraints
// aren't met.
type ListArticleReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListArticleReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListArticleReplyMultiError) AllErrors() []error { return m }

// ListArticleReplyValidationError is the validation error returned by
// ListArticleReply.Validate if the designated constraints aren't met.
type ListArticleReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListArticleReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListArticleReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListArticleReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListArticleReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListArticleReplyValidationError) ErrorName() string { return "ListArticleReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListArticleReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListArticleReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListArticleReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListArticleReplyValidationError{}

// Validate checks the field values on Author with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Author) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Author with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AuthorMultiError, or nil if none found.
func (m *Author) ValidateAll() error {
	return m.validate(true)
}

func (m *Author) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Bio

	// no validation rules for Image

	// no validation rules for Following

	if len(errors) > 0 {
		return AuthorMultiError(errors)
	}

	return nil
}

// AuthorMultiError is an error wrapping multiple validation errors returned by
// Author.ValidateAll() if the designated constraints aren't met.
type AuthorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthorMultiError) AllErrors() []error { return m }

// AuthorValidationError is the validation error returned by Author.Validate if
// the designated constraints aren't met.
type AuthorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthorValidationError) ErrorName() string { return "AuthorValidationError" }

// Error satisfies the builtin error interface
func (e AuthorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthor.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthorValidationError{}

// Validate checks the field values on ArticleReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ArticleReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ArticleReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ArticleReplyMultiError, or
// nil if none found.
func (m *ArticleReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ArticleReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Slug

	// no validation rules for Title

	// no validation rules for Description

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	// no validation rules for Favorited

	// no validation rules for FavoritesCount

	if all {
		switch v := interface{}(m.GetAuthor()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ArticleReplyValidationError{
					field:  "Author",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ArticleReplyValidationError{
					field:  "Author",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAuthor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ArticleReplyValidationError{
				field:  "Author",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ArticleReplyMultiError(errors)
	}

	return nil
}

// ArticleReplyMultiError is an error wrapping multiple validation errors
// returned by ArticleReply.ValidateAll() if the designated constraints aren't met.
type ArticleReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ArticleReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ArticleReplyMultiError) AllErrors() []error { return m }

// ArticleReplyValidationError is the validation error returned by
// ArticleReply.Validate if the designated constraints aren't met.
type ArticleReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArticleReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArticleReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArticleReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArticleReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArticleReplyValidationError) ErrorName() string { return "ArticleReplyValidationError" }

// Error satisfies the builtin error interface
func (e ArticleReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArticleReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArticleReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArticleReplyValidationError{}

// Validate checks the field values on EmptyReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EmptyReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EmptyReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EmptyReplyMultiError, or
// nil if none found.
func (m *EmptyReply) ValidateAll() error {
	return m.validate(true)
}

func (m *EmptyReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EmptyReplyMultiError(errors)
	}

	return nil
}

// EmptyReplyMultiError is an error wrapping multiple validation errors
// returned by EmptyReply.ValidateAll() if the designated constraints aren't met.
type EmptyReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptyReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptyReplyMultiError) AllErrors() []error { return m }

// EmptyReplyValidationError is the validation error returned by
// EmptyReply.Validate if the designated constraints aren't met.
type EmptyReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyReplyValidationError) ErrorName() string { return "EmptyReplyValidationError" }

// Error satisfies the builtin error interface
func (e EmptyReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmptyReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyReplyValidationError{}

// Validate checks the field values on CreateArticleRequest_Article with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateArticleRequest_Article) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateArticleRequest_Article with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateArticleRequest_ArticleMultiError, or nil if none found.
func (m *CreateArticleRequest_Article) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateArticleRequest_Article) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTitle()) < 1 {
		err := CreateArticleRequest_ArticleValidationError{
			field:  "Title",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetDescription() != "" {

	}

	if m.GetBody() != "" {

	}

	if len(m.GetTagList()) > 0 {

	}

	if len(errors) > 0 {
		return CreateArticleRequest_ArticleMultiError(errors)
	}

	return nil
}

// CreateArticleRequest_ArticleMultiError is an error wrapping multiple
// validation errors returned by CreateArticleRequest_Article.ValidateAll() if
// the designated constraints aren't met.
type CreateArticleRequest_ArticleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateArticleRequest_ArticleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateArticleRequest_ArticleMultiError) AllErrors() []error { return m }

// CreateArticleRequest_ArticleValidationError is the validation error returned
// by CreateArticleRequest_Article.Validate if the designated constraints
// aren't met.
type CreateArticleRequest_ArticleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateArticleRequest_ArticleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateArticleRequest_ArticleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateArticleRequest_ArticleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateArticleRequest_ArticleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateArticleRequest_ArticleValidationError) ErrorName() string {
	return "CreateArticleRequest_ArticleValidationError"
}

// Error satisfies the builtin error interface
func (e CreateArticleRequest_ArticleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateArticleRequest_Article.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateArticleRequest_ArticleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateArticleRequest_ArticleValidationError{}

// Validate checks the field values on SaveArticleRequest_Article with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SaveArticleRequest_Article) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SaveArticleRequest_Article with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SaveArticleRequest_ArticleMultiError, or nil if none found.
func (m *SaveArticleRequest_Article) ValidateAll() error {
	return m.validate(true)
}

func (m *SaveArticleRequest_Article) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetTitle() != "" {

		if utf8.RuneCountInString(m.GetTitle()) < 1 {
			err := SaveArticleRequest_ArticleValidationError{
				field:  "Title",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetDescription() != "" {

	}

	if m.GetBody() != "" {

	}

	if len(m.GetTagList()) > 0 {

	}

	if len(errors) > 0 {
		return SaveArticleRequest_ArticleMultiError(errors)
	}

	return nil
}

// SaveArticleRequest_ArticleMultiError is an error wrapping multiple
// validation errors returned by SaveArticleRequest_Article.ValidateAll() if
// the designated constraints aren't met.
type SaveArticleRequest_ArticleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SaveArticleRequest_ArticleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SaveArticleRequest_ArticleMultiError) AllErrors() []error { return m }

// SaveArticleRequest_ArticleValidationError is the validation error returned
// by SaveArticleRequest_Article.Validate if the designated constraints aren't met.
type SaveArticleRequest_ArticleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveArticleRequest_ArticleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveArticleRequest_ArticleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveArticleRequest_ArticleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveArticleRequest_ArticleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveArticleRequest_ArticleValidationError) ErrorName() string {
	return "SaveArticleRequest_ArticleValidationError"
}

// Error satisfies the builtin error interface
func (e SaveArticleRequest_ArticleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveArticleRequest_Article.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveArticleRequest_ArticleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveArticleRequest_ArticleValidationError{}
