package server

import (
	"fmt"
	"net/http"

	"github.com/Oloruntobi1/walletservice/constants"
	"github.com/Oloruntobi1/walletservice/db"
	"github.com/Oloruntobi1/walletservice/log"

	//"github.com/Oloruntobi1/walletservice/monnify"
	//"github.com/Oloruntobi1/walletservice/safegate"
	"github.com/Oloruntobi1/walletservice/wallet"
	//"github.com/Oloruntobi1/walletservice/walletafrica"
	_ "github.com/Oloruntobi1/payload"
	"github.com/Oloruntobi1/tobibank"
	"github.com/gin-gonic/gin"
)

type PayoutRequest struct {
	Name string `json:"name" binding:"required" example:"tobi"`
}

// HandlePayout godoc
// @Summary handle payout
// @Description handle payout with name request
// @Tags payout
// @ID payout-by-name
// @Accept  json
// @Produce  json
// @Param name body CreateWalletRequest false "person name"
// @Success 200 {object} payload.WalletPayoutResponse
// @Router /payout [post]
func (s *Server) HandlePayout(ctx *gin.Context) {
	var req *PayoutRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		log.LogRequest(req, err)
		return
	}

	// if req.Name != "tobi"{
	// 	log.LogRequest(req, "not error")
	// 	return
	// }

	log.LogRequest(req, "not error")

	res, err := getPayoutProvider()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		log.LogRequest(req, err)
		return
	}
	rem, err := wallet.PayoutUser(res)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		log.LogRequest(req, err)
		return
	}
	ctx.JSON(http.StatusOK, rem)

}

func getPayoutProvider() (wallet.Wallet, error) {

	res := db.CheckPayoutProvider()
	if res == constants.TOBIBANK {
		return tobibank.WireUpTobiBank(), nil
	}
	return nil, fmt.Errorf("no provider available")
}
