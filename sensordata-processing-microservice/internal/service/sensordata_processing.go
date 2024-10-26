package service

import (
	"context"
	v1 "sensordata-processing-microservice/api/sensor/v1"
	"sensordata-processing-microservice/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// SensorService 实现传感器处理的 gRPC 服务
type SensorService struct {
	v1.UnimplementedSensorServiceServer
	mgr *biz.SensorDataProcessingManager
	log *log.Helper
}

func NewSensorService(mgr *biz.SensorDataProcessingManager, logger log.Logger) *SensorService {
	return &SensorService{
		mgr: mgr,
		log: log.NewHelper(logger),
	}
}

// SetThreshold 设置传感器的报警阈值
func (s *SensorService) SetThreshold(ctx context.Context, req *v1.SetThresholdRequest) (*v1.SetThresholdResponse, error) {
	// DTO -> DO
	err := s.mgr.SetThreshold(ctx, req.SensorId, req.Threshold)
	if err != nil {
		s.log.Errorf("Error setting threshold: %v", err)
		return nil, err
	}
	return &v1.SetThresholdResponse{Success: true}, nil
}

// CheckAlarm 检查传感器的报警状态
func (s *SensorService) CheckAlarm(ctx context.Context, req *v1.AlarmRequest) (*v1.AlarmResponse, error) {
	// DTO -> DO
	sensorID := req.SensorId
	newValue := req.NewValue

	triggered, message, err := s.mgr.CheckAlarm(ctx, sensorID, newValue)
	if err != nil {
		s.log.Errorf("Error triggering alarm: %v", err)
		return nil, err
	}

	return &v1.AlarmResponse{AlarmTriggered: triggered, Message: message}, nil
}
