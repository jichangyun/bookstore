package logic

import (
	"context"
	"time"
	"fmt"
	"net/http"
	"errors"

	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckLogic {
	return CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(w http.ResponseWriter, req types.CheckReq) (*types.CheckResp, error) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		l.Logger.Errorf("Writer is not a Flusher. w: %v.", w)
		return nil, errors.New("ResponseWriter unsupport Flusher!")
	}

	// 先返回头部信息
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.WriteHeader(http.StatusOK)
	flusher.Flush()
	
	fmt.Fprintf(w, "Start return value.\n")
	flusher.Flush()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(i) * time.Second)
		fmt.Fprintf(w, "Return value %d after sleep %d second.\n", i, i)
		flusher.Flush()
	}

	return &types.CheckResp{}, nil
}
