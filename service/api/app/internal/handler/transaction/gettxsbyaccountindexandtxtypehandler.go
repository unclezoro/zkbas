package transaction

import (
	"net/http"

	"github.com/zecrey-labs/zecrey-legend/service/api/app/internal/logic/transaction"
	"github.com/zecrey-labs/zecrey-legend/service/api/app/internal/svc"
	"github.com/zecrey-labs/zecrey-legend/service/api/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTxsByAccountIndexAndTxTypeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqGetTxsByAccountIndexAndTxType
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := transaction.NewGetTxsByAccountIndexAndTxTypeLogic(r.Context(), svcCtx)
		resp, err := l.GetTxsByAccountIndexAndTxType(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
