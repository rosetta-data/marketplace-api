// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package organization

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonBc289ab0DecodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization(in *jlexer.Lexer, out *UpdateOrganizationInput) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			if in.IsNull() {
				in.Skip()
				out.ID = nil
			} else {
				if out.ID == nil {
					out.ID = new(string)
				}
				*out.ID = string(in.String())
			}
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		case "address":
			if in.IsNull() {
				in.Skip()
				out.Address = nil
			} else {
				if out.Address == nil {
					out.Address = new(string)
				}
				*out.Address = string(in.String())
			}
		case "phone":
			if in.IsNull() {
				in.Skip()
				out.Phone = nil
			} else {
				if out.Phone == nil {
					out.Phone = new(string)
				}
				*out.Phone = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBc289ab0EncodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization(out *jwriter.Writer, in UpdateOrganizationInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		if in.ID == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.ID))
		}
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	{
		const prefix string = ",\"address\":"
		out.RawString(prefix)
		if in.Address == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Address))
		}
	}
	{
		const prefix string = ",\"phone\":"
		out.RawString(prefix)
		if in.Phone == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Phone))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateOrganizationInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateOrganizationInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateOrganizationInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateOrganizationInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization(l, v)
}
func easyjsonBc289ab0DecodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization1(in *jlexer.Lexer, out *Organization) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "userid":
			out.UserID = string(in.String())
		case "address":
			out.Address = string(in.String())
		case "phone":
			out.Phone = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBc289ab0EncodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization1(out *jwriter.Writer, in Organization) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"userid\":"
		out.RawString(prefix)
		out.String(string(in.UserID))
	}
	{
		const prefix string = ",\"address\":"
		out.RawString(prefix)
		out.String(string(in.Address))
	}
	{
		const prefix string = ",\"phone\":"
		out.RawString(prefix)
		out.String(string(in.Phone))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Organization) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Organization) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Organization) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Organization) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization1(l, v)
}
func easyjsonBc289ab0DecodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization2(in *jlexer.Lexer, out *ListOrganization) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(ListOrganization, 0, 0)
			} else {
				*out = ListOrganization{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Organization
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBc289ab0EncodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization2(out *jwriter.Writer, in ListOrganization) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v ListOrganization) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ListOrganization) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ListOrganization) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ListOrganization) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization2(l, v)
}
func easyjsonBc289ab0DecodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization3(in *jlexer.Lexer, out *CreateOrganizationInput) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "userid":
			out.UserID = string(in.String())
		case "address":
			out.Address = string(in.String())
		case "phone":
			out.Phone = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBc289ab0EncodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization3(out *jwriter.Writer, in CreateOrganizationInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"userid\":"
		out.RawString(prefix)
		out.String(string(in.UserID))
	}
	{
		const prefix string = ",\"address\":"
		out.RawString(prefix)
		out.String(string(in.Address))
	}
	{
		const prefix string = ",\"phone\":"
		out.RawString(prefix)
		out.String(string(in.Phone))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateOrganizationInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateOrganizationInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateOrganizationInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateOrganizationInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammerEmenuApiInternalDomainOrganization3(l, v)
}
