package domain

// type Song ports.Song

type Song struct {
	ID          int
	Artist      string
	Title       string
	Lyrics      string
	ReleaseDate string
	Link        string
}

type UpdateSongData struct {
	Artist      *string
	Title       *string
	Lyrics      *string
	ReleaseDate *string
	Link        *string
}

func (s *Song) Update(d UpdateSongData) error {
	if d.Artist != nil {
		s.Artist = *d.Artist
	}

	if d.Title != nil {
		s.Title = *d.Title
	}

	if d.Lyrics != nil {
		s.Lyrics = *d.Lyrics
	}

	if d.ReleaseDate != nil {
		s.ReleaseDate = *d.ReleaseDate
	}

	if d.Link != nil {
		s.Link = *d.Link
	}

	return nil
}
