package service

import (
	"context"
	"piefiredire/baconipsum"
	"strings"
)

type PieFireDireService interface {
	BeefSummary(ctx context.Context) (BeefSummary, error)
}

type Service struct {
	baconipsumService baconipsum.PieFireDireBaconipsum
}

func NewService(
	baconipsumService baconipsum.PieFireDireBaconipsum,
) PieFireDireService {
	return &Service{
		baconipsumService: baconipsumService,
	}
}

type BeefSummary struct {
	Beef map[string]int `json:"beef"`
}

// BeefSummary implements PieFireDireService.
func (s *Service) BeefSummary(ctx context.Context) (BeefSummary, error) {
	message, err := s.baconipsumService.Get()
	if err != nil {
		return BeefSummary{}, err
	}

	result := BeefSummary{
		Beef: make(map[string]int),
	}
	words := strings.FieldsFunc(message, func(r rune) bool {
		return !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r == '-')
	})
	for _, w := range words {
		beefType := strings.ToLower(w)
		result.Beef[beefType]++
	}

	return result, nil
}
