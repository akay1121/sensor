package biz

import (
	"context"
	"sensordata-processing-microservice/internal/biz/entity"
	"time"
)

type SensorDataProcessing = entity.SensorDataProcessing

type SensorDataProcessingRepository interface {
	GetSensorByID(ctx context.Context, id string) (*entity.SensorDataProcessing, error)
	UpdateSensorStatus(ctx context.Context, id string, status string) error
	//if already have the rule just delete the relative of SetSensorThreshold
	SetSensorThreshold(ctx context.Context, id string, threshold float64) error
	UpdateSensor(ctx context.Context, user *User) error
}

type SensorDataProcessingManager struct {
	repo SensorDataProcessingRepository
}

// init SensorProcessingManager
func NewSensorProcessingManager(repo SensorDataProcessingRepository) *SensorDataProcessingManager {
	return &SensorDataProcessingManager{repo: repo}
}

// check whether have alarm
func (m *SensorDataProcessingManager) CheckAlarm(ctx context.Context, sensorID string, newValue float64) (bool, string, error) {
	sensor, err := m.repo.GetSensorByID(ctx, sensorID)
	if err != nil {
		return false, "", err
	}

	//if you have rule replace it to the below
	//if sensor.Status == "alarm" {
	//	return true, "Sensor is in alarm state!", nil
	//}
	//return false, "Sensor is operating normally.", nil
	//check whether excedding the threshold
	if abs(newValue-sensor.PreviousValue) > sensor.Threshold {
		return true, "Sensor data abnormal!", nil
	}

	// update the value
	sensor.PreviousValue = newValue
	sensor.LastUpdate = time.Now()
	if err := m.repo.UpdateSensor(ctx, sensor); err != nil {
		return false, "", err
	}

	return false, "Sensor data normal.", nil
}

// abs 返回浮点数的绝对值
func abs(value float64) float64 {
	if value < 0 {
		return -value
	}
	return value
}

// interpolate data when the data missing
func (m *SensorDataProcessingManager) InterpolateData(ctx context.Context, sensorID string) (bool, string, error) {
	sensor, err := m.repo.GetSensorByID(ctx, sensorID)
	if err != nil {
		return false, "", err
	}

	//judge whether need to interpolate data
	if time.Since(sensor.LastUpdate) > 10*time.Minute {
		return true, "Interpolation data applied for missing entries.", nil
	}
	return false, "No interpolation needed.", nil
}
func (m *SensorDataProcessingManager) SetThreshold(ctx context.Context, sensorID string, threshold float64) error {
	return m.repo.SetSensorThreshold(ctx, sensorID, threshold)
}
