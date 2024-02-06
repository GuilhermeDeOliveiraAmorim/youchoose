package entity

type MovieWriter struct {
	SharedEntity
	MovieID  string `json:"movie_id"`
	WriterID string `json:"writer_id"`
}

func NewMovieWriter(movieID, writerID string) *MovieWriter {
	return &MovieWriter{
		SharedEntity: *NewSharedEntity(),
		MovieID:      movieID,
		WriterID:     writerID,
	}
}

func (mw *MovieWriter) Equals(other *MovieWriter) bool {
	return mw.MovieID == other.MovieID && mw.WriterID == other.WriterID
}
