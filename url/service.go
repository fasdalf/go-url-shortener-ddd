package url

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/vokinneberg/go-url-shortener-ddd/domain"
)

type URLService struct {
	finder Finder
	saver  Saver
}

type Finder interface {
	Find(id string) (*domain.URL, error)
}
type Saver interface {
	Save(url *domain.URL) error
}

func NewURLService(finder Finder, saver Saver) *URLService {
	return &URLService{
		finder: finder,
		saver:  saver,
	}
}

func (h *URLService) Shorten(original string) (*domain.URL, error) {
	hash := sha1.New()
	hash.Write([]byte(original))
	short := hex.EncodeToString(hash.Sum(nil))[:8]
	url := domain.NewURL(short, original)
	err := h.saver.Save(url)
	return url, err
}

func (h *URLService) Find(id string) (*domain.URL, error) {
	return h.finder.Find(id)
}
