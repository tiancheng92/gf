package gf

import (
	"testing"
	"time"
)

func TestPTSliceUnion(t *testing.T) {
	tests := []struct {
		name    string
		wantSum int64
	}{
		{
			name:    "success_union",
			wantSum: 46800,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPeriodTime()
			p.Append(time.Now(), time.Now().Add(2*time.Hour))
			p.Append(time.Now().Add(-1*time.Hour), time.Now().Add(1*time.Hour))
			p.Append(time.Now().Add(10*time.Hour), time.Now().Add(20*time.Hour))
			duration := p.Union().Duration()
			if duration != tt.wantSum {
				t.Errorf("Duration() = %v, want %v", duration, tt.wantSum)
			}
		})
	}
}

func TestPTSliceIntersect(t *testing.T) {
	tests := []struct {
		name    string
		wantSum int64
	}{
		{
			name:    "success_union",
			wantSum: 3600,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPeriodTime()
			p.Append(time.Now(), time.Now().Add(2*time.Hour))
			p.Append(time.Now().Add(-1*time.Hour), time.Now().Add(1*time.Hour))
			duration := p.Intersect().Duration()
			if duration != tt.wantSum {
				t.Errorf("Duration() = %v, want %v", duration, tt.wantSum)
			}
		})
	}
}
