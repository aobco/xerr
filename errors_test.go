package xerr

import (
	"errors"
	"fmt"
	"testing"
)

func thirdErr() error {
	return errors.New("origin 3rd source error: hey it is me")
}

func dependencyCall() error {
	return thirdErr()
}

func bizCodeWrap() error {
	cause := dependencyCall()
	wrapErr := BizWrap(cause, "FakeBizCode", "DebugMsg")
	return wrapErr
}

func callChainTrace1() error {
	wrap := bizCodeWrap()
	return Wrap(wrap, "call chain trace 1")
}

func callChainTrace2() error {
	wrap := callChainTrace1()
	return Wrap(wrap, "call chain trace 2")
}

var FakeBizRsp = BizResponseMap{
	BizCode("FakeBizCode"): BizMsg("This a fake biz response msg"),
}

func TestXer(t *testing.T) {
	if err := callChainTrace2(); err != nil {
		fmt.Printf("%+v\n", err)
		response := FakeBizRsp.Response(err)
		println("------------- response --------------")
		fmt.Printf("%+v\n", response)
	}
}
