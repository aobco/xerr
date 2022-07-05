package xerr

type BizCode string

type BizMsg string

type BizResponse struct {
	Code BizCode
	Msg  BizMsg
}

var BizResponseUnknown = &BizResponse{
	Code: "unknown",
	Msg:  "unknown error",
}

type BizResponseMap map[BizCode]BizMsg

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
