// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package order

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

func easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder(in *jlexer.Lexer, out *UpdateOrderInput) {
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
		case "table":
			if in.IsNull() {
				in.Skip()
				out.TableID = nil
			} else {
				if out.TableID == nil {
					out.TableID = new(string)
				}
				*out.TableID = string(in.String())
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
		case "orderitems":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				if out.Items == nil {
					out.Items = new([]OrderItem)
				}
				if in.IsNull() {
					in.Skip()
					*out.Items = nil
				} else {
					in.Delim('[')
					if *out.Items == nil {
						if !in.IsDelim(']') {
							*out.Items = make([]OrderItem, 0, 1)
						} else {
							*out.Items = []OrderItem{}
						}
					} else {
						*out.Items = (*out.Items)[:0]
					}
					for !in.IsDelim(']') {
						var v1 OrderItem
						(v1).UnmarshalEasyJSON(in)
						*out.Items = append(*out.Items, v1)
						in.WantComma()
					}
					in.Delim(']')
				}
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(int)
				}
				*out.Status = int(in.Int())
			}
		case "totalsum":
			if in.IsNull() {
				in.Skip()
				out.TotalSum = nil
			} else {
				if out.TotalSum == nil {
					out.TotalSum = new(float32)
				}
				*out.TotalSum = float32(in.Float32())
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
func easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder(out *jwriter.Writer, in UpdateOrderInput) {
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
		const prefix string = ",\"table\":"
		out.RawString(prefix)
		if in.TableID == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.TableID))
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
		const prefix string = ",\"orderitems\":"
		out.RawString(prefix)
		if in.Items == nil {
			out.RawString("null")
		} else {
			if *in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
				out.RawString("null")
			} else {
				out.RawByte('[')
				for v2, v3 := range *in.Items {
					if v2 > 0 {
						out.RawByte(',')
					}
					(v3).MarshalEasyJSON(out)
				}
				out.RawByte(']')
			}
		}
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		if in.Status == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.Status))
		}
	}
	{
		const prefix string = ",\"totalsum\":"
		out.RawString(prefix)
		if in.TotalSum == nil {
			out.RawString("null")
		} else {
			out.Float32(float32(*in.TotalSum))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateOrderInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateOrderInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateOrderInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateOrderInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder(l, v)
}
func easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder1(in *jlexer.Lexer, out *OrderItem) {
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
		case "order":
			out.OrderID = string(in.String())
		case "item":
			out.ItemID = string(in.String())
		case "quantity":
			out.Quantity = float32(in.Float32())
		case "unitprise":
			out.UnitPrice = float32(in.Float32())
		case "totalprice":
			out.TotalPrice = float32(in.Float32())
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
func easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder1(out *jwriter.Writer, in OrderItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"order\":"
		out.RawString(prefix)
		out.String(string(in.OrderID))
	}
	{
		const prefix string = ",\"item\":"
		out.RawString(prefix)
		out.String(string(in.ItemID))
	}
	{
		const prefix string = ",\"quantity\":"
		out.RawString(prefix)
		out.Float32(float32(in.Quantity))
	}
	{
		const prefix string = ",\"unitprise\":"
		out.RawString(prefix)
		out.Float32(float32(in.UnitPrice))
	}
	{
		const prefix string = ",\"totalprice\":"
		out.RawString(prefix)
		out.Float32(float32(in.TotalPrice))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v OrderItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OrderItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OrderItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OrderItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder1(l, v)
}
func easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder2(in *jlexer.Lexer, out *Order) {
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
		case "user":
			out.UserID = string(in.String())
		case "table":
			out.TableID = string(in.String())
		case "organization":
			out.OrganizationID = string(in.String())
		case "created":
			out.CreatedAt = string(in.String())
		case "updated":
			out.UpdatedAt = string(in.String())
		case "orderitems":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]OrderItem, 0, 1)
					} else {
						out.Items = []OrderItem{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v4 OrderItem
					(v4).UnmarshalEasyJSON(in)
					out.Items = append(out.Items, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "status":
			out.StatusID = int(in.Int())
		case "totalsum":
			out.TotalSum = float32(in.Float32())
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
func easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder2(out *jwriter.Writer, in Order) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.UserID != "" {
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		out.String(string(in.UserID))
	}
	if in.TableID != "" {
		const prefix string = ",\"table\":"
		out.RawString(prefix)
		out.String(string(in.TableID))
	}
	{
		const prefix string = ",\"organization\":"
		out.RawString(prefix)
		out.String(string(in.OrganizationID))
	}
	if in.CreatedAt != "" {
		const prefix string = ",\"created\":"
		out.RawString(prefix)
		out.String(string(in.CreatedAt))
	}
	if in.UpdatedAt != "" {
		const prefix string = ",\"updated\":"
		out.RawString(prefix)
		out.String(string(in.UpdatedAt))
	}
	if len(in.Items) != 0 {
		const prefix string = ",\"orderitems\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v5, v6 := range in.Items {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.StatusID != 0 {
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.Int(int(in.StatusID))
	}
	{
		const prefix string = ",\"totalsum\":"
		out.RawString(prefix)
		out.Float32(float32(in.TotalSum))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Order) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Order) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Order) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Order) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder2(l, v)
}
func easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder3(in *jlexer.Lexer, out *ListOrder) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(ListOrder, 0, 0)
			} else {
				*out = ListOrder{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v7 Order
			(v7).UnmarshalEasyJSON(in)
			*out = append(*out, v7)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder3(out *jwriter.Writer, in ListOrder) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v8, v9 := range in {
			if v8 > 0 {
				out.RawByte(',')
			}
			(v9).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v ListOrder) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ListOrder) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ListOrder) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ListOrder) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder3(l, v)
}
func easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder4(in *jlexer.Lexer, out *CreateOrderItemInput) {
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
		case "quantity":
			out.Quantity = float32(in.Float32())
		case "unitprise":
			out.UnitPrice = float32(in.Float32())
		case "totalprice":
			out.TotalPrice = float32(in.Float32())
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
func easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder4(out *jwriter.Writer, in CreateOrderItemInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"item\":"
		out.RawString(prefix[1:])
		out.String(string(in.ItemID))
	}
	{
		const prefix string = ",\"quantity\":"
		out.RawString(prefix)
		out.Float32(float32(in.Quantity))
	}
	{
		const prefix string = ",\"unitprise\":"
		out.RawString(prefix)
		out.Float32(float32(in.UnitPrice))
	}
	{
		const prefix string = ",\"totalprice\":"
		out.RawString(prefix)
		out.Float32(float32(in.TotalPrice))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateOrderItemInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateOrderItemInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateOrderItemInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateOrderItemInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder4(l, v)
}
func easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder5(in *jlexer.Lexer, out *CreateOrderInput) {
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
		case "user":
			out.UserID = string(in.String())
		case "table":
			out.TableID = string(in.String())
		case "organization":
			out.OrganizationID = string(in.String())
		case "orderitems":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]CreateOrderItemInput, 0, 2)
					} else {
						out.Items = []CreateOrderItemInput{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v10 CreateOrderItemInput
					(v10).UnmarshalEasyJSON(in)
					out.Items = append(out.Items, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "status":
			out.StatusID = int(in.Int())
		case "totalsum":
			out.TotalSum = float32(in.Float32())
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
func easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder5(out *jwriter.Writer, in CreateOrderInput) {
	out.RawByte('{')
	first := true
	_ = first
	if in.UserID != "" {
		const prefix string = ",\"user\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.UserID))
	}
	if in.TableID != "" {
		const prefix string = ",\"table\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.TableID))
	}
	{
		const prefix string = ",\"organization\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OrganizationID))
	}
	if len(in.Items) != 0 {
		const prefix string = ",\"orderitems\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v11, v12 := range in.Items {
				if v11 > 0 {
					out.RawByte(',')
				}
				(v12).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.StatusID != 0 {
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.Int(int(in.StatusID))
	}
	{
		const prefix string = ",\"totalsum\":"
		out.RawString(prefix)
		out.Float32(float32(in.TotalSum))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateOrderInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateOrderInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateOrderInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateOrderInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComEvgeniyDammermarketplaceApiInternalDomainOrder5(l, v)
}