// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package conf

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

func easyjson1d75023bDecodeGithubComKingsgrouposArchivistLibGoOverwriteConf(in *jlexer.Lexer, out *KConfItem) {
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
		case "Kxa":
			out.Kxa = int64(in.Int64())
		case "Kxb":
			out.Kxb = int64(in.Int64())
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
func easyjson1d75023bEncodeGithubComKingsgrouposArchivistLibGoOverwriteConf(out *jwriter.Writer, in KConfItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Kxa\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Kxa))
	}
	{
		const prefix string = ",\"Kxb\":"
		out.RawString(prefix)
		out.Int64(int64(in.Kxb))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v KConfItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1d75023bEncodeGithubComKingsgrouposArchivistLibGoOverwriteConf(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v KConfItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1d75023bEncodeGithubComKingsgrouposArchivistLibGoOverwriteConf(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *KConfItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1d75023bDecodeGithubComKingsgrouposArchivistLibGoOverwriteConf(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *KConfItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1d75023bDecodeGithubComKingsgrouposArchivistLibGoOverwriteConf(l, v)
}
func easyjson1d75023bDecodeGithubComKingsgrouposArchivistLibGoOverwriteConf1(in *jlexer.Lexer, out *KConf) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
	} else {
		in.Delim('{')
		*out = make(KConf)
		for !in.IsDelim('}') {
			key := string(in.String())
			in.WantColon()
			var v1 *KConfItem
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(KConfItem)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			(*out)[key] = v1
			in.WantComma()
		}
		in.Delim('}')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1d75023bEncodeGithubComKingsgrouposArchivistLibGoOverwriteConf1(out *jwriter.Writer, in KConf) {
	if in == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v2First := true
		for v2Name, v2Value := range in {
			if v2First {
				v2First = false
			} else {
				out.RawByte(',')
			}
			out.String(string(v2Name))
			out.RawByte(':')
			if v2Value == nil {
				out.RawString("null")
			} else {
				(*v2Value).MarshalEasyJSON(out)
			}
		}
		out.RawByte('}')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v KConf) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1d75023bEncodeGithubComKingsgrouposArchivistLibGoOverwriteConf1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v KConf) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1d75023bEncodeGithubComKingsgrouposArchivistLibGoOverwriteConf1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *KConf) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1d75023bDecodeGithubComKingsgrouposArchivistLibGoOverwriteConf1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *KConf) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1d75023bDecodeGithubComKingsgrouposArchivistLibGoOverwriteConf1(l, v)
}
