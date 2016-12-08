package search

type HitType string

const (
  SegmentHitDB       HitType = "segment-db"
  SegmentHitHome     HitType = "segment-home"
  SegmentHitJson     HitType = "segment-json"
  SegmentHitScripted HitType = "segment-scripted"
)

type Hit struct {
	Id        int64    `json:"id"`
	Title     string   `json:"title"`
	Uri       string   `json:"uri"`
	Type      HitType  `json:"type"`
	Tags      []string `json:"tags"`
	IsStarred bool     `json:"isStarred"`
}

type HitList []*Hit

func (s HitList) Len() int           { return len(s) }
func (s HitList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s HitList) Less(i, j int) bool { return s[i].Title < s[j].Title }

type SegmentQuery struct {
  Title     string
  Tags      []string
  OrgId     int64
  UserId    int64
  Limit     int
  IsStarred bool

  Result HitList
}

type FindPersistedSegmentsQuery struct {
  Title     string
  OrgId     int64
  UserId    int64
  IsStarred bool

  Result HitList
}
