// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

// Encode encodes bool as json.
func (o OptBool) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Bool(bool(o.Value))
}

// Decode decodes bool from json.
func (o *OptBool) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptBool to nil")
	}
	o.Set = true
	v, err := d.Bool()
	if err != nil {
		return err
	}
	o.Value = bool(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptBool) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptBool) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes float32 as json.
func (o OptFloat32) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Float32(float32(o.Value))
}

// Decode decodes float32 from json.
func (o *OptFloat32) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptFloat32 to nil")
	}
	o.Set = true
	v, err := d.Float32()
	if err != nil {
		return err
	}
	o.Value = float32(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptFloat32) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptFloat32) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int64 as json.
func (o OptInt64) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int64(int64(o.Value))
}

// Decode decodes int64 from json.
func (o *OptInt64) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt64 to nil")
	}
	o.Set = true
	v, err := d.Int64()
	if err != nil {
		return err
	}
	o.Value = int64(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt64) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt64) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes User as json.
func (o OptUser) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes User from json.
func (o *OptUser) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptUser to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptUser) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptUser) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *User) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *User) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Name.Set {
			e.FieldStart("name")
			s.Name.Encode(e)
		}
	}
	{
		if s.Height.Set {
			e.FieldStart("height")
			s.Height.Encode(e)
		}
	}
	{
		if s.Male.Set {
			e.FieldStart("male")
			s.Male.Encode(e)
		}
	}
}

var jsonFieldsNameOfUser = [4]string{
	0: "id",
	1: "name",
	2: "height",
	3: "male",
}

// Decode decodes User from json.
func (s *User) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode User to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "name":
			if err := func() error {
				s.Name.Reset()
				if err := s.Name.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "height":
			if err := func() error {
				s.Height.Reset()
				if err := s.Height.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"height\"")
			}
		case "male":
			if err := func() error {
				s.Male.Reset()
				if err := s.Male.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"male\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode User")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *User) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *User) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
