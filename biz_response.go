package xerr

type BizResponse struct {
	Code string
	Msg  string
}

var BizResponseUnknown = &BizResponse{
	Code: "unknown",
	Msg:  "unknown error",
}

type BizResponseMap map[string]string

func (biz BizResponseMap) Response(err error) *BizResponse {
	if err == nil {
		return BizResponseUnknown
	}
	errWrap := new(ErrWrap)
	as := As(err, &errWrap)
	if as {
		if bizMsg, ok := biz[errWrap.Code()]; ok {
			return &BizResponse{
				Code: errWrap.code,
				Msg:  bizMsg,
			}
		}
	}
	return BizResponseUnknown
}
