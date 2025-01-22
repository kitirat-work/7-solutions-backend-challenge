package baconipsum

import "github.com/imroc/req/v3"

type PieFireDireBaconipsum interface {
	Get() (string, error)
}

type Baconipsum struct {
	client *req.Client
}

func NewBaconipsum(client *req.Client) PieFireDireBaconipsum {
	return &Baconipsum{client: client}
}

// Get implements PieFireDireBaconipsum.
func (b *Baconipsum) Get() (string, error) {
	resp, err := b.client.R().Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}
