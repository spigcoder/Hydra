package log

type ProduceRequest struct {
	Record Record `json:"record"`
}

type ProduceResponse struct {
	Offset uint64 `json:"offset"`
}

type ConsumeRequest struct {
	Offset uint64 `form:"offset"`
}

type ConsumeResponse struct {
	Record Record `json:"record"`
}
