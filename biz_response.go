package xerr

import (
	"github.com/aobco/log"
)

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
	println("as", as)
	if as {
		log.Errorf("%+v", errWrap)
		if bizMsg, ok := biz[errWrap.Code()]; ok {
			return &BizResponse{
				Code: errWrap.code,
				Msg:  bizMsg,
			}
		}
	}
	return BizResponseUnknown
}
