package internal

type MovieWriter struct {
	SharedEntity
	MovieID  int `json:"movie_id"`
	WriterID int `json:"writer_id"`
}

func NewMovieWriter(movieID, writerID int) *MovieWriter {
	return &MovieWriter{
		SharedEntity: *NewSharedEntity(),
		MovieID:      movieID,
		WriterID:     writerID,
	}
}

func (mw *MovieWriter) Equals(other *MovieWriter) bool {
	return mw.MovieID == other.MovieID && mw.WriterID == other.WriterID
}
