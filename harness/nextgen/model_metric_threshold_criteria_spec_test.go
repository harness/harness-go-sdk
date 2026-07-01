package nextgen

import (
	"encoding/json"
	"testing"
)

func TestMetricThresholdCriteriaSpec_ZeroValueSerialization(t *testing.T) {
	zero := float64(0)
	nonZero := float64(5.5)

	tests := []struct {
		name     string
		spec     MetricThresholdCriteriaSpec
		expected string
	}{
		{
			name:     "greaterThan zero should be included",
			spec:     MetricThresholdCriteriaSpec{GreaterThan: &zero},
			expected: `{"greaterThan":0}`,
		},
		{
			name:     "lessThan zero should be included",
			spec:     MetricThresholdCriteriaSpec{LessThan: &zero},
			expected: `{"lessThan":0}`,
		},
		{
			name:     "both zero should be included",
			spec:     MetricThresholdCriteriaSpec{LessThan: &zero, GreaterThan: &zero},
			expected: `{"lessThan":0,"greaterThan":0}`,
		},
		{
			name:     "non-zero values should be included",
			spec:     MetricThresholdCriteriaSpec{GreaterThan: &nonZero},
			expected: `{"greaterThan":5.5}`,
		},
		{
			name:     "nil fields should be omitted",
			spec:     MetricThresholdCriteriaSpec{},
			expected: `{}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.spec)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}
			if string(data) != tt.expected {
				t.Errorf("got %s, want %s", string(data), tt.expected)
			}
		})
	}
}

func TestMetricThresholdCriteriaSpec_ZeroValueDeserialization(t *testing.T) {
	input := `{"lessThan":0,"greaterThan":0}`

	var spec MetricThresholdCriteriaSpec
	err := json.Unmarshal([]byte(input), &spec)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if spec.LessThan == nil {
		t.Fatal("LessThan should not be nil")
	}
	if *spec.LessThan != 0 {
		t.Errorf("LessThan = %v, want 0", *spec.LessThan)
	}

	if spec.GreaterThan == nil {
		t.Fatal("GreaterThan should not be nil")
	}
	if *spec.GreaterThan != 0 {
		t.Errorf("GreaterThan = %v, want 0", *spec.GreaterThan)
	}
}

func TestMetricThresholdCriteria_FullRoundTrip(t *testing.T) {
	zero := float64(0)
	criteria := MetricThresholdCriteria{
		Type_: "Absolute",
		Spec: &MetricThresholdCriteriaSpec{
			GreaterThan: &zero,
		},
	}

	data, err := json.Marshal(criteria)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var decoded MetricThresholdCriteria
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if decoded.Spec == nil {
		t.Fatal("Spec should not be nil")
	}
	if decoded.Spec.GreaterThan == nil {
		t.Fatal("GreaterThan should not be nil after round-trip")
	}
	if *decoded.Spec.GreaterThan != 0 {
		t.Errorf("GreaterThan = %v, want 0", *decoded.Spec.GreaterThan)
	}
}
