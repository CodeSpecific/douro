package service

import (
	"github.com/CodeSpecific/douro/infra/algo"
)

type RedEnvelopAlgo func(count, amount int) int

type EnvelopeService interface {
	GetMonery(count, amount int) int
}

type envelopeService struct {
	RedEnvelopAlgo
}

func (s *envelopeService) GetMonery(count, amount int) int {
	return s.RedEnvelopAlgo(count, amount)
}

func NewEnvelopService() EnvelopeService {
	return &envelopeService{algo.RandomValue}
}


