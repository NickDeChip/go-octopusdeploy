package octopusdeploy

type triggerAction struct {
	Type ActionType `json:"ActionType"`

	Resource
}

func newTriggerAction(actionType ActionType) *triggerAction {
	return &triggerAction{
		Type:     actionType,
		Resource: *newResource(),
	}
}
