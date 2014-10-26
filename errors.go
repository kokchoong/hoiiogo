package hoiiogo

type HoiioError struct {
	Status string `json:"status"`
}

func (e HoiioError) Error() string {
  return "Hoiio API Error: " + e.Status
}