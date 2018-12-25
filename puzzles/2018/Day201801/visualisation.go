package Day201801

func (td dayEntry) Visualise(updates chan<- []string) {
	defer close(updates)
	updates <- []string{"some", "stuff", "to", "display"}
}
