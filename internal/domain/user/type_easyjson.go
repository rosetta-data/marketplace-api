// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package user

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

func easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser(in *jlexer.Lexer, out *User) {
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
		case "phone":
			out.Phone = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "firstname":
			out.FirstName = string(in.String())
		case "lastname":
			out.LastName = string(in.String())
		case "role":
			out.RoleName = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "roleid":
			out.RoleID = int(in.Int())
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
func easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"phone\":"
		out.RawString(prefix)
		out.String(string(in.Phone))
	}
	if in.Password != "" {
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	{
		const prefix string = ",\"firstname\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"lastname\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	if in.RoleName != "" {
		const prefix string = ",\"role\":"
		out.RawString(prefix)
		out.String(string(in.RoleName))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	if in.RoleID != 0 {
		const prefix string = ",\"roleid\":"
		out.RawString(prefix)
		out.Int(int(in.RoleID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser(l, v)
}
func easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser1(in *jlexer.Lexer, out *UpdateUserInput) {
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
		case "firstname":
			if in.IsNull() {
				in.Skip()
				out.FirstName = nil
			} else {
				if out.FirstName == nil {
					out.FirstName = new(string)
				}
				*out.FirstName = string(in.String())
			}
		case "lastname":
			if in.IsNull() {
				in.Skip()
				out.LastName = nil
			} else {
				if out.LastName == nil {
					out.LastName = new(string)
				}
				*out.LastName = string(in.String())
			}
		case "password":
			if in.IsNull() {
				in.Skip()
				out.Password = nil
			} else {
				if out.Password == nil {
					out.Password = new(string)
				}
				*out.Password = string(in.String())
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
func easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser1(out *jwriter.Writer, in UpdateUserInput) {
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
		const prefix string = ",\"firstname\":"
		out.RawString(prefix)
		if in.FirstName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.FirstName))
		}
	}
	{
		const prefix string = ",\"lastname\":"
		out.RawString(prefix)
		if in.LastName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.LastName))
		}
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		if in.Password == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Password))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateUserInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateUserInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateUserInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateUserInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser1(l, v)
}
func easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser2(in *jlexer.Lexer, out *SignInInput) {
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
		case "phone":
			out.Phone = string(in.String())
		case "password":
			out.Password = string(in.String())
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
func easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser2(out *jwriter.Writer, in SignInInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"phone\":"
		out.RawString(prefix[1:])
		out.String(string(in.Phone))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SignInInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SignInInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SignInInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SignInInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser2(l, v)
}
func easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser3(in *jlexer.Lexer, out *ListUser) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(ListUser, 0, 0)
			} else {
				*out = ListUser{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 User
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
func easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser3(out *jwriter.Writer, in ListUser) {
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
func (v ListUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ListUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ListUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ListUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser3(l, v)
}
func easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser4(in *jlexer.Lexer, out *CreateUserInput) {
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
		case "phone":
			out.Phone = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "firstname":
			out.FirstName = string(in.String())
		case "lastname":
			out.LastName = string(in.String())
		case "roleid":
			out.RoleID = int(in.Int())
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
func easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser4(out *jwriter.Writer, in CreateUserInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"phone\":"
		out.RawString(prefix[1:])
		out.String(string(in.Phone))
	}
	if in.Password != "" {
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	{
		const prefix string = ",\"firstname\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"lastname\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	if in.RoleID != 0 {
		const prefix string = ",\"roleid\":"
		out.RawString(prefix)
		out.Int(int(in.RoleID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateUserInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateUserInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateUserInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateUserInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainUser4(l, v)
}