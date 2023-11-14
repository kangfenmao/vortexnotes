package indexer

type Indexer interface {
	ListAllNotes() []string
	IndexExist(path string) bool
	ParseNote(path string) string
	AddNoteToIndex(path string, note string) bool
}

func StartIndexer(indexer Indexer) {
	notes := indexer.ListAllNotes()
	for _, note := range notes {
		if !indexer.IndexExist(note) {
			parsedNote := indexer.ParseNote(note)
			indexer.AddNoteToIndex(note, parsedNote)
		}
	}
}
