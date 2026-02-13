package data

import (
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

// func (m Movie) MarshalJSON() ([]byte, error) {
// 	var runtime string

// 	if m.Runtime != 0 {
// 		runtime = fmt.Sprintf("%d mins", m.Runtime)
// 	}

// 	aux := struct {
// 		ID      int64    `json:"id"`
// 		Title   string   `json:"title"`
// 		Year    int32    `json:"year,omitempty"`
// 		Runtime string   `json:"runtime,omitempty"`
// 		Genres  []string `json:"genres,omitempty"`
// 		Version int32    `json:"version"`
// 	}{
// 		ID:      m.ID,
// 		Title:   m.Title,
// 		Year:    m.Year,
// 		Runtime: runtime,
// 		Genres:  m.Genres,
// 		Version: m.Version,
// 	}

// 	return json.Marshal(aux)
// }
