package enums

func (a *analytics) handleEmailBounce(em email) error {
	em.recipient.updateStatus(em.status)
	a.track(em.status)
	return nil
}
