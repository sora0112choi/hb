// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package retention

import (
	"github.com/goharbor/harbor/src/pkg/retention/policy"
	"github.com/goharbor/harbor/src/pkg/retention/policy/rule"
	"github.com/goharbor/harbor/src/pkg/retention/q"
	"time"
)

// Manager defines operations of managing policy
type Manager interface {
	// Create new policy and return uuid
	CreatePolicy(p *policy.Metadata) (string, error)
	// Update the existing policy
	// Full update
	UpdatePolicy(p *policy.Metadata) error
	// Delete the specified policy
	// No actual use so far
	DeletePolicy(ID int64) error
	// Get the specified policy
	GetPolicy(ID int64) (*policy.Metadata, error)
	// Create a new retention execution
	CreateExecution(execution *Execution) (string, error)
	// Update the specified execution
	UpdateExecution(execution *Execution) error
	// Get the specified execution
	GetExecution(eid int64) (*Execution, error)
	// List execution histories
	ListExecutions(query *q.Query) ([]*Execution, error)
	// Add new history
	AppendHistory(history *History) error
	// List all the histories marked by the specified execution
	ListHistories(executionID int64, query *q.Query) ([]*History, error)
}

type DefaultManager struct {
}

func (d *DefaultManager) CreatePolicy(p *policy.Metadata) (string, error) {
	panic("implement me")
}

func (d *DefaultManager) UpdatePolicy(p *policy.Metadata) error {
	panic("implement me")
}

func (d *DefaultManager) DeletePolicy(ID int64) error {
	panic("implement me")
}

func (d *DefaultManager) GetPolicy(ID int64) (*policy.Metadata, error) {
	return &policy.Metadata{
		ID:        1,
		Algorithm: "OR",
		Rules: []rule.Metadata{
			{
				ID:       1,
				Priority: 1,
				Template: "recentXdays",
				Parameters: rule.Parameters{
					"num": 10,
				},
				TagSelectors: []*rule.Selector{
					{
						Kind:       "label",
						Decoration: "with",
						Pattern:    "latest",
					},
					{
						Kind:       "regularExpression",
						Decoration: "matches",
						Pattern:    "release-[\\d\\.]+",
					},
				},
				ScopeSelectors: []*rule.Selector{
					{
						Kind:       "regularExpression",
						Decoration: "matches",
						Pattern:    ".+",
					},
				},
			},
		},
		Trigger: &policy.Trigger{
			Kind: "Schedule",
			Settings: map[string]interface{}{
				"cron": "* 22 11 * * *",
			},
		},
		Scope: &policy.Scope{
			Level:"project",
			Reference: 1,
		},
	}, nil
}

func (d *DefaultManager) CreateExecution(execution *Execution) (string, error) {
	panic("implement me")
}

func (d *DefaultManager) UpdateExecution(execution *Execution) error {
	panic("implement me")
}

func (d *DefaultManager) ListExecutions(query *q.Query) ([]*Execution, error) {
	return []*Execution{
		{
			ID:        1,
			PolicyID:  1,
			StartTime: time.Now().Add(-time.Minute),
			EndTime:   time.Now(),
			Status:    "Success",
		},
		{
			ID:        2,
			PolicyID:  1,
			StartTime: time.Now().Add(-time.Minute),
			EndTime:   time.Now(),
			Status:    "Failed",
		},
		{
			ID:        3,
			PolicyID:  1,
			StartTime: time.Now().Add(-time.Minute),
			EndTime:   time.Now(),
			Status:    "Running",
		},
	}, nil
}

func (d *DefaultManager) GetExecution(eid int64) (*Execution, error) {
	return &Execution{
		ID        :1,
		PolicyID  :1,
		StartTime :time.Now().Add(-time.Minute),
		EndTime   :time.Now(),
		Status    :"Success",
	}, nil
}

func (d *DefaultManager) ListHistories(executionID int64, query *q.Query) ([]*History, error) {
	panic("implement me")
}

func (d *DefaultManager) AppendHistory(history *History) error {
	panic("implement me")
}

func NewManager() Manager {
	return &DefaultManager{}
}
