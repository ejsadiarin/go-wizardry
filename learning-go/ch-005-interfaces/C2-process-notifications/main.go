package interfaces

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	}

	return d.priorityLevel
}

func (g groupMessage) importance() int {
	return g.priorityLevel
}

func (s systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch v := n.(type) {
	case directMessage:
		return v.senderUsername, v.importance()
	case groupMessage:
		return v.groupName, v.importance()
	case systemAlert:
		return v.alertCode, v.importance()
	default:
		return "", 0
	}
}
