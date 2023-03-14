// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package specification

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

func easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification(in *jlexer.Lexer, out *UpdateSpecificationInput) {
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
		case "item":
			if in.IsNull() {
				in.Skip()
				out.ItemID = nil
			} else {
				if out.ItemID == nil {
					out.ItemID = new(string)
				}
				*out.ItemID = string(in.String())
			}
		case "organization":
			if in.IsNull() {
				in.Skip()
				out.OrganizationID = nil
			} else {
				if out.OrganizationID == nil {
					out.OrganizationID = new(string)
				}
				*out.OrganizationID = string(in.String())
			}
		case "nametm":
			if in.IsNull() {
				in.Skip()
				out.NameTm = nil
			} else {
				if out.NameTm == nil {
					out.NameTm = new(string)
				}
				*out.NameTm = string(in.String())
			}
		case "nameru":
			if in.IsNull() {
				in.Skip()
				out.NameRu = nil
			} else {
				if out.NameRu == nil {
					out.NameRu = new(string)
				}
				*out.NameRu = string(in.String())
			}
		case "nametr":
			if in.IsNull() {
				in.Skip()
				out.NameTr = nil
			} else {
				if out.NameTr == nil {
					out.NameTr = new(string)
				}
				*out.NameTr = string(in.String())
			}
		case "nameen":
			if in.IsNull() {
				in.Skip()
				out.NameEn = nil
			} else {
				if out.NameEn == nil {
					out.NameEn = new(string)
				}
				*out.NameEn = string(in.String())
			}
		case "descriptiontm":
			if in.IsNull() {
				in.Skip()
				out.DescriptionTm = nil
			} else {
				if out.DescriptionTm == nil {
					out.DescriptionTm = new(string)
				}
				*out.DescriptionTm = string(in.String())
			}
		case "descriptionru":
			if in.IsNull() {
				in.Skip()
				out.DescriptionRu = nil
			} else {
				if out.DescriptionRu == nil {
					out.DescriptionRu = new(string)
				}
				*out.DescriptionRu = string(in.String())
			}
		case "descriptiontr":
			if in.IsNull() {
				in.Skip()
				out.DescriptionTr = nil
			} else {
				if out.DescriptionTr == nil {
					out.DescriptionTr = new(string)
				}
				*out.DescriptionTr = string(in.String())
			}
		case "descriptionen":
			if in.IsNull() {
				in.Skip()
				out.DescriptionEn = nil
			} else {
				if out.DescriptionEn == nil {
					out.DescriptionEn = new(string)
				}
				*out.DescriptionEn = string(in.String())
			}
		case "value":
			if in.IsNull() {
				in.Skip()
				out.Value = nil
			} else {
				if out.Value == nil {
					out.Value = new(string)
				}
				*out.Value = string(in.String())
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
func easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification(out *jwriter.Writer, in UpdateSpecificationInput) {
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
		const prefix string = ",\"item\":"
		out.RawString(prefix)
		if in.ItemID == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.ItemID))
		}
	}
	{
		const prefix string = ",\"organization\":"
		out.RawString(prefix)
		if in.OrganizationID == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.OrganizationID))
		}
	}
	{
		const prefix string = ",\"nametm\":"
		out.RawString(prefix)
		if in.NameTm == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.NameTm))
		}
	}
	{
		const prefix string = ",\"nameru\":"
		out.RawString(prefix)
		if in.NameRu == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.NameRu))
		}
	}
	{
		const prefix string = ",\"nametr\":"
		out.RawString(prefix)
		if in.NameTr == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.NameTr))
		}
	}
	{
		const prefix string = ",\"nameen\":"
		out.RawString(prefix)
		if in.NameEn == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.NameEn))
		}
	}
	{
		const prefix string = ",\"descriptiontm\":"
		out.RawString(prefix)
		if in.DescriptionTm == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.DescriptionTm))
		}
	}
	{
		const prefix string = ",\"descriptionru\":"
		out.RawString(prefix)
		if in.DescriptionRu == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.DescriptionRu))
		}
	}
	{
		const prefix string = ",\"descriptiontr\":"
		out.RawString(prefix)
		if in.DescriptionTr == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.DescriptionTr))
		}
	}
	{
		const prefix string = ",\"descriptionen\":"
		out.RawString(prefix)
		if in.DescriptionEn == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.DescriptionEn))
		}
	}
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		if in.Value == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Value))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateSpecificationInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateSpecificationInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateSpecificationInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateSpecificationInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification(l, v)
}
func easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification1(in *jlexer.Lexer, out *Specification) {
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
		case "item":
			out.ItemID = string(in.String())
		case "organization":
			out.OrganizationID = string(in.String())
		case "nametm":
			out.NameTm = string(in.String())
		case "nameru":
			out.NameRu = string(in.String())
		case "nametr":
			out.NameTr = string(in.String())
		case "nameen":
			out.NameEn = string(in.String())
		case "descriptiontm":
			out.DescriptionTm = string(in.String())
		case "descriptionru":
			out.DescriptionRu = string(in.String())
		case "descriptiontr":
			out.DescriptionTr = string(in.String())
		case "descriptionen":
			out.DescriptionEn = string(in.String())
		case "value":
			out.Value = string(in.String())
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
func easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification1(out *jwriter.Writer, in Specification) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"item\":"
		out.RawString(prefix)
		out.String(string(in.ItemID))
	}
	{
		const prefix string = ",\"organization\":"
		out.RawString(prefix)
		out.String(string(in.OrganizationID))
	}
	{
		const prefix string = ",\"nametm\":"
		out.RawString(prefix)
		out.String(string(in.NameTm))
	}
	{
		const prefix string = ",\"nameru\":"
		out.RawString(prefix)
		out.String(string(in.NameRu))
	}
	{
		const prefix string = ",\"nametr\":"
		out.RawString(prefix)
		out.String(string(in.NameTr))
	}
	{
		const prefix string = ",\"nameen\":"
		out.RawString(prefix)
		out.String(string(in.NameEn))
	}
	{
		const prefix string = ",\"descriptiontm\":"
		out.RawString(prefix)
		out.String(string(in.DescriptionTm))
	}
	{
		const prefix string = ",\"descriptionru\":"
		out.RawString(prefix)
		out.String(string(in.DescriptionRu))
	}
	{
		const prefix string = ",\"descriptiontr\":"
		out.RawString(prefix)
		out.String(string(in.DescriptionTr))
	}
	{
		const prefix string = ",\"descriptionen\":"
		out.RawString(prefix)
		out.String(string(in.DescriptionEn))
	}
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		out.String(string(in.Value))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Specification) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Specification) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Specification) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Specification) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification1(l, v)
}
func easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification2(in *jlexer.Lexer, out *ListSpecification) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(ListSpecification, 0, 0)
			} else {
				*out = ListSpecification{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Specification
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
func easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification2(out *jwriter.Writer, in ListSpecification) {
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
func (v ListSpecification) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ListSpecification) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ListSpecification) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ListSpecification) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification2(l, v)
}
func easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification3(in *jlexer.Lexer, out *CreateSpecificationInput) {
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
		case "item":
			out.ItemID = string(in.String())
		case "organization":
			out.OrganizationID = string(in.String())
		case "nametm":
			out.NameTm = string(in.String())
		case "nameru":
			out.NameRu = string(in.String())
		case "nametr":
			out.NameTr = string(in.String())
		case "nameen":
			out.NameEn = string(in.String())
		case "descriptiontm":
			out.DescriptionTm = string(in.String())
		case "descriptionru":
			out.DescriptionRu = string(in.String())
		case "descriptiontr":
			out.DescriptionTr = string(in.String())
		case "descriptionen":
			out.DescriptionEn = string(in.String())
		case "value":
			out.Value = string(in.String())
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
func easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification3(out *jwriter.Writer, in CreateSpecificationInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"item\":"
		out.RawString(prefix[1:])
		out.String(string(in.ItemID))
	}
	{
		const prefix string = ",\"organization\":"
		out.RawString(prefix)
		out.String(string(in.OrganizationID))
	}
	{
		const prefix string = ",\"nametm\":"
		out.RawString(prefix)
		out.String(string(in.NameTm))
	}
	{
		const prefix string = ",\"nameru\":"
		out.RawString(prefix)
		out.String(string(in.NameRu))
	}
	{
		const prefix string = ",\"nametr\":"
		out.RawString(prefix)
		out.String(string(in.NameTr))
	}
	{
		const prefix string = ",\"nameen\":"
		out.RawString(prefix)
		out.String(string(in.NameEn))
	}
	{
		const prefix string = ",\"descriptiontm\":"
		out.RawString(prefix)
		out.String(string(in.DescriptionTm))
	}
	{
		const prefix string = ",\"descriptionru\":"
		out.RawString(prefix)
		out.String(string(in.DescriptionRu))
	}
	{
		const prefix string = ",\"descriptiontr\":"
		out.RawString(prefix)
		out.String(string(in.DescriptionTr))
	}
	{
		const prefix string = ",\"descriptionen\":"
		out.RawString(prefix)
		out.String(string(in.DescriptionEn))
	}
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		out.String(string(in.Value))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateSpecificationInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateSpecificationInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateSpecificationInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateSpecificationInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainSpecification3(l, v)
}