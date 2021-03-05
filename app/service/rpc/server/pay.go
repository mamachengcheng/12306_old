package server

import (
	"context"
	"errors"
	"github.com/mamachengcheng/12306/app/models"
	pb "github.com/mamachengcheng/12306/app/service/rpc/message"
	"github.com/mamachengcheng/12306/app/static"
	"github.com/mamachengcheng/12306/app/utils"
	"gorm.io/gorm"
)

type PayServer struct {
	pb.UnimplementedPayServer
}

func (s *OrderServer) PayMoney(ctx context.Context, in *pb.PayMoneyRequest) (*pb.PayMoneyReply, error) {

	// 获取Reds分布式锁
	res := utils.AcquireLockWithTimeout(in.Username)
	msg := ""

	if !res {
		// 更新订单状态
		err := utils.MysqlDB.Transaction(func(tx *gorm.DB) error {
			var order models.Order
			result := utils.MysqlDB.Model(&order).Where("id", in.OrderID).Update("order_status", static.PaidOrder)

			if result.RowsAffected == 0 {
				return errors.New("PaymentFailed")
			}

			return nil
		})

		// 如果支付成功，则释放Reds分布式锁
		if err == nil {
			utils.ReleaseLock(in.Username)
		} else {
			msg = "支付失败"
		}
	} else {
		msg = "无待支付订单"
	}

	return &pb.PayMoneyReply{
		Result: res,
		Msg:    msg,
	}, nil
}

func (s *OrderServer) RefundMoney(ctx context.Context, in *pb.RefundMoneyRequest) (*pb.RefundMoneyReply, error) {

	// 更新订单状态
	err := utils.MysqlDB.Transaction(func(tx *gorm.DB) error {
		var order models.Order
		result := utils.MysqlDB.Model(&order).Where("id = ? AND order_status = ?", in.OrderID, static.PaidOrder).Update("order_status", static.Refunded)

		if result.RowsAffected == 0 {
			return errors.New("PaymentFailed")
		}

		return nil
	})

	res := true
	msg := ""
	if err != nil {
		res = false
		msg = "退款失败"
	}

	return &pb.RefundMoneyReply{
		Result: res,
		Msg: msg,
	}, nil
}
