// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package executeHistoryButton

import (
	"context"

	"github.com/erda-project/erda/apistructs"
	protocol "github.com/erda-project/erda/modules/openapi/component-protocol"
)

type ComponentAction struct{}

func getProps(visible bool) map[string]interface{} {
	return map[string]interface{}{
		"text":    "执行历史",
		"visible": visible,
	}
}

func (ca *ComponentAction) Render(ctx context.Context, c *apistructs.Component, scenario apistructs.ComponentProtocolScenario, event apistructs.ComponentEvent, gs *apistructs.GlobalStateData) error {
	//return json.Unmarshal([]byte(`{"text":"执行历史"}`), &c.Props)
	visible := true
	if _, ok := c.State["visible"]; ok {
		visible = c.State["visible"].(bool)
	}
	c.Props = getProps(visible)
	return nil
}

func RenderCreator() protocol.CompRender {
	return &ComponentAction{}
}