package unit

import "reflect"

// DecodeHookFunc implements the github.com/mitchellh/mapstructure DecodeHookFunc interface.
func DecodeHookFunc(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
	if from.Kind() != reflect.String {
		return data, nil
	}
	switch to {
	case reflect.TypeOf(Acceleration(0)):
		var a Acceleration
		if err := a.UnmarshalString(data.(string)); err != nil {
			return nil, err
		}
		return a, nil
	case reflect.TypeOf(Radians(0)):
		var r Radians
		if err := r.UnmarshalString(data.(string)); err != nil {
			return nil, err
		}
		return r, nil
	case reflect.TypeOf(Degree(0)):
		var d Degree
		if err := d.UnmarshalString(data.(string)); err != nil {
			return nil, err
		}
		return d, nil
	case reflect.TypeOf(AngularSpeed(0)):
		var a AngularSpeed
		if err := a.UnmarshalString(data.(string)); err != nil {
			return nil, err
		}
		return a, nil
	case reflect.TypeOf(Distance(0)):
		var d Distance
		if err := d.UnmarshalString(data.(string)); err != nil {
			return nil, err
		}
		return d, nil
	case reflect.TypeOf(Force(0)):
		var f Force
		if err := f.UnmarshalString(data.(string)); err != nil {
			return nil, err
		}
		return f, nil
	case reflect.TypeOf(Frequency(0)):
		var f Frequency
		if err := f.UnmarshalString(data.(string)); err != nil {
			return nil, err
		}
		return f, nil
	case reflect.TypeOf(Jerk(0)):
		var j Jerk
		if err := j.UnmarshalString(data.(string)); err != nil {
			return nil, err
		}
		return j, nil
	case reflect.TypeOf(Mass(0)):
		var m Mass
		if err := m.UnmarshalString(data.(string)); err != nil {
			return nil, err
		}
		return m, nil
	case reflect.TypeOf(Speed(0)):
		var s Speed
		if err := s.UnmarshalString(data.(string)); err != nil {
			return nil, err
		}
		return s, nil
	case reflect.TypeOf(Torque(0)):
		var t Torque
		if err := t.UnmarshalString(data.(string)); err != nil {
			return nil, err
		}
		return t, nil
	}
	return data, nil
}
