package filter

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"sync/atomic"

	"github.com/Liyanbing/filter/cache"
)

type Manger interface {
	Execute(ctx context.Context, data interface{}) (ret interface{}, err error)
	Refresh(ctx context.Context, jsonStr string) error
}

type Reporter interface {
	Report(ctx context.Context, data interface{}, filterID string)
}

type ReportFunc func(ctx context.Context, data interface{}, filterID string)

func (rf ReportFunc) Report(ctx context.Context, data interface{}, filterID string) {
	rf(ctx, data, filterID)
}

// -------------
type Config struct {
	Filters map[string]SingleConfig `json:"filters"`
	Version string                  `json:"version"`
}

type SingleConfig struct {
	FilterData []interface{} `json:"filter_data"`
	Weight     int64         `json:"weight"`
	Priority   int64         `json:"priority"`
}

type container struct {
	filterGroup *GroupFilter
	reporter    Reporter
}

func NewFilter(ctx context.Context, jsonStr string, reporter Reporter) (Manger, error) {
	con, err := newContainerWithJSON(ctx, jsonStr)
	if err != nil {
		return nil, err
	}
	con.reporter = reporter

	filterValue := atomic.Value{}
	filterValue.Store(con)

	return &manager{
		filterValue: filterValue,
	}, nil
}

func newContainerWithJSON(ctx context.Context, jsonStr string) (*container, error) {
	var cnf Config
	err := json.NewDecoder(strings.NewReader(jsonStr)).Decode(&cnf)
	if err != nil {
		return nil, err
	}

	filterGroup, err := NewGroupFilterWithConfig(ctx, &cnf)
	if err != nil {
		return nil, err
	}

	return &container{
		filterGroup: filterGroup,
		reporter:    nil,
	}, nil
}

// -----------
type manager struct {
	filterValue atomic.Value
}

func (s *manager) Execute(ctx context.Context, data interface{}) (interface{}, error) {
	filterValue, ok := s.filterValue.Load().(*container)
	if !ok {
		return nil, errors.New("invalid manager")
	}

	if data == nil {
		data = make(map[string]interface{})
	}

	_, filterID := filterValue.filterGroup.Run(ctx, data, cache.NewCache())
	if filterValue.reporter != nil {
		filterValue.reporter.Report(ctx, data, filterID)
	}
	return data, nil
}

func (s *manager) Refresh(ctx context.Context, jsonStr string) error {
	newCon, err := newContainerWithJSON(ctx, jsonStr)
	if err != nil {
		return err
	}

	if container, ok := s.filterValue.Load().(*container); ok {
		newCon.reporter = container.reporter
	}

	s.filterValue.Store(newCon)
	return nil
}

func CheckConfig(ctx context.Context, jsonStr string) error {
	_, err := newContainerWithJSON(ctx, jsonStr)
	return err
}
